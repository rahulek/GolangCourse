package Stringsandrunes

import "fmt"

// Function StdOperationsDemo accepts a title parameter
// and demonstrates operations of strings and runes
func StdOperationsDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v: START ***\n", title)

	//strings support concat and slicing and many other operations
	//see: https://pkg.go.dev/strings
	s1 := "Hello, "
	s2 := "world"
	fmt.Printf("\ts1: %q, s2: %q, s1+s2: %q\n", s1, s2, s1+s2)

	//Convert to a plain byte array: Take only 5 bytes of s1
	b1 := []byte(s1[:5])
	fmt.Printf("\tb1: %v (%[1]T)\n", b1)
	//Convert back to string
	fmt.Printf("\tback to string again: %q\n", string(b1))

	fmt.Printf("*** %v: END ***\n", title)
	fmt.Printf("\n")
}
