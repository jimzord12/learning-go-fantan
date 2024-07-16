package tests

import (
	"fmt"

	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models"
)

type Character = models.Character

var NewPlayer = models.NewPlayer

func CreatePlayer() *Character {
	return NewPlayer("PL-1", "John Wick", models.HUMAN)
}

func CreateEnemy() *Character {
	return models.NewArcDemonEnemy("BS-01", 1)
}

func SimulateBattle() {
	player := CreatePlayer()
	enemy := CreateEnemy()

	fmt.Println(" === [ENEMY DATA] ===")
	fmt.Println()
	enemy.DisplayAllStats()
	enemy.DisplayEquipment()
	enemy.DisplayInventory()

	fmt.Println(" === [PLAYER DATA #1] ===")
	fmt.Println()
	player.DisplayAllStats()
	player.DisplayEquipment()
	player.DisplayInventory()

	weapon := NewWeapon(models.GREATSWORD, models.STEEL)
	armor := NewArmor(models.BRONZE)

	player.Equip(weapon)
	player.Equip(armor)

	fmt.Println(" === [PLAYER DATA #2] ===")
	fmt.Println()
	player.DisplayAllStats()
	player.DisplayEquipment()
	player.DisplayInventory()

	fmt.Println(" === [PLAYER ATK 1/2 - Heavy] ===")
	fmt.Println()
	player.Attack(enemy, models.HEAVY_ATTACK)
	enemy.DisplayAllStats()

	fmt.Println(" === [PLAYER ATK 2/2 - Light] ===")
	fmt.Println()
	player.Attack(enemy, models.LIGHT_ATTACK)
	enemy.DisplayAllStats()

}
