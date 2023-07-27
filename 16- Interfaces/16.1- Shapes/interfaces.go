package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

func calcArea(f shape) {
	fmt.Printf("The shape area is %0.2f", f.area())
}

type rectangle struct {
	height float64
	width  float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}


type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

func main() {
	r := rectangle{15, 25}
	calcArea(r)

	c := circle{10}
	calcArea(c)
}