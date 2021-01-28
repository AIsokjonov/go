package main

import (
	"fmt"
	"time"
)

func Reverse(s string) string {
	sl := []byte(s)
	rev := [100]byte{}
	j := 0
	for i := len(sl) - 1; i >= 0; i-- {
		rev[j] = sl[i]
		j++
	}
	str := string(rev[:len(sl)])
	return str
}

func InPlaceReverse(s string) string {
	sl := []byte(s)
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
	return string(sl)
}

func REVERSE() {
	str := "Googler"

	fmt.Printf("Original: %v\n", str)

	start := time.Now()
	fmt.Printf("Reversed(slice of bytes): %v\n", Reverse(str))
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("It took %v to reverse\n", delta)

	start = time.Now()
	fmt.Printf("\nReversed(in-place): %v\n", InPlaceReverse(str))
	end = time.Now()
	delta = end.Sub(start)
	fmt.Printf("It took %v to reverse\n", delta)
}
