package d08

import (
	"fmt"
	"log"
	"testing"
)

/*
--- Day 05:  ---

*/

func Test_Grid(t *testing.T) {
	g := NewGrid(TEST_DATA)
	if g.Get(0, 0) != 3 {
		t.Fatalf("Get(0,0) expected 3, got %v\n", g.Get(0, 0))
	}

	if g.Get(0, 2) != 3 {
		t.Fatalf("Get(0,2) expected 3, got %v\n", g.Get(0, 2))
	}

	if g.Get(0, 4) != 3 {
		t.Fatalf("Get(0,4) expected 3, got %v\n", g.Get(0, 2))
	}

	if g.Get(1, 0) != 2 {
		t.Fatalf("Get(1,0) expected 2, got %v\n", g.Get(1, 0))
	}

	if g.Get(1, 1) != 5 {
		t.Fatalf("Get(1,1) expected 5, got %v\n", g.Get(1, 1))
	}
}

func checkVisibility(grid *Grid, row int, col int, height int, top bool, bottom bool, left bool, right bool, t *testing.T) {
	actual_height := grid.Get(row, col)
	if actual_height != height {
		t.Fatalf("(%v,%v) expected height %v, actual height %v\n", row, col, height, actual_height)
	}

	actual_top, actual_bottom, actual_left, actual_right := grid.IsVisible(row, col)
	if actual_top != top {
		t.Fatalf("(%v,%v) expected top %v, actual top %v\n", row, col, top, actual_top)
	}
	if actual_bottom != bottom {
		t.Fatalf("(%v,%v) expected bottom %v, actual bottom %v\n", row, col, bottom, actual_bottom)
	}
	if actual_left != left {
		t.Fatalf("(%v,%v) expected left %v, actual left %v\n", row, col, left, actual_left)
	}
	if actual_right != right {
		t.Fatalf("(%v,%v) expected right %v, actual right %v\n", row, col, right, actual_right)
	}
}

func Test_GridIsVisible(t *testing.T) {

	g := NewGrid(TEST_DATA)
	if !g.IsVisibleFromTop(1, 1) {
		t.Fatalf("IsVisibleFromTop(1,1) expected true, got %v\n", g.IsVisibleFromTop(1, 1))
	}

	if !g.IsVisibleFromLeft(1, 1) {
		t.Fatalf("IsVisibleFromLeft(1,1) expected true, got %v\n", g.IsVisibleFromLeft(1, 1))
	}

	if g.IsVisibleFromRight(1, 1) {
		t.Fatalf("IsVisibleFromRight(1,1) expected false, got %v\n", g.IsVisibleFromRight(1, 1))
	}

	if g.IsVisibleFromBottom(1, 1) {
		t.Fatalf("IsVisibleFromBottom(1,1) expected false, got %v\n", g.IsVisibleFromBottom(1, 1))
	}

}

func Test_GridIsVisible2(t *testing.T) {
	grid := NewGrid(TEST_DATA)

	checkVisibility(grid, 1, 1, 5, true, false, true, false, t)

	checkVisibility(grid, 1, 2, 5, true, false, false, true, t)

	checkVisibility(grid, 1, 3, 1, false, false, false, false, t)

	checkVisibility(grid, 2, 1, 5, false, false, false, true, t)

	checkVisibility(grid, 2, 2, 3, false, false, false, false, t)

	checkVisibility(grid, 2, 3, 3, false, false, false, true, t)

	checkVisibility(grid, 2, 3, 3, false, false, false, true, t)

}

func Test_GridVisibility(t *testing.T) {
	g := NewGrid(TEST_DATA)
	count := 0
	for row := 0; row < g.rows; row++ {
		for col := 0; col < g.cols; col++ {
			top, bottom, left, right := g.IsVisible(row, col)
			if top || bottom || left || right {
				count += 1
			}
		}
	}
	log.Fatalf("Total for test is %v\n", count)

}

func Test_GridVisibilityReal(t *testing.T) {
	g := NewGrid(REAL_DATA)
	count := 0
	for row := 0; row < g.rows; row++ {
		for col := 0; col < g.cols; col++ {
			top, bottom, left, right := g.IsVisible(row, col)
			if top || bottom || left || right {
				count += 1
			}
		}
	}
	log.Fatalf("Total for test is %v\n", count)

}

func Test_GridView(t *testing.T) {
	grid := NewGrid(TEST_DATA)

	value := grid.CountVisibleTop(1, 2)
	if value != 1 {
		t.Fatalf("expect visibility top to be 1, was %v\n", value)
	}

	value = grid.CountVisibleLeft(1, 2)
	if value != 1 {
		t.Fatalf("expect visibility left to be 1, was %v\n", value)
	}

	value = grid.CountVisibleRight(1, 2)
	if value != 2 {
		t.Fatalf("expect visibility right to be 2, was %v\n", value)
	}

	value = grid.CountVisibleBottom(1, 2)
	if value != 2 {
		t.Fatalf("expect visibility bottom to be 2, was %v\n", value)
	}

}

func Test_GridView2(t *testing.T) {
	grid := NewGrid(TEST_DATA)
	bottom := grid.CountVisibleBottom(3, 2)
	left := grid.CountVisibleLeft(3, 2)
	top := grid.CountVisibleTop(3, 2)
	right := grid.CountVisibleRight(3, 2)

	score := left * right * top * bottom
	if score != 8 {
		t.Fatalf("Expected 8, got %v\n", score)
	}
}

func Test_GridView4(t *testing.T) {
	grid := NewGrid(REAL_DATA)
	max_score := 0
	for row := 0; row < grid.rows; row++ {
		for col := 0; col < grid.cols; col++ {
			bottom := grid.CountVisibleBottom(row, col)
			top := grid.CountVisibleTop(row, col)
			left := grid.CountVisibleLeft(row, col)
			right := grid.CountVisibleRight(row, col)
			score := bottom * top * left * right
			if score > max_score {
				max_score = score
			}
		}
	}
	fmt.Printf("max score is %v\n", max_score)
	t.Fatalf("x")
}
