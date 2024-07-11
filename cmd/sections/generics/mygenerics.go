package generics

import "fmt"

/// GENERIC FUNC

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

/// Another Example from the STD Lib
func DeleteFunc[E any, S ~[]E](s S, del func(E) bool) S {
	for i, v := range s {
		if del(v) {
			j := 1
			for i++; i < len(s); i++ {
				v = s[i]
				if !del(v) {
					s[j] = v
					j++
				}
			}
			return s[:j]
		}
	}
	return s
}

/// GENERIC TYPES

type Pair[K, V any] struct {
	Key   K
	Value V
}

func (p Pair[K, V]) Describe() string {
	return fmt.Sprintf("Key:%v, Value: %v", p.Key, p.Value)
}

func Main() {
	intStringPair := Pair[int, string]{Key: 1, Value: "hello"}
	fmt.Println(intStringPair.Describe())

	stringFloatPair := Pair[string, float64]{Key: "Pi", Value: 3.141}
	fmt.Println(stringFloatPair.Describe())
}
