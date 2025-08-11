package d15

import "fmt"

type Sensor struct {
	Point    *Point
	Beacon   *Beacon
	Strength int
}

func (s *Sensor) String() string {
	return fmt.Sprintf("S(%v,%v)", s.Point.X, s.Point.Y)
}

func NewSensor(input string) *Sensor { // 0,4
	p := NewPoint(input)
	s := Sensor{}
	s.Point = p
	p.Sensor = &s
	return &s
}

type Beacon struct {
	Point *Point
}

func NewBeacon(input string) *Beacon {
	p := NewPoint(input)
	b := Beacon{}
	b.Point = p
	p.Beacon = &b
	return &b
}

func (b *Beacon) String() string {
	return fmt.Sprintf("B(%v,%v)", b.Point.X, b.Point.Y)
}
