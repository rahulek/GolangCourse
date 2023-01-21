package main

import (
	"fmt"
	"reflect"
)

// Multiple smaller interfaces
// Walker-Eater-Sleeper
type Walker interface {
	Walk() string
}

type Eater interface {
	Eat(food string) string
}

type Sleeper interface {
	Sleep(hours int) string
}

// A composed interface that's
// composed of other smaller interfaces
type Animal interface {
	Walker
	Eater
	Sleeper
}

// A Concrete type that implements Animal behavior
type Dog struct {
	Name          string
	Tail          string
	NumTeeth      int
	MinSleepHours int
}

// Another Concrete type that implements the Animal interface
type Human struct {
	Name string
}

// ---- Dog implementing the Animal behavior ----------
func (d *Dog) Walk() string {
	return fmt.Sprintf("\tDog %s is walking on all 4 with its %s tail\n", d.Name, d.Tail)
}

func (d *Dog) Eat(food string) string {
	return fmt.Sprintf("\tDog %s is eating %s with its %d teeth\n", d.Name, food, d.NumTeeth)
}

func (d *Dog) Sleep(numHours int) string {
	var ret string
	if numHours < d.MinSleepHours {
		ret = fmt.Sprintf("\tDog %s will be unhappy, it needs min %d hours of sleep\n", d.Name, d.MinSleepHours)
	} else {
		ret = fmt.Sprintf("\tDog %s Happy, its going to sleep now.\n", d.Name)
	}

	return ret
}

// ---- Human implementing the Animal behavior ----------
func (h *Human) Walk() string {
	return fmt.Sprintf("\tHuman %s is walking on his 2.\n", h.Name)
}

func (h *Human) Eat(food string) string {
	return fmt.Sprintf("\tHuman %s is eating %s happily.\n", h.Name, food)
}

func (h *Human) Sleep(numHours int) string {
	return fmt.Sprintf("\tHuman %s really Happy, its going to sleep now.\n", h.Name)
}

// InterfaceCompositionDemo function accepts title parameter
// and demonstrates the interface composition technique
func InterfaceCompositionDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	var d Dog = Dog{Name: "Jacky", NumTeeth: 40, Tail: "Curly", MinSleepHours: 10}
	fmt.Printf("\tDog Walking: %s", d.Walk())
	fmt.Printf("\tDog Eating: %s", d.Eat("Dog food"))
	fmt.Printf("\tDog Sleeping: %s", d.Sleep(4))

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}

func animalOperations(a Animal) {
	fmt.Printf("\tanimalOperations: Start\n")
	fmt.Printf("\t\t%s", a.Walk())
	fmt.Printf("\t\t%s", a.Eat("Meat and Eggs"))
	fmt.Printf("\t\t%s", a.Sleep(12))
	fmt.Printf("\tanimalOperations: End\n")
}

// TypeConversionDemo function accepts title parameter
// and demonstrates the type conversion of interfaces
func TypeConversionDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	var d Dog = Dog{Name: "Jacky", NumTeeth: 40, Tail: "Curly", MinSleepHours: 10}
	var a Animal = Animal(&d)
	var w Walker = Walker(a)
	e, s := Eater(a), Sleeper(a)
	fmt.Printf("\td: %T, a: %T, w: %T, e: %T, s: %T\n", d, a, w, e, s)

	//'a' is an Animal with an underlying type of a Dog
	animalOperations(a)

	var h Human = Human{Name: "James"}
	//Now h is also an animal
	animalOperations(Animal(&h))

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}

// EmptyInterfaceDemo function accepts title parameter
// and demonstrates the use of empty parameters
func EmptyInterfaceDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//An empty interface is an interface that does not have any method
	//Therefore, all the types are automatically considered to have implemented
	//the empty interface. Therefore, an empty interface (interface{}) can be
	//used as a placeholder for "any" type. Go 1.18+ program can use "any" keyword
	//to mean the same thing.
	var x interface{} = 10
	var y interface{} = 240.34

	//Type conversion OR reflection can be used to find the underlying type of
	//the empty interface vars
	switch x.(type) {
	case int:
		fmt.Printf("\tx: %v is an int\n", x)
	default:
		fmt.Printf("\tx: %v, unknown type\n", x)
	}

	//Use reflection to find an underlying type
	fmt.Printf("\tY: %v, its type is: %v\n", y, reflect.ValueOf(y).Kind())

	//Go 1.18+, use 'any' instead of interface{}
	var anyVar any = "a string"
	fmt.Printf("\tanyVar: %v, type: %[1]T\n", anyVar)
	anyVar = 12
	fmt.Printf("\tNOW: anyVar: %v, type: %[1]T\n", anyVar)

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}
