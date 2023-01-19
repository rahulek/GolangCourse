package maps

import "fmt"

// OperationsDemo function accepts a title parameter
// and demonstrates various operations of the maps.
func OperationsDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	capitalsOfIndianStates := make(map[string]int)
	capitalsOfIndianStates["Uttar Pradesh"] = 199812341
	capitalsOfIndianStates["Maharashtra"] = 112374333
	capitalsOfIndianStates["Bihar"] = 104099452
	capitalsOfIndianStates["West Bengal"] = 91276115
	capitalsOfIndianStates["Madhya Pradesh"] = 72626809

	//Operation: Getting the value of a given key
	key := "Maharashtra"
	value := capitalsOfIndianStates[key]
	fmt.Printf("\tKey: %v, Value: %v\n", key, value)

	//Operations: Getting the value of the key that does not exist
	//in the map
	key = "Mahrashtra" //An intention typo in the spelling of the key
	value = capitalsOfIndianStates[key]
	//This returns 0
	// --> This is confusing because
	// --> Does it mean that the key does not exist?
	// --> OR the population is really 0?
	fmt.Printf("\tKey: %v, Value: %v\n", key, value)

	//To fix the confusion in the above scenario,
	//use the double return value syntax
	value, ok := capitalsOfIndianStates[key] //key is still a mistyped "Maharashtra"
	if ok {
		fmt.Printf("\tKey: %v, Value: %v\n", key, value)
	} else {
		fmt.Printf("\tKey: %v: VALUE NOT AVAILABLE\n", key)
	}

	//Operation: Add a new value to the map
	capitalsOfIndianStates["Tamil Nadu"] = 72147030
	fmt.Printf("\tTamil Nadu's population: %v\n", capitalsOfIndianStates["Tamil Nadu"])

	//Opeartion: Modifying an existing key
	capitalsOfIndianStates["Tamil Nadu"] = 81245678
	fmt.Printf("\tTamil Nadu's NEW population: %v\n", capitalsOfIndianStates["Tamil Nadu"])

	//Operations: Deleting a key
	//Use a delete function
	delete(capitalsOfIndianStates, "Tamil Nadu")
	fmt.Printf("\tDELETED: Tamil Nadu's population: %v\n", capitalsOfIndianStates["Tamil Nadu"])

	fmt.Println("~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ")

	//Operation: Using a looping construct (covered later)
	for key, value := range capitalsOfIndianStates {
		fmt.Printf("\tIteration: key= %v, value= %v\n", key, value)
	}

	fmt.Println("~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ")

	//Operation: Size of the map using the len() function
	mapSize := len(capitalsOfIndianStates)
	fmt.Printf("\tSize of the Capitals Map: %v\n", mapSize)

	//Operation: Assignments/Copy - Unlike arrays, maps are passed by reference
	//assign map to another map
	copiedMap := capitalsOfIndianStates
	fmt.Printf("\tMaharashtra in both the maps: %v, %v\n",
		capitalsOfIndianStates["Maharashtra"],
		copiedMap["Maharashtra"])
	//Now delete "Maharashtra" from copiedMap
	delete(copiedMap, "Maharashtra")
	//Check the deleted key again in both the maps
	fmt.Printf("\tMaharashtra(DELETED) in both the maps: %v, %v\n",
		capitalsOfIndianStates["Maharashtra"],
		copiedMap["Maharashtra"])

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}
