package parser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/mathieudottir/poker_tracker/backend/internal/models"
)

var (
	// Tournament summary patterns
	summaryTournamentIDRegex = regexp.MustCompile(`Tournament summary : Expresso Nitro\((\d+)\)`)
	summaryPlayerRegex       = regexp.MustCompile(`Player : (.+)`)
	summaryBuyinRegex        = regexp.MustCompile(`Buy-In : ([\d.]+)€ \+ ([\d.]+)€`)
	summaryPrizepoolRegex    = regexp.MustCompile(`Prizepool : ([\d.]+)€`)
	summaryStartTimeRegex    = regexp.MustCompile(`Tournament started (.+)`)
	summaryFinishRegex       = regexp.MustCompile(`You finished in (\d+)(?:st|nd|rd|th) place`)
	summaryPrizeWonRegex     = regexp.MustCompile(`You won ([\d.]+)€`)

	// Hand history patterns
	handHeaderRegex     = regexp.MustCompile(`Winamax Poker - Tournament "([^"]+)" buyIn: ([\d.]+)€ \+ ([\d.]+)€ level: (\d+) - HandId: #([0-9-]+) - Holdem no limit \((\d+)/(\d+)\) - (.+)`)
	tableRegex          = regexp.MustCompile(`Table: '([^']+)' 3-max \(real money\) Seat #(\d+) is the button`)
	playerRegex         = regexp.MustCompile(`Seat (\d+): (.+) \((\d+)\)`)
	dealtToRegex        = regexp.MustCompile(`Dealt to ([^ ]+) \[(.+)\]`)
	actionRegex         = regexp.MustCompile(`^([^ ]+) (folds|checks|calls|bets|raises|shows|collected)`)
	potRegex            = regexp.MustCompile(`Total pot (\d+)`)
)

// ParseHandHistoryFile parses a complete hand history file
func ParseHandHistoryFile(filepath string) ([]models.HandHistory, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var hands []models.HandHistory
	var currentHand *models.HandHistory
	var currentStreet string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			if currentHand != nil {
				hands = append(hands, *currentHand)
				currentHand = nil
			}
			continue
		}

		// Check for hand header
		if matches := handHeaderRegex.FindStringSubmatch(line); matches != nil {
			if currentHand != nil {
				hands = append(hands, *currentHand)
			}

			buyin, _ := strconv.ParseFloat(matches[2], 64)
			rake, _ := strconv.ParseFloat(matches[3], 64)
			level, _ := strconv.Atoi(matches[4])
			sb, _ := strconv.Atoi(matches[6])
			bb, _ := strconv.Atoi(matches[7])
			timestamp := parseWinamaxTime(matches[8])

			currentHand = &models.HandHistory{
				TournamentName: matches[1],
				BuyinCents:     int(buyin * 100),
				RakeCents:      int(rake * 100),
				Level:          level,
				SmallBlind:     sb,
				BigBlind:       bb,
				Timestamp:      timestamp,
				HandID:         matches[5],
				Players:        []models.Player{},
				Actions:        []models.Action{},
			}
			currentStreet = "PREHAND"
			continue
		}

		if currentHand == nil {
			continue
		}

		// Parse table info
		if matches := tableRegex.FindStringSubmatch(line); matches != nil {
			currentHand.TableName = matches[1]
			currentHand.ButtonSeat, _ = strconv.Atoi(matches[2])
			continue
		}

		// Parse players
		if matches := playerRegex.FindStringSubmatch(line); matches != nil {
			seat, _ := strconv.Atoi(matches[1])
			stack, _ := strconv.Atoi(matches[3])
			currentHand.Players = append(currentHand.Players, models.Player{
				Seat:  seat,
				Name:  matches[2],
				Stack: stack,
			})
			continue
		}

		// Parse dealt cards
		if matches := dealtToRegex.FindStringSubmatch(line); matches != nil {
			currentHand.HeroName = matches[1]
			currentHand.HeroCards = matches[2]
			continue
		}

		// Detect streets
		if strings.Contains(line, "*** PRE-FLOP ***") {
			currentStreet = "PREFLOP"
			continue
		}
		if strings.Contains(line, "*** FLOP ***") {
			currentStreet = "FLOP"
			if matches := regexp.MustCompile(`\[(.+)\]`).FindStringSubmatch(line); matches != nil {
				currentHand.Board = matches[1]
			}
			continue
		}
		if strings.Contains(line, "*** TURN ***") {
			currentStreet = "TURN"
			if matches := regexp.MustCompile(`\[.+\]\[(.+)\]`).FindStringSubmatch(line); matches != nil {
				currentHand.Board += " " + matches[1]
			}
			continue
		}
		if strings.Contains(line, "*** RIVER ***") {
			currentStreet = "RIVER"
			if matches := regexp.MustCompile(`\[.+\]\[(.+)\]`).FindStringSubmatch(line); matches != nil {
				currentHand.Board += " " + matches[1]
			}
			continue
		}
		if strings.Contains(line, "*** SHOW DOWN ***") {
			currentStreet = "SHOWDOWN"
			continue
		}
		if strings.Contains(line, "*** SUMMARY ***") {
			currentStreet = "SUMMARY"
			continue
		}

		// Parse pot
		if matches := potRegex.FindStringSubmatch(line); matches != nil {
			currentHand.Pot, _ = strconv.Atoi(matches[1])
			continue
		}

		// Parse actions
		if currentStreet != "PREHAND" && currentStreet != "SUMMARY" {
			action := parseAction(line, currentStreet)
			if action != nil {
				currentHand.Actions = append(currentHand.Actions, *action)
			}
		}
	}

	if currentHand != nil {
		hands = append(hands, *currentHand)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return hands, nil
}

