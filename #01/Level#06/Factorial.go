package main

import(
	"fmt"
)

func main() {
	fmt.Println(factorial(4))
}

func factorial(number int) int{
	
	if number == 0 {
		return number
	}else{
		return number * factorial(number-1)
	}

}