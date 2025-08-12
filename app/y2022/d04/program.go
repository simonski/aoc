package d04

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/goutils"
)

/*
--- Day 05:  ---

*/

type Puzzle struct {
	title string
	year  string
	day   string
	input string
	lines []string
	Pairs []*Pair
}

func NewPuzzle(input string) Puzzle {
	p := Puzzle{year: "2022", day: "04", title: "Camp Cleanup"}
	p.Load(input)
	return p
}

func (puzzle *Puzzle) Load(input string) {
	lines := strings.Split(input, "\n")
	puzzle.input = input
	puzzle.lines = lines

	pairs := make([]*Pair, 0)
	for _, pinput := range puzzle.lines {
		p := NewPair(pinput)
		pairs = append(pairs, p)
	}
	puzzle.Pairs = pairs

}

func (puzzle *Puzzle) Part2() {
	puzzle.Load(REAL_DATA)
}

func (puzzle Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}

func (puzzle *Puzzle) Part1() {
	puzzle.Load(REAL_DATA)
	overlaps := 0
	for _, pair := range puzzle.Pairs {
		if pair.Overlap() {
			overlaps += 1
		}
	}
	fmt.Printf("2022-04 Part 1: Overlaps: %v\n", overlaps)

	max := 0
	for _, pair := range puzzle.Pairs {
		max = goutils.Max(max, pair.task1.high)
		max = goutils.Max(max, pair.task2.high)
	}
	max += 1
	fmt.Printf("max is %v\n", max)
	for _, pair := range puzzle.Pairs {
		fmt.Println(pair.task1.Debug(max, "X"))
	}

}

type Task struct {
	input string
	low   int
	high  int
}

func (t *Task) Debug(max_length int, debug_string string) string {
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

func (t1 *Task) Overlap(t2 *Task) bool {
	return t1.low >= t2.low && t1.high <= t2.high
}

func (t1 *Task) Intersect(t2 *Task) bool {
	if t1.low >= t2.low && t1.low <= t2.high {
		return true
	}
	if t1.high >= t2.low && t1.high <= t2.high {
		return true
	}
	return false
}

func NewTask(input string) *Task {
	splits := strings.Split(input, "-")
	low, _ := strconv.Atoi(splits[0])
	high, _ := strconv.Atoi(splits[1])
	t := Task{input: input, low: low, high: high}
	return &t
}

type Pair struct {
	input string
	task1 *Task
	task2 *Task
}

func (p *Pair) Overlap() bool {
	return p.task1.Overlap(p.task2) || p.task2.Overlap(p.task1)
}

func (p *Pair) Intersect() bool {
	return p.task1.Intersect(p.task2) || p.task2.Intersect(p.task1)
}

func NewPair(input string) *Pair {
	splits := strings.Split(input, ",")
	t1 := NewTask(splits[0])
	t2 := NewTask(splits[1])
	p := Pair{input: input, task1: t1, task2: t2}
	return &p
}
