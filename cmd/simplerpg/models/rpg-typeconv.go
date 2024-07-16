package models

func DifficultyToMaterial(d Difficulty) Material {
	switch d {
	case WOOD_DIF:
		return WOOD
	case BRONZE_DIF:
		return BRONZE
	case IRON_DIF:
		return IRON
	case STEEL_DIF:
		return STEEL
	case TITANIUM_DIF:
		return TITANIUM
	case MYTHRIL_DIF:
		return MYTHRIL
	default:
		return NO_MATERIAL
	}
}

// ConvertMaterialToDifficulty converts a Material to a Difficulty
func MaterialToDifficulty(m Material) Difficulty {
	switch m {
	case WOOD:
		return WOOD_DIF
	case BRONZE:
		return BRONZE_DIF
	case IRON:
		return IRON_DIF
	case STEEL:
		return STEEL_DIF
	case TITANIUM:
		return TITANIUM_DIF
	case MYTHRIL:
		return MYTHRIL_DIF
	default:
		return 0 // No corresponding difficulty
	}
}
