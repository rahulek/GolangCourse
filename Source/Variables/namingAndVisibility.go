package main

import (
	"Variables/mymathpackage"
	"fmt"
)

func namingAndVisibilityDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v: START ***\n", title)

	//Use the "mymathpackage" and use its exports
	addition := mymathpackage.MyAdder(10, 30)
	fmt.Printf("\tAddition using mymathpackage.MyAdder(): %v\n", addition)
	fmt.Printf("\tExported PI: %v\n", mymathpackage.MyMathPI)
	//Note: the cantBeSeen() function inside the mymathpackage(see mymathpackage\myadder.go)
	//can't be used because its not exported due to its initial letter being a lowercase

	fmt.Printf("*** %v: END ***\n", title)
	fmt.Printf("\n")
}
