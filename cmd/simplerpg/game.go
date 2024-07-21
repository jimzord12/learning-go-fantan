package simplerpg

import (
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models"
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/logging"
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/tests"
)

func RunGame() {
	// Create a logger that writes to the standard output
	logging.LogInit()
	models.DungeonInit(models.WOOD_DIF)

	player := tests.CreatePlayer()
	enemy := tests.CreateEnemy()

	// tests.TestWeapons()
	// logging.StdDivider("*", 100)
	// tests.TestArmors()
	// logging.StdDivider("*", 100)
	// tests.TestPotions()
	// logging.StdDivider("*", 100)
	// tests.TestAccessories()
	tests.TextInfoWrapper(tests.TestLeveling, "LEVELING")
	// tests.TestBattleRng()
	// tests.TestLootRng()
	tests.TextInfoWrapper(tests.SimulateBattle, "BATTLE")

}
