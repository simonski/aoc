package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	cli "github.com/simonski/cli"
)

type AppLogic interface {
	Run(cli *cli.CLI)
	Render(cli *cli.CLI)
	Help(cli *cli.CLI)
	GetMethod(methodName string) (reflect.Value, reflect.Value, bool)
	GetName() string
	Api(day int) string
	Summary(year int, day int) *Summary
}

/*
Converts a decimal string to an integer value
e.g. "11" -> utisl.BinaryStringToInt("11") = 3
*/
func BinaryStringToInt(v string) int {
	if len(v) > 64 {
		fmt.Print("\n\n\n\n\n\n\n\n\n\n")
		fmt.Printf("BinaryStringToInt converting a %v-bit number !!!\n", len(v))
		fmt.Println(v)
		fmt.Print("\n\n\n\n\n\n\n\n\n\n")
	}
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

func BinaryStringToInt64(v string) int64 {
	if len(v) > 64 {
		fmt.Print("\n\n\n\n\n\n\n\n\n\n")
		fmt.Printf("BinaryStringToInt converting a %v-bit number !!!\n", len(v))
		fmt.Println(v)
		fmt.Print("\n\n\n\n\n\n\n\n\n\n")
	}
	result := int64(0)
	pow := 1
	for index := len(v) - 1; index >= 0; index-- {
		value, _ := strconv.Atoi(v[index : index+1])
		value *= pow
		result += int64(value)
		pow += pow
	}
	return result
}

func BinaryStringToUInt64(v string) uint64 {
	if len(v) > 64 {
		fmt.Print("\n\n\n\n\n\n\n\n\n\n")
		fmt.Printf("BinaryStringToInt converting a %v-bit number !!!\n", len(v))
		fmt.Println(v)
		fmt.Print("\n\n\n\n\n\n\n\n\n\n")
	}
	result := uint64(0)
	pow := uint64(1)
	for index := len(v) - 1; index >= 0; index-- {
		value, _ := strconv.Atoi(v[index : index+1])
		ival := uint64(value)
		ival *= pow
		result += ival
		pow += pow
	}
	return result
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

// func NewH() []string {
// 	// an H is a 5x8 grid
// 	// #...#
// 	// #...#
// 	// #...#
// 	// #####
// 	// #...#
// 	// #...#
// 	// #...#
// 	// #...#
// 	var letter []string
// 	letter = append(letter, "#...#")
// 	letter = append(letter, "#...#")
// 	letter = append(letter, "#...#")
// 	letter = append(letter, "#####")
// 	letter = append(letter, "#...#")
// 	letter = append(letter, "#...#")
// 	letter = append(letter, "#...#")
// 	letter = append(letter, "#...#")
// 	return letter
// }

// func NewI() []string {
// 	// an I is a 3x8 grid
// 	// ###
// 	// .#.
// 	// .#.
// 	// .#.
// 	// .#.
// 	// .#.
// 	// .#.
// 	// ###
// 	var letter []string
// 	letter = append(letter, "###")
// 	letter = append(letter, ".#.")
// 	letter = append(letter, ".#.")
// 	letter = append(letter, ".#.")
// 	letter = append(letter, ".#.")
// 	letter = append(letter, ".#.")
// 	letter = append(letter, ".#.")
// 	letter = append(letter, "###")
// 	return letter

// }

// func DebugLetter(letter []string) {
// 	for index := 0; index < len(letter); index++ {
// 		line := letter[index]
// 		fmt.Printf("%v\n", line)
// 	}
// }

// func DrawLetter(l []string) {
// 	for _, line := range l {
// 		fmt.Printf("%v\n", line)
// 	}
// }

// func Applyrange(value int, changeby int, lowerbound int, upperbound int) int {
// 	value += changeby
// 	if value < lowerbound {
// 		// we went under the lowerbound, wrap around to the upperbound
// 		diff := goutils.Abs(value - lowerbound)
// 		value = upperbound - diff
// 	} else if value > upperbound {
// 		// we went over the upperbound, wrap aroudn to the lowerbound
// 		diff := goutils.Abs(value - upperbound)
// 		value = lowerbound + diff
// 	}
// 	return value
// }
