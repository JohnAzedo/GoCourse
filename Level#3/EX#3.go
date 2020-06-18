package main
// Year born until chosen year
import (
	"fmt"
)

func main() {
	
	yearBorn := 1999
	chosenYear := 2077
	
	for yearBorn <= chosenYear {
		fmt.Println(yearBorn)
		yearBorn++;
	}	
}
