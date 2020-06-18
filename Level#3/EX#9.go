package main
// Just a switch declaration......
import (
	"fmt"
)

func main() {
	
	favoriteSport := "basketball"
	
	switch favoriteSport {
		case "soccer": fmt.Println("Not")
		case "football": fmt.Println("Not")
		case "e-Sports": fmt.Println("Not")
		default: fmt.Println("It's basketball")
	}
}