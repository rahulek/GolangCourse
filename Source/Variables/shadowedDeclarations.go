package main

import "fmt"

var (
	//Error scenario 1:
	//package level var can't be redeclared
	//i below is already declared inside packageLevel.go in the same "main" package
	//i int = 9669

	j int = 9669
)

func shadowedDeclarationsDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v: START ***\n", title)

	//j below is in the block scope (of the surrounding function)
	//and shadows the package level j which is initialized to 9669
	var j int = 768
	fmt.Printf("\tj: %d\n", j)

	fmt.Printf("*** %v: END ***\n", title)
	fmt.Printf("\n")
}
