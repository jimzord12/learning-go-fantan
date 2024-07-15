package generalhelpers

import "math"

func Round2Dec(num float64) float64 {
	return math.Round(num*100) / 100
}
