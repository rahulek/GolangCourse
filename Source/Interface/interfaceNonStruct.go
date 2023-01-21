package main

import (
	"fmt"
	"strings"
)

// InterfaceOnNonStructTypeDemo function accepts title parameter
// and demonstrates the use of interfaces on non-struct type.
func InterfaceOnNonStructTypeDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	var aString AString = AString("This is a test string")
	fmt.Printf("\tBefore Uppercasing: %v\n", aString)
	aString.ToUppercase()

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}

type UpperCaser interface {
	ToUppercase()
}

// A new type representing string MUST Be defined
// Even though its just a string, compiler does not
// allow std types to be used to implement interfaces
type AString string

func (s AString) ToUppercase() {
	fmt.Printf("\tUppercased: %v\n", strings.ToUpper(string(s)))
}
