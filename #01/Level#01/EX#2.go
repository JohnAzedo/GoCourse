package main
// Package level scope

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