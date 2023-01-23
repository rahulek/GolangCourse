package main

import (
	"fmt"
	"time"
)

// global data
var counter int = 0

// RaceConditionDemo function accepts a title
// parameter and demonstrates how the race
// conditions appear and cause problems in the Go program.
func RaceConditionDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//5 sets of readers and writers are started
	//reader simply read the counter (global) and
	//the write goroutine only increments the counter
	//
	//Because the gorotines don't access the counter
	//without any synchronization, reader may see the
	//counter value incrementd by more than one when it
	//executes next
	//
	//See the program output for the wrong results.
	for i := 0; i < 5; i++ {
		go reader()
		go writer()
	}

	time.Sleep(10 * time.Millisecond)

	fmt.Printf("*** %v : START ***\n", title)
	fmt.Printf("\n")
}

func reader() {
	fmt.Printf("\t\treader Goroutine: counter = %v\n", counter)
}

func writer() {
	counter++
}
