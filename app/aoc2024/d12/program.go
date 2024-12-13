package d12

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
Day 12: Garden Groups
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
	s.DateStarted = "2024-12-13 11:25:17"
	return s
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("2024")
	iday, _ := strconv.Atoi("12")
	p := Puzzle{year: iyear, day: iday, title: "Day 12: Garden Groups"}
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
	g := NewGrid(TEST_DATA_1)
	g.findRegions()
	fmt.Printf("There are %v region for TEST_DATA_1\n", g.region)
	value := 0
	new_value := 0
	for _, r := range g.regions {
		fmt.Printf("region %v has area %v perimeter %v = %v, sides=%v, new_price=%v\n", r.id, r.area(), r.perimeter(g), r.price(g), r.sides(g), r.new_price(g))
		value += r.price(g)
		new_value += r.new_price(g)
	}
	fmt.Printf("grid value is %v, new_value=%v\n", value, new_value)
	fmt.Println()

	g = NewGrid(TEST_DATA_2)
	g.findRegions()
	fmt.Printf("There are %v region for TEST_DATA_2\n", g.region)
	value = 0
	new_value = 0
	for _, r := range g.regions {
		fmt.Printf("region %v has area %v perimeter %v = %v, sides=%v, new_price=%v\n", r.id, r.area(), r.perimeter(g), r.price(g), r.sides(g), r.new_price(g))
		value += r.price(g)
		new_value += r.new_price(g)
	}
	fmt.Printf("grid value is %v, new_value=%v\n", value, new_value)
	fmt.Println()

	g = NewGrid(TEST_DATA_3)
	g.findRegions()
	value = 0
	new_value = 0
	for _, r := range g.regions {
		fmt.Printf("region %v has area %v perimeter %v = %v, sides=%v, new_price=%v\n", r.id, r.area(), r.perimeter(g), r.price(g), r.sides(g), r.new_price(g))
		value += r.price(g)
		new_value += r.new_price(g)
	}
	fmt.Printf("grid value is %v, new_value=%v\n", value, new_value)
	fmt.Println()

	g = NewGrid(TEST_DATA_E)
	g.findRegions()
	value = 0
	new_value = 0
	for _, r := range g.regions {
		fmt.Printf("region %v has area %v perimeter %v = %v, sides=%v, new_price=%v\n", r.id, r.area(), r.perimeter(g), r.price(g), r.sides(g), r.new_price(g))
		value += r.price(g)
		new_value += r.new_price(g)
	}
	fmt.Printf("grid value is %v, new_value=%v\n", value, new_value)
	fmt.Println()

	// g = NewGrid(REAL_DATA)
	// g.findRegions()
	// value = 0
	// fmt.Printf("There are %v region for REAL_DATA\n", g.region)
	// for _, r := range g.regions {
	// 	fmt.Printf("region %v has area %v perimeter %v = %v, size=%v\n", r.id, r.area(), r.perimeter(g), r.price(g), r.size(g))
	// 	value += r.price(g)
	// }
	// fmt.Printf("grid value is %v\n", value)
	// fmt.Println()
}

func (puzzle *Puzzle) Part2() {
	puzzle.Load(REAL_DATA)
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}
