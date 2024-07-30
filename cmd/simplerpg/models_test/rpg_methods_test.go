package modelstest

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models"
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/logging"
)

type Item = models.Item
type BattleAction = models.BattleAction

// TestMain allows for setup and teardown before and after tests
func TestMain(m *testing.M) {
	// Setup code
	logging.LogInit()
	models.DungeonInit(models.WOOD_DIF)

	// Run the tests
	exitCode := m.Run()

	// Teardown code if needed

	// Exit with the appropriate code
	os.Exit(exitCode)
}

func TestDisplayInventory(t *testing.T) {
	player := models.NewPlayer("001", "Takis", models.ELF, 5)

	// Capture the standard output
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}
	defer reader.Close()
	defer writer.Close()

	stdout := os.Stdout
	defer func() { os.Stdout = stdout }() // Restore original stdout after test
	os.Stdout = writer

	// Channel to capture output
	outputCh := make(chan string)
	go func() {
		var buf bytes.Buffer
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			buf.WriteString(scanner.Text() + "\n")
		}
		outputCh <- buf.String()
	}()

	// Call the method
	player.DisplayInventory()
	writer.Close() // Close the writer to signal EOF

	// Get the captured output
	output := <-outputCh

	// Check if the output is as expected
	expectedOutput := "Inventory is Empty\n"
	if output != expectedOutput {
		t.Errorf("expected %q but got %q", expectedOutput, output)
	}

	// Test with items in the inventory
	player.Inventory.Items = []*models.Item{models.NewWeapon(models.SWORD, models.IRON), models.NewAccessory(models.TITANIUM)}

	// Capture the standard output again
	reader, writer, err = os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}
	defer reader.Close()
	defer writer.Close()

	os.Stdout = writer

	go func() {
		var buf bytes.Buffer
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			buf.WriteString(scanner.Text() + "\n")
		}
		outputCh <- buf.String()
	}()

	// Call the method again
	player.DisplayInventory()
	writer.Close() // Close the writer to signal EOF

	// Get the captured output
	output = <-outputCh

	// Check if the output is as expected
	expectedOutput = fmt.Sprintf("Inv Slot ( 0 ): %v\nInv Slot ( 1 ): %v\n", player.Inventory.Items[0], player.Inventory.Items[1])
	if output != expectedOutput {
		t.Errorf("expected %q but got %q", expectedOutput, output)
	}
}

func TestCharacter_Attack(t *testing.T) {
	// logging.LogInit()
	// models.DungeonInit(models.WOOD_DIF)

	weaponSpear := models.NewWeapon(models.SPEAR, models.MYTHRIL)
	notAWeapon := models.NewAccessory(models.WOOD)

	player := models.NewPlayer("PL-001", "Pl_Test_1", models.ELF, 5)

	player2 := models.NewPlayer("PL-002", "Pl_Test_2", models.ELF, 5)
	player2.Equipment.WeaponSlot = notAWeapon

	player3 := models.NewPlayer("PL-003", "Pl_Test_3", models.ELF, 5) // 1. Create Player
	err := player3.MoveToInventory(weaponSpear)                       // 2. Move Weapon to Inv
	if err != nil {
		log.Println(err)
	}
	err = player3.Equip(weaponSpear) // 2. Equip the Weapon, From Inv -> Player's WeaponSlot
	if err != nil {
		log.Println(err)
	}

	player4 := models.NewPlayer("PL-004", "Pl_Test_4", models.ELF, 5)
	player4.Stamina = 1
	err = player4.MoveToInventory(weaponSpear)
	if err != nil {
		log.Println(err)
	}
	err = player4.Equip(weaponSpear)
	if err != nil {
		log.Println(err)
	}

	playerSuccess := models.NewPlayer("PL-005", "Pl_Test_5", models.ELF, 5)
	err = playerSuccess.MoveToInventory(weaponSpear)
	if err != nil {
		log.Println(err)
	}
	err = playerSuccess.Equip(weaponSpear)
	if err != nil {
		log.Println(err)
	}

	tests := []struct {
		name        string
		char        *models.Character
		enemy       *models.Character
		atkType     models.BattleAction
		expectedErr string
	}{
		{
			name:        "No weapon equipped",
			char:        player,
			enemy:       models.NewGoblinEnemy("MON-001", 2),
			atkType:     models.LIGHT_ATTACK,
			expectedErr: "[ERROR]: You cannot attack, you are NOT holding a Weapon",
		},
		{
			name:        "Equipped item is not a weapon",
			char:        player2,
			enemy:       models.NewGoblinEnemy("MON-002", 2),
			atkType:     models.LIGHT_ATTACK,
			expectedErr: "[ERROR]: You cannot attack with (Wooden Accessory), its NOT a Weapon",
		},
		{
			name:        "Wrong battle action",
			char:        player3,
			enemy:       models.NewGoblinEnemy("MON-003", 2),
			atkType:     models.DEFEND,
			expectedErr: "[ERROR]: You need to select 'LIGHT_ATTACK' or 'HEAVY_ATTACK' as BattleAction, not (DEFEND)",
		},
		{
			name:        "Not enough stamina",
			char:        player4,
			enemy:       models.NewArcDemonEnemy("MON-004", 15),
			atkType:     models.LIGHT_ATTACK,
			expectedErr: "[ERROR]: You do not have enough stamina for this action (LIGHT ATTACK)",
		},
		{
			name:        "Successful attack",
			char:        playerSuccess,
			enemy:       models.NewImpEnemy("MON-002", 2),
			atkType:     models.LIGHT_ATTACK,
			expectedErr: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.char.Attack(tt.enemy, tt.atkType)
			if err != nil {
				if err.Error() != tt.expectedErr {
					t.Errorf("expected error %q but got %q", tt.expectedErr, err.Error())
				} else {
					t.Logf("PASS: %s", tt.name)
				}
			} else {
				if tt.expectedErr != "" {
					t.Errorf("expected error %q but got nil", tt.expectedErr)
				} else {
					t.Logf("PASS: %s", tt.name)
				}
			}
		})
	}
}

