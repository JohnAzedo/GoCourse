package main
// Link: https://www.youtube.com/watch?v=Q7ZrIDMj9zc&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=19



import (
	"fmt"
)

// Package level scope
// := doesn`t work here

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\n%v\n%v\n", x, y, z)
	fmt.Println(s)
}
