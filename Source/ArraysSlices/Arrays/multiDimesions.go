package arrays

import "fmt"

// MultiDimensionsDemo function gets title parameter
// and demonstrates declaration of the multi-dimensional array
func MultiDimensionsDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//Declare a 2x2 Identity Matrix
	//and initialize later
	var idMatrix [2][2]int
	//The array is not initialized and therefore all 0
	fmt.Printf("\tidMaxtrix: %v\n", idMatrix)

	//Now, initialize the array and print
	idMatrix[0] = [2]int{1, 0}
	idMatrix[1] = [2]int{0, 1}
	fmt.Printf("\tidMaxtrix (after init): %v\n", idMatrix)

	//What's the size of the idMatrix
	//Its calculated using the len() method
	fmt.Printf("\tsize of idMatrix: [%vx%v]\n", len(idMatrix), len(idMatrix[0]))

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}
