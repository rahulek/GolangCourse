package Floatingtypes

import "fmt"

// Function StdOperationsDemo accepts a title parameter
// and demonstrates the standard arithmetic operations
// of the floating point variables.
func StdOperationsDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	f1 := 12.34
	f2 := 2.01
	fmt.Printf("\tAddition: f1: %.2f, f2: %.2f, f1 + f2 = %.2f\n", f1, f2, f1+f2)
	fmt.Printf("\tSubtraction: f1: %.2f, f2: %.2f, f1 - f2 = %.2f\n", f1, f2, f1-f2)
	fmt.Printf("\tMultiplication: f1: %.2f, f2: %.2f, f1 * f2 = %.2f\n", f1, f2, f1*f2)
	fmt.Printf("\tDivision: f1: %.2f, f2: %.2f, f1 / f2 = %.2f\n", f1, f2, f1/f2)

	//There is no % on floating data variables
	//There are no shift operations on floating data variables
	//There are no bitwise operations on floating data variables

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}
