package main

import (
	"BuiltInTypes/Boolean"
	"BuiltInTypes/Complex"
	"BuiltInTypes/Floatingtypes"
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

	Floatingtypes.DeclarationDemo("Declaring floating point/decimals")
	Floatingtypes.StdOperationsDemo("Arithmetic operations on floating point variables")

	Complex.DeclarationDemo("Complex Data variables declarations")
	Complex.ComposeDecomposeDemo("Compose and Decompose Complex numbers")

	fmt.Println("----- Builtin Types End -----")

}
