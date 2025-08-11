package d01

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
	"github.com/simonski/cli"
)

/*
The Tyranny of the Rocket Equation
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
	iyear, _ := strconv.Atoi("2019")
	iday, _ := strconv.Atoi("01")
	p := Puzzle{year: iyear, day: iday, title: "The Tyranny of the Rocket Equation"}
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
	s.ProgressP1 = utils.Completed
	s.ProgressP2 = utils.Completed
	s.DateStarted = "d01"
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

/*
For a mass of 12, divide by 3 and round down to get 4, then subtract 2 to get 2.
For a mass of 14, dividing by 3 and rounding down still yields 4, so the fuel required is also 2.
For a mass of 1969, the fuel required is 654.
For a mass of 100756, the fuel required is 33583.

Fuel required to launch a given module is based on its mass. Specifically, to
find the fuel required for a module,
take its mass, divide by three, round down, and subtract 2.
*/
func (puzzle *Puzzle) FuelRequired(mass int) int {
	divided := (mass / 3) - 2
	return divided
}

func (puzzle *Puzzle) FuelRequiredPart2(mass int) int {
	// recursive
	result := (mass / 3) - 2
	if result <= 0 {
		// fmt.Printf("discounted %v\n", result)
		return mass
	} else {
		// fmt.Printf("remaining  %v\n", result)
		return mass + puzzle.FuelRequiredPart2(result)
	}
}

// rename this to the year and day in question
func (puzzle *Puzzle) Part1() {
	fmt.Printf("Y2019 D01 ok!\n")
	lines := strings.Split(REAL_DATA, "\n")
	total := 0
	for _, line := range lines {
		mass, _ := strconv.Atoi(line)
		fuelRequired := puzzle.FuelRequired(mass)
		total += fuelRequired
	}
	fmt.Printf("Part 1 total: %v\n", total)

}

// rename this to the year and day in question
func (puzzle *Puzzle) Part2() {
	fmt.Printf("Y2019 D02 ok!\n")
	lines := strings.Split(REAL_DATA, "\n")
	total := 0
	for _, line := range lines {
		mass, _ := strconv.Atoi(line)
		fuelRequired := puzzle.FuelRequiredPart2(mass) - mass
		total += fuelRequired
	}
	fmt.Printf("Part 2 total: %v\n", total)

}
