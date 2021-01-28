package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string) *Person {
	p := Person{name: name}
	p.age = 21
	return &p
}

func Structs() {
	fmt.Println(Person{"James", 24})
	fmt.Println(&Person{"Johnson", 29})
	fmt.Println(newPerson("Mark"))

	s := Person{"Robert", 53}
	fmt.Println(s.name)

	sp := &s
	sp.age = 49
	fmt.Println(sp)
}
