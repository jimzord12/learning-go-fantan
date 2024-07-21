package models

import (
	"errors"
	"fmt"
	"log"
	"math/rand"

	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/generalhelpers"
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/logging"
)

////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////// CHARACTER METHODS /////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////

func (char *Character) DisplayInventory() {
	if len(char.Inventory.Items) == 0 {
		fmt.Println("Inventory is Empty")
	}
	for i, v := range char.Inventory.Items {
		fmt.Println("1.", i, ":", v)
	}
}

func (char *Character) DisplayEquipment() {
	fmt.Println("Weapon:", char.Equipment.WeaponSlot, "\nArmor:", char.Equipment.ArmorSlot)
}

func (char *Character) DisplayAllStats() {
	fmt.Println("Name:", char.Name)
	fmt.Printf("HP: %.2f/%.2f\n", char.Hp, char.Stats.MaxHp)
	fmt.Printf("STM: %.2f/%.2f\n", char.Stamina, char.Stats.MaxStamina)
	fmt.Printf("WGT: %.2f/%.2f\n", char.Weight, char.Stats.MaxWeight)
	fmt.Printf("Crit: %.2f\n", char.Stats.CritStrike)
	fmt.Printf("Dodge: %.2f\n", char.Stats.DodgeChance)
	fmt.Println("Level:", char.Level)
	fmt.Println("EXP:", char.Exp)
}

func (char *Character) isPlayer() bool {
	return generalhelpers.ExistsInSlice(PlayerTypes, char.Type)
}

// TODO: Refactor using Composition. First seperate Player and Enemy from Character
func (char *Character) Attack(enemy *Character, atkType BattleAction) error {
	equippedWeapon := char.Equipment.WeaponSlot

	// Has a Weapon?
	if equippedWeapon == nil {
		logging.LogError(logging.Logger, "(func (char *Character) Attack(enemy *Character, atkType BattleAction) error) -> No Weapon is equipped")

		return fmt.Errorf("[ERROR]: You cannot attack, you are NOT holding a Weapon")
	}

	// is that a Weapon?
	if char.Equipment.WeaponSlot.ItemType != WEAPON {
		logging.LogError(logging.Logger, "(func (char *Character) Attack(enemy *Character, atkType BattleAction) error) -> The equipped Item is NOT a Weapon")
		return fmt.Errorf("[ERROR]: You cannot attack with (%s), its NOT a Weapon", equippedWeapon.Name)
	}

	// Are you performing the correct action?
	if atkType != LIGHT_ATTACK && atkType != HEAVY_ATTACK {
		logging.LogError(logging.Logger, "(func (char *Character) Attack(enemy *Character, atkType BattleAction) error) -> Wrong Battle Action")
		return fmt.Errorf("[ERROR]: You need to select 'LIGHT_ATTACK' or 'HEAVY_ATTACK' as BattleAction, not (%v)", atkType)
	}

	// Has enough Stamina?
	reqStamina, err := GetRequiredStamina(equippedWeapon, atkType)
	if err != nil {
		logging.LogError(logging.Logger, "(func (char *Character) Attack(enemy *Character, atkType BattleAction) error) -> Something went wrong with GetRequiredStamina(equippedWeapon, atkType)")
		return err
	}

	if char.Stamina < reqStamina {
		logging.LogError(logging.Logger, "(func (char *Character) Attack(enemy *Character, atkType BattleAction) error) -> Not enough Stamina")
		return fmt.Errorf("[ERROR]: You do not have enough stamina for this action (%v)", atkType)
	}

	if WillEnemyDodge(enemy) {
		fmt.Println("=== The Defender Successfully DODGE! ===")
		char.Stamina -= reqStamina * 2

		return nil
	}

	var attackTypeFactor float64

	if atkType == LIGHT_ATTACK {
		attackTypeFactor = 1.0 // Light Atk Multi
	} else {
		attackTypeFactor = 1.5 // Heavy Atk Multi
	}

	// Calculate the Luck Factor for the Attack
	luckFactor := BattleLuckRoll(char.isPlayer())
	fmt.Printf("[%s] Action Effectiveness (luck): [%.2f%%]\n", char.Name, luckFactor*100)

	// Calculate if its a Crit!
	var crit int
	if char.isCrit() {
		crit = 2
	} else {
		crit = 1
	}

	// Calculate the Damage based on: Weapon, luckFactor and the AttackType (Light or Heavy)
	damage := equippedWeapon.Value * luckFactor * attackTypeFactor * float64(crit)
	fmt.Printf("[%s], [atkType: %d], Atk Power is: (%.2f)\n", char.Name, atkType, damage)

	// Decrease Player's Stamina
	char.Stamina -= reqStamina

	// Decrease Enemy's HP
	enemy.Hp -= damage

	return nil
}

