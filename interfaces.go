package main
import (
  "fmt"
  "math"
)

// interface
type geometry interface {
  area() float64
  perim() float64
}

// rectangle struct
type rect struct {
  width, height float64
}

// circle struct
type circle struct {
  radius float64
}

// rectangle methods
func(r rect) area() float64 {
  return r.width * r.height
}

func(r rect) perim() float64 {
  return 2 * r.width + 2 * r.height
}

// circle methods
func(c circle) area() float64 {
  return math.Pi * c.radius * c.radius
}

func(c circle) perim() float64 {
  return 2 * math.Pi * c.radius
}

func measure(g geometry) {
  fmt.Println("Value:", g)
  fmt.Println("Area:", g.area())
  fmt.Println("Perimeter:", g.perim())
}

func main() {
  r := rect{6, 4}
  c := circle{5}

  measure(r)
  measure(c)
}
