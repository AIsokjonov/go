package main

import "fmt"

func ByteSlices() {
	str := "Hello, World!"

	c := []byte(str)
	fmt.Printf("Original string: %s\n", str)
	fmt.Printf("Bytes: %v\n", c)
	c[7] = 'C'
	c[8] = '+'
	c[9] = '+'

	fmt.Printf("Bytes: %v\n", c)
}
