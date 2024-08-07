package tests

import (
	"fmt"
	"log"

	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models"
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/itemhelpers"
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/logging"
)

// Aliases for better readability
type Item = models.Item

var NewWeapon = models.NewWeapon
var NewArmor = models.NewArmor
var NewPotion = models.NewPotion
var NewAccessory = models.NewAccessory

func TestWeapons() {
	// Create All Available Weapons
	var allWeapons []*Item

	for weaponType := range models.WeaponTypes {
		for materialType := range models.MaterialTypes {
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

	for material := range models.MaterialTypes {
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
	for potionType := range models.PotionTypesToNames {
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

func TestAccessories(args ...any) {
	// player, ok := args[0].(*Character)

	// if !ok {
	// 	logging.LogError(logging.Logger, "(func GivePotions(args ...any)) -> Issue with type assertion")
	// }

	var allAccessories []*Item
	for accessory := range models.MaterialTypes {
		allAccessories = append(allAccessories, NewAccessory(accessory))
	}

	itemhelpers.SortByType(allAccessories)
	for idx, accessory := range allAccessories {
		fmt.Printf("Accessory: (#%d): %+v\n", idx, *accessory)
	}

	fmt.Println("========================================")
	fmt.Println("================ VALUE ================")

	itemhelpers.SortByValue(allAccessories)
	for idx, accessory := range allAccessories {
		convertedValue, err := accessory.GetAccessoryStats()
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("Accessory: (#%d): %+v | Final Value: %.02f\n", idx, *accessory, convertedValue)
	}

}

func GivePotions(args ...any) {
	player, ok := args[0].(*Character)

	if !ok {
		logging.LogError(logging.Logger, "(func GivePotions(args ...any)) -> Issue with type assertion")
	}

	for i := 0; i < 3; i++ {
		potion_S := models.NewPotion(models.SMALL)
		potion_M := models.NewPotion(models.MEDIUM)
		potion_L := models.NewPotion(models.LARGE)

		player.MoveToInventory(potion_S)
		player.MoveToInventory(potion_M)
		player.MoveToInventory(potion_L)

	}
	fmt.Println("")
	player.DisplayInventory()
	fmt.Println("")
}
