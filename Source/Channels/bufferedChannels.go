package main

import (
	"fmt"
	"sync"
	"time"
)

// BufferedChannelDemo function accepts a title parameter
// and demonstrates the use of the buffered channel.
func BufferedChannelDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//Buffered channel
	//When creating a channel, a non-zero 2nd parameter creates
	//a buffered channel. That's the channel capacity until which
	//the writer goroutine doesn't block. The reader goroutine can
	//continue at a slower pace than the writer with the use of the
	//buffered channels. The reader goroutine will block only when
	//the buffer is fully drained
	bufferedCh := make(chan int, 10)
	wg := sync.WaitGroup{}

	//2 goroutines will work
	wg.Add(2)

	go readerWithBufferedCh(&wg, bufferedCh)
	go writerWithBufferedCh(&wg, bufferedCh)

	//wait for both the goroutines to finish
	wg.Wait()

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}

func readerWithBufferedCh(wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()

	//pull of exact 10 numbers from the channel
	//put an artifical delay to simulate a slow
	//data consumer/reader
	for i := 0; i < 10; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("\t\tBufferedChannelDemo: Reader: read: %v\n", <-ch)
	}
}

func writerWithBufferedCh(wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()

	//pull of exact 10 numbers from the channel
	//put an artifical delay to simulate a slow
	//data consumer/reader
	for i := 0; i < 10; i++ {
		fmt.Printf("\t\tBufferedChannelDemo: Writer: writing: %v\n", i)
		ch <- i
	}
}
