package tests

import (
	"fmt"

	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models"
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/generalhelpers"
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/logging"
)

func TestBattleRng() {
	var playerRolls []float64
	var enemyRolls []float64

	for i := 0; i < 50; i++ {
		playerRoll := models.BattleLuckRoll(true) * 100
		enemyRoll := models.BattleLuckRoll(false) * 100

		playerRolls = append(playerRolls, playerRoll)
		enemyRolls = append(enemyRolls, enemyRoll)

		fmt.Printf("(%d) Player Rolled: [%f%%]\n", i, playerRoll)
		fmt.Printf("(%d) Enemy Rolled: [%f%%]\n", i, enemyRoll)
		fmt.Println()
		fmt.Println("*************************************************")
	}

	fmt.Println()
	fmt.Println("*************************************************")

	fmt.Printf("Player AVG Roll: [%.2f]\n", generalhelpers.GetAvgFromSlice(playerRolls))
	fmt.Printf("Enemy AVG Roll: [%.2f]\n", generalhelpers.GetAvgFromSlice(enemyRolls))

	fmt.Println()
	fmt.Println("*************************************************")
}

func TestLootRng() {
	for i := 0; i < 50; i++ {
		enemyDrops, hasLoot := models.CalcDrops()

		if !hasLoot {
			fmt.Printf("(%d) [=> NO LOOT <=]", i)
		} else {
			items := enemyDrops.GetLoot()
			fmt.Printf("(%d) -= HAS LOOT =-", i)
			fmt.Println()

			for i, item := range items {
				fmt.Printf("   (%d) %+v", i, *item)
			}
		}

		fmt.Println()
		logging.StdDivider("-", 75)
	}
}
