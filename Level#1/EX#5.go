package main
// Link: https://www.youtube.com/watch?v=Q7ZrIDMj9zc&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=21


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