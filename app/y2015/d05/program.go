package d05

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
	"github.com/simonski/cli"
)

/*
Doesn't He Have Intern-Elves For This?
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
	iday, _ := strconv.Atoi("05")
	p := Puzzle{year: iyear, day: iday, title: "Doesn't He Have Intern-Elves For This?"}
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
	s.DateStarted = "d05"
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
	splits := strings.Split(REAL_DATA, "\n")
	niceCount := 0
	naughtyCount := 0
	for _, line := range splits {
		if IsNice(line) {
			niceCount++
		} else {
			naughtyCount++
		}
	}
	fmt.Printf("Part 1 Nice count %v, naughty count %v\n", niceCount, naughtyCount)
}

func (puzzle *Puzzle) Part2() {
	splits := strings.Split(REAL_DATA, "\n")
	niceCount := 0
	naughtyCount := 0
	for _, line := range splits {
		if IsNice2(line) {
			niceCount++
		} else {
			naughtyCount++
		}
	}
	fmt.Printf("Part 2 Nice count %v, naughty count %v\n", niceCount, naughtyCount)
}

func IsNice(line string) bool {
	if strings.Index(line, "ab") > -1 || strings.Index(line, "cd") > -1 || strings.Index(line, "pq") > -1 || strings.Index(line, "xy") > -1 {
		return false
	}

	vowelless := strings.ReplaceAll(line, "a", "")
	vowelless = strings.ReplaceAll(vowelless, "e", "")
	vowelless = strings.ReplaceAll(vowelless, "i", "")
	vowelless = strings.ReplaceAll(vowelless, "o", "")
	vowelless = strings.ReplaceAll(vowelless, "u", "")

	diff := len(line) - len(vowelless)
	if diff < 3 {
		return false
	}

	doubleCount := 0
	for iletter := 97; iletter < 97+26; iletter++ {
		c := string(rune(iletter))
		c += c
		if strings.Index(line, c) > -1 {
			// fmt.Printf("%v in %v? : yes!\n", c, line)
			doubleCount++
			// contains a double
		} else {
			// fmt.Printf("%v in %v? : no!\n", c, line)

		}
	}

	if doubleCount > 0 {
		return true
	} else {
		return false
	}

}

func IsNice2(line string) bool {

	/*
			It contains a pair of any two letters that appears at least twice in the string without overlapping, like xyxy (xy) or aabcdefgaa (aa), but not like aaa (aa, but it overlaps).
		It contains at least one letter which repeats with exactly one letter between them, like xyx, abcdefeghi (efe), or even aaa.
	*/

	foundDouble := false
	for iletter1 := 97; iletter1 < 97+26; iletter1++ {
		l1 := string(rune(iletter1))
		for iletter2 := 97; iletter2 < 97+26; iletter2++ {
			l2 := string(rune(iletter2))
			l := l1 + l2
			s := strings.ReplaceAll(line, l, "")
			if len(line)-len(s) >= 4 {
				foundDouble = true
				break
			}
		}
		if foundDouble {
			break
		}
	}

	foundRegex := false
	for iletter1 := 97; iletter1 < 97+26; iletter1++ {
		l1 := string(rune(iletter1))
		pattern := l1 + "." + l1
		match, _ := regexp.MatchString(pattern, line)
		if match {
			foundRegex = true
			break
		}
	}

	if foundDouble && foundRegex {
		return true
	}
	return false

}
