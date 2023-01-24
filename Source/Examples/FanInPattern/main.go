package main

import "fmt"

func main() {
	fmt.Println("------------- FanIn Pattern : START --------------")
	fmt.Println()

	//Create a hiChannel that will pubslihed Hi messages
	hiCh := generator("Hi")
	//Create a byeChannel that will pubslihed Hi messages
	byeCh := generator("Bye")
	//Create a muxed channel which handles hiCh and byeCh
	muxedCh := fanInImpl(hiCh, byeCh)

	//Read first 10 messages coming from either hiCh/byCh
	//via the muxedCh
	for i := 0; i < 10; i++ {
		fmt.Printf("\t\tMUX CHANNEL: Recd: %v\n", <-muxedCh)
	}

	fmt.Println()
	fmt.Println("------------- FanIn Pattern : END --------------")
}

// fanInImpl will take any number of string producing channels
// it will create a muxed channels that will get any of the
// strings produced by any of the channels passed in.
func fanInImpl(channels ...<-chan string) <-chan string {
	//Create a muxed channel
	fanInCh := make(chan string)

	//Range over all the channels provides
	for _, ch := range channels {
		//for each input channel, launch a goroutine to handle that channel only
		//the goroutine will wait for any input coming from the provided channel
		//Once any input arrives, it simply writes it to the muxed channel
		go func(ch <-chan string) {
			for {
				fanInCh <- <-ch
			}
		}(ch)
	}

	//return the muxed channel
	return fanInCh
}

// generator function implements the Generator pattern
// that creates and returns a channel which keeps
// producing data and then closes that channel
func generator(input string) <-chan string {
	//Create a generator channel
	ch := make(chan string)

	//Start a goroutine that will produce
	//a string and write it to the producer channel
	//and then close the channel
	go func(inp string, ch chan<- string) {
		for i := 0; i < 10; i++ {
			ch <- fmt.Sprintf("%s: %d\n", inp, i+1)
		}
		close(ch)
	}(input, ch)

	//return the generator channel
	return ch
}
