package main
import "fmt"

func main() {
	days := map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}

	var choice int
	fmt.Print("Enter day number: ")
	fmt.Scanf("%d\n", &choice)

	fmt.Println(days[choice])
}