func WillEnemyDodge(enemy *Character) bool {
	// Monsters can not Dodge
	if !enemy.isPlayer() {
		return false
	}

	luck := rand.Intn(100)
	fmt.Printf("(Dodge Calculations): [%s] has [%d%%] Dodge Chance\n", enemy.Name, int(enemy.Stats.DodgeChance*100))
	fmt.Printf("(Dodge Calculations): Luck is [%d], will Dodge [%t]\n", luck, luck < int(enemy.Stats.DodgeChance*100))
	return luck < int(enemy.Stats.DodgeChance*100)
}

func (char *Character) isCrit() bool {
	luck := rand.Intn(100)
	fmt.Printf("(Crit Calculations): [%s] has [%d%%] Crit Chance\n", char.Name, int(char.Stats.CritStrike*100))
	fmt.Printf("(Crit Calculations): Luck is [%d], will Crit [%t]\n", luck, luck < int(char.Stats.CritStrike*100))
	return luck < int(char.Stats.CritStrike*100)
}

func (char *Character) Unequip(item *Item) error {
	if item.ItemType != WEAPON && item.ItemType != ARMOR && item.ItemType != ACCESSORY {
		return fmt.Errorf("[ERROR]: You cannot equip this item type (%s)", item.ItemType)
	}

	char.MoveToInventory(item)

	switch item.ItemType {
	case WEAPON:
		char.Equipment.WeaponSlot = nil
	case ARMOR:
		char.Equipment.ArmorSlot = nil
	}

	return nil
}

// This Function is meant to be used ONLY by Enemies.
func (enemy *Character) EnemyEquipRandWeapon() error {
	// Get Randmon Weapon Type
	randWeaponType := rand.Intn(5)

	// Create the Weapon
	weapon := NewWeapon(WeaponType(randWeaponType), DifficultyToMaterial(ActiveDungeon.Difficulty))

	// Add the Weapon to the Inventory
	enemy.MoveToInventory(weapon)

	// Equip the Weapon
	err := enemy.Equip(weapon)
	if err != nil {
		logging.LogError(logging.Logger, "[func EquipRandWeapon]")
		return err
	}

	return nil
}

func (enemy *Character) EnemyEquipArmor() error {
	armor := NewArmor(DifficultyToMaterial(ActiveDungeon.Difficulty))
	enemy.MoveToInventory(armor)

	err := enemy.Equip(armor)
	if err != nil {
		logging.LogError(logging.Logger, "[func EquipRandArmor]")
		return err
	}

	return nil
}

