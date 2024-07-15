package rpgcharacters

var HumanBaseStats = BaseStats{
	MaxHp:       100,
	MaxStamina:  100,
	MaxWeight:   40,
	StmRecovery: 40,
	CritStrike:  0.15,
	DodgeChance: 0.2,
}

var ElfBaseStats = BaseStats{
	MaxHp:       80,
	MaxStamina:  120,
	MaxWeight:   25,
	StmRecovery: 35,
	CritStrike:  0.2,
	DodgeChance: 0.3,
}

var DwarfBaseStats = BaseStats{
	MaxHp:       120,
	MaxStamina:  90,
	MaxWeight:   60,
	StmRecovery: 50,
	CritStrike:  0.1,
	DodgeChance: 0.1,
}
