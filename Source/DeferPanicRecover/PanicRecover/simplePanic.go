package panicrecover

import (
	"fmt"
	"net/http"
)

// SimplePanicDemo function accepts title parameter
// and demonstrates how panic is caused
func SimplePanicDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	a, b := 1, 0
	//Create a Div-by-0 situation which will induces a panic
	divResult := a / b
	fmt.Printf("\tdivResult: %v\n", divResult)
}

func DevPanicDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//Developer can use panic() call to create a panic
	//panic() is used for an exception condition in the program

	panic("Exception situation: SOS")
}

// GenuinePanicDemo function accepts title parameter
// and becomes a simple HTTP server to handle a GET
// request. It starts itself on port 8080 and panics
// when it can't be started successfully
func GenuinePanicDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : BEGIN ***\n", title)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Go Http Server"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		//Server could not start listening on 8080
		//Really an exception situation and reason to panic
		panic(err.Error())
	}
}