func (char *Character) Equip(item *Item) error {
	// 1. Check if the Item is of Type WEAPON or ARMOR or ACCESSORY
	if item.ItemType != WEAPON && item.ItemType != ARMOR && item.ItemType != ACCESSORY {
		logging.LogError(logging.Logger, "(func (char *Character) Equip(item *Item) error) -> passed wrong Item type")
		return fmt.Errorf("[ERROR]: You cannot equip this item type (%s)", item.Name)
	}

	// 2. Check if the Item exists in the Inventory
	if !generalhelpers.ExistsInSlice(char.Inventory.Items, item) {
		logging.LogError(logging.Logger, "(func (char *Character) Equip(item *Item) error) -> Item does not exist in Inventoy")
		return fmt.Errorf("[ERROR]: The item you are trying to equip does not exist in the Character's Inventory (%s)", item.Name)
	}

	// 3. Check if the Player already has an Item of the same type equipped
	switch item.ItemType {
	case WEAPON:
		// Player already has a Weapon equipped
		if char.Equipment.WeaponSlot != nil {
			char.Unequip(char.Equipment.WeaponSlot) // Unequip Current Weapon
			char.Equipment.WeaponSlot = item        // Equip Selected Weapon
		} else {
			char.Equipment.WeaponSlot = item // Equip Selected Weapon
		}
	case ARMOR:
		if char.Equipment.ArmorSlot != nil {
			char.Unequip(char.Equipment.ArmorSlot)
			char.Equipment.ArmorSlot = item
		} else {
			char.Equipment.ArmorSlot = item
		}
	case ACCESSORY:
		if char.Equipment.AccessorySlot != nil {
			char.Unequip(char.Equipment.AccessorySlot)
			char.Equipment.AccessorySlot = item
		} else {
			char.Equipment.AccessorySlot = item
		}

	default:
		return fmt.Errorf("[ERROR]: Something went wrong while trying to EQUIP the item: (%v)", *item)
	}

	char.RemoveFromInventory(item)
	char.Weight += item.Weight
	return nil
}

func (char *Character) UseItem(item *Item) error {
	switch item.ItemType {
	case WEAPON, ARMOR:
		char.Equip(item)
	case POTION:
		fmt.Printf("Potion (%s) is used by [%s]", item.Name, char.Name)
		char.Hp += item.Value
	default:
		return fmt.Errorf("[ERROR]: This item is not supported yet: (%v)", *item)
	}

	return nil
}

func (char *Character) Rest() error {
	if char.Stamina == char.Stats.MaxStamina {
		return fmt.Errorf("[ERROR]: You can NOT rest, your Stamina is full")
	}

	newStamina := char.Stamina + char.Stats.StmRecovery

	if newStamina > char.Stats.MaxStamina {
		char.Stamina = char.Stats.MaxStamina
		return nil
	}

	char.Stamina += char.Stats.StmRecovery
	return nil
}

func (race CharacterType) GetBaseStats() BaseStats {
	switch race {
	case HUMAN:
		return HumanBaseStats
	case ELF:
		return ElfBaseStats
	case DWARF:
		return DwarfBaseStats
	case SIMPLE:
		return SimpleBaseStats
	case ELITE:
		return ElfBaseStats
	case BOSS:
		return BossBaseStats
	default:
		logging.LogError(logging.Logger, "Something went wrong while getting the base stats from a character race")
		panic("[ERROR]: At char-types.go, func signature: 'func (race CharacterRace) GetBaseStats() BaseStats'")
	}
}

func (char *Character) MoveToInventory(item *Item) error {
	newSize := len(char.Inventory.Items) + 1
	newWeight := char.Weight + item.Weight

	if newSize > char.Inventory.MaxSize {
		logging.LogError(logging.Logger, "(func (char *Character) MoveToInventory(item *Item) error) -> Inv Max Size reached")
		return fmt.Errorf("[ERROR]: Character's Inventory is full (no more slots)")
	}

	if newWeight > char.Stats.MaxWeight {
		logging.LogError(logging.Logger, "(func (char *Character) MoveToInventory(item *Item) error) -> Inv Max Weight reached")
		return fmt.Errorf("[ERROR]: Character's Inventory is too heavy (no more slots)")
	}

	char.Inventory.Items = append(char.Inventory.Items, item)
	char.Weight += item.Weight

	return nil
}

func (char *Character) RemoveFromInventory(item *Item) error {
	// 1. Find the Item in the Character's Inventory
	index, err := char.Inventory.FindItemIndex(item.Name)
	if err != nil {
		log.Println(err)
		return err
	}

	// 2. Remove the Item from the Character's Inventory
	char.Inventory.Items = generalhelpers.RemoveFromSlice(char.Inventory.Items, index)

	// 3. Update the Character's Inventory stats
	char.Weight -= item.Weight
	return nil
}

