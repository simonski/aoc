package d17

import (
	"fmt"
	"os"
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
	p := Puzzle{year: "2022", day: "17", title: "Pyroclastic Flow"}
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
	c := NewChamber(TEST_DATA)
	// fmt.Println(c.Debug())
	c.Run(false, false, 1000000000000)
	fmt.Printf("Rock Count %v, Height is %v\n", len(c.Rocks), c.Height)
}

func (puzzle *Puzzle) Part2() {
	c := NewChamber(REAL_DATA)
	fmt.Println(c.Debug())
	c.Run(false, false, 1000000000000)
	fmt.Printf("Rock Count %v, Height is %v\n", len(c.Rocks), c.Height)
}

func (puzzle *Puzzle) Run() {
	USAGE := "Usage: aoc run 2022 17 P1|P2 test|real (-v)"
	if len(os.Args) != 6 {
		fmt.Println(USAGE)
		os.Exit(1)
	} else if os.Args[4] == "P1" {
		puzzle.Part1()
	} else if os.Args[4] == "P2" {
		puzzle.Part2()
	} else {
		fmt.Println(USAGE)
		os.Exit(1)
	}
}
