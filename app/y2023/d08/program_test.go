package d8

import (
	"fmt"
	"testing"

	"github.com/simonski/aoc/utils"
)

// func Test_1(t *testing.T) {
// 	p := NewPuzzleWithData(TEST_DATA)
// 	fmt.Printf("There are %v lines.\n", len(p.lines))
// }

func Test_Instruction(t *testing.T) {

	i := NewInstruction("LR")

	expect := func(expected string) {
		actual := i.Next()
		if actual != expected {
			t.Fatalf("Expected %v got %v\n", expected, actual)
		}
	}

	expectIndex := func(expected int) {
		actual := i.Index
		if actual != expected {
			t.Fatalf("Expected %v got %v\n", expected, actual)
		}
	}

	expectIndex(-1)
	expect("L")
	expectIndex(0)
	expect("R")
	expectIndex(1)
	expect("L")
	expectIndex(0)
	expect("R")
	expectIndex(1)

}

func Test_InstructionLLR(t *testing.T) {

	i := NewInstruction("LLR")

	expect := func(expected string) {
		actual := i.Next()
		if actual != expected {
			t.Fatalf("Expected %v got %v\n", expected, actual)
		}
	}

	expectIndex := func(expected int) {
		actual := i.Index
		if actual != expected {
			t.Fatalf("Expected %v got %v\n", expected, actual)
		}
	}

	expectIndex(-1)
	expect("L")
	expectIndex(0)
	expect("L")
	expectIndex(1)
	expect("R")
	expectIndex(2)
	expect("L")
	expectIndex(0)
	expect("L")
	expectIndex(1)
	expect("R")
	expectIndex(2)
	expect("L")
	expectIndex(0)

}

func Test_X(t *testing.T) {

	i, d := Load(TEST_DATA_1)
	if i == nil {
		t.Fatalf("Instructions should not be nil.")
	}

	if d == nil {
		t.Fatalf("Directions should not be nil.")
	}

	for name, c := range d.Choices {
		t.Logf("Rule: %v = (%v,%v)\n", name, c.Left, c.Right)
	}
	if len(d.Choices) != 7 {
		t.Fatalf("Expected 7 choices, got %v\n", len(d.Choices))
	}

}

func Test_P1_TestData(t *testing.T) {

	steps, moves := Steps("AAA", "ZZZ", TEST_DATA_1)
	fmt.Printf("%v steps\n", steps)
	for _, move := range moves {
		fmt.Println(move)
	}
	fmt.Printf("%v steps\n\n", steps)
	t.Fatalf("x")

}

func Test_P1_TestData2(t *testing.T) {

	steps, moves := Steps("AAA", "ZZZ", TEST_DATA_2)
	fmt.Printf("%v steps\n", steps)
	for _, move := range moves {
		fmt.Println(move)
	}
	fmt.Printf("%v steps\n\n", steps)
	t.Fatalf("x")

}

func Test_P1_RealData(t *testing.T) {

	steps, moves := Steps("AAA", "ZZZ", REAL_DATA)
	fmt.Printf("%v steps\n", steps)
	for _, move := range moves {
		fmt.Println(move)
	}
	fmt.Printf("%v steps\n\n", steps)
	t.Fatalf("x")

}

func Test_P2_TestData(t *testing.T) {

	p2 := NewP2("11A", TEST_DATA_P2_1, true)
	if p2.Position != "11A" {
		t.Fatalf("Position should be AAA, was %v\n", p2.Position)
	}
	if len(p2.Moves) != 0 {
		t.Fatalf("moves size should be 0, was %v\n", len(p2.Moves))
	}
	if len(p2.Instructions.Data) != 2 {
		t.Fatalf("instructions size should be 2, was %v\n", len(p2.Instructions.Data))
	}
	if len(p2.Directions.Choices) != 8 {
		t.Fatalf("choices size should be 8, was %v\n", len(p2.Directions.Choices))
	}

	if p2.Position != "11A" {
		t.Fatalf("Position should be 11A, was %v\n", p2.Position)
	}

	p2.Step()
	if p2.Position != "11B" {
		t.Fatalf("Position should be BBB, was %v\n", p2.Position)
	}

	p2.Step()
	if p2.Position != "11Z" {
		t.Fatalf("Position should be 11Z, was %v\n", p2.Position)
	}

	p2.Step()
	if p2.Position != "11B" {
		t.Fatalf("Position should be 11B, was %v\n", p2.Position)
	}

	// i, d := Load(TEST_DATA_P2_1)
	// startChoices := d.FindEndsWith("A")

	// count := 0
	// for index, c := range startChoices {

	// }

}

func Test_P2_TestData_X(t *testing.T) {
	p1 := NewP2("11A", TEST_DATA_P2_1, true)
	p2 := NewP2("22A", TEST_DATA_P2_1, true)
	p1.StepToEndingZ()
	p2.StepToEndingZ()
	fmt.Printf("P1 (11A) size %v, ending on %v - %v\n", len(p1.Moves), p1.Position, p1.Moves)
	fmt.Printf("P2 (22A) size %v, ending on %v - %v\n", len(p2.Moves), p2.Position, p2.Moves)

	s1 := uint(len(p1.Moves))
	s2 := uint(len(p2.Moves))
	fmt.Printf("LCM(%v,%v)=%v\n", s1, s2, utils.Lcm(s1, s2))
	t.Fatalf("x")

}

func Test_P2_TestData_YTT(t *testing.T) {

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
			t.Logf("%v -> ?.\n", name)
			counter, _ := p1.StepToEndingZ()
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
			counter2, _ := p1.StepToEndingZ()
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
	t.Fatalf("x")

}
