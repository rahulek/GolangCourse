package Boolean

import "fmt"

// BooleanUsageDemo functions accept the title string and then
// demonstrates the usage of the boolean variables
func BooleanUsageDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//Booleans are mostly used in conditional logic
	//where the code block is executes depending upon
	//the true/false value of the boolean variable

	t := true
	f := false

	if t {
		// t is true
		fmt.Println("\tt is TRUE")
	}

	//! operator negates true to false and vice versa
	if !f {
		//f is false
		fmt.Println("\tf is FALSE")
	}

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}
