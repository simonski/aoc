package d17

type Entry struct {
	Line   string
	Number int
	Rocks  int
}

func NewEntry(line string, number int, rocks int) *Entry {
	return &Entry{Line: line, Number: number, Rocks: rocks}
}

type Dequeue struct {
	data []*Entry
}

func NewDequeue() *Dequeue {
	return &Dequeue{data: make([]*Entry, 0)}
}

func (s *Dequeue) Push(li *Entry) {
	s.data = append(s.data, li)
}

func (s *Dequeue) Pop() *Entry {
	if s.Size() == 0 {
		return nil
	}
	index := len(s.data) - 1
	li := s.data[index]
	s.data = s.data[0 : index-1]
	return li
}

func (s *Dequeue) PrePush(li *Entry) {
	d := make([]*Entry, 0)
	d = append(d, li)
	d = append(d, s.data...)
	s.data = d
}

func (s *Dequeue) PrePop() *Entry {
	if s.Size() == 0 {
		return nil
	}
	d := s.data[0]
	s.data = s.data[1:]
	return d
}

func (s *Dequeue) Size() int {
	return len(s.data)
}
