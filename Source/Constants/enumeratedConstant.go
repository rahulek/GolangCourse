package main

import "fmt"

// In this const block, a is set to an expression resulting in 1
// All subsequent consts use new value of iota and expression used
// to create the new value
const (
	a = iota + 1
	b
	c
	d
)

// Multiple const blocks can exist
// Here, iota is used and its value (0) is unassigned (_) and ignored
// All subsequent constants get incremental value of 1, 2, 3....
const (
	_ = iota
	saturday
	sunday
	monday
	tuesday
	wednesday
	thursday
	friday
)

// EnumeratedConstDemo function takes a title parameter
// and demonstrates the use of enumeration based constants.
func EnumeratedConstDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v: START ***\n", title)

	const localEnumConst = iota
	fmt.Printf("\tSimple iota based constant: %v (%[1]T)\n", localEnumConst)

	const (
		zero = iota
		one
		two
		three
	)
	fmt.Printf("\tzero = %v, one = %v, two = %v, threee = %v\n", zero, one, two, three)

	//package level constants are most useful
	fmt.Printf("\ta = %v, b = %v, c = %v, d = %v\n", a, b, c, d)
	fmt.Printf("\tSunday = %v, Monday = %v, Friday = %v\n", sunday, monday, friday)

	fmt.Printf("*** %v: END ***\n", title)
	fmt.Printf("\n")

}
