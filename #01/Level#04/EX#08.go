package main
// Map

import (
	"fmt"
)

func main() {
	// In this map, string is a key type and slice of string is a value type
	myMap := map[string][]string {
		"John": []string {
			"Violin",
			"Basketball",
			"Work out",
		},

		"Obama": []string {
			"Basketball",
			"Watch TV",
		},
	}

	// When you get map elements with range, this starts at the end
	for name, hobbies := range myMap {
		fmt.Println(name)
		for _, hobby := range hobbies {
			fmt.Println("\t", hobby)
		}
	}

}