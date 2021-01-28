package main

import "fmt"

func MapSlice() {
	// slice of maps
	items := make([]map[int]int, 5)

	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = i
	}
	fmt.Printf("Slice of maps: %v\n", items)
}
