package main
import "fmt"

func main() {
  // break statement
  fmt.Println("Break statement: ")
  for i := 0; i < 3; i++ {
    for j := 0; j < 10; j++ {
      if j > 4 {
        break
      }
      fmt.Printf("%d", j)
    }
    fmt.Print(" ")
  }

  // continue statement
  fmt.Println("\n\nContinue statement: ")
  students := []string{"Abubakr", "Bahrom", "Abdulaziz", "Abdurauf", "Durbek", "Sarvar", "Bekzod"}

  for i := range students {
    if students[i] == "Abdurauf" {
      fmt.Printf("\nFound him %s\n", students[i])
      continue
    }
    fmt.Printf("%s", students[i])
    fmt.Print(" ")
  }
}
