package main

import "fmt"

type RECT struct {
	width, height int
}

func (r *RECT) Area() int {
	return r.width * r.height
}

func (r RECT) Perim() int {
	return 2*r.width + 2*r.height
}

func Methods() {
	r := Rect{width: 8, height: 3}
	fmt.Println("Area:", r.Area())
	fmt.Println("Perimeter:", r.Perim())

	rp := &r
	fmt.Println("\nPointer Area:", rp.Area())
	fmt.Println("Pointer Perimeter:", rp.Perim())
}