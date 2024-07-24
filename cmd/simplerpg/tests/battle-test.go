package tests

import (
	"fmt"
	"log"

	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models"
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/battlehelpers"
)

func CreateBattle(args ...any) {
	player, ok := args[0].(*Character)

	if !ok {
		log.Fatalf("[ERROR]: CreateBattle type assertion issue")
	}

	// When a Battle is created, an Enemy is created based on the provided MonsterType
	battle, enemy := models.NewBattle("BTL-001", player, models.BRONZE_DIF, models.ELITE)

	hasBattleEnded := false
	for i := 1; true; i++ {
		if hasBattleEnded {
			break
		}

		// The player Selects a BattleAction from the UI
		playerBattleAction, consumable, err := battlehelpers.GetBattleAction(player)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// The enemy selects an action from their Pattern
		enemyPattern := models.EnemyNamesToPatterns[enemy.Name]
		enemyAction := battle.GetPatternIndex(enemyPattern)
		var enemyBattleAction models.BattleAction = enemyPattern[enemyAction]

		// Creating the Round
		round := models.NewBattleRound(fmt.Sprintf("Round-%d", i), player, enemy, playerBattleAction, enemyBattleAction, consumable)

		// Execute selected Actions
		//TODO: Check the: func PerformRound(round BattleRound) (hasBattleEnded bool) in METHODS
		hasBattleEnded = models.PerformRound(*round)
		battle.AddRound(*round)
	}

	fmt.Println()
	fmt.Println()
	fmt.Println("==== BATTLE ENDED ==== ")
	fmt.Println(battle)
}
