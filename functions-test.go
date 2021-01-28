package main

import "fmt"

type flt func(int) bool

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func filter(sl []int, f flt) (yes, no []int) {
	for _, v := range sl {
		if f(v) {
			yes = append(yes, v)
		} else {
			no = append(no, v)
		}
	}
	return
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Evens or odds test")
	even, odd := filter(slice, isEven)
	fmt.Println("evens:", even)
	fmt.Println("odds:", odd)
}
