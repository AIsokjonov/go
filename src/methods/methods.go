package main
import (
	"fmt"
	"math"
)

// methods with struct types
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// pointer receiver
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// methods with non-struct types
type MyFloat float64

func (m MyFloat) Sum() float64 {
	if m < 0 {
		return float64(-m)
	}
	return float64(m)
}

func main() {
	// methods for struct types
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	// methods for non-struct types
	f := MyFloat(-3.02345)
	fmt.Println(f.Sum())

	// pointer receiver
	v.Scale(7)
	fmt.Println(v.Abs())
}
