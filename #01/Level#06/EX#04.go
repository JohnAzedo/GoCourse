package main
// Using methods

import (
	"fmt"
)

type Person struct {
	name string
	lastName string
	age int
}

func (person Person) toString() {
	fmt.Println(person.name, person.lastName, ":", person.age)
} 

func main() {
	person := Person{
		name: "John",
		lastName: "Lemon",
		age: 21,
	}

	person.toString()
}
