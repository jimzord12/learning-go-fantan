package rpgitems

import "log"

type ItemType int
type PotionType int
type Material int
type WeaponType int

// Global logger variable
var Logger *log.Logger

const (
	WEAPON ItemType = iota
	ARMOR
	POTION
	MISC
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
