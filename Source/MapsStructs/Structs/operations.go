package structs

import "fmt"

// OperationsDemo function accepts title parameter
// and demonstrates use of the access, assignment
// operations.
func OperationsDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	type movieRating struct {
		movieTitle string
		rating     int
	}

	//Create a struct to store a movie rating
	rating1 := movieRating{movieTitle: "Godfather - 1", rating: 5}

	//Access the struct fields using a dot(.) notation
	fmt.Printf("\tMovie rating: Title: %v, Rating: %v\n", rating1.movieTitle, rating1.rating)

	//Structs assignments involve copy (like arrays) and unlike (maps)
	copiedRating := rating1
	fmt.Printf("\tRating: %#v, Copied Rating: %#v\n", rating1, copiedRating)

	//Modifyind the copy DOES NOT modify the original
	//They are independent copies
	copiedRating.rating *= 2
	fmt.Printf("\tModified the copiedRating's rating field\n")
	fmt.Printf("\tRating: %#v, Copied Rating: %#v\n", rating1, copiedRating)

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}
