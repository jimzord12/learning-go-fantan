package models

import (
	"fmt"
	"log"

	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/generalhelpers"
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/logging"
)

////////////////////////////////////////////////////////////////////////////////////
/////////////////////////// CHARACTER CONSTRUCTORS /////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////

func newCharacter(id string, name string, race CharacterType, baseStats BaseStats, hp float64, stm float64, lvl int) *Character {
	return &Character{
		ID:   id,
		Name: name,
		Type: race,
		Stats: BaseStats{
			MaxHp:       baseStats.MaxHp,
			MaxStamina:  baseStats.MaxStamina,
			MaxWeight:   baseStats.MaxWeight,
			StmRecovery: baseStats.StmRecovery,
			CritStrike:  baseStats.CritStrike,
			DodgeChance: baseStats.DodgeChance,
		},
		Hp:      hp,
		Stamina: stm,
		Level:   lvl,
	}
}

func NewPlayer(id string, name string, race CharacterType) *Character {
	var baseStats BaseStats = race.GetBaseStats()

	return newCharacter(id, name, race, baseStats, baseStats.MaxHp, baseStats.MaxStamina, 1)
}

func newEnemy(id string, name string, monsterType CharacterType, lvl int) *Character {
	var baseStats BaseStats = monsterType.GetBaseStats()

	return newCharacter(id, name, monsterType, baseStats, baseStats.MaxHp, baseStats.MaxStamina, lvl)
}

// Simple Enemies
func NewSpiderEnemy(id string, lvl int) *Character {
	// Creating the Enemy
	enemy := newEnemy(id, "Spider", SIMPLE, lvl)

	// Creating & Equiping them with a random Weapon
	enemy.EnemyEquipRandWeapon(ActiveDungeon.Difficulty)

	// Creating & Equiping them with Armor
	enemy.EnemyEquipArmor(ActiveDungeon.Difficulty)

	return enemy
}

func NewSlimeEnemy(id string, lvl int) *Character {
	enemy := newEnemy(id, "Slime", SIMPLE, lvl)
	enemy.EnemyEquipRandWeapon(ActiveDungeon.Difficulty)
	enemy.EnemyEquipArmor(ActiveDungeon.Difficulty)

	return enemy
}
func NewRatEnemy(id string, lvl int) *Character {
	enemy := newEnemy(id, "Rat", SIMPLE, lvl)
	enemy.EnemyEquipRandWeapon(ActiveDungeon.Difficulty)
	enemy.EnemyEquipArmor(ActiveDungeon.Difficulty)

	return enemy
}

func NewMushroomEnemy(id string, lvl int) *Character {
	enemy := newEnemy(id, "Mushroom", SIMPLE, lvl)
	enemy.EnemyEquipRandWeapon(ActiveDungeon.Difficulty)
	enemy.EnemyEquipArmor(ActiveDungeon.Difficulty)

	return enemy
}

// Elite Enemies
func NewGoblinEnemy(id string, lvl int) *Character {
	enemy := newEnemy(id, "Goblin", ELITE, lvl)
	enemy.EnemyEquipRandWeapon(ActiveDungeon.Difficulty)
	enemy.EnemyEquipArmor(ActiveDungeon.Difficulty)

	return enemy
}

func NewKoboldEnemy(id string, lvl int) *Character {
	enemy := newEnemy(id, "Kobold", ELITE, lvl)
	enemy.EnemyEquipRandWeapon(ActiveDungeon.Difficulty)
	enemy.EnemyEquipArmor(ActiveDungeon.Difficulty)

	return enemy
}

func NewImpEnemy(id string, lvl int) *Character {
	enemy := newEnemy(id, "Imp", ELITE, lvl)
	enemy.EnemyEquipRandWeapon(ActiveDungeon.Difficulty)
	enemy.EnemyEquipArmor(ActiveDungeon.Difficulty)

	return enemy
}

// Boss Enemies
func NewDragonEnemy(id string, lvl int) *Character {
	enemy := newEnemy(id, "Dragon", BOSS, lvl)
	enemy.EnemyEquipRandWeapon(ActiveDungeon.Difficulty)
	enemy.EnemyEquipArmor(ActiveDungeon.Difficulty)

	return enemy
}

func NewArcDemonEnemy(id string, lvl int) *Character {
	enemy := newEnemy(id, "Arc Demon", BOSS, lvl)
	enemy.EnemyEquipRandWeapon(ActiveDungeon.Difficulty)
	enemy.EnemyEquipArmor(ActiveDungeon.Difficulty)

	return enemy
}

////////////////////////////////////////////////////////////////////////////////////
///////////////////////////// DUNGEON CONSTRUCTORS /////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////

var ActiveDungeon *Dungeon

func DungeonInit(diff Difficulty) {
	ActiveDungeon = &Dungeon{
		Difficulty: diff,
	}
}

////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////// ITEM CONSTRUCTORS /////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////

func newItem(name string, weight float64, itemType ItemType, value float64, material Material) *Item {
	return &Item{
		Name:     name,
		Weight:   generalhelpers.Round2Dec(weight),
		ItemType: itemType,
		Value:    generalhelpers.Round2Dec(value),
		Material: material,
	}
}

// Think of Material as Rarity
func NewWeapon(weaponType WeaponType, material Material) *Item {
	var weight float64
	var value float64
	var name string

	switch weaponType {
	case SWORD:
		name = material.String() + " Sword"
		weight = material.GetWeight(SWORD_VOLUME)
		value = weaponType.GetDmg(weight, material)
	case GREATSWORD:
		name = material.String() + " Greatsword"
		weight = material.GetWeight(GREATSWORD_VOLUME)
		value = weaponType.GetDmg(weight, material)
	case DAGGER:
		name = material.String() + " Dagger"
		weight = material.GetWeight(DAGGER_VOLUME)
		value = weaponType.GetDmg(weight, material)
	case MACE:
		name = material.String() + " Mace"
		weight = material.GetWeight(MACE_VOLUME)
		value = weaponType.GetDmg(weight, material)
	case SPEAR:
		name = material.String() + " Spear"
		weight = material.GetWeight(SPEAR_VOLUME)
		value = weaponType.GetDmg(weight, material)
	default:
		logging.LogError(logging.Logger, "The params:")
		fmt.Println(weaponType, material)
		log.Fatalln("[ERROR]: something went wrong while creating a new weapon")
	}

	return newItem(name, weight, WEAPON, value, material)

}

// I made Armor simpler than the Weapons
func NewArmor(material Material) *Item {
	var weight float64
	var value float64
	var name string

	weight = material.GetWeight(ARMOR_VOLUME)
	value = material.GetToughness() * (ARMOR_BASE_DEF + (weight * ARMOR_BASE_MULTI_PER_KG))
	name = material.String() + " Armor"

	return newItem(name, weight, ARMOR, value, material)
}

func NewAccessory(material Material) *Item {
	var weight float64
	var value float64
	var name string

	weight = material.GetWeight(ACC_VOLUME)
	value = material.GetToughness() * ACC_BASE_VALUE
	name = material.String() + " Accessory"

	return newItem(name, weight, ACCESORY, value, material)
}

func NewPotion(size PotionType) *Item {
	weight := float64(size) * POTION_BASE_WEIGHT
	value := float64(size) * POTION_BASE_VALUE
	var name string

	switch size {
	case SMALL:
		name = PotionTypes[1]
	case MEDIUM:
		name = PotionTypes[2]
	case LARGE:
		name = PotionTypes[3]
	default:
		log.Fatalln("[ERROR]: something went wrong while creating a new potion ")
	}

	return newItem(name, weight, POTION, value, NO_MATERIAL)
}