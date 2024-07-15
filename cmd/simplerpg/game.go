package simplerpg

import (
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/logging"
	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/tests"
)

func RunGame() {
	// Create a logger that writes to the standard output
	logging.LogInit()

	tests.TestWeapons()
	logging.StdDivider("*", 100)
	tests.TestArmors()
	logging.StdDivider("*", 100)
	tests.TestPotions()
}
