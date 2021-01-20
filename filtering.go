package main
import "fmt"

func isEven(n int) bool {
  if n % 2 == 0 {
    return true
  }
  return false
}

func filter(s []int, fn func(int) bool) []int {
  res := []int{}
  for _, i := range s {
    if fn(i) {
      res = append(res, i)
    }
  }
  return res
}

func main() {
  sl := make([]int, 10, 10)
  for i := 0; i < len(sl); i++ {
    sl[i] = i * 3
  }

  fmt.Printf("Original: %v\n" ,sl)
  fmt.Printf("Filtered: %v\n", filter(sl, isEven))
}
