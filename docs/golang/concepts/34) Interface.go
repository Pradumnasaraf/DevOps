package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}
type circle struct {
	radius float64
}

type rectangle struct {
	width  float64
	length float64
}

func (r rectangle) area() float64 {
	return r.length * r.width
}

func (c circle) area() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	myRectangle := rectangle{width: 8, length: 10}
	fmt.Println(myRectangle.area())

	myCircle := circle{radius: 10}
	fmt.Println(myCircle.area())

	
}
