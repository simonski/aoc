package d03

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
	"github.com/simonski/cli"
	"github.com/simonski/goutils"
)

/*
Perfectly Spherical Houses in a Vacuum
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
	iyear, _ := strconv.Atoi("2015")
	iday, _ := strconv.Atoi("03")
	p := Puzzle{year: iyear, day: iday, title: "Perfectly Spherical Houses in a Vacuum"}
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
	s.DateStarted = "d03"
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

func (puzzle *Puzzle) Part1() {
	grid := NewGrid201503()
	grid.Increment(0, 0)
	for index := 0; index < len(REAL_DATA); index++ {
		c := REAL_DATA[index : index+1]
		if c == "<" {
			grid.West()
		} else if c == ">" {
			grid.East()
		} else if c == "^" {
			grid.North()
		} else if c == "v" {
			grid.South()
		} else {
			panic("wtf")
		}
	}

	total_single_presents := 0
	for _, value := range grid.Counter.Data {
		if value == 1 {
			total_single_presents++
		}
	}

	fmt.Printf("Part1: Single present households: %v\n", total_single_presents)
	fmt.Printf("Part1: Total households: %v\n", len(grid.Counter.Data))
}

/*
--- Part Two ---
The next year, to speed up the process, Santa creates a robot version of himself, Robo-Santa, to deliver presents with him.

Santa and Robo-Santa start at the same location (delivering two presents to the same starting house), then take turns moving based on instructions from the elf, who is eggnoggedly reading from the same script as the previous year.

This year, how many houses receive at least one present?

For example:

^v delivers presents to 3 houses, because Santa goes north, and then Robo-Santa goes south.
^>v< now delivers presents to 3 houses, and Santa and Robo-Santa end up back where they started.
^v^v^v^v^v now delivers presents to 11 houses, with Santa going one direction and Robo-Santa going the other.
*/

func (puzzle *Puzzle) Part2() {
	grid := NewGrid201503()
	grid.Increment(0, 0)
	grid.Increment(0, 0)
	data := REAL_DATA
	for index := 0; index < len(data); index += 2 {
		c1 := data[index : index+1]
		if c1 == "<" {
			grid.West()
		} else if c1 == ">" {
			grid.East()
		} else if c1 == "^" {
			grid.North()
		} else if c1 == "v" {
			grid.South()
		} else {
			panic("wtf")
		}

		idx := index + 1
		c2 := data[idx : idx+1]
		if c2 == "<" {
			grid.RoboWest()
		} else if c2 == ">" {
			grid.RoboEast()
		} else if c2 == "^" {
			grid.RoboNorth()
		} else if c2 == "v" {
			grid.RoboSouth()
		} else {
			panic("wtf")
		}

	}

	total_single_presents := 0
	for _, value := range grid.Counter.Data {
		if value == 1 {
			total_single_presents++
		}
	}

	fmt.Printf("Part2: Single present households: %v\n", total_single_presents)
	fmt.Printf("Part2: Total households: %v\n", len(grid.Counter.Data))
}

type Grid201503 struct {
	x       int
	y       int
	robo_x  int
	robo_y  int
	Counter *goutils.Counter
}

func NewGrid201503() *Grid201503 {
	g := Grid201503{x: 0, y: 0, Counter: goutils.NewCounter()}
	return &g
}

func (grid *Grid201503) Increment(x int, y int) {
	key := fmt.Sprintf("%v,%v", x, y)
	grid.Counter.Increment(key)
}

func (grid *Grid201503) North() {
	y := grid.y + 1
	x := grid.x
	grid.Increment(x, y)
	grid.x = x
	grid.y = y
}

func (grid *Grid201503) South() {
	y := grid.y - 1
	x := grid.x
	grid.Increment(x, y)
	grid.x = x
	grid.y = y
}

func (grid *Grid201503) East() {
	y := grid.y
	x := grid.x + 1
	grid.Increment(x, y)
	grid.x = x
	grid.y = y
}

func (grid *Grid201503) West() {
	y := grid.y
	x := grid.x - 1
	grid.Increment(x, y)
	grid.x = x
	grid.y = y
}

func (grid *Grid201503) RoboNorth() {
	y := grid.robo_y + 1
	x := grid.robo_x
	grid.Increment(x, y)
	grid.robo_x = x
	grid.robo_y = y
}

func (grid *Grid201503) RoboSouth() {
	y := grid.robo_y - 1
	x := grid.robo_x
	grid.Increment(x, y)
	grid.robo_x = x
	grid.robo_y = y
}

func (grid *Grid201503) RoboEast() {
	y := grid.robo_y
	x := grid.robo_x + 1
	grid.Increment(x, y)
	grid.robo_x = x
	grid.robo_y = y
}

func (grid *Grid201503) RoboWest() {
	y := grid.robo_y
	x := grid.robo_x - 1
	grid.Increment(x, y)
	grid.robo_x = x
	grid.robo_y = y
}
