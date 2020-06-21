package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

func (square Square) area() {
	result := math.Pow(square.side, 2)
	fmt.Println("Square area: ", result)
}

type Circle struct {
	radius float64
}

func (circle Circle) area() {
	result := math.Pi * 2 * circle.radius
	fmt.Println("Circle area: ", result)
}

type info interface {
	area()
}

func measure(i info){
	i.area()
}

func main() {
	square := Square{side: 15.0}
	circle := Circle{radius: 10.0}

	measure(square)
	measure(circle)
}

