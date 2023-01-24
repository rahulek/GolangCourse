package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("------- URLFetcher Example : START ------")
	fmt.Println()

	//Check if usable arguments provided
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "\tUSAGE: URLFetcher ...urlsToFetch\n")
		os.Exit(-1)
	}

	//Start the timer
	start := time.Now()

	//Data Channel receives the formatted result from the fetcher goroutine
	dataCh := make(chan string)

	//For all the arguments provides, start a goroutine to
	//fetch the URL
	for _, url := range os.Args[1:] {
		go fetchURL(url, dataCh)
	}

	//Wait for all the goroutines to finish.
	//We started len(os.Args[1:]) goroutines.
	//We should read only those mant returns
	numURLs := len(os.Args[1:])
	for i := 0; i < numURLs; i++ {
		res := <-dataCh
		fmt.Printf("\tResult from fetchURL: %v\n", res)
	}

	//Print the total time required to fetch all the URLs.
	fmt.Printf("\t*** %.2fs: total time.***\n\n", time.Since(start).Seconds())

	fmt.Println("------- URLFetcher Example : END ------")
}

func fetchURL(url string, ch chan<- string) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("URL: %s: Error: %v\n", url, err)
		return
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("URL: %s: Error: %v\n", url, err)
		return
	}

	ch <- fmt.Sprintf("URL: %s, Read: %v bytes successfully.\n", url, len(data))
}
