package deferred

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// HttpGetDemo function accepts title parameter
// and demonstrates the use of defer statement.
func HttpGetDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatal(err)
	}

	//resp.Body is a socket resource and must be freed
	//Can't simply do resp.close before reading the response
	//Therefore, defer can be used.
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\tHTTP Response Below: ")
	fmt.Printf("\t%v\n", string(data))

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}
