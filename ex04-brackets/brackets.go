package brackets

type Stack struct {
	stack   []byte
	counter int
}

func New() *Stack {
	return &Stack{}
}

func (s *Stack) Push(value byte) {
	s.stack = append(s.stack, value)
	s.counter++
}

func (s *Stack) Pop() byte {
	leng := len(s.stack)
	elem := s.stack[leng-1]
	s.stack = s.stack[:leng-1]
	s.counter--
	return elem
}

func Bracket(str string) (bool, error) {
	var stack *Stack = New()
	var value byte
	if len(str) < 1 {
		return true, nil
	}
	for _, c := range str {
		if byte(c) == '(' || byte(c) == '[' || byte(c) == '{' {
			stack.Push(byte(c))
		} else if stack.counter > 0 {
			value = stack.Pop()
			if value == '(' {
				if byte(c) != ')' {
					return false, nil
				}
			} else if value != byte(c)-2 {
				return false, nil
			}
		}
	}
	if stack.counter != 0 {
		return false, nil
	}
	return true, nil
}
