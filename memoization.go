package main

import "fmt"

func main() {
	result := 1
	a := 1
	b := 1

	for i := 0; i <= 10; i++ {
		if i > 1 {
			result = a + b
		}
		fmt.Printf("fibonacci(%d) is: %d:\n", i, result)

		// this method is memoization
		a = b
		b = result
	}
}
