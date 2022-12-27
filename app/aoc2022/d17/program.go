package d17

import (
	"fmt"
	"os"
	"strings"

	"github.com/simonski/cli"
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

var FLOOR_SIZE = 1000

func (puzzle *Puzzle) Part1() {
	cli := cli.New(os.Args)
	log_level := 0
	if cli.Contains("-v") {
		log_level = 1
	} else if cli.Contains("-vv") {
		log_level = 2
	}
	resetFloor := cli.Contains("-floor")
	if os.Args[5] == "test" {
		c := NewChamber(TEST_DATA, log_level)
		// fmt.Println(c.Debug())
		c.RunPart1(2022, resetFloor, FLOOR_SIZE)
		fmt.Println(c.Debug())
		fmt.Printf("Rock Count %v, Height is %v\n", len(c.Rocks), c.Height)

	} else if os.Args[5] == "real" {
		c := NewChamber(REAL_DATA, log_level)
		c.RunPart1(2022, resetFloor, FLOOR_SIZE)
		fmt.Println(c.Debug())
		fmt.Printf("Rock Count %v, Height is %v\n", len(c.Rocks), c.Height)

	}
}

func (puzzle *Puzzle) Part2() {
	if os.Args[5] == "test" {
		c := NewChamber(TEST_DATA, 0)
		fmt.Println(c.Debug())
		c.RunPart2(1000000000000, true, FLOOR_SIZE)
		fmt.Printf("Rock Count %v, Height is %v\n", len(c.Rocks), c.Height)
	} else if os.Args[5] == "real" {
		c := NewChamber(REAL_DATA, 0)
		fmt.Println(c.Debug())
		c.RunPart2(1000000000000, true, FLOOR_SIZE)
		fmt.Printf("Rock Count %v, Height is %v\n", len(c.Rocks), c.Height)
	}
}

func (puzzle *Puzzle) Run() {
	USAGE := "Usage: aoc run 2022 17 P1|P2 test|real (-v)"
	if len(os.Args) < 6 {
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
