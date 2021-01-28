package main

import "fmt"

func NestedIfs() {
	var a int = 100
	var b int = 200

	if a == 100 {
		if b == 200 {
			fmt.Println("Value of a is 100 and b is 200")
		}
	}

	fmt.Printf("Exact value of a: %d\n", a)
	fmt.Printf("Exact value of b: %d\n", b)
}
