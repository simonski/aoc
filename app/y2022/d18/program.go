package d18

import (
	"fmt"
	"strings"

	"github.com/simonski/aoc/utils"
	"github.com/simonski/cli"
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
	p := Puzzle{year: 2022, day: 18, title: "Boiling Boulders"}
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
	g := NewGrid(TEST_DATA)
	external, _ := g.get_external_surface_area()
	_, not_connected := g.CountSides()
	fmt.Println("TEST")
	fmt.Printf("total_surface_area : %v\nget_external_size  : %v\n", not_connected, external)

}

func (puzzle *Puzzle) Part2() {
	fmt.Println()
	fmt.Println("REAL")
	g2 := NewGrid(REAL_DATA)
	external2, _ := g2.get_external_surface_area()
	_, not_connected2 := g2.CountSides()
	fmt.Printf("total_surface_area : %v\nget_external_size  : %v\n", not_connected2, external2)
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}

func (puzzle *Puzzle) GetSummary() utils.Summary {
	s := utils.Summary{Day: puzzle.day, Year: puzzle.year, Name: puzzle.title}
	s.ProgressP1 = utils.Started
	s.ProgressP2 = utils.Started
	s.DateStarted = "d01"
	return s
}

func (puzzle *Puzzle) RunWWW() string {
	return ""
}

func (puzzle *Puzzle) Render(cli *cli.CLI) {
}
