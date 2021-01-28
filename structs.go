package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 21
	return &p
}

func main() {
	fmt.Println(person{"James", 24})
	fmt.Println(&person{"Johnson", 29})
	fmt.Println(newPerson("Mark"))

	s := person{"Robert", 53}
	fmt.Println(s.name)

	sp := &s
	sp.age = 49
	fmt.Println(sp)
}
