package main
// Slice with 5 int values

import (
	"fmt"
)

func main() {
	slice := []int{10, 20, 30, 40, 50}

	for i, v := range slice {
		fmt.Println(i, "-", v)
	}
}