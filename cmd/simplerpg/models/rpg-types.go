package models

import (
	"log"
)

///////////////////////////////// CHARACTER TYPES /////////////////////////////////

type Character struct {
	ID        string
	Name      string
	Type      CharacterType
	Stats     BaseStats
	Hp        float64
	Stamina   float64
	Weight    float64
	Level     int
	Exp       float64
	Inventory Inventory
	Equipment Equipment
}

type BaseStats struct {
	MaxHp       float64
	MaxStamina  float64
	MaxWeight   float64
	StmRecovery float64
	CritStrike  float64
	DodgeChance float64
}

type CharacterType int

const (
	// Player Char Type (Races)
	HUMAN CharacterType = iota
	ELF
	DWARF
	// Below are the Monsters' Types
	SIMPLE
	ELITE
	BOSS
)

// Created this Var to make single point from
// where to check if a Character Type is an Player type or not
var PlayerTypes = []CharacterType{
	HUMAN,
	ELF,
	DWARF,
}

///////////////////////////////// DUNGEON TYPES /////////////////////////////////

type Difficulty int

const (
	WOOD_DIF Difficulty = iota + 1
	BRONZE_DIF
	IRON_DIF
	STEEL_DIF
	TITANIUM_DIF
	MYTHRIL_DIF
)

type Dungeon struct {
	Difficulty
}

///////////////////////////////// INVENTORY TYPES /////////////////////////////////

type Inventory struct {
	MaxSize int
	Items   []*Item
}

type Equipment struct {
	WeaponSlot    *Item
	ArmorSlot     *Item
	AccessorySlot *Item
}

///////////////////////////////// ITEM TYPES /////////////////////////////////

type ItemType int
type PotionType int
type Material int
type WeaponType int

// Global logger variable
var Logger *log.Logger

// When Adding a New ItemType
// Check if the Character type's fields or methods need update
const (
	WEAPON ItemType = iota
	ARMOR
	ACCESSORY
	POTION
)

const (
	SWORD WeaponType = iota
	GREATSWORD
	DAGGER
	MACE
	SPEAR
)

const (
	NO_MATERIAL Material = iota
	WOOD
	BRONZE
	IRON
	STEEL
	TITANIUM
	MYTHRIL
)

const (
	SMALL PotionType = iota + 1
	MEDIUM
	LARGE
)

type Item struct {
	ID     string
	Name   string
	Weight float64
	ItemType
	Material
	Value float64
}

///////////////////////////////// LOOT TYPES /////////////////////////////////

type EnemyDrops struct {
	EquipmentType     ItemType   // 0-2
	EquipmentMaterial Material   // 0-100
	PotionDrop        PotionType // 0-100

}

///////////////////////////////// BATTLE TYPES /////////////////////////////////

type Battle struct {
	ID           string
	Player       *Character
	Enemy        *Character
	BattleRounds []BattleRound
}

type BattleRound struct {
	ID             string
	Attacker       *Character
	Defender       *Character
	AttackerAction BattleAction
	DefenderAction BattleAction
	Consumable     *Item
}

type BattleAction int

const (
	LIGHT_ATTACK BattleAction = iota + 1
	HEAVY_ATTACK
	DEFEND
	REST
	HEAL
)

type EnemyBattlePattern []BattleAction
