package d12

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
	p := Puzzle{year: "2022", day: "12", title: "Hill Climbing Algorithm"}
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
	grid := NewGrid(REAL_DATA)
	path := HillClimb(true, grid)
	fmt.Printf("%v\n", path)
	fmt.Printf("\nBest path size is %v\n", path.Size())
	for index, p := range path.Points {
		fmt.Printf("[%v] %v\n", index, p)
	}
}

func (puzzle *Puzzle) Part2() {
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}
