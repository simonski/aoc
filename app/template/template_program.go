package TOKEN_PACKAGE

import (
	"strconv"
	"strings"
)

/*
--- Day 05:  ---

*/

type Puzzle struct {
	title string
	year  int
	day   int
	input string
	lines []string
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("TOKEN_YEAR")
	iday, _ := strconv.Atoi("TOKEN_DAY")
	p := Puzzle{year: iyear, day: iday, title: "TOKEN_TITLE"}
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

func (puzzle *Puzzle) Part1() {
	puzzle.Load(REAL_DATA)
}

func (puzzle *Puzzle) Part2() {
	puzzle.Load(REAL_DATA)
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}
