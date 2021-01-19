package main
import (
  "fmt"
  "time"
)

func Calculation() {
  for i := 0; i < 100; i++ {
    fmt.Printf("Hello, World %d times\n", i)
  }
}

func main() {
  start := time.Now()
  Calculation()
  end := time.Now()
  delta := end.Sub(start)
  fmt.Printf("\nCalculation took this amount of time: %s\n", delta)
}
