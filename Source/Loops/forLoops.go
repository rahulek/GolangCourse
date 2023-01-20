package main

import "fmt"

// ForLoopDemo function accepts the title parameter
// and demonstrates the use of the for loops
func ForLoopDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//Simple traditional for loop
	for i := 0; i < 5; i++ {
		fmt.Printf("\tSimple for loop: i: %v\n", i)
	}

	fmt.Printf("\t~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~\n")

	//Multiple initializers
	//Note: i := 0, j:= 0 in the initializer list wont compile
	//Note: i++, j++ in the trailing increment/decrement list wont compile either
	for i, j := 0, 0; i < 5; i++ {
		fmt.Printf("\tMultiple initializers: i: %v, j: %v\n", i, j)
		//also deal with the j
		j--
	}

	fmt.Printf("\t~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~\n")

	//initializer and/or training increment/decrement block can be skipped
	counter := 0
	for counter < 10 {
		fmt.Printf("\tSkipped init and trail: counter: %v\n", counter)

		//counter must be modified so as to not create an infinite loop
		counter++
	}

	fmt.Printf("\t~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~\n")

	//infinite loops can be written. Care must be taken to break
	counter = 0
	for {
		counter++
		fmt.Printf("\tInfinite Loop: counter: %v\n", counter)
		if counter == 10 {
			break
		}
	}

	fmt.Printf("\t~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~\n")

	//continuous keyword allows loop logic to be skipped
	for counter := 0; counter < 10; counter++ {
		if counter%2 == 0 {
			continue
		}

		fmt.Printf("\tcontinue demo: counter: %v\n", counter)
	}

	fmt.Printf("\t~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~\n")

	//Nested for loops
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("\tNested loop: i = %v, j = %v\n", i, j)
		}
	}

	fmt.Printf("\t~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~\n")

	//Nested for loop with labelled break/continue
	//break's label allows the control flow to jump to a specified for location
OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if i == 0 && j == 0 {
				continue OuterLoop
			}
			fmt.Printf("\tNested loop: with labelled break i = %v, j = %v\n", i, j)
		}
	}

	fmt.Printf("\t~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~\n")

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}
