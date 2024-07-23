package tests

import (
	"log"

	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models"
)

func CreateBattle(args ...any) {
	player, ok := args[0].(*Character)

	if !ok {
		log.Fatalf("[ERROR]: CreateBattle type assertion issue")
	}

	// When a Battle is created, an Enemy is created based on the provided MonsterType
	battle, enemy := models.NewBattle("BTL-001", player, models.BRONZE_DIF, models.ELITE)

	// The player Selects a BattleAction from the UI
	var playerBattleAction models.BattleAction = models.LIGHT_ATTACK

	// The enemy selects an action from their Pattern
	enemyPattern := models.EnemyNamesToPatterns[enemy.Name]
	enemyAction := battle.GetPatternIndex(enemyPattern)
	var enemyBattleAction models.BattleAction = enemyPattern[enemyAction]

	round := models.NewBattleRound("Round-001", player, enemy, playerBattleAction, enemyBattleAction, nil)

}
