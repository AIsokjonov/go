package main
import "fmt"

type Shaper interface {
	Area() float32
	Perim() float32
}

type Square struct {
	side float32
}

func (s *Square) Area() float32 {
	return s.side * s.side
}

func (s *Square) Perim() float32 {
	return 2 * (s.side + s.side)
}

func main() {
	sq1 := &Square{5}
	areaIntf := Shaper(sq1)
	fmt.Println(areaIntf.Area())
	fmt.Println(areaIntf.Perim())
}