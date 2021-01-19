package main
import (
  "fmt"
  "strconv"
)

func sum(a []int) int {
  s := 0
  for i := 0; i < len(a); i++ {
    s += a[i]
  }
  return s
}

func main() {
  // array
  var students = [5]string{"James","Bob","Mark","Mike","Robert"}
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

  // slices as parameters
  var arr2 = [5]int{2,4,5,6,7}
  // arr2[:] is a "slice"
  fmt.Printf("\nSlices as parameters: %d\n", sum(arr2[:]))

  // slice with make()
  slice2 := make([]int, 10)
  for i := range slice2 {
    slice2[i] = i * 5
  }
  fmt.Printf("Slice with make(): %d\n", slice2)
  fmt.Printf("The length of the slice: %d\n", len(slice2))
  fmt.Printf("The capacity of the slice: %d\n", cap(slice2))

  // another slice
  slice3 := make([]string, 10, 25)
  for i := 0; i < len(slice3); i++ {
    slice3[i] = strconv.Itoa(i * 3)
  }
  fmt.Printf("\nAnother slice: %s\n", slice3)
  fmt.Printf("The length of the slice: %d\n", len(slice3))
  fmt.Printf("The capacity of the slice: %d\n", cap(slice3))
}
