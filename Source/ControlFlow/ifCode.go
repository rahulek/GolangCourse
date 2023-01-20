package main

import "fmt"

// IfDemo function accepts a title parameter
// and demonstrates the use of if() control flow
// statement
func IfDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : BEGIN ***\n", title)

	//General syntax
	// if condition {
	// code to execute when the condition is evaluated to true
	//}
	// NOTE:
	//  1. Condition is NOT surrounded by the parantheses
	//  2. The opening curly brace { must be on the same line as the if statement
	if true {
		fmt.Printf("\tcondition true: will always get printed\n")
	}

	if false {
		fmt.Printf("\tcondition false: this will NOT get printed")
	}

	// Code that shows how to use comparion operator(s) to generate
	// a boolean result that can be used inside the if condition
	// 6 Comparison operators:
	// <, <=, >, >=, =, !=
	i := -100
	if i < 0 {
		fmt.Printf("\ti: %v, i < 0 ?: %v\n", i, i < 0)
	}

	// if block with the initiazer code
	// Syntax:
	//  if someVar := someExpr; condition_with_someVar {
	// someVar in the scope and can be used
	//}

	romanEmperorsMap := map[int]string{
		1: "Julius Caesar",
		2: "Augustus",
		3: "Tiberius",
		4: "Caligula",
		5: "Claudius",
	}
	if name, ok := romanEmperorsMap[1]; ok {
		fmt.Printf("\tif block with initializer syntax: 1st Roman emperor name: %v\n", name)
	}
	//Condition below uses a negation operator !
	if _, ok := romanEmperorsMap[10]; !ok {
		fmt.Printf("\tif block with initializer syntax: 10th Roman emperor name NOT found in the map.\n")
	}

	//if blocks are most useful with the logical operators
	//&& - Logical AND (both conditions suurrounding the && must be true)
	//|| - Logical OR (either the left OR the right condition must be true)
	//!  - Logical Negation( true -> false, false -> true)
	gradeSecured := 78
	//&& - if both conditions evaluated to be true, the if block is executed
	if gradeSecured >= 40 && gradeSecured <= 100 {
		fmt.Printf("\tgradeSecured: %v (>=40 and <= 100)\n", gradeSecured)
	}

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}
