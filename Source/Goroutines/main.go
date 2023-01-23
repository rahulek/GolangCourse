package main

import "fmt"

func main() {
	fmt.Println("------------ Goroutines : START ---------------")

	//Execute only one function at a time
	//SimpleGoroutineDemo("Simple Goroutine")
	//SimpleSolutionDemo("Simple solution to the goroutine not executing problem")
	//AnonymousGoroutineDemo("Anonymous function as a goroutine")
	//AnonymousGoroutineDemo2("Anonymous function2 as a goroutine inside a for loop")
	//AnonymousGoroutineDemo3("Anonymous function2 as a goroutine inside a for loop: FIXED")
	//WaitGroupSyncDemo("WaitGroup based synchronization")
	//WaitGroupSyncDemo2("WaitGroup based synchronization - with a for loop")
	//RaceConditionDemo("Race Condition")
	//SyncWithMutexDemo("Synchronizatin with a RWMutex")
	SyncWithMutexDemo2("Synchronizatin with a RWMutex - FIXED")

	fmt.Println("------------ Goroutines : END ---------------")
}
