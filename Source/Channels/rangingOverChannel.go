package main

import (
	"fmt"
	"sync"
)

// ChannelRangingWithDeadlockDemo function accepts title parameter
// and demonstrates the use of range keyword with a channel but
// because the range does not complete ever, the reader deadlocks
func ChannelRangingWithDeadlockDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v: START ***\n", title)

	//Create a non-buffered channel
	ch := make(chan int)
	//Create a waitGroup
	wg := sync.WaitGroup{}

	//2 goroutines to complete
	wg.Add(2)

	//Anonymous function based reader goroutine
	go func(ch <-chan int) {
		//Use the range keyword to range over the channel
		//range will read the data coming over channel
		//until the channel is *closed*
		//If the channel does not close, range will remain
		//blocked.
		// **** THIS IS A DEADLOCK EXAMPLE *****
		for i := range ch {
			fmt.Printf("\t\tChannelRangingWithDeadlockDemo: Got: %v\n", i)
		}
	}(ch)

	//Anonymous function based writer goroutine
	go func(ch chan<- int) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}(ch)

	//Wait for all the goroutines to complete
	wg.Wait()

	fmt.Printf("*** %v: END ***\n", title)
	fmt.Printf("\n")
}

// ChannelRangingSolutionDemo function accepts title parameter
// and demonstrates the use of range keyword with a channel. It
// close the write side of the channel which is detected by the
// read end and that breaks the for range loop.
func ChannelRangingSolutionDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v: START ***\n", title)

	//Create a non-buffered channel
	ch := make(chan int)
	//Create a waitGroup
	wg := sync.WaitGroup{}

	//2 goroutines to complete
	wg.Add(2)

	//Anonymous function based reader goroutine
	go func(ch <-chan int) {
		defer wg.Done()

		//Use the range keyword to range over the channel
		//range will read the data coming over channel
		//until the channel is *closed*
		//Now, the writer goroutine uses close() to close
		//the channel. That is detected by the range and
		//it breaks the for loop and thus the deadlock
		for i := range ch {
			fmt.Printf("\t\tChannelRangingSolutionDemo: Got: %v\n", i)
		}
	}(ch)

	//Anonymous function based writer goroutine
	go func(ch chan<- int) {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			ch <- i
		}
		//All items written, now close the channel
		close(ch)
	}(ch)

	//Wait for all the goroutines to complete
	wg.Wait()

	fmt.Printf("*** %v: END ***\n", title)
	fmt.Printf("\n")
}
