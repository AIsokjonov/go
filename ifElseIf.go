package main
import "fmt"

func IFELSEIF() {
	var age int = 21
	if age < 18 {
		fmt.Println("Underaged!")
	} else if age == 18 {
		fmt.Println("Normal")
	} else {
		fmt.Println("Adult")
	}
}
