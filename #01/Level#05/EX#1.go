package main
// Struct person
import (
	"fmt"
)

type Person struct {
	name string
	lastName string
	flavors []string
}

func main() {
	person := Person{
		name: "John",
		lastName: "Lemon",
		flavors: []string{"chocolate", "vanilla"},
	}

	fmt.Println(person)
}