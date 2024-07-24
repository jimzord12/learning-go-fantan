package battlehelpers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models"
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/inputhelpers"
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/logging"
)

func GetBattleAction(player *models.Character) (models.BattleAction, *models.Item, error) {
	input := inputhelpers.GetTerminalInput("Select a Battle Action:\n1. LIGHT ATTACK\n2. HEAVY ATTACK\n3. DEFEND (Not Supported yet)\n4. REST\n5. HEAL")

	selectedAction, err := strconv.Atoi(input)
	if err != nil {
		logging.LogError(logging.Logger, "(GetBattleAction) -> Error while trying to convert input string to int")
		return -1, nil, errors.New("(GetBattleAction) -> Error while trying to convert input string to int")
	}

	if selectedAction < 1 || selectedAction > 5 {
		logging.LogError(logging.Logger, "(GetBattleAction) -> User provided an int which is less than 1 or greater than 5")
		return -1, nil, errors.New("(GetBattleAction) -> User provided an int which is less than 1 or greater than 5")
	}

	if selectedAction == int(models.HEAL) {
		potionStock, amountOfPotions := player.Inventory.GetPotionStock()

		if amountOfPotions == 0 {
			return -1, nil, errors.New("(GetBattleAction) -> HEAL -> User does not have any potions")
		}

	

		msg := fmt.Sprintf("Select a Potion Type:\n1. SMALL (%d)\n2. MEDIUM (%d)\n3. LARGE (%d)", potionStock[models.SMALL], potionStock[models.MEDIUM], potionStock[models.LARGE])

		playerPotionSelectionStr := inputhelpers.GetTerminalInput(msg)

		playerPotionSelection, err := strconv.Atoi(playerPotionSelectionStr)
		if err != nil {
			logging.LogError(logging.Logger, "(GetBattleAction) -> (HEAL) -> Error while trying to convert input string to int")
			return -1, nil, errors.New("(GetBattleAction) -> (HEAL) -> Error while trying to convert input string to int")
		}

		if playerPotionSelection < 1 || playerPotionSelection > 3 {
			logging.LogError(logging.Logger, "(GetBattleAction) -> (HEAL) -> User provided an int which is less than 1 or greater than 3")
			return -1, nil, errors.New("(GetBattleAction) -> (HEAL) -> User provided an int which is less than 1 or greater than 3")
		}

		amountOfSpecificPotion := potionStock[models.PotionType(playerPotionSelection)]
		if amountOfSpecificPotion < 1 {
			logging.LogError(logging.Logger, "(GetBattleAction) -> (HEAL) -> User selected a Potion Type whose amount is less than 1")
			return -1, nil, errors.New("(GetBattleAction) -> (HEAL) -> User selected a Potion Type whose amount is less than 1")
		}

		potion, err := player.Inventory.FindItem(models.PotionTypesToNames[models.PotionType(playerPotionSelection)])
		if err != nil {
			logging.LogError(logging.Logger, "(GetBattleAction) -> (HEAL) -> User selected a Potion Type whose amount is less than 1")
			return -1, nil, err
		}

		return models.BattleAction(selectedAction), potion, nil
	}
	return models.BattleAction(selectedAction), nil, nil

}
