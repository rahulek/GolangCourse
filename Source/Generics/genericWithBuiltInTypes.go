package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// GenericsWithBuiltInTypesDemo function accepts title
// parameter and demonstrates the use of Go Generics
// with maps/structs, etc.
func GenericsWithBuiltInTypesDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	var capitalsMap map[string]string = map[string]string{
		"MH": "Mumbai",
		"MP": "Bhopal",
		"GJ": "Gandhi Nagar",
		"TN": "Chennai",
		"KL": "Thiruananthpuram",
		"KA": "Bengaluru",
	}
	var literacyPercentMap map[string]int = map[string]int{
		"MH": 76,
		"MP": 45,
		"KL": 100,
		"KA": 72,
		"TN": 77,
	}

	//Deal with a map[string]string
	genericMap[string, string](capitalsMap)

	fmt.Println("\t~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~")

	//Now use the same function to deal with map[string]int
	genericMap[string, int](literacyPercentMap)

	fmt.Println("\t~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~")

	//Create an IntCounter of int8
	int8Counter := Counter[int8]{counter: 0}
	fmt.Printf("\tint8Counter: initial value: %#v\n", int8Counter)
	int8Counter.incrementCounterBy(int8(126))
	fmt.Printf("\tint8Counter: incremented by 126: %#v\n", int8Counter)
	int8Counter.incrementCounterBy(int8(1))
	fmt.Printf("\tint8Counter: incremented by 1: %#v\n", int8Counter)
	int8Counter.incrementCounterBy(int8(1))
	fmt.Printf("\tint8Counter: incremented by 1: %#v\n", int8Counter)

	fmt.Println("\t~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~")

	//Create an IntCounter of int8
	//Instead of 0, start counting from 10
	uint8Counter := Counter[uint8]{counter: 10}
	fmt.Printf("\tuint8Counter: initial value: %v\n", uint8Counter)
	uint8Counter.incrementCounterBy(uint8(243))
	fmt.Printf("\tuint8Counter: incremented by 243: %v\n", uint8Counter)
	uint8Counter.incrementCounterBy(uint8(2))
	fmt.Printf("\tuint8Counter: incremented by 2: %v\n", uint8Counter)
	uint8Counter.incrementCounterBy(uint8(1))
	fmt.Printf("\tuint8Counter: incremented by 1: %v\n", uint8Counter)

	//Invalid type of counter can't be instantiated.
	//For example - we constrained the counter's type to all integer variants.
	//Therefore, float counter should not get created/
	//Uncomment the following 2 lines to see the compiler error.
	// floatCounter := Counter[float64]{counter: 10.34}
	// fmt.Printf("\tfloatCounter: initialValue: %v\n", floatCounter)

	fmt.Println("\t~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~")

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")

}

// genericMap function has 2 type parameters -
//
//	K - used for defining the type of the map's keys
//	V - used for defining the type of the map's values
//
// # Both K and V are constrained to be constraits.Order
//
// Constrains are go interfaces. For example - order is
// declared in the constraints package. See - https://pkg.go.dev/golang.org/x/exp/constraints
func genericMap[K comparable, V constraints.Ordered](inputMap map[K]V) {
	fmt.Printf("\t\tgenericMap():\n")

	for k, v := range inputMap {
		fmt.Printf("\t\t (%v): %v\n", k, v)
	}
}

// IntCounter is just an int type
// constraints.Integer is defined as an interface -
//
//	Signed | Unsigned
//
// Signed Is - ~int | ~int8 | ~int16 | ~int32 | ~int64
// Unsigned is - ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
type Counter[T constraints.Integer] struct {
	counter T
}

func (ic *Counter[T]) incrementCounterBy(by T) T {
	ic.counter += by
	return ic.counter
}
