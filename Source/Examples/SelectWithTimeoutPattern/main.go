package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("------------ Select With Timeout : START -----------")
	fmt.Println()

	//Create 2 talkative channels - each produces data after the specified delay
	hiCh := talkativeProducer("Hi", time.Duration(100*time.Millisecond))
	byeCh := talkativeProducer("Bye", time.Duration(50*time.Millisecond))

	//Overall timeout: the main goroutine will pick messages from hiCh and byeCh
	//and after timeout, stops processing the data
	timeOutCh := time.After(2000 * time.Millisecond)
outer:
	for {
		select {
		case hiMsg := <-hiCh:
			fmt.Printf("\t%s\n", hiMsg)
		case byeMsg := <-byeCh:
			fmt.Printf("\t%s\n", byeMsg)
		case <-timeOutCh:
			fmt.Printf("\t** TIMED OUT **\n")
			break outer
		}
	}

	fmt.Println()
	fmt.Println("------------ Select With Timeout : START -----------")
}

func talkativeProducer(input string, delay time.Duration) <-chan string {
	ch := make(chan string)

	go func(inp string, ch chan<- string) {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s : %d", inp, i+1)
			time.Sleep(delay)
		}

	}(input, ch)

	return ch
}
