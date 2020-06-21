package main
// Create your own type 

import (
	"fmt"
)

type lemon int
var x lemon = 7

func main() { 
	fmt.Printf("Value: %v\nType: %T\n", x, x)
	x = 42
	fmt.Println(x)
}
