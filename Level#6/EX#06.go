package main
// Anonymous function

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("Hello, playground")
	}()
}
