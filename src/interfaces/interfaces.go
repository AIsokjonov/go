// detect and convert the type of an interface variable
package main
import (
	"fmt"
	"math"
)

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

func main() {
	var area Shaper
	sq := &Circle{5}
	area = sq

	switch t := area.(type) {
	case *Square:
		fmt.Printf("Type Square %T with value: %v\n", t, t)
	case *Circle:
		fmt.Printf("Type Circle %T with value: %v\n", t, t)

	default:
		fmt.Printf("Unexpected type: %T", t)
	}
}

func (s *Square) Area() float32 {
	return s.side * s.side
}

func (c *Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}