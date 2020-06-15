package main
// Shifting bits
import (
	"fmt"
)

func main() {
	x := 200
	fmt.Printf("%d\t%b\t%#x", x, x, x)
	
	y := x << 1
	fmt.Printf("\n%d\t%b\t%#x", y, y, y)
	
	z := x << 1
	fmt.Printf("\n%d\t%b\t%#x", z, z, z)

}