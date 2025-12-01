-- Initial database schema for Poker Tracker

-- Users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    player_name TEXT NOT NULL,
    wina_status TEXT DEFAULT 'Aluminium',
    hh_directory TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    dev_mode BOOLEAN DEFAULT FALSE
);

-- Tournaments table
CREATE TABLE IF NOT EXISTS tournaments (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    buyin_cents INTEGER NOT NULL,
    rake_cents INTEGER NOT NULL,
    multiplier FLOAT NOT NULL,
    prize1_cents INTEGER NOT NULL,
    prize2_cents INTEGER NOT NULL,
    prize3_cents INTEGER NOT NULL,
    hero_rank INTEGER NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    hands_count INTEGER DEFAULT 0,
    net_result_cents INTEGER NOT NULL,
    ev_cents FLOAT NOT NULL,
    chips_ev FLOAT DEFAULT 0,
    import_file TEXT,
    tournament_id TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, tournament_id)
);

-- Hands table
CREATE TABLE IF NOT EXISTS hands (
    id SERIAL PRIMARY KEY,
    tournament_id INTEGER NOT NULL REFERENCES tournaments(id) ON DELETE CASCADE,
    hand_number BIGINT NOT NULL,
    hero_stack_start INTEGER NOT NULL,
    hero_stack_end INTEGER NOT NULL,
    chips_won INTEGER NOT NULL,
    ev_chips FLOAT DEFAULT 0,
    action_json JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(tournament_id, hand_number)
);

-- Multipliers expected (hardcoded Winamax data)
CREATE TABLE IF NOT EXISTS multipliers_expected (
    id SERIAL PRIMARY KEY,
    buyin_cents INTEGER NOT NULL,
    multiplier FLOAT NOT NULL,
    probability FLOAT NOT NULL,
    prize1_cents INTEGER NOT NULL,
    prize2_cents INTEGER NOT NULL,
    prize3_cents INTEGER NOT NULL,
    UNIQUE(buyin_cents, multiplier)
);

-- Rakeback status (hardcoded Winamax data)
CREATE TABLE IF NOT EXISTS rakeback_status (
    id SERIAL PRIMARY KEY,
    status_name TEXT UNIQUE NOT NULL,
    miles_required INTEGER NOT NULL,
    rake_equiv INTEGER NOT NULL,
    annual_bonus INTEGER NOT NULL,
    rakeback_pct FLOAT NOT NULL
);

-- Opponents table (future use)
CREATE TABLE IF NOT EXISTS opponents (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    player_name TEXT NOT NULL,
    vpip FLOAT DEFAULT 0,
    pfr FLOAT DEFAULT 0,
    steal FLOAT DEFAULT 0,
    limp_pct FLOAT DEFAULT 0,
    check_raise FLOAT DEFAULT 0,
    hands_count INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, player_name)
);

-- Opponent hands table (future use)
CREATE TABLE IF NOT EXISTS opponent_hands (
    id SERIAL PRIMARY KEY,
    opponent_id INTEGER NOT NULL REFERENCES opponents(id) ON DELETE CASCADE,
    hand_id INTEGER NOT NULL REFERENCES hands(id) ON DELETE CASCADE,
    action_json JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_tournaments_user_id ON tournaments(user_id);
CREATE INDEX IF NOT EXISTS idx_tournaments_start_time ON tournaments(start_time);
CREATE INDEX IF NOT EXISTS idx_tournaments_buyin ON tournaments(buyin_cents);
CREATE INDEX IF NOT EXISTS idx_hands_tournament_id ON hands(tournament_id);
CREATE INDEX IF NOT EXISTS idx_opponents_user_id ON opponents(user_id);
CREATE INDEX IF NOT EXISTS idx_opponent_hands_opponent_id ON opponent_hands(opponent_id);
CREATE INDEX IF NOT EXISTS idx_opponent_hands_hand_id ON opponent_hands(hand_id);
