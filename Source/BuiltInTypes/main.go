package main

import (
	"BuiltInTypes/Boolean"
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

	fmt.Println("----- Builtin Types End -----")

}
