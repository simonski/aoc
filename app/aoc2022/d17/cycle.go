package d17

import "fmt"

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

func (cd *Cycledetector) Buildkey(position int, size int) string {
	k := ""
	for index := position; index < position+size; index++ {
		line := cd.data.data[index].Line
		k += line
		if index > 0 {
			k += "\n"
		}
	}
	return k
}

func (cd *Cycledetector) Debugkey(position int, size int) string {
	k := ""
	for index := position; index < position+size; index++ {
		line := cd.data.data[index].Line
		k = fmt.Sprintf("%v%v [index=%v]", k, line, index)
		if index > 0 {
			k += "\n"
		}
	}
	return k
}

func (cd *Cycledetector) FindRepeatingKeys(key string, keysize int) []int {
	results := make([]int, 0)
	for index := 0; index < cd.data.Size()-keysize; index++ {
		candidate := cd.Buildkey(index, keysize)
		if candidate == key {
			results = append(results, index)
		}
	}
	return results
}
