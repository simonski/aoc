package d05

import (
	"testing"
)

func requireStack(puzzle *Puzzle, index int, expected string, t *testing.T) {
	stack := puzzle.Stacks[index]
	if stack.Flatten() != expected {
		t.Fatalf("Stack[%v] should be %v, was %v.\n", index, expected, stack.Flatten())
	}
}

func Test_Part1(t *testing.T) {
	p := NewPuzzleRealData()
	p.RunInstructions()
	t.Fatalf("Line was %v\n", p.Line())
}

func Test_Part2(t *testing.T) {
	p := NewPuzzleRealData()
	p.RunInstructions9001()
	t.Fatalf("Line was %v\n", p.Line())
}

func Test_1_9000(t *testing.T) {
	p := NewPuzzleTestData()
	p.RunInstructions()

	line := p.Line()
	if line != "CMZ" {
		t.Fatalf("Line should be CMZ, was %v\n", line)
	}
}

func Test_1_9001(t *testing.T) {
	p := NewPuzzleTestData()
	p.RunInstructions9001()

	line := p.Line()
	if line != "MCD" {
		t.Fatalf("Line should be MCD, was %v\n", line)
	}
}

func Test_1(t *testing.T) {
	p := NewPuzzleTestData()
	requireStack(p, 0, "ZN", t)
	requireStack(p, 1, "MCD", t)
	requireStack(p, 2, "P", t)

	i1 := p.Instructions[0]
	p.RunInstruction(i1)

	requireStack(p, 0, "ZND", t)
	requireStack(p, 1, "MC", t)
	requireStack(p, 2, "P", t)

	i2 := p.Instructions[1]
	p.RunInstruction(i2)

	requireStack(p, 0, "", t)
	requireStack(p, 1, "MC", t)
	requireStack(p, 2, "PDNZ", t)

	i3 := p.Instructions[2]
	p.RunInstruction(i3)

	requireStack(p, 0, "CM", t)
	requireStack(p, 1, "", t)
	requireStack(p, 2, "PDNZ", t)

	i4 := p.Instructions[3]
	p.RunInstruction(i4)

	requireStack(p, 0, "C", t)
	requireStack(p, 1, "M", t)
	requireStack(p, 2, "PDNZ", t)

	line := p.Line()
	if line != "CMZ" {
		t.Fatalf("Line should be CMZ, was %v\n", line)
	}

}
