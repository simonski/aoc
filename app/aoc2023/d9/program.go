package d9

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
Day 9: Mirage Maintenance
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
	s.DateStarted = "2023-12-09 07:45:16"
	s.DateStarted = "2023-12-09 17:45:16"
	return s
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("2023")
	iday, _ := strconv.Atoi("9")
	p := Puzzle{year: iyear, day: iday, title: "Day 9: Mirage Maintenance"}
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
	fmt.Printf("D9P1: %v\n", part1(strings.Split(REAL_DATA, "\n")))
}

func (puzzle *Puzzle) Part2() {
	fmt.Printf("D9P2: %v\n", part2(strings.Split(REAL_DATA, "\n")))
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}
