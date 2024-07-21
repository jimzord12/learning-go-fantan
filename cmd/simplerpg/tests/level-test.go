package tests

import (
	"fmt"
	"log"

	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models"
)

// type FuncWrapper interface {
// 	Call()
// }

// type FuncWrap struct {
// 	Fn func(...any)
// }

func TestLeveling(args ...any) {
	var player *Character
	var ok bool

	fmt.Println("EXP Requirements for 1 to 10")

	if player, ok = args[0].(*Character); !ok {
		// Debug: print the type of args[0]
		fmt.Printf("Type of args[0]: %T\n", args[0])
		log.Fatalf("[ERROR]: TestLeveling type assertion OK: (%v)", ok)
	}

	for level := 1; level <= 10; level++ {
		fmt.Printf("Level %d -> %d: %.2f EXP required\n", level, level+1, models.ExpForNextLevel(level))
	}

	// player := models.NewPlayer("PL-1", "John Wick", models.HUMAN)
	// enemy := models.NewArcDemonEnemy("BS-01", 1)

	fmt.Println()
	fmt.Printf("Player's current EXP: (%.2f)\n", player.Exp)
	fmt.Println()

	// expFromBattle := models.ExpGainedFromEnemy(player.Level, enemy)
	fmt.Printf("[TESTING]: Giving 5000 EXP to Player")
	player.GainEXP(5000)

	fmt.Println()
	fmt.Printf("Player's current EXP: (%.2f)\n", player.Exp)
	fmt.Println()
}
