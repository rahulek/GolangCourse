package main

import (
	"fmt"
	"time"
)

type MyTimeoutUser struct {
	Name string
	Age  int
}

// TimeoutChannelDemo function accepts title parameter
// and demonstartes use of the timeout channels
func TimeoutChannelDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//In this example, we don't make use of the waitGroups.
	//Here, the reader will create a timeout channel. Once
	//it times out, it will write to the done channel to
	//indicate that it has finished.
	//This function will use done channel to wait for the
	//reader to be over and timeout.

	//Create a channel for sending a User struct
	userCh := make(chan MyTimeoutUser)
	doneCh := make(chan struct{})

	go readerWithSelectWithTimeoutCh(userCh, doneCh)

	//Write 10 users and then send the done signal
	for i := 0; i < 10; i++ {
		aUser := MyTimeoutUser{}
		aUser.Name = fmt.Sprintf("User %d", i+1)
		aUser.Age = i + 10
		userCh <- aUser
	}

	//All writing done. Wait for the done channel
	//it will be signaled by the writer when it
	//times out
	<-doneCh

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}

func readerWithSelectWithTimeoutCh(userCh <-chan MyTimeoutUser, doneCh chan<- struct{}) {
	//Create a timeout channel that will timeout after 5000ms (5s)
	timeoutCh := time.After(5000 * time.Millisecond)
	fmt.Printf("\t\tTimeoutChannelDemo: Timeout is set to 5s.\n\n")
	for {
		//The select is a multiplexer like switch. If picks data from
		//whichever channel is ready. If there is no data, it continues
		//to wait.
		select {
		case user := <-userCh:
			fmt.Printf("\t\tTimeoutChannelDemo: reader received: %v\n", user)
		case <-timeoutCh:
			fmt.Printf("\t\tTimeoutChannelDemo: **TIMED OUT, setting the done channel.**\n")
			doneCh <- struct{}{}
			break
		}
	}
}
