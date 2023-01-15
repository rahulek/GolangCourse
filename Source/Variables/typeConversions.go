package main

import (
	"fmt"
	"strconv"
)

func typeConversionDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v: START ***\n", title)

	var i int = 98
	var f float32 = 234.

	//Simple assignment won't compile
	//LHS is float and RHS is int
	//Uncomment to see the error
	//f = i

	//Conversion possible by an explicit cast only
	//NOTE - the casting uses a function call-like syntax
	f = float32(i)
	fmt.Printf("\tType Coversion (int to float): f= %.2f, i=%v\n", f, i)

	//Converting from float to int would result in data loss
	//but possible
	f = 123.45
	i = int(f)
	fmt.Printf("\tType Coversion (float to int): f= %.2f, i=%v\n", f, i)

	//how to print an interger as a string?
	//This will intepret year as a ascii/utf8 code byte
	//63 is ":" ASCII
	//But we need it to print "63"
	var year = 63
	fmt.Printf("\t(int as a string): %s\n", string(year))

	//strconv package is useful for such conversions
	//See: https://pkg.go.dev/strconv
	stringifiedYear := strconv.Itoa(year)
	fmt.Printf("\t(int as a string - with strconv): %s\n", stringifiedYear)

	fmt.Printf("*** %v: END ***\n", title)
	fmt.Printf("\n")
}
