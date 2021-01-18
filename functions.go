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

func Multiply(a, b int, reply *int) {
  *reply = a * b
}

func myPrint(a ...int) {
  if len(a) == 0 {
    return
  }

  for _, v := range a {
    fmt.Printf("The number is: %d\n", v)
  }
}

func main() {
  // calling function inside function
  fmt.Printf("Value: %d\n", sum(arithmetic(20, 10)))

  //fmt.Print("Enter you name: ")
  //var name string
  //fmt.Scanf("%s", &name)

  // assign a function to a variable
  //sayHello := greet(name)
  //fmt.Println(sayHello)

  // blank identifier example
  fmt.Println("\nArithmetic operations")
  a1, _, a3 := arithmetic(3, 5) // the second variable is the blank identifier
  fmt.Printf("Addition: %d\nMultiplication: %d\n", a1, a3)

  // changing an outside variable
  n := 0 // int variable
  reply := &n // saves the memory address of n == pointer
  Multiply(10, 5, reply)
  fmt.Println("\nMultiply:", *reply)
  fmt.Println("Multiply:", n)

  // variadic functions
  myPrint(1,23, 134, 41, 42, 0, -23)
}
