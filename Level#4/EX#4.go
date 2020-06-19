package main
// Put data in slice

import (
	"fmt"
)

func main() {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	slice = append(slice, 52, 1000, 232)
	fmt.Println(slice)

	sliceTwo := []int{23, 54, 34, 12}
	// Use ... after variable to get data from slice or array
	slice = append(slice, sliceTwo...)
	fmt.Println(slice)
}