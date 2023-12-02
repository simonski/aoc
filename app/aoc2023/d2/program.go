package d2

import (
	"fmt"
	"strings"

	"github.com/simonski/aoc/utils"
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
	games []*Game
}

func NewPuzzleWithData(input string) *Puzzle {
	p := Puzzle{year: 2023, day: 2, title: "--- Day 2: Cube Conundrum ---"}
	p.Load(input)
	return &p
}

func NewPuzzle() *Puzzle {
	return NewPuzzleWithData(REAL_DATA)
}

func (puzzle *Puzzle) GetSummary() utils.Summary {
	s := utils.Summary{Day: puzzle.day, Year: puzzle.year, Name: puzzle.title}
	s.ProgressP1 = utils.Completed
	s.ProgressP2 = utils.Completed
	return s
}

func (puzzle *Puzzle) Load(input string) {
	lines := strings.Split(input, "\n")
	puzzle.input = input
	puzzle.lines = lines
	puzzle.games = make([]*Game, 0)
	for _, line := range lines {
		g := NewGame(line)
		puzzle.games = append(puzzle.games, g)
	}
}

func (puzzle *Puzzle) Part1() {
	puzzle.Load(REAL_DATA)
	games := puzzle.FindGames(12, 13, 14)
	total := 0
	for _, g := range games {
		total += g.ID
	}
	fmt.Printf("Part1: %v\n", total)
}

func (puzzle *Puzzle) Part2() {
	puzzle.Load(REAL_DATA)
	totalPower := 0
	for _, g := range puzzle.games {
		totalPower += g.Power()
	}
	fmt.Printf("Part2: %v\n", totalPower)

}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}

func (puzzle *Puzzle) FindGames(red int, green int, blue int) []*Game {
	retain := make([]*Game, 0)
	for _, g := range puzzle.games {
		if g.Passes(red, green, blue) {
			retain = append(retain, g)
		}
	}
	return retain
}
