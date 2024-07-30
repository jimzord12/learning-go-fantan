package models

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/generalhelpers"
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/logging"
)

////////////////////////////////////////////////////////////////////////////////////
/////////////////////////// CHARACTER CONSTRUCTORS /////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////

func newCharacter(id string, name string, race CharacterType, baseStats BaseStats, hp float64, stm float64, lvl int) *Character {
	newChar := Character{
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
		Hp:        hp,
		Stamina:   stm,
		Inventory: *NewInventory(),
		Level:     1,
	}

	// This is made only for the Monsters. Because we need to create Monsters
	// of specific levels. Players Level up gradually
	if lvl > 1 {
		newChar.LevelUpTo(lvl)
	}

	return &newChar
}

func NewPlayer(id string, name string, race CharacterType, lvl int) *Character {
	var baseStats BaseStats = race.GetBaseStats()
	return newCharacter(id, name, race, baseStats, baseStats.MaxHp, baseStats.MaxStamina, lvl)
}

func newEnemy(id string, name string, monsterType CharacterType, lvl int) *Character {
	var baseStats BaseStats = monsterType.GetBaseStats()
	return newCharacter(id, name, monsterType, baseStats, baseStats.MaxHp, baseStats.MaxStamina, lvl)
}

// Simple Enemies
func NewSpiderEnemy(id string, lvl int) *Character {
	// Creating the Enemy
	enemy := newEnemy(id, "Spider", SIMPLE, lvl)

	// Creating & Equiping them with a random Weapon
	enemy.EnemyEquipRandWeapon()

	// Creating & Equiping them with Armor
	enemy.EnemyEquipArmor()

	return enemy
}

func NewSlimeEnemy(id string, lvl int) *Character {
	enemy := newEnemy(id, "Slime", SIMPLE, lvl)
	enemy.EnemyEquipRandWeapon()
	enemy.EnemyEquipArmor()

	return enemy
}
func NewRatEnemy(id string, lvl int) *Character {
	enemy := newEnemy(id, "Rat", SIMPLE, lvl)
	enemy.EnemyEquipRandWeapon()
	enemy.EnemyEquipArmor()

	return enemy
}

func NewMushroomEnemy(id string, lvl int) *Character {
	enemy := newEnemy(id, "Mushroom", SIMPLE, lvl)
	enemy.EnemyEquipRandWeapon()
	enemy.EnemyEquipArmor()

	return enemy
}

// Elite Enemies
func NewGoblinEnemy(id string, lvl int) *Character {
	enemy := newEnemy(id, "Goblin", ELITE, lvl)
	enemy.EnemyEquipRandWeapon()
	enemy.EnemyEquipArmor()

	return enemy
}

func NewKoboldEnemy(id string, lvl int) *Character {
	enemy := newEnemy(id, "Kobold", ELITE, lvl)
	enemy.EnemyEquipRandWeapon()
	enemy.EnemyEquipArmor()

	return enemy
}

func NewImpEnemy(id string, lvl int) *Character {
	enemy := newEnemy(id, "Imp", ELITE, lvl)
	enemy.EnemyEquipRandWeapon()
	enemy.EnemyEquipArmor()

	return enemy
}

// Boss Enemies
func NewDragonEnemy(id string, lvl int) *Character {
	enemy := newEnemy(id, "Dragon", BOSS, lvl)
	enemy.EnemyEquipRandWeapon()
	enemy.EnemyEquipArmor()

	return enemy
}

func NewArcDemonEnemy(id string, lvl int) *Character {
	enemy := newEnemy(id, "Arc Demon", BOSS, lvl)
	enemy.EnemyEquipRandWeapon()
	enemy.EnemyEquipArmor()

	return enemy
}

