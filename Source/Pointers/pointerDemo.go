package main

import "fmt"

// PointersDemo function accepts title parameter
// and demonstrates use of Go pointers
func PointersDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	secretOfLife := 42
	//Creating a reference means setting a pointer
	//NOTE: & operator takes an address of another variable
	secretOfLifePtr := &secretOfLife
	fmt.Printf("\tsecretOfLife: %v, secretOfLife's address: %p\n", secretOfLife, &secretOfLife)
	//NOTE: %p format specifier is used to print the addresses
	fmt.Printf("\tsecretOfLifePtr: %p\n", secretOfLifePtr)

	//Read the data where the pointer is pointing at.
	//Because the ptr is pointing at secretOfLine, we expect 42 back
	//NOTE: A pointer is dereferenced using a * operator
	fmt.Printf("\tReading the ptr yields: %v\n", *secretOfLifePtr)

	//Similarly, a value can be changed using the pointer.
	//When that happens, the original data to which the pointer points is
	//effectively changed.
	*secretOfLifePtr = 424242
	fmt.Printf("\tData modified via the ptr -> original data: %v\n", secretOfLife)

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}

// ArrayWithPointers accepts title parameter
// and demonstrates how Go stores arrays and how
// to use pointer to point to its elements
func ArrayWithPointers(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//Create an array of 5 elements
	arr1 := [5]int{10, 20, 30, 40, 50}

	//Create 5 pointers, each pointing to a successive element in the array
	var ptr1 *int = &arr1[0]
	var ptr2 *int = &arr1[1]
	var ptr3 *int = &arr1[2]
	var ptr4 *int = &arr1[3]
	var ptr5 *int = &arr1[4]

	//Print all the address to which the 5 pointers are pointing
	//NOTE - because on this computer architecture, int is stored in
	//8 bytes, each address is 8-bytes apart.
	fmt.Printf("\tarr1 is: %v\n", arr1)
	fmt.Printf("\tarr1 is at: %p\n", &arr1)
	fmt.Printf("\t\tptr1 is pointing to: %p\n", ptr1)
	fmt.Printf("\t\tptr2 is pointing to: %p\n", ptr2)
	fmt.Printf("\t\tptr3 is pointing to: %p\n", ptr3)
	fmt.Printf("\t\tptr4 is pointing to: %p\n", ptr4)
	fmt.Printf("\t\tptr5 is pointing to: %p\n", ptr5)

	//Change the underlying array via each pointer (double each element)
	fmt.Printf("\tDoubling each array element through the pointers.\n")
	*ptr1 *= 2
	*ptr2 *= 2
	*ptr3 *= 2
	*ptr4 *= 2
	*ptr5 *= 2

	//Now print the original array
	fmt.Printf("\tMODIFIED array: %v\n", arr1)

	//Go does not allow pointer arithmetic as in C
	//Go's unsafe package (https://pkg.go.dev/unsafe) can be used
	//if this type of functionality is required.
	//Uncomment below to see the error
	//ptr1++

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}

// CustomTypePtrDemo function accepts title parameter
// and demonstrates the use of pointer to custom defined
// types such as structs and maps.
func CustomTypePtrDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//Declare a custom struct type
	type Book struct {
		title     string
		totalSold int
	}

	//Create a book
	aBook := Book{title: "My Book", totalSold: 1003}
	fmt.Printf("\tBook: %#v, Type: %[1]T\n", aBook)

	//Create a pointer/refer to the book
	var bookPtr *Book = &aBook
	fmt.Printf("\tBook via bookPtr: %#v\n", *bookPtr)

	//Modify the totalSold via pointer
	bookPtr.totalSold += 10
	fmt.Printf("\tIncremented totalSold by 10, new book: %#v\n", aBook)

	//make a map
	var myRatingsMap map[string]int = make(map[string]int)
	myRatingsMap["Amar Akbar Anthony"] = 2
	myRatingsMap["Godfather"] = 5
	myRatingsMap["The Kashmir Files"] = 5

	//create a pointer/ref to the map
	var mapPtr *map[string]int = &myRatingsMap
	fmt.Printf("\tMap via its pointer: %v\n", *mapPtr)

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}
