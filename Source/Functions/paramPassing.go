package main

import "fmt"

// ParamPassingDemo function accepts title parameter
// and demonstrates pass-by-value of the parameters
func ParamPassingDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	a, b := 10, 20
	//variables are passed by-value
	//copies are made when a function is invoked
	//any changes made to them inside that function
	//WILL NOT BE VISIBLE in this function
	fmt.Printf("\tBefore function call : a: %v, b: %v\n", a, b)
	tryModifying(a, b)
	fmt.Printf("\tAfter function's return: a: %v, b: %v\n", a, b)

	fmt.Println("~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~")

	//The effect that the called function changes the data inside the
	//calling function can be achieved by passing the pointer to the function
	fmt.Println("\t--- Trying the modifications by passing the pointers ---")
	fmt.Printf("\tBefore function call : a: %v, b: %v\n", a, b)
	tryModifyingWithAPointer(&a, &b)
	fmt.Printf("\tAfter function's return: a: %v, b: %v\n", a, b)

	fmt.Println("~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~")

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}

func tryModifying(a, b int) {
	fmt.Printf("\t\ttryModifying(a, b): a: %v, b: %v\n", a, b)

	//a, b are local variables and are
	//assigned place on the stack frame that
	//is ephemeral - ie - lasts only until the
	//function is executing. Therefore, these
	//variables have no existence once the function returns

	//try modiying a and b
	a *= 2
	b *= 10

	fmt.Printf("\t\ttryModifying(a, b), after modification: a: %v, b: %v\n", a, b)
}

func tryModifyingWithAPointer(a, b *int) {
	fmt.Printf("\t\ttryModifying(a, b): a: %v, b: %v\n", *a, *b)

	//In this version, a are b are pointers to the a, b locations
	//inside the calling function's local frame. Modifications applied
	//by dereferencing the pointers will be reflected therefore in
	//the calling function

	//try modiying a and b
	*a *= 2
	*b *= 10

	fmt.Printf("\t\ttryModifying(a, b), after modification: a: %v, b: %v\n", *a, *b)
}
