package main
// Link: https://www.youtube.com/watch?v=Q7ZrIDMj9zc&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=20


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
