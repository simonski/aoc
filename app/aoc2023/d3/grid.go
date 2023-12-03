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
		// fmt.Printf("IsSymbol('[%v,%v]=%v')=%v\n", col, row, c, false)
		return false
	}
	_, err := strconv.Atoi(c)
	fmt.Printf("IsSymbol('[%v,%v]=%v')=%v\n", col, row, c, err != nil)
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
				candidate := CandidateNumber{Value: value, ColStart: startIndex, ColEnd: endIndex, Row: row}
				results = append(results, &candidate)
			}

		}

	}
	return results
}

type CandidateNumber struct {
	Value    int
	ColStart int
	ColEnd   int
	Row      int
}

func (c *CandidateNumber) Debug() string {
	return fmt.Sprintf("[%v-%v,%v] %v", c.ColStart, c.ColEnd, c.Row, c.Value)
}

func (g *Grid) FindNumbersAdjacent() []*CandidateNumber {
	results := make([]*CandidateNumber, 0)
	numbers := g.FindNumbers()

	for _, candidate := range numbers {
		isAdjacent := false

		if g.IsSymbol(candidate.ColStart-1, candidate.Row) || g.IsSymbol(candidate.ColEnd+1, candidate.Row) {
			isAdjacent = true
		} else {
			for col := candidate.ColStart - 1; col <= candidate.ColEnd+1; col++ {
				if g.IsSymbol(col, candidate.Row-1) || g.IsSymbol(col, candidate.Row+1) {
					isAdjacent = true
					break
				}
			}
		}

		if isAdjacent {
			fmt.Printf("IsAdjacent: %v\n", candidate.Debug())
			results = append(results, candidate)
		} else {
			fmt.Printf("NOT IsAdjacent: %v\n", candidate.Debug())

		}
	}
	return results
}
