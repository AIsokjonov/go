package main
import "fmt"

func reslice(n []int,l int) ([]int) {
	m := make([]int, len(n) * l)
	copy(m, n)
	return m
}

func main() {
	sl := []int{1,2,3}
	fmt.Println(reslice(sl, 5))
}