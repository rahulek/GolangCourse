package slices

import "fmt"

// CreationDemo function accepts a title parameter
// and demonstrates the slice creation operations.
func CreationDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//Define an underlying array
	arr := [5]int{1, 2, 3, 4, 5}

	//Slice creation syntax: arr[m:n]
	//This creates a slice of the underlying array
	//at index m (inclusive) to index n (exclusive)
	//That is: m to n-1
	//if m unspecified, its 0
	//if n unspecified, its len(arr)
	aSlice := arr[:3]  //slice of arr elements 0, 1, 2 (3 is excluded)
	bSlice := arr[4:]  //slice of arr elements 4 onwards
	cSlice := arr[2:5] //Slice of arr elements from index 2 to index 4
	dSlice := arr[:]   //Slice of arr that contains all the elements

	fmt.Printf("\tarr: %v\n", arr)
	fmt.Printf("\taSlice: %v\n", aSlice)
	fmt.Printf("\tbSlice: %v\n", bSlice)
	fmt.Printf("\tcSlice: %v\n", cSlice)
	fmt.Printf("\tdSlice: %v\n", dSlice)

	//Each slice is a different data structure maintained by
	//Go runtime BUT all of them has a pointer that's pointing
	//to various sections within the underlying array (arr, in this case)
	fmt.Printf("\tAddress of arr: %p, Length: %v\n", &arr, len(arr))
	fmt.Printf("\taSlice(%p) Backing Array(%p) Len(%d) Cap(%d)\n", &aSlice, aSlice, len(aSlice), cap(aSlice))
	fmt.Printf("\tbSlice(%p) Backing Array(%p) Len(%d) Cap(%d)\n", &bSlice, bSlice, len(bSlice), cap(bSlice))
	fmt.Printf("\tcSlice(%p) Backing Array(%p) Len(%d) Cap(%d)\n", &cSlice, cSlice, len(cSlice), cap(cSlice))
	fmt.Printf("\tdSlice(%p) Backing Array(%p) Len(%d) Cap(%d)\n", &dSlice, dSlice, len(dSlice), cap(dSlice))

	fmt.Printf("\t~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~\n")

	//Start with a nil slice
	var intSlice []int
	fmt.Printf("\tintSlice: %v (%[1]T)\n", intSlice)

	//append() function appends the slice with the new data
	//append() can append to the nil slice too
	intSlice = append(intSlice, 1)
	fmt.Printf("\tintSlice: %v\n", intSlice)

	//multiple elements can be appended
	intSlice = append(intSlice, 2, 3, 4, 5)
	fmt.Printf("\tintSlice: %v\n", intSlice)

	//New slice can be appended to the existing slice
	//But needs a spread-like syntax that uses ...
	intSlice = append(intSlice, []int{10, 20, 30}...)
	fmt.Printf("\tintSlice: %v\n", intSlice)

	fmt.Printf("\t~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~\n")

	//A Go make() function can be used to create the slice of
	//desired len and cap
	madeSlice := make([]int, 5) //only len
	madeSlice[0] = 10
	madeSlice[1] = 20
	madeSlice[2] = 30
	madeSlice[3] = 40
	madeSlice[4] = 50
	fmt.Printf("\tmadeSlice: %v\n", madeSlice)

	//append
	madeSlice = append(madeSlice, -10, -20, -30, -40, -50)
	fmt.Printf("\tmadeSlice: %v\n", madeSlice)

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")

}
