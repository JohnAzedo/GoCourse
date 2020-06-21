package main

import (
	"fmt"
)

func main() {
	fmt.Println(returnInt())
	fmt.Println(returnIntAndString())
}

// Return Int
func returnInt() int {
	return 7
}

// Return int and string
func returnIntAndString() (int, string) {
	return 7, "Lemon"
}

