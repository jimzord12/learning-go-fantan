package tests

import (
	"fmt"

	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models"
)

func TestLeveling() {
	for level := 1; level <= 10; level++ {
		fmt.Printf("Level %d -> %d: %.2f EXP required\n", level, level+1, models.ExpForNextLevel(level))
	}

	player := models.NewPlayer("PL-1", "John Wick", models.HUMAN)
	// enemy := models.NewArcDemonEnemy("BS-01", 1)

	fmt.Println()
	fmt.Printf("Player's current EXP: (%.2f)\n", player.Exp)
	fmt.Println()

	// expFromBattle := models.ExpGainedFromEnemy(player.Level, enemy)
	player.GainEXP(5000)

	fmt.Println()
	fmt.Printf("Player's current EXP: (%.2f)\n", player.Exp)
	fmt.Println()
}
