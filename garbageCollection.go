package main

import "fmt"

var arr []byte = []byte{'a', 'b', 'a', 'a', 'a', 'c', 'd', 'e', 'f', 'g'}

func main() {
	sl := make([]byte, len(arr))
	i := 0
	tmp := byte(0)

	fmt.Printf("Before: %v\n", arr)
	for _, v := range arr {
		if v != tmp {
			sl[i] = v
			i++
		}
		tmp = v
	}
	fmt.Printf("After: %v\n", sl)
}
