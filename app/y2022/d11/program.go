package d11

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
	p := Puzzle{year: 2022, day: 11, title: "Monkey In The Middle"}
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
}

func (puzzle *Puzzle) Part2() {
	troupe := NewTroupe(REAL_DATA)

	for index := 0; index < 10000; index++ {
		troupe.Round(false, 1)
		fmt.Printf("\nRound %v\n", troupe.RoundNum)

		for index, monkey := range troupe.Monkeys {
			fmt.Printf("Monkey[%v] %v inspections, %v entries.\n", index, monkey.InspectCount, len(monkey.Items))
		}
	}
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
