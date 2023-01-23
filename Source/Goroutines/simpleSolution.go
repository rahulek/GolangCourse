package main

import (
	"fmt"
	"time"
)

// SimpleSolutionDemo function accepts a title parameter
// and demonstrates how a temporary solution solves the
// problem exhibited by the simpleGoroutine.go example.
func SimpleSolutionDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//Start a goroutine
	go helloWorld2()

	//Just so that the function doesn't return to main and finish
	//sleep for some time so that the helloWorld2 goroutine gets a
	//chance to execute.
	//
	//This, of course, is a not a scalable solution and putting an
	//artificial sleep is of no practial use to synchronize the
	//execution of the goroutines
	time.Sleep(3 * time.Millisecond)

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}

func helloWorld2() {
	fmt.Printf("\t\thelloWorld2 goroutine: HELLO WORLD\n")
}
