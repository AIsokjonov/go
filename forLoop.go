package main
import "fmt"

func main() {
  // array of string values
  students := []string{"James", "Bob", "Robert"}

  // iterate through each item in students array
  for i := 0; i < 3; i++ {
    fmt.Println("Student #", i+1, ": ", students[i])
  }

  fmt.Println("Iterating from the end till the beginning")
  // iterate from the end till the beginning
  for k := len(students) - 1; k >= 0; k-- {
    fmt.Println("Student #", k+1, ":", students[k])
  }

  // array of integers
  nums := []int{12,312,-2, 324,0,123}

  // iterate through each item in nums array
  for j := 0; j < len(nums); j++ {
    fmt.Println("Num #",j+1,": ",nums[j])
  }
}
