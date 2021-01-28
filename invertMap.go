package main

import "fmt"

func InvertMap() {
	barVal := map[string]int{
		"alpha":   34,
		"bravo":   56,
		"charlie": 23,
		"delta":   87,
		"echo":    56,
		"foxtrot": 12,
		"golf":    34,
		"hotel":   16,
		"indio":   87,
		"juliet":  65,
		"kilo":    43,
		"lima":    98,
	}

	invMap := make(map[int]string, len(barVal))

	fmt.Printf("Before inversion:\n")
	for k, v := range barVal {
		fmt.Printf("key: %v\tvalue: %v\n", k, v)
	}

	for k, v := range barVal {
		invMap[v] = k
	}

	fmt.Printf("\nAfter inversion:\n")
	for k, v := range invMap {
		fmt.Printf("key: %v\tvalue: %v\n", k, v)
	}
}
