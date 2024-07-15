package tests

import rpgcharacters "github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models/rpg-characters"

type Character = rpgcharacters.Character

var NewPlayer = rpgcharacters.NewPlayer
var NewEnemy = rpgcharacters.NewEnemy

func CreatePlayer() *Character {
	return NewPlayer("PL-1", "John Wick", rpgcharacters.HUMAN)
}

func CreateEnemy() *Character {
	return NewEnemy("ENM-1", "Swamp Slime", rpgcharacters.SIMPLE, 120, 60, 0.05, 0.15, 20, 1)
}

func SimulateBattle(player, enemy *Character) {}
