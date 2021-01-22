package main
import "fmt"

func main() {
	students := map[int][]int{
		1: {1,2,3},
		2: {4,5,6},
	}
	fmt.Printf("Map with slices: %v\n", students)
}