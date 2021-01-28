package main

import (
	"fmt"
	"log"
	"runtime"
)

func f() {
	for i := 0; i < 5; i++ {
		g := func(i int) { fmt.Printf("%d", i) }
		g(i)
		fmt.Printf(" - g is of type %T and has value %v\n", g, &g)
	}
}

func main() {
	f()

	func() {
		sum := 0.0
		for i := 1; i <= 5; i++ {
			sum += float64(i)
		}
		fmt.Printf("\nSum: %d\n", int(sum))
	}()

	// debugging with runtime
	fmt.Println("\nDebuging with runtime")
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()

	a := 2 + 4
	where()
	b := 2 * 7

	fmt.Println("value `a`:", a)
	fmt.Println("value `b`:", b)

	// debugging with log
	fmt.Println("\nDebugging with log")
	where2 := log.Print
	where2()
	c := 9 / 3
	where2()

	fmt.Println("value `c`:", c)
}
