package methods
import "fmt"

type S struct {
	a int
}

type SType S
type SAlias = S
type IntType int
type IntAlias = int

func (recv S) print() {
	fmt.Printf("%t: %[1]v\n", recv)
}

func (recv SType) print() {
	fmt.Printf("%t: %[1]v\n", recv)
}

func (recv IntType) print() {
	fmt.Printf("%t: %[1]v\n", recv)
}