func (char *Character) GainEXP(gainedExp float64) {
	char.Exp += gainedExp // Add gained EXP
	forNextLevel := ExpForNextLevel(char.Level)
	fmt.Printf("EXP Required For Next Level: %.2f\n", forNextLevel)
	newExp := char.Exp

	for newExp >= forNextLevel {
		char.LevelUp()
		newExp -= forNextLevel
		forNextLevel = ExpForNextLevel(char.Level)
		fmt.Printf("EXP that Remained : %.2f\n", newExp)
	}

	char.Exp = newExp
	fmt.Println("Current EXP: ", newExp)
}

func (char *Character) LevelUp() {
	char.Level += 1
	char.Stats.MaxHp += 20
	char.Stats.MaxStamina += 5
	char.Stats.MaxWeight += 2
	char.Stats.StmRecovery += 2

	char.Exp = 0

	fmt.Printf("[%s] just Leveled Up! (%d) -> (%d)\n", char.Name, char.Level-1, char.Level)
}

func PerformBattleAction(action BattleAction, attacker *Character, defender *Character, consumable *Item) {
	switch action {
	case LIGHT_ATTACK:
		attacker.Attack(defender, LIGHT_ATTACK)
	case HEAVY_ATTACK:
		attacker.Attack(defender, HEAVY_ATTACK)
	case DEFEND:
		//TODO:
	case REST:
		attacker.Rest()
	case HEAL:
		attacker.UseItem(consumable)
	default:
		logging.LogError(logging.Logger, "Provided PlayerAction through the battleround param is nsupported")
	}
}

func PerformRound(round BattleRound) (hasBattleEnded bool) {
	// if generalhelpers.ExistsInSlice(PlayerTypes, char.Type) {
	// 	PerformBattleAction(round.PlayerAction, round.Attacker, round.Defender, round.Consumable)
	// } else {
	// 	PerformBattleAction(round.EnemyAction, round.Attacker, round.Defender, round.Consumable)
	// }

	// 1. Attacker (Player) performs his/her Action
	fmt.Printf("[ATTACK 1/2]: (%s) -> (%s)", round.Attacker.Name, round.Defender.Name)
	fmt.Println()
	PerformBattleAction(round.AttackerAction, round.Attacker, round.Defender, round.Consumable)

	// 2. Checking Enemy
	if round.Defender.Hp <= 0 {
		fmt.Printf("Battle Ended, Winner: (%s), Losser: (%s)", round.Attacker.Name, round.Defender.Name)
		return true
	}

	// 3. Defender (Monster) perform its Action
	fmt.Printf("[ATTACK 2/2]: (%s) -> (%s)", round.Defender.Name, round.Attacker.Name)
	fmt.Println()
	PerformBattleAction(round.DefenderAction, round.Defender, round.Attacker, round.Consumable)

	// 4. Checking Player
	if round.Attacker.Hp <= 0 {
		fmt.Printf("Battle Ended, Winner: (%s), Losser: (%s)", round.Defender.Name, round.Attacker.Name)
		return true
	}

	return false
}

////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////// INVENTORY METHODS /////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////

func (it Inventory) FindItemIndex(desiredItemName string) (int, error) {
	for i, v := range it.Items {
		if v.Name == desiredItemName {
			return i, nil
		}
	}

	return -1, errors.New("[ERROR]: Could not find Index")
}

////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////// ITEMS METHODS //////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////

var logger = logging.Logger

///// WeaponType Methods /////

func (wt WeaponType) GetDmg(weight float64, material Material) float64 {
	return material.GetToughness() * (WEAPON_BASE_DMG + (weight * WEAPON_BASE_MULTI_PER_KG))
}

///// Material Methods /////

// This could be stored as a const, but its not for practise purposes
func (mat Material) GetWeight(volume float64) float64 {
	switch mat {
	case WOOD:
		return volume * WOOD_DENSITY
	case BRONZE:
		return volume * BRONZE_DENSITY
	case IRON:
		return volume * IRON_DENSITY
	case STEEL:
		return volume * STEEL_DENSITY
	case TITANIUM:
		return volume * TITANIUM_DENSITY
	case MYTHRIL:
		return volume * MYTHRIL_DENSITY
	default:
		logging.LogError(logger, fmt.Sprintf("Something went wrong while attemping to calculate this material's (%v) weight\n", mat))
		return -1
	}
}

