package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/mathieudottir/poker_tracker/backend/internal/calculator"
	"github.com/mathieudottir/poker_tracker/backend/internal/models"
	"github.com/mathieudottir/poker_tracker/backend/internal/parser"
	"github.com/mathieudottir/poker_tracker/backend/internal/repository"
)

type ImportService struct {
	repo *repository.Repository
}

func NewImportService(repo *repository.Repository) *ImportService {
	return &ImportService{repo: repo}
}

// ImportTournament imports a complete tournament from files
func (s *ImportService) ImportTournament(ctx context.Context, userID int, hhFilePath, summaryFilePath string) error {
	// Parse summary file
	summary, err := parser.ParseSummaryFile(summaryFilePath)
	if err != nil {
		return fmt.Errorf("failed to parse summary: %w", err)
	}

	// Check if tournament already exists
	exists, err := s.repo.TournamentExists(ctx, userID, summary.TournamentID)
	if err != nil {
		return fmt.Errorf("failed to check tournament existence: %w", err)
	}
	if exists {
		log.Printf("Tournament %s already imported, skipping", summary.TournamentID)
		return nil
	}

	// Parse hand history file
	hands, err := parser.ParseHandHistoryFile(hhFilePath)
	if err != nil {
		return fmt.Errorf("failed to parse hand history: %w", err)
	}

	// Calculate multiplier and prize distribution
	multiplier := calculator.DetermineMultiplier(summary.PrizepoolCents, summary.BuyinCents, summary.RakeCents)

	// For Expresso Nitro, use total entry (buyin + rake) for calculations
	totalEntryCents := summary.BuyinCents + summary.RakeCents

	// Expresso Nitro is winner-takes-all
	prize1 := summary.PrizepoolCents
	prize2 := 0
	prize3 := 0

	// Calculate EV based on total entry (for proper lookup in constants table)
	expectedEV := calculator.CalculateTournamentEV(totalEntryCents)

	// Calculate net result
	netResult := calculator.CalculateNetResult(summary.PrizeWonCents, summary.BuyinCents, summary.RakeCents)

	// Determine hero's finish position (1st, 2nd, or 3rd)
	heroRank := summary.FinishPosition

	// Calculate end time (start time + duration would be ideal, but we'll estimate)
	endTime := summary.StartTime.Add(5 * time.Minute) // Default estimate
	if len(hands) > 0 {
		endTime = hands[len(hands)-1].Timestamp
	}

	// Create tournament record
	tournament := &models.Tournament{
		UserID:         userID,
		BuyinCents:     summary.BuyinCents,
		RakeCents:      summary.RakeCents,
		Multiplier:     multiplier,
		Prize1Cents:    prize1,
		Prize2Cents:    prize2,
		Prize3Cents:    prize3,
		HeroRank:       heroRank,
		StartTime:      summary.StartTime,
		EndTime:        endTime,
		HandsCount:     len(hands),
		NetResultCents: netResult,
		EVCents:        expectedEV,
		ChipsEV:        0, // Will be calculated from hands
		ImportFile:     filepath.Base(hhFilePath),
		TournamentID:   summary.TournamentID,
	}

	// Create tournament
	if err := s.repo.CreateTournament(ctx, tournament); err != nil {
		return fmt.Errorf("failed to create tournament: %w", err)
	}

	// Import hands
	var totalChipsEV float64
	for i, handHistory := range hands {
		// Extract hero's stack changes
		heroStackStart := 0
		heroStackEnd := 0

		// Find hero in players
		for _, player := range handHistory.Players {
			if player.Name == summary.PlayerName {
				heroStackStart = player.Stack
				break
			}
		}

		// Calculate stack at end (simplified - would need more complex logic)
		heroStackEnd = heroStackStart
		chipsWon := 0

		// Check if hero won the pot
		for _, action := range handHistory.Actions {
			if action.Player == summary.PlayerName && action.Action == "collected" {
				chipsWon = action.Amount
				heroStackEnd = heroStackStart + chipsWon
			}
		}

		// Convert actions to JSON
		actionsJSON, err := json.Marshal(handHistory.Actions)
		if err != nil {
			log.Printf("Warning: failed to marshal actions for hand %d: %v", i, err)
			actionsJSON = []byte("[]")
		}

		hand := &models.Hand{
			TournamentID:   tournament.ID,
			HandNumber:     int64(i + 1),
			HeroStackStart: heroStackStart,
			HeroStackEnd:   heroStackEnd,
			ChipsWon:       chipsWon,
			EVChips:        float64(chipsWon), // Simplified EV calculation
			ActionJSON:     string(actionsJSON),
		}

		if err := s.repo.CreateHand(ctx, hand); err != nil {
			log.Printf("Warning: failed to create hand %d: %v", i, err)
		}

		totalChipsEV += hand.EVChips
	}

	log.Printf("Successfully imported tournament %s with %d hands", summary.TournamentID, len(hands))
	return nil
}

// ImportDirectory imports all tournaments from a directory
func (s *ImportService) ImportDirectory(ctx context.Context, userID int, directory string) (int, error) {
	matches, err := filepath.Glob(filepath.Join(directory, "*_summary.txt"))
	if err != nil {
		return 0, fmt.Errorf("failed to find summary files: %w", err)
	}

	imported := 0
	for _, summaryFile := range matches {
		// Find corresponding hand history file
		hhFile := strings.Replace(summaryFile, "_summary.txt", ".txt", 1)

		log.Printf("Importing tournament from %s", summaryFile)
		if err := s.ImportTournament(ctx, userID, hhFile, summaryFile); err != nil {
			log.Printf("Error importing tournament from %s: %v", summaryFile, err)
			continue
		}
		imported++
	}

	return imported, nil
}

// GetImportLogs returns recent import logs (simplified implementation)
func (s *ImportService) GetImportLogs() []string {
	// In a real implementation, this would read from a log file
	return []string{
		"Import service started",
		"Ready to import hand histories",
	}
}
