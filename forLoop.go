package main

import "fmt"

func Forloop() {
	// counter-controlled iteration
	for i := 0; i < 3; i++ {
		fmt.Println("Hello, World!")
	}

	// nested counter-controlled iteration
	for i := 0; i < 5; i++ {
		for j := 5 - i; j >= 0; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// condition controlled iteration
	var i int = 0
	for i < 5 {
		fmt.Printf("This is the %d iteration\n", i+1)
		i++
	}
}
