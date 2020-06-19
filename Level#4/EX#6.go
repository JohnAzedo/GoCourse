package main
// Using make to create a slice

import (
	"fmt"
)

func main() {
	// Second arg is length and third is capacity
	slice := make([]string, 5, 5)
	slice = []string{"RN", "SP", "AM", "RJ", "RS"}
	fmt.Println(len(slice), cap(slice))
	fmt.Println(slice)
}