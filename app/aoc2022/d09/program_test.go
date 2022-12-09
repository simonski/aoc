package d09

import (
	"fmt"
	"testing"
)

func Test_1(t *testing.T) {
	// p := NewPuzzleWithData(TEST_DATA_01)
	// fmt.Printf("There are %v lines.\n", len(p.lines))

	b := NewBoard(TEST_DATA, false)
	b.RunInstructions(false)
	for key, value := range b.Tail.data {
		fmt.Printf("%v=%v\n", key, value)
	}
	fmt.Printf("Tail Visits: %v\n", b.CountTailVisits())

	t.Fatalf("mm")
}

func Test_Part1(t *testing.T) {
	// p := NewPuzzleWithData(TEST_DATA_01)
	// fmt.Printf("There are %v lines.\n", len(p.lines))

	b := NewBoard(REAL_DATA, false)
	b.RunInstructions(false)
	for key, value := range b.Tail.data {
		fmt.Printf("%v=%v\n", key, value)
	}
	fmt.Printf("Tail Visits: %v\n", b.CountTailVisits())

	t.Fatalf("mm")
}

func Test_2(t *testing.T) {
	// p := NewPuzzleWithData(TEST_DATA_01)
	// fmt.Printf("There are %v lines.\n", len(p.lines))

	b := NewBoard(TEST_DATA, true)
	instructions := b.Instructions
	for _, i := range instructions {
		b.RunInstruction(i, true)
	}

	t.Fatalf("mm")
}

func Test_1_BoardB2(t *testing.T) {
	// p := NewPuzzleWithData(TEST_DATA_01)
	// fmt.Printf("There are %v lines.\n", len(p.lines))

	size := 2
	b := NewBoardP2(TEST_DATA, false, size, 6, 6)
	b.RunInstructions(false)
	tail := b.Knots[size-1]
	for key, value := range tail.data {
		fmt.Printf("%v=%v\n", key, value)
	}
	fmt.Printf("Tail Visits: %v\n", b.CountTailVisits())

	t.Fatalf("mm")
}

func Test_1_BoardB02_10(t *testing.T) {
	// p := NewPuzzleWithData(TEST_DATA_01)
	// fmt.Printf("There are %v lines.\n", len(p.lines))

	size := 10
	b := NewBoardP2(TEST_DATA, false, size, 6, 6)
	b.RunInstructions(true)
	tail := b.Knots[size-1]
	for key, value := range tail.data {
		fmt.Printf("%v=%v\n", key, value)
	}
	fmt.Printf("Tail Visits: %v\n", b.CountTailVisits())

	t.Fatalf("mm")
}

func Test_1_BoardB02_10_p2(t *testing.T) {
	// p := NewPuzzleWithData(TEST_DATA_01)
	// fmt.Printf("There are %v lines.\n", len(p.lines))

	size := 10
	b := NewBoardP2(TEST_DATA_02, false, size, 26, 26)
	b.RunInstructions(true)
	tail := b.Knots[size-1]
	for key, value := range tail.data {
		fmt.Printf("%v=%v\n", key, value)
	}
	fmt.Printf("Tail Visits: %v\n", b.CountTailVisits())

	t.Fatalf("mm")
}

func Test_1_BoardB02_10_p2_real(t *testing.T) {
	// p := NewPuzzleWithData(TEST_DATA_01)
	// fmt.Printf("There are %v lines.\n", len(p.lines))

	size := 10
	b := NewBoardP2(REAL_DATA, false, size, 26, 26)
	b.RunInstructions(true)
	tail := b.Knots[size-1]
	for key, value := range tail.data {
		fmt.Printf("%v=%v\n", key, value)
	}
	fmt.Printf("Tail Visits: %v\n", b.CountTailVisits())

	t.Fatalf("mm")
}
