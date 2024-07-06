package numbers

import "fmt"

func Numbers() {
	fmt.Print("3. Running numbers.go\n\n")

	var tinyUint uint8 = 255
	var smallUint uint16 = 65535
	var mediumUint uint32 = 4294967295
	var bigUint uint64 = 18446744073709551615

	var tinyInt int8 = 127
	var smallInt int16 = 32767
	var mediumInt int32 = 2147483647
	var bigInt int64 = 9223372036854775807

	var float32Var float32 = 3.40282346638528859811704183484516925440e+38
	var float64Var float64 = 1.797693134862315708145274237317043567981e+308

	var complex64Var complex64 = 1 + 2i
	var complex128Var complex128 = 1 + 2i

	fmt.Println("3.1 Unsigned Integers")

	fmt.Println("3.1.1 uint8: ", tinyUint)
	fmt.Println("3.1.2 uint16: ", smallUint)
	fmt.Println("3.1.3 uint32: ", mediumUint)
	fmt.Println("3.1.4 uint64: ", bigUint)

	fmt.Println("3.2 Signed Integers")

	fmt.Println("3.2.1 int8: ", tinyInt)
	fmt.Println("3.2.2 int16: ", smallInt)
	fmt.Println("3.2.3 int32: ", mediumInt)
	fmt.Println("3.2.4 int64: ", bigInt)

	fmt.Println("3.3 Floating Point Numbers")

	fmt.Println("3.3.1 float32: ", float32Var)
	fmt.Println("3.3.2 float64: ", float64Var)

	fmt.Println("3.4 Complex Numbers")

	fmt.Println("3.4.1 complex64: ", complex64Var)
	fmt.Println("3.4.2 complex128: ", complex128Var)
}
