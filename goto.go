package main

import "fmt"

func main() {
	students := []string{"Bob", "Alex", "Mark", "Abdurauf", "Bahrom", "Abubakr"}
	i := 0
HERE:
	fmt.Println("Not found him")
	i++
	if students[i] == "Abdurauf" {
		fmt.Printf("\nFound %s\n", students[i])
		return
	}
	goto HERE
}
