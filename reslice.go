package main
import "fmt"

func main() {
	sl := make([]int, 0, 10)

	fmt.Printf("Length: %d\nCapacity: %d\n", len(sl), cap(sl))

	for i := 0; i < cap(sl); i++ {
		sl = sl[0:i+1]
		fmt.Printf("Length: %d\tCapacity: %d\n", len(sl), cap(sl))
	}

	// iterate over each item
	fmt.Printf("\nIterating over slice\n")
	for i := 0; i < len(sl); i++ {
		fmt.Printf("item #%d: %d\n", i, sl[i])
	}
}