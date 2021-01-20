package main
import (
	"fmt"
	"time"
	"sort"
)

func main() {
	// searching strings
	sl := []string{"James","Bob","Robert","Mark","Alex","Mike"}
	x := "Robert"

	start := time.Now()
	
	fmt.Printf("Original: %v\n", sl)
	sort.Strings(sl)
	fmt.Printf("Sorted: %v\n", sl)
	
	fmt.Printf("\nSearch item: %v\n", x)
	fmt.Printf("Search Result: %v\n", sort.SearchStrings(sl, x))

	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("\nIt took %v to complete this task\n", delta)

	// append
	fmt.Println("\nAppend")
	sl2 := []int{1,2,3,4,5}
	sl3 := []int{6,7,8,9,10}
	fmt.Printf("slice #2: %v\n", sl2)
	fmt.Printf("slice #3: %v\n", sl3)

	sl2 = append(sl2, sl3...)

	fmt.Printf("appended: %v\n", sl2)

	// append - delete item at index i
	sl2 = append(sl2[:2], sl2[3:]...)
	fmt.Printf("delete item at index 2: %v\n", sl2)

	// append - extends the slice length
	fmt.Printf("\nBefore extend: %v\n", sl2)
	sl2 = append(sl2[:], make([]int, 5)...)
	fmt.Printf("After extend: %v\n", sl2)

	// insert a new item at index i
	fmt.Printf("\nBefore insert: %v\n", sl2)
	sl2 = append(sl2[:len(sl2)], append([]int{6100}, sl2[len(sl2):]...)...)
	fmt.Printf("After insert: %v\n", sl2)

	// insert a new slice at index i
	sl4 := []int{1554, 1997}
	fmt.Printf("\nslice into a slice(before): %v\n", sl2)
	sl2 = append(sl2[:len(sl2)], append(sl4, sl2[len(sl2):]...)...)
	fmt.Printf("slice into a slice(after): %v\n", sl2)
}