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
	fmt.Println("")
	player.DisplayEquipment()
	fmt.Println("")

	///////////////////////////////////////////////////////////////////////
	// tests.TextInfoWrapper("GIVING POTIONS", tests.GivePotions, player)
	///////////////////////////////////////////////////////////////////////

	// enemy := tests.CreateEnemy()

	// tests.TestWeapons()
	// logging.StdDivider("*", 100)
	// tests.TestArmors()
	// logging.StdDivider("*", 100)
	// tests.TestPotions()
	// logging.StdDivider("*", 100)

	// tests.TextInfoWrapper("ACCESSORY", tests.TestAccessories, player)

	///////////////////////////////////////////////////////////////////////
	// tests.TextInfoWrapper("LEVELING", tests.TestLeveling, player)
	///////////////////////////////////////////////////////////////////////

	// tests.TextInfoWrapper("RNG-LOOT", tests.TestLootRng, player)
	// tests.TestBattleRng()
	// tests.TestLootRng()
	// tests.TextInfoWrapper("BATTLE", tests.SimulateBattle, player, enemy)

	///////////////////////////////////////////////////////////////////////
	// tests.TextInfoWrapper("BATTLE", tests.CreateBattle, player)
	///////////////////////////////////////////////////////////////////////

	///////////////////////////////////////////////////////////////////////
	// tests.TextInfoWrapper("SHOP TESTING", tests.SetUpShop, player)
	///////////////////////////////////////////////////////////////////////

	///////////////////////////////////////////////////////////////////////
	tests.TextInfoWrapper("DUNGEON TESTING", tests.CreateDungeonMap, player)
	///////////////////////////////////////////////////////////////////////

}
