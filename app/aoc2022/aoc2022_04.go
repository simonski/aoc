package aoc2022

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
	goutils "github.com/simonski/goutils"
)

/*
--- Day 04: Description ---

*/

func (app *Application) Y2022D04_Summary() *utils.Summary {
	s := &utils.Summary{}
	s.Name = "Camp Cleanup"
	s.Year = 2022
	s.Day = 04
	s.Entries = make([]*utils.Entry, 0)

	s.ProgressP1 = utils.Completed
	s.ProgressP2 = utils.Completed

	entry := &utils.Entry{}
	entry.Date = "2022-12-04"
	entry.Title = "TDD to the rescue."
	entry.Notes = `Day four, 7am, 38mins. A visualisation is possible! I will try to go quickly, solviing with some TDD, present an ascii output then once complete go back and do a webpage vis.  Some notes:

	<p>
	I did a silly thing working out if it intersects so I started again. Not super keen on the <code>.Overlap()</code> and <code>.Intesect()</code> logic but it works and I'll live.
	</p>

	<p>
	TDD totally helped in solving it first time.
	
	<p>
	The visualisation is just the stdout of all as html - keep it simple.

	<p>
	As always, the tools (this blogging ahem, <i>engine></i> is now interesting to me).
	<ul>
	<li>code formatting or referencing somehow
	<li>inlining html in a sensible manner
	<li>putting teh blog entries into a separate file that is go:embed read at startup rather than using go structs and super-hardcoding the blog itself
	`

	entry.Summary = s
	s.Entries = append(s.Entries, entry)

	return s

}

type Day4Task struct {
	input string
	low   int
	high  int
}

func (t *Day4Task) Debug(max_length int, debug_string string) string {
	line := ""
	for index := 0; index <= t.low; index++ {
		line = fmt.Sprintf("%v%v", line, ".")
	}
	for index := t.low; index <= t.high; index++ {
		s := ""
		if debug_string == "" {
			s = fmt.Sprintf("%v", index)
		} else {
			s = debug_string
		}
		line = fmt.Sprintf("%v%v", line, s)
	}
	for index := t.high; index <= max_length; index++ {
		line = fmt.Sprintf("%v%v", line, ".")
	}
	return line
}

func (t1 *Day4Task) Overlap(t2 *Day4Task) bool {
	return t1.low >= t2.low && t1.high <= t2.high
}

func (t1 *Day4Task) Intersect(t2 *Day4Task) bool {
	if t1.low >= t2.low && t1.low <= t2.high {
		return true
	}
	if t1.high >= t2.low && t1.high <= t2.high {
		return true
	}
	return false
}

func NewDay4Task(input string) *Day4Task {
	splits := strings.Split(input, "-")
	low, _ := strconv.Atoi(splits[0])
	high, _ := strconv.Atoi(splits[1])
	t := Day4Task{input: input, low: low, high: high}
	return &t
}

type Day4Pair struct {
	input string
	task1 *Day4Task
	task2 *Day4Task
}

func (p *Day4Pair) Overlap() bool {
	return p.task1.Overlap(p.task2) || p.task2.Overlap(p.task1)
}

func (p *Day4Pair) Intersect() bool {
	return p.task1.Intersect(p.task2) || p.task2.Intersect(p.task1)
}

func NewDay4Pair(input string) *Day4Pair {
	splits := strings.Split(input, ",")
	t1 := NewDay4Task(splits[0])
	t2 := NewDay4Task(splits[1])
	p := Day4Pair{input: input, task1: t1, task2: t2}
	return &p
}

type Day4 struct {
	input string
	Pairs []*Day4Pair
}

func NewDay4(input string) *Day4 {
	splits := strings.Split(input, "\n")
	pairs := make([]*Day4Pair, 0)
	for _, pinput := range splits {
		p := NewDay4Pair(pinput)
		pairs = append(pairs, p)
	}
	return &Day4{input: input, Pairs: pairs}
}

func (app *Application) Y2022D04P1() {
	input := DAY_2022_04_DATA
	day4 := NewDay4(input)
	overlaps := 0
	for _, pair := range day4.Pairs {
		if pair.Overlap() {
			overlaps += 1
		}
	}
	fmt.Printf("2022-04 Part 1: Overlaps: %v\n", overlaps)

	max := 0
	for _, pair := range day4.Pairs {
		max = goutils.Max(max, pair.task1.high)
		max = goutils.Max(max, pair.task2.high)
	}
	max += 1
	fmt.Printf("max is %v\n", max)
	for _, pair := range day4.Pairs {
		fmt.Println(pair.task1.Debug(max, "X"))
	}

}

// // rename this to the year and day in question
// func (app *Application) Y2022D04P2_inprogress() {
// }

// rename and uncomment this to the year and day in question once complete for a gold star!
// func (app *Application) Y20XXDXXP1Render() {
// }

// rename and uncomment this to the year and day in question once complete for a gold star!
// func (app *Application) Y20XXDXXP2Render() {
// }

// this is what we will reflect and call - so both parts with run. It's up to you to make it print nicely etc.
// The app reference has a CLI for logging.
// func (app *Application) Y2022D04() {
// 	app.Y2022D04P1_inprogress()
// 	app.Y2022D04P2_inprogress()
// }
