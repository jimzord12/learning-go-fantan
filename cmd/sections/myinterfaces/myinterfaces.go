package myinterfaces

import "fmt"

type ItemType int

const (
	WEAPON ItemType = iota
	ARMOR
	POTION
)

/// RPG EXAMPLE ///
type Character struct {
	Name      string
	Hp        float64
	Level     int
	Exp       float64
	Inventory Inventory
}

type Inventory struct {
	Size      int
	Weight    float64
	MaxWeight float64
	Items     []Item
}

type Item struct {
	Name   string
	Weight float64
	ItemType
	value float64
}

func (char *Character) PickUp(item *Item) error {
	var newWeight float64 = char.Inventory.Weight + item.Weight

	if len(char.Inventory.Items) >= char.Inventory.Size {
		return fmt.Errorf("[ERROR]: Your Inventory has not available Slots")
	}

	if newWeight > char.Inventory.MaxWeight {
		return fmt.Errorf("[ERROR]: Your Inventory is too heavy")
	}

	fmt.Println("++++++++++++++")
	char.Inventory.Items = append(char.Inventory.Items, *item)
	return nil
}

func (char *Character) DisplayInventory() {
	for i, v := range char.Inventory.Items {
		fmt.Println("1.", i, ":", v)
	}
}

func (char *Character) Attack(weapon Item, enemy *Character) error {
	if weapon.ItemType != WEAPON {
		return fmt.Errorf("[ERROR]: (%s) in NOT a Weapon", weapon.Name)
	}

	enemy.Hp -= weapon.value
	return nil

}

func (item *Item) Use() {
	switch item.ItemType {
	case WEAPON:
		fmt.Println("Using Weapon")
	case ARMOR:
		fmt.Println("Using Armor")
	case POTION:
		fmt.Println("Using Potion")
	}
}

func Main() {
	// 1. Creating Player
	player := Character{
		Name:  "Player",
		Hp:    100.0,
		Level: 1,
		Inventory: Inventory{
			Size:      10,
			MaxWeight: 120.0,
		},
	}

	// 2. Creating Enemy
	enemy := Character{
		Name:  "Enemy",
		Hp:    300.0,
		Level: 3,
	}

	// 3. Creating a Weapon, a Weapon
	sword := Item{
		Name:     "Emerald-Sword",
		Weight:   52.5,
		ItemType: WEAPON,
		value:    220,
	}

	// 4. Picking Up the Sword
	player.PickUp(&sword)

	// 5. Displaying Inventory
	player.DisplayInventory()

	// 6. Using the Sword
	player.Inventory.Items[0].Use()

	// 7. Attacking the Enemy
	fmt.Println("Enemy HP BEFORE:", enemy.Hp)
	player.Attack(player.Inventory.Items[0], &enemy)
	fmt.Println("Enemy HP AFTER:", enemy.Hp)

}
