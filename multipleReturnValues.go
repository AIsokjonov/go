package main

import "fmt"

func Vals() (int, int) {
	return 3, 7
}

func MultipleReturnValues() {
	a, b := Vals()
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	_, c := Vals()
	fmt.Println("c:", c)
}
