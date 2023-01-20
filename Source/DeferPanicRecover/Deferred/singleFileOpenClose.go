package deferred

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// SingleFileOpenCloseDemo function accepts title parameter
// and demonstrates the application of the defer functionality.
func SingleFileOpenCloseDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//Open a file
	testfileloc, err := filepath.Abs("./testData/fileOne.txt")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(testfileloc)
	if err != nil {
		//bail out if error
		log.Fatal(err)
	}

	//*** defer closing the file ***
	//This will close the file when the SingleFileOpenCloseDemo() function
	//returns. This means that a precious resources such as a file handle is
	//properly released.
	defer func(fp *os.File) {
		fmt.Printf("\tINSIDE THE DEFERRED FUNCTION: now closing the file.\n")
		fp.Close()
		fmt.Println("~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ")
	}(f)

	//Scanner takes the file handle f for the reading purposes
	scanner := bufio.NewScanner(f)

	//Continue to read from the scanner (that is the file f provided to it)
	for scanner.Scan() {
		fmt.Printf("\tTEXT READ: %v\n", scanner.Text())
	}

	//if at all scanner has any errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}
