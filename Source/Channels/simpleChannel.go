package main

import (
	"fmt"
	"sync"
)

// SimpleChannelDemo function accepts title parameter
// and demonstrates use of the simple channel to communicate
// between two goroutines
func SimpleChannelDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//Create a channel
	//Note: Channels are created ONLY with make() function
	//Data type that flows through the channel must be specified
	ch := make(chan string)
	wg := sync.WaitGroup{}

	//2 goroutines to complete
	wg.Add(2)

	//anonymous function based goroutine
	//that reads from the channel
	go func() {
		defer wg.Done()

		//Read from the chanel
		str := <-ch
		fmt.Printf("\t\tSimpleChannelDemo: reader received: %v\n", str)
	}()

	//anonymous function based goroutine
	//that writes to the channel
	go func() {
		defer wg.Done()
		ch <- "Hi, Welcome to Go Programming!!!"
	}()

	//Wait for the the waitgroup to be complete
	wg.Wait()

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}

// SimpleChannelWithDirectionDemo function accepts title parameter
// and demonstrates use of the simple channel to communicate
// between two goroutines. Channel declaration for the goroutines
// indicate the direction in which the data flows.
func SimpleChannelWithDirectionDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//Create a channel
	//Note: Channels are created ONLY with make() function
	//Data type that flows through the channel must be specified
	ch := make(chan string)
	wg := sync.WaitGroup{}

	//2 goroutines to complete
	wg.Add(2)

	//anonymous function based goroutine
	//that reads from the channel
	go func(ch <-chan string) { //DECLARE the channel direction (Read from the channel)
		defer wg.Done()

		//Read from the chanel
		str := <-ch
		fmt.Printf("\t\tSimpleChannelWithDirectionDemo: reader received: %v\n", str)
	}(ch)

	//anonymous function based goroutine
	//that writes to the channel
	go func(ch chan<- string) { //chan<- string : (Write into the channel)
		defer wg.Done()
		ch <- "Hi, Welcome to Go Programming!!!"
	}(ch)

	//Wait for the the waitgroup to be complete
	wg.Wait()

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}
