package main
import "fmt"

func main() {
	// local variable definition
	var a int = 10

	if (a < 20) {
		fmt.Println("a is less than 20")
	}
	fmt.Printf("value of a is: %d\n", a)
}