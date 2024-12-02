package d2

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
Day 2: Red-Nosed Reports
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
	s.DateStarted = "2024-12-01 16:13:01"
	return s
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("2024")
	iday, _ := strconv.Atoi("2")
	p := Puzzle{year: iyear, day: iday, title: "Day 2: Red-Nosed Reports"}
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

type Report struct {
	Levels []int
}

func (r *Report) IsAllDecrease() bool {
	last := math.MaxInt
	failCount := 0
	for _, v := range r.Levels {
		if v > last {
			failCount += 1
			return false
		}
		last = v
	}
	return failCount == 0
}

func (r *Report) IsAllIncrease() bool {
	last := 0
	failCount := 0
	for _, v := range r.Levels {
		if v < last {
			failCount += 1
		}
		last = v

	}
	return failCount == 0
}

func (r *Report) Permutations() []*Report {
	reports := make([]*Report, 0)
	// fmt.Println(r.Levels)
	for index := range r.Levels {
		// fmt.Printf("index=%v\n", index)
		report := Report{}
		levels := make([]int, 0)
		for i, v := range r.Levels {
			// fmt.Printf("i=%v, v=%v\n", i, v)
			if i == index {
				continue
			}
			levels = append(levels, v)
			reports = append(reports, &report)
		}
		report.Levels = levels
		// copy(r.Levels, levels)
	}
	return reports
}

func (r *Report) Difference(min int, max int) bool {
	failCount := 0
	for index, value1 := range r.Levels {
		if index+1 == len(r.Levels) {
			break
		}
		value2 := r.Levels[index+1]
		diff := value1 - value2
		if diff < 0 {
			diff *= -1
		}
		if diff >= min && diff <= max {
			// ok
		} else {
			failCount += 1
		}
	}
	return failCount == 0
}

func (r *Report) IsSafe() bool {
	allIncrease := r.IsAllIncrease()
	allDecrease := r.IsAllDecrease()
	differOk := r.Difference(1, 3)
	return (allIncrease || allDecrease) && differOk
}

type Data struct {
	Reports []*Report
}

func NewData(lines []string) *Data {
	data := Data{}
	data.Load(lines)
	return &data
}

func (d *Data) Load(lines []string) {
	d.Reports = make([]*Report, 0)
	for _, line := range lines {
		report := NewReport(line)
		d.Reports = append(d.Reports, report)
	}
}

func NewReport(line string) *Report {
	r := Report{}
	r.Levels = make([]int, 0)
	splits := strings.Split(line, " ")
	for _, sv := range splits {
		v, _ := strconv.Atoi(sv)
		r.Levels = append(r.Levels, v)
	}
	// sort.Ints(r.Levels)
	return &r
}

func (puzzle *Puzzle) Part1() {
	puzzle.Load(REAL_DATA)
	// puzzle.Load(TEST_DATA)
	data := NewData(puzzle.lines)
	safeCount := 0
	for _, report := range data.Reports {
		if report.IsSafe() {
			fmt.Printf("%v [safe]\n", report.Levels)
			safeCount += 1
		} else {
			fmt.Printf("%v [unsafe]\n", report.Levels)
		}
	}

	fmt.Printf("Safe count: %v\n", safeCount)
}

func (puzzle *Puzzle) Part2() {
	puzzle.Load(REAL_DATA)
	// puzzle.Load(TEST_DATA)
	data := NewData(puzzle.lines)
	safeCount := 0
	for _, report := range data.Reports {
		if report.IsSafe() {
			fmt.Printf("%v [safe]\n", report.Levels)
			safeCount += 1
		} else {
			// remove 1 from the levels and try again
			foundSafe := false
			perms := report.Permutations()
			for _, r := range perms {
				if r.IsSafe() {
					safeCount += 1
					foundSafe = true
					break
				}
			}
			if foundSafe {
				fmt.Printf("%v [safe (with 1 level changed)]\n", report.Levels)
			} else {
				fmt.Printf("%v [unsafe]\n", report.Levels)
			}
		}
	}

	fmt.Printf("Safe count: %v\n", safeCount)
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}
