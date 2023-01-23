package main

import "fmt"

// SimpleGoroutineDemo function accepts a title
// parameter and demonstrates the use of simplistics
// single goroutine
func SimpleGoroutineDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//goroutine is just a "go" keyword placed in
	//front of a function call OR an anonymous
	//function declaration executed immediately
	go helloWorld()

	//goroutine is a concurrent execution tool which is
	//similar to to traditional thread in a multi-threaded environment
	//but they are very lightweight and can be created quickly and in
	//large number.
	//Go runtime muliplxes these goroutines on its' thread pool of OS
	//threads created at the runtime initialization time.

	//NOTE :
	// In this example, even though the helloWorld goroutine that prints
	// a hello world message is started, nothing gets printed out
	//
	// The reason for this problem is -
	//  when the helloWorld() goroutine is started, it doesn't immediately
	//  execute. The calling function (SimpleGoroutineDemo) which is started by
	//  the main function is also part of the main() goroutine. This function
	//  and the main function continue to execute and finish early before the helloWorld
	//  goroutine executes.
	//  When the main goroutine finishes, the program is finished the helloWorld
	//  goroutine was never executed.

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}

func helloWorld() {
	fmt.Printf("\t\thelloWorld goroutine: HELLO WORLD\n")
}
