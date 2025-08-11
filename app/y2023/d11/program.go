package d11

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
Day 11: Cosmic Expansion
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
	s.ProgressP1 = utils.Completed
	s.ProgressP2 = utils.Completed
	s.DateStarted = "2023-12-11 11:00:52"
	s.DateCompleted = "2023-12-11 11:50:52"
	return s
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("2023")
	iday, _ := strconv.Atoi("11")
	p := Puzzle{year: iyear, day: iday, title: "Day 11: Cosmic Expansion"}
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
	u := NewUniverse(REAL_DATA)
	u.Expand()
	pairs := u.Pairs()
	total := 0
	for _, p := range pairs {
		cell1 := p[0]
		cell2 := p[1]
		distance := cell1.Distance(cell2)
		total += distance
	}
	fmt.Printf("D11P1: %v\n", total)
}

func (puzzle *Puzzle) Part2() {
	u := NewUniverse(REAL_DATA)
	// u.Expand()
	pairs := u.Pairs()
	total := 0
	for _, p := range pairs {
		cell1 := p[0]
		cell2 := p[1]
		distance := cell1.DistanceP2(u, cell2, 1000000)
		total += distance
	}
	fmt.Printf("D11P2: %v\n", total)
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}