// ParseSummaryFile parses a tournament summary file
func ParseSummaryFile(filepath string) (*models.TournamentSummary, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	summary := &models.TournamentSummary{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if matches := summaryTournamentIDRegex.FindStringSubmatch(line); matches != nil {
			summary.TournamentID = matches[1]
		}
		if matches := summaryPlayerRegex.FindStringSubmatch(line); matches != nil {
			summary.PlayerName = matches[1]
		}
		if matches := summaryBuyinRegex.FindStringSubmatch(line); matches != nil {
			buyin, _ := strconv.ParseFloat(matches[1], 64)
			rake, _ := strconv.ParseFloat(matches[2], 64)
			summary.BuyinCents = int(buyin * 100)
			summary.RakeCents = int(rake * 100)
		}
		if matches := summaryPrizepoolRegex.FindStringSubmatch(line); matches != nil {
			prizepool, _ := strconv.ParseFloat(matches[1], 64)
			summary.PrizepoolCents = int(prizepool * 100)
		}
		if matches := summaryStartTimeRegex.FindStringSubmatch(line); matches != nil {
			summary.StartTime = parseWinamaxTime(matches[1])
		}
		if matches := summaryFinishRegex.FindStringSubmatch(line); matches != nil {
			summary.FinishPosition, _ = strconv.Atoi(matches[1])
		}
		if matches := summaryPrizeWonRegex.FindStringSubmatch(line); matches != nil {
			prize, _ := strconv.ParseFloat(matches[1], 64)
			summary.PrizeWonCents = int(prize * 100)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return summary, nil
}

// parseAction parses an action line
func parseAction(line, street string) *models.Action {
	// Parse blinds
	if strings.Contains(line, "posts small blind") {
		parts := strings.Fields(line)
		if len(parts) >= 4 {
			amount, _ := strconv.Atoi(parts[3])
			return &models.Action{
				Street: "BLINDS",
				Player: parts[0],
				Action: "posts small blind",
				Amount: amount,
			}
		}
	}
	if strings.Contains(line, "posts big blind") {
		parts := strings.Fields(line)
		if len(parts) >= 4 {
			amount, _ := strconv.Atoi(parts[3])
			return &models.Action{
				Street: "BLINDS",
				Player: parts[0],
				Action: "posts big blind",
				Amount: amount,
			}
		}
	}

	// Parse regular actions
	parts := strings.Fields(line)
	if len(parts) < 2 {
		return nil
	}

	action := &models.Action{
		Street: street,
		Player: parts[0],
		Action: parts[1],
	}

	// Extract amounts
	if strings.Contains(line, "raises") && len(parts) >= 4 {
		// Format: "player raises X to Y"
		if len(parts) >= 4 {
			amount, _ := strconv.Atoi(parts[3])
			action.Amount = amount
		}
	} else if strings.Contains(line, "bets") || strings.Contains(line, "calls") {
		if len(parts) >= 3 {
			amount, _ := strconv.Atoi(parts[2])
			action.Amount = amount
		}
	} else if strings.Contains(line, "collected") {
		if len(parts) >= 3 {
			amount, _ := strconv.Atoi(parts[2])
			action.Amount = amount
		}
	}

	return action
}

// parseWinamaxTime parses Winamax timestamp format
func parseWinamaxTime(timeStr string) time.Time {
	// Format: "2025/11/30 10:26:11 UTC"
	t, err := time.Parse("2006/01/02 15:04:05 MST", timeStr)
	if err != nil {
		return time.Now()
	}
	return t
}

// GetTournamentIDFromFilename extracts tournament ID from filename
func GetTournamentIDFromFilename(filename string) string {
	re := regexp.MustCompile(`\((\d+)\)`)
	matches := re.FindStringSubmatch(filename)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}
