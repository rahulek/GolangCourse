package main

import "fmt"

// CollectionIterationDemo function accepts a title parameter
// and demonstrates the use of using for loops for iterating
// over various collection types
func CollectionsIterationDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//for loops can be used to iterate over
	//arrays, slices, strings, maps channels(not covered yet)

	//Arrays
	arr1 := [5]int{10, 20, 30, 40, 50}
	for i, v := range arr1 {
		fmt.Printf("\tRanging over an array: i: %v, v: %v\n", i, v)
	}

	fmt.Printf("\t~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ \n")

	//Slices
	slice1 := [5]byte{0x12, 0x18, 0x98, 0x7F, 0x56}
	for i, v := range slice1 {
		fmt.Printf("\tRanging over a slice: i: %v, v: %#X\n", i, v)
	}

	fmt.Printf("\t~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ \n")

	//Maps
	map1 := make(map[int]string)
	map1[0] = "Unused"
	map1[1] = "English"
	map1[2] = "Algebra"
	map1[3] = "History"
	map1[4] = "Social Sciences"

	for k, v := range map1 {
		fmt.Printf("\tRanging over a map: key: %v, value: %v\n", k, v)
	}

	//strings
	s := "Hello World!"
	for i, v := range s {
		fmt.Printf("\tRanging over a string: index: %v, value: %v\n", i, string(v))
	}

	fmt.Printf("\t~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ \n")

	//Discarding unused values with an underscore (_)
	s = "Hi"
	for _, v := range s { //NOTE: _ is used to ignore the index value
		fmt.Printf("\tRanging over a string: value: %v\n", string(v))
	}

	fmt.Printf("\t~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ \n")

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}
