package main
// Using defer

import (
	"fmt"
)

func main() {
	wakeUp()
}

func wakeUp() {
	defer fmt.Println("Wait, You must wake up first!")
	fmt.Println("Go out bed")
	fmt.Println("Brushing teeth")
	fmt.Println("Start studying")
}