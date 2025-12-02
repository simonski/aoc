package d01

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
	"github.com/simonski/cli"
)

/*
Day 1: Historian Hysteria
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
	s.DateStarted = "2025-12-01 15:30:49"
	return s
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("2025")
	iday, _ := strconv.Atoi("1")
	p := Puzzle{year: iyear, day: iday, title: "Day 1: Secret Entrance"}
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
	// puzzle.Load(TEST_DATA)

	position := 50
	incremental := 0
	total := 0
	for _, line := range puzzle.lines {
		// original_position := position
		direction := line[0]
		distanceStr := line[1:]
		distance, _ := strconv.Atoi(distanceStr)

		for counter := 0; counter < distance; counter++ {
			if direction == 'L' {
				position -= 1
				if position == -1 {
					position = 99
				}
			} else if direction == 'R' {
				position += 1
				if position == 100 {
					position = 0
				}
			}
			if position == 0 && counter < distance-1 {
				incremental += 1
			}
		}

		if position == 0 {
			total += 1
		}

	}
	fmt.Printf("Total: %v, Incremental: %v, Total: %v\n", total, incremental, total+incremental)

}

func (puzzle *Puzzle) Part2() {
	puzzle.Load(REAL_DATA)
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}

func (puzzle *Puzzle) Render(cli *cli.CLI) {
}
func (puzzle *Puzzle) RunWWW() string {
	return ""
}
