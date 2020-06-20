package main
// Anonymous struct

import (
	"fmt"
)

func main() {
	anon := struct{
		name string
		flavor string
	}{
		name: "Pizza",
		flavor: "Marguerita",
	}

	fmt.Println(anon)


}