func CreateRandomEnemy(monsterType CharacterType, dungeonDiff Difficulty, playerLvl int) *Character {
	if generalhelpers.ExistsInSlice(PlayerTypes, monsterType) {
		logging.LogError(logging.Logger, "While CreateRandomEnemy, in 'monsterType' param a PlayerType was given")
		panic("CreateRandomEnemy -> generalhelpers.ExistsInSlice(PlayerTypes, monsterType) -> true")
	}

	enemyTypesAmount := len(EnemyTypesToEnemyNames[monsterType])
	randMonster := rand.Intn(enemyTypesAmount)
	monsterName := EnemyTypesToEnemyNames[monsterType][randMonster]

	randLevelBooster := rand.Intn(3) // 0, 1, 2
	logging.GiveVertSpace(fmt.Sprintf("Created Monster of type [%s], Level [%d + %d]", monsterName, playerLvl, randLevelBooster))

	switch monsterName {
	case "Spider":
		return NewSpiderEnemy("no-id-yet", playerLvl+randLevelBooster)
	case "Slime":
		return NewSlimeEnemy("no-id-yet", playerLvl+randLevelBooster)
	case "Rat":
		return NewRatEnemy("no-id-yet", playerLvl+randLevelBooster)
	case "Mushroom":
		return NewMushroomEnemy("no-id-yet", playerLvl+randLevelBooster)
	case "Goblin":
		return NewGoblinEnemy("no-id-yet", playerLvl+randLevelBooster)
	case "Kobold":
		return NewKoboldEnemy("no-id-yet", playerLvl+randLevelBooster)
	case "Imp":
		return NewImpEnemy("no-id-yet", playerLvl+randLevelBooster)
	case "ArcDemon":
		return NewArcDemonEnemy("no-id-yet", playerLvl+randLevelBooster)
	case "Dragon":
		return NewDragonEnemy("no-id-yet", playerLvl+randLevelBooster)
	default:
		logging.LogError(logging.Logger, "At func CreateRandomEnemy, the retrieved (monsterName) from the EnemyTypesToEnemyNames map, is not supported")
		panic("CreateRandomEnemy -> monsterName -> not supported")
	}

}

////////////////////////////////////////////////////////////////////////////////////
/////////////////////////// INVENTORY CONSTRUCTORS /////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////

func NewInventory() *Inventory {
	return &Inventory{
		MaxSize: 20,
	}
}

////////////////////////////////////////////////////////////////////////////////////
///////////////////////////// DUNGEON CONSTRUCTORS /////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////

var ActiveDungeon *Dungeon

func DungeonInit(diff Difficulty) {
	fmt.Println("Creating Dungeon with Difficulty:", diff)
	ActiveDungeon = &Dungeon{
		Difficulty: diff,
	}
}

func NewDungeonNode(id string, encounter DungeonEncounter, nextEncounters []string, prevEncounters []string) *DungeonNode {
	return &DungeonNode{
		ID:             id,
		Encounter:      encounter,
		NextEncounters: nextEncounters,
		PrevEncounters: prevEncounters,
		Completed:      false,
	}
}

func NewDungeonMap(nodes ...*DungeonNode) *DungeonMap {
	var firstNodesIDs []string

	for _, node := range nodes {
		if node.IsAfterStartingNode() {
			firstNodesIDs = append(firstNodesIDs, node.ID)
		}
	}

	fmt.Printf("\nfirst Nodes IDs: %v\n", firstNodesIDs)

	startingNode := NewDungeonNode("Starting-Node", STARTING_NODE, firstNodesIDs, []string{})
	fmt.Printf("\nStarting Node: %v\n", startingNode)

	fmt.Printf("\nAll Nodes:\n")
	for _, node := range nodes {
		fmt.Printf("%+v\n", *node)
	}

	return &DungeonMap{
		AllNodes:     nodes,
		CurrentNode:  startingNode,
		PreviousNode: nil,
	}
}

////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////// ITEM CONSTRUCTORS /////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////

func newItem(name string, weight float64, itemType ItemType, value float64, material Material) *Item {
	return &Item{
		Name:     name,
		Weight:   generalhelpers.Round2Dec(weight),
		ItemType: itemType,
		Value:    generalhelpers.Round2Dec(value),
		Material: material,
	}
}

