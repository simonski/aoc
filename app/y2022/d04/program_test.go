package d04

import (
	"fmt"
	"testing"
)

func Test_AOC2022_04_Part1_ShouldOverlap(t *testing.T) {
	input := "2-4,6-8"
	pair := NewPair(input)
	t.Logf(pair.task1.Debug(10, ""))
	t.Logf(pair.task2.Debug(10, ""))
	if pair.Overlap() {
		t.Fatalf("%v should not overlap.n", input)
	}
}

func Test_AOC2022_04_Part1_SholdNotOverlap(t *testing.T) {
	input := "2-8,3-7"
	pair := NewPair(input)
	t.Logf(pair.task1.Debug(10, ""))
	t.Logf(pair.task2.Debug(10, ""))
	if !pair.Overlap() {
		t.Fatalf("%v should overlap.n", input)
	}
}

func Test_AOC2022_04_Part1_ShouldNotIntersect_1(t *testing.T) {
	input := "2-4,6-8"
	pair := NewPair(input)
	t.Logf(pair.task1.Debug(10, ""))
	t.Logf(pair.task2.Debug(10, ""))
	if pair.Intersect() {
		t.Fatalf("%v should not intesect.n", input)
	}
}

func Test_AOC2022_04_Part1_ShouldNotIntersect_2(t *testing.T) {
	input := "2-3,4-5"
	pair := NewPair(input)
	t.Logf(pair.task1.Debug(10, ""))
	t.Logf(pair.task2.Debug(10, ""))
	if pair.Intersect() {
		t.Fatalf("%v should not intesect.n", input)
	}
}

func Test_AOC2022_04_Part1_ShouldIntersect_3(t *testing.T) {
	input := "5-7,7-9"
	pair := NewPair(input)
	t.Logf(pair.task1.Debug(10, ""))
	t.Logf(pair.task2.Debug(10, ""))
	if !pair.Intersect() {
		t.Fatalf("%v should intesect.n", input)
	}
}

func Test_AOC2022_04_Part1_ShouldIntersect_4(t *testing.T) {
	input := "2-8,3-7"
	pair := NewPair(input)
	t.Logf(pair.task1.Debug(10, ""))
	t.Logf(pair.task2.Debug(10, ""))
	if !pair.Intersect() {
		t.Fatalf("%v should intesect.n", input)
	}
}

func Test_AOC2022_04_Part1_ShouldIntersect_5(t *testing.T) {
	input := "6-6,4-6"
	pair := NewPair(input)
	t.Logf(pair.task1.Debug(10, ""))
	t.Logf(pair.task2.Debug(10, ""))
	if !pair.Intersect() {
		t.Fatalf("%v should intesect.n", input)
	}
}

func Test_AOC2022_04_Part1_Verify(t *testing.T) {
	input := TEST_DATA
	day4 := NewPuzzle(input)
	overlaps := 0
	for _, pair := range day4.Pairs {
		if pair.Overlap() {
			overlaps += 1
		}
	}
	if overlaps != 2 {
		t.Fatalf("Number of overlaps should be 2, was %v\n", overlaps)
	}
}

func Test_AOC2022_04_Part2_Verify(t *testing.T) {
	input := TEST_DATA
	day4 := NewPuzzle(input)
	intersections := 0
	for _, pair := range day4.Pairs {
		if pair.Intersect() {
			intersections += 1
		}
	}
	if intersections != 4 {
		t.Fatalf("Number of overlaps should be 2, was %v\n", intersections)
	}
}

func Test_AOC2022_04_Part1(t *testing.T) {
	input := REAL_DATA
	day4 := NewPuzzle(input)
	overlaps := 0
	for _, pair := range day4.Pairs {
		if pair.Overlap() {
			overlaps += 1
		}
	}
	fmt.Printf("2022-04 Part 1: Overlaps: %v\n", overlaps)
}

func Test_AOC2022_04_Part2(t *testing.T) {
	input := REAL_DATA
	day4 := NewPuzzle(input)
	overlaps := 0
	for _, pair := range day4.Pairs {
		if pair.Intersect() {
			overlaps += 1
		}
	}
	fmt.Printf("2022-04 Part 2: Overlaps: %v\n", overlaps)
}
