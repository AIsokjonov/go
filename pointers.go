package main
import "fmt"

type User struct {
	Name	string
	Pets	[]string
}

func (u *User) newPet() {
	u.Pets = append(u.Pets, "Lucy")
	fmt.Println("NewPet():", u)
}

func main() {
	u := User{Name: "Anna", Pets: []string{"Bailey"}}
	p := &u
	p.newPet()
	fmt.Println("Main method():", u)
	fmt.Println("newPet method():", p)
}