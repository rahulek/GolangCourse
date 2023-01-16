package main

import (
	"ArraysSlices/arrays"
	"ArraysSlices/slices"
	"fmt"
)

func main() {
	fmt.Println("---------- Arrays and Slices: START ----------")

	arrays.DeclarationsDemo("Array Declarations")
	arrays.MultiDimensionsDemo("Multi-Dimensional Arrays")
	arrays.OperationsDemo("Array Operations")

	slices.DeclarationsDemo("Slice Declarations")
	slices.CreationDemo("Slice Creation")
	slices.SliceAsSharedViewDemo("Slices as Shared View")
	slices.SliceLenCapDemo("Slice Length and Capacity")
	slices.AppendWeirdnessDemo("Slices can refer to stale data")
	slices.CopySliceDemo("copy() to copy the slices")

	fmt.Println("---------- Arrays and Slices: END ----------")
}
