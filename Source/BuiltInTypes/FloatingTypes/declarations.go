package Floatingtypes

import (
	"fmt"
	"math"
)

// Function DeclarationDemo accepts the title parameter
// and demonstrates the declaration of the floating data types
func DeclarationDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	var pi = 3.1415927
	fmt.Printf("\tpi is default float: %v (%[1]T)\n", pi)

	//Explcit float32
	var piAnother float32 = 3.1415927
	fmt.Printf("\tpiAnother is explicit float32: %v (%[1]T)\n", piAnother)

	//Ranges
	fmt.Printf("\tfloat32 range: %.5e -> %.5e\n", math.SmallestNonzeroFloat32, math.MaxFloat32)
	fmt.Printf("\tfloat64 range: %.5e -> %.5e\n", math.SmallestNonzeroFloat64, math.MaxFloat64)

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}
