package d8

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
Day 8: Haunted Wasteland
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
	s.DateStarted = "2023-12-08 06:52:12"
	s.DateCompleted = "2023-12-08 10:52:12"
	return s
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("2023")
	iday, _ := strconv.Atoi("8")
	p := Puzzle{year: iyear, day: iday, title: "Day 8: Haunted Wasteland"}
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
}

func (puzzle *Puzzle) Part2() {
	puzzle.Load(REAL_DATA)
	P2X_X()
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}

func P2X() {

	// foreach
	// first step to Z is the "offset"
	// NEXT step to Z should be the repeating
	// so I think it is lcm + offsets as the number of steps?

	_, d := Load(REAL_DATA)
	repeats := make([]uint, 0)
	offsets := make([]uint, 0)
	for name, c := range d.Choices {
		if c.EndA {
			p1 := NewP2(name, REAL_DATA, false)
			fmt.Printf("%v -> ?.\n", name)
			counter, circularMap := p1.StepToEndingZ()
			for key, value := range circularMap {
				fmt.Printf("%v: %v\n", key, value)
			}
			fmt.Println("")
			moveCount := p1.MoveCount
			offset := moveCount // (p1.Moves)
			offsets = append(offsets, uint(offset))

			fmt.Printf("(%v), %v -> %v, %v moves to get to first Z.\n", counter, name, p1.Position, moveCount) //len(p1.Moves))

			// p2 := NewP2(p1.Position, REAL_DATA, true)
			// var p2 *P2
			p2 := p1
			p2.Directions = p1.Directions
			p2.Instructions = p1.Instructions
			p2.Position = p1.Position
			p1.Moves = make([]string, 0)
			p1.MoveCount = 0
			counter2, circularMap := p1.StepToEndingZ()
			for key, value := range circularMap {
				fmt.Printf("%v: %v\n", key, value)
			}
			fmt.Println("")
			fmt.Println("")

			repeat := p1.MoveCount
			fmt.Printf("(%v), %v -> %v, %v moves to get to second Z.\n", counter2, p1.Position, p1.Position, moveCount)

			repeats = append(repeats, uint(repeat))
			// break
		}
	}

	for _, offset := range offsets {
		fmt.Printf("offset: %v\n", offset)
	}
	lcm_value := utils.Lcm_x(repeats)
	fmt.Printf("lcm(above)=%v\n", lcm_value)

	final_value := lcm_value
	for _, offset := range offsets {
		final_value += offset
	}

	fmt.Printf("final = %v\n", final_value)

	// 219111877451 too low
	// 219111963718 too low
	// 219111963718
	// fmt.Printf("moves are %v\n", value)

	// fmt.Printf("LCM(%v,%v)=%v\n", s1, s2, utils.Lcm(s1, s2))
	// t.Fatalf("x")

}

func P2X_X() {

	// foreach
	// first step to Z is the "offset"
	// NEXT step to Z should be the repeating
	// so I think it is lcm + offsets as the number of steps?

	graph, instructions := LoadGraph(REAL_DATA)
	a_nodes := graph.FindA()
	all := make([]uint, 0)
	for _, a_node := range a_nodes {
		fmt.Printf("%v -> ?\n", a_node.Name)
		node, steps := graph.FindZ(a_node.Name, instructions)
		all = append(all, uint(steps))
		fmt.Printf("%v -> %v, %v steps.\n", a_node.Name, node.Name, steps)
		instructions.Reset()
	}

	lcm := utils.Lcm_x(all)
	fmt.Printf("LCM: %v\n", lcm)
	// instructions, directions := Load(REAL_DATA)

	// 219111877451 too low
	// 219111963718 too low
	// 219111963718
	// fmt.Printf("moves are %v\n", value)

	// fmt.Printf("LCM(%v,%v)=%v\n", s1, s2, utils.Lcm(s1, s2))
	// t.Fatalf("x")

}
