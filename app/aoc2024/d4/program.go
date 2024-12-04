package d4

import (
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
Day 4: Ceres Search
*/

type Puzzle struct {
	title string
	year  int
	day   int
	input string
	lines []string
}

func (puzzle *Puzzle) GetSummary() utils.Summary {
	s := utils.Summary{Day: puzzle.day, Year: puzzle.year, Name: puzzle.title}
	s.ProgressP1 = utils.NotStarted
	s.ProgressP2 = utils.NotStarted
	s.DateStarted = "2024-12-04 06:28:26"
	return s
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("2024")
	iday, _ := strconv.Atoi("4")
	p := Puzzle{year: iyear, day: iday, title: "Day 4: Ceres Search"}
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
