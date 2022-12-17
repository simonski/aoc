package d15

import (
	"fmt"
	"strings"
)

/*
--- Day 05:  ---

*/

type Puzzle struct {
	title string
	year  string
	day   string
	input string
	lines []string
}

func NewPuzzleWithData(input string) *Puzzle {
	p := Puzzle{year: "2022", day: "15", title: "Beacon Exclusion Zone"}
	p.Load(input)
	return &p
}

func NewPuzzle() *Puzzle {
	return NewPuzzleWithData(REAL_DATA)
}

func (puzzle *Puzzle) Load(input string) {
	lines := strings.Split(input, "\n")
	puzzle.input = input
	puzzle.lines = lines
}

func (puzzle *Puzzle) Part1_x() {
	g := NewGrid(TEST_DATA)
	VERBOSE := false
	cannot := g.CountCannotBePresent(VERBOSE, 10)
	fmt.Printf("Segments for row %v: %v\n", 10, cannot)
}

func (puzzle *Puzzle) Part1_y() {
	g := NewGrid(REAL_DATA)
	VERBOSE := false
	cannot := g.CountCannotBePresent(VERBOSE, 2000000)
	fmt.Printf("Segments for row %v: %v\n", 2000000, cannot)
}

func (puzzle *Puzzle) Part2_x() {
	g := NewGrid(TEST_DATA)
	VERBOSE := true
	answer_row := -1
	answer_col := -1
	VERBOSE = false
	for row := 0; row < 20; row++ {
		gaps, col := g.CountGaps(VERBOSE, row)
		fmt.Printf("Gaps for row[%v]=%v, col=%v\n", row, gaps, col)
		if gaps == 1 && col > -1 {
			if g.Get(col, row) == nil {
				answer_row = row
				answer_col = col
			}
		}
	}
	fmt.Printf("(%v,%v)", answer_col, answer_row)
}

func (puzzle *Puzzle) Part2_y() {
	g := NewGrid(REAL_DATA)
	VERBOSE := true
	answer_row := -1
	answer_col := -1
	VERBOSE = false
	for row := 0; row < 4000000; row++ {
		gaps, col := g.CountGaps(VERBOSE, row)
		fmt.Printf("Gaps for row[%v]=%v, col=%v\n", row, gaps, col)
		if gaps == 1 && col > -1 {
			if g.Get(col, row) == nil {
				answer_row = row
				answer_col = col
			}
		}
	}
	fmt.Printf("(%v,%v)", answer_col, answer_row)
}

func (puzzle *Puzzle) Run() {

	puzzle.Part1_x()
	puzzle.Part1_y()

	puzzle.Part2_x()
	puzzle.Part2_y()

}
