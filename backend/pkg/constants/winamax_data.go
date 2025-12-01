package constants

import "github.com/mathieudottir/poker_tracker/backend/internal/models"

// WinamaxMultipliers contains all the hardcoded multiplier data for each buy-in level
// Multipliers: 2x, 3x, 4x, 5x, 10x, 50x, 100x, 1000x
var WinamaxMultipliers = map[int][]models.MultiplierExpected{
	// €0.25 (25 cents)
	25: {
		{BuyinCents: 25, Multiplier: 2, Probability: 0.665, Prize1Cents: 38, Prize2Cents: 13, Prize3Cents: 0},
		{BuyinCents: 25, Multiplier: 3, Probability: 0.16, Prize1Cents: 56, Prize2Cents: 19, Prize3Cents: 0},
		{BuyinCents: 25, Multiplier: 4, Probability: 0.08, Prize1Cents: 75, Prize2Cents: 25, Prize3Cents: 0},
		{BuyinCents: 25, Multiplier: 5, Probability: 0.04, Prize1Cents: 94, Prize2Cents: 31, Prize3Cents: 0},
		{BuyinCents: 25, Multiplier: 10, Probability: 0.03, Prize1Cents: 188, Prize2Cents: 63, Prize3Cents: 0},
		{BuyinCents: 25, Multiplier: 50, Probability: 0.015, Prize1Cents: 938, Prize2Cents: 313, Prize3Cents: 0},
		{BuyinCents: 25, Multiplier: 100, Probability: 0.008, Prize1Cents: 1875, Prize2Cents: 625, Prize3Cents: 0},
		{BuyinCents: 25, Multiplier: 1000, Probability: 0.002, Prize1Cents: 18750, Prize2Cents: 6250, Prize3Cents: 0},
	},
	// €0.50 (50 cents)
	50: {
		{BuyinCents: 50, Multiplier: 2, Probability: 0.665, Prize1Cents: 75, Prize2Cents: 25, Prize3Cents: 0},
		{BuyinCents: 50, Multiplier: 3, Probability: 0.16, Prize1Cents: 113, Prize2Cents: 38, Prize3Cents: 0},
		{BuyinCents: 50, Multiplier: 4, Probability: 0.08, Prize1Cents: 150, Prize2Cents: 50, Prize3Cents: 0},
		{BuyinCents: 50, Multiplier: 5, Probability: 0.04, Prize1Cents: 188, Prize2Cents: 63, Prize3Cents: 0},
		{BuyinCents: 50, Multiplier: 10, Probability: 0.03, Prize1Cents: 375, Prize2Cents: 125, Prize3Cents: 0},
		{BuyinCents: 50, Multiplier: 50, Probability: 0.015, Prize1Cents: 1875, Prize2Cents: 625, Prize3Cents: 0},
		{BuyinCents: 50, Multiplier: 100, Probability: 0.008, Prize1Cents: 3750, Prize2Cents: 1250, Prize3Cents: 0},
		{BuyinCents: 50, Multiplier: 1000, Probability: 0.002, Prize1Cents: 37500, Prize2Cents: 12500, Prize3Cents: 0},
	},
	// €1 (100 cents)
	100: {
		{BuyinCents: 100, Multiplier: 2, Probability: 0.665, Prize1Cents: 150, Prize2Cents: 50, Prize3Cents: 0},
		{BuyinCents: 100, Multiplier: 3, Probability: 0.16, Prize1Cents: 225, Prize2Cents: 75, Prize3Cents: 0},
		{BuyinCents: 100, Multiplier: 4, Probability: 0.08, Prize1Cents: 300, Prize2Cents: 100, Prize3Cents: 0},
		{BuyinCents: 100, Multiplier: 5, Probability: 0.04, Prize1Cents: 375, Prize2Cents: 125, Prize3Cents: 0},
		{BuyinCents: 100, Multiplier: 10, Probability: 0.03, Prize1Cents: 750, Prize2Cents: 250, Prize3Cents: 0},
		{BuyinCents: 100, Multiplier: 50, Probability: 0.015, Prize1Cents: 3750, Prize2Cents: 1250, Prize3Cents: 0},
		{BuyinCents: 100, Multiplier: 100, Probability: 0.008, Prize1Cents: 7500, Prize2Cents: 2500, Prize3Cents: 0},
		{BuyinCents: 100, Multiplier: 1000, Probability: 0.002, Prize1Cents: 75000, Prize2Cents: 25000, Prize3Cents: 0},
	},
	// €2 (200 cents)
	200: {
		{BuyinCents: 200, Multiplier: 2, Probability: 0.665, Prize1Cents: 300, Prize2Cents: 100, Prize3Cents: 0},
		{BuyinCents: 200, Multiplier: 3, Probability: 0.16, Prize1Cents: 450, Prize2Cents: 150, Prize3Cents: 0},
		{BuyinCents: 200, Multiplier: 4, Probability: 0.08, Prize1Cents: 600, Prize2Cents: 200, Prize3Cents: 0},
		{BuyinCents: 200, Multiplier: 5, Probability: 0.04, Prize1Cents: 750, Prize2Cents: 250, Prize3Cents: 0},
		{BuyinCents: 200, Multiplier: 10, Probability: 0.03, Prize1Cents: 1500, Prize2Cents: 500, Prize3Cents: 0},
		{BuyinCents: 200, Multiplier: 50, Probability: 0.015, Prize1Cents: 7500, Prize2Cents: 2500, Prize3Cents: 0},
		{BuyinCents: 200, Multiplier: 100, Probability: 0.008, Prize1Cents: 15000, Prize2Cents: 5000, Prize3Cents: 0},
		{BuyinCents: 200, Multiplier: 1000, Probability: 0.002, Prize1Cents: 150000, Prize2Cents: 50000, Prize3Cents: 0},
	},
	// €5 (500 cents)
	500: {
		{BuyinCents: 500, Multiplier: 2, Probability: 0.665, Prize1Cents: 750, Prize2Cents: 250, Prize3Cents: 0},
		{BuyinCents: 500, Multiplier: 3, Probability: 0.16, Prize1Cents: 1125, Prize2Cents: 375, Prize3Cents: 0},
		{BuyinCents: 500, Multiplier: 4, Probability: 0.08, Prize1Cents: 1500, Prize2Cents: 500, Prize3Cents: 0},
		{BuyinCents: 500, Multiplier: 5, Probability: 0.04, Prize1Cents: 1875, Prize2Cents: 625, Prize3Cents: 0},
		{BuyinCents: 500, Multiplier: 10, Probability: 0.03, Prize1Cents: 3750, Prize2Cents: 1250, Prize3Cents: 0},
		{BuyinCents: 500, Multiplier: 50, Probability: 0.015, Prize1Cents: 18750, Prize2Cents: 6250, Prize3Cents: 0},
		{BuyinCents: 500, Multiplier: 100, Probability: 0.008, Prize1Cents: 37500, Prize2Cents: 12500, Prize3Cents: 0},
		{BuyinCents: 500, Multiplier: 1000, Probability: 0.002, Prize1Cents: 375000, Prize2Cents: 125000, Prize3Cents: 0},
	},
	// €10 (1000 cents)
	1000: {
		{BuyinCents: 1000, Multiplier: 2, Probability: 0.665, Prize1Cents: 1500, Prize2Cents: 500, Prize3Cents: 0},
		{BuyinCents: 1000, Multiplier: 3, Probability: 0.16, Prize1Cents: 2250, Prize2Cents: 750, Prize3Cents: 0},
		{BuyinCents: 1000, Multiplier: 4, Probability: 0.08, Prize1Cents: 3000, Prize2Cents: 1000, Prize3Cents: 0},
		{BuyinCents: 1000, Multiplier: 5, Probability: 0.04, Prize1Cents: 3750, Prize2Cents: 1250, Prize3Cents: 0},
		{BuyinCents: 1000, Multiplier: 10, Probability: 0.03, Prize1Cents: 7500, Prize2Cents: 2500, Prize3Cents: 0},
		{BuyinCents: 1000, Multiplier: 50, Probability: 0.015, Prize1Cents: 37500, Prize2Cents: 12500, Prize3Cents: 0},
		{BuyinCents: 1000, Multiplier: 100, Probability: 0.008, Prize1Cents: 75000, Prize2Cents: 25000, Prize3Cents: 0},
		{BuyinCents: 1000, Multiplier: 1000, Probability: 0.002, Prize1Cents: 750000, Prize2Cents: 250000, Prize3Cents: 0},
	},
}

// WinamaxRakebackStatus contains the rakeback percentages for each Winamax status level
var WinamaxRakebackStatus = []models.RakebackStatus{
	{StatusName: "Aluminium", RakebackPct: 0},
	{StatusName: "Bronze", RakebackPct: 10.0},
	{StatusName: "Argent", RakebackPct: 15.0},
	{StatusName: "Or", RakebackPct: 20.0},
	{StatusName: "Platine", RakebackPct: 25.0},
	{StatusName: "Diamant", RakebackPct: 30.0},
	{StatusName: "Red Diamond", RakebackPct: 33.0},
}

// GetMultiplierProbabilities returns the multiplier data for a given buy-in
func GetMultiplierProbabilities(buyinCents int) []models.MultiplierExpected {
	if data, ok := WinamaxMultipliers[buyinCents]; ok {
		return data
	}
	return []models.MultiplierExpected{}
}

// GetExpectedEV calculates the expected value for a given buy-in based on multiplier probabilities
// For Expresso Nitro (winner-takes-all): EV = sum(probability * (prizepool/3 - entry))
func GetExpectedEV(totalEntryCents int) float64 {
	multipliers := GetMultiplierProbabilities(totalEntryCents)
	if len(multipliers) == 0 {
		return 0
	}

	var ev float64
	for _, m := range multipliers {
		// Expresso Nitro is winner-takes-all: 1/3 chance to win prizepool
		// Prizepool for this multiplier = total_entry * multiplier
		prizepool := float64(totalEntryCents) * m.Multiplier
		// EV for this multiplier = prob * (1/3 * prizepool - total_entry)
		evForMultiplier := (prizepool / 3.0) - float64(totalEntryCents)
		ev += m.Probability * evForMultiplier
	}

	return ev
}

// GetRakebackForStatus returns the rakeback percentage for a given status
func GetRakebackForStatus(status string) float64 {
	for _, s := range WinamaxRakebackStatus {
		if s.StatusName == status {
			return s.RakebackPct
		}
	}
	return 0.0
}
