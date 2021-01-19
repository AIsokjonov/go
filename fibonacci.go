package main
import "fmt"

var fib [10]int

func fibs() [10]int {
  fib[0] = 1
  fib[1] = 1

  for i := 2; i < 10; i++ {
    fib[i] = fib[i-1] + fib[i-2]
  }
  return fib
}

func main() {
  fmt.Println(fibs())
}
