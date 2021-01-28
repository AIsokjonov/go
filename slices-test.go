package main

import "fmt"

func slice_fib(n int) []int {
	slice := make([]int, n)
	slice[0], slice[1] = 1, 1
	for i := 2; i < n; i++ {
		slice[i] = slice[i-1] + slice[i-2]
	}
	return slice
}

func main() {
	fmt.Println(slice_fib(10))
}
