package main

import (
	"fmt"
	"sync"
)

// ChannelCloseDetectionDemo function accepts title parameter
// and shows an alternate syntax to the range keyword. This
// is similar to detecting the closing of the channel but without
// the range keyword
func ChannelCloseDetectionDemo(title string) {
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

		//Instead of using a range keyword to range over
		//the channel, use an alternate syntax to read from the
		//channel and returns whether the close happened
		for {
			if i, ok := <-ch; ok {
				fmt.Printf("\t\tChannelCloseDetectionDemo: Got: %v\n", i)
			} else {
				fmt.Printf("\t\tChannelCloseDetectionDemo: **DETECTED CHANNEL CLOSE**\n")
				break
			}
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
