package main
// Link: https://www.youtube.com/watch?v=Q7ZrIDMj9zc&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=18

import (
	"fmt"
)

// Package level scope
// := doesn`t work here
var x int
var y string
var z bool

func main() {
	fmt.Printf("%v\n%v\n%v\n", x, y, z)
	// Esses valores são chamados de nulos ou zeros
}