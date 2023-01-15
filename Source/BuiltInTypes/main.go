package main

import (
	"BuiltInTypes/Boolean"
	"BuiltInTypes/Integertypes"
	"fmt"
)

func main() {
	fmt.Println("----- Builtin Types Start -----")

	Boolean.BoolDeclarationDemo("Declaring the Booleans")
	Boolean.BooleanUsageDemo("Using the Boolean Variables")

	Integertypes.DeclarationDemo("Declaring integer variables")
	Integertypes.StdOperationsDemo("Standard operations on integers")
	Integertypes.BitwiseOperationsDemo("Bitwise operations on integers")
	Integertypes.ShiftOperationsDemo("Bit Shift operations on integers")

	fmt.Println("----- Builtin Types End -----")

}
