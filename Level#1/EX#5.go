package main
// Convert types


import (
	"fmt"
)

type lemon int
var x lemon = 7

var y int

func main() { 
	fmt.Printf("Value x: %v\nType x: %T\n", x, x)
	x = 42
	fmt.Println(x)
	
	y = int(x)
	fmt.Printf("Value y: %v\nType y: %T\n", y, y)
	
}