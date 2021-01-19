package main
import (
  "fmt"
  "strconv"
  "time"
)

// pass pointer to array
func change(a *[4]int) (*[4]int) {
  for i := range a {
    a[i] = i + 1
  }
  return a
}

func sum(a *[1000]int) (s int) {
  for i := range a {
    s += i
  }
  return
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

  var bigArr [1000]int
  for i := 0; i < 1000; i++ {
    bigArr[i] = i + 1
  }
  start := time.Now()
  fmt.Printf("\nSum: %d\n", sum(&bigArr))
  end := time.Now()
  delta := end.Sub(start)
  fmt.Printf("Execution time: %s\n", delta)

}
