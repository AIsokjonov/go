package main
import "fmt"

type List []int

func (l List) Len() int { return len(l) }

func (l *List) Append(val int) { *l = append(*l, val) }

func main() {
	// a bare value
	var lst List
	lst.Append(1)
	fmt.Printf("%v (len: %d)\n", lst, lst.Len())
	fmt.Printf("Type: %T\n\n", lst)

	// a pointer value
	plst := new(List)
	plst.Append(2)
	fmt.Printf("%v (len: %d)\n", plst, lst.Len())
	fmt.Printf("Type: %T\n", plst)
}
