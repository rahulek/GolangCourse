package main

import (
	"fmt"
	"sync"
)

// SimpleDeadlockDemo function accepts title parameter
// and demonstrates how goroutines can deadlock.
func SimpleDeadlockDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	wg := sync.WaitGroup{}
	ch := make(chan float32)

	wg.Add(2)
	go reader(&wg, ch)
	go writer(&wg, ch)
	wg.Wait()

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}

// The reader goroutine reads only 1 float from the channel
func reader(wg *sync.WaitGroup, ch <-chan float32) {
	fmt.Printf("\t\tSimpleDeadlockDemo receiver: Got: %v\n", <-ch)

	wg.Done()
}

// The writer goroutine writes two floats to the channel
//   - the first value written to the channel will be read by the reader
//   - the deadlock is created when the writer another float to the channel
//     but neither the reader NOR the main() goroutine have any mechanism
//     to read that value
//   - Because the writer gets blocked, it doesn't complete
//   - The waitGroup waiting for the reader and writer to complete keeps waiting
//     and the program does not proceed
//   - This is the DEADLOCK scenario
func writer(wg *sync.WaitGroup, ch chan<- float32) {
	defer wg.Done()

	fmt.Printf("\t\tSimpleDeadlockDemo writer: Writing: 3.1415927\n")
	ch <- 3.1415925

	//Write another value
	fmt.Printf("\t\tSimpleDeadlockDemo writer: Now Writing: 1.233409\n")
	ch <- 1.233409
}
