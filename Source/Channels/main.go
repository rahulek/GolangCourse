package main

import "fmt"

func main() {
	fmt.Println("-------- Channels : START ---------")

	//SimpleChannelDemo("Simple channel to communicate between two goroutines")
	//SimpleChannelWithDirectionDemo("Simple channel to communicate between two goroutines with channel direction")
	//SimpleDeadlockDemo("Goroutines that use channels can deadlock")
	//BufferedChannelDemo("Buffered Channels")
	//ChannelRangingWithDeadlockDemo("Channel ranging that demonstrates a deadlock")
	//ChannelRangingSolutionDemo("Channel ranging deadlock solution")
	//ChannelCloseDetectionDemo("Channel close without the range keyword")
	//CustomTypeChannelDemo("Channels can transport data of any custom type")
	//SelectDemo("Data of custom type exchange using the select statement")
	TimeoutChannelDemo("Using a timeout channel")

	fmt.Println("-------- Channels : END ---------")
}
