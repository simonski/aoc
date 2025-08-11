package d3

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Day 3: Gear Ratios
*/

type Grid struct {
	data []string
	cols int
	rows int
}

func NewGrid(data string) *Grid {
	rows := strings.Split(data, "\n")
	g := Grid{data: rows, cols: len(rows[0]), rows: len(rows)}
	return &g
}

func (g *Grid) Get(col int, row int) string {
	if col >= 0 && col+1 <= g.cols && row >= 0 && row+1 <= g.rows {
		return g.data[row][col : col+1]
	} else {
		return "."
	}
}

func (g *Grid) IsSymbol(col int, row int) bool {
	c := g.Get(col, row)
	if c == "." {
		return false
	}
	_, err := strconv.Atoi(c)
	return err != nil
}

func (g *Grid) FindNumbers() []*CandidateNumber {
	results := make([]*CandidateNumber, 0)
	for row := 0; row < g.rows; row++ {
		for col := 0; col < g.cols; col++ {
			c := g.Get(col, row)
			// fmt.Printf("[%v,%v]=%v\n", col, row, c)
			_, isDigitErr := strconv.Atoi(c)
			startIndex := col
			endIndex := col
			number := c
			if isDigitErr == nil {
				// begin finding the whole number
				for index := col + 1; index < g.cols; index++ {
					nextC := g.Get(index, row)
					_, isDigitErr := strconv.Atoi(nextC)
					if isDigitErr == nil {
						number = fmt.Sprintf("%v%v", number, nextC)
					} else {
						// not a digit, we are finished
						endIndex = index - 1
						col = endIndex
						break
					}
				}
				value, _ := strconv.Atoi(number)
				candidate := NewCandidateNumber(value, startIndex, endIndex, row)
				results = append(results, candidate)
			}

		}

	}
	return results
}

type Point struct {
	col int
	row int
}

func NewPoint(col int, row int) *Point {
	return &Point{col: col, row: row}
}

func (p *Point) Equals(op *Point) bool {
	return p.col == op.col && p.row == op.row
}

type CandidateNumber struct {
	Value    int
	ColStart int
	ColEnd   int
	Row      int
	Symbols  []*Point
	Paired   bool
	Pairing  *CandidateNumber
}

func NewCandidateNumber(value int, colStart int, colEnd int, row int) *CandidateNumber {
	cn := CandidateNumber{Value: value, ColStart: colStart, ColEnd: colEnd, Row: row, Symbols: make([]*Point, 0)}
	return &cn
}

func (c *CandidateNumber) AddSymbol(col int, row int) {
	s := NewPoint(col, row)
	c.Symbols = append(c.Symbols, s)
}

func (c *CandidateNumber) IsGearedWith(oc *CandidateNumber) bool {
	if len(c.Symbols) == 1 && len(oc.Symbols) == 1 {
		c1 := c.Symbols[0]
		c2 := oc.Symbols[0]
		if c1.Equals(c2) {
			// fmt.Printf("%v - IsGearedWith()? %v  - both have 1 symbol each that matches\n", c.Debug(), oc.Debug())
			return true
		} else {
			// fmt.Printf("%v - IsGearedWith()? %v  - both have 1 symbol each -but does not match\n", c.Debug(), oc.Debug())

		}
	} else {
		// fmt.Printf("%v - IsGearedWith()? %v - No, 1 has %v symbols, 2 has %v symbols\n", c.Debug(), oc.Debug(), len(c.Symbols), len(oc.Symbols))

	}
	return false
}

func (c *CandidateNumber) Debug() string {
	return fmt.Sprintf("[%v-%v,%v] %v", c.ColStart, c.ColEnd, c.Row, c.Value)
}

func (g *Grid) FindNumbersAdjacent() []*CandidateNumber {
	results := make([]*CandidateNumber, 0)
	numbers := g.FindNumbers()

	for _, candidate := range numbers {
		isAdjacent := false

		if g.IsSymbol(candidate.ColStart-1, candidate.Row) {
			isAdjacent = true
			candidate.AddSymbol(candidate.ColStart-1, candidate.Row)
		} else if g.IsSymbol(candidate.ColEnd+1, candidate.Row) {
			isAdjacent = true
			candidate.AddSymbol(candidate.ColEnd+1, candidate.Row)
		} else {
			for col := candidate.ColStart - 1; col <= candidate.ColEnd+1; col++ {
				if g.IsSymbol(col, candidate.Row-1) {
					candidate.AddSymbol(col, candidate.Row-1)
					isAdjacent = true
				} else if g.IsSymbol(col, candidate.Row+1) {
					isAdjacent = true
					candidate.AddSymbol(col, candidate.Row+1)
				}

			}
		}

		if isAdjacent {
			results = append(results, candidate)
		}
	}
	return results
}

type GearNumbers struct {
	Number1 *CandidateNumber
	Number2 *CandidateNumber
}

func (g *Grid) FindGearNumbers() []*GearNumbers {
	results := make([]*GearNumbers, 0)
	numbers := g.FindNumbersAdjacent()
	for index := 0; index < len(numbers); index++ {
		c1 := numbers[index]
		if c1.Paired {
			continue
		}
		for index2 := 0; index2 < len(numbers); index2++ {
			if index2 == index {
				continue
			}
			c2 := numbers[index2]
			if c2.Paired {
				continue
			} else if c1.IsGearedWith(c2) {
				c1.Paired = true
				c1.Pairing = c2
				c2.Paired = true
				c2.Pairing = c1
				gn := GearNumbers{Number1: c1, Number2: c2}
				results = append(results, &gn)
			}
		}
	}
	return results
}
