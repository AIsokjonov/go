package main
import (
	"fmt"
	"strconv"
)

const SIZE = 4

type Stack struct {
	ix 	int
	data [SIZE]int
}

func (s *Stack) Push(n int) {
	if s.ix == SIZE {
		return
	}
	s.data[s.ix] = n
	s.ix++
}

func (s *Stack) Pop() int {
	if s.ix > 0 {
		s.ix--
		element := s.data[s.ix]
		s.data[s.ix] = 0
		return element
	}
	return -1
}

func (s Stack) String() string {
	str := ""
	for ix := 0; ix < s.ix; ix++ {
		str += "[" + strconv.Itoa(ix) + ":" + strconv.Itoa(s.data[ix]) + "]"
	}
	return str
}

func main() {
	st := new(Stack)
	st.Push(4)
	st.Push(1)
	st.Push(5)
	st.Push(7)
	fmt.Println(st.data)
}