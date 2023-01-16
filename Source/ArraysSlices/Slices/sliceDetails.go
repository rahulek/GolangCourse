package slices

import "fmt"

// SliceAsSharedViewDemo function accepts title parameter
// and demonstrates the sharing of the underlying memory
// among of the slices and how modifying the underlying
// backing array modifies each slice.
func SliceAsSharedViewDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//Backing array
	arr := [5]int{10, 20, 30, 40, 50}
	fmt.Printf("\tBacking array: %v\n", arr)

	//Create two slices which are overlapping
	//the underlying backing storage
	slice1 := arr[:3]  //Elements 0, 1 and 2
	slice2 := arr[1:4] //Elements 1, 2 and 3
	fmt.Printf("\tslice1: %v\n", slice1)
	fmt.Printf("\tslice2: %v\n", slice2)

	//Modify slice1's 2nd element (at its index 1)
	//Note that : slice1[1] == slice2[0] and slice1[2]= slice2[1]
	//Therefore, modifying the underlying array through slice1
	//reflects into slice2 too.
	slice1[1] *= 10
	slice1[2] *= 10
	fmt.Printf("\tslice1: %v\n", slice1)
	fmt.Printf("\tslice2: %v\n", slice2)

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}

// SliceLenCapDemo function accepts a title parameter
// and demonstrates how the slice's len and cap get
// changed
func SliceLenCapDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//Start with an empty int slice
	intSlice := []int{}
	fmt.Printf("\tEmpty intSlice: %v, len: %v, cap: %v\n", intSlice, len(intSlice), cap(intSlice))
	fmt.Printf("\t\tintSlice's backing ptr: %p\n", intSlice)

	//Append 1 element and check the len and cap
	intSlice = append(intSlice, 1)
	fmt.Printf("\tAfter inserting 1: intSlice: %v, len: %v, cap: %v\n", intSlice, len(intSlice), cap(intSlice))
	fmt.Printf("\t\tintSlice's backing ptr: %p\n", intSlice)

	//Append more elements
	intSlice = append(intSlice, 2, 3, 4, 5)
	fmt.Printf("\tAfter inserting 4 elements: intSlice: %v, len: %v, cap: %v\n", intSlice, len(intSlice), cap(intSlice))
	fmt.Printf("\t\tintSlice's backing ptr: %p\n", intSlice)

	//Add one element again
	intSlice = append(intSlice, 99)
	fmt.Printf("\tAfter inserting 1 element again: intSlice: %v, len: %v, cap: %v\n", intSlice, len(intSlice), cap(intSlice))
	fmt.Printf("\t\tintSlice's backing ptr: %p\n", intSlice)

	//NOTE:
	//In the above sequence:
	// - Note that we started with an empty slice: { ptr: singleton{}, len: 0, cap: 0 }
	// - After adding one element, the slice was out of space because it's cap was 0
	//   Go runtime will have to allocate a new memory space enough to accomodate it and copy
	//   the element to it. Now the slice: { ptr: <PTR1>, len: 1, cap: 1 }
	// - After adding 4 new elements, the slice was again out of space (its cap was 1)
	//   Go runtime again has to reallocate the array and move its contents.
	//   Now the slice: { ptr: <PTR2>, len: 5, cap: 6 }
	// - After the 2nd insertion of four elements, Go runtime allocates space for one extra element
	//   so that the next one element push should NOT cause reallocation:
	//   slice: { ptr: <PTR2>, len: 6, cap: 6 }
	//   Note that PTR2 has remained the same and no reallocation has occurred

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}

// AppendWeirdNessDemo function accepts a title parameter
// and demonstrates how append() function that has potention
// to cause reallocation of the underlying backing storage can
// keep the stale references to underlying backing storage and
// cause the program to use the memory that's not safe to use.
func AppendWeirdnessDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//Empty slice
	intSlice := []int{}
	fmt.Printf("\tEmpty intSlice: %v\n", intSlice)

	//Append 5 elements
	intSlice = append(intSlice, 1, 2, 3, 4, 5)
	fmt.Printf("\tModified intSlice: %v, len: %v, cap: %v\n", intSlice, len(intSlice), cap(intSlice))
	fmt.Printf("\t\tBacking storage for intSlice: %p\n", intSlice)

	//Another slice same as intSlice
	newSlice := intSlice
	fmt.Printf("\tintSlice: %v, newSlice: %v\n", intSlice, newSlice)
	fmt.Printf("\t\tBacking storage for intSlice: %p\n", intSlice)
	fmt.Printf("\t\tBacking storage for newSlice: %p\n", newSlice)

	//Cause append() to reallocate because its out of capacity
	intSlice = append(intSlice, 100, 200, 300)
	fmt.Printf("\tintSlice: %v, newSlice: %v\n", intSlice, newSlice)
	fmt.Printf("\t\tBacking storage for intSlice: %p\n", intSlice)
	fmt.Printf("\t\tBacking storage for newSlice: %p\n", newSlice)

	//NOTE:
	// - check the program output
	// - Its important to note that, after the last append of (100, 200, 300) to
	//   intSlice, the underlying storage was out of capacity (was 6 and 3 more appended).
	//   Therefore, Go runtime reallocates the backing storage and updates the intSlice.
	//   However, newSlice's backing storage pointers continues to point to the old address.
	//
	// - Go developers MUST remain vigilant. When working with multiple slices which may
	//   refer to the same underlying backing storage, the developers must ensure that all
	//   slice's backing storage pointer remains valid. Strange results/crash might occur if
	//   this is not handled correctly.

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}

// CopySliceDemo function accepts the title parameter
// and demonstrates the copy() function to copy the slices
func CopySliceDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//Copy slice from a src slice to a dest slice
	// - dest and src slices of same size
	// - dest slice size < src slice size
	// - dest slice size > src slice size
	src := []int{1, 2, 3, 4, 5}

	// dest slice same size as src file
	dest := make([]int, 5)
	copy(dest, src)
	fmt.Printf("\tsrc: %v, dest: %v\n", src, dest)
	//The slices are distinct. Their backing store pointers are different
	fmt.Printf("\t\tsrc: %p, dest: %p\n", src, dest)

	//dest slice size < src slice size
	//copy() will copy only 3 elements to src to match its size
	dest = make([]int, 3)
	copy(dest, src)
	fmt.Printf("\tsrc: %v, dest: %v\n", src, dest)
	//The slices are distinct. Their backing store pointers are different
	fmt.Printf("\t\tsrc: %p, dest: %p\n", src, dest)

	//dest slice size > src slice size
	//copy() will copy all src element to dest and extra will be 0s
	dest = make([]int, 10)
	copy(dest, src)
	fmt.Printf("\tsrc: %v, dest: %v\n", src, dest)
	//The slices are distinct. Their backing store pointers are different
	fmt.Printf("\t\tsrc: %p, dest: %p\n", src, dest)

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")

}
