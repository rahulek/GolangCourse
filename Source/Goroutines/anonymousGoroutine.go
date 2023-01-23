package main

import (
	"fmt"
	"time"
)

// AnonymousGoroutineDemo function accepts a title
// parameter and demonstrates the use of anonymous
// functions as Goroutines.
func AnonymousGoroutineDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//Start an anonymous function as a goroutine
	//The goroutine use greetMsg local variable that
	//is defined in outside function. That is, the greetMsg
	//is closed over and being used inside the anonymous function
	greetMsg := "Welcome! How are you?"
	go func() {
		fmt.Printf("\t\tAnonymous function: greetMsg: %v\n", greetMsg)
	}()

	//Artificial sleep to let the anonynous goroutine execute
	time.Sleep(2 * time.Millisecond)

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}

// AnonymousGoroutineDemo2 function accepts a title
// parameter and demonstrates the use of anonymous
// functions as Goroutines.
func AnonymousGoroutineDemo2(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//Start an anonymous function as a goroutine
	//inside a for loop. The loop variable is is now
	//captured/closed-over and used inside that goroutine
	//The goroutine will execute concurrent to the for loop
	//and therefore the i counter may see a diffrernt value
	//when the gorountine excutes which is a problem

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("\t\tAnonymous function2 PROBLEMATIC: i: %v\n", i)
		}()
	}

	//To fix this problem, anonynous function can be passed
	//a value of the loop counter. Now, the goroutine gets it
	//own copy of the loop counter and the output reflect the
	//intended result.
	for i := 0; i < 10; i++ {
		go func(counter int) {
			fmt.Printf("\t\tAnonymous function2 OK: counter: %v\n", counter)
		}(i)
	}

	//Artificial sleep to let the anonynous goroutine execute
	time.Sleep(3 * time.Millisecond)

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}

// AnonymousGoroutineDemo3 function accepts a title
// parameter and demonstrates the use of anonymous
// functions as Goroutines.
func AnonymousGoroutineDemo3(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//To fix the problem demonstarted by Anonynous2Demo(), the loop counter's
	//current value can be passed in as a function parameter. Now, the goroutine gets it
	//own copy of the loop counter and the output reflect the
	//intended result.
	for i := 0; i < 10; i++ {
		go func(counter int) {
			fmt.Printf("\t\tAnonymous function3 OK: counter: %v\n", counter)
		}(i)
	}

	//Artificial sleep to let the anonynous goroutine execute
	time.Sleep(2 * time.Millisecond)

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}
