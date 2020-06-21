package main
// Get function scope: Closers

import(
	"fmt"
)

func main(){
	// Get function scope
	a := function()

	for i:=0; i<10; i++ {
		fmt.Println(a())
	}

}


func function() func() int{
	x := 0
	return func() int{
		x++
		return x
	}
}