package tests

import (
	"fmt"

	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models"
)

func TestLeveling() {
	for level := 1; level <= 10; level++ {
		fmt.Printf("Level %d: %.2f EXP required\n", level, models.CalcRequiredExp(level))
	}
}
