package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

const TEST_09_DATA = `35
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
576`

func Test_AOC2020_09_NewInstruction(t *testing.T) {

	// load test data
	ints := make([]int, 0)
	splits := strings.Split(TEST_09_DATA, "\n")
	for _, line := range splits {
		value, _ := strconv.Atoi(line)
		ints = append(ints, value)
	}

	preamble := 5

	s := Sequence{values: ints, preamble: preamble}
	for index := preamble; index < s.Size(); index++ {
		// at position, check the previous N entries as a subsequence and find a combo.
		currentValue := s.Get(index)
		subSequence := s.Subsequence(index)
		contains, index1, index2 := subSequence.ContainsCombination(currentValue)
		if contains {
			value1 := subSequence.Get(index1)
			value2 := subSequence.Get(index2)
			fmt.Printf("index %v, value %v, sequence %v contains values %v and %v\n", index, currentValue, subSequence, value1, value2)
		} else {
			fmt.Printf("index %v, value %v, sequence %v does not contain a combination.\n", index, currentValue, subSequence)
		}
	}

}
