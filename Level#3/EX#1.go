package main
// Put on screen the number from 1 to 10000
import (
	"fmt"
)

func main() {
	for x := 1; x <= 10000; x++ {
		fmt.Print(" ", x)
	}
}
