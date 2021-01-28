package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubblesort(sl []int) []int {
	for i := 1; i < len(sl); i++ {
		for j := 0; j < len(sl)-i; j++ {
			if sl[j] > sl[j+1] {
				sl[j], sl[j+1] = sl[j+1], sl[j]
			}
		}
	}
	return sl
}

func mySort(sl []int) []int {
	tmp := 0
	for i := 0; i < len(sl); i++ {
		for j := 1; j < len(sl); j++ {
			if sl[j-1] > sl[j] {
				tmp = sl[j-1]
				sl[j-1] = sl[j]
				sl[j] = tmp
			}
		}
	}
	return sl
}

func BubbleSort() {
	sl := rand.Perm(1000000)

	start := time.Now()
	bubblesort(sl)
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("It took %v to sort\n", delta)
}
