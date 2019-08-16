package stack


type Stack struct {
 stack []int
}

func New() *Stack {
 return &Stack{}
}

func (s *Stack) Push(value int) {
 s.stack = append(s.stack, value)
}

func (s *Stack) Pop() int {
	leng := len(s.stack)
	elem := s.stack[leng - 1]
 	s.stack = s.stack[:leng - 1]
	return elem
}
