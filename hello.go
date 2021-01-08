package main
import "fmt"

func main() {
	fmt.Println("Hello, World")

	// static type declaration
	var x float64
	x = 20.0
	fmt.Println("x: ",x)
	fmt.Printf("x is of type %T\n", x)

	// dynamic type declaration
	y := 42
	fmt.Println("y: ", y)
	fmt.Printf("y is of Type %T\n", y)
}

// variables
var i, j, k int
var c, ch byte
var f, salary float32