package main

import "fmt"

func main() {
	fmt.Println("-------------  Pointers: START -----------")

	PointersDemo("Simple Pointers")
	ArrayWithPointers("Pointers to array elements")
	CustomTypePtrDemo("Pointers to custom types - structs/maps")

	fmt.Println("-------------  Pointers: END -----------")
}
