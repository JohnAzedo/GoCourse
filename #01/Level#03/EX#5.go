package main
// Module 4 from 10 to 100 with continue
import (
	"fmt"
)

func main() {
	
	for x := 10; x<=100; x++ {
		if x%4 != 0 {
			continue  
		}
		
		fmt.Println(x)
	}
}