package main
import "fmt"

// variables
var i, j, k int
var sch, ch byte
var f, salary float32

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

	// mixed variable declaration
	// mixed variable types can be declared using type inference
	var a, b, c = 3, 5.0, "A Prayer Before Dawn"
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
	fmt.Printf("a is type of %T\n", a)
	fmt.Printf("b is type of %T\n", b)
	fmt.Printf("c is type of %T\n", c)

	// constants
	const LENGTH int = 10
	const WIDTH int = 5
	var area int

	area = LENGTH * WIDTH
	fmt.Println("value of area: ", area)
}