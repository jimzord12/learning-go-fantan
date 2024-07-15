package battlesystem

import rpgcharacters "github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models/rpg-characters"

type Character = rpgcharacters.Character

type Battle struct {
	ID           string
	Player       *Character
	Enemy        *Character
	BattleRounds []BattleRound
}

type BattleRound struct {
	ID       string
	Attacker *Character
	Defender *Character
	Action   BattleAction
}

type BattleAction int

const (
	ATTACK BattleAction = iota
	DEFEND
	REST
	HEAL
)
