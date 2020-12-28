package main

import (
	"fmt"
	"testing"
)

func Test_AOC2020_20_Test1(t *testing.T) {
	image := NewImageFromString(DAY_20_TEST_DATA)
	if image.Size() != 9 {
		t.Errorf("Day 20 Part 1: Test data should have 6 tiles, was %v.\n", image.Size())
	}

	tile2311 := image.GetTile("2311")
	fmt.Printf("\nRotate(0)\n")
	fmt.Printf("%v\n", tile2311.Debug())

	tile2311_cw90 := image.GetTile("2311").Rotate()
	fmt.Printf("\nRotate(90)\n")
	fmt.Printf("%v\n", tile2311_cw90.Debug())

	tile2311_cw180 := image.GetTile("2311").Rotate()
	fmt.Printf("\nRotate(180)\n")
	fmt.Printf("%v\n", tile2311_cw180.Debug())

	tile2311_cw270 := image.GetTile("2311").Rotate()
	fmt.Printf("\nRotate(270)\n")
	fmt.Printf("%v\n", tile2311_cw270.Debug())

	tile2311_cw360 := image.GetTile("2311").Rotate()
	fmt.Printf("\nRotate(360)\n")
	fmt.Printf("%v\n", tile2311_cw360.Debug())

}

func Test_AOC2020_20_TestFlipVertical(t *testing.T) {
	image := NewImageFromString(DAY_20_TEST_DATA)
	tile2311 := image.GetTile("2311")

	expected_original := `..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###`

	expected_flipped := `..###..###
###...#.#.
..#....#..
.#.#.#..##
##...#.###
##.##.###.
####.#...#
#...##..#.
##..#.....
..##.#..#.`

	actual_original := tile2311.Debug()
	if actual_original != expected_original {
		t.Errorf("FlipVertical: originals do not match\n.")
	}

	tile2311.FlipVertical()
	actual_flipped := tile2311.Debug()
	if actual_flipped != expected_flipped {
		t.Errorf("FlipVertical: flipped do not match\n.")
		t.Errorf("Pre-flip\n\n%v\n\n", actual_original)
		t.Errorf("Expected\n\n%v\n\n", expected_flipped)
		t.Errorf("Actual\n\n%v\n\n", actual_flipped)
	}

}

func Test_AOC2020_20_TestFlipHorizontal(t *testing.T) {
	image := NewImageFromString(DAY_20_TEST_DATA)
	tile2311 := image.GetTile("2311")

	expected_original := `..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###`

	expected_flipped := `.#..#.##..
.....#..##
.#..##...#
#...#.####
.###.##.##
###.#...##
##..#.#.#.
..#....#..
.#.#...###
###..###..`

	actual_original := tile2311.Debug()
	if actual_original != expected_original {
		t.Errorf("FlipHorizontal: originals do not match\n.")
	}

	tile2311.FlipHorizontal()
	actual_flipped := tile2311.Debug()
	if actual_flipped != expected_flipped {
		t.Errorf("FlipHorizontal: flipped do not match\n.")
		t.Errorf("Pre-flip\n\n%v\n\n", actual_original)
		t.Errorf("Expected\n\n%v\n\n", expected_flipped)
		t.Errorf("Actual\n\n%v\n\n", actual_flipped)
	}

}
