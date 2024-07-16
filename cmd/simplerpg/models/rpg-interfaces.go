package models

type ICharacter interface {
	Attack(enemy *Character, luck float64, action BattleAction) error
}
