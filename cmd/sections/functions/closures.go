package functions

type fib func() int // 1st Try, no Docs or AI :P

func FibonacciSequence(a int, b int) fib {
	return func() int {
		a, b = b, a+b
		return a
	}
}
