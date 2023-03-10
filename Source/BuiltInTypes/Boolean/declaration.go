package Boolean

import "fmt"

// BoolDeclarationDemo function accepts the title parameter and then
// demonstrates the declaration of the boolean variables
func BoolDeclarationDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	t := true
	f := false
	fmt.Printf("\tt= %v, %[1]T\n", t)
	fmt.Printf("\tf= %v, %[1]T\n", f)

	weekDay := 4
	isSunday := weekDay == 7
	fmt.Printf("\tweekDay: %v, isSunday: %v (%[2]T)\n", weekDay, isSunday)

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}
