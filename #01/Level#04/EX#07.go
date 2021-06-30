package main
// Slice of slices

import (
	"fmt"
)

func main() {
	// Second arg is length and third is capacity
	sliceOfSlices := [][]string{
		[]string{
			"name",
			"last name",
			"hobby",
		},
		[]string{
			"John",
			"Lemon",
			"Music",
		},
		[]string{
			"Barack",
			"Obama",
			"Basketball",
		},
	}

	for _, value := range sliceOfSlices {
		fmt.Println(value[0])
			for _, valueTwo := range value {
				fmt.Println("\t", valueTwo)
			}
	}

}