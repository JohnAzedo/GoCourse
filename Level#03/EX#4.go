package main
// Year born until chosen year with infinity for
import (
	"fmt"
)

func main() {
	
	yearBorn := 1999
	chosenYear := 2077
	
	for {
		fmt.Println(yearBorn)
		yearBorn++;
		
		if yearBorn > chosenYear {
			break
		}
		
		
	}	
}