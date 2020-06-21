package main
// Callbacks

import(
	"fmt"
)

func main(){
	function(callback)
}

func callback() {
	fmt.Println("This is a test!")
}

func function(cb func()) {
	fmt.Println("First Function!")
	cb()
}