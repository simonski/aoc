package d1

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/simonski/aoc/utils"
)

/*
Day 1: Not Quite Lisp
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
	s.DateStarted = "2023-12-02 07:41:42"
	return s
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("2015")
	iday, _ := strconv.Atoi("1")
	p := Puzzle{year: iyear, day: iday, title: "Day 1: Not Quite Lisp"}
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

	//	func AOC_2015_01_part1_attempt1(cli *goutils.CLI) {

	// total_length := len(DAY_2015_01_DATA)
	only_down := len(strings.ReplaceAll(DAY_2015_01_DATA, "(", ""))
	only_up := len(strings.ReplaceAll(DAY_2015_01_DATA, ")", ""))

	floor := 0
	floor += only_up
	floor -= only_down

	fmt.Printf("From 0 go up %v and down %v ending up on %v\n", only_up, only_down, floor)

}

func (puzzle *Puzzle) Part2() {
	puzzle.Load(REAL_DATA)

	/*
	   --- Part Two ---
	   Now, given the same instructions, find the position of the first character that causes him to enter the basement (floor -1). The first character in the instructions has position 1, the second character has position 2, and so on.

	   For example:

	   ) causes him to enter the basement at character position 1.
	   ()()) causes him to enter the basement at character position 5.
	   What is the position of the character that causes Santa to first enter the basement?
	*/

	start := time.Now()
	end := time.Now()
	floor := 0
	for index := 0; index < len(DAY_2015_01_DATA); index++ {
		c := DAY_2015_01_DATA[index : index+1]
		if c == "(" {
			floor++
		} else if c == ")" {
			floor--
		}
		fmt.Printf("Char[%v] %v, Floor %v\n", index, c, floor)
		if floor == -1 {
			break
		}
	}
	fmt.Printf("%v\n", end.Sub(start))
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}
