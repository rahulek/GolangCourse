package main

import (
	"fmt"
	"sync"
)

type User struct {
	Name string
	Age  int
}

// CustomTypeChannelDemo function accepts title parameter
// and demonstrates the use of exchaning custom data types
// over the channel
func CustomTypeChannelDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//Channel go exchange the User struct
	ch := make(chan User)
	//WaitGroup
	wg := sync.WaitGroup{}

	//2 goroutines
	wg.Add(2)

	//Reader gorotine reads User structs from the channel
	go func(userCh <-chan User) {
		defer wg.Done()

		for u := range userCh {
			fmt.Printf("\t\tCustomTypeChannelDemo: User: %v\n", u)
		}
	}(ch)

	//Writer goroutine writes 5 Users and then closes the channel
	//to indicate to the reader that all the data has been written!
	go func(userCh chan<- User) {
		defer wg.Done()

		for i := 0; i < 5; i++ {
			aUser := User{}
			aUser.Name = fmt.Sprintf("User %d", i+1)
			aUser.Age = i + 10
			userCh <- aUser
		}

		//Close the ch
		close(userCh)

		//One the channel is closed, it can't be written to!
		//Uncomment below line to see the panic
		//userCh <- User{}

	}(ch)

	//Wait for the waitGroup to complete
	wg.Wait()

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}
