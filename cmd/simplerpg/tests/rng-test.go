package tests

import (
	"fmt"

	rngbattle "github.com/jimzord12/learning-go-fantan/cmd/simplerpg/mechanics/rng-systems/rng-battle"
	rngloot "github.com/jimzord12/learning-go-fantan/cmd/simplerpg/mechanics/rng-systems/rng-loot"
	rpgcharacters "github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models/rpg-characters"
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/generalhelpers"
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/logging"
)

func TestBattleRng() {
	var playerRolls []float64
	var enemyRolls []float64

	for i := 0; i < 50; i++ {
		playerRoll := rngbattle.BattleLuckRoll(rpgcharacters.ELF) * 100
		enemyRoll := rngbattle.BattleLuckRoll(rpgcharacters.BOSS) * 100

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
		enemyDrops, hasLoot := rngloot.CalcDrops()

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
