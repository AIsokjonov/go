package main
import (
	"container/list"
	"fmt"
)

func doubleList(n int) (*list.List) {
	l := list.New()
	for i := 1; i <= n; i++ {
		l.PushBack(i)
	}
	return l
}

func main() {
	myList := doubleList(5)
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}