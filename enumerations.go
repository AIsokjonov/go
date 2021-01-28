package main

import "fmt"

type Gender int

const (
	UNKNOWN = iota
	FEMALE
	MALE
)

func Enumerations() {
	fmt.Println("Value:", MALE)
}
