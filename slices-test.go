package main

import "fmt"

func Slice_fib(n int) []int {
	slice := make([]int, n)
	slice[0], slice[1] = 1, 1
	for i := 2; i < n; i++ {
		slice[i] = slice[i-1] + slice[i-2]
	}
	return slice
}

func SliceMapTest() {
	fmt.Println(Slice_fib(10))
}
