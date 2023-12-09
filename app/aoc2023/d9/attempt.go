package d9

import (
	"github.com/simonski/aoc/utils"
)

func part1(lines []string) int {
	result := 0
	for _, line := range lines {
		numbers := utils.SplitDataToListOfInts(line, " ")
		result += next(numbers)
	}

	return result
}

func part2(lines []string) int {
	result := 0
	for _, line := range lines {
		numbers := utils.SplitDataToListOfInts(line, " ")
		result += previous(numbers)
	}

	return result
}

func next(numbers []int) int {
	diffs := make([]int, len(numbers)-1)
	zeros := false
	for i := 1; i < len(numbers); i++ {
		diffs[i-1] = numbers[i] - numbers[i-1]
		if diffs[i-1] != 0 {
			zeros = true
		}
	}
	if !zeros {
		return numbers[len(numbers)-1]
	}
	return numbers[len(numbers)-1] + next(diffs)
}

func previous(numbers []int) int {
	diffs := make([]int, len(numbers)-1)
	zeros := false
	for i := 1; i < len(numbers); i++ {
		diffs[i-1] = numbers[i] - numbers[i-1]
		if diffs[i-1] != 0 {
			zeros = true
		}
	}
	if !zeros {
		return numbers[0]
	}
	return numbers[0] - previous(diffs)
}
