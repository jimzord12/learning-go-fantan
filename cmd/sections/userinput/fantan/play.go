package fantan

import (
	"fmt"
	"hello-world/cmd/sections/userinput/fantan/helper"
	"hello-world/cmd/sections/userinput/fantan/types"
	"os"
)

func PlayFantan() {
	fmt.Print("Running main.go\n\n")
	// var wins, losses int
	// var profit float64
	// var deposit int
	player := types.Player{
		Wins:    0,
		Losses:  0,
		Profit:  0,
		Balance: 0,
	}

	player.Balance = float64(helper.GetAndParseUserInputInt("Plzzz Give me Money"))
	fmt.Print("Thank you for your money! (", player.Balance, ")\n\n")

	// 1st Game
	result := FanTanGame(player.Balance)
	if result > 0 {
		player.Wins += 1
	} else {
		player.Losses += 1
	}
	player.Profit += result
	player.Balance += result

	fmt.Println("-----------------------------------------")
	fmt.Println("WINS: ", player.Wins, " | LOSSES: ", player.Losses)
	fmt.Println("-----------------------------------------")
	fmt.Println("Profit: (", player.Profit, ")")
	fmt.Println("-----------------------------------------")
	fmt.Println("Balance: (", player.Balance, ")")
	fmt.Println("-----------------------------------------")

	if player.Balance <= 0 {
		fmt.Println("Thank you for giving us your Money!!!")
		fmt.Println("Have a good day and come again soon!")
		os.Exit(0)
	}

	fmt.Println("To Quit press 'q' and to Play Again press 'p'...")

	// Next Games
	for {
		fmt.Println("To Play Again press 'p/P' | To Quit press 'q/Q'")
		var userChoise string
		fmt.Scanln(&userChoise)
		if userChoise == "p" || userChoise == "P" {
			result = FanTanGame(player.Balance)
			if result > 0 {
				player.Wins += 1
			} else {
				player.Losses += 1
			}
			player.Profit += result
			player.Balance += result

			fmt.Println("-----------------------------------------")
			fmt.Println("WINS: ", player.Wins, " | LOSSES: ", player.Losses)
			fmt.Println("-----------------------------------------")
			fmt.Println("Profit: (", player.Profit, ")")
			fmt.Println("-----------------------------------------")
			fmt.Println("Balance: (", player.Balance, ")")
			fmt.Println("-----------------------------------------")

			if player.Balance <= 0 {
				fmt.Println("Thank you for giving us your Money!!!")
				fmt.Println("Have a good day and come again soon!")
				os.Exit(0)
			}
		} else if userChoise == "q" || userChoise == "Q" {
			fmt.Println("Thank you for Playing!")
			os.Exit(0)
		} else {
			fmt.Println("Please enter either 'q' or 'y'...")
		}
	}
}
