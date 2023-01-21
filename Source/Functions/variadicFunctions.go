package main

import "fmt"

// VariadicDemo function accepts title parameter
// and demonstrates the use of variable number
// of function arguments
func VariadicDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v: START ***\n", title)

	simpleVariadic("First", "Middle", "Last")

	fmt.Println("~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~")

	simpleVariadic("Vegetables", "Fruits", "Eggs", "Bakery", "Oils")

	fmt.Println("~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~")

	teamAverage("LA Lakers", 110, 240, 50, 100, 85, 90)

	fmt.Println("~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~")

	fmt.Printf("*** %v: START ***\n", title)
	fmt.Printf("\n")
}

func simpleVariadic(names ...string) {
	fmt.Printf("\t\tsimpleVariadic()\n")
	for i, name := range names {
		fmt.Printf("\t\t(%d) %v\n", i, name)
	}
}

func teamAverage(teamName string, scores ...int) {
	numOfScores := len(scores)
	fmt.Printf("\t\tteamAverage(): %v team scores -> %v\n", numOfScores, scores)

	sum := 0
	for _, score := range scores {
		sum += score
	}
	fmt.Printf("\t\tteamAverage(): Total: %v, Scores: %v, Average: %v\n", sum, numOfScores, float64(sum)/float64(numOfScores))
}
