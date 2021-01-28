package main

import "fmt"

func Enlarge(n []int, l int) []int {
	ns := make([]int, len(n)*l)
	copy(ns, n)
	return ns
}

func Insert(a []string, b []string, i int) []string {
	res := make([]string, len(a)+len(b))
	at := copy(res, a[:i])
	at += copy(res[at:], b)
	copy(res[at:], a[i:])
	return res
}

func MagnifySlice() {
	sl := []int{1, 2, 3, 4, 5}
	fmt.Println(Enlarge(sl, 4))

	// insert a slice into another slice
	sl2 := []string{"james", "bob"}
	sl3 := []string{"robert", "mike"}
	fmt.Println(Insert(sl2, sl3, 1))
}
