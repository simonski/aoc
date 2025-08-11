package d5

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
Day 5: If You Give A Seed A Fertilizer
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
	s.ProgressP1 = utils.Completed
	s.ProgressP2 = utils.Completed
	s.DateStarted = "2023-12-05 07:54:44"
	s.DateCompleted = "2023-12-05 11:53:44"
	return s
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("2023")
	iday, _ := strconv.Atoi("5")
	p := Puzzle{year: iyear, day: iday, title: "Day 5: If You Give A Seed A Fertilizer"}
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
	sm := NewSeedMap(REAL_DATA)
	lowvalue := math.MaxInt
	lowseed := math.MaxInt
	for _, seed := range sm.Seeds {
		location := sm.GetLocationFromSeed(seed)
		if location < lowvalue {
			lowvalue = location
			lowseed = int(seed)
		}
	}
	fmt.Printf("Part1: Seed %v, location %v\n", lowseed, lowvalue)

}

func (puzzle *Puzzle) Part2() {
	puzzle.Load(REAL_DATA)
	sm := NewSeedMap(REAL_DATA)
	min_seed, min_location := sm.GetMinSeedAndLocation()
	fmt.Printf("MinSeed: %v, MinLocation: %v\n", min_seed, min_location)

}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}
