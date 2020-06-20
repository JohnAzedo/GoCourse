package main

import(
	"fmt"
)

func main() {
	factorial(5)
}

func factorial(int number) int{
	if number == 1 {
		return number
	}else{
		return number * factorial(number-1)
	}

}