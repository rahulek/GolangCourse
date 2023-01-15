package main

import (
	"fmt"
)

func main() {
	fmt.Println("---- Variable Declarations Start ----")

	simpleDeclarationsDemo("Simple Declarations")
	packageLevelDemo("Package Level Declarations")
	packageLevelBlockDemo("Package Level BLOCK Declarations")
	shadowedDeclarationsDemo("Shadowed Variable Declarations")
	declarationErrorsDemo("Declaration Errors")
	namingAndVisibilityDemo("Variable Naming and Visibility")
	typeConversionDemo("Variable Type Conversions")

	fmt.Println("---- Variable Declarations End ---")

}
