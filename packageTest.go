package main
import "fmt"
import "myPackage/hello.go"

func main() {
	fmt.Println(hello.Greet("James"))
}