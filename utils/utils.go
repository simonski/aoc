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
	for i := range splits {
		iline, _ := strconv.Atoi(splits[i])
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

// # Recursive function to return gcd of a and b
func Gcd(a uint, b uint) uint {
	if a == 0 {
		return b
	}
	return Gcd(b%a, a)
}

// # Function to return LCM of two numbers
func Lcm(a uint, b uint) uint {
	return (a / Gcd(a, b)) * b
}

func Lcm_x(data []uint) uint {
	result := Lcm(data[0], data[1])
	for index := 1; index < len(data); index++ {
		result = Lcm(result, data[index])
	}
	return result
}

// func Compute_lcms(data []uint) uint {
// 	lcm := Compute_lcm(data[0], data[1])
// 	for index := 1; index < len(data); index++ {
// 		lcm = Compute_lcm(lcm, data[index])
// 	}
// 	return lcm
// }

// func Compute_lcm(x uint, y uint) uint {

// 	//    # choose the greater number
// 	var greater uint
// 	if x > y {
// 		greater = x
// 	} else {
// 		greater = y
// 	}

// 	var lcm uint
// 	for {
// 		if (greater%x == 0) && (greater%y == 0) {
// 			lcm = greater
// 			break
// 		}
// 		greater += 1
// 	}
// 	return lcm
// }

func MaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// a semi-functional collapsing function which finds overlaps and merges indexes
func CollapseRanges(ranges [][]int) [][]int {

	inside := func(r1 []int, r2 []int) bool {
		if r1[0] > r2[0] && r1[1] < r2[0] {
			return true
		}
		return false
	}

	extends_left := func(r1 []int, r2 []int) bool {
		if r1[0] < r2[0] && r1[1] > r2[0] && r1[1] < r2[1] {
			return true
		}
		return false
	}

	extends_right := func(r1 []int, r2 []int) bool {
		if r1[1] > r2[1] && r1[0] > r2[0] && r1[0] < r2[1] {
			return true
		}
		return false
	}

	filter := func(r1 []int, r2 []int) (bool, []int) {
		if inside(r1, r2) {
			// r1 is inside r2
			// drop r1
			return true, r2
		} else if extends_left(r1, r2) {
			// r1 extends r2 left
			result := []int{r1[0], r2[1]}
			return true, result

		} else if extends_right(r1, r2) {
			// r1 extends r2 right
			result := []int{r2[0], r1[1]}
			return true, result

		} else {
			return false, nil
		}
	}

	filter_down := func(ranges [][]int) (bool, [][]int) {

		for index := 0; index < len(ranges); index++ {
			r1 := ranges[index]
			for index2 := index + 1; index2 < len(ranges); index2++ {
				if index == index2 {
					// dont compare thyself
					continue
				}
				r2 := ranges[index2]

				fmt.Printf("checking (%v) (%v)\n", r1, r2)

				modified, r3 := filter(r1, r2)
				if modified {
					fmt.Printf("modified (%v) (%v)\n", r1, r2)
					// remove r1 and r2, replace with r3
					new_ranges := make([][]int, 0)
					new_ranges = append(new_ranges, r3)
					for i := 0; i < len(ranges); i++ {
						if i != index && i != index2 {
							new_ranges = append(new_ranges, ranges[i])
						}
					}
					return true, new_ranges
				} else {
					fmt.Printf("not modified (%v) (%v)\n", r1, r2)

				}

			}
		}
		return false, ranges

	}

	filter_all := func(ranges [][]int) [][]int {
		results := ranges
		modified := true
		for {
			modified, results = filter_down(results)
			if !modified {
				break
			}
		}
		return results
	}

	merged := filter_all(ranges)
	return merged
}

func ToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
