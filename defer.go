package main
import (
  "fmt"
  "log"
  "io"
)

func func1() {
  fmt.Println("in Function #1 at the top!")
  defer func2()
  fmt.Println("in Function #2 at the bottom")
}

func func2() {
  fmt.Println("Function #2: deferred until the end of the calling function!")
}

// tracing with defer
func trace(s string) {
  fmt.Println("Entering:", s)
}

func untrace(s string) {
  fmt.Println("Leaving:", s)
}

func a() {
  trace("a")
  defer untrace("a")
  fmt.Println("in a")
}

func b() {
  trace("b")
  defer untrace("b")
  fmt.Println("in b")
  a()
}

// loggging parameters with defer
func logging(s string) (n int, err error) {
  defer func() {
    log.Printf("logging(%q) = %d, %v", s, n, err)
  }()

  return 7, io.EOF
}

func main() {
  func1()

  // tracing with defer
  fmt.Println("\nTracing with defer")
  b()

  // logging parameters with defer
  fmt.Println("\nLogging parameters with defer")
  logging("Go")
}
