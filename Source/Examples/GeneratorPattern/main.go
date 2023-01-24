package main

import "fmt"

func main() {
	fmt.Println("--------- Channel Generator Pattern : START ----------")
	fmt.Println()

	//Genete two channels - one keeps saying "Hi"
	//Other keeps saying "Bye"
	hiCh := generator("Hi")
	byeCh := generator("Bye")

	//Range over hiChannel and read all its messages
	for msg := range hiCh {
		fmt.Printf("\t\tHi Channel: %s\n", msg)
	}

	//Range over byeChannel and read all its messages
	for msg := range byeCh {
		fmt.Printf("\t\tBye Channel: %s\n", msg)
	}

	fmt.Println()
	fmt.Println("--------- Channel Generator Pattern : END ----------")
}

// generator is a function that *generates* a channel
// This channel is returned to the caller and the caller
// can use the channel to listen to the messages published
// on that channel
func generator(input string) <-chan string {
	//create a channel that's returned
	retCh := make(chan string)

	//Start a goroutine as an anonymous function
	//The goroutune keeps on publishing a message
	//on the created channel and then closes the channel
	go func(inp string, ch chan<- string) {
		for i := 0; i < 5; i++ {
			ch <- fmt.Sprintf("%s: %d\n", inp, i+1)
		}

		//close the ch
		close(ch)
	}(input, retCh)

	//return the created channel
	return retCh
}
