package d12

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/simonski/aoc/utils"
)

const FAILURE = -1
const TOTAL_SUCCESS = 0
const PARTIAL_SUCCESS = 1

func eval(current string, rules []int) int {
	// so we need to check, does it match each rule and then have more time to spare?
	actuals := buildBlocks(current, rules)
	_, _, invalid := evaluateBlocks(0, "", "", actuals, rules)
	fmt.Printf("actuals=%v, invalid=%v\n", actuals, invalid)

	// now we have have
	// a line #....#....#
	// the rules 3,3,1
	// the blocks 1,1,1
	// this means we can quickly decide if this is going to work, as the blocks 1,1,1 won't work on the first 3,3,1

	pass_count := 0
	if len(actuals) > len(rules) {
		fmt.Printf("fail, more actuals (%v) that rules (%v).\n", len(actuals), len(rules))
		return FAILURE
	} else {
		for index := range actuals {
			if actuals[index] > rules[index] {
				fmt.Printf("fail, actual[%v]=%v ()>)  != rules[%v]=%v\n", index, actuals[index], index, rules[index])
				return FAILURE
			} else if actuals[index] < rules[index] && index < len(actuals) {
				fmt.Printf("fail, actual[%v]=%v (<) != rules[%v]=%v\n", index, actuals[index], index, rules[index])
				return FAILURE
			} else if actuals[index] == rules[index] {
				fmt.Printf("pass1, actual[%v]=%v == rules[%v]=%v\n", index, actuals[index], index, rules[index])
				pass_count += 1
			} else {
				fmt.Printf("dropped %v\n", index)
			}
		}
	}

	if pass_count == len(rules) {
		fmt.Print("passed rules = len rules, TOTAL_SUCCESS\n")
		return TOTAL_SUCCESS
	} else {
		fmt.Print("passed rules != len rules, PARTIAL_SUCCESS\n")
		return PARTIAL_SUCCESS
	}

}

func Test_3(t *testing.T) {
	example1 := "######...#.#.#.#.#..##....#....#.#####.....#..#..#####.....#...## "
	example2 := "######...#.#.#."
	rules := utils.SplitDataToListOfInts("6 1 1 6 1 1 6 1 1 6 1 1 6 1 1", " ")

	result1 := eval(example1, rules)
	result2 := eval(example2, rules)

	if result1 != FAILURE {
		t.Fatalf("Result1 got %v expected %v\n", result1, FAILURE)
	}

	if result2 != FAILURE {
		t.Fatalf("Result2 got %v expected %v\n", result2, FAILURE)
	}

	fmt.Printf("Result1: %v\n", result1)
	fmt.Printf("Result2: %v\n", result2)
	t.Fail()

}

func Test_X(t *testing.T) {
	line := ".######......#.#.######......#.#.######......#.#.######......#.#.######..#.#.. "
	rules := []int{6, 1, 1, 6, 1, 1, 6, 1, 1, 6, 1, 1, 6, 1, 1}
	fmt.Printf("%v\n", eval(line, rules))
	t.Fail()
}

// [6 1 1 6 1 1 6 1 1 6 1 1 6 1 1] false

func Test_arr(t *testing.T) {
	a := []int{1, 2, 5}
	b := []int{1, 2, 3}
	if !reflect.DeepEqual(a, b) {
		t.Fail()
	}
}
