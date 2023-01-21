package main

import "fmt"

// FuncAsTypeDemo function accepts title parameter
// and demonstrates the use of functions as a type
func FuncAsTypeDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//Create an anonynous function
	//because it has no name, it must be invoked immediately to be
	//of any use.
	func(a, b int) {
		fmt.Printf("\t\tInside an anonymous function: a: %v, b: %v\n", a, b)
	}(10, 20)

	//Store a func in a variable
	//multiplyFn is a func that takes 2 integers parameters and return an integer
	var multiplyFn func(int, int) int = func(a, b int) int {
		return a * b
	}
	//Print the type of multiplyFn
	fmt.Printf("\tmultiplyFn: %T\n", multiplyFn)

	//now that the multiplyFn is a function type, the function it represents
	//can be invoked
	result := multiplyFn(10, 20)
	fmt.Printf("\tmultiplyFn(10, 20) returned: %v\n", result)

	//multiplyFn variable can be set to nil (which means the function is not defined)
	//obviously, calling multiplyFn now will cause a panic.
	multiplyFn = nil

	fmt.Println("~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ")

	//Functions can return other functions
	//create a base 5 adder
	base5Adder := createAdder(5)
	fmt.Printf("\tbase5Adder's type: %T\n", base5Adder)
	//Now, we can call the base5Adder function. Whatever
	//value we pass in gets added to the base value of 5
	//becuase it has closed over the baseValue of 5 in its
	//context
	result = base5Adder(25)
	fmt.Printf("\tbase5Adder(25): result: %v\n", result)
	result = base5Adder(-5)
	fmt.Printf("\tbase5Adder(-5): result: %v\n", result)

	fmt.Println("~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ")

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}

// Adder is a function that accepts some base
// as a parameter. It return a function type.
// This returned function, when invoked will take
// one integer parameter and return an integer value.
//
// The returned function needs baseValue defined in
// the context of the adder function and its implementation
// closes over baseValue. Therefore, its called as a closure
func createAdder(baseValue int) func(int) int {

	//return a function type
	//the returned function, when invoked
	//will take an int parameter that it will
	//add to the baseValue
	return func(v int) int {
		return baseValue + v
	}
}
