package rpgitems

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
	BASE_DMG                 = 10.0
	BASE_WEAPON_MULTI_PER_KG = 2.0

	// Armor Shape Volume (m^3)
	ARMOR_VOLUME = 0.005

	// Armor's Defense
	BASE_DEF                = 5
	BASE_ARMOR_MULTI_PER_KG = 1.5

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
