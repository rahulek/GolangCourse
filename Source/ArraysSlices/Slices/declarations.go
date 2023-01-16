package slices

import "fmt"

//A Slice is a "view" into an underlying array.
//Slices are based on Arrays because the arrays
// provide underlying backing storage for the slices
//Consider, these declarations -
//   arr := [5]int { 1, 2, 3, 4, 5}
//   slice := arr[:3]   --> "slice" is a slice of the first 3 elements of "arr"
//A Slice is a Go data structure maintained by the Go runtime. Developers
// don't need to create it but it's understanding is crucial in dealing with
// the slices and how they behave.
//Slice's data structure -
//    ----------------
//   |  ptr to array -|----|            (arr)
//	 |				  |    |     |-------------------|
//   | ---------------|    |---->| 1 | 2 | 3 | 4 | 5 |
//   |     len(3)     |          |___________________|
//   | _______________|
//   |     cap(5)     |
//   ------------------

// DeclarationsDemo function accepts a title parameter
// and demonstrates the use of slice declarations
func DeclarationsDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//A Slice is declared with empty square brackets - [] unlike the arrays
	//A nil slice
	//A nil slice's data structure = { ptr: nil, len: 0, cap: 0}
	var nilSlice []int
	fmt.Printf("\tnilSlice: %v, len: %v, cap: %v\n", nilSlice, len(nilSlice), cap(nilSlice))
	fmt.Printf("\tBacking array address: %p\n", nilSlice)

	//An empty slice
	//A empty slice's data structure = { ptr: some_address, len: 0, cap: 0}
	emptySlice := []int{}
	fmt.Printf("\temptySlice: %v, len: %v, cap: %v\n", emptySlice, len(emptySlice), cap(emptySlice))
	fmt.Printf("\tBacking array address: %p\n", emptySlice)

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}
