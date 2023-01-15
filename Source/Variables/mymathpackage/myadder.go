package mymathpackage

import "fmt"

// NOTE: Capitalized var is exported from this package
// and visible to other packages inside this module.
var MyMathPI = 3.1415927

// NOTE: Capitalized function is exported from this package
// and visible to other packages inside this module.
func MyAdder(a, b int) int {
	return a + b
}

// NOTE: Unexported because its first letter is not Capitalized
func cantBeSeen() {
	fmt.Println("This is a cantBeSeen function.")
}
