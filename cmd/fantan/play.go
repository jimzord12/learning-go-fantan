package fantan

import (
	"fmt"
	"hello-world/cmd/fantan/helper"
	"hello-world/cmd/fantan/types"
	"os"
)

// balance will come from DB
func PlayFantan(balance float64) {

	// Creating Player
	player := types.Player{
		Games:   0,
		Wins:    0,
		Losses:  0,
		Profit:  0,
		Balance: balance,
	}

	// Looping for the ability to play multiple rounds of the game.
	for {
		if player.Games < 1 { // 1st Game
			println("++++++++++++++++")
			FanTanGame(&player)
		} else { // Next Games
			println("-----------------")
			FanTanRound(&player)
		}

	}
}

func FanTanRound(player *types.Player) {
	printRoundStats(player)

	if player.Balance <= 1.0 {
		fmt.Println("Thank you for giving us your Money!!!")
		fmt.Println("Have a good day and come again soon!")
		os.Exit(0)
	}

	const playAgainMsg = "To Quit press 'q' and to Play Again press 'p'..."
	userInput := helper.GetAndValidateInput(helper.GetUserInput, playAgainMsg, []string{"p", "P", "q", "Q"})

	switch userInput {
	case "p", "P":
		FanTanGame(player)
	case "q", "Q":
		fmt.Println("Goodbye, Thank you for Playing!")
		os.Exit(0)
	default:
		panic(1)
	}

}

func printRoundStats(player *types.Player) {
	fmt.Println("-----------------------------------------")
	fmt.Println("WINS: ", player.Wins, " | LOSSES: ", player.Losses)
	fmt.Println("-----------------------------------------")
	fmt.Println("Profit: (", player.Profit, ")")
	fmt.Println("-----------------------------------------")
	fmt.Println("Balance: (", player.Balance, ")")
	fmt.Println("-----------------------------------------")
}
