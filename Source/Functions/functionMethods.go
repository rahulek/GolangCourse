package main

import "fmt"

// FuncAsTypeMethodsDemo function accepts title parameter
// and demonstrates the use of function as type methods
func FuncAsTypeMethodsDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//Create a User type and invoke a printInfo() method on it.
	aUser := User{Name: "James Bond", Age: 100}
	aUser.printInfo()

	//Invoke another method that accepts a pointer receiver
	(&aUser).incrementAge()
	//Print the user info to confirm if the Age was really incremented
	aUser.printInfo()

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}

// define a new type
type User struct {
	Name string
	Age  int
}

// printInfo is now not an ordinary function
// Its called as a method that accepts a
// value receiver u. Threfore, it can also
// be seen as a method that can invoked on a User
// type.
// The value receiver means that the User variable
// is passed by-value
func (u User) printInfo() {
	//Inside the method body, it has access to the User type
	//and all its underlying data/fields (in case of structs)
	fmt.Printf("\t\tMethod printInfo() : Name: %v, Age: %v\n", u.Name, u.Age)
}

// Methods can also receive pointer receivers
// This allows the method to make modifications to the User object
func (uPtr *User) incrementAge() {
	uPtr.Age++
	fmt.Printf("\t\tMethod incrementAge() : Name: %v, Age: %v\n", uPtr.Name, uPtr.Age)
}
