package main
import "fmt"

// variadic function accepts any number of arguments
// variadic function example
func sum(nums ...int) {
	fmt.Print(nums, " => ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("Total:", total)
}

func main() {
	sum(1, 3, 0)
	sum(13,4312,235,-23434,234,542,134352)

	// if you already have multiple arguments in a slice
	// apply them to a variadic function using (slice...) like this
	nums := []int{1,2,3,4,5}
	sum(nums...)
}