package main

import (
	"fmt"

	"MapsStructs/maps"
	"MapsStructs/structs"
)

func main() {
	fmt.Println("---------- Maps and Structs : START ------------")

	maps.DeclarationsDemo("Map Declarations ")
	maps.OperationsDemo("Map Operations")

	structs.DeclarationsDemo("Struct Declarations")
	structs.OperationsDemo("Struct Access and Assignment")
	structs.EmbeddingDemo("Structs Embedding/Composition")
	structs.TaggingDemo("Struct Field Tagging")

	fmt.Println("---------- Maps and Structs : END ------------")
}
