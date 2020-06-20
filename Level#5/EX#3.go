package main
// Vehicle built-in struct

import (
	"fmt"
)

type Vehicle struct {
	doors int
	color string
}

type Truck struct {
	Vehicle
	fourWheels bool
}

type Sedan struct {
	Vehicle
	luxuryModel bool
}


func main() {
	carOne := Sedan{
		Vehicle{4, "blue"},
		true,
	}

	carTwo := Truck{
		Vehicle{2, "red"},
		false,
	}

	fmt.Println(carOne)
	fmt.Println(carTwo)
}