func (mat Material) GetToughness() float64 {
	switch mat {
	case WOOD:
		return WOOD_TOUGHNESS_MULTI
	case BRONZE:
		return BRONZE_TOUGHNESS_MULTI
	case IRON:
		return IRON_TOUGHNESS_MULTI
	case STEEL:
		return STEEL_TOUGHNESS_MULTI
	case TITANIUM:
		return TITANIUM_TOUGHNESS_MULTI
	case MYTHRIL:
		return MYTHRIL_TOUGHNESS_MULTI
	default:
		logging.LogError(logger, fmt.Sprintf("Something went wrong while getting this material's (%v) toughness\n", mat))
		return -1
	}
}

func (mat Material) String() string {
	switch mat {
	case NO_MATERIAL:
		return "None"
	case WOOD:
		return "Wooden"
	case BRONZE:
		return "Bronze"
	case IRON:
		return "Iron"
	case STEEL:
		return "Steel"
	case TITANIUM:
		return "Titanium"
	case MYTHRIL:
		return "Mythril"
	default:
		logging.LogError(logger, "Something went wrong while attemping to return the string represensation of a material")
		return "[ERROR]"
	}
}

///// ItemType Methods /////

func (i ItemType) String() string {
	switch i {
	case WEAPON:
		return "Weapon"
	case ARMOR:
		return "Armor"
	case POTION:
		return "Potion"
	case ACCESSORY:
		return "Accessory"
	default:
		logging.LogError(logger, "Something went wrong while attemping to return the represensation of an itemType")
		return "[ERROR]"
	}
}

////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////// DUNGEON METHODS /////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////// BATTLE METHODS //////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////

func GetRequiredStamina(weapon *Item, atkType BattleAction) (float64, error) {
	if weapon == nil {
		logging.LogError(logging.Logger, "(func GetRequiredStaminaFor(weapon *Item) float64) you passed a empty pointer.")
		return -1, errors.New("weapon *Item -> nil")
	}

	if atkType != LIGHT_ATTACK && atkType != HEAVY_ATTACK {
		logging.LogError(logging.Logger, "(func GetRequiredStaminaFor(weapon *Item) float64) you passed a wrong BattleAction value, supports only (LIGHT_ATTACK | HEAVY_ATTACK).")
		return -1, errors.New("atkType is not (LIGHT_ATTACK | HEAVY_ATTACK)")
	}

	var atkTypeFactor float64

	if atkType == LIGHT_ATTACK {
		atkTypeFactor = 1.5
	} else {
		atkTypeFactor = 2
	}

	return weapon.Weight * atkTypeFactor, nil
}

func RandBattleAction() BattleAction {
	action := rand.Intn(4)
	return BattleAction(action)
}

func rollLuck(isPlayer bool) int {
	var weights []int

	if isPlayer {
		// Player has higher chances to get 3-7
		weights = PlayerBattleRollChances
	} else {
		// Enemy has higher chances to get 1-4
		weights = EnemyBattleRollChances
	}

	cumulativeWeights := make([]int, len(weights))

	cumulativeWeights[0] = weights[0]
	for i := 1; i < len(weights); i++ {
		cumulativeWeights[i] = cumulativeWeights[i-1] + weights[i]
	}

	randValue := rand.Intn(cumulativeWeights[len(cumulativeWeights)-1]) + 1

	for i, cw := range cumulativeWeights {
		if randValue <= cw {
			return i
		}
	}
	return 0 // fallback
}

// Function to calculate the boost percentage based on the luck roll
func luckBoost(luck int) float64 {
	return float64(luck) / 4.0 // 0 to 2.0 (0% to 200%)
}

func BattleLuckRoll(isPlayer bool) float64 {
	luck := rollLuck(isPlayer)

	return luckBoost(luck)
}

////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////// LEVEL METHODS //////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////

