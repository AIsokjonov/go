package main

import (
	"fmt"
	"log"
	"os"
)

func Errors() {

	// print error with fmt
	_, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("error occurred:", err)
	}

	// print error with log
	if err != nil {
		log.Println("error occurred:", err)
	}

	// print error with Fatalln
	if err != nil {
		log.Fatalln("Error occurred:", err)
	}

	// print error with panic
	if err != nil {
		panic(err)
	}
}
