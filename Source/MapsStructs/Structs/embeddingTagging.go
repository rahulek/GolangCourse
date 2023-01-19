package structs

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// EmbeddingDemo function accepts a title parameter
// and demonstrates the composition technique
func EmbeddingDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	//Embedding is a composition technique
	//Instead of using classic OOP's inheritance to
	//reuse the functionality, composition allows
	//the reuse using embedding

	//a 2D point has (x, y) co-ordinates
	type Point2D struct {
		x, y int
	}

	//a 3D point has an additional z co-ordinate
	//but we use 2D's (x, y) by embedding it and
	//declaring additional z only
	type Point3D struct {
		Point2D
		z int
	}

	point2D := Point2D{x: 10, y: 20}
	//NOTE: "point2D" is copied into point3d.Point2D (structs are value-types)
	point3D := Point3D{Point2D: point2D, z: 100}
	fmt.Printf("\t2D Point: %#v, 3D Point: %#v\n", point2D, point3D)
	//Chang point2D and see if it affects points3D
	point2D.x *= 2
	fmt.Printf("\t2D Point: %#v, 3D Point: %#v\n", point2D, point3D)

	//Accessing a 3D point
	//NOTE - x, y are directly avaialble of the Point3D struct
	//No need to do - point3D.Point2D.x, etc
	fmt.Printf("\t3D Point Access simple: (x, y, z)= %v, %v, %v\n", point3D.x, point3D.y, point3D.z)

	//But can be done, as shown below
	fmt.Printf("\t3D Point Access complex: (x, y, z)= %v, %v, %v\n", point3D.Point2D.x, point3D.Point2D.y, point3D.z)

	fmt.Printf("*** %v : END ***\n", title)
	fmt.Printf("\n")
}

// TaggingDemo function accepts a title parameter
// and demonstrates the use of field tagging.
func TaggingDemo(title string) {
	fmt.Printf("\n")
	fmt.Printf("*** %v : START ***\n", title)

	// Tagging the field name allows the program to
	// add a meta-data for each field. This meta-data
	// is for the use of program indirectly through the
	// 3rd party package (like JSON) OR directly to the
	// program using Go's "reflect" package
	// This mechanism is not ONLY for JSON processing but
	// for any meta-data that's processable using specific
	// packages.
	type User struct {
		Name     string `json:"name"`
		Password string `json:"passowrd"`
		Address  string `json:"address,omitempty"`
	}

	user1 := User{Name: "User1", Password: "Password1", Address: "Not-empty address"}
	user2 := User{Name: "User2", Password: "Password2", Address: ""}

	//JSON Encode User1
	//All fields should get JSON encoded
	encodedUser1, err := json.MarshalIndent(user1, "", "\t\t")
	if err != nil {
		fmt.Printf("\tTaggingDemo: JSON Marshaling failed")
	} else {
		fmt.Printf("\tTaggingDemo: Encoded user1: %v\n", string(encodedUser1))
	}

	//JSON Encode user2
	//Because the address field is empty, it won't get encoded due to
	//"omitempty" flag
	encodedUser2, err := json.MarshalIndent(user2, "", "\t\t")
	if err != nil {
		fmt.Printf("\tTaggingDemo: JSON Marshaling failed")
	} else {
		fmt.Printf("\tTaggingDemo: Encoded user1: %v\n", string(encodedUser2))
	}

	//Use Reflection to get the meta-data
	t := reflect.TypeOf(User{})
	f, _ := t.FieldByName("Name")
	fmt.Printf("\tUser.Name field's tags: %v\n", f.Tag)
	f, _ = t.FieldByName("Address")
	fmt.Printf("\tUser.Address field's tags: %v\n", f.Tag)

	fmt.Printf("*** %v : START ***\n", title)
	fmt.Printf("\n")
}