func ExpForNextLevel(currentLevel int) float64 {
	if currentLevel < 1 {
		return 0
	}

	if currentLevel == 1 {
		return BaseExp
	}

	if currentLevel == 2 {
		return BaseExp * 2
	}

	return float64(int(BaseExp) * (currentLevel * (currentLevel - 1) / 2))
}

func ExpGainedFromEnemy(playerLevel int, enemy *Character) float64 {
	if playerLevel < 1 {
		logging.LogError(logging.Logger, "(func ExpGainedFromEnemy(playerLevel int, enemy *Character) float64) -> The provided player Level is less than 1, which should NEVER happen")
		panic("Inserted player level was less than 1")
	}

	typeFactor := float64(enemy.Type) - 2.0 // Simple x1, Elite x2, Boss x3
	typeFactorLinear := 2*typeFactor - 1
	playerLvlDiff := playerLevel - enemy.Level
	lvlDiffFactor := 0.25 * float64(playerLvlDiff)
	base := (enemy.Hp * lvlDiffFactor) + enemy.Hp
	result := base * typeFactorLinear

	fmt.Printf("[%s] provided: %.2f XP\n", enemy.Name, result)

	return result
}

////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////// LOOT METHODS //////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////

func getMaterialDropChance() (Material, bool) {
	var luck = rand.Intn(101)
	var dungeon = ActiveDungeon

	if luck >= 95 {
		return DifficultyToMaterial(dungeon.Difficulty + 2), true
	} else if luck >= 70 && luck < 95 {
		return DifficultyToMaterial(dungeon.Difficulty + 1), true
	} else if luck >= 20 && luck < 70 {
		return DifficultyToMaterial(dungeon.Difficulty), true
	} else {
		return -1, false
	}
}

func getPotionDropChance() (PotionType, bool) {
	var luck = rand.Intn(101)

	if luck >= 95 {
		return LARGE, true
	} else if luck >= 70 && luck < 95 {
		return MEDIUM, true
	} else if luck >= 20 && luck < 70 {
		return SMALL, true
	} else {
		return -1, false
	}
}

func getEquipmentTypeDropChance() ItemType {
	var luck = rand.Intn(3)

	switch luck {
	case 0:
		return WEAPON
	case 1:
		return ARMOR
	case 2:
		return ACCESSORY
	default:
		return -1
	}
}

func getWeaponTypeDropChance() WeaponType {
	var luck = rand.Intn(5)

	switch luck {
	case 0:
		return SWORD
	case 1:
		return DAGGER
	case 2:
		return GREATSWORD
	case 3:
		return MACE
	case 4:
		return SPEAR
	default:
		logging.LogError(logging.Logger, "In 'getWeaponTypeDropChance()' luck is -1, but this should not be possible")
		return -1
	}
}

// False => No Drops
// True => You must check if struct field are not -1
// (-1) means No Drop for that Field
func CalcDrops() (EnemyDrops, bool) {
	equipment, hasEquipDrop := getMaterialDropChance()
	potion, hasPotDrop := getPotionDropChance()

	if hasEquipDrop || hasPotDrop {

		return EnemyDrops{
			EquipmentType:     getEquipmentTypeDropChance(),
			EquipmentMaterial: equipment,
			PotionDrop:        potion,
		}, true
	}

	return EnemyDrops{}, false

}

func (loot EnemyDrops) GetLoot() []*Item {
	var drops []*Item

	if loot.EquipmentMaterial != -1 {
		switch loot.EquipmentType {
		case WEAPON:
			weaponType := getWeaponTypeDropChance()
			drops = append(drops, NewWeapon(weaponType, loot.EquipmentMaterial))
		case ARMOR:
			drops = append(drops, NewArmor(loot.EquipmentMaterial))
		case ACCESSORY:
			drops = append(drops, NewAccessory(loot.EquipmentMaterial))
		default:
			logging.LogError(Logger, "While getting loot, signature: 'func (loot EnemyDrops) GetLoot() [2]*Item'")
			panic("[ERROR]: getting loot")
		}
	}

	if loot.PotionDrop != -1 {
		drops = append(drops, NewPotion(loot.PotionDrop))
	}

	return drops
}
