package main
import "fmt"

func main() {
  var grade string = "B"
  var marks int = 90

  switch (marks) {
    case 90: grade = "A"
    case 80: grade = "B"
    case 50, 60, 70: grade = "C"
    default: grade = "D"
  }

  switch {
    case grade == "A":
      fmt.Printf("Excellent!")
    case grade == "B", grade == "C":
      fmt.Printf("Well done!")
    case grade == "D":
      fmt.Printf("You passed!")
    case grade == "F":
      fmt.Printf("Better try again!")
    default:
      fmt.Printf("Invalid grade!")
  }

  fmt.Printf("Your grade is %s\n", grade)
}
