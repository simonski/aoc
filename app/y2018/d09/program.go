package d09

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
	"github.com/simonski/cli"
)

/*
Marble Mania
*/

type Puzzle struct {
	title     string
	year      int
	day       int
	input     string
	lines     []string
	HasRender bool
	HasWWW    bool
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("2018")
	iday, _ := strconv.Atoi("09")
	p := Puzzle{year: iyear, day: iday, title: "Marble Mania"}
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

func (puzzle *Puzzle) GetSummary() utils.Summary {
	s := utils.Summary{Day: puzzle.day, Year: puzzle.year, Name: puzzle.title}
	s.ProgressP1 = utils.Started
	s.ProgressP2 = utils.Started
	s.DateStarted = "d09"
	return s
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}

func (puzzle *Puzzle) RunWWW() string {
	return ""
}

func (puzzle *Puzzle) Render(cli *cli.CLI) {
}

// rename this to the year and day in question
func (puzzle *Puzzle) Part1() {
	rl := NewRouteLogic(REAL_DATA)
	fmt.Printf("There are %v locations, %v test routes.\n", len(rl.Locations), len(rl.Routes))

	firstLocation, lastLocation := rl.FindFirstAndLastLocations()
	fmt.Printf("firstLocationis %v, lastLocation is %v\n", firstLocation.Name, lastLocation.Name)
	// var keys []string
	for _, route := range rl.Routes {
		fmt.Printf("\t%v\n", route.Debug())
	}

	// at each leaf we ask: is this path calculated and if so what was the depth and distance
	rl = NewRouteLogic(REAL_DATA)
	fmt.Printf("\nThere are %v locations, there are %v routes.\n", len(rl.Locations), len(rl.Routes))
	for _, route := range rl.Routes {
		fmt.Printf("\t%v\n", route.Debug())
	}
	fmt.Printf("\n")

	// for index := uint64(1); index <= uint64(len(rl.Routes)); index++ {
	// 	fmt.Printf("!%v = %v\n", index, Factorial(index))
	// }
}

// rename this to the year and day in question
func (puzzle *Puzzle) Part2() {
}

// rename and uncomment this to the year and day in question once complete for a gold star!
// func (puzzle *Puzzle) Y202018D09P1Render() {
// }

// rename and uncomment this to the year and day in question once complete for a gold star!
// func (puzzle *Puzzle) Y202018D09P2Render() {
// }
