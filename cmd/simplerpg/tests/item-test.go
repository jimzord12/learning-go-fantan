package tests

import (
	"fmt"

	rpgitems "github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models/rpg-items"
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/itemhelpers"
)

// Aliases for better readability
type Item = rpgitems.Item

var NewWeapon = rpgitems.NewWeapon
var NewArmor = rpgitems.NewArmor
var NewPotion = rpgitems.NewPotion

func TestWeapons() {
	// Create All Available Weapons
	var allWeapons []*Item

	for weaponType := range rpgitems.WeaponTypes {
		for materialType := range rpgitems.MaterialTypes {
			allWeapons = append(allWeapons, NewWeapon(weaponType, materialType))
		}
	}

	itemhelpers.SortByType(allWeapons)
	for idx, weapon := range allWeapons {
		fmt.Printf("Weapon: (#%d): %+v\n", idx, *weapon)
	}

	fmt.Println("========================================")
	fmt.Println("=============== MATERIAL ===============")

	itemhelpers.SortByMaterial(allWeapons)
	for idx, weapon := range allWeapons {
		fmt.Printf("Weapon: (#%d): %+v\n", idx, *weapon)
	}

	fmt.Println("========================================")
	fmt.Println("================ VALUE ================")

	itemhelpers.SortByValue(allWeapons)
	for idx, weapon := range allWeapons {
		fmt.Printf("Weapon: (#%d): %+v\n", idx, *weapon)
	}
}

func TestArmors() {
	var allArmors []*Item

	for material := range rpgitems.MaterialTypes {
		allArmors = append(allArmors, NewArmor(material))
	}

	itemhelpers.SortByType(allArmors)
	for idx, armor := range allArmors {
		fmt.Printf("Armor: (#%d): %+v\n", idx, *armor)
	}

	fmt.Println("========================================")
	fmt.Println("=============== MATERIAL ===============")

	itemhelpers.SortByMaterial(allArmors)
	for idx, armor := range allArmors {
		fmt.Printf("Armor: (#%d): %+v\n", idx, *armor)
	}

	fmt.Println("========================================")
	fmt.Println("================ VALUE ================")

	itemhelpers.SortByValue(allArmors)
	for idx, armor := range allArmors {
		fmt.Printf("Armor: (#%d): %+v\n", idx, *armor)
	}
}

func TestPotions() {
	var allPotions []*Item
	for potionType := range rpgitems.PotionTypes {
		allPotions = append(allPotions, NewPotion(potionType))
	}

	itemhelpers.SortByType(allPotions)
	for idx, potion := range allPotions {
		fmt.Printf("Potion: (#%d): %+v\n", idx, *potion)
	}

	fmt.Println("========================================")
	fmt.Println("================ VALUE ================")

	itemhelpers.SortByValue(allPotions)
	for idx, potion := range allPotions {
		fmt.Printf("Potion: (#%d): %+v\n", idx, *potion)
	}

}
