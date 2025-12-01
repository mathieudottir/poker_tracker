-- Seed Winamax multipliers and rakeback data

-- Insert multipliers for €0.25
INSERT INTO multipliers_expected (buyin_cents, multiplier, probability, prize1_cents, prize2_cents, prize3_cents) VALUES
(25, 2, 0.80, 38, 13, 0),
(25, 4, 0.10, 75, 25, 0),
(25, 10, 0.05, 188, 63, 0),
(25, 100, 0.04, 1875, 625, 0),
(25, 1000, 0.01, 18750, 6250, 0)
ON CONFLICT (buyin_cents, multiplier) DO NOTHING;

-- Insert multipliers for €0.50
INSERT INTO multipliers_expected (buyin_cents, multiplier, probability, prize1_cents, prize2_cents, prize3_cents) VALUES
(50, 2, 0.80, 75, 25, 0),
(50, 4, 0.10, 150, 50, 0),
(50, 10, 0.05, 375, 125, 0),
(50, 100, 0.04, 3750, 1250, 0),
(50, 1000, 0.01, 37500, 12500, 0)
ON CONFLICT (buyin_cents, multiplier) DO NOTHING;

-- Insert multipliers for €1
INSERT INTO multipliers_expected (buyin_cents, multiplier, probability, prize1_cents, prize2_cents, prize3_cents) VALUES
(100, 2, 0.80, 150, 50, 0),
(100, 4, 0.10, 300, 100, 0),
(100, 10, 0.05, 750, 250, 0),
(100, 100, 0.04, 7500, 2500, 0),
(100, 1000, 0.01, 75000, 25000, 0)
ON CONFLICT (buyin_cents, multiplier) DO NOTHING;

-- Insert multipliers for €2
INSERT INTO multipliers_expected (buyin_cents, multiplier, probability, prize1_cents, prize2_cents, prize3_cents) VALUES
(200, 2, 0.80, 300, 100, 0),
(200, 4, 0.10, 600, 200, 0),
(200, 10, 0.05, 1500, 500, 0),
(200, 100, 0.04, 15000, 5000, 0),
(200, 1000, 0.01, 150000, 50000, 0)
ON CONFLICT (buyin_cents, multiplier) DO NOTHING;

-- Insert multipliers for €5
INSERT INTO multipliers_expected (buyin_cents, multiplier, probability, prize1_cents, prize2_cents, prize3_cents) VALUES
(500, 2, 0.80, 750, 250, 0),
(500, 4, 0.10, 1500, 500, 0),
(500, 10, 0.05, 3750, 1250, 0),
(500, 100, 0.04, 37500, 12500, 0),
(500, 1000, 0.01, 375000, 125000, 0)
ON CONFLICT (buyin_cents, multiplier) DO NOTHING;

-- Insert multipliers for €10
INSERT INTO multipliers_expected (buyin_cents, multiplier, probability, prize1_cents, prize2_cents, prize3_cents) VALUES
(1000, 2, 0.80, 1500, 500, 0),
(1000, 4, 0.10, 3000, 1000, 0),
(1000, 10, 0.05, 7500, 2500, 0),
(1000, 100, 0.04, 75000, 25000, 0),
(1000, 1000, 0.01, 750000, 250000, 0)
ON CONFLICT (buyin_cents, multiplier) DO NOTHING;

-- Insert multipliers for €25
INSERT INTO multipliers_expected (buyin_cents, multiplier, probability, prize1_cents, prize2_cents, prize3_cents) VALUES
(2500, 2, 0.80, 3750, 1250, 0),
(2500, 4, 0.10, 7500, 2500, 0),
(2500, 10, 0.05, 18750, 6250, 0),
(2500, 100, 0.04, 187500, 62500, 0),
(2500, 1000, 0.01, 1875000, 625000, 0)
ON CONFLICT (buyin_cents, multiplier) DO NOTHING;

-- Insert multipliers for €50
INSERT INTO multipliers_expected (buyin_cents, multiplier, probability, prize1_cents, prize2_cents, prize3_cents) VALUES
(5000, 2, 0.80, 7500, 2500, 0),
(5000, 4, 0.10, 15000, 5000, 0),
(5000, 10, 0.05, 37500, 12500, 0),
(5000, 100, 0.04, 375000, 125000, 0),
(5000, 1000, 0.01, 3750000, 1250000, 0)
ON CONFLICT (buyin_cents, multiplier) DO NOTHING;

-- Insert multipliers for €100
INSERT INTO multipliers_expected (buyin_cents, multiplier, probability, prize1_cents, prize2_cents, prize3_cents) VALUES
(10000, 2, 0.75, 15000, 5000, 0),
(10000, 4, 0.15, 30000, 10000, 0),
(10000, 10, 0.05, 75000, 25000, 0),
(10000, 100, 0.04, 750000, 250000, 0),
(10000, 1000, 0.01, 7500000, 2500000, 0)
ON CONFLICT (buyin_cents, multiplier) DO NOTHING;

-- Insert multipliers for €250
INSERT INTO multipliers_expected (buyin_cents, multiplier, probability, prize1_cents, prize2_cents, prize3_cents) VALUES
(25000, 2, 0.75, 37500, 12500, 0),
(25000, 4, 0.15, 75000, 25000, 0),
(25000, 10, 0.05, 187500, 62500, 0),
(25000, 100, 0.04, 1875000, 625000, 0),
(25000, 1000, 0.01, 18750000, 6250000, 0)
ON CONFLICT (buyin_cents, multiplier) DO NOTHING;

-- Insert multipliers for €500
INSERT INTO multipliers_expected (buyin_cents, multiplier, probability, prize1_cents, prize2_cents, prize3_cents) VALUES
(50000, 2, 0.75, 75000, 25000, 0),
(50000, 4, 0.15, 150000, 50000, 0),
(50000, 10, 0.05, 375000, 125000, 0),
(50000, 100, 0.04, 3750000, 1250000, 0),
(50000, 1000, 0.01, 37500000, 12500000, 0)
ON CONFLICT (buyin_cents, multiplier) DO NOTHING;

-- Insert rakeback status levels
INSERT INTO rakeback_status (status_name, miles_required, rake_equiv, annual_bonus, rakeback_pct) VALUES
('Aluminium', 0, 0, 0, 0),
('Bronze', 2012, 201, 20, 10.0),
('Silver', 6000, 600, 75, 12.5),
('Gold', 20000, 2000, 300, 15.0),
('Platinum', 80000, 8000, 1400, 17.5),
('D1', 150000, 15000, 3000, 20.0),
('D2', 300000, 30000, 6750, 22.5),
('D3', 600000, 60000, 15000, 25.0),
('D4', 1200000, 120000, 33000, 27.5),
('D5', 2400000, 240000, 72000, 30.0),
('Red Diamond', 5000000, 500000, 165000, 33.0)
ON CONFLICT (status_name) DO NOTHING;
