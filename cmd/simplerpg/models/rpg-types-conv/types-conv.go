package typeconvpkg

import (
	dungpkg "github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models/rpg-dungeon"
	itemspkg "github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models/rpg-items"
)

type Material = itemspkg.Material
type Difficulty = dungpkg.Difficulty

// ConvertDifficultyToMaterial converts a Difficulty to a Material
func DifficultyToMaterial(d Difficulty) Material {
	switch d {
	case dungpkg.WOOD_DIF:
		return itemspkg.WOOD
	case dungpkg.BRONZE_DIF:
		return itemspkg.BRONZE
	case dungpkg.IRON_DIF:
		return itemspkg.IRON
	case dungpkg.STEEL_DIF:
		return itemspkg.STEEL
	case dungpkg.TITANIUM_DIF:
		return itemspkg.TITANIUM
	case dungpkg.MYTHRIL_DIF:
		return itemspkg.MYTHRIL
	default:
		return itemspkg.NO_MATERIAL
	}
}

// ConvertMaterialToDifficulty converts a Material to a Difficulty
func MaterialToDifficulty(m Material) Difficulty {
	switch m {
	case itemspkg.WOOD:
		return dungpkg.WOOD_DIF
	case itemspkg.BRONZE:
		return dungpkg.BRONZE_DIF
	case itemspkg.IRON:
		return dungpkg.IRON_DIF
	case itemspkg.STEEL:
		return dungpkg.STEEL_DIF
	case itemspkg.TITANIUM:
		return dungpkg.TITANIUM_DIF
	case itemspkg.MYTHRIL:
		return dungpkg.MYTHRIL_DIF
	default:
		return 0 // No corresponding difficulty
	}
}
