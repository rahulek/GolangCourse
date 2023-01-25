package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("-------------- Quit Signal Pattern : START ----------")
	fmt.Println()

	//Create a Quit channel to signal the producer gorotine
	//to quit
	quitCh := make(chan struct{})
	//Timeout of 2s
	timeOutCh := time.After(2000 * time.Millisecond)

	//Create a talkative goroutine
	talker := talkativeProducerThatQuits("CHIRP", quitCh)

outer:
	//keep polling with select
	//If talker's data is available, grab it and print it
	//If timeout occurs, send a quit signal on quit channel
	for {
		select {
		case data := <-talker:
			fmt.Printf("\tData Received: %s\n", data)
		case <-timeOutCh:
			fmt.Printf("\t***TIMED OUT: Sending the Quit Signal **\n")
			quitCh <- struct{}{}
			break outer
		}
	}

	fmt.Println()
	fmt.Println("-------------- Quit Signal Pattern : END  ----------")
}

func talkativeProducerThatQuits(input string, quitCh chan struct{}) <-chan string {
	//make a channel to talk on.
	ch := make(chan string)

	//Start a goroutine that will keep
	//write a string to the channel every 100ms
	//If there is a quit signal received on that
	//channel, quit!
	go func(inp string, ch chan<- string, qCh <-chan struct{}) {
	outer:
		for i := 0; ; i++ {
			select {
			case ch <- fmt.Sprintf("%s : %d", inp, i+1):
				//do nothing
			case <-qCh:
				fmt.Printf("\ttalkativeProducer: ** RECEIVED QUIT SIGNAL **\n")
				break outer
			}
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Printf("\ttalkingProducer: ** goroutine finished ***\n")
	}(input, ch, quitCh)

	//return the talking channel
	return ch
}
