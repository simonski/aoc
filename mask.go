package main

type Mask struct {
	data string
}

func (b *Mask) Get(index int) string {
	if len(b.data) > index {
		position := len(b.data) - index
		return b.data[position-1 : position]
	} else {
		return "X"
	}
}

func NewMaskFromValue(value string) *Mask {
	m := Mask{data: value}
	return &m
}
