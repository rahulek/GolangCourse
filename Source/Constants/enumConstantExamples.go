package main

import "fmt"

// Define a const block for various memory sizes
// (Used from: https://go.dev/doc/effective_go#constants)
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)

// Define another const block that uses
// the bitshift operations that are used
// to define a bit-wide role
const (
	NoRole          = iota
	RegularUserRole = 1 << iota
	AdminUserRole
	SuperUserRole
)

// EnumeratedConstExamplesDemo function accepts a title string
// and demonstrates the usage of enumerated constant in a practical
// application settings
func EnumeratedConstExamplesDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	fmt.Printf("\tKB= %v, MB= %v, GB= %v, TB= %v\n", KB, MB, GB, TB)
	fmt.Printf("\tNoRole= %#08b\n", NoRole)
	fmt.Printf("\tRegularUserRole= %#08b\n", RegularUserRole)
	fmt.Printf("\tAdminUserRole= %#08b\n", AdminUserRole)
	fmt.Printf("\tSuperUseRole= %#08b\n", SuperUserRole)

	//Current Role is only a RegularUser
	var roleByte byte = RegularUserRole
	fmt.Printf("\troleByte : %#08b, isAdminUser? : %v, isRegularUser? : %v\n",
		roleByte, roleByte&AdminUserRole != 0, roleByte&RegularUserRole != 0)

	//Use now also becomes and AdminUse
	roleByte |= AdminUserRole
	fmt.Printf("\troleByte : %#08b, isAdminUser? : %v, isRegularUser? : %v\n",
		roleByte, roleByte&AdminUserRole != 0, roleByte&RegularUserRole != 0)

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}
