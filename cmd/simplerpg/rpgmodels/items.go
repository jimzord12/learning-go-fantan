package rpgmodels

import (
	"errors"
	"fmt"
	"log"

	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/global"
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpghelpers"
)

type ItemType int
type PotionSize int
type Material int
type WeaponType int

func (wt WeaponType) GetDmg(weight float64) float64 {
	return global.BASE_DMG + (weight * global.BASE_WEAPON_MULTI_PER_KG)
}

// This could be stored as a const, but its not for practise purposes
func (mat Material) GetWeight(volume float64) float64 {
	switch mat {
	case WOOD:
		return volume * global.WOOD_DENSITY
	case BRONZE:
		return volume * global.BRONZE_DENSITY
	case IRON:
		return volume * global.IRON_DENSITY
	case STEEL:
		return volume * global.STEEL_DENSITY
	case TITANIUM:
		return volume * global.TITANIUM_DENSITY
	case MYTHRIL:
		return volume * global.MYTHRIL_DENSITY
	default:
		log.Fatalf("[ERROR]: Something went wrong while attemping to calculate this material's (%v) weight\n", mat)
		return -1
	}
}

func (mat Material) String() string {
	switch mat {
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
		log.Fatalf("[ERROR]: Something went wrong while attemping to return the represensation of an material \n", mat)
		return "[ERROR]"
	}
}

func (i ItemType) String() string {
	switch i {
	case WEAPON:
		return "Weapon"
	case ARMOR:
		return "Armor"
	case POTION:
		return "Potion"
	case MISC:
		return "Misc"
	default:
		log.Fatalln("[ERROR]: Something went wrong while attemping to return the represensation of an itemType")
		return "[ERROR]"
	}
}

func (it Item) String() string {
	return it.Name
}

const (
	WEAPON ItemType = iota
	ARMOR
	POTION
	MISC
)

const (
	WOOD Material = iota + 1
	BRONZE
	IRON
	STEEL
	TITANIUM
	MYTHRIL
)

const (
	SWORD WeaponType = iota
	GREATSWORD
	DAGGER
	MACE
	SPEAR
)

const (
	SMALL PotionSize = iota + 1
	MEDIUM
	LARGE
)

type Inventory struct {
	Size      int
	Weight    float64 // This is the Current Weight, is used for caching
	MaxWeight float64
	Items     []*Item
}

type Item struct {
	Name   string
	Weight float64
	ItemType
	Material
	Value float64
}

// / METHODS ///
func (inv *Inventory) Add(item *Item) error {
	newSize := inv.Size + 1
	newWeight := inv.Weight + item.Weight

	if newSize > inv.Size {
		return fmt.Errorf("[ERROR]: Inventory is full (no more slots)")
	}

	if newWeight > inv.MaxWeight {
		return fmt.Errorf("[ERROR]: Inventory is too heavy (no more slots)")
	}

	inv.Items = append(inv.Items, item)
	inv.Weight += item.Weight
	inv.Size += 1

	return nil
}

func (inv *Inventory) Remove(item *Item) error {
	// 1. Find the Item in the Inventory
	index, err := inv.FindItemIndex(item.Name)
	if err != nil {
		log.Println(err)
		return err
	}

	// 2. Remove the Item from the Inventory
	rpghelpers.RemoveFromSlice(inv.Items, index)

	// 3. Update the Inventory stats
	inv.Weight -= item.Weight
	inv.Size -= 1
	return nil
}

func (it Inventory) FindItemIndex(desiredItemName string) (int, error) {
	for i, v := range it.Items {
		if v.Name == desiredItemName {
			return i, nil
		}
	}

	return -1, errors.New("[ERROR]: Could not find Index")
}

func newItem(name string, weight float64, itemType ItemType, value float64) *Item {
	return &Item{
		Name:     name,
		Weight:   weight,
		ItemType: itemType,
		Value:    value,
	}
}

// Think of Material as Rarity
func NewWeapon(weaponType WeaponType, material Material) *Item {
	var weight float64
	var value float64
	var name string

	switch weaponType {
	case SWORD:
		name = material.String() + "Sword"
		weight = material.GetWeight(global.SWORD_VOLUME)
		value = weaponType.GetDmg(weight)
	case GREATSWORD:
		name = material.String() + "Greatsword"
		weight = material.GetWeight(global.GREATSWORD_VOLUME)
		value = weaponType.GetDmg(weight)
	case DAGGER:
		name = material.String() + "Dagger"
		weight = material.GetWeight(global.DAGGER_VOLUME)
		value = weaponType.GetDmg(weight)
	case MACE:
		name = material.String() + "Mace"
		weight = material.GetWeight(global.MACE_VOLUME)
		value = weaponType.GetDmg(weight)
	case SPEAR:
		name = material.String() + "Spear"
		weight = material.GetWeight(global.SPEAR_VOLUME)
		value = weaponType.GetDmg(weight)
	default:
		log.Fatalln("[ERROR]: something went wrong while creating a new weapon")
	}

	return newItem(name, weight, WEAPON, value)

}

// I made Armor simpler than the Weapons
func NewArmor(material Material) *Item {
	var weight float64
	var value float64
	var name string

	weight = material.GetWeight(global.ARMOR_VOLUME)
	value = global.BASE_DEF + (weight * global.BASE_ARMOR_MULTI_PER_KG)
	name = material.String() + "Armor"

	return newItem(name, weight, ARMOR, value)
}

func NewPotion(size PotionSize) *Item {
	weight := float64(size) * global.POTION_BASE_WEIGHT
	value := float64(size) * global.POTION_BASE_VALUE
	var name string

	switch size {
	case SMALL:
		name = "Small HP Potion"
	case MEDIUM:
		name = "Medium HP Potion"
	case LARGE:
		name = "Large HP Potion"
	default:
		log.Fatalln("[ERROR]: something went wrong while creating a new potion ")
	}

	return newItem(name, weight, POTION, value)
}
