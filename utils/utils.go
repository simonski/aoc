package utils

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Converts a decimal string to an integer value
e.g. "11" -> utisl.BinaryStringToInt("11") = 3
*/
func BinaryStringToInt(v string) int {
	result := 0
	pow := 1
	for index := len(v) - 1; index >= 0; index-- {
		value, _ := strconv.Atoi(v[index : index+1])
		value *= pow
		result += value
		pow += pow
	}
	return result
}

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

func Factorial(a uint64) uint64 {
	if a > 1 {
		a = a * Factorial(a-1)
		return a
	} else {
		return a
	}
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

// reads some test data to a slice of ints
func SplitDataToListOfInts(data string, delim string) []int {
	results := make([]int, 0)
	splits := strings.Split(data, delim)
	for _, line := range splits {
		iline, _ := strconv.Atoi(line)
		results = append(results, iline)
	}
	return results
}
