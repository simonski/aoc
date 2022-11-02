package aoc2020

/*
Day 09 - Encoding Error
https://adventofcode.com/2020/day/9

--- Day 9: Encoding Error ---
With your neighbor happily enjoying their video game, you turn your attention to an open data port on the little screen in the seat in front of you.

Though the port is non-standard, you manage to connect it to your computer through the clever use of several paperclips. Upon connection, the port outputs a series of numbers (your puzzle input).

The data appears to be encrypted with the eXchange-Masking Addition System (XMAS) which, conveniently for you, is an old cypher with an important weakness.

XMAS starts by transmitting a preamble of 25 numbers. After that, each number you receive should be the sum of any two of the 25 immediately previous numbers. The two numbers will have different values, and there might be more than one such pair.

For example, suppose your preamble consists of the numbers 1 through 25 in a random order. To be valid, the next number must be the sum of two of those numbers:

26 would be a valid next number, as it could be 1 plus 25 (or many other pairs, like 2 and 24).
49 would be a valid next number, as it is the sum of 24 and 25.
100 would not be valid; no two of the previous 25 numbers sum to 100.
50 would also not be valid; although 25 appears in the previous 25 numbers, the two numbers in the pair must be different.
Suppose the 26th number is 45, and the first number (no longer an option, as it is more than 25 numbers ago) was 20. Now, for the next number to be valid, there needs to be some pair of numbers among 1-19, 21-25, or 45 that add up to it:

26 would still be a valid next number, as 1 and 25 are still within the previous 25 numbers.
65 would not be valid, as no two of the available numbers sum to it.
64 and 66 would both be valid, as they are the result of 19+45 and 21+45 respectively.
Here is a larger example which only considers the previous 5 numbers (and has a preamble of length 5):

35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576
In this example, after the 5-number preamble, almost every number is the sum of two of the previous 5 numbers; the only number that does not follow this rule is 127.

The first step of attacking the weakness in the XMAS data is to find the first number in the list (after the preamble) which is not the sum of two of the 25 numbers before it. What is the first number that does not have this property?
*/
import (
	"fmt"
	"sort"

	"github.com/simonski/aoc/utils"
	goutils "github.com/simonski/goutils"
)

func (app *Application) Y2020D09_Summary() *utils.Summary {
	s := utils.NewSummary(2020, 9)
	s.Name = "Encoding Error"
	s.ProgressP1 = utils.Completed
	s.ProgressP2 = utils.Completed
	return s
}

// AOC_2020_09 is the entrypoint
func (app *Application) Y2020D09P1() {
	AOC_2020_09_part1_attempt1(app)
}
func (app *Application) Y2020D09P2() {
	AOC_2020_09_part2_attempt1(app)
}

func AOC_2020_09_part2_attempt1(app *Application) {
	cli := app.CLI
	searchFor := 373803594
	// find the first contiguous block of numbers that sums to our number
	filename := cli.GetFileExistsOrDie("-input")
	ints := goutils.Load_file_to_ints(filename)

	for index1, _ := range ints {
		for index2, _ := range ints {
			if index2 <= index1 {
				continue
			}
			slice := ints[index1:index2]
			total := 0
			minvalue := 99999999
			maxvalue := 0
			for _, value3 := range slice {
				total += value3
				minvalue = goutils.Min(minvalue, value3)
				maxvalue = goutils.Max(maxvalue, value3)
			}
			if total == searchFor {
				value1 := ints[index1]
				value2 := ints[index2]
				fmt.Printf("index1: %v [value=%v], index2 %v [value=%v], total=%v, slice %v\n", index1, value1, index2, value2, total, slice)
				fmt.Printf("Weakness is [%v + %v] = %v\n", minvalue, maxvalue, minvalue+maxvalue)
				return
			} else {
				fmt.Printf("index1: %v, index2 %v, total=%v\n", index1, index2, total)
			}

		}

	}

}

func AOC_2020_09_part1_attempt1(app *Application) {
	cli := app.CLI
	filename := cli.GetFileExistsOrDie("-input")

	ints := goutils.Load_file_to_ints(filename)
	preamble := 25

	s := Sequence{values: ints, preamble: preamble}
	for index := preamble; index < s.Size(); index++ {
		// at position, check the previous N entries as a subsequence and find a combo.
		currentValue := s.Get(index)
		subSequence := s.Subsequence(index)
		contains, _, _ := subSequence.ContainsCombination(currentValue)
		if contains {
			// value1 := subSequence.Get(index1)
			// value2 := subSequence.Get(index2)
			//fmt.Printf("index %v, value %v, sequence %v contains values %v and %v\n", index, currentValue, subSequence, value1, value2)
		} else {
			fmt.Printf("index %v, value %v, sequence %v does not contain a combination.\n", index, currentValue, subSequence)
		}
	}
}

type Sequence struct {
	values   []int
	preamble int
}

func NewSequenceFromFilename(filename string, preamble int) *Sequence {
	inty_list := goutils.Load_file_to_ints(filename)
	s := Sequence{values: inty_list, preamble: preamble}
	return &s
}

func (s *Sequence) IsValid(index int) bool {
	sub := s.Subsequence(index)
	value := s.values[index]
	result, _, _ := sub.ContainsCombination(value)
	return result
}

func (s *Sequence) Size() int {
	return len(s.values)
}

func (s *Sequence) Get(index int) int {
	return s.values[index]
}

// Searches this sequence for the combination, indicating true/false, index1, index2
func (s *Sequence) ContainsCombination(totalToFind int) (bool, int, int) {
	for index1, value := range s.values {
		// for the current value, given the list is sorted, attempt to find the difference
		differenceRequired := totalToFind - value
		contains, index2 := s.ContainsValue(differenceRequired)
		if contains {
			return true, index1, index2
		}
	}
	return false, -1, -1
}

func (s *Sequence) ContainsValue(valueToFind int) (bool, int) {
	values := s.values
	i := sort.Search(len(values), func(i int) bool { return values[i] >= valueToFind })
	if i < len(values) && values[i] == valueToFind {
		return true, i
	}
	return false, -1

}

func (s *Sequence) IndexOf(valueToFind int) int {
	values := s.values
	i := sort.Search(len(values), func(i int) bool { return values[i] >= valueToFind })
	if i < len(values) && values[i] == valueToFind {
		return i
	}
	return -1
}

func NewSequenceFromInts(values []int, preamble int) *Sequence {
	s := Sequence{values: values, preamble: preamble}
	return &s
}

// Returns the previous N entries from index X. ordered ascending
func (s *Sequence) Subsequence(position int) *Sequence {
	left := position - s.preamble
	// right := position
	slice := s.values[left : left+s.preamble]
	values := make([]int, s.preamble)
	copy(values, slice)
	// fmt.Printf("Subsequence(%v), [%v:%v] = %v\n", s.values, left, right, values)
	sort.Ints(values)
	return NewSequenceFromInts(values, s.preamble)
}
