package main

// Put on screen all capital letters (from 65 to 90)
import (
	"fmt"
)

func main() {
	for i := 65; i <= 90; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(string(i), " ")
			// fmt.Printf("%#U ", i)
		}

		fmt.Println()
	}
}
