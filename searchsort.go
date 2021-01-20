package main
import (
  "fmt"
  "sort"
  "time"
)

func main() {
  unsorted := []int{134,12,-14,14,0,124,1,124143}

  start := time.Now()
  fmt.Printf("Original: %v\n", unsorted)
  sort.Ints(unsorted)
  fmt.Printf("Original: %v\n", unsorted)
  end := time.Now()
  delta := end.Sub(start)
  fmt.Printf("Sorted: %v\n", sort.IntsAreSorted(unsorted))
  fmt.Printf("It took %v to complete this task\n", delta)

  // sorting strings
  students := []string{"James","Bob","Mike"}
  fmt.Println("\nSorting strings")
  fmt.Printf("Original: %v\n", students)
  sort.Strings(students)
  fmt.Printf("Sorted: %v\n", students)
}
