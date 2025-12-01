package models

import (
	"time"
)

// User represents a poker tracker user
type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"-"`
	PlayerName   string    `json:"player_name"`
	WinaStatus   string    `json:"wina_status"`
	HHDirectory  string    `json:"hh_directory"`
	CreatedAt    time.Time `json:"created_at"`
	DevMode      bool      `json:"dev_mode"`
}

// Tournament represents a single Expresso tournament
type Tournament struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	BuyinCents    int       `json:"buyin_cents"`
	RakeCents     int       `json:"rake_cents"`
	Multiplier    float64   `json:"multiplier"`
	Prize1Cents   int       `json:"prize1_cents"`
	Prize2Cents   int       `json:"prize2_cents"`
	Prize3Cents   int       `json:"prize3_cents"`
	HeroRank      int       `json:"hero_rank"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
	HandsCount    int       `json:"hands_count"`
	NetResultCents int      `json:"net_result_cents"`
	EVCents       float64   `json:"ev_cents"`
	ChipsEV       float64   `json:"chips_ev"`
	ImportFile    string    `json:"import_file"`
	TournamentID  string    `json:"tournament_id"`
	CreatedAt     time.Time `json:"created_at"`
}

// Hand represents a single poker hand
type Hand struct {
	ID              int       `json:"id"`
	TournamentID    int       `json:"tournament_id"`
	HandNumber      int64     `json:"hand_number"`
	HeroStackStart  int       `json:"hero_stack_start"`
	HeroStackEnd    int       `json:"hero_stack_end"`
	ChipsWon        int       `json:"chips_won"`
	EVChips         float64   `json:"ev_chips"`
	ActionJSON      string    `json:"action_json"`
	CreatedAt       time.Time `json:"created_at"`
}

// MultiplierExpected represents the expected multiplier probabilities for each buy-in
type MultiplierExpected struct {
	ID          int     `json:"id"`
	BuyinCents  int     `json:"buyin_cents"`
	Multiplier  float64 `json:"multiplier"`
	Probability float64 `json:"probability"`
	Prize1Cents int     `json:"prize1_cents"`
	Prize2Cents int     `json:"prize2_cents"`
	Prize3Cents int     `json:"prize3_cents"`
}

// RakebackStatus represents Winamax status levels and rakeback info
type RakebackStatus struct {
	ID            int     `json:"id"`
	StatusName    string  `json:"status_name"`
	MilesRequired int     `json:"miles_required"`
	RakeEquiv     int     `json:"rake_equiv"`
	AnnualBonus   int     `json:"annual_bonus"`
	RakebackPct   float64 `json:"rakeback_pct"`
}

// Opponent represents a poker opponent (future use)
type Opponent struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	PlayerName string    `json:"player_name"`
	VPIP       float64   `json:"vpip"`
	PFR        float64   `json:"pfr"`
	Steal      float64   `json:"steal"`
	LimpPct    float64   `json:"limp_pct"`
	CheckRaise float64   `json:"check_raise"`
	HandsCount int       `json:"hands_count"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// OpponentHand represents a hand involving an opponent (future use)
type OpponentHand struct {
	ID         int       `json:"id"`
	OpponentID int       `json:"opponent_id"`
	HandID     int       `json:"hand_id"`
	ActionJSON string    `json:"action_json"`
	CreatedAt  time.Time `json:"created_at"`
}

// Stats represents aggregated statistics for a user
type Stats struct {
	TotalTournaments int     `json:"total_tournaments"`
	TotalHands       int     `json:"total_hands"`
	NetResultCents   int     `json:"net_result_cents"`
	EVCents          float64 `json:"ev_cents"`
	ChipsEV          float64 `json:"chips_ev"`
	TotalRakeCents   int     `json:"total_rake_cents"`
	TotalRakeback    int     `json:"total_rakeback_cents"`
	Winrate          float64 `json:"winrate"`
	EVROI            float64 `json:"ev_roi"`
	NetROI           float64 `json:"net_roi"`
	AvgChipsPerGame  float64 `json:"avg_chips_per_game"`
}

// TournamentSummary represents parsed tournament summary data
type TournamentSummary struct {
	TournamentID     string
	PlayerName       string
	BuyinCents       int
	RakeCents        int
	RegisteredPlayers int
	PrizepoolCents   int
	StartTime        time.Time
	Duration         string
	FinishPosition   int
	PrizeWonCents    int
}

// HandHistory represents a parsed hand
type HandHistory struct {
	HandID         string
	TournamentName string
	BuyinCents     int
	RakeCents      int
	Level          int
	SmallBlind     int
	BigBlind       int
	Timestamp      time.Time
	TableName      string
	ButtonSeat     int
	Players        []Player
	HeroName       string
	HeroCards      string
	Actions        []Action
	Board          string
	Pot            int
	Rake           int
}

// Player represents a player at the table
type Player struct {
	Seat  int
	Name  string
	Stack int
}

// Action represents a poker action
type Action struct {
	Street string
	Player string
	Action string
	Amount int
}
