package main

import (
	"fmt"
	"sync"
)

// WaitGroupSyncDemo function accepts title parameter
// and demonstrates the use of the WaitGroup sync method
// to synchronized the goroutines
func WaitGroupSyncDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	wg := sync.WaitGroup{}
	greetMsg := "Welcome to the Club!"

	//Add 1 to the waitgroup to indicate one goroutine
	//needs to complete before the waitgroup is considered
	//finished
	wg.Add(1)
	go func(msg string) {
		fmt.Printf("\t\tWaitGroupSyncDemo - anonymous fn as a goroutine: msg = %v\n", msg)

		//Signal the waitgroup that this goroutine is done
		wg.Done()
	}(greetMsg)

	//No time-based artifical sync required
	//This function will block here until the whole
	//workgroup is not finished with. (here, 1 goroutune
	//must finish for the following Wait to be over!!)
	wg.Wait()

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}

// WaitGroupSyncDemo2 function accepts title parameter
// and demonstrates the use of the WaitGroup sync method
// to synchronized the goroutines. The Gorotuines are started
// within a for loop
func WaitGroupSyncDemo2(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	wg := sync.WaitGroup{}
	greetMsg := "Welcome to the Club!"

	//Start 10 goroutines and sync them with the demo function
	//using the WaitGroup
	for i := 0; i < 10; i++ {
		//Add 1 to the waitgroup to indicate one goroutine
		//needs to complete before the waitgroup is considered
		//finished
		wg.Add(1)
		go func(msg string, counter int) {
			fmt.Printf("\t\tWaitGroupSyncDemo2 - msg = %v, counter = %v\n", msg, counter)

			//Signal the waitgroup that this goroutine is done
			wg.Done()
		}(greetMsg, i)
	}

	//No time-based artifical sync required
	//This function will block here until the whole
	//workgroup is not finished with. (here, 1 goroutune
	//must finish for the following Wait to be over!!)
	wg.Wait()

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}
