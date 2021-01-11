package main
import "fmt"

func BMW() string {
	fav := "BMW"
	return fav
}

func McLaren() string {
	fav := "McLaren"
	return fav
}

func main() {
	favoriteCar := "McLaren"

	if favoriteCar == "BMW" {
		fmt.Println("My favorite car:", BMW())
	} else if favoriteCar == "McLaren" {
		fmt.Println("My favorite car:", McLaren())
	} else {
		fmt.Println("I like Mustang")
	}
}
