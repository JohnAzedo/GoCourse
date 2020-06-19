package main
// Deleting data from slice

import (
	"fmt"
)

func main() {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	slice = append(slice[:3], slice[5:]...)
	fmt.Println(slice)
}