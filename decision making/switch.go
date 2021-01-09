package main
import (
  "fmt"
  "time"
)

func main() {
  i := 3
  fmt.Print("Write ", i, " as ")
  switch i {
    case 1:
      fmt.Println("one")
    case 2:
      fmt.Println("two")
    case 3:
      fmt.Println("three")
  }

  switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
      fmt.Println("It's the weekend and today is ", time.Now().Weekday())
    default:
      fmt.Println("It's a weekday and today is ", time.Now().Weekday())
  }

  t := time.Now()
  switch {
    case t.Hour() < 12:
      fmt.Println("It's before noon and the time is ", time.Now().Hour())
    default:
      fmt.Println("It's after noon and the time is ", time.Now().Hour())
  }

  whatAmI := func(i interface{}) {
    switch t := i.(type) {
      case bool:
        fmt.Println("I'm a boolean")
      case int:
        fmt.Println("I'm an integer")
      default:
        fmt.Printf("Don't know %T\n", t)
    }
  }
  whatAmI(23)
}
