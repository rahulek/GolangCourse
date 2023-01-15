package Complex

import (
	"fmt"
)

// ComposeDecomposeDemo accepts a title parameter
// and demonstrates how to get real and imaginary portions
// of the complex number. It also demonstrates composing
// a complex number from the real and imaginary parts
func ComposeDecomposeDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	c1 := 100 + 20i
	fmt.Printf("\tDECOMPOSE: c1: %v, real(c1): %v, imag(c1): %v\n", c1, real(c1), imag(c1))

	c2 := complex(99, 1)
	fmt.Printf("\tCOMPOSE: real: %v, imag: %v, complex: %v\n", 99, 1, c2)

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}
