package main

import (
	"fmt"

	"DeferPanicRecover/deferred"
	"DeferPanicRecover/panicrecover"
)

func main() {
	fmt.Println("---------- Deferred Execution : START ---------")

	deferred.SimpleDeferDemo("Simple Defer Usage")
	deferred.MultiDeferDemo("Multiple Defer Usage")
	deferred.DeferContextDemo("Defer Context")
	deferred.SingleFileOpenCloseDemo("Single File Open and Close with Defer")
	deferred.HttpGetDemo("HTTP GET and Defer")

	//Uncomment the following line to create a program panic scenario
	//panicrecover.SimplePanicDemo("Inducing panic by creating a Divide-by-0 error")

	//Uncomment the following line to create a program panic scenario
	//panicrecover.DevPanicDemo("Go program uses panic to create a panic")

	//Use Ctrl+C to stop the server started by the following code
	//Run the same program again in a separate terminal.
	//Because the HTTP server is already started on port 8080, second attempt
	//will cause an error and panic() by the HTTP server
	//Uncomment to see the demo
	//panicrecover.GenuinePanicDemo("HTTP Get server that panics")

	panicrecover.RecoverDemo("Recovering from a panic")

	fmt.Println("---------- Deferred Execution : END ---------")
}
