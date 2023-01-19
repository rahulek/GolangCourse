package maps

import "fmt"

// DeclarationsDemo function accepts title parameter
// and demonstrates the declarations of the maps.
func DeclarationsDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//Declare a simple map
	//Map's keys are of string type
	//Map's values are of of string type
	//Syntax: mapVar := map[k]v
	governingSystems := map[string]string{
		"India":                    "Democracy",
		"United States of America": "Democracy",
		"Libya":                    "Dictator",
		"China":                    "Democracy with Communism",
		"United Kigdom":            "Constitutional Monarchy", //Ending comma is EXPECTED
	}

	fmt.Printf("\tGoverningSystems: %v\n", governingSystems)
	fmt.Printf("\tType of GoverningSystems: %T\n", governingSystems)

	//Create a map using the make() function
	//type(K) = string, type(V) = int
	capitalsOfIndianStates := make(map[string]int)
	capitalsOfIndianStates["Uttar Pradesh"] = 199812341
	capitalsOfIndianStates["Maharashtra"] = 112374333
	capitalsOfIndianStates["Bihar"] = 104099452
	capitalsOfIndianStates["West Bengal"] = 91276115
	capitalsOfIndianStates["Madhya Pradesh"] = 72626809

	fmt.Printf("\tcapitals: %v\n", capitalsOfIndianStates)
	fmt.Printf("\tcapitals: %T\n", capitalsOfIndianStates)

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}
