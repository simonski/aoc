package d3

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
Day 3: Mull It Over
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
	s.DateStarted = "2024-12-03 06:29:08"
	return s
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("2024")
	iday, _ := strconv.Atoi("3")
	p := Puzzle{year: iyear, day: iday, title: "Day 3: Mull It Over"}
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
	// puzzle.Load(TEST_DATA)
	puzzle.Load(REAL_DATA)

	total := 0
	for _, line := range puzzle.lines {
		fmt.Println(line)
		re := regexp.MustCompile("mul\\(\\d+,\\d+\\)")
		results := re.FindAllString(line, -1)
		for index, r := range results {
			r = strings.ReplaceAll(r, "mul", "")
			r = strings.ReplaceAll(r, "(", "")
			r = strings.ReplaceAll(r, ")", "")
			splits := strings.Split(r, ",")
			v1, _ := strconv.Atoi(splits[0])
			v2, _ := strconv.Atoi(splits[1])
			value := v1 * v2
			total += value
			fmt.Printf("[%v/%v] %v, %v*%v=%v, total=%v\n", index, len(results), r, v1, v2, value, total)

		}
	}
	fmt.Printf("total=%v\n", total)

}

func (puzzle *Puzzle) Part2() {
	// puzzle.Load(TEST_DATA)
	puzzle.Load(REAL_DATA)

	total := 0
	enabled := true
	for _, line := range puzzle.lines {
		fmt.Println(line)
		re := regexp.MustCompile("mul\\(\\d+,\\d+\\)|do\\(\\)|don\\'t\\(\\)")
		results := re.FindAllString(line, -1)
		for index, r := range results {
			if r == "do()" {
				enabled = true
				continue
			} else if r == "don't()" {
				enabled = false
				continue
			}

			if !enabled {
				continue
			}

			fmt.Println(r)
			r = strings.ReplaceAll(r, "mul", "")
			r = strings.ReplaceAll(r, "(", "")
			r = strings.ReplaceAll(r, ")", "")
			splits := strings.Split(r, ",")
			v1, _ := strconv.Atoi(splits[0])
			v2, _ := strconv.Atoi(splits[1])
			value := v1 * v2
			total += value
			fmt.Printf("[%v/%v] %v, %v*%v=%v, total=%v\n", index, len(results), r, v1, v2, value, total)

		}
	}
	fmt.Printf("total=%v\n", total)
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}
