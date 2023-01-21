package main

import "fmt"

// Go - Standard Interface in the io package - https://pkg.go.dev/io#Writer
//
// interfaces in Go are similar to OOP interface features in C++/Java, etc
// An interfaces simply defines a behavaior and leave it to the type/s to
// implement those.
type Writer interface {
	//only behaviors, no data
	Write(data []byte) (int, error)
}

// Custom Printer type
type MyPrinter struct {
	ModelName     string
	isLaserTech   bool
	isDoubleSided bool
}

//Print prints on paper. Printing
//is a Write operation. Therefore,
//if MyPrinter implements the Write
//method of the Writer interface, it
//automatically becomes an implementor
//of the Writer interface
//NOTE *** NO specific implements keyword
//required in Go to implement an in interface

func (p MyPrinter) Write(data []byte) (int, error) {
	fmt.Printf("\tPrinter %s is printing %d bytes of data.\n", p.ModelName, len(data))
	//all well, return the set of number of bytes written and nil error
	return len(data), nil
}

// WriterInterfaceDemo function accepts title parameter
// and demonstrates the use of a custom interface
func WriterInterfaceDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//Create a Printer var
	var laserPrinter MyPrinter = MyPrinter{ModelName: "My LaserJet Printer", isLaserTech: true, isDoubleSided: false}

	//MyPrinter implements the Writer interface
	laserPrinter.Write([]byte("Hello Go interfaces"))

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}
