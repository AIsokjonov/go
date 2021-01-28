package main

import "fmt"

func copyAndAppendSlice() {
	// copy a slice into another slice
	sl := []int{1, 2, 3}
	sl_copy := make([]int, 10)
	n := copy(sl_copy, sl)

	fmt.Printf("Original slice: %v\n", sl)
	fmt.Printf("Copy slice: %v\n", sl_copy)
	fmt.Printf("Number of copied elements: %v\n", n)

	// append a slice into another slice
	sl2 := []int{2345, 234, -234}
	fmt.Println("\nAppend slice")
	fmt.Printf("Original slice: %v\n", sl2)

	sl2 = append(sl2, 15, 54)
	fmt.Printf("Appended slice: %v\n", sl2)
}
