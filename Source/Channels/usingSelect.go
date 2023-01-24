package main

import "fmt"

type MyUser struct {
	Name string
	Age  int
}

// SelectDemo function accepts title parameter
// and demonstrates the use of the *done* channels.
// These are not special Go channels but typically used
// in Go program to indicate that some processing should be
// done (over!!)
func SelectDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//In this example, we don't make use of the waitGroups.
	//Instead, the function will signal to the reader goroutine
	//to stop processing by sending a *done* signal

	//Create a channel for sending a User struct
	userCh := make(chan MyUser)
	//Create a channel that's used to send a *done*
	//signal
	doneCh := make(chan struct{}) //Empty struct{} does not take any memory

	go readerWithSelect(userCh, doneCh)

	//Write 10 users and then send the done signal
	for i := 0; i < 10; i++ {
		aUser := MyUser{}
		aUser.Name = fmt.Sprintf("User %d", i+1)
		aUser.Age = i + 10
		userCh <- aUser
	}

	//All done
	doneCh <- struct{}{}

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}

func readerWithSelect(userCh <-chan MyUser, doneCh <-chan struct{}) {
	for {
		//The select is a multiplexer like switch. If picks data from
		//whichever channel is ready. If there is no data, it continues
		//to wait.
		select {
		case user := <-userCh:
			fmt.Printf("\t\tSelectDemo: reader received: %v\n", user)
		case <-doneCh:
			fmt.Printf("\t\tSelectDemo: doneChan signal received. **FINISHED**\n")
			break
		}
	}
}
