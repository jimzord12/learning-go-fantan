package rngbattle

import (
	"math/rand"

	rpgcharacters "github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models/rpg-characters"
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/generalhelpers"
)

type CharacterType = rpgcharacters.CharacterType

var PlayerTypes = rpgcharacters.PlayerTypes

func rollLuck(isPlayer bool) int {
	var weights []int

	if isPlayer {
		// Player has higher chances to get 3-7
		weights = PlayerBattleRollChances
	} else {
		// Enemy has higher chances to get 1-4
		weights = EnemyBattleRollChances
	}

	cumulativeWeights := make([]int, len(weights))

	cumulativeWeights[0] = weights[0]
	for i := 1; i < len(weights); i++ {
		cumulativeWeights[i] = cumulativeWeights[i-1] + weights[i]
	}

	randValue := rand.Intn(cumulativeWeights[len(cumulativeWeights)-1]) + 1

	for i, cw := range cumulativeWeights {
		if randValue <= cw {
			return i
		}
	}
	return 0 // fallback
}

// Function to calculate the boost percentage based on the luck roll
func luckBoost(luck int) float64 {
	return float64(luck) / 4.0 // 0 to 2.0 (0% to 200%)
}

func BattleLuckRoll(charType CharacterType) float64 {
	luck := rollLuck(generalhelpers.ExistsInSlice(PlayerTypes, charType))

	return luckBoost(luck)
}
