package d6

import (
	"fmt"
	"strings"
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
	p := Puzzle{year: "2022", day: "06", title: "Tuning Trouble"}
	p.Load(input)
	return &p
}

func NewPuzzle() *Puzzle {
	return NewPuzzleWithData(REAL_DATA)
}

func (p *Puzzle) NextPacketMarker(buffer_size int) int {
	length := len(p.input)
	for index := 0; index < length; index++ {
		if index+buffer_size > length {
			break
		}
		counts := make(map[string]int)
		slice := p.input[index : index+buffer_size]
		fmt.Printf("slice[%v:%v] = %v\n", index, index+buffer_size, slice)

		for i := 0; i < len(slice); i++ {
			c := slice[i : i+1]
			fmt.Printf("%v[%v] = %v\n", slice, i, c)
			v := counts[c]
			v += 1
			counts[c] = v
		}
		success := true
		for _, v := range counts {
			if v != 1 {
				success = false
				break
			}
		}
		if success {
			return index + buffer_size
		}

	}
	return -1
}

func (puzzle *Puzzle) Load(input string) {
	lines := strings.Split(input, "\n")
	puzzle.input = input
	puzzle.lines = lines
}

func (puzzle *Puzzle) Part1() {
	puzzle.Load(REAL_DATA)
	fmt.Printf("202206P1: %v\n", puzzle.NextPacketMarker(4))
}

func (puzzle *Puzzle) Part2() {
	puzzle.Load(REAL_DATA)
	fmt.Printf("202206P2: %v\n", puzzle.NextPacketMarker(14))
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}
