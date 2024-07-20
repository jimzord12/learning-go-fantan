package main

import (
	"fmt"

	"github.com/jimzord12/learning-go-fantan/cmd/fantan"
)

func main() {
	fantan.PlayFantan(500)

	fmt.Print("Running main.go\n\n")

	// functions.Main()
	// structs.Main()
	// mytypes.Main()
	// mytypes.Main2() // Type Assertion
	// myinterfaces.Main()
	// whatever.Main()
	// goroutines.Main1()
	// goroutines.Main2()
	// flags.Main()
	// env.Main()
	// simplewebserver.SimpleWebServer()
	// simplechiserver.Main()
	// simplerpg.RunGame()

	// simpleRPG()

}

// func fantanWrapper() {
// 	userBalance := helper.GetAndParseUserInputInt("How much money do you wish to DEPOSIT?")
// 	fantan.PlayFantan(float64(userBalance))

// }

// func simpleRPG() {
// 	// 1. Create a Player
// 	player := rpgmodels.NewCharacter("Player", 150, 1)
// 	fmt.Println(player)

// 	// 2. Create an Enemy
// 	enemy := rpgmodels.NewCharacter("Slime", 50, 1)
// 	fmt.Println(enemy)

// 	// 3. Create a Sword
// 	sword := rpgmodels.NewItem("Rusty Sword", 10.5, rpgmodels.WEAPON, 23.5)
// 	fmt.Println(sword.Name)

// 	// 4. Create Armor
// 	armor := rpgmodels.NewItem("Old Rags", 3.2, rpgmodels.ARMOR, 5)
// 	fmt.Println(armor.Name)

// 	// 	// 5. Equip Weapon & Armor
// 	player.Equip(sword)
// 	player.Equip(armor)

// 	player.DisplayInventory()
// 	player.DisplayEquipment()

// 	// 6. Attack Enemy
// 	enemy.DisplayHp()
// 	fmt.Println("\nPlayer attacking Slime...")
// 	player.Attack(enemy)
// 	enemy.DisplayHp()
// }
