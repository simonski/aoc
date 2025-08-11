package d1

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
Day 1: Historian Hysteria
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
	s.DateStarted = "2024-12-01 15:30:49"
	return s
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("2024")
	iday, _ := strconv.Atoi("1")
	p := Puzzle{year: iyear, day: iday, title: "Day 1: Historian Hysteria"}
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

func (puzzle *Puzzle) LoadData() ([]int, []int, map[int]int) {
	v1list := make([]int, 0)
	v2list := make([]int, 0)

	v2map := make(map[int]int)

	for _, line := range puzzle.lines {
		line = strings.ReplaceAll(line, "   ", " ")
		splits := strings.Split(line, " ")
		fmt.Println(splits)
		v1, e1 := strconv.Atoi(splits[0])
		v2, e2 := strconv.Atoi(splits[1])

		v2map[v2] += 1

		if e1 != nil {
			panic(e1)
		}
		if e2 != nil {
			panic(e2)
		}
		v1list = append(v1list, v1)
		v2list = append(v2list, v2)
	}

	sort.Ints(v1list)
	sort.Ints(v2list)
	return v1list, v2list, v2map
}
func (puzzle *Puzzle) Part1() {
	puzzle.Load(REAL_DATA)

	v1list, v2list, v2map := puzzle.LoadData()

	fmt.Println(v2map)

	total := 0
	for index, v1 := range v1list {
		v2 := v2list[index]
		diff := v2 - v1
		fmt.Println(diff)
		if diff < 0 {
			diff *= -1
		}
		total += diff
	}

	fmt.Println(v1list)
	fmt.Println(v2list)
	fmt.Printf("total=%v\n", total)

	total_score := 0
	for _, v1 := range v1list {
		score := v1 * v2map[v1]
		total_score += score
		fmt.Printf("value=%v, frequency=%v, score=%v\n", v1, v2map[v1], score)

	}
	fmt.Printf("score=%v\n", total_score)

}

func (puzzle *Puzzle) Part2() {
	puzzle.Load(REAL_DATA)
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}
