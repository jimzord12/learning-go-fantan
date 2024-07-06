package numbers

import (
	"fmt"
	"math/rand"
	"time"
)

func KelvinToCelsius(input int8) float64 {
	fmt.Println("Temperature in Celsius: ", input)
	kelvinTemp := float64(input) - 273.15

	// roundedTemp := math.Round(kelvinTemp*100) / 100
	celsius := fmt.Sprintf("(Rounded) Temperature in Kelvin: %.2f \u00B0C", kelvinTemp)

	println(celsius)

	return kelvinTemp

}

func RandomNumbers() {
	for {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		fmt.Println("-> : ", r.Int31n(4)+1)
		time.Sleep(1 * time.Second)
	}

}
