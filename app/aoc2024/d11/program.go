package d11

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
Day 11: Plutonian Pebbles
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
	s.DateStarted = "2024-12-11 07:31:41"
	return s
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("2024")
	iday, _ := strconv.Atoi("11")
	p := Puzzle{year: iyear, day: iday, title: "Day 11: Plutonian Pebbles"}
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
	puzzle.p1(REAL_DATA)
	// puzzle.p1(REAL_DATA)
}

func (puzzle *Puzzle) p1(data string) {
	values := strings.Split(data, " ")
	fmt.Printf("%v: %v %v\n", 0, len(values), values)
	for count := 0; count < 75; count++ {
		values = p1_blink(values)
		if len(values) < 10 {
			fmt.Printf("%v: %v %v\n", count+1, len(values), values)
		} else {
			fmt.Printf("%v: %v\n", count+1, len(values))
		}
	}
}

func p1_blink(values []string) []string {
	result := make([]string, 0)
	for index := 0; index < len(values); index++ {
		value := values[index]
		if value == "0" {
			// rule 1
			value = strings.ReplaceAll(value, "0", "1")
			result = append(result, value)
		} else if len(value)%2 == 0 {
			// rule 2
			lpart := value[0 : len(value)/2]
			rpart := value[len(value)/2:]
			lpart_int, _ := strconv.Atoi(lpart)
			rpart_int, _ := strconv.Atoi(rpart)
			lpart = fmt.Sprintf("%v", lpart_int)
			rpart = fmt.Sprintf("%v", rpart_int)
			result = append(result, lpart)
			result = append(result, rpart)
		} else {
			// rule 3
			value_int, _ := strconv.Atoi(value)
			value_int *= 2024
			value = fmt.Sprintf("%v", value_int)
			result = append(result, value)
		}
	}
	return result
}

func (puzzle *Puzzle) p2(data string) {
}

func (puzzle *Puzzle) Part2() {
	puzzle.Load(REAL_DATA)
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}
