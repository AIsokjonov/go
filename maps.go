package main
import "fmt"

func main() {
	// map declaration
	students := map[int]string{}
	
	// maps with integer slices as values
	processes := map[int][]int{}

	// maps with string slices as values
	pros := map[int][]string{}


	// initialization
	students = map[int]string{
		1: "Abdurauf",
		2: "Bob",
		3: "Jasur",
	}

	processes = map[int][]int{
		1: {124,542,53465},
		2: {235,524,213},
		3: {756, 875,7452},
	}

	pros = map[int][]string{
		1: {"mysqld", "systemctl","code"},
		2: {"xcode", "terminal","chrome"},
		3: {"code-helper","docker"},
	}

	// check the existence of key-value item
	_, isPresent := students[5]
	
	fmt.Println(students)
	fmt.Println(processes)
	fmt.Println(pros)
	fmt.Printf("Existence: %v\n", isPresent)
}