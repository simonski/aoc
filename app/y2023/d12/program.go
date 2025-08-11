package d12

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
	"github.com/simonski/cli"
)

/*
Day 12: Hot Springs
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
	s.DateStarted = "2023-12-12 08:01:20"
	return s
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("2023")
	iday, _ := strconv.Atoi("12")
	p := Puzzle{year: iyear, day: iday, title: "Day 12: Hot Springs"}
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

	c := &cli.CLI{Args: os.Args}
	c.IS_VERBOSE = c.Contains("-v")

	g := NewGrid(TEST_DATA_2)
	g.Part1(c.IS_VERBOSE)

	g = NewGrid(REAL_DATA)
	g.Part1(c.IS_VERBOSE)

}

func (puzzle *Puzzle) Part2() {

	c := cli.New(os.Args)
	var g *Grid
	if c.Contains("-line") {
		data := c.GetStringOrDie("-line")
		g = NewGrid(data)
	} else {
		g = NewGrid(REAL_DATA)
		// g = NewGrid(TEST_DATA_2)
	}
	total := 0
	verbose := c.Contains("-v")
	caching := c.Contains("-c")

	fmt.Printf("Verbose : %v\n", verbose)
	fmt.Printf("Caching : %v\n", caching)

	// the cache should be a fail fast where we store
	// the result at this posiiton and whether it CAN continue
	// not the line itself, which is not important as a negative
	COUNTER = 0
	for index, row := range g.data {
		context := NewContext(c)
		context.cacheEnabled = caching
		context.verbose = verbose
		big_row := row.Grow()
		row = big_row
		line := ""
		// results := descend(0, line, row.left, cache, row.rules, verbose)
		// results := descend(0, line, big_row.left, cache, big_row.rules, verbose)
		if verbose {
			fmt.Printf("[%v] '%v' %v, 2^%v:\n", index, row.left, row.rules, row.bits)
		}
		results := descend(context, 0, line, row.left, row.rules)

		if caching {
			fmt.Printf("Cache %v\n", context.Debug())
		}

		total += results
		fmt.Printf("%v\n", results)
		// break
	}
	fmt.Printf("Total is %v\n", total)
	fmt.Printf("Counter is %v\n", COUNTER)
}

func (puzzle *Puzzle) Part2X() {

	c := cli.New(os.Args)
	var g *Grid
	if c.Contains("-line") {
		data := c.GetStringOrDie("-line")
		g = NewGrid(data)
	} else {
		// g = NewGrid(REAL_DATA)
		g = NewGrid(TEST_DATA_2)
	}

	solution1 := 0
	solution2 := 0
	for _, row := range g.data {
		conditions := row.left
		rules := row.rules
		fmt.Printf("%v\n", row.line)

		// rules = eval(rules)
		solution1 += countup(conditions, rules)

		// conditions2 := conditions
		// rules2 := rules
		// for index := 0; index < 5; index++ {
		// 	conditions2 = fmt.Sprintf("%v?%v", conditions2, conditions)
		// 	rules2 = append(rules2, rules...)
		// }

		// solution2 += countup(conditions2, rules2)
	}

	fmt.Printf("P1 : %v\n", solution1)
	fmt.Printf("P2 : %v\n", solution2)
}

func (puzzle *Puzzle) Run() {
	// puzzle.Part1()
	// puzzle.Part2()
	puzzle.Part2X()

}
