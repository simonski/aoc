package TOKEN_PACKAGE

import (
	"os"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
--- Day TOKEN_DAY: TOEKN_TITLE ---

*/

type Puzzle struct {
	title string
	year  string
	day   string
	input string
	lines []string
}

func (p *Puzzle) Summary() *utils.Summary {
	year, _ := strconv.Atoi("TOKEN_YEAR")
	day, _ := strconv.Atoi("TOKEN_DAY")
	s := utils.NewSummary(year, day)
	return s
}

func NewPuzzleWithData(input string) *Puzzle {
	p := Puzzle{year: "TOKEN_YEAR", day: "TOKEN_DAY", title: "TOKEN_TITLE"}
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
	c := utils.NewCLI(os.Args)
	if c.Contains("P1") {
		// puzzle.Part1()
	} else if c.Contains("P2") {
		// puzzle.Part2()
	} else {
		// puzzle.Part2()
	}
}
