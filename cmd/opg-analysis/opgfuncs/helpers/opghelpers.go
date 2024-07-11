package helpers

import "math"

func Round2Dec(value float64) float64 {
	return math.Round(value*100) / 100
}
