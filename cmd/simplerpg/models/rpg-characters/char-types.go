package charpkg

import (
	rpginventory "github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models/rpg-inventory"
)

type Inventory = rpginventory.Inventory
type Equipment = rpginventory.Equipment

type Character struct {
	ID        string
	Name      string
	Type      CharacterType
	Stats     BaseStats
	Hp        float64
	Stamina   float64
	Weight    float64
	Level     int
	Exp       float64
	Inventory Inventory
	Equipment Equipment
}

type BaseStats struct {
	MaxHp       float64
	MaxStamina  float64
	MaxWeight   float64
	StmRecovery float64
	CritStrike  float64
	DodgeChance float64
}

type CharacterType int

const (
	// Player Char Type (Races)
	HUMAN CharacterType = iota
	ELF
	DWARF
	// Below are the Monsters' Types
	SIMPLE
	ELITE
	BOSS
)

// Created this Var to make single point from
// where to check if a Character Type is an Player type or not
var PlayerTypes = []CharacterType{
	HUMAN,
	ELF,
	DWARF,
}
