package main

import "fmt"

// fibonacci
func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

// factorial
func factorial(n uint64) (res uint64) {
	if n <= 1 {
		return 1
	}
	res = n * factorial(n-1)
	return
}

func main() {
	result := 0
	for i := 0; i <= 10; i++ {
		result = fibonacci(i)
		fmt.Printf("Fibonacci(%d) is: %d\n", i, result)
	}

	// factorial
	fmt.Printf("\nFactorial of %d is: %d\n", 5, factorial(5))
}
