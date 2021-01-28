package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	Area()
	Perim()
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perim() float64 {
	return 2*r.width + 2*r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perim() float64 {
	return 2 * math.Pi * c.radius
}

func Interfaces() {
	r := Rect{width: 6, height: 3}
	c := Circle{radius: 3}
	fmt.Println("Rectangle:", r)
	fmt.Println("Circle:", c)

	fmt.Println("\nArea of rectangle:", r.Area())
	fmt.Println("Perimeter of rectangle:", r.Perim())

	fmt.Println("\nArea of circle:", c.Area())
	fmt.Println("Perimeter of circle:", c.Perim())
}
