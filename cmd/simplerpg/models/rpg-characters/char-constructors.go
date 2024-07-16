package charpkg

func newCharacter(id string, name string, race CharacterType, baseStats BaseStats, hp float64, stm float64, lvl int) *Character {
	return &Character{
		ID:   id,
		Name: name,
		Type: race,
		Stats: BaseStats{
			MaxHp:       baseStats.MaxHp,
			MaxStamina:  baseStats.MaxStamina,
			MaxWeight:   baseStats.MaxWeight,
			StmRecovery: baseStats.StmRecovery,
			CritStrike:  baseStats.CritStrike,
			DodgeChance: baseStats.DodgeChance,
		},
		Hp:      hp,
		Stamina: stm,
		Level:   lvl,
	}
}

func NewPlayer(id string, name string, race CharacterType) *Character {
	var baseStats BaseStats = race.GetBaseStats()

	return newCharacter(id, name, race, baseStats, baseStats.MaxHp, baseStats.MaxStamina, 1)
}

func NewEnemy(id string, name string, monsterType CharacterType, baseHP float64, baseSTM float64, crit float64, dodgechn float64, stmRec float64, lvl int) *Character {
	return newCharacter(id, name, monsterType, BaseStats{
		MaxHp:       baseHP,
		MaxStamina:  baseSTM,
		CritStrike:  crit,
		DodgeChance: dodgechn,
		StmRecovery: stmRec,
	}, baseHP, baseSTM, lvl)
}
