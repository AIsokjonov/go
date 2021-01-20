package main
import (
	"time"
	"fmt"
)

func main() {
	value := 0
	screen := [1080][1920]int{}

	start := time.Now()
	for row := range screen {
		for column := range screen[row] {
			screen[row][column] = value
			value += 1
		}
	}

	for row := range screen {
		for column := range screen[row] {
			fmt.Printf("row: %d, column: %d\n", row, column)
		}
	}

	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("It took %v to complete this task\n", delta)
}