package main

import "fmt"

var i int = 42
var salary = 456.50
var taxRate float32 = 0.33

// Error Scenario 1
//Note that the shorthand declarations at the package level not allowed.
//Uncomment the line below to see the compiler error
//weekDay := "Monday"

// Error Scenario 2
// Type mismatches are flagged as compiler errors
//Uncomment the line below to see the compiler error
// LHS is integer type but RHS is the string type
//var counter int = "one"

func packageLevelDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v: START ***\n", title)

	fmt.Printf("\ti: %v\n", i)
	fmt.Printf("\tsalary: %v\n", salary)
	fmt.Printf("\ttaxRate: %v\n", taxRate)

	fmt.Printf("*** %v: END ***\n", title)
	fmt.Printf("\n")
}
