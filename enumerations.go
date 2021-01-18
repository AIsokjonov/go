package main
import "fmt"

type Gender int
const (
  UNKNOWN = iota
  FEMALE
  MALE
)

func main() {
  fmt.Println("Value:", MALE)
}
