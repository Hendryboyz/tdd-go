package main

import "fmt"

func sumInts(m map[string]int64) int64 {
	var sum int64
	for _, val := range m {
		sum += val
	}
	return sum
}

func sumFloats(m map[string]float64) float64 {
	var sum float64
	for _, val := range m {
		sum += val
	}
	return sum
}

type Number interface {
	int64 | float64
}

// func genericFunction[type parameters](ordinary function arguments) returnType {}
func sumIntsOrFloats[K comparable, V Number](m map[K]V) V {
	var sum V
	for _, val := range m {
		sum += val
	}
	return sum
}

func main() {
	ints := map[string]int64{
		"first":  1,
		"second": 2,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf(
		"Non-Generic Sums: %v and %v\n",
		sumInts(ints),
		sumFloats(floats))

	fmt.Printf(
		"Generic Sums: %v and %v\n",
		sumIntsOrFloats[string, int64](ints),
		sumIntsOrFloats(floats))

}
