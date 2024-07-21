package tests

import (
	"fmt"

	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models"
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/logging"
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

	weapon := NewWeapon(models.MACE, models.BRONZE)
	armor := NewArmor(models.BRONZE)

	// Move Items to Inventory
	player.MoveToInventory(weapon)
	player.MoveToInventory(armor)

	player.Equip(weapon)
	player.Equip(armor)

	fmt.Println()
	fmt.Println(" === [PLAYER DATA] ===")
	fmt.Println()
	player.DisplayAllStats()
	logging.StdDivider("*", 75)
	player.DisplayEquipment()
	logging.StdDivider("*", 75)
	player.DisplayInventory()

	fmt.Println()
	fmt.Println(" === [ENEMY DATA] ===")
	fmt.Println()
	enemy.DisplayAllStats()
	enemy.DisplayEquipment()
	enemy.DisplayInventory()

	// fmt.Println()
	// fmt.Println(" === [PLAYER ATK 1/2 - Heavy] ===")
	// fmt.Println()
	// player.Attack(enemy, models.HEAVY_ATTACK)
	// enemy.DisplayAllStats()
	// logging.StdDivider("-", 75)
	// player.DisplayAllStats()

	// fmt.Println()
	// fmt.Println(" === [DEMON ATK 1/1 - Heavy] ===")
	// fmt.Println()
	// enemy.Attack(player, models.HEAVY_ATTACK)
	// enemy.DisplayAllStats()
	// logging.StdDivider("-", 75)
	// player.DisplayAllStats()

	// fmt.Println()
	// fmt.Println(" === [PLAYER ATK 2/2 - Light] ===")
	// fmt.Println()
	// player.Attack(enemy, models.LIGHT_ATTACK)
	// enemy.DisplayAllStats()
	// logging.StdDivider("-", 75)
	// player.DisplayAllStats()

	firstRound := models.NewBattleRound("RD-1", player, enemy, models.HEAVY_ATTACK, models.LIGHT_ATTACK, &models.Item{})

	models.PerformRound(*firstRound)
	fmt.Println("")
	fmt.Println("")
	player.DisplayAllStats()
	fmt.Println("")
	fmt.Println("")
	enemy.DisplayAllStats()
}
