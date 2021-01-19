package main
import "fmt"

func f() {
	for i := 0; i < 1e6; i++ {
		g := func(i int) {fmt.Printf("%d", i)}
		g(i)
		fmt.Printf(" - g is of type %T and has value %v\n", g, &g)
	}
}

func main() {
	f()

	func() {
		sum := 0.0
		for i := 1; i <= 1e6; i++ {
			sum += float64(i)
		}
	}()
}