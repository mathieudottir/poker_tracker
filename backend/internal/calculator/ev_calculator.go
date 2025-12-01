package calculator

import (
	"github.com/mathieudottir/poker_tracker/backend/internal/models"
	"github.com/mathieudottir/poker_tracker/backend/pkg/constants"
)

// CalculateTournamentEV calculates the expected value for a tournament based on buy-in
func CalculateTournamentEV(buyinCents int) float64 {
	return constants.GetExpectedEV(buyinCents)
}

// CalculateNetResult calculates net result for a tournament
func CalculateNetResult(prizeWon, buyin, rake int) int {
	return prizeWon - buyin - rake
}

// CalculateROI calculates return on investment
func CalculateROI(netResult, buyin, rake int) float64 {
	investment := float64(buyin + rake)
	if investment == 0 {
		return 0
	}
	return (float64(netResult) / investment) * 100
}

// CalculateEVROI calculates expected value ROI
func CalculateEVROI(evCents float64, buyin int) float64 {
	if buyin == 0 {
		return 0
	}
	return ((evCents - float64(buyin)) / float64(buyin)) * 100
}

// CalculateRakeback calculates rakeback based on status
func CalculateRakeback(totalRakeCents int, status string) int {
	rakebackPct := constants.GetRakebackForStatus(status)
	return int(float64(totalRakeCents) * (rakebackPct / 100.0))
}

// DetermineMultiplier determines the multiplier based on prizepool and buy-in
func DetermineMultiplier(prizepoolCents, buyinCents int) float64 {
	if buyinCents == 0 {
		return 0
	}
	// Prizepool = buyin * 3 * multiplier (for 3 players)
	// So multiplier = prizepool / (buyin * 3)
	return float64(prizepoolCents) / float64(buyinCents*3)
}

// GetPrizeDistribution returns prize distribution for a given multiplier and buy-in
func GetPrizeDistribution(buyinCents int, multiplier float64) (int, int, int) {
	multipliers := constants.GetMultiplierProbabilities(buyinCents)
	for _, mult := range multipliers {
		if mult.Multiplier == multiplier {
			return mult.Prize1Cents, mult.Prize2Cents, mult.Prize3Cents
		}
	}

	// If exact multiplier not found, calculate based on prizepool
	prizepool := int(float64(buyinCents*3) * multiplier)
	prize1 := int(float64(prizepool) * 0.75)
	prize2 := int(float64(prizepool) * 0.25)
	return prize1, prize2, 0
}

// CalculateChipsEV calculates simple chips EV for a hand (basic implementation)
// This is a simplified version - can be enhanced with proper equity calculations later
func CalculateChipsEV(hand *models.Hand) float64 {
	// Basic implementation: EV = chips won
	// TODO: Implement proper equity calculations based on cards and board
	return float64(hand.ChipsWon)
}

// CalculateWinrate calculates winrate in bb/100 hands
func CalculateWinrate(totalHands, totalChipsWon, averageBB int) float64 {
	if totalHands == 0 || averageBB == 0 {
		return 0
	}
	return (float64(totalChipsWon) / float64(averageBB)) / float64(totalHands) * 100
}

// CalculateStats calculates aggregated statistics for a user
func CalculateStats(tournaments []models.Tournament, userStatus string) models.Stats {
	stats := models.Stats{}

	var totalBuyinAndRake int
	for _, t := range tournaments {
		stats.TotalTournaments++
		stats.NetResultCents += t.NetResultCents
		stats.EVCents += t.EVCents
		stats.ChipsEV += t.ChipsEV
		stats.TotalRakeCents += t.RakeCents
		stats.TotalHands += t.HandsCount
		totalBuyinAndRake += t.BuyinCents + t.RakeCents
	}

	stats.TotalRakeback = CalculateRakeback(stats.TotalRakeCents, userStatus)

	if totalBuyinAndRake > 0 {
		stats.NetROI = (float64(stats.NetResultCents) / float64(totalBuyinAndRake)) * 100
		stats.EVROI = ((stats.EVCents - float64(totalBuyinAndRake-stats.TotalRakeCents)) / float64(totalBuyinAndRake-stats.TotalRakeCents)) * 100
	}

	if stats.TotalTournaments > 0 {
		stats.AvgChipsPerGame = stats.ChipsEV / float64(stats.TotalTournaments)
	}

	// Winrate is typically measured per tournament in SNGs
	if stats.TotalTournaments > 0 {
		stats.Winrate = float64(stats.NetResultCents) / float64(stats.TotalTournaments)
	}

	return stats
}
