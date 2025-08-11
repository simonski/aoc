package d6

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
Day 6: Wait For It
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
	s.DateStarted = "2023-12-06 06:15:05"
	s.DateCompleted = "2023-12-06 06:49:05"
	return s
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("2023")
	iday, _ := strconv.Atoi("6")
	p := Puzzle{year: iyear, day: iday, title: "Day 6: Wait For It"}
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

	times1 := get_best_times(35, 213)
	times2 := get_best_times(69, 1168)
	times3 := get_best_times(68, 1086)
	times4 := get_best_times(87, 1248)

	result := len(times1) * len(times2) * len(times3) * len(times4)
	fmt.Printf("Part1: %v\n", result)

}

func (puzzle *Puzzle) Part2() {
	test_times := get_best_times(71530, 940200)
	real_times := get_best_times(35696887, 213116810861248)
	fmt.Printf("Part2: (test) : %v\n", len(test_times))
	fmt.Printf("Part2: (real) : %v\n", len(real_times))

}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}

func distance(downtime int, racetime int) int {
	return (racetime - downtime) * downtime
}

func get_best_times(racetime int, max_distance int) []int {
	results := make([]int, 0)
	for downtime := 0; downtime < racetime; downtime++ {
		d := distance(downtime, racetime)
		if d > max_distance {
			results = append(results, d)
		}
	}
	return results
}
