package tests

import "github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models"

type Character = models.Character

var NewPlayer = models.NewPlayer

func CreatePlayer() *Character {
	return NewPlayer("PL-1", "John Wick", models.HUMAN)
}

func CreateEnemy() *Character {
	return NewEnemy("ENM-1", "Swamp Slime", models.SIMPLE, 120, 60, 0.05, 0.15, 20, 1)
}

func SimulateBattle(player, enemy *Character) {}
