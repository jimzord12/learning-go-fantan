package fantan

import (
	"fmt"
	"hello-world/cmd/sections/userinput/fantan/helper"
	"hello-world/cmd/sections/userinput/fantan/types"
	"os"
)

func PlayFantan(balance float64) {

	// Creating Player
	player := types.Player{
		Games:   0,
		Wins:    0,
		Losses:  0,
		Profit:  0,
		Balance: balance,
	}

	for {
		if player.Games < 1 { // 1st Game
			println("++++++++++++++++")
			FanTanGame(&player)
		} else { // Next Games
			println("-----------------")
			FanTanRound(&player)
		}

	}

	// Next Games
	// 	for {
	// 		fmt.Println("To Play Again press 'p/P' | To Quit press 'q/Q'")
	// 		var userChoise string
	// 		fmt.Scanln(&userChoise)
	// 		if userChoise == "p" || userChoise == "P" {
	// 			result = FanTanGame(player.Balance)
	// 			if result > 0 {
	// 				player.Wins += 1
	// 			} else {
	// 				player.Losses += 1
	// 			}
	// 			player.Profit += result
	// 			player.Balance += result

	// 			fmt.Println("-----------------------------------------")
	// 			fmt.Println("WINS: ", player.Wins, " | LOSSES: ", player.Losses)
	// 			fmt.Println("-----------------------------------------")
	// 			fmt.Println("Profit: (", player.Profit, ")")
	// 			fmt.Println("-----------------------------------------")
	// 			fmt.Println("Balance: (", player.Balance, ")")
	// 			fmt.Println("-----------------------------------------")

	// 			if player.Balance <= 0 {
	// 				fmt.Println("Thank you for giving us your Money!!!")
	// 				fmt.Println("Have a good day and come again soon!")
	// 				os.Exit(0)
	// 			}
	// 		} else if userChoise == "q" || userChoise == "Q" {
	// 			fmt.Println("Thank you for Playing!")
	// 			os.Exit(0)
	// 		} else {
	// 			fmt.Println("Please enter either 'p/P' or 'q/Q'...")
	// 		}
	// 	}
}

func FanTanRound(player *types.Player) {
	fmt.Println("-----------------------------------------")
	fmt.Println("WINS: ", player.Wins, " | LOSSES: ", player.Losses)
	fmt.Println("-----------------------------------------")
	fmt.Println("Profit: (", player.Profit, ")")
	fmt.Println("-----------------------------------------")
	fmt.Println("Balance: (", player.Balance, ")")
	fmt.Println("-----------------------------------------")

	if player.Balance <= 1.0 {
		fmt.Println("Thank you for giving us your Money!!!")
		fmt.Println("Have a good day and come again soon!")
		os.Exit(0)
	}

	const playAgainMsg = "To Quit press 'q' and to Play Again press 'p'..."
	userInput := helper.GetAndValidateInput(helper.GetUserInput, playAgainMsg, "p", "P", "q", "Q")

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
