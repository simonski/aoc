package d17

import (
	"fmt"
	"os"
	"strings"

	"github.com/simonski/cli"
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
}

func NewPuzzleWithData(input string) *Puzzle {
	p := Puzzle{year: "2022", day: "17", title: "Pyroclastic Flow"}
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

var FLOOR_SIZE = 1000

func (puzzle *Puzzle) Part1() {
	cli := cli.New(os.Args)
	log_level := 0
	if cli.Contains("-v") {
		log_level = 1
	} else if cli.Contains("-vv") {
		log_level = 2
	}
	resetFloor := cli.Contains("-floor")
	if os.Args[5] == "test" {
		c := NewChamber(TEST_DATA, log_level)
		// fmt.Println(c.Debug())
		c.RunPart1(2022, resetFloor, FLOOR_SIZE)
		fmt.Println(c.Debug())
		fmt.Printf("Rock Count %v, Height is %v\n", len(c.Rocks), c.Height)

	} else if os.Args[5] == "real" {
		c := NewChamber(REAL_DATA, log_level)
		c.RunPart1(2022, resetFloor, FLOOR_SIZE)
		fmt.Println(c.Debug())
		fmt.Printf("Rock Count %v, Height is %v\n", len(c.Rocks), c.Height)

	}
}

func (puzzle *Puzzle) Part2_a() {
	cl := cli.New(os.Args)
	c := NewChamber(REAL_DATA, 0)
	if cl.Contains("-v") {
		c.LOG_LEVEL = 1
	} else if cl.Contains("-vv") {
		c.LOG_LEVEL = 2
	}
	// c.AddRocks(5000)
	rocks := cl.GetIntOrDefault("-rocks", 5000)
	key_size := cl.GetIntOrDefault("-key_size", 1615)
	first_index, key_size := c.Part2_FindFirstKey(rocks, key_size)
	if first_index > -1 {
		fmt.Printf("First index of key size %v is %v\n", key_size, first_index)
	}

	fmt.Println("")

	// c2 := NewChamber(REAL_DATA, 0)
	// c2.Part2_VerifySequences(5000, first_index, key_size)

	// fmt.Printf("Rock Count %v, Height is %v\n", len(c.Rocks), c.Height)
}

func (puzzle *Puzzle) Part2() {

	cl := cli.New(os.Args)
	rocks := cl.GetIntOrDefault("-rocks", 1000)
	max_key_size := cl.GetIntOrDefault("-max_key_size", 100)
	min_key_size := cl.GetIntOrDefault("-min_key_size", 20)
	fmt.Printf("-rocks=%v, -max_key_size=%v, -min_key_size=%v\n", rocks, max_key_size, min_key_size)
	c2 := NewChamber(REAL_DATA, 0)
	if cl.Contains("-v") {
		c2.LOG_LEVEL = 1
	} else if cl.Contains("-vv") {
		c2.LOG_LEVEL = 2
	}
	c2.AddRocks(rocks)
	fmt.Println(c2.Debug())
	c2.Part2_FindSequences(rocks, max_key_size, min_key_size)

	// if os.Args[5] == "test" {
	// 	c := NewChamber(TEST_DATA, 0)
	// 	fmt.Println(c.Debug())
	// 	// c.RunPart2(1000000000000, true, FLOOR_SIZE)
	// 	fmt.Printf("Rock Count %v, Height is %v\n", len(c.Rocks), c.Height)
	// } else if os.Args[5] == "real" {
	// 	c := NewChamber(REAL_DATA, 0)
	// 	fmt.Println(c.Debug())
	// 	// c.RunPart2(1000000000000, true, FLOOR_SIZE)
	// 	fmt.Printf("Rock Count %v, Height is %v\n", len(c.Rocks), c.Height)
	// }
}

func (puzzle *Puzzle) Run() {
	USAGE := "Usage: aoc run 2022 17 P1|P2 test|real (-v)"
	if len(os.Args) < 6 {
		fmt.Println(USAGE)
		os.Exit(1)
	} else if os.Args[4] == "P1" {
		puzzle.Part1()
	} else if os.Args[4] == "P2" {
		puzzle.Part2()
	} else {
		fmt.Println(USAGE)
		os.Exit(1)
	}
}
