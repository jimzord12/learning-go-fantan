package rpgmodels

import (
	"errors"
	"fmt"
	"log"

	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpghelpers"
)

type ItemType int

func (i ItemType) String() string {
	switch i {
	case WEAPON:
		return "Weapon"
	case ARMOR:
		return "Armor"
	case POTION:
		return "Potion"
	default:
		return ""
	}
}

func (it Item) String() string {
	return it.Name
}

const (
	WEAPON ItemType = iota
	ARMOR
	POTION
)

type Inventory struct {
	Size      int
	Weight    float64 // This is the Current Weight, is used for caching
	MaxWeight float64
	Items     []Item
}

type Item struct {
	Name   string
	Weight float64
	ItemType
	Value float64
}

/// METHODS ///

func (inv *Inventory) Add(item *Item) error {
	newSize := inv.Size + 1
	newWeight := inv.Weight + item.Weight

	if newSize > inv.Size {
		return fmt.Errorf("[ERROR]: Inventory is full (no more slots)")
	}

	if newWeight > inv.MaxWeight {
		return fmt.Errorf("[ERROR]: Inventory is too heavy (no more slots)")
	}

	inv.Items = append(inv.Items, *item)
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

func NewItem(name string, weight float64, itemType ItemType, value float64) *Item {
	return &Item{
		Name:     name,
		Weight:   weight,
		ItemType: itemType,
		Value:    value,
	}
}
