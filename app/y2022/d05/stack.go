package d5

type Stack struct {
	data []string
}

func NewStack() *Stack {
	s := Stack{}
	data := make([]string, 0)
	s.data = data
	return &s
}

func (s *Stack) Push(entry string) *Stack {
	s.data = append(s.data, entry)
	return s
}

func (s *Stack) Peek() string {
	if s.Size() == 0 {
		return ""
	}
	value := s.data[len(s.data)-1]
	return value
}

func (s *Stack) Flatten() string {
	line := ""
	for _, e := range s.data {
		line += e
	}
	return line
}

func (s *Stack) Pop() string {
	if s.Size() == 0 {
		return ""
	}
	value := s.Peek()
	s.data = s.data[0 : len(s.data)-1]
	return value
}

func (s *Stack) Size() int {
	return len(s.data)
}
