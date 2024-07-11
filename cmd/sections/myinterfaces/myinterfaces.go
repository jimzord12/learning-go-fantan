package myinterfaces

import "fmt"

type baseInfo struct {
	model    string
	brand    string
	topSpeed float64
}

type gasEngine struct {
	baseInfo
	mpg     uint8 // how much Miles does a Gallon permit
	gallons uint8
}

type electricEngine struct {
	baseInfo
	mpkwh uint8 // how much Miles does a KWh permit
	kwh   uint8
}

func (g gasEngine) milesLeft() uint8 {
	return g.gallons * g.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

// / THE INTERFACE ///
type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) bool {
	if miles <= e.milesLeft() {
		return true
	} else {
		return false
	}
}

func Main() {
	engine1 := gasEngine{
		baseInfo: baseInfo{
			model:    "v1",
			brand:    "Toyota",
			topSpeed: 250,
		},
		mpg:     10,
		gallons: 25,
	}

	engine2 := electricEngine{
		baseInfo: baseInfo{model: "v2", brand: "Tesla", topSpeed: 210},
		mpkwh:    12,
		kwh:      20,
	}

	fmt.Println("Can Engine 1 make it?", canMakeIt(engine1, 150))
	fmt.Println("Can Engine 2 make it?", canMakeIt(engine2, 255))
}
