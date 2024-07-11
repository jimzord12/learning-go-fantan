package mytypes

import (
	"fmt"
)

/// CUSTOM TYPES ///

type AccountNumber string

// Adding a Method to a Type
func (a AccountNumber) IsValid() bool {
	return len(a) == 10
}

// 2nd Example
type ByteSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

// Here we extend the ByteSize Type by providing a String() method,
// Now it implements the Stringer Interface
func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}

// / TYPES ALIAS ///
type Balance = float64


/// TYPE ASSERTION ///
var i any = "Hello" // var i interface{} = "Hello"

func Main() {
	accNum := AccountNumber("0123456789")
	a := accNum.IsValid()
	fmt.Println(a)
	fmt.Println(5e3) // => 5 * 10^3
	fmt.Println(5*10 ^ 3)

	// fmt.Println(KB)               // Uses ByteSize.String() method, prints "1.00KB"
	// fmt.Println(MB)               // Uses ByteSize.String() method, prints "1.00MB"
	// fmt.Println(GB)               // Uses ByteSize.String() method, prints "1.00GB"
	// fmt.Println(ByteSize(2 * GB)) // Uses ByteSize.String() method, prints "2.00GB"
	// fmt.Println(ByteSize(1525))
}