// Think of Material as Rarity
func NewWeapon(weaponType WeaponType, material Material) *Item {
	var weight float64
	var value float64
	var name string

	switch weaponType {
	case SWORD:
		name = material.String() + " Sword"
		weight = material.GetWeight(SWORD_VOLUME)
		value = weaponType.GetDmg(weight, material)
	case GREATSWORD:
		name = material.String() + " Greatsword"
		weight = material.GetWeight(GREATSWORD_VOLUME)
		value = weaponType.GetDmg(weight, material)
	case DAGGER:
		name = material.String() + " Dagger"
		weight = material.GetWeight(DAGGER_VOLUME)
		value = weaponType.GetDmg(weight, material)
	case MACE:
		name = material.String() + " Mace"
		weight = material.GetWeight(MACE_VOLUME)
		value = weaponType.GetDmg(weight, material)
	case SPEAR:
		name = material.String() + " Spear"
		weight = material.GetWeight(SPEAR_VOLUME)
		value = weaponType.GetDmg(weight, material)
	default:
		logging.LogError(logging.Logger, "The params:")
		fmt.Println(weaponType, material)
		log.Fatalln("[ERROR]: something went wrong while creating a new weapon")
	}

	return newItem(name, weight, WEAPON, value, material)

}

// I made Armor simpler than the Weapons
func NewArmor(material Material) *Item {
	var weight float64
	var value float64
	var name string

	weight = material.GetWeight(ARMOR_VOLUME)
	value = material.GetToughness() * (ARMOR_BASE_DEF + (weight * ARMOR_BASE_MULTI_PER_KG))
	name = material.String() + " Armor"

	return newItem(name, weight, ARMOR, value, material)
}

func NewAccessory(material Material) *Item {
	var weight float64
	var value float64
	var name string

	weight = material.GetWeight(ACC_VOLUME)
	value = material.GetToughness() * ACC_BASE_VALUE
	name = material.String() + " Accessory"

	return newItem(name, weight, ACCESSORY, value, material)
}

func NewPotion(size PotionType) *Item {
	weight := float64(size) * POTION_BASE_WEIGHT
	value := float64(size) * POTION_BASE_VALUE
	var name string

	switch size {
	case SMALL:
		name = PotionTypesToNames[1]
	case MEDIUM:
		name = PotionTypesToNames[2]
	case LARGE:
		name = PotionTypesToNames[3]
	default:
		log.Fatalln("[ERROR]: something went wrong while creating a new potion ")
	}

	return newItem(name, weight, POTION, value, NO_MATERIAL)
}

////////////////////////////////////////////////////////////////////////////////////
////////////////////////////// BATTLE CONSTRUCTORS /////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////

func NewBattleRound(id string, attacker *Character, defender *Character, atkAction BattleAction, defAction BattleAction, consumable *Item) *BattleRound {
	return &BattleRound{
		ID:             id,
		Attacker:       attacker,
		Defender:       defender,
		AttackerAction: atkAction,
		DefenderAction: defAction,
		Consumable:     consumable,
	}
}

func NewBattle(id string, player *Character, diff Difficulty, monsterType CharacterType) (*Battle, *Character) {
	if generalhelpers.ExistsInSlice(PlayerTypes, monsterType) {
		logging.LogError(logging.Logger, "While creating a NewBattle, in 'monsterType' param a PlayerType was given")
		panic("NewBattle -> generalhelpers.ExistsInSlice(PlayerTypes, monsterType) -> true")
	}

	monster := CreateRandomEnemy(monsterType, diff, player.Level)

	return &Battle{
		ID:         id,
		Player:     player,
		Enemy:      monster,
		Difficulty: diff,
	}, monster
}

////////////////////////////////////////////////////////////////////////////////////
////////////////////////////// SHOP CONSTRUCTORS //////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////

func NewShop(player *Character) *Shop {
	potionsStock := MakePotionsForShop(player)
	equipment := MakeEquipmentForShop()

	allPotions := make(map[PotionType][]*Item)
	for potType, potAmount := range potionsStock {
		var potions []*Item
		for i := 0; i < potAmount; i++ {
			potions = append(potions, NewPotion(potType))
		}
		allPotions[potType] = potions
	}

	return &Shop{
		Potions:   allPotions,
		Equipment: equipment,
	}
}
