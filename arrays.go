package main
import (
  "fmt"
  "strconv"
)

// pass pointer to array
func change(a *[4]int) (*[4]int) {
  for i := range a {
    a[i] = i + 1
  }
  return a
}

func main() {
  var students [5]string
  for i := 0; i < len(students); i++ {
    students[i] = strconv.Itoa(i)
  }
  fmt.Printf("Array: %s\n", students)

  var nums [4]int
  for i := range nums {
    nums[i] = i * 2
  }
  fmt.Printf("Nums: %d\n", nums)

  // pass pointer
  fmt.Printf("original(before): %d\n", nums)
  fmt.Println(change(&nums))
  fmt.Printf("Original(after): %d\n", nums)
}
