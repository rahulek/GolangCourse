package main

import "fmt"

// DeclareAndUseDemo function accepts title
// parameter and demonstrates how to declare
// and use functions
func DeclareAndUseDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//Invoke the sayHi() function.
	//No parameters, no return value
	sayHi()

	//Invoke function that needs 2 parameters and returns nothing
	greetUser("Sachin", "What a masterful batting!!")
	greetUser2("Rahul", "Hats off!!")

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}

// Function that accepts no parameter and returns nothing
// Declaration syntax:
//
//	func keyword - function_name - parentheses () - opening curly brace {
//
// NOTE - Open brace { must start on the same line as the function declaration
func sayHi() {
	fmt.Printf("\t\tsayHi() function: Hi\n")
}

// Function that accepts 2 parameters and returns nothing
// syntax same as above except
//
//	the parameters are declared within the parantheses ()
//	Each param is declared with the syntax: param_name - param_type
func greetUser(name string, message string) {
	fmt.Printf("\t\tgreetUser(name, message): Hi %v, %v\n", name, message)
}

// Function same as above except
// the params name and message are declared differently
// Both the params are of string type and therefore
// the type is declared only once after the parameter names
func greetUser2(name, message string) {
	fmt.Printf("\t\tgreetUser2(name, message): Hi %v, %v\n", name, message)
}
