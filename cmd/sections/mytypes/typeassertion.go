package mytypes

import (
	"fmt"
)

type sword struct {
	Weight int
	Damage float64
}

type shield struct {
	Weight  int
	Defence float64
}

func do(i any) {
	switch v := i.(type) {
	case int:
		fmt.Println("The parameter is of type (Int) | ", v*2)
	case float64:
		fmt.Println("The parameter is of type (float64) | ", v*2.5)
	case string:
		fmt.Println("The parameter is of type (string) | ", len(v))
	case sword:
		fmt.Println("The parameter is of type (sword) | ", v.Damage)
	default:
		fmt.Printf("The parameter's type is not supported! | %T\n", i)

	}
}

func Main2() {
	// var i interface{} = "Hello"
	var i any = "Hello"

	s := i.(string)
	fmt.Println("1:", s)

	s, ok := i.(string)
	fmt.Println("2:", s, ok)

	f, ok := i.(float64)
	fmt.Println("3:", f, ok)

	// f1 := i.(float64) // It Panics!
	// fmt.Println("4:", f1)

	/// Asserting Types with Switch Statement
	do(21)
	do(21.22)
	do("21")
	do(sword{
		Weight: 125,
		Damage: 502.5,
	})
	do(shield{
		Weight:  512,
		Defence: 150.2,
	})

}
