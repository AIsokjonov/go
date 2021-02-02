package main
import "fmt"

const SIZE = 5

type Stack struct {
	index	int
	data	[SIZE]int
}

func (s *Stack) Push(n int) {
	if s.index == SIZE {
		return
	}
	s.data[s.index] = n
	s.index++
}

func (s *Stack) Pop() int {
	if s.index > 0 {
		s.index--
		element := s.data[s.index]
		s.data[s.index] = 0
		return element
	}
	return -1
}

func main() {
	s := &Stack{}
	s.data = [5]int{23,52,12,76,2}
	fmt.Println(s.data)

	fmt.Println(s.Pop())
}