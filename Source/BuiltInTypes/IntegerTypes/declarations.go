package Integertypes

import (
	"fmt"
	"math"
)

// DeclarationDemo function receives the title as a parameter and
// demonstrates the signed and unsigned integer declarations and the
// type ranges associates with each.
func DeclarationDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	i := 123
	fmt.Printf("\t---------- SIGNED --------------\n")
	fmt.Printf("\tDefault integer type: i: %d (%[1]T)\n", i)
	fmt.Printf("\tint range: %d -> %d\n", math.MinInt, math.MaxInt)
	fmt.Printf("\tint8 range: %d -> %d\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("\tint16 range: %d -> %d\n", math.MinInt16, math.MaxInt16)
	fmt.Printf("\tint32 range: %d -> %d\n", math.MinInt32, math.MaxInt32)
	fmt.Printf("\tint64range: %d -> %d\n", math.MinInt64, math.MaxInt64)
	fmt.Printf("\n")

	fmt.Printf("\t---------- UNSIGNED --------------\n")
	fmt.Printf("\tuint range: %d -> %d\n", 0, uint(math.MaxUint))
	fmt.Printf("\tuint8 range: %d -> %d\n", 0, math.MaxUint8)
	fmt.Printf("\tuint16 range: %d -> %d\n", 0, math.MaxUint16)
	fmt.Printf("\tuint32 range: %d -> %d\n", 0, math.MaxUint32)
	fmt.Printf("\tuint64range: %d -> %d\n", 0, uint64(math.MaxUint64))
	fmt.Printf("\n")

	//Special byte type (alias for uint8)
	var b byte = 'a'
	fmt.Printf("\tSpecial type byte, b: %v (%[1]T)\n", b)

	//Individual integer types can be explicitly declared and used
	var u8 uint8 = 127 //the value must be within the valid uint8 range
	var i64 int32 = 1212222444
	fmt.Printf("\tu8 = %v (%[1]T), i64 = %v (%[2]T)\n", u8, i64)

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}
