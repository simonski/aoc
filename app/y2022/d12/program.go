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
	result := dijkstra_v2(grid, true)
	fmt.Printf("min steps : %v\n", result)
}

func (puzzle *Puzzle) Part2() {
	grid := NewGrid(REAL_DATA)
	min_result := 100000000
	as := make([]*Point, 0)
	for row := 0; row < grid.Rows; row++ {
		for col := 0; col < grid.Cols; col++ {
			p := grid.Get(row, col)
			if p.Letter == "a" {
				as = append(as, p)
			}
		}
	}

	fmt.Printf("Part 2: There are %v letter a's.\n", len(as))

	for index, p := range as {
		fmt.Printf("%v/%v (%v)\n", index+1, len(as), p)
		a_grid := NewGrid(REAL_DATA)
		ap := a_grid.Get(p.Row, p.Col)
		a_grid.Start.IsStart = false
		a_grid.Start.Letter = "a"

		ap.IsStart = true
		ap.Letter = "S"
		ap.VisitDirection = "."
		a_grid.Start = ap

		result := dijkstra_v2(a_grid, false)
		if result < min_result {
			min_result = result
			fmt.Printf("%v/%v (%v), min steps : %v\n", index+1, len(as), p, min_result)
		} else {
			fmt.Printf("%v/%v (%v), steps : %v\n", index+1, len(as), p, result)

		}
	}
	fmt.Printf("min steps : %v\n", min_result)
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}
