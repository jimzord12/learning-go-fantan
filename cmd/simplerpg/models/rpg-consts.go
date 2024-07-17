package models

////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////// CHARACTER CONSTS /////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////

var HumanBaseStats = BaseStats{
	MaxHp:       100,
	MaxStamina:  50,
	MaxWeight:   35,
	StmRecovery: 15,
	CritStrike:  0.15,
	DodgeChance: 0.2,
}

var ElfBaseStats = BaseStats{
	MaxHp:       80,
	MaxStamina:  60,
	MaxWeight:   25,
	StmRecovery: 10,
	CritStrike:  0.25,
	DodgeChance: 0.3,
}

var DwarfBaseStats = BaseStats{
	MaxHp:       120,
	MaxStamina:  45,
	MaxWeight:   50,
	StmRecovery: 20,
	CritStrike:  0.1,
	DodgeChance: 0.1,
}

var SimpleBaseStats = BaseStats{
	MaxHp:       100,
	MaxStamina:  60,
	MaxWeight:   999,
	StmRecovery: 20,
	CritStrike:  0.1,
	DodgeChance: 0.1,
}

var EliteBaseStats = BaseStats{
	MaxHp:       150,
	MaxStamina:  80,
	MaxWeight:   999,
	StmRecovery: 30,
	CritStrike:  0.05,
	DodgeChance: 0.15,
}

var BossBaseStats = BaseStats{
	MaxHp:       200,
	MaxStamina:  100,
	MaxWeight:   999,
	StmRecovery: 40,
	CritStrike:  0.1,
	DodgeChance: 0.0,
}

////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////// ITEM CONSTS ////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////

const (
	INV_MAX_WEIGHT = 40

	// Material Densities (KG/m^3)
	WOOD_DENSITY     = 1250.0
	BRONZE_DENSITY   = 2500.0
	IRON_DENSITY     = 3000.0
	STEEL_DENSITY    = 3500.0
	TITANIUM_DENSITY = 4250.0
	MYTHRIL_DENSITY  = 3000.0

	// Material Toughness
	WOOD_TOUGHNESS_MULTI     = 1.0
	BRONZE_TOUGHNESS_MULTI   = 1.25
	IRON_TOUGHNESS_MULTI     = 1.75
	STEEL_TOUGHNESS_MULTI    = 3
	TITANIUM_TOUGHNESS_MULTI = 5
	MYTHRIL_TOUGHNESS_MULTI  = 7.5

	// Weapon Shape Volumes (m^3)
	SWORD_VOLUME      = 0.0013
	GREATSWORD_VOLUME = 0.0036
	DAGGER_VOLUME     = 0.0005
	MACE_VOLUME       = 0.0022
	SPEAR_VOLUME      = 0.0029

	// Weapon' Damage
	WEAPON_BASE_DMG          = 10.0
	WEAPON_BASE_MULTI_PER_KG = 2.0

	// Armor Shape Volume (m^3)
	ARMOR_VOLUME = 0.005

	// Armor's Defense
	ARMOR_BASE_DEF          = 5
	ARMOR_BASE_MULTI_PER_KG = 1.5

	// Accessory
	ACC_VOLUME     = 0.0002
	ACC_BASE_VALUE = 6

	// Potions
	POTION_BASE_WEIGHT = 0.5
	POTION_BASE_VALUE  = 50.0
)

var WeaponTypes = map[WeaponType]string{
	SWORD:      "Sword",
	GREATSWORD: "Greatsword",
	DAGGER:     "Dagger",
	MACE:       "Mace",
	SPEAR:      "Spear",
}

var MaterialTypes = map[Material]string{
	WOOD:     WOOD.String(),
	BRONZE:   BRONZE.String(),
	IRON:     IRON.String(),
	STEEL:    STEEL.String(),
	TITANIUM: TITANIUM.String(),
	MYTHRIL:  MYTHRIL.String(),
}

var PotionTypes = map[PotionType]string{
	SMALL:  "HP Small Potion",
	MEDIUM: "HP Medium Potion",
	LARGE:  "HP Large Potion",
}

////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////// BATTLE CONSTS ////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////

// Simple
var SpiderPattern = EnemyBattlePattern{LIGHT_ATTACK, HEAVY_ATTACK, REST, DEFEND}
var SlimePattern EnemyBattlePattern = EnemyBattlePattern{HEAVY_ATTACK, LIGHT_ATTACK, DEFEND, REST}
var RatPattern EnemyBattlePattern = EnemyBattlePattern{HEAVY_ATTACK, REST, LIGHT_ATTACK, DEFEND}
var MushroomPattern EnemyBattlePattern = EnemyBattlePattern{DEFEND, LIGHT_ATTACK, REST, HEAVY_ATTACK}

// Elite
var GoblinPattern EnemyBattlePattern = EnemyBattlePattern{LIGHT_ATTACK, HEAVY_ATTACK, REST, DEFEND}
var KoboldPattern EnemyBattlePattern = EnemyBattlePattern{DEFEND, LIGHT_ATTACK, REST, HEAVY_ATTACK, DEFEND}
var ImpPattern EnemyBattlePattern = EnemyBattlePattern{HEAVY_ATTACK, DEFEND, LIGHT_ATTACK, REST}

// Bosses have random BattleActions

var EnemyNamesToPatterns = map[string]EnemyBattlePattern{
	"Spider":   SpiderPattern,
	"Slime":    SlimePattern,
	"Rat":      RatPattern,
	"Mushroom": MushroomPattern,
	"Goblin":   GoblinPattern,
	"Kobold":   KoboldPattern,
	"Imp":      ImpPattern,
}

var PlayerBattleRollChances = []int{1, 2, 3, 4, 6, 5, 4, 3, 2}
var EnemyBattleRollChances = []int{1, 2, 3, 4, 6, 5, 4, 3, 2}

////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////// LEVEL CONSTS ////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////

const BaseExp float64 = 500.0
