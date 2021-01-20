package main
import (
	"fmt"
	"time"
	"sort"
)

func main() {
	// searching int
	sl := []float64{1.2, -0.73, 4.42, 5.42, 23.53}
	x := 3.5
	fmt.Printf("Original: %v\n", sl)

	start := time.Now()

	fmt.Printf("Search item: %v\n", x)
	sort.Float64s(sl)
	fmt.Printf("Sorted slice: %v\n", sl)
	fmt.Printf("Searching float64: %v\n", sort.SearchFloat64s(sl, x))

	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("It took %v to complete this task\n", delta)
}