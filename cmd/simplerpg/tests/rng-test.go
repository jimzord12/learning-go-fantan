package tests

import (
	"fmt"
	"log"

	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models"
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/generalhelpers"
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

func TestLootRng(args ...any) {
	player, ok := args[0].(*Character)

	if !ok {
		log.Fatalf("[ERROR]: CreateBattle type assertion issue")
	}

	for i := 0; i < 50; i++ {
		models.LootEnemy(player)
	}
}
