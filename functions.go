package main
import "fmt"

func sum(a, b, c int) int {
  return a + b + c
}

func arithmetic(a, b int) (int, int, int) {
  n1 := a + b
  n2 := a - b
  n3 := a * b
  return n1, n2, n3
}

func greet(name string) string {
  return "Hello " + name
}

func main() {
  // calling function inside function
  fmt.Printf("Value: %d\n", sum(arithmetic(20, 10)))

  fmt.Print("Enter you name: ")
  var name string
  fmt.Scanf("%s", &name)

  // assign a function to a variable
  sayHello := greet(name)
  fmt.Println(sayHello)

  // blank identifier example
  fmt.Println("\nArithmetic operations")
  a1, _, a3 := arithmetic(3, 5) // the second variable is the blank identifier
  fmt.Printf("Addition: %d\nMultiplication: %d\n", a1, a3)
}
