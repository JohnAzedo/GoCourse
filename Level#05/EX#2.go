package main
// Struct Map
import (
	"fmt"
)

type Person struct {
	name string
	lastName string
	flavors []string
}

func main() {
	myMap := map[string]Person{}

	personOne := Person{
		name: "John",
		lastName: "Lemon",
		flavors: []string{"chocolate", "vanilla"},
	}

	personTwo := Person{
		name: "Louis",
		lastName: "Lemon",
		flavors: []string{"chocolate", "Banana"},
	}

	myMap[personOne.name] = personOne
	myMap[personTwo.name] = personTwo

	
	for key, person := range myMap{
		fmt.Println(key, person.lastName, ":")
		for _, flavor := range person.flavors{
			fmt.Println("\t", flavor)
		}
	}
}