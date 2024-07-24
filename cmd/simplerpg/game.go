package simplerpg

import (
	"fmt"

	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models"
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/logging"
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/tests"
)

func RunGame() {
	// Create a logger that writes to the standard output
	logging.LogInit()
	models.DungeonInit(models.WOOD_DIF)

	player := tests.CreatePlayer()
	for i := 0; i < 3; i++ {
		potion_S := models.NewPotion(models.SMALL)
		potion_M := models.NewPotion(models.MEDIUM)
		potion_L := models.NewPotion(models.LARGE)

		player.MoveToInventory(potion_S)
		player.MoveToInventory(potion_M)
		player.MoveToInventory(potion_L)

	}
	fmt.Println("")
	player.DisplayInventory()
	fmt.Println("")
	// enemy := tests.CreateEnemy()

	// tests.TestWeapons()
	// logging.StdDivider("*", 100)
	// tests.TestArmors()
	// logging.StdDivider("*", 100)
	// tests.TestPotions()
	// logging.StdDivider("*", 100)
	// tests.TestAccessories()
	tests.TextInfoWrapper("LEVELING", tests.TestLeveling, player)
	// tests.TextInfoWrapper("RNG-LOOT", tests.TestLootRng, player)
	// tests.TestBattleRng()
	// tests.TestLootRng()
	// tests.TextInfoWrapper("BATTLE", tests.SimulateBattle, player, enemy)
	tests.TextInfoWrapper("BATTLE", tests.CreateBattle, player)

}
