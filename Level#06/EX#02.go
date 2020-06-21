package main
// Func plus with variable parameter and Slice objects

import (
	"fmt"
)

func main() {
	
	mySlice := []int{1, 34, 43, 39, 102}
	// Get slice objects
	fmt.Println(plus(mySlice...))
	fmt.Println(plusWithSlice(mySlice))

}

// With variable parameters
func plus(values ...int) int {
	var soma int
	for _, value := range values{
		soma += value
	}

	return soma
}

// With slice argument
func plusWithSlice(slice []int) int {
	var soma int
	for _, value := range slice{
		soma += value
	}

	return soma
}





