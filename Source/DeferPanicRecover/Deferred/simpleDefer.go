package deferred

import "fmt"

// SimpleDeferDemo function accepts a title parameter
// and demonstrates the use of simplistic deferred execution
func SimpleDeferDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//defer defers the function code exection
	//the deferred function is invoked when the containing function
	//is finished.

	//First, let's see simple sequential execution
	//Each line will print sequentially
	fmt.Println("First Line")
	fmt.Println("Second Line")
	fmt.Println("Third Line")

	//Now, defer the execution of the second line
	fmt.Println("~~~~ Now with the defer ~~~")
	fmt.Println("First Line")
	defer fmt.Println("DEFERRED: Second Line") //Defer the execution of the Println() call
	fmt.Println("Third Line")

	fmt.Printf("*** %v : START ***\n", title)
	fmt.Printf("\n")
}

// MultiDeferDemo function accepts the title parameter
// and demonstrates the use of multiple defer statements.
func MultiDeferDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//When multiple defer statements appear, they are
	//executed in the LIFO (Last-in-First-Out) order

	fmt.Println("First Line")
	fmt.Println("Second Line")
	defer fmt.Println("~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~")
	defer fmt.Println("DEFERRED: Third Line")
	defer fmt.Println("DEFERRED: Fourth Line")
	fmt.Println("Fifth Line")

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}

// DeferContextDemo function accepts title parameter
// and demonstrates the context for the defer statement
func DeferContextDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	ctx := "Defer Context at the Start"
	//What will the deferred function print?
	defer func(dc string) {
		fmt.Printf("\tCtx is: %v\n", dc)
		defer fmt.Println("~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~")
	}(ctx)
	ctx = "Defer Context at the End"

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}
