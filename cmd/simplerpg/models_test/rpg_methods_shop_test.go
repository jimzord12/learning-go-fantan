package modelstest

import (
	"reflect"
	"testing"

	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models"
)

type Item = models.Item
type BattleAction = models.BattleAction

func TestNewShop(t *testing.T) {
	yoda := models.NewPlayer("PL-001", "Yoda", models.ELF, 999)
	correctShop, errYoda := models.NewShop(yoda)

	type testType struct {
		name          string
		player        *models.Character
		expectedErr   error
		expectedValue *models.Shop
	}

	testCases := []testType{
		{
			name:          "Fail - Empty Player",
			player:        &models.Character{},
			expectedErr:   models.NewNotCorrectlyInitError(&models.Character{}),
			expectedValue: nil,
		},
		{
			name:          "Success - Shop Creation",
			player:        yoda,
			expectedErr:   errYoda,
			expectedValue: correctShop,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := models.NewShop(tc.player)
			if err != nil {
				if reflect.TypeOf(tc.expectedErr) != reflect.TypeOf(err) {
					t.Errorf("expected -> (%s),\nbut got -> (%s)", tc.expectedErr, err)
				} else {
					t.Logf("PASS: (%s)", tc.name)
				}
			} else {
				if err != tc.expectedErr {
					t.Errorf("expected -> (%s),\nbut got -> (%s)", tc.expectedErr, err)
				} else {
					t.Logf("PASS: (%s)", tc.name)
				}
			}
		})
	}
}

func TestBuy(t *testing.T) {
	player := models.NewPlayer("PL-001", "Pl_Test_1", models.ELF, 5)
	emptyPlayer := &models.Character{}

	shop, _ := models.NewShop(player)

	itemToBuy := models.NewWeapon(models.DAGGER, models.MYTHRIL)
	// wrongItem := &Item{
	// 	Name:     "Non Existing Item",
	// 	Weight:   500,
	// 	ItemType: models.ItemType(27), // Does not exist in Enum
	// 	Material: models.BRONZE,
	// 	Value:    200,
	// }
	emptyItem := &Item{}

	type testType struct {
		name        string
		shop        *models.Shop
		player      *models.Character
		item        *Item
		expectedErr string
	}

	testCases := []testType{
		{
			name:        "Not Enough Gold",
			shop:        shop,
			player:      player,
			item:        itemToBuy,
			expectedErr: models.NewNotEnoughGoldError(itemToBuy.GetGoldBuyValue(), player.Gold).Error(),
		},
		{
			name:        "Empty Player",
			shop:        shop,
			player:      emptyPlayer,
			item:        itemToBuy,
			expectedErr: models.NewNotCorrectlyInitError(emptyPlayer).Error(),
		},
		// {
		// 	name:        "Wrong Item Type",
		// 	shop:        shop,
		// 	player:      player,
		// 	item:        wrongItem,
		// 	expectedErr: models.NewWrongItemType(models.AllItemTypes, wrongItem.ItemType).Error(),
		// },
		{
			name:        "Empty Item",
			shop:        shop,
			player:      player,
			item:        emptyItem,
			expectedErr: models.NewNotCorrectlyInitError(emptyItem).Error(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.shop.Buy(tc.player, tc.item); err != nil {
				if err.Error() != tc.expectedErr {
					t.Errorf("expected error -> (%s), but got -> (%s)", tc.expectedErr, err.Error())
				} else {
					t.Logf("PASS: %s", tc.name)
				}
			} else {
				t.Logf("PASS: Error is nil, (%s)", tc.name)

			}
		})
	}
}

func TestSell(t *testing.T) {
	player := models.NewPlayer("PL-001", "Pl_Test_1", models.ELF, 5)
	playerWithEmptyInv := models.NewPlayer("PL-002", "Pl_Test_2", models.ELF, 5)
	emptyPlayer := &models.Character{}

	shop, _ := models.NewShop(player)

	itemToSell := models.NewWeapon(models.DAGGER, models.MYTHRIL)

	player.MoveToInventory(itemToSell)

	// wrongItem := &Item{
	// 	Name:     "Non Existing Item",
	// 	Weight:   500,
	// 	ItemType: models.ItemType(27), // Does not exist in Enum
	// 	Material: models.BRONZE,
	// 	Value:    200,
	// }
	emptyItem := &Item{}

	type testType struct {
		name        string
		shop        *models.Shop
		player      *models.Character
		item        *Item
		expectedErr string
	}

	testCases := []testType{
		{
			name:        "Correct Item Sale",
			shop:        shop,
			player:      player,
			item:        itemToSell,
			expectedErr: "models.NewNotEnoughGoldError(itemToBuy.GetGoldBuyValue(), player.Gold).Error()",
		},
		{
			name:        "Item Sale Fails - Empty Player",
			shop:        shop,
			player:      emptyPlayer,
			item:        itemToSell,
			expectedErr: models.NewNotCorrectlyInitError(emptyPlayer).Error(),
		},
		{
			name:        "Item Sale Fails - Item Does Not Exist",
			shop:        shop,
			player:      playerWithEmptyInv,
			item:        itemToSell,
			expectedErr: models.NewDoesNotExistInInventory(itemToSell, playerWithEmptyInv).Error(),
		},
		// {
		// 	name:        "Wrong Item Type",
		// 	shop:        shop,
		// 	player:      player,
		// 	item:        wrongItem,
		// 	expectedErr: models.NewWrongItemType(models.AllItemTypes, wrongItem.ItemType).Error(),
		// },
		{
			name:        "Empty Item",
			shop:        shop,
			player:      player,
			item:        emptyItem,
			expectedErr: models.NewNotCorrectlyInitError(emptyItem).Error(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.shop.Sell(tc.player, tc.item); err != nil {
				if err.Error() != tc.expectedErr {
					t.Errorf("expected error -> (%s), but got -> (%s)", tc.expectedErr, err.Error())
				} else {
					t.Logf("PASS: %s", tc.name)
				}
			} else {
				t.Logf("PASS: Error is nil, (%s)", tc.name)

			}
		})
	}
}
