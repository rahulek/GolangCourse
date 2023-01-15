package Complex

import "fmt"

// DeclarationsDemo function accepts title parameter
// and demonstrates the declaration of the complex data variables.
func DeclarationDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	c1 := 12 + 7i
	c2 := 10i
	fmt.Printf("\tc1: %v (%[1]T)\n", c1)
	fmt.Printf("\tc2: %v (%[1]T)\n", c2)

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}
