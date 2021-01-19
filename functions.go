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

// functions as values and parameters
func Add(a, b int) {
  fmt.Printf("The sum of %d and %d is: %d\n", a, b, a + b)
}

func callback(y int, f func(int, int)) {
  f(y, 2)
}

// functions used as a filter
type flt func(int) bool

func isOdd(n int) bool {
  if n % 2 == 0 {
    return false
  }
  return true
}

func isEven(n int) bool {
  if n % 2 == 0 {
    return true
  }
  return false
}

func filter(sl[] int, f flt)[] int {
  var res[] int
  for _, v := range sl {
    if f(v) {
      res = append(res, v)
    }
  }
  return res
}

func Add2() (func(b int) int) {
  return func(b int) int {
    return b + 2
  }
}

func Adder(a int) (func(b int) int) {
  return func(b int) int {
    return a + b
  }
}

// non-recursive fibonacci function
func fib() func() int {
  a, b := 1, 1
  return func() int {
    a, b = b, a + b
    return b
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
  //fmt.Println(sayHello)``

  // blank identifier example
  fmt.Println("\nArithmetic operations")
  a1, _, a3 := arithmetic(3, 5)
  fmt.Printf("Addition: %d\nMultiplication: %d\n", a1, a3)

  // changing an outside variable
  n := 0
  reply := &n
  Multiply(10, 5, reply)
  fmt.Println("\nMultiply:", *reply)
  fmt.Println("Multiply:", n)

  // variadic functions
  myPrint(1,23, 134, 41, 42, 0, -23)

  // functions as parameters
  callback(1, Add)

  // functions as filters
  fmt.Println("\nFunctions as filters")
  slice := [] int{1,2,3,4,5,7}
  fmt.Println("slice =", slice)
  odd := filter(slice, isOdd)
  fmt.Println("Odd elements of slice are:", odd)
  even := filter(slice, isEven)
  fmt.Println("Even elements of slice are:", even)

  // functions as return values
  p2 := Add2()
  fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))
  TwoAdder := Adder(2)
  fmt.Printf("The result is: %v\n", TwoAdder(3))

  // non-recursive fibonacci
  f := fib()
  for i := 0; i <= 9; i++ {
    fmt.Println(i + 2, f())
  }
}