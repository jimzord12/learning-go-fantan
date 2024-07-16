package tests

import (
	"fmt"

	levelsystem "github.com/jimzord12/learning-go-fantan/cmd/simplerpg/mechanics/level-system"
)

func TestLeveling() {
	for level := 1; level <= 10; level++ {
		fmt.Printf("Level %d: %.2f EXP required\n", level, levelsystem.CalcRequiredExp(level))
	}
}
