package dungeonhelpers

import (
	"fmt"

	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models"
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/inputhelpers"
)

func GetMovementTypeFromPlayer() bool {
	passFn := func(i int) bool {
		if i > 2 || i < 1 {
			return false
		}
		return true
	}
	moveType := inputhelpers.GetTerminalInputInt("Select Type of Movement:\n1. Forward\n2. Backwards", passFn)

	return moveType == 1
}

func GetSelectedNodeFromPlayer(nodes []*models.DungeonNode) *models.DungeonNode {
	passFn := func(i int) bool {
		return i <= len(nodes)-1
	}

	for j := 0; j < len(nodes); j++ {
		fmt.Printf("\n (%d) For Node: (%s) of Encouter: (%d)", j, nodes[j].ID, nodes[j].Encounter)
	}

	selectedNodeIndex := inputhelpers.GetTerminalInputInt("\nSelect a Node:\n", passFn)

	return nodes[selectedNodeIndex]
}

// func DisplayDungeonNodes(nodes []*models.DungeonNode) {
// 	for idx, node := range nodes {
// 		fmt.Printf("\n(%d) Dungeon Node: %+v\n", idx, node)
// 	}
// }

func SelectNextNode(dungeonMap *models.DungeonMap) (*models.DungeonNode, bool) {
	direction := GetMovementTypeFromPlayer()

	// Directon is Forward
	if direction {
		nextNodes, err := dungeonMap.GetNextNodes()
		if err != nil || len(nextNodes) == 0 {
			fmt.Println(err)
			return nil, false

		}

		return GetSelectedNodeFromPlayer(nextNodes), true
	}

	// Directon is Backwards
	if !direction {
		prevNodes, err := dungeonMap.GetPrevNodes()
		println(prevNodes)
		if err != nil || len(prevNodes) == 0 {
			fmt.Println(err)
			return nil, false

		}

		return GetSelectedNodeFromPlayer(prevNodes), false
	}

	return nil, false
}
