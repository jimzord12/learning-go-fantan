package rngloot

import (
	"math/rand"

	rpgdungeon "github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models/rpg-dungeon"
	itemspkg "github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models/rpg-items"
	typeconvpkg "github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models/rpg-types-conv"
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/logging"
)

type Item = itemspkg.Item

var DifficultyToMaterial = typeconvpkg.DifficultyToMaterial
var MaterialToDifficulty = typeconvpkg.MaterialToDifficulty

func getMaterialDropChance() (Material, bool) {
	var luck = rand.Intn(101)
	var dungeon = rpgdungeon.ActiveDungeon

	if luck >= 95 {
		return DifficultyToMaterial(dungeon.Difficulty + 2), true
	} else if luck >= 70 && luck < 95 {
		return DifficultyToMaterial(dungeon.Difficulty + 1), true
	} else if luck >= 20 && luck < 70 {
		return DifficultyToMaterial(dungeon.Difficulty), true
	} else {
		return -1, false
	}
}

func getPotionDropChance() (PotionType, bool) {
	var luck = rand.Intn(101)

	if luck >= 95 {
		return itemspkg.LARGE, true
	} else if luck >= 70 && luck < 95 {
		return itemspkg.MEDIUM, true
	} else if luck >= 20 && luck < 70 {
		return itemspkg.SMALL, true
	} else {
		return -1, false
	}
}

func getEquipmentTypeDropChance() ItemType {
	var luck = rand.Intn(3)

	switch luck {
	case 0:
		return itemspkg.WEAPON
	case 1:
		return itemspkg.ARMOR
	case 2:
		return itemspkg.ACCESORY
	default:
		return -1
	}
}

func getWeaponTypeDropChance() itemspkg.WeaponType {
	var luck = rand.Intn(5)

	switch luck {
	case 0:
		return itemspkg.SWORD
	case 1:
		return itemspkg.DAGGER
	case 2:
		return itemspkg.GREATSWORD
	case 3:
		return itemspkg.MACE
	case 4:
		return itemspkg.SPEAR
	default:
		logging.LogError(logging.Logger, "In 'getWeaponTypeDropChance()' luck is -1, but this should not be possible")
		return -1
	}
}

// False => No Drops
// True => You must check if struct field are not -1
// (-1) means No Drop for that Field
func CalcDrops() (EnemyDrops, bool) {
	equipment, hasEquipDrop := getMaterialDropChance()
	potion, hasPotDrop := getPotionDropChance()

	if hasEquipDrop || hasPotDrop {

		return EnemyDrops{
			EquipmentType:     getEquipmentTypeDropChance(),
			EquipmentMaterial: equipment,
			PotionDrop:        potion,
		}, true
	}

	return EnemyDrops{}, false

}

func (loot EnemyDrops) GetLoot() []*Item {
	var drops []*Item

	if loot.EquipmentMaterial != -1 {
		switch loot.EquipmentType {
		case itemspkg.WEAPON:
			weaponType := getWeaponTypeDropChance()
			drops = append(drops, itemspkg.NewWeapon(weaponType, loot.EquipmentMaterial))
		case itemspkg.ARMOR:
			drops = append(drops, itemspkg.NewArmor(loot.EquipmentMaterial))
		case itemspkg.ACCESORY:
			drops = append(drops, itemspkg.NewAccessory(loot.EquipmentMaterial))
		default:
			logging.LogError(itemspkg.Logger, "While getting loot, signature: 'func (loot EnemyDrops) GetLoot() [2]*Item'")
			panic("[ERROR]: getting loot")
		}
	}

	if loot.PotionDrop != -1 {
		drops = append(drops, itemspkg.NewPotion(loot.PotionDrop))
	}

	return drops
}
