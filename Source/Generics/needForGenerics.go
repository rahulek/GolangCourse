package main

import "fmt"

//Go Generics are available from Go v1.18 onwards.
//Earlier Go versions donot support Generics

// NeedForGenericsDemo function accepts title parameter
// and demonstrates the need for the Generics
func NeedForGenericsDemo(title string) {
	fmt.Printf("*** %v : START ***\n", title)
	fmt.Printf("*** %v : START ***\n", title)

	//Without Generics, we need to write a code that handles one type of
	//inputs only. This introduces code repetition and related bugs
	fmt.Printf("\tsumAllInts(1, 2, 3, 4, 5): %v\n", sumAllInts(1, 2, 3, 4, 5))
	fmt.Printf("\tsumAllFloats(1.1, 2.2, 3.3, 4.4, 5.5): %.2f\n", sumAllFloats(1.1, 2.2, 3.3, 4.4, 5.5))

	fmt.Printf("*** %v : START ***\n", title)
	fmt.Printf("*** %v : START ***\n", title)
}

// Sum All Integer inputs and return an integer result
func sumAllInts(intInputs ...int) int {
	sum := 0
	for _, input := range intInputs {
		sum += input
	}

	return sum
}

// Sum All Float inputs and return an Float result
func sumAllFloats(floatInputs ...float64) float64 {
	sum := 0.
	for _, input := range floatInputs {
		sum += input
	}
	return sum
}
