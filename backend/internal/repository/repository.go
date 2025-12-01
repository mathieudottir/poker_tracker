package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/mathieudottir/poker_tracker/backend/internal/models"
	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(dsn string) (*Repository, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Repository{db: db}, nil
}

func (r *Repository) Close() error {
	return r.db.Close()
}

// User operations

func (r *Repository) CreateUser(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (username, password_hash, player_name, wina_status, hh_directory, dev_mode)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at
	`
	return r.db.QueryRowContext(ctx, query,
		user.Username, user.PasswordHash, user.PlayerName, user.WinaStatus, user.HHDirectory, user.DevMode,
	).Scan(&user.ID, &user.CreatedAt)
}

func (r *Repository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	query := `
		SELECT id, username, password_hash, player_name, wina_status, hh_directory, created_at, dev_mode
		FROM users
		WHERE username = $1
	`
	user := &models.User{}
	err := r.db.QueryRowContext(ctx, query, username).Scan(
		&user.ID, &user.Username, &user.PasswordHash, &user.PlayerName,
		&user.WinaStatus, &user.HHDirectory, &user.CreatedAt, &user.DevMode,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repository) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	query := `
		SELECT id, username, password_hash, player_name, wina_status, hh_directory, created_at, dev_mode
		FROM users
		WHERE id = $1
	`
	user := &models.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID, &user.Username, &user.PasswordHash, &user.PlayerName,
		&user.WinaStatus, &user.HHDirectory, &user.CreatedAt, &user.DevMode,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repository) UpdateUser(ctx context.Context, user *models.User) error {
	query := `
		UPDATE users
		SET player_name = $1, wina_status = $2, hh_directory = $3, dev_mode = $4
		WHERE id = $5
	`
	_, err := r.db.ExecContext(ctx, query,
		user.PlayerName, user.WinaStatus, user.HHDirectory, user.DevMode, user.ID,
	)
	return err
}

func (r *Repository) DeleteUserData(ctx context.Context, userID int) error {
	// Cascading deletes will handle tournaments, hands, opponents, etc.
	query := `DELETE FROM tournaments WHERE user_id = $1`
	_, err := r.db.ExecContext(ctx, query, userID)
	return err
}

// Tournament operations

func (r *Repository) CreateTournament(ctx context.Context, tournament *models.Tournament) error {
	query := `
		INSERT INTO tournaments (
			user_id, buyin_cents, rake_cents, multiplier, prize1_cents, prize2_cents, prize3_cents,
			hero_rank, start_time, end_time, hands_count, net_result_cents, ev_cents, chips_ev,
			import_file, tournament_id
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
		ON CONFLICT (user_id, tournament_id) DO NOTHING
		RETURNING id, created_at
	`
	return r.db.QueryRowContext(ctx, query,
		tournament.UserID, tournament.BuyinCents, tournament.RakeCents, tournament.Multiplier,
		tournament.Prize1Cents, tournament.Prize2Cents, tournament.Prize3Cents, tournament.HeroRank,
		tournament.StartTime, tournament.EndTime, tournament.HandsCount, tournament.NetResultCents,
		tournament.EVCents, tournament.ChipsEV, tournament.ImportFile, tournament.TournamentID,
	).Scan(&tournament.ID, &tournament.CreatedAt)
}

func (r *Repository) GetTournamentsByUserID(ctx context.Context, userID int) ([]models.Tournament, error) {
	query := `
		SELECT id, user_id, buyin_cents, rake_cents, multiplier, prize1_cents, prize2_cents, prize3_cents,
			hero_rank, start_time, end_time, hands_count, net_result_cents, ev_cents, chips_ev,
			import_file, tournament_id, created_at
		FROM tournaments
		WHERE user_id = $1
		ORDER BY start_time DESC
	`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tournaments []models.Tournament
	for rows.Next() {
		var t models.Tournament
		if err := rows.Scan(
			&t.ID, &t.UserID, &t.BuyinCents, &t.RakeCents, &t.Multiplier,
			&t.Prize1Cents, &t.Prize2Cents, &t.Prize3Cents, &t.HeroRank,
			&t.StartTime, &t.EndTime, &t.HandsCount, &t.NetResultCents,
			&t.EVCents, &t.ChipsEV, &t.ImportFile, &t.TournamentID, &t.CreatedAt,
		); err != nil {
			return nil, err
		}
		tournaments = append(tournaments, t)
	}
	return tournaments, rows.Err()
}

func (r *Repository) GetTournamentByID(ctx context.Context, id int) (*models.Tournament, error) {
	query := `
		SELECT id, user_id, buyin_cents, rake_cents, multiplier, prize1_cents, prize2_cents, prize3_cents,
			hero_rank, start_time, end_time, hands_count, net_result_cents, ev_cents, chips_ev,
			import_file, tournament_id, created_at
		FROM tournaments
		WHERE id = $1
	`
	t := &models.Tournament{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&t.ID, &t.UserID, &t.BuyinCents, &t.RakeCents, &t.Multiplier,
		&t.Prize1Cents, &t.Prize2Cents, &t.Prize3Cents, &t.HeroRank,
		&t.StartTime, &t.EndTime, &t.HandsCount, &t.NetResultCents,
		&t.EVCents, &t.ChipsEV, &t.ImportFile, &t.TournamentID, &t.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (r *Repository) TournamentExists(ctx context.Context, userID int, tournamentID string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM tournaments WHERE user_id = $1 AND tournament_id = $2)`
	var exists bool
	err := r.db.QueryRowContext(ctx, query, userID, tournamentID).Scan(&exists)
	return exists, err
}

// Hand operations

func (r *Repository) CreateHand(ctx context.Context, hand *models.Hand) error {
	query := `
		INSERT INTO hands (
			tournament_id, hand_number, hero_stack_start, hero_stack_end,
			chips_won, ev_chips, action_json
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (tournament_id, hand_number) DO NOTHING
		RETURNING id, created_at
	`
	return r.db.QueryRowContext(ctx, query,
		hand.TournamentID, hand.HandNumber, hand.HeroStackStart, hand.HeroStackEnd,
		hand.ChipsWon, hand.EVChips, hand.ActionJSON,
	).Scan(&hand.ID, &hand.CreatedAt)
}

func (r *Repository) GetHandsByTournamentID(ctx context.Context, tournamentID int) ([]models.Hand, error) {
	query := `
		SELECT id, tournament_id, hand_number, hero_stack_start, hero_stack_end,
			chips_won, ev_chips, action_json, created_at
		FROM hands
		WHERE tournament_id = $1
		ORDER BY hand_number ASC
	`
	rows, err := r.db.QueryContext(ctx, query, tournamentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hands []models.Hand
	for rows.Next() {
		var h models.Hand
		if err := rows.Scan(
			&h.ID, &h.TournamentID, &h.HandNumber, &h.HeroStackStart, &h.HeroStackEnd,
			&h.ChipsWon, &h.EVChips, &h.ActionJSON, &h.CreatedAt,
		); err != nil {
			return nil, err
		}
		hands = append(hands, h)
	}
	return hands, rows.Err()
}

// Opponent operations (future use)

func (r *Repository) CreateOrUpdateOpponent(ctx context.Context, opponent *models.Opponent) error {
	query := `
		INSERT INTO opponents (user_id, player_name, vpip, pfr, steal, limp_pct, check_raise, hands_count, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (user_id, player_name)
		DO UPDATE SET
			vpip = $3, pfr = $4, steal = $5, limp_pct = $6, check_raise = $7,
			hands_count = $8, updated_at = $9
		RETURNING id
	`
	return r.db.QueryRowContext(ctx, query,
		opponent.UserID, opponent.PlayerName, opponent.VPIP, opponent.PFR,
		opponent.Steal, opponent.LimpPct, opponent.CheckRaise, opponent.HandsCount, time.Now(),
	).Scan(&opponent.ID)
}

func (r *Repository) GetOpponentsByUserID(ctx context.Context, userID int) ([]models.Opponent, error) {
	query := `
		SELECT id, user_id, player_name, vpip, pfr, steal, limp_pct, check_raise,
			hands_count, created_at, updated_at
		FROM opponents
		WHERE user_id = $1
		ORDER BY hands_count DESC
	`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var opponents []models.Opponent
	for rows.Next() {
		var o models.Opponent
		if err := rows.Scan(
			&o.ID, &o.UserID, &o.PlayerName, &o.VPIP, &o.PFR,
			&o.Steal, &o.LimpPct, &o.CheckRaise, &o.HandsCount,
			&o.CreatedAt, &o.UpdatedAt,
		); err != nil {
			return nil, err
		}
		opponents = append(opponents, o)
	}
	return opponents, rows.Err()
}

// Stats operations

func (r *Repository) GetMultiplierDistribution(ctx context.Context, userID int) (map[float64]int, error) {
	query := `
		SELECT multiplier, COUNT(*) as count
		FROM tournaments
		WHERE user_id = $1
		GROUP BY multiplier
		ORDER BY multiplier
	`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	distribution := make(map[float64]int)
	for rows.Next() {
		var multiplier float64
		var count int
		if err := rows.Scan(&multiplier, &count); err != nil {
			return nil, err
		}
		distribution[multiplier] = count
	}
	return distribution, rows.Err()
}

func (r *Repository) GetResultsByBuyin(ctx context.Context, userID int) (map[int]models.Stats, error) {
	query := `
		SELECT
			buyin_cents,
			COUNT(*) as total_tournaments,
			SUM(hands_count) as total_hands,
			SUM(net_result_cents) as net_result,
			SUM(ev_cents) as ev_cents,
			SUM(chips_ev) as chips_ev,
			SUM(rake_cents) as total_rake
		FROM tournaments
		WHERE user_id = $1
		GROUP BY buyin_cents
		ORDER BY buyin_cents
	`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make(map[int]models.Stats)
	for rows.Next() {
		var buyinCents int
		var stats models.Stats
		if err := rows.Scan(
			&buyinCents,
			&stats.TotalTournaments,
			&stats.TotalHands,
			&stats.NetResultCents,
			&stats.EVCents,
			&stats.ChipsEV,
			&stats.TotalRakeCents,
		); err != nil {
			return nil, err
		}
		results[buyinCents] = stats
	}
	return results, rows.Err()
}

// Utility to convert actions to JSON
func ActionsToJSON(actions []models.Action) (string, error) {
	data, err := json.Marshal(actions)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
