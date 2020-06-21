package main
// Pointers with struct

import(
	"fmt"
)

type Person struct {
	name string
	lastName string
}

func main(){
	person := Person{
		name: "John",
		lastName: "Lemon",
	}

	fmt.Println(person)
	changeMe(&person)
	fmt.Println(person)
}

func changeMe(person *Person){
	person.name = "Nobody cares"
	person.lastName = "with you"
}