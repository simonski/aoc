package d10

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
Day 10: Pipe Maze
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
	s.ProgressP1 = utils.Started
	s.ProgressP2 = utils.NotStarted
	s.DateStarted = "2023-12-10 07:00:44"
	return s
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("2023")
	iday, _ := strconv.Atoi("10")
	p := Puzzle{year: iyear, day: iday, title: "Day 10: Pipe Maze"}
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
	// puzzle.Load(REAL_DATA)
	// g := NewGrid(TEST_DATA_2)
	// fmt.Println(g.Debug())

	// g = NewGrid(TEST_DATA_3)
	// fmt.Println(g.Debug())

	// g = NewGrid(TEST_DATA_4)
	// fmt.Println(g.Debug())

	// g = NewGrid(TEST_DATA_5)
	// fmt.Println(g.Debug())

	// g = NewGrid(TEST_DATA_6)
	// fmt.Println(g.Debug())

	// g = NewGrid(TEST_DATA_7)
	// fmt.Println(g.Debug())

	// g = NewGrid(TEST_DATA_8)
	// fmt.Println(g.Debug())

	// g = NewGrid(TEST_DATA_9)
	// fmt.Println(g.Debug())

	// g = NewGrid(REAL_DATA)
	// fmt.Println(g.Debug())

	g := NewGrid(TEST_DATA_8)
	// s_point := g.current_s_pos
	fmt.Println(g.Debug())
	fmt.Println("")
	fmt.Printf("Walk: start at (%v,%v)\n", g.start_s_pos.x, g.start_s_pos.y)
	g.Walk(g.start_s_pos)
	fmt.Println(g.DebugDistance())

	g = NewGrid(REAL_DATA)
	// s_point := g.current_s_pos
	fmt.Println(g.Debug())
	fmt.Println("")
	fmt.Printf("Walk: start at (%v,%v)\n", g.start_s_pos.x, g.start_s_pos.y)
	result := g.Walk(g.start_s_pos)
	fmt.Println(g.DebugDistance())
	fmt.Printf("P1: %v\n", result)

}

func (puzzle *Puzzle) Part2() {
	fmt.Println("PART2")
	g := NewGrid(REAL_DATA)
	fmt.Println("PART2")
	g.Walk(g.start_s_pos)
	fmt.Println("PART2")
	fmt.Println(g.DebugPathOnly())
	fmt.Println("/PART2")

	// 	take a corner, anything that is connected to the corner and NOT a visited tile is outside
	// subtract all the path entries
	// == inside/enclosed

}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	// puzzle.Part2()
}
