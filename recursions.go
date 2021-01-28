package main

import "fmt"

// fibonacci
func FIBONACCI(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = FIBONACCI(n-1) + FIBONACCI(n-2)
	}
	return
}

// factorial
func Factorial(n uint64) (res uint64) {
	if n <= 1 {
		return 1
	}
	res = n * Factorial(n-1)
	return
}

func Main() {
	result := 0
	for i := 0; i <= 10; i++ {
		result = FIBONACCI(i)
		fmt.Printf("Fibonacci(%d) is: %d\n", i, result)
	}

	// factorial
	fmt.Printf("\nFactorial of %d is: %d\n", 5, Factorial(5))
}
