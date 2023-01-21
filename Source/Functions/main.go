package main

import "fmt"

func main() {
	fmt.Println("----------  Functions : START ----------")

	DeclareAndUseDemo("Function Declaration and Invocation")
	ParamPassingDemo("Parameters are passed by value")
	VariadicDemo("Variadic functions take variable number of arguments")
	ReturnValueDemo("Return values from the function call")
	FuncAsTypeDemo("Functions are first-class types")
	FuncAsTypeMethodsDemo("Type methods are functions associates with that type")

	fmt.Println("----------  Functions : END ----------")
}
