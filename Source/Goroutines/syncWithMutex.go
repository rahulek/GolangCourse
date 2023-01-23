package main

import (
	"fmt"
	"sync"
)

// global data
var counter2 int = 0

// waitgroup for overall sync
var wg sync.WaitGroup = sync.WaitGroup{}

// a mutex
var m sync.RWMutex = sync.RWMutex{}

// SyncWithMutexDemo function accepts title parameter
// and demonstrates the use of the Mutex based
// synchnization between the goroutines - this is
// an example of a brut-force synchronization with no
// use of concurrent mechanisms such as channels.
func SyncWithMutexDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	for i := 0; i < 10; i++ {
		//2 goroutines started with every loop iteration
		wg.Add(2)
		go reader2()
		go writer2()
	}

	//The result is synchronized output (increasing seq) but still random
	//across executions.

	wg.Wait()

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")

}

func reader2() {
	//Readers use a read-lock
	m.RLock()
	fmt.Printf("\t\treader Goroutine: counter = %v\n", counter2)
	//Reader Unlock
	m.RUnlock()

	//Signal the waitgroup
	wg.Done()
}

func writer2() {
	//Writers use only Lock
	m.Lock()
	counter2++
	//Wruter Unlock
	m.Unlock()

	//Signal the waitgroup
	wg.Done()
}

// SyncWithMutexDemo function accepts title parameter
// and demonstrates the use of the Mutex based
// synchnization between the goroutines
func SyncWithMutexDemo2(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	for i := 0; i < 10; i++ {
		//2 goroutines started with every loop iteration
		wg.Add(2)
		m.RLock()
		go reader22()
		m.Lock()
		go writer22()
	}

	//The result is synchronized output (increasing seq) but still random
	//across executions.

	wg.Wait()

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")

}

func reader22() {
	fmt.Printf("\t\treader Goroutine: counter = %v\n", counter2)
	//Reader Unlock
	m.RUnlock()

	//Signal the waitgroup
	wg.Done()
}

func writer22() {
	counter2++
	//Wruter Unlock
	m.Unlock()

	//Signal the waitgroup
	wg.Done()
}
