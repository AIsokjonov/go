package main
import "fmt"

func main() {
  var i1 = 5
  fmt.Printf("An integer: %d, and its location in memory:%p\n", i1, &i1)
  var intP *int
  intP = &i1
  fmt.Printf("The value at memory location %p is %d\n", intP, *intP)

  // values changes
  // when we assign a new value to the pointer

  name := "Abdurauf"
  var p *string = &name
  fmt.Printf("\nMemory address: %p\nValue: %s\n", p, *p)
  *p = "James"
  fmt.Printf("\nMemory address: %p\nChanged value: %s\n", p, *p)
}
