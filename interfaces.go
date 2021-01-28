package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area()
	perim()
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	r := rect{width: 6, height: 3}
	c := circle{radius: 3}
	fmt.Println("Rectangle:", r)
	fmt.Println("Circle:", c)

	fmt.Println("\nArea of rectangle:", r.area())
	fmt.Println("Perimeter of rectangle:", r.perim())

	fmt.Println("\nArea of circle:", c.area())
	fmt.Println("Perimeter of circle:", c.perim())
}
