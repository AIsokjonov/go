package main
import (
	"fmt"
	"time"
	"sort"
)

func mySearch(sl []int, n int) int {
	var res int
	for i := 0; i < len(sl); i++ {
		if sl[i] == n {
			res = i
		}
	}
	return res
}

func main() {
	// searching int
	sl := []int{24,1,4,0,234,13,5,6}
	x := 5

	fmt.Printf("Original: %v\n", sl)
	sort.Ints(sl)
	fmt.Printf("Sorted: %v\n", sl)

	start := time.Now()
	fmt.Printf("\nSearch integer: %v\n", sort.SearchInts(sl, x))
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("It took %v to complete this task\n", delta)

	start = time.Now()
	fmt.Printf("\nMy Search: %v\n", mySearch(sl, x))
	end = time.Now()
	delta = end.Sub(start)
	fmt.Printf("It took %v to complete this task\n", delta)
}