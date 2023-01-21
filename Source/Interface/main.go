package main

import "fmt"

func main() {
	fmt.Println("---------  Interfaces : START -----------")

	WriterInterfaceDemo("Writer Interface")
	InterfaceOnNonStructTypeDemo("Interface on a non-struct type")
	InterfaceCompositionDemo("Composing interfaces from other smaller interfaces")
	TypeConversionDemo("Type conversions for the interfaces")
	EmptyInterfaceDemo("Empty Interface")

	fmt.Println("---------  Interfaces : END -----------")
}
