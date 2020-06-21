package main
// Put a funtion in a variable

import(
	"fmt"
)

func main(){
	x := func() int {
		return 10
	}

	fmt.Println(x())
}