func TestGetRequiredStamina(t *testing.T) {
	correctWeapon := models.NewWeapon(models.SWORD, models.IRON)
	wrongWeapon := models.NewArmor(models.STEEL)
	emptyItem := &models.Item{}

	correctAction := models.HEAVY_ATTACK
	wrongAction := models.REST
	whateverInt := models.BattleAction(33)

	type testCase struct {
		name             string
		weapon           *Item
		action           models.BattleAction
		expectedErr      string
		expectedRetValue float64
		shouldPanic      bool
	}

	tests := []testCase{
		{
			name:             "Empty Weapon - Correct Action",
			weapon:           emptyItem,
			action:           correctAction,
			expectedErr:      models.NewNotCorrectlyInitError(emptyItem).Error(),
			expectedRetValue: -1,
			shouldPanic:      false,
		},
		{
			name:             "Empty Weapon - Wrong Action",
			weapon:           emptyItem,
			action:           wrongAction,
			expectedErr:      models.NewNotCorrectlyInitError(emptyItem).Error(),
			expectedRetValue: -1,
			shouldPanic:      false,
		},
		{
			name:             "Correct Weapon - Non Existing Action",
			weapon:           correctWeapon,
			action:           whateverInt,
			expectedErr:      models.NewWrongBattleActionError([]models.BattleAction{models.LIGHT_ATTACK, models.HEAVY_ATTACK}, whateverInt).Error(),
			expectedRetValue: -1,
			shouldPanic:      true,
		},
		{
			name:             "Wrong Weapon - Correct Action",
			weapon:           wrongWeapon,
			action:           correctAction,
			expectedErr:      models.NewWrongItemType([]models.ItemType{models.WEAPON}, wrongWeapon.ItemType).Error(),
			expectedRetValue: -1,
			shouldPanic:      false,
		},
		{
			name:             "Correct Weapon - Wrong Action",
			weapon:           correctWeapon,
			action:           wrongAction,
			expectedErr:      models.NewWrongBattleActionError([]models.BattleAction{models.LIGHT_ATTACK, models.HEAVY_ATTACK}, wrongAction).Error(),
			expectedRetValue: -1,
			shouldPanic:      false,
		},
		{
			name:             "Correct Weapon - Correct Action",
			weapon:           correctWeapon,
			action:           correctAction,
			expectedErr:      "",
			expectedRetValue: correctWeapon.Weight * 2,
			shouldPanic:      false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := models.GetRequiredStamina(tc.weapon, tc.action)

			if err != nil {
				if err.Error() != tc.expectedErr {
					t.Errorf("expected error -> (%q) | but got -> (%q)", tc.expectedErr, err.Error())
				} else {
					t.Logf("PASS: %s", tc.name)
				}
			} else {
				if tc.expectedRetValue != result {
					t.Errorf("expected result -> (%.2f) | but got -> (%.2f)", tc.expectedRetValue, result)
				} else {
					t.Logf("PASS: %s, RESULT: (%.2f)", tc.name, result)
				}
			}
		})
	}
}
