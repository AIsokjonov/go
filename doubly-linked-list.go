package main

import (
	"container/list"
	"fmt"
)

func dll(n int) *list.List {
	l := list.New()
	for i := 0; i < n; i++ {
		l.PushBack(i)
	}
	return l
}

func main() {
	myList := dll(5)
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
