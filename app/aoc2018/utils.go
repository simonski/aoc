package aoc2018

import (
	"fmt"
)

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func NewH() []string {
	// an H is a 5x8 grid
	// #...#
	// #...#
	// #...#
	// #####
	// #...#
	// #...#
	// #...#
	// #...#
	var letter []string
	letter = append(letter, "#...#")
	letter = append(letter, "#...#")
	letter = append(letter, "#...#")
	letter = append(letter, "#####")
	letter = append(letter, "#...#")
	letter = append(letter, "#...#")
	letter = append(letter, "#...#")
	letter = append(letter, "#...#")
	return letter
}

func NewI() []string {
	// an I is a 3x8 grid
	// ###
	// .#.
	// .#.
	// .#.
	// .#.
	// .#.
	// .#.
	// ###
	var letter []string
	letter = append(letter, "###")
	letter = append(letter, ".#.")
	letter = append(letter, ".#.")
	letter = append(letter, ".#.")
	letter = append(letter, ".#.")
	letter = append(letter, ".#.")
	letter = append(letter, ".#.")
	letter = append(letter, "###")
	return letter

}

func DebugLetter(letter []string) {
	for index := 0; index < len(letter); index++ {
		line := letter[index]
		fmt.Printf("%v\n", line)
	}
}

func DrawLetter(l []string) {
	for _, line := range l {
		fmt.Printf("%v\n", line)
	}
}

func applyrange(value int, changeby int, lowerbound int, upperbound int) int {
	value += changeby
	if value < lowerbound {
		// we went under the lowerbound, wrap around to the upperbound
		diff := Abs(value - lowerbound)
		value = upperbound - diff
	} else if value > upperbound {
		// we went over the upperbound, wrap aroudn to the lowerbound
		diff := Abs(value - upperbound)
		value = lowerbound + diff
	}
	return value
}
