package modelstest

import (
	"testing"

	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models"
)

type Character = models.Character

func TestPerformBattleAction(t *testing.T) {
	goodPlayer := models.NewPlayer("PL-001", "Yellow Samurai", models.HUMAN, 250)
	emptyPlayer := &Character{}

	goodEnemy := models.NewGoblinEnemy("MON-001", 230)
	emptyEnemy := &Character{}

	goodConsumable := models.NewPotion(models.MEDIUM)
	badConsumable := models.NewArmor(models.BRONZE)
	emptyConsumable := &Item{}

	goodAction := models.LIGHT_ATTACK
	badAction := BattleAction(86)
	battleAction := AllBattleActions

	type testType struct {
		name        string
		action      BattleAction
		attacker    *Character
		defender    *Character
		consumable  *Item
		expectedErr error
	}

	testCases := []testType{
		{
			name:       "Fail - Bad Action",
			action:     badAction,
			attacker:   goodPlayer,
			defender:   goodEnemy,
			consumable: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

		})
	}
}
