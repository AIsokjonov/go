package main
import "fmt"

type Simple struct {
	num int
}

type RSimple struct {
	n int
}

type Simpler interface {
	Get() int
	Set(int)
}

func (s *Simple) Get() int {
	return s.num
}

func (s *Simple) Set(n int) {
	s.num = n
}

func (rs *RSimple) Get() int {
	return rs.n
}

func (rs *RSimple) Set(num int) {
	rs.n = num
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
	var r RSimple
	fmt.Println(fI(&r))
}