package main

import "fmt"

// IfElseDemo function accepts a title parameter
// and demonstrates the use of if..else blocks
func IfElseDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	age := 72
	if age >= 1 && age < 10 {
		fmt.Printf("\tage: %d is a child\n", age)
	} else if age >= 10 && age < 19 {
		fmt.Printf("\tage: %d is a teenager\n", age)
	} else if age >= 19 && age < 60 {
		fmt.Printf("\tage: %d is an adult\n", age)
	} else if age >= 60 && age <= 100 {
		fmt.Printf("\tage: %d is a senior citizen\n", age)
	}

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}
