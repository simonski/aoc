package d17

type Cycledetector struct {
	data *Dequeue
}

func NewCycledetector() *Cycledetector {
	cd := Cycledetector{data: NewDequeue()}
	return &cd
}

func (cd *Cycledetector) Add(line string, line_number int, rocks int) {
	e := NewEntry(line, line_number, rocks)
	cd.data.Push(e)
}
