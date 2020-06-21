package main
// IOTA

import (
	"fmt"
)

const (
	_ = 1999 + iota
	b
	c
	d
	e
)

func main() {
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}