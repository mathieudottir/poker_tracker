package constants

import "github.com/mathieudottir/poker_tracker/backend/internal/models"

// WinamaxMultipliers contains all the hardcoded multiplier data for each buy-in level
var WinamaxMultipliers = map[int][]models.MultiplierExpected{
	// €0.25 (25 cents)
	25: {
		{BuyinCents: 25, Multiplier: 2, Probability: 0.80, Prize1Cents: 38, Prize2Cents: 13, Prize3Cents: 0},
		{BuyinCents: 25, Multiplier: 4, Probability: 0.10, Prize1Cents: 75, Prize2Cents: 25, Prize3Cents: 0},
		{BuyinCents: 25, Multiplier: 10, Probability: 0.05, Prize1Cents: 188, Prize2Cents: 63, Prize3Cents: 0},
		{BuyinCents: 25, Multiplier: 100, Probability: 0.04, Prize1Cents: 1875, Prize2Cents: 625, Prize3Cents: 0},
		{BuyinCents: 25, Multiplier: 1000, Probability: 0.01, Prize1Cents: 18750, Prize2Cents: 6250, Prize3Cents: 0},
	},
	// €0.50 (50 cents)
	50: {
		{BuyinCents: 50, Multiplier: 2, Probability: 0.80, Prize1Cents: 75, Prize2Cents: 25, Prize3Cents: 0},
		{BuyinCents: 50, Multiplier: 4, Probability: 0.10, Prize1Cents: 150, Prize2Cents: 50, Prize3Cents: 0},
		{BuyinCents: 50, Multiplier: 10, Probability: 0.05, Prize1Cents: 375, Prize2Cents: 125, Prize3Cents: 0},
		{BuyinCents: 50, Multiplier: 100, Probability: 0.04, Prize1Cents: 3750, Prize2Cents: 1250, Prize3Cents: 0},
		{BuyinCents: 50, Multiplier: 1000, Probability: 0.01, Prize1Cents: 37500, Prize2Cents: 12500, Prize3Cents: 0},
	},
	// €1 (100 cents)
	100: {
		{BuyinCents: 100, Multiplier: 2, Probability: 0.80, Prize1Cents: 150, Prize2Cents: 50, Prize3Cents: 0},
		{BuyinCents: 100, Multiplier: 4, Probability: 0.10, Prize1Cents: 300, Prize2Cents: 100, Prize3Cents: 0},
		{BuyinCents: 100, Multiplier: 10, Probability: 0.05, Prize1Cents: 750, Prize2Cents: 250, Prize3Cents: 0},
		{BuyinCents: 100, Multiplier: 100, Probability: 0.04, Prize1Cents: 7500, Prize2Cents: 2500, Prize3Cents: 0},
		{BuyinCents: 100, Multiplier: 1000, Probability: 0.01, Prize1Cents: 75000, Prize2Cents: 25000, Prize3Cents: 0},
	},
	// €2 (200 cents)
	200: {
		{BuyinCents: 200, Multiplier: 2, Probability: 0.80, Prize1Cents: 300, Prize2Cents: 100, Prize3Cents: 0},
		{BuyinCents: 200, Multiplier: 4, Probability: 0.10, Prize1Cents: 600, Prize2Cents: 200, Prize3Cents: 0},
		{BuyinCents: 200, Multiplier: 10, Probability: 0.05, Prize1Cents: 1500, Prize2Cents: 500, Prize3Cents: 0},
		{BuyinCents: 200, Multiplier: 100, Probability: 0.04, Prize1Cents: 15000, Prize2Cents: 5000, Prize3Cents: 0},
		{BuyinCents: 200, Multiplier: 1000, Probability: 0.01, Prize1Cents: 150000, Prize2Cents: 50000, Prize3Cents: 0},
	},
	// €5 (500 cents)
	500: {
		{BuyinCents: 500, Multiplier: 2, Probability: 0.80, Prize1Cents: 750, Prize2Cents: 250, Prize3Cents: 0},
		{BuyinCents: 500, Multiplier: 4, Probability: 0.10, Prize1Cents: 1500, Prize2Cents: 500, Prize3Cents: 0},
		{BuyinCents: 500, Multiplier: 10, Probability: 0.05, Prize1Cents: 3750, Prize2Cents: 1250, Prize3Cents: 0},
		{BuyinCents: 500, Multiplier: 100, Probability: 0.04, Prize1Cents: 37500, Prize2Cents: 12500, Prize3Cents: 0},
		{BuyinCents: 500, Multiplier: 1000, Probability: 0.01, Prize1Cents: 375000, Prize2Cents: 125000, Prize3Cents: 0},
	},
	// €10 (1000 cents)
	1000: {
		{BuyinCents: 1000, Multiplier: 2, Probability: 0.80, Prize1Cents: 1500, Prize2Cents: 500, Prize3Cents: 0},
		{BuyinCents: 1000, Multiplier: 4, Probability: 0.10, Prize1Cents: 3000, Prize2Cents: 1000, Prize3Cents: 0},
		{BuyinCents: 1000, Multiplier: 10, Probability: 0.05, Prize1Cents: 7500, Prize2Cents: 2500, Prize3Cents: 0},
		{BuyinCents: 1000, Multiplier: 100, Probability: 0.04, Prize1Cents: 75000, Prize2Cents: 25000, Prize3Cents: 0},
		{BuyinCents: 1000, Multiplier: 1000, Probability: 0.01, Prize1Cents: 750000, Prize2Cents: 250000, Prize3Cents: 0},
	},
	// €25 (2500 cents)
	2500: {
		{BuyinCents: 2500, Multiplier: 2, Probability: 0.80, Prize1Cents: 3750, Prize2Cents: 1250, Prize3Cents: 0},
		{BuyinCents: 2500, Multiplier: 4, Probability: 0.10, Prize1Cents: 7500, Prize2Cents: 2500, Prize3Cents: 0},
		{BuyinCents: 2500, Multiplier: 10, Probability: 0.05, Prize1Cents: 18750, Prize2Cents: 6250, Prize3Cents: 0},
		{BuyinCents: 2500, Multiplier: 100, Probability: 0.04, Prize1Cents: 187500, Prize2Cents: 62500, Prize3Cents: 0},
		{BuyinCents: 2500, Multiplier: 1000, Probability: 0.01, Prize1Cents: 1875000, Prize2Cents: 625000, Prize3Cents: 0},
	},
	// €50 (5000 cents)
	5000: {
		{BuyinCents: 5000, Multiplier: 2, Probability: 0.80, Prize1Cents: 7500, Prize2Cents: 2500, Prize3Cents: 0},
		{BuyinCents: 5000, Multiplier: 4, Probability: 0.10, Prize1Cents: 15000, Prize2Cents: 5000, Prize3Cents: 0},
		{BuyinCents: 5000, Multiplier: 10, Probability: 0.05, Prize1Cents: 37500, Prize2Cents: 12500, Prize3Cents: 0},
		{BuyinCents: 5000, Multiplier: 100, Probability: 0.04, Prize1Cents: 375000, Prize2Cents: 125000, Prize3Cents: 0},
		{BuyinCents: 5000, Multiplier: 1000, Probability: 0.01, Prize1Cents: 3750000, Prize2Cents: 1250000, Prize3Cents: 0},
	},
	// €100 (10000 cents)
	10000: {
		{BuyinCents: 10000, Multiplier: 2, Probability: 0.75, Prize1Cents: 15000, Prize2Cents: 5000, Prize3Cents: 0},
		{BuyinCents: 10000, Multiplier: 4, Probability: 0.15, Prize1Cents: 30000, Prize2Cents: 10000, Prize3Cents: 0},
		{BuyinCents: 10000, Multiplier: 10, Probability: 0.05, Prize1Cents: 75000, Prize2Cents: 25000, Prize3Cents: 0},
		{BuyinCents: 10000, Multiplier: 100, Probability: 0.04, Prize1Cents: 750000, Prize2Cents: 250000, Prize3Cents: 0},
		{BuyinCents: 10000, Multiplier: 1000, Probability: 0.01, Prize1Cents: 7500000, Prize2Cents: 2500000, Prize3Cents: 0},
	},
	// €250 (25000 cents)
	25000: {
		{BuyinCents: 25000, Multiplier: 2, Probability: 0.75, Prize1Cents: 37500, Prize2Cents: 12500, Prize3Cents: 0},
		{BuyinCents: 25000, Multiplier: 4, Probability: 0.15, Prize1Cents: 75000, Prize2Cents: 25000, Prize3Cents: 0},
		{BuyinCents: 25000, Multiplier: 10, Probability: 0.05, Prize1Cents: 187500, Prize2Cents: 62500, Prize3Cents: 0},
		{BuyinCents: 25000, Multiplier: 100, Probability: 0.04, Prize1Cents: 1875000, Prize2Cents: 625000, Prize3Cents: 0},
		{BuyinCents: 25000, Multiplier: 1000, Probability: 0.01, Prize1Cents: 18750000, Prize2Cents: 6250000, Prize3Cents: 0},
	},
	// €500 (50000 cents)
	50000: {
		{BuyinCents: 50000, Multiplier: 2, Probability: 0.75, Prize1Cents: 75000, Prize2Cents: 25000, Prize3Cents: 0},
		{BuyinCents: 50000, Multiplier: 4, Probability: 0.15, Prize1Cents: 150000, Prize2Cents: 50000, Prize3Cents: 0},
		{BuyinCents: 50000, Multiplier: 10, Probability: 0.05, Prize1Cents: 375000, Prize2Cents: 125000, Prize3Cents: 0},
		{BuyinCents: 50000, Multiplier: 100, Probability: 0.04, Prize1Cents: 3750000, Prize2Cents: 1250000, Prize3Cents: 0},
		{BuyinCents: 50000, Multiplier: 1000, Probability: 0.01, Prize1Cents: 37500000, Prize2Cents: 12500000, Prize3Cents: 0},
	},
}

// WinamaxRakebackStatus contains all rakeback status levels
var WinamaxRakebackStatus = []models.RakebackStatus{
	{StatusName: "Aluminium", MilesRequired: 0, RakeEquiv: 0, AnnualBonus: 0, RakebackPct: 0},
	{StatusName: "Bronze", MilesRequired: 2012, RakeEquiv: 201, AnnualBonus: 20, RakebackPct: 10.0},
	{StatusName: "Silver", MilesRequired: 6000, RakeEquiv: 600, AnnualBonus: 75, RakebackPct: 12.5},
	{StatusName: "Gold", MilesRequired: 20000, RakeEquiv: 2000, AnnualBonus: 300, RakebackPct: 15.0},
	{StatusName: "Platinum", MilesRequired: 80000, RakeEquiv: 8000, AnnualBonus: 1400, RakebackPct: 17.5},
	{StatusName: "D1", MilesRequired: 150000, RakeEquiv: 15000, AnnualBonus: 3000, RakebackPct: 20.0},
	{StatusName: "D2", MilesRequired: 300000, RakeEquiv: 30000, AnnualBonus: 6750, RakebackPct: 22.5},
	{StatusName: "D3", MilesRequired: 600000, RakeEquiv: 60000, AnnualBonus: 15000, RakebackPct: 25.0},
	{StatusName: "D4", MilesRequired: 1200000, RakeEquiv: 120000, AnnualBonus: 33000, RakebackPct: 27.5},
	{StatusName: "D5", MilesRequired: 2400000, RakeEquiv: 240000, AnnualBonus: 72000, RakebackPct: 30.0},
	{StatusName: "Red Diamond", MilesRequired: 5000000, RakeEquiv: 500000, AnnualBonus: 165000, RakebackPct: 33.0},
}

// GetExpectedEV calculates the expected EV for a given buy-in in cents
func GetExpectedEV(buyinCents int) float64 {
	multipliers, exists := WinamaxMultipliers[buyinCents]
	if !exists {
		return 0
	}

	var expectedEV float64
	for _, mult := range multipliers {
		// For 3-max, 1st place gets 75% of prizepool, 2nd gets 25%
		expectedValue := (float64(mult.Prize1Cents) * 0.333) + (float64(mult.Prize2Cents) * 0.333) + (float64(mult.Prize3Cents) * 0.333)
		expectedEV += expectedValue * mult.Probability
	}

	return expectedEV
}

// GetRakebackForStatus returns the rakeback percentage for a given status
func GetRakebackForStatus(status string) float64 {
	for _, rb := range WinamaxRakebackStatus {
		if rb.StatusName == status {
			return rb.RakebackPct
		}
	}
	return 0
}

// GetMultiplierProbabilities returns all multiplier probabilities for a buy-in
func GetMultiplierProbabilities(buyinCents int) []models.MultiplierExpected {
	return WinamaxMultipliers[buyinCents]
}
