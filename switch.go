package main
import "fmt"

func Season(month int) string {
  switch month {
    case 12, 1, 2:
      fmt.Println("Winter")
    case 3, 4, 5:
      fmt.Println("Spring")
    case 6,7,8:
      fmt.Println("Summer")
    case 9,10,11:
      fmt.Println("Autumn")

    default:
      fmt.Println("Unknown season!")
  }
  return ""
}

func main() {
  var monthNum int
  fmt.Print("Enter month Number: ")
  fmt.Scanf("%d", &monthNum)
  Season(monthNum)
}
