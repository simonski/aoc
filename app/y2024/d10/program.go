package d10

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
Day 10: Hoof It
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
	s.DateStarted = "2024-12-10 06:32:44"
	return s
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("2024")
	iday, _ := strconv.Atoi("10")
	p := Puzzle{year: iyear, day: iday, title: "Day 10: Hoof It"}
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
	puzzle.p1(TEST_DATA)
	puzzle.p1(TEST_DATA_2)
	puzzle.p1(REAL_DATA)
}

func (puzzle *Puzzle) p1(data string) {
	tm := NewTrailMap(data)
	for _, start := range tm.starts {
		path := make([]*Cell, 0)
		tm.walk(start, path)
	}
	paths := tm.paths
	fmt.Printf("There are %v starts, %v ends and %v ways of getting to the end trails.\n", len(tm.starts), len(tm.ends), len(paths))

	// a complete trail is the start cell and end cell as a key
	complete_trails := make(map[string]bool)
	for _, p := range paths {
		pathKey := p[0].key + "," + p[9].key
		complete_trails[pathKey] = true
	}

	values := make(map[string]int)
	for _, p := range paths {
		pathKey := p[0].key
		value := values[pathKey]
		value++
		values[pathKey] = value
	}

	total := 0
	for _, v := range values {
		total += v
	}

	// for k := range complete_trails {
	// 	fmt.Println(k)
	// }
	fmt.Printf("Count of unique start/ends is %v\n\n", len(complete_trails))
	fmt.Printf("total=%v\n", total)
}

func (puzzle *Puzzle) Part2() {
	puzzle.Load(REAL_DATA)
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}
