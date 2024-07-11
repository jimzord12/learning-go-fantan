package rpgmodels

import (
	"fmt"
)

type Character struct {
	Name      string
	Hp        float64
	Level     int
	Exp       float64
	Inventory Inventory
	Equipment Equipment
}

type Equipment struct {
	WeaponSlot *Item
	ArmorSlot  *Item
}

func NewCharacter(name string, hp float64, lvl int) *Character {
	return &Character{
		Name:  name,
		Hp:    hp,
		Level: lvl,
	}
}

func (char *Character) DisplayInventory() {
	if len(char.Inventory.Items) == 0 {
		fmt.Println("Inventory is Empty")
	}
	for i, v := range char.Inventory.Items {
		fmt.Println("1.", i, ":", v)
	}
}

func (char *Character) DisplayEquipment() {
	fmt.Println("Weapon:", char.Equipment.WeaponSlot, "\nArmor:", char.Equipment.ArmorSlot)
}

func (char *Character) DisplayStats() {
	fmt.Println("Name:", char.Name)
	fmt.Println("HP:", char.Hp)
	fmt.Println("Level:", char.Level)
	fmt.Println("EXP:", char.Exp)
}

func (char *Character) DisplayHp() {
	fmt.Printf("[%s], HP: %.2f", char.Name, char.Hp)
}

func (char *Character) Attack(enemy *Character) error {
	//TODO: When Killing an Enemy increase EXP
	equippedWeapon := char.Equipment.WeaponSlot
	if equippedWeapon == nil {
		return fmt.Errorf("[ERROR]: You cannot attack, you are NOT holding a Weapon")

	}

	if char.Equipment.WeaponSlot.ItemType != WEAPON {
		return fmt.Errorf("[ERROR]: You cannot attack with (%s), its NOT a Weapon", equippedWeapon)
	}

	enemy.Hp -= equippedWeapon.Value
	return nil
}

func (char *Character) Unequip(item *Item) error {
	if item.ItemType != WEAPON && item.ItemType != ARMOR {
		return fmt.Errorf("[ERROR]: You cannot equip this item type (%s)", item.ItemType)
	}

	char.Inventory.Add(item)

	switch item.ItemType {
	case WEAPON:
		char.Equipment.WeaponSlot = nil
	case ARMOR:
		char.Equipment.ArmorSlot = nil
	}

	return nil
}

func (char *Character) Equip(item *Item) error {
	// 1. Check if the Item is of Type WEAPON or ARMOR
	if item.ItemType != WEAPON && item.ItemType != ARMOR {
		return fmt.Errorf("[ERROR]: You cannot equip this item type (%s)", item)
	}

	// 2. Check if the Player already has an Item of the same type equipped
	switch item.ItemType {
	case WEAPON:
		// Player already has a Weapon equipped
		if char.Equipment.WeaponSlot != nil {
			char.Unequip(char.Equipment.WeaponSlot) // Unequip Current Weapon
			char.Equipment.WeaponSlot = item        // Equip Selected Weapon
		} else {
			char.Equipment.WeaponSlot = item // Equip Selected Weapon
		}
	case ARMOR:
		if char.Equipment.ArmorSlot != nil {
			char.Unequip(char.Equipment.ArmorSlot)
			char.Equipment.ArmorSlot = item
		} else {
			char.Equipment.ArmorSlot = item
		}
	default:
		return fmt.Errorf("[ERROR]: Something went wrong while trying to EQUIP the item: (%v)", *item)
	}

	return nil
}

func (char *Character) UseItem(item *Item) error {
	switch item.ItemType {
	case WEAPON, ARMOR:
		char.Equip(item)
	case POTION:
		char.Hp += item.Value
	default:
		return fmt.Errorf("[ERROR]: This item is not supported yet: (%v)", *item)
	}

	return nil
}
