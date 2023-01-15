package Integertypes

import "fmt"

// StdOperationsDemo function takes a title argument and
// demonstrates the standard operations on the integer data type
func StdOperationsDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v: START ***\n", title)

	//add, subtract, multiply, divide, remainder
	a := 136
	b := 20
	fmt.Printf("\tAddition: a: %v, b: %v, result: %v\n", a, b, a+b)
	fmt.Printf("\tSubtraction: a: %v, b: %v, result: %v\n", a, b, a-b)
	fmt.Printf("\tMultiplication: a: %v, b: %v, result: %v\n", a, b, a*b)
	//Integer division because a and b are both ints, the result also must be an int
	//and not a float/decimal
	fmt.Printf("\tDivision: a: %v, b: %v, result: %v\n", a, b, a/b)
	fmt.Printf("\tRemainder: a: %v, b: %v, result: %v\n", a, b, a%b)

	fmt.Printf("*** %v: END ***\n", title)
	fmt.Printf("\n")
}

// BitwiseOperationsDemo function accepts the title string and
// demonstrates the bitwise operators
func BitwiseOperationsDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	op1 := 0x33 //0011 0011
	op2 := 0xCC //1100 1100
	fmt.Printf("\tAND: op1: %#08b, op2: %#08b, op1 & op2: %#08b\n", op1, op2, op1&op2)
	fmt.Printf("\tOR: op1: %#08b, op2: %#08b, op1 | op2: %#08b\n", op1, op2, op1|op2)
	fmt.Printf("\tXOR: op1: %#08b, op2: %#08b, op1 ^ op2: %#08b\n", op1, op2, op1^op2)
	//x &^ y == x & (~y)
	fmt.Printf("\tAND-NOT: op1: %#08b, op2: %#08b, op1 &^ op2: %#08b\n", op1, op2, op1&^op2)
	fmt.Printf("\tNOT (unary): op1: %#08b, ~op1: %#08b\n", op1, uint8(^op1))

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}

func ShiftOperationsDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	a := 10     //00001010
	m := a << 2 //equivalent to a * 4
	d := a >> 1 //equivalent to a / 2
	fmt.Printf("\ta: %v, a << 2: %v\n", a, m)
	fmt.Printf("\ta: %v, a >> 1: %v\n", a, d)

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}
