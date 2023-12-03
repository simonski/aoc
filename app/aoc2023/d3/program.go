package d3

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
Day 3: Gear Ratios
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
	s.DateStarted = "2023-12-03 07:30:19"
	s.DateCompleted = "2023-12-03 08:51:00"
	return s
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("2023")
	iday, _ := strconv.Atoi("3")
	p := Puzzle{year: iyear, day: iday, title: "Day 3: Gear Ratios"}
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

	grid := NewGrid(REAL_DATA)
	numbers := grid.FindNumbersAdjacent()
	total := 0
	for _, n := range numbers {
		total += n.Value
	}
	fmt.Printf("Part 1 : %v\n", total)

}

func (puzzle *Puzzle) Part2() {
	grid := NewGrid(TEST_DATA)
	numbers := grid.FindGearNumbers()
	total := 0
	for _, n := range numbers {
		value := n.Number1.Value * n.Number2.Value
		total += value
	}
	fmt.Printf("Part 2 (test) : %v\n", total)

	grid = NewGrid(REAL_DATA)
	numbers = grid.FindGearNumbers()
	total = 0
	for _, n := range numbers {
		value := n.Number1.Value * n.Number2.Value
		total += value
	}
	fmt.Printf("Part 2 (real) : %v\n", total)

}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}
