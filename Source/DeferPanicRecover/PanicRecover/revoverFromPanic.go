package panicrecover

import "fmt"

// RecoverDemo function accepts title parameter
// and demonstrates the recover function.
func RecoverDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//Call the function that may/may not panic
	//recover() is used within the panickingFunction
	//to handle the panic. When the execution returns
	//to this function, panic scenario is recovered from
	//(if panic really happended)
	panickingFunction()

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}

func panickingFunction() {
	fmt.Printf("\tInside panickingFunction\n")

	//deferred block here uses recover() function to detect
	//any panic errors. This function returns any error causes
	//by the panic. If there is no panic, it returns nil
	//
	//NOTE : Its important to see that the defer block executes
	//even in case of a panic.
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("\t ---> recover() has detected an error:  %v <---- \n", err)
		} else {
			fmt.Printf("\tpanickingFunction returning normally\n")
		}
	}()

	//Create a panic
	//Comment the below line to see the defer block's else condition
	//running which is when the recover detects no error and everything
	//returns back to the caller
	panic("Exception situation is detected: SOS")
}
