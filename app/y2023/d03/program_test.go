package d3

import (
	"fmt"
	"testing"
)

func Test_FindNumbers(t *testing.T) {
	grid := NewGrid(TEST_DATA)
	numbers := grid.FindNumbers()
	for _, n := range numbers {
		fmt.Println(n.Debug())
	}
	if len(numbers) != 10 {
		t.Fatalf("Expected 10 numbers, got %v\n", len(numbers))
	}
}

func Test_FindNumbersAdjacent(t *testing.T) {
	grid := NewGrid(TEST_DATA)
	numbers := grid.FindNumbersAdjacent()
	for _, n := range numbers {
		fmt.Println(n.Debug())
	}
	if len(numbers) != 8 {
		t.Fatalf("Expected 8 adjacent numbers, got %v\n", len(numbers))
	}

	total := 0
	for _, n := range numbers {
		total += n.Value
	}
	fmt.Printf("Test total is %v\n", total)
}

func Test_FindNumbersAdjacent_Part1(t *testing.T) {
	grid := NewGrid(REAL_DATA)
	numbers := grid.FindNumbersAdjacent()
	for _, n := range numbers {
		fmt.Println(n.Debug())
	}

	total := 0
	for _, n := range numbers {
		total += n.Value
	}
	fmt.Printf("Test total is %v\n", total)
}

func Test_FindGearNumbers_Part2(t *testing.T) {
	grid := NewGrid(TEST_DATA)
	numbers := grid.FindGearNumbers()
	if len(numbers) != 2 {
		t.Fatalf("Expected 2 gearnumbers, got %v\n", len(numbers))
	}
}
