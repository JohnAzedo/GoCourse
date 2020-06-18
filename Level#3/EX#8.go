package main
// Just a switch declaration......
import (
	"fmt"
)

func main() {
	
	x := 3
	
	switch true {
		case x == 1: fmt.Println("x is 1")
		case x == 2: fmt.Println("x is 2")
		default: fmt.Println("It's none of your business")
	}
}