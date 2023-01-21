package main

import "fmt"

// ReturnValueDemo function accepts title parameter
// and demonstrates the use of return values from a function
func ReturnValueDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v: START ***\n", title)

	//Call a function that doesn't take any parameter
	//but returns a string
	ret := singleReturn()
	fmt.Printf("\tsingleReturn(): returned value: %v\n", ret)

	fmt.Println("~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~")

	ret = singleReturnParam("My Friend")
	fmt.Printf("\tsingleReturnParam(): returned value: %v\n", ret)

	fmt.Println("~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~")

	val := singleNamedReturn(40, 50)
	fmt.Printf("\tsingleNamedReturn: val: %v\n", val)

	fmt.Println("~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~")

	a := -10
	b := 50
	fmt.Println("\tTrying multipleReturnValues with valid pointer parameters.")
	sum, err := multiplReturnValues(&a, &b)
	if err != nil {
		fmt.Printf("\tmultipleReturnValues: Error (%v) occurred.\n", err)
	} else {
		fmt.Printf("\tmultipleReturnValues: Result: %v\n", sum)
	}

	fmt.Println("~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~")

	// Now try an error code
	fmt.Println("\tTrying multipleReturnValues with invalid pointer parameters.")
	sum, err = multiplReturnValues(&a, nil)
	if err != nil {
		fmt.Printf("\tmultipleReturnValues: Error (%v) occurred.\n", err)
	} else {
		fmt.Printf("\tmultipleReturnValues: Result: %v\n", sum)
	}
	fmt.Println("~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~")

	fmt.Printf("*** %v: END ***\n", title)
	fmt.Printf("\n")
}

func singleReturn() string {
	return "I welcome you!"
}

func singleReturnParam(name string) string {
	return fmt.Sprintf("I welcome you, %s", name)
}

func singleNamedReturn(a, b int) (result int) {
	//when the return value is named (for example: result, in this example),
	//function can use it as any local variable and set it.
	//Then, return statement need not be followed by any variable/value
	result = a * b

	return //NOTE: no  return result only, plain return
}

// NOTE - Go functions have ability to return multiple values.
// Its idiomatic to write a function that returns multiple values.
// and used mostly to return (value, noError) OR (noValue, Error)
// combinations. Other usages of multiple returns are also possible.
func multiplReturnValues(aPtr, bPtr *int) (int, error) {
	//If aPtr or bPtr is nil, return an error
	if aPtr == nil || bPtr == nil {
		//Error case, return (0, error)
		return 0, fmt.Errorf("Non-nil pointers are needed.")
	}

	//all OK, calculate sum
	sum := *aPtr + *bPtr
	return sum, nil //Return multiple values, sum and nil to indicate no error
}
