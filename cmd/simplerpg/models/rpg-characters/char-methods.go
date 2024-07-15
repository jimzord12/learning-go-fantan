package rpgcharacters

import (
	"fmt"
	"log"

	rpgitems "github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models/rpg-items"
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/generalhelpers"
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/logging"
)

type Item = rpgitems.Item

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
	// TODO: Decrease Character's Stamina based on another Func
	var dmg float64

	switch char.Type {
	case HUMAN, ELF, DWARF:
		equippedWeapon := char.Equipment.WeaponSlot
		if equippedWeapon == nil {
			return fmt.Errorf("[ERROR]: You cannot attack, you are NOT holding a Weapon")

		}

		if char.Equipment.WeaponSlot.ItemType != rpgitems.WEAPON {
			return fmt.Errorf("[ERROR]: You cannot attack with (%s), its NOT a Weapon", equippedWeapon.Name)
		}
		dmg = equippedWeapon.Value

	case SIMPLE:

	case ELITE:

	case BOSS:

	default:
	}

	enemy.Hp -= dmg
	return nil
}

func (char *Character) Unequip(item *Item) error {
	if item.ItemType != rpgitems.WEAPON && item.ItemType != rpgitems.ARMOR {
		return fmt.Errorf("[ERROR]: You cannot equip this item type (%s)", item.ItemType)
	}

	char.TakeItem(item)

	switch item.ItemType {
	case rpgitems.WEAPON:
		char.Equipment.WeaponSlot = nil
	case rpgitems.ARMOR:
		char.Equipment.ArmorSlot = nil
	}

	return nil
}

func (char *Character) Equip(item *Item) error {
	// 1. Check if the Item is of Type WEAPON or ARMOR
	if item.ItemType != rpgitems.WEAPON && item.ItemType != rpgitems.ARMOR {
		return fmt.Errorf("[ERROR]: You cannot equip this item type (%s)", item.Name)
	}

	// 2. Check if the Player already has an Item of the same type equipped
	switch item.ItemType {
	case rpgitems.WEAPON:
		// Player already has a Weapon equipped
		if char.Equipment.WeaponSlot != nil {
			char.Unequip(char.Equipment.WeaponSlot) // Unequip Current Weapon
			char.Equipment.WeaponSlot = item        // Equip Selected Weapon
		} else {
			char.Equipment.WeaponSlot = item // Equip Selected Weapon
		}
	case rpgitems.ARMOR:
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
	case rpgitems.WEAPON, rpgitems.ARMOR:
		char.Equip(item)
	case rpgitems.POTION:
		char.Hp += item.Value
	default:
		return fmt.Errorf("[ERROR]: This item is not supported yet: (%v)", *item)
	}

	return nil
}

func (char *Character) Rest() error {
	if char.Stamina == char.Stats.MaxStamina {
		return fmt.Errorf("[ERROR]: You can NOT rest, your Stamina is full")
	}

	newStamina := char.Stamina + char.Stats.StmRecovery

	if newStamina > char.Stats.MaxStamina {
		char.Stamina = char.Stats.MaxStamina
		return nil
	}

	char.Stamina += char.Stats.StmRecovery
	return nil
}

func (race CharacterType) GetBaseStats() BaseStats {
	switch race {
	case HUMAN:
		return HumanBaseStats
	case ELF:
		return ElfBaseStats
	case DWARF:
		return DwarfBaseStats
	case SIMPLE, ELITE, BOSS:
		logging.LogError(logging.Logger, "You called GetBaseStats() on a MONSTER CharacterRace type. Don't do that.")
		panic("[ERROR]: At char-types.go, func signature: 'func (race CharacterRace) GetBaseStats() BaseStats'")
	default:
		logging.LogError(logging.Logger, "Something went wrong while getting the base stats from a character race")
		panic("[ERROR]: At char-types.go, func signature: 'func (race CharacterRace) GetBaseStats() BaseStats'")
	}
}

func (char *Character) TakeItem(item *Item) error {
	newSize := char.Inventory.Size + 1
	newWeight := char.Weight + item.Weight

	if newSize > char.Inventory.Size {
		return fmt.Errorf("[ERROR]: Character's Inventory is full (no more slots)")
	}

	if newWeight > char.Stats.MaxWeight {
		return fmt.Errorf("[ERROR]: Character's Inventory is too heavy (no more slots)")
	}

	char.Inventory.Items = append(char.Inventory.Items, item)
	char.Weight += item.Weight
	char.Inventory.Size += 1

	return nil
}

func (char *Character) RemoveItem(item *Item) error {
	// 1. Find the Item in the Character's Inventory
	index, err := char.Inventory.FindItemIndex(item.Name)
	if err != nil {
		log.Println(err)
		return err
	}

	// 2. Remove the Item from the Character's Inventory
	generalhelpers.RemoveFromSlice(char.Inventory.Items, index)

	// 3. Update the Character's Inventory stats
	char.Weight -= item.Weight
	char.Inventory.Size -= 1
	return nil
}
