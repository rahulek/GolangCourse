package Stringsandrunes

import "fmt"

// DeclarationDemo function accepts a title parameter
// and demonstrates how to declare strings and runes
func DeclarationDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v: START ***\n", title)

	//string is a bunch of bytes (uint8) which are utf8 encoded
	//string literal is always declared within double quotes ("")
	//string is indexable with [ndx]
	s := "A sample string."
	fmt.Printf("\ts: %q (%[1]T)\n", s)
	fmt.Printf("\ts[2]: %v (%[1]T)\n", s[2])

	//rune is an alias for uint32
	//rune encodes a char in utf32 (all utf8 are part of utf32)
	//rune's size of 32-bits allow it to be used to store any character
	//  in any language or script.
	//rune is declared inside a single quote ('')

	r := '\U0001F549' //Unicode for AUM
	fmt.Printf("\tr: %c (%[1]T)\n", r)

	fmt.Printf("*** %v: END ***\n", title)
	fmt.Printf("\n")
}
