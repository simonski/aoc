package d01

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
	"github.com/simonski/cli"
)

/*
--- Day 01: Trebuchet?! ---
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
	return s
}

func NewPuzzleWithData(input string) *Puzzle {
	p := Puzzle{year: 2023, day: 1, title: "Trebuchet?!"}
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

func (puzzle *Puzzle) Part1(data string) {
	puzzle.Load(data)
	p := NewPuzzleWithData(data)

	fmt.Printf("There are %v lines.\n", len(p.lines))

	grandTotal := 0
	for _, line := range p.lines {
		lvar := -1
		rvar := -1
		for index := 0; index < len(line)-1; index++ {
			candidate := line[index : index+1]
			i, isInt := IsInt(candidate)
			if isInt {
				lvar = i
				break
			}
		}

		for index := len(line) - 1; index >= 0; index-- {
			candidate := line[index : index+1]
			i, isInt := IsInt(candidate)
			if isInt {
				rvar = i
				break
			}
		}

		if lvar == -1 {
			lvar = rvar
		}
		if rvar == -1 {
			rvar = lvar
		}

		total := fmt.Sprintf("%v%v", lvar, rvar)
		sum, _ := strconv.Atoi(total)
		grandTotal += sum
		fmt.Printf("Line: %v, lvar: %v, rvar: %v, sum: %v\n", line, lvar, rvar, sum)
	}
	fmt.Printf("GrantTotal: %v\n", grandTotal)

}

type NumberThing struct {
	Value  string
	Number int
}

func (puzzle *Puzzle) Part2(data string) {
	puzzle.Load(data)
	p := NewPuzzleWithData(data)

	output := make([]string, 0)
	for _, original := range p.lines {
		line := EdgeCases(original)
		line = strings.ReplaceAll(line, "zero", "0")
		line = strings.ReplaceAll(line, "one", "1")
		line = strings.ReplaceAll(line, "two", "2")
		line = strings.ReplaceAll(line, "three", "3")
		line = strings.ReplaceAll(line, "four", "4")
		line = strings.ReplaceAll(line, "five", "5")
		line = strings.ReplaceAll(line, "six", "6")
		line = strings.ReplaceAll(line, "seven", "7")
		line = strings.ReplaceAll(line, "eight", "8")
		line = strings.ReplaceAll(line, "nine", "9")
		line = strings.ReplaceAll(line, "ten", "10")
		output = append(output, line)
		fmt.Printf("%v -> %v\n", original, line)
	}
	new_data := strings.Join(output, "\n")

	puzzle.Part1(new_data)

}

func (puzzle *Puzzle) Run() {
	// USAGE := "Usage: aoc run 2023 01 (P1|P2)"
	c := cli.New(os.Args)
	if c.Contains("P1") {
		puzzle.Part1(REAL_DATA)
	} else if c.Contains("P2") {
		puzzle.Part2(REAL_DATA)
	} else {
		puzzle.Part2(REAL_DATA)
		puzzle.Part2(TEST_DATA_2)
	}

}

func IsInt(candidate string) (int, bool) {
	i, err := strconv.Atoi(candidate)
	return i, err == nil
}

func EdgeCases(line string) string {
	line = strings.ReplaceAll(line, "oneight", "one_eight")
	line = strings.ReplaceAll(line, "twone", "two_one")
	line = strings.ReplaceAll(line, "threeight", "three_eight")
	line = strings.ReplaceAll(line, "fiveight", "five_eight")
	line = strings.ReplaceAll(line, "sevenine", "seven_nine")
	line = strings.ReplaceAll(line, "eightwo", "eight_two")
	line = strings.ReplaceAll(line, "nineight", "nine_eight")
	return line
}

func FirstWord(line string) (string, int) {

	words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	minIndex := 1000
	// firstIndex := -1
	// firstWord := 0
	minWord := ""
	minWordIndex := -1
	// idx := 0
	for index, word := range words {
		wordIndex := strings.Index(line, word)
		if wordIndex > -1 && wordIndex < minIndex {
			minIndex = wordIndex
			minWord = word
			minWordIndex = index + 1

			// idx = index
		}
	}
	return minWord, minWordIndex

}

func IsMin(candidate int, values ...int) bool {
	for _, v := range values {
		if v < candidate {
			return false
		}
	}
	return true
}
