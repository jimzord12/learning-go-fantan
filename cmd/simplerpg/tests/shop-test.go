package tests

import (
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models"
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/logging"
)

func SetUpShop(args ...any) {
	player, ok := args[0].(*Character)

	if !ok {
		logging.LogError(logging.Logger, "(func GivePotions(args ...any)) -> Issue with type assertion")
	}

	logging.GiveVertSpace("======== SETTING UP SHOP ========")

	shop := models.NewShop(player)

	shop.DisplayGoods()

	logging.GiveVertSpace(" (1) Giving 200 Gold to Player")
	player.Gold = 200

	logging.GiveVertSpace(" (2) Player attempts to buy the a Weapon")

	shop.Buy(player, shop.Equipment[1])

	player.DisplayInventory()

	logging.GiveVertSpace(" (3) Player attempts to sell the a Potion")

	shop.Sell(player, player.Inventory.Items[len(player.Inventory.Items)-1])

	player.DisplayInventory()

}
