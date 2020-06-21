package main

import(
	"fmt"
)

func main(){
	array := [5]int{100, 200, 300, 400, 500}
	// Get a pointer of 400
	fmt.Println(&array[3])
}