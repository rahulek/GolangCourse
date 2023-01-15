package main

import "fmt"

// Package scope vars
// Behave differently than the local scope declarations
// - Local scope unused vars are flagged as errors
// - Package scope unused vars are flagged by the linter and not treated as errors
var (
	packageLevelUnused = 34
)

func declarationErrorsDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v START ***\n", title)

	//Error Scenario:
	//Variable declared in local scope MUST be used
	//i below is declated but not used
	//Uncomment to see the error
	//var i int

	//j is declared and used, no problem
	var used string = "used"
	fmt.Printf("\tUsed: %v\n", used)

	fmt.Printf("*** %v END ***\n", title)
	fmt.Printf("\n")
}
