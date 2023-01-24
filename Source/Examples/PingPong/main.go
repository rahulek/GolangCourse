package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("------------ PingPong Example: START -----------")
	fmt.Println()

	wg := sync.WaitGroup{}

	pingCh := make(chan string)
	pongCh := make(chan string)

	wg.Add(2)

	go pinger(&wg, pingCh, pongCh)
	go ponger(&wg, pingCh, pongCh)

	wg.Wait()

	fmt.Println()
	fmt.Println("------------ PingPong Example: END -----------")
}

func pinger(wg *sync.WaitGroup, pingCh <-chan string, pongCh chan<- string) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		pongCh <- fmt.Sprintf("PING %d", i+1)
		data := <-pingCh
		fmt.Printf("\t\tPINGER: %s\n", data)
	}
}

func ponger(wg *sync.WaitGroup, pingCh chan<- string, pongCh <-chan string) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		data := <-pongCh
		fmt.Printf("\t\tPONGER: %s\n", data)
		pingCh <- fmt.Sprintf("PONG %d", i+1)
	}
}
