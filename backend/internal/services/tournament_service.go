package services

import (
	"context"
	"time"

	"github.com/mathieudottir/poker_tracker/backend/internal/calculator"
	"github.com/mathieudottir/poker_tracker/backend/internal/models"
	"github.com/mathieudottir/poker_tracker/backend/internal/repository"
)

type TournamentService struct {
	repo *repository.Repository
}

func NewTournamentService(repo *repository.Repository) *TournamentService {
	return &TournamentService{repo: repo}
}

// GetTournaments returns all tournaments for a user
func (s *TournamentService) GetTournaments(ctx context.Context, userID int) ([]models.Tournament, error) {
	return s.repo.GetTournamentsByUserID(ctx, userID)
}

// GetTournament returns a specific tournament
func (s *TournamentService) GetTournament(ctx context.Context, id int) (*models.Tournament, error) {
	return s.repo.GetTournamentByID(ctx, id)
}

// GetHands returns all hands for a tournament
func (s *TournamentService) GetHands(ctx context.Context, tournamentID int) ([]models.Hand, error) {
	return s.repo.GetHandsByTournamentID(ctx, tournamentID)
}

// GetStats returns aggregated statistics for a user
func (s *TournamentService) GetStats(ctx context.Context, userID int, userStatus string) (*models.Stats, error) {
	tournaments, err := s.repo.GetTournamentsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	stats := calculator.CalculateStats(tournaments, userStatus)
	return &stats, nil
}

// GetMultiplierDistribution returns multiplier distribution
func (s *TournamentService) GetMultiplierDistribution(ctx context.Context, userID int) (map[float64]int, error) {
	return s.repo.GetMultiplierDistribution(ctx, userID)
}

// GetResultsByBuyin returns results grouped by buy-in
func (s *TournamentService) GetResultsByBuyin(ctx context.Context, userID int) (map[int]models.Stats, error) {
	return s.repo.GetResultsByBuyin(ctx, userID)
}

// GetBankrollHistory returns bankroll evolution over time
func (s *TournamentService) GetBankrollHistory(ctx context.Context, userID int) ([]BankrollPoint, error) {
	tournaments, err := s.repo.GetTournamentsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var history []BankrollPoint
	var cumulative int

	for _, t := range tournaments {
		cumulative += t.NetResultCents
		history = append(history, BankrollPoint{
			Time:    t.EndTime,
			Balance: cumulative,
		})
	}

	return history, nil
}

// GetEVHistory returns EV evolution over time
func (s *TournamentService) GetEVHistory(ctx context.Context, userID int) ([]EVPoint, error) {
	tournaments, err := s.repo.GetTournamentsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var history []EVPoint
	var cumulativeEV float64

	for _, t := range tournaments {
		cumulativeEV += t.EVCents
		history = append(history, EVPoint{
			Time: t.EndTime,
			EV:   cumulativeEV,
		})
	}

	return history, nil
}

type BankrollPoint struct {
	Time    time.Time `json:"time"`
	Balance int       `json:"balance"`
}

type EVPoint struct {
	Time time.Time `json:"time"`
	EV   float64   `json:"ev"`
}
