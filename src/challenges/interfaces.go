package main
import "fmt"

type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	n int
}

type RSimple struct {
	n2 int
}

func (s *Simple) Get() int {
	return s.n
}

func (s *Simple) Set(num int) {
	s.n = num
}

func (rs *RSimple) Get() int {
	return rs.n2
}

func (rs *RSimple) Set(num2 int) {
	rs.n2 = num2
}

func fI(it Simpler) int {
	switch it.(type) {
	case *Simple:
		it.Set(5)
		return it.Get()
	case *RSimple:
		it.Set(50)
		return it.Get()

	default:
		return 99
	}
	return 0
}

func main() {
	var s Simple
	fmt.Println(fI(&s))

	var rs RSimple
	fmt.Println(fI(&rs))
}