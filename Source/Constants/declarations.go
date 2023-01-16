package main

import (
	"fmt"
	"math"
)

// A package-scoped const
const packageScopedUnexportedConst = "UnExportedPackageScoped" //unexported due to a unCapitailized name
const ExportedConst = math.Pi                                  //exported because of the Capitalized name
const shadowedConst = 999

// ConstantDeclarationsDemo accepts a title parameter
// and demonstrates the declaration of various constants.
func ConstantDeclarationsDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v: START ***\n", title)

	//A simple const can be declared in local scope
	//No type defined explicitly. It will be inferred by the compiler
	const localConst = 9876
	fmt.Printf("\tPackage Level Unexported Const: %v (%[1]T)\n", packageScopedUnexportedConst)
	fmt.Printf("\tPackage Level Exported Const: %v (%[1]T)\n", ExportedConst)
	fmt.Printf("\tLocal Scope Const: %v (%[1]T)\n", localConst)

	//A const can be defined with an explicit type too
	const explicitConst uint8 = 112
	fmt.Printf("\tExplicit Local Const: %v (%[1]T)\n", explicitConst)

	//Const can't change (it's constant)
	//Uncomment below line to see the error
	//explicitConst = 92

	//A const MUST BE a compile-time constant
	//That is: it must be a value that's either a literal constant
	//OR calculated at compile-time
	const expressionBasedConst int = (int(explicitConst) * 2)
	fmt.Printf("\tExpression based Const: %v (%[1]T)\n", expressionBasedConst)

	//But the following won't compile
	//because we're trying to execute a function (runtime process) to derive
	//the const's value
	//Uncomment below line to see the error
	//const erroredOut = strings.Clone("Hi")

	//constant can be of any built-in type
	const alwaysFun bool = false
	fmt.Printf("\talwaysFun %v (%[1]T)\n", alwaysFun)

	//Package level const (shadowedConst above is an int type) can be shadowed
	//locally. The local const can have a completely different type
	const shadowedConst = 45.67
	fmt.Printf("\tshadowedConst: %v (%[1]T)\n", shadowedConst)

	fmt.Printf("*** %v: END ***\n", title)
	fmt.Printf("\n")
}
