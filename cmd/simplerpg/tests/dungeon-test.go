package tests

import (
	"fmt"

	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models"
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/dungeonhelpers"
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/logging"
)

func CreateDungeonMap(args ...any) {
	player, ok_pl := args[0].(*Character)
	// activeDun, ok_dun := args[0].(*Character)

	if !ok_pl {
		logging.LogError(logging.Logger, "(CreateDungeonMap(args ...any)) -> Problem with asserting player's type from provided args")
		panic("tests -> CreateDungeonMap -> player type assertion error")
	}

	// if !ok_dun {
	// 	logging.LogError(logging.Logger, "(CreateDungeonMap(args ...any)) -> Problem with asserting dungeon's type from provided args")
	// 	panic("tests -> CreateDungeonMap -> dungeon type assertion error")
	// }

	// 1. Creating the Map's Nodes
	node_1 := models.NewDungeonNode("Node-001", models.SIMPLE_ENEMY_ENCOUNTER, []string{"Node-004"}, []string{"Starting-Node"})
	node_2 := models.NewDungeonNode("Node-002", models.SIMPLE_ENEMY_ENCOUNTER, []string{"Node-006", "Node-007"}, []string{"Starting-Node"})
	node_3 := models.NewDungeonNode("Node-003", models.SIMPLE_ENEMY_ENCOUNTER, []string{"Node-007"}, []string{"Starting-Node"})

	node_4 := models.NewDungeonNode("Node-004", models.ELITE_ENEMY_ENCOUNTER, []string{"Node-005"}, []string{"Node-001"})
	node_5 := models.NewDungeonNode("Node-005", models.SIMPLE_ENEMY_ENCOUNTER, []string{"Node-006", "Node-009"}, []string{"Node-004"})
	node_6 := models.NewDungeonNode("Node-006", models.SHOP_ENCOUNTER, []string{"Node-009"}, []string{"Node-002", "Node-005", "Node-008"})

	node_7 := models.NewDungeonNode("Node-007", models.ELITE_ENEMY_ENCOUNTER, []string{"Node-008"}, []string{"Node-002", "Node-003"})
	node_8 := models.NewDungeonNode("Node-008", models.ELITE_ENEMY_ENCOUNTER, []string{"Node-006", "Node-009"}, []string{"Node-007"})

	node_9 := models.NewDungeonNode("Node-009", models.BOSS_ENEMY_ENCOUNTER, []string{}, []string{"Node-005", "Node-006", "Node-008"})

	allNodes := []*models.DungeonNode{node_1, node_2, node_3, node_4, node_5, node_6, node_7, node_8, node_9}

	// 2. Creating the Map
	dungeonMap := models.NewDungeonMap(allNodes...)

	// 3. Adding the Map to the Active Dungeon
	models.ActiveDungeon.Map = dungeonMap

	logging.GiveVertSpace(fmt.Sprintln("The Player's Current Position is:", models.ActiveDungeon.Map.CurrentNode))

	logging.GiveVertSpace("You have 5 moves to make...")
	for i := 0; i < 6; i++ {
		playerNodeSelection, isForward := dungeonhelpers.SelectNextNode(models.ActiveDungeon.Map)
		if playerNodeSelection == nil {
			logging.LogError(logging.Logger, "(CreateDungeonMap -> for loop) -> playerNodeSelection == nil\n")

			fmt.Printf("\nCurrent Position: (%s)\n", models.ActiveDungeon.Map.CurrentNode.ID)
			continue
		}

		if isForward {
			// Check if can go forward
			if dungeonMap.CurrentNode.IsFinalNode() {
				continue
			}
			models.ActiveDungeon.Map.MovePlayerForward(player, playerNodeSelection)
		} else {
			// Check if can go back
			if dungeonMap.CurrentNode.IsAfterStartingNode() {
				continue
			}
			models.ActiveDungeon.Map.MovePlayerBackwards(player, playerNodeSelection)
		}

	}

	logging.StdDivider("*", 100)
	logging.GiveVertSpace(models.ActiveDungeon.Map.CurrentNode.ID)
	logging.StdDivider("*", 100)
	logging.GiveVertSpace("THE END!")

}
