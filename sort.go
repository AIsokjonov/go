package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	unsorted := []int{134, 12, -14, 14, 0, 124, 1, 124143}

	start := time.Now()
	fmt.Printf("Original: %v\n", unsorted)
	sort.Ints(unsorted)
	fmt.Printf("Original: %v\n", unsorted)
	fmt.Printf("Sorted: %v\n", sort.IntsAreSorted(unsorted))

	// sorting strings
	students := []string{"James", "Bob", "Mike"}
	fmt.Println("\nSorting strings")
	fmt.Printf("Original: %v\n", students)
	sort.Strings(students)
	fmt.Printf("Sorted: %v\n", students)

	// sorting float64
	f := []float64{243.1, 0.31, -13.314, 23.12, 552.1, 65.1}
	fmt.Printf("\nOriginal: %v\n", f)
	sort.Float64s(f)
	fmt.Printf("Sorted: %v\n", f)

	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("\nIt took %v to complete this task\n", delta)
}
