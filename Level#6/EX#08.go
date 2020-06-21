package main
// Function that return another function

import(
	"fmt"
)

func main(){
	x := function()

	x()
}

func function() func() {
	return func() {
		fmt.Println("This is a test!")
	}
}