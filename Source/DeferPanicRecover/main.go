package main

import (
	"fmt"

	"DeferPanicRecover/deferred"
)

func main() {
	fmt.Println("---------- Deferred Execution : START ---------")

	deferred.SimpleDeferDemo("Simple Defer Usage")
	deferred.MultiDeferDemo("Multiple Defer Usage")
	deferred.DeferContextDemo("Defer Context")
	deferred.SingleFileOpenCloseDemo("Single File Open and Close with Defer")
	deferred.HttpGetDemo("HTTP GET and Defer")

	fmt.Println("---------- Deferred Execution : END ---------")
}
