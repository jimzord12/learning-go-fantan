package itemspkg

import (
	"fmt"

	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/logging"
)

var logger = logging.Logger

///// WeaponType Methods /////

func (wt WeaponType) GetDmg(weight float64, material Material) float64 {
	return material.GetToughness() * (WEAPON_BASE_DMG + (weight * WEAPON_BASE_MULTI_PER_KG))
}

///// Material Methods /////

// This could be stored as a const, but its not for practise purposes
func (mat Material) GetWeight(volume float64) float64 {
	switch mat {
	case WOOD:
		return volume * WOOD_DENSITY
	case BRONZE:
		return volume * BRONZE_DENSITY
	case IRON:
		return volume * IRON_DENSITY
	case STEEL:
		return volume * STEEL_DENSITY
	case TITANIUM:
		return volume * TITANIUM_DENSITY
	case MYTHRIL:
		return volume * MYTHRIL_DENSITY
	default:
		logging.LogError(logger, fmt.Sprintf("Something went wrong while attemping to calculate this material's (%v) weight\n", mat))
		return -1
	}
}

func (mat Material) GetToughness() float64 {
	switch mat {
	case WOOD:
		return WOOD_TOUGHNESS_MULTI
	case BRONZE:
		return BRONZE_TOUGHNESS_MULTI
	case IRON:
		return IRON_TOUGHNESS_MULTI
	case STEEL:
		return STEEL_TOUGHNESS_MULTI
	case TITANIUM:
		return TITANIUM_TOUGHNESS_MULTI
	case MYTHRIL:
		return MYTHRIL_TOUGHNESS_MULTI
	default:
		logging.LogError(logger, fmt.Sprintf("Something went wrong while getting this material's (%v) toughness\n", mat))
		return -1
	}
}

func (mat Material) String() string {
	switch mat {
	case NO_MATERIAL:
		return "None"
	case WOOD:
		return "Wooden"
	case BRONZE:
		return "Bronze"
	case IRON:
		return "Iron"
	case STEEL:
		return "Steel"
	case TITANIUM:
		return "Titanium"
	case MYTHRIL:
		return "Mythril"
	default:
		logging.LogError(logger, "Something went wrong while attemping to return the string represensation of a material")
		return "[ERROR]"
	}
}

///// ItemType Methods /////

func (i ItemType) String() string {
	switch i {
	case WEAPON:
		return "Weapon"
	case ARMOR:
		return "Armor"
	case POTION:
		return "Potion"
	case ACCESORY:
		return "Accessory"
	default:
		logging.LogError(logger, "Something went wrong while attemping to return the represensation of an itemType")
		return "[ERROR]"
	}
}

// func (pot PotionType) String() string {
// 	switch pot {
// 	case SMALL:
// 		return "Small"
// 	case MEDIUM:
// return "Med"
// 	case LARGE:

// 	default:
// 		logging.LogError(logger, "Something went wrong while attemping to return the represensation of an itemType")
// 		return "[ERROR]"
// 	}
// }
