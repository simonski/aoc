package d15

import (
	"fmt"
	"strconv"

	"github.com/simonski/aoc/utils"
)

/*
Day 15: Warehouse Woes
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
	s.DateStarted = "2024-12-21 17:46:55"
	return s
}

func NewPuzzleWithData() *Puzzle {
	iyear, _ := strconv.Atoi("2024")
	iday, _ := strconv.Atoi("15")
	p := Puzzle{year: iyear, day: iday, title: "Day 15: Warehouse Woes"}
	return &p
}

func NewPuzzle() *Puzzle {
	return NewPuzzleWithData()
}

func (puzzle *Puzzle) Load(input string) {
}

func (puzzle *Puzzle) Part1() {
	// grid, instructions := Parse(TEST_DATA_00, TEST_DATA_01)
	// grid, instructions := Parse(TEST_DATA_1, TEST_DATA_2)
	grid, instructions := Parse(REAL_DATA_1, REAL_DATA_2)
	fmt.Printf("Start>>>, robot(%v,%v)\n", grid.robot.x, grid.robot.y)
	fmt.Println(grid.debug())
	fmt.Println("")
	for index := 0; index < len(instructions); index++ {
		instruction := instructions[index : index+1]
		fmt.Printf("[%v] %v, robot=%v,%v\n", index+1, instruction, grid.robot.x, grid.robot.y)
		grid.execute(instruction)
		fmt.Println(grid.debug())
		fmt.Println()
	}
	fmt.Printf("GPS: %v\n", grid.gps())
}

func (puzzle *Puzzle) Part2() {
	// grid, instructions := ParseP2(TEST_DATA_00, TEST_DATA_01)
	grid, instructions := ParseP2(TEST_DATA_20, TEST_DATA_21)

	// grid, instructions := ParseP2(TEST_DATA_1, TEST_DATA_2)
	// grid, instructions := ParseP2(REAL_DATA_1, REAL_DATA_2)
	fmt.Printf("Start>>>, robot(%v,%v)\n", grid.robot.x, grid.robot.y)
	fmt.Println(grid.debug())
	fmt.Println("")
	for index := 0; index < len(instructions); index++ {
		instruction := instructions[index : index+1]
		fmt.Printf("[%v] %v, robot=%v,%v\n", index+1, instruction, grid.robot.x, grid.robot.y)
		grid.execute(instruction)
		fmt.Println(grid.debug())
		fmt.Println()
	}
	fmt.Printf("GPS: %v\n", grid.gps())
}

func (puzzle *Puzzle) Run() {
	// puzzle.Part1()
	puzzle.Part2()
}
