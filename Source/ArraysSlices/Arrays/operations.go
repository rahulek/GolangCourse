package arrays

import "fmt"

//IMPORTANT:
//Arrays are passed to functions as value types in Go
//This has an adverse performance implication
//If the array is large, its assignment to other array
//involves a copy operation and that may affect the program's
//performance adversely.

// OperationsDemo function accepts a title parameter
// and demonstrates the use of len() and effect of copying.
func OperationsDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : BEGIN ***\n", title)

	//The len() function allows to get the size of the array
	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("\tarr1: %v, size: %v\n", arr1, len(arr1))

	//Assignement - involves copy of arr1 into arr2
	arr2 := arr1
	fmt.Printf("\tarr2: %v, size: %v\n", arr2, len(arr2))

	//Change only arr1 and print arr1 and arr2
	//Note that arr1's modification is NOT reflected into arr2
	//arr1 and arr2 have independent existences and are placed
	//in memory at different addresses
	arr1[0] = 10
	fmt.Printf("\tarr1: %v, arr2: %v\n", arr1, arr2)

	//Confirm the existence of arr1 and arr2 at separate
	//memory addresses.
	//%p allows us to print the address using a pointer format
	//& is a reference/addressof() operator
	fmt.Printf("\tarr1: %p, arr2: %p\n", &arr1, &arr2)

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}
