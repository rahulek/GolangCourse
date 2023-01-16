package arrays

import "fmt"

//Go arrays are collection types
//They store data of same type
//The size of the array must be specified within the [n]
//The size of the 2D/3D etc is specified as [m][n], etc
//Elements the array are stored adjacent to each other
//and accessed using the 0-based index. The fact that
//the data elements are stored one after the other in
//memory makes it fast to retrieve.

// DeclarationsDemo function accepts a title parameter
// and demonstrates the declarations of the array type.
func DeclarationsDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v: START ***\n", title)

	//Declare as an explicit type
	var taxRates [3]int = [3]int{10, 20, 30}
	fmt.Printf("\ttaxRates: %v (%[1]T)\n", taxRates)

	//Shorthand.
	//[...] indicates the size to be determined by the compiler
	workDays := [...]string{"M", "W", "F"}
	fmt.Printf("\tworkDays: %v (%[1]T)\n", workDays)

	//Accesses are 0-index based
	//Index 0 is the 1st element, etc.
	firstWorkDay := workDays[0]
	secondWorkDay := workDays[1]
	fmt.Printf("\tfirstWorkDay: %v, secondWorkDay: %v\n", firstWorkDay, secondWorkDay)

	//Array elements can be modified
	taxRates[0] *= 2
	taxRates[1] *= 2
	taxRates[2] *= 2
	fmt.Printf("\tDouble taxRates: %v\n", taxRates)

	fmt.Printf("*** %v: END ***\n", title)
	fmt.Printf("\n")
}
