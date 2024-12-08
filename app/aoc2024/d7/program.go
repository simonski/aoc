package d7

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
Day 7: Bridge Repair
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
	s.DateStarted = "2024-12-07 08:15:09"
	return s
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("2024")
	iday, _ := strconv.Atoi("7")
	p := Puzzle{year: iyear, day: iday, title: "Day 7: Bridge Repair"}
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
	puzzle.p1_test(TEST_DATA)
	puzzle.p1_test(REAL_DATA)
}

func (puzzle *Puzzle) p1_test(data string) {
	puzzle.Load(data)
	total := 0
	for _, line := range puzzle.lines {
		splits := strings.Split(line, ":")
		target, _ := strconv.Atoi(splits[0])
		v := strings.Trim(splits[1], " ")
		values := utils.SplitDataToListOfInts(v, " ")
		if isValid(target, values) {
			total += target
		} else {
		}
	}
	fmt.Println(total)
}

func isValid(target int, values []int) bool {
	return descend(values[0], target, 1, values)
}

func descend(total int, target int, depth int, values []int) bool {
	value := values[depth]
	add_total := total + value
	mul_total := total * value
	// fmt.Printf("total=%v, target=%v, value=%v, depth=%v/%v, values=%v, add_total=%v, mul_total=%v\n", total, target, value, depth, len(values), values, add_total, mul_total)
	if depth+1 == len(values) {
		if add_total == target || mul_total == target {
			return true
		} else {
			return false
		}
	} else {
		new_depth := depth + 1
		a := descend(add_total, target, new_depth, values)
		b := descend(mul_total, target, new_depth, values)
		return a || b
	}

}

func (puzzle *Puzzle) Part2() {
	puzzle.p2_test(TEST_DATA)
	puzzle.p2_test(REAL_DATA)
}

func (puzzle *Puzzle) p2_test(data string) {
	puzzle.Load(data)
	total := 0
	for _, line := range puzzle.lines {
		splits := strings.Split(line, ":")
		target, _ := strconv.Atoi(splits[0])
		v := strings.Trim(splits[1], " ")
		values := utils.SplitDataToListOfInts(v, " ")
		if isValidP2(target, values) {
			total += target
		}
		// if isValid(target, values) {
		// 	total += target
		// }
	}
	fmt.Println(total)
}

func isValidP2(target int, values []int) bool {
	return descendP2(values[0], target, 1, values)
}

func descendP2(total int, target int, depth int, values []int) bool {
	value := values[depth]
	cc_total_str := fmt.Sprintf("%v%v", total, value)
	cc_total, _ := strconv.Atoi(cc_total_str)
	add_total := total + value
	mul_total := total * value

	// prefix := strings.Repeat(" ", depth)
	// fmt.Printf(prefix+"total=%v, target=%v, value=%v, depth=%v/%v, values=%v, add_total=%v, mul_total=%v, cc_total=%v\n", total, target, value, depth, len(values), values, add_total, mul_total, cc_total)
	if depth+1 == len(values) {
		if cc_total == target || add_total == target || mul_total == target {
			return true
		} else {
			return false
		}
	} else {
		new_depth := depth + 1
		a := descendP2(add_total, target, new_depth, values)
		b := descendP2(mul_total, target, new_depth, values)
		c := descendP2(cc_total, target, new_depth, values)
		return a || b || c
	}

}

func (puzzle *Puzzle) Run() {
	// puzzle.Part1()
	// fmt.Println()
	puzzle.Part2()
}
