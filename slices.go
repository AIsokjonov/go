package main
import "fmt"

func main() {
  // array
  var students [5]string
  students[0] = "James"
  students[1] = "Bob"
  students[2] = "Mike"
  students[3] = "George"
  students[4] = "Robert"

  fmt.Printf("array: %s\ntype: %T\n", students, students)

  // slice
  var projectMembers []string = students[1:3]
  fmt.Printf("slice: %s\ntype: %T\n", projectMembers, projectMembers)

  // expand the capacity of a slice
  var arr1 [6]int
  var slice1 []int = arr1[2:5]

  for i := 0; i < len(arr1); i++ {
    arr1[i] = i
  }

  for i := 0; i < len(slice1); i++ {
    fmt.Printf("Slice at %d is %d\n", i, slice1[i])
  }
  fmt.Printf("The length of arr1 is %d\n", len(arr1))
  fmt.Printf("The length of slice1 is %d\n", len(slice1))
  fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

  slice1 = slice1[0:4]
  for i := 0; i < len(slice1); i++ {
    fmt.Printf("Slice at %d is %d\n", i, slice1[i])
  }
  fmt.Printf("The length of slice1 is %d\n", len(slice1))
  fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
}
