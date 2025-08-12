package d05

import (
	"fmt"
	"strconv"
	"strings"
)

/*
--- Day 05:  ---

*/

type Puzzle struct {
	title        string
	year         string
	day          string
	input        string
	Instructions []string
	Stacks       []*Stack
}

func NewPuzzle(input string) *Puzzle {
	p := Puzzle{year: "2022", day: "05", title: "Supply Stacks"}
	p.Load(input)
	p.Stacks = make([]*Stack, 0)
	return &p
}

func NewPuzzleTestData() *Puzzle {
	p := NewPuzzle(TEST_INSTRUCTIONS)
	s1 := NewStack()
	s1.Push("Z").Push("N")
	s2 := NewStack()
	s2.Push("M").Push("C").Push("D")
	s3 := NewStack()
	s3.Push("P")
	p.AddStack(s1)
	p.AddStack(s2)
	p.AddStack(s3)
	return p
}

func NewPuzzleRealData() *Puzzle {

	// [P]     [C]         [M]
	// [D]     [P] [B]     [V] [S]
	// [Q] [V] [R] [V]     [G] [B]
	// [R] [W] [G] [J]     [T] [M]     [V]
	// [V] [Q] [Q] [F] [C] [N] [V]     [W]
	// [B] [Z] [Z] [H] [L] [P] [L] [J] [N]
	// [H] [D] [L] [D] [W] [R] [R] [P] [C]
	// [F] [L] [H] [R] [Z] [J] [J] [D] [D]
	//  1   2   3   4   5   6   7   8   9

	p := NewPuzzle(REAL_INSTRUCTIONS)
	s1 := NewStack()
	s1.Push("F").Push("H").Push("B").Push("V").Push("R").Push("Q").Push("D").Push("P")
	p.AddStack(s1)

	s2 := NewStack()
	s2.Push("L").Push("D").Push("Z").Push("Q").Push("W").Push("V")
	p.AddStack(s2)

	s3 := NewStack()
	s3.Push("H").Push("L").Push("Z").Push("Q").Push("G").Push("R").Push("P").Push("C")
	p.AddStack(s3)

	s4 := NewStack()
	s4.Push("R").Push("D").Push("H").Push("F").Push("J").Push("V").Push("B")
	p.AddStack(s4)

	s5 := NewStack()
	s5.Push("Z").Push("W").Push("L").Push("C")
	p.AddStack(s5)

	s6 := NewStack()
	s6.Push("J").Push("R").Push("P").Push("N").Push("T").Push("G").Push("V").Push("M")
	p.AddStack(s6)

	s7 := NewStack()
	s7.Push("J").Push("R").Push("L").Push("V").Push("M").Push("B").Push("S")
	p.AddStack(s7)

	s8 := NewStack()
	s8.Push("D").Push("P").Push("J")
	p.AddStack(s8)

	s9 := NewStack()
	s9.Push("D").Push("C").Push("N").Push("W").Push("V")
	p.AddStack(s9)

	return p
}

func (puzzle *Puzzle) AddStack(stack *Stack) {
	puzzle.Stacks = append(puzzle.Stacks, stack)
}

func (puzzle *Puzzle) Line() string {
	line := ""
	for _, stack := range puzzle.Stacks {
		line += stack.Peek()
	}
	return line
}

func (puzzle *Puzzle) RunInstruction(instruction string) {
	instruction = strings.ReplaceAll(instruction, "\n", "")
	if instruction == "" {
		return
	}
	splits := strings.Split(instruction, " ")
	// move 1 from 3 to 2
	count, _ := strconv.Atoi(splits[1])
	from, _ := strconv.Atoi(splits[3])
	to, _ := strconv.Atoi(splits[5])

	from -= 1
	to -= 1

	fromStack := puzzle.Stacks[from]
	toStack := puzzle.Stacks[to]
	for index := 0; index < count; index++ {
		value := fromStack.Pop()
		toStack.Push(value)
	}

}

func (puzzle *Puzzle) RunInstruction9001(instruction string) {
	instruction = strings.ReplaceAll(instruction, "\n", "")
	if instruction == "" {
		return
	}
	splits := strings.Split(instruction, " ")
	// move 1 from 3 to 2
	count, _ := strconv.Atoi(splits[1])
	from, _ := strconv.Atoi(splits[3])
	to, _ := strconv.Atoi(splits[5])

	from -= 1
	to -= 1

	fromStack := puzzle.Stacks[from]
	toStack := puzzle.Stacks[to]
	tempStack := NewStack()
	for index := 0; index < count; index++ {
		value := fromStack.Pop()
		tempStack.Push(value)
	}
	for index := 0; index < count; index++ {
		value := tempStack.Pop()
		toStack.Push(value)
	}

}

func (puzzle *Puzzle) RunInstructions9001() {
	for _, instruction := range puzzle.Instructions {
		puzzle.RunInstruction9001(instruction)
	}
}

func (puzzle *Puzzle) RunInstructions() {
	for _, instruction := range puzzle.Instructions {
		puzzle.RunInstruction(instruction)
	}
}

func (puzzle *Puzzle) Load(input string) {
	lines := strings.Split(input, "\n")
	puzzle.input = input
	puzzle.Instructions = lines
}

func (puzzle *Puzzle) Part1() {
	p := NewPuzzleRealData()
	p.RunInstructions()
	fmt.Println(p.Line())
}

func (puzzle *Puzzle) Part2() {
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}
