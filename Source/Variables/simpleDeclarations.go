package main

import "fmt"

func simpleDeclarationsDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v: START ***\n", title)

	//1. Declare a variable without initializing it
	//   Initialize it later
	var i int
	fmt.Printf("\t i: %v (%T)\n", i, i)

	//2. Declare a variable and initize it
	var weekDay int = 1
	fmt.Printf("\t weekDay: %v (%[1]T)\n", weekDay)

	//3. Shorthand Declaration, compiler infers the type
	today := "Sunday"
	salary := 20450.
	cmplx := 1 + 2i
	fmt.Printf("\t today: %v (%[1]T)\n", today)
	fmt.Printf("\t salary: %v (%[1]T)\n", salary)
	fmt.Printf("\t cmplx: %v (%[1]T)\n", cmplx)

	fmt.Printf("*** %v: END ***\n", title)
	fmt.Printf("\n")
}
