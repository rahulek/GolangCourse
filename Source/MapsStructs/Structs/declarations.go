package structs

import "fmt"

// DeclarationsDemo function accepts a title string
// and demonstrates the struct declarations
func DeclarationsDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v: START ***\n", title)

	//Structs are type declared first
	type aBook struct {
		title         string
		author        string
		priceRs       float32
		publishedYear int
	}

	//Based on the type, multiple structs can be instantiated
	//Literal initilization: STYLE NOT RECOMMENDED
	//because fields are not-named.
	//If the struct definition changes, declared struct may
	//assign a different value to the same field
	book1 := aBook{
		"Shahnameh",
		"Abu-al-Qasem Firdawsi",
		1456.87,
		2001,
	}
	fmt.Printf("\tbook1: %v\n", book1)

	//Structs are intialized using the field-names
	//With the field names specified, ordering of
	//initialization could be totally different
	book2 := aBook{
		priceRs:       675.98,
		publishedYear: 2006,
		title:         "Masnawi",
		author:        "Maulana Rumi",
	}
	fmt.Printf("\tbook2: %v\n", book2)
	fmt.Printf("\tbook2: %#v\n", book2)

	//Initialize lazily.
	//Declare a struct type variable
	//Access each field with a dot (.) and set its value
	book3 := aBook{}
	book3.title = "Hindu Iconography"
	book3.author = "T.A.Gopinath Rao"
	book3.priceRs = 2329.00
	book3.publishedYear = 1912
	fmt.Printf("\tbook3: %#v\n", book3)

	//Anonymous structs are declared and initialized in-place
	firtRomanEmperor := struct {
		name string
		died int
	}{
		"Julius Caesar",
		-44,
	}
	fmt.Printf("\tfirstRomanEmperor: %#v\n", firtRomanEmperor)

	fmt.Printf("*** %v: END ***\n", title)
	fmt.Printf("\n")

}
