package main

import "fmt"

const (
	WIDTH  = 1920
	HEIGHT = 1080
)

type pixel int // aliasing
var screen [WIDTH][HEIGHT]pixel

func main() {
	// multidimensional arrays
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = 0
		}
	}
	//fmt.Printf("Screen 2-D array: %d\n", screen)

	// multidimensional slices
	students := [][]string{}

	BBA := []string{"James", "Bob", "Robert"}
	IT := []string{"Mike", "Charles", "Edmon"}

	students = append(students, BBA)
	students = append(students, IT)

	fmt.Printf("Students: %s\n", students)
	fmt.Printf("IT dept. students: %s\n", students[1])
	fmt.Printf("BBA dept. students: %s\n", students[0])
}
