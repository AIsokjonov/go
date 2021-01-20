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
}