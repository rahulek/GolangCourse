package main

import "fmt"

// SwitchDemo function accepts a title parameter
// and demonstrates the use of the switch statement
func SwitchDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//switch blocks are more convenient when
	//multiple if.else if blocks are involved

	//Simple switch: with the tag
	role := 2
	switch role { //age is the tag
	case 1:
		fmt.Printf("\trole: %d is a regular user.\n", role)
	case 2:
		fmt.Printf("\trole: %d is an admin user.\n", role)
	case 3:
		fmt.Printf("\trole: %d is a super user.\n", role)
	}

	//Now the same with the default case
	role = 10
	switch role { //age is the tag
	case 1:
		fmt.Printf("\trole: %d is a regular user.\n", role)
	case 2:
		fmt.Printf("\trole: %d is an admin user.\n", role)
	case 3:
		fmt.Printf("\trole: %d is a super user.\n", role)
	default:
		fmt.Printf("\trole: %d is an unknown user type.\n", role)
	}

	//initialize syntax
	var roleFlagByte byte = 0xC
	//Bit 3: SuperUser, Bit 2: Admin, Bit 1: Regular User, Bit 0: No role
	switch role := roleFlagByte&0x8 != 0; role {
	case true:
		fmt.Printf("\troleFlagByte: %#X, Is a super user.\n", roleFlagByte)
	case false:
		fmt.Printf("\troleFlagByte: %#X, Is NOT a super user.\n", roleFlagByte)
	}

	//Tagless switch
	age := 22
	switch { //NOTE: no tag after the switch keyword
	case age >= 1 && age < 10:
		fmt.Printf("\tage: %d is a child.\n", age)
	case age >= 10 && age < 20:
		fmt.Printf("\tage: %d is a teenager.\n", age)
	case age >= 20 && age < 60:
		fmt.Printf("\tage: %d is an adult.\n", age)
	default:
		fmt.Printf("\tage: %d is a senior citizen.\n", age)
	}

	//Type switches
	//Useful when switching on the type of the tag
	//var price interface{} = 12.34 //interface{} reperesent any valid type
	var price any = 12.34 //any is available with Go 1.18 and above

	switch price.(type) {
	case int:
		fmt.Printf("\tprice: %v is an int type\n", price)
	case float32:
		fmt.Printf("\tprice: %v is a float32 type\n", price)
	case float64:
		fmt.Printf("\tprice: %v is a float64 type\n", price)
	default:
		fmt.Printf("\tprice: %v is an unknown type\n", price)
	}

	//Type can be used on custom types too
	myStruct := struct {
		price int
	}{
		10,
	}
	switch any(myStruct).(type) {
	case int:
		fmt.Printf("\tmystruct: %v is an int type\n", myStruct)
	default:
		fmt.Printf("\tmyStruct: %v is of type: %[1]T\n", myStruct)
	}

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}
