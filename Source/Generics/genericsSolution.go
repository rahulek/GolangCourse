package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// GenericsSolutionDemo function accepts title parameter
// and demonstrates the use of Go Generics
func GenericsSolutionDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	resultIntSummation := sumAll[int](1, 2, 3, 4, 5)
	resultFloatSummation := sumAll[float64](1.1, 2.2, 3.3, 4.4, 5.5)
	fmt.Printf("\tWith Generics: intSummation: %v, floatSummation: %.2f\n", resultIntSummation, resultFloatSummation)

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}

// Generic Summation
func sumAll[T constraints.Ordered](inputs ...T) T {
	var sum T
	for _, input := range inputs {
		sum += input
	}
	return sum
}
