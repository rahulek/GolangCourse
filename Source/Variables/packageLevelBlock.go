package main

import "fmt"

// combine related declarations inside a package level block
// Note:
// 1. var is required only once
// 2. Each variable is declared as usual inside the () which defines a block
var (
	country         string = "India"
	indepenceYear   int    = 1947
	governingSystem string = "Democratic"
)

// Multiple blocks can exist at the package level
var (
	bookTitle  string = "Meditations"
	bookAuthor string = "Marcus Aurelius"
	bookAge    int    = 1900
)

func packageLevelBlockDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v: START ***\n", title)

	//Note: the format specifier changed from %v (determined by the compiler)
	//to %d(integers), %s(string), etc.
	//See: https://pkg.go.dev/fmt
	fmt.Printf("\tCountry: %s, Independece: %d, System: %s\n", country, indepenceYear, governingSystem)
	fmt.Println("\tBook Info: ", bookTitle, bookAuthor, bookAge)

	fmt.Printf("*** %v: END ***\n", title)
	fmt.Printf("\n")
}
