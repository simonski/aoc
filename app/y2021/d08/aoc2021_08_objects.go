package aoc2021

import (
	"strings"
)

// contains 4 digits
type Display struct {
	Digits []*Digit
}

type Helper struct {
	Zero  *Pattern
	One   *Pattern
	Two   *Pattern
	Three *Pattern
	Four  *Pattern
	Five  *Pattern
	Six   *Pattern
	Seven *Pattern
	Eight *Pattern
	Nine  *Pattern
}

func NewHelper(data string) *Helper {
	lines := strings.Split(data, "\n")
	h := Helper{}
	for _, line := range lines {
		p := NewPattern(line)
		for _, ov := range p.OutputValue {
			value := p.DigitValue(ov)
			if value == 1 {
				h.One = p
			} else if value == 4 {
				h.Four = p
			} else if value == 7 {
				h.Seven = p
			} else if value == 8 {
				h.Eight = p
			}
		}
	}
	return &h
}

// a single digit
type Digit struct {
	A bool
	B bool
	C bool
	D bool
	E bool
	F bool
	G bool
}

// returns number of segments switched on
func (d *Digit) CountSegmentsOn() int {
	count := 0
	if d.A {
		count++
	}
	if d.B {
		count++
	}
	if d.C {
		count++
	}
	if d.D {
		count++
	}
	if d.E {
		count++
	}
	if d.F {
		count++
	}
	if d.G {
		count++
	}
	return count
}

func (d *Digit) Value() int {
	if d.A && d.B && d.C && !d.D && d.E && d.F && d.G {
		return 0
	} else if !d.A && !d.B && d.C && !d.D && !d.E && d.F && !d.G {
		return 1
	} else if d.A && !d.B && d.C && d.D && d.E && !d.F && d.G {
		return 2
	} else if d.A && !d.B && d.C && d.D && !d.E && d.F && d.G {
		return 4
	} else if d.A && d.B && !d.C && d.D && !d.E && d.F && d.G {
		return 5
	} else if d.A && d.B && !d.C && d.D && d.E && d.F && d.G {
		return 6
	} else if d.A && !d.B && d.C && !d.D && !d.E && d.F && !d.G {
		return 7
	} else if d.A && d.B && d.C && d.D && d.E && d.F && d.G {
		return 8
	} else if d.A && d.B && d.C && d.D && !d.E && d.F && d.G {
		return 9
	} else {
		return -1
	}
}

func (d *Digit) ValueFromCount() int {
	segments := d.CountSegmentsOn()
	if segments == 2 {
		return 1
	} else if segments == 4 {
		return 4
	} else if segments == 3 {
		return 7
	} else if segments == 7 {
		return 8
	}
	return -1
}

type Pattern struct {
	// acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf
	Line           string
	SignalPatterns []string
	OutputValue    []string

	One   string
	Two   string
	Three string
	Four  string
	Fives []string
	Sixes []string
	Seven string
	Eight string
	Nine  string
	Zero  string

	A string
	B string
	C string
	D string
	E string
	F string
	G string
}

func (p *Pattern) Reprogram(a string, b string, c string, d string, e string, f string, g string) {
	p.A = a
	p.B = b
	p.C = c
	p.D = d
	p.E = e
	p.F = f
	p.G = g
}

func (p *Pattern) GetValue(candidate string) int {
	zero := p.A + p.B + p.C + p.E + p.F + p.G
	one := p.C + p.F
	two := p.A + p.C + p.D + p.E + p.G
	three := p.A + p.C + p.D + p.F + p.G
	four := p.B + p.C + p.D + p.F
	five := p.A + p.B + p.D + p.F + p.G
	six := p.A + p.B + p.D + p.E + p.F + p.G
	seven := p.A + p.C + p.F
	eight := p.A + p.B + p.C + p.D + p.E + p.F + p.G
	nine := p.A + p.B + p.C + p.D + p.F + p.G
	// fmt.Printf("%v %v %v %v %v %v %v %v %v %v\n", zero, one, two, three, four, five, six, seven, eight, nine)

	if p.Subtract(candidate, one) == "" {
		return 1
	} else if p.Subtract(candidate, seven) == "" {
		return 7
	} else if p.Subtract(candidate, two) == "" {
		return 2
	} else if p.Subtract(candidate, three) == "" {
		return 3
	} else if p.Subtract(candidate, four) == "" {
		return 4
	} else if p.Subtract(candidate, five) == "" {
		return 5
	} else if p.Subtract(candidate, nine) == "" {
		return 9
	} else if p.Subtract(candidate, six) == "" {
		return 6
	} else if p.Subtract(candidate, zero) == "" {
		return 0
	} else if p.Subtract(candidate, eight) == "" {
		return 8
	} else {
		return -1
	}
}

func (p *Pattern) DigitValue(v string) int {
	segments := len(v)
	if segments == 2 {
		return 1
	} else if segments == 4 {
		return 4
	} else if segments == 3 {
		return 7
	} else if segments == 7 {
		return 8
	}
	return -1
}

// 1 = 2
// 4 = 4
// 7 = 3
// 8 = 7

func (p *Pattern) Subtract(signal1 string, signal2 string) string {
	// subtract signal2 from signal1
	result := signal1
	for index := 0; index < len(signal2); index++ {
		c := signal2[index : index+1]
		result = strings.ReplaceAll(result, c, "")
	}
	return result
}

func NewPattern(line string) *Pattern {
	p := Pattern{Line: line}
	splits := strings.Split(line, "|")
	s1 := strings.Trim(splits[0], " ")
	s2 := strings.Trim(splits[1], " ")
	p.SignalPatterns = strings.Split(s1, " ")
	p.OutputValue = strings.Split(s2, " ")

	for _, sp := range p.SignalPatterns {
		value := p.DigitValue(sp)
		if value == 1 {
			p.One = sp
		} else if value == 4 {
			p.Four = sp
		} else if value == 7 {
			p.Seven = sp
		} else if value == 8 {
			p.Eight = sp
		}

		if len(sp) == 5 {
			p.Fives = append(p.Fives, sp)
		} else if len(sp) == 6 {
			p.Sixes = append(p.Sixes, sp)
		}
	}
	return &p
}
