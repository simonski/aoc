package d15

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
	p := Puzzle{year: "2022", day: "15", title: "Beacon Exclusion Zone"}
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
	// not 5129731 (too low)
	//     5129730
	//     5129730
	//     5129730

	// data := TEST_DATA
	// row := 10
	// 174,580,378
	data := REAL_DATA
	row := 2000000

	g := NewGrid(data)
	minx, miny, maxx, maxy := g.Bounds()
	width := g.Width()
	height := g.Height()
	fmt.Printf("Grid bounds (%v,%v,%v,%v)\nHeight=%v\nWidth=%v\nVolume=%v\nSensors=%v\nBeacons=%v\n", minx, miny, maxx, maxy, height, width, height*width, len(g.Sensors), len(g.Beacons))
	could_be_beacon, could_not_be_beacon := g.CountCannotsForRow_V2(row)
	// could_be_beacon, could_not_be_beacon := g.CountCannotsForRow(row)
	fmt.Printf("could_be_beacon=%v, could_not_be_beacon=%v\n", could_be_beacon, could_not_be_beacon)
}

func (puzzle *Puzzle) Part2() {
	puzzle.Load(REAL_DATA)
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}
