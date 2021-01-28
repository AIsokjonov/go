package main
import "fmt"

type flt func(int) bool

func IsEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func Filter(sl []int, f flt) (yes, no []int) {
	for _, v := range sl {
		if f(v) {
			yes = append(yes, v)
		} else {
			no = append(no, v)
		}
	}
	return
}

func FunctionTest() {
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Evens or odds test")
	even, odd := Filter(slice, IsEven)
	fmt.Println("evens:", even)
	fmt.Println("odds:", odd)
}
