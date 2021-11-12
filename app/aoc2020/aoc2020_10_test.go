package aoc2020

import (
	"fmt"
	"strings"
	"testing"

	goutils "github.com/simonski/goutils"
)

const TEST_10_DATA_01 = `16
10
15
5
1
11
7
19
6
12
4`

const TEST_10_DATA_02 = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`

func Test_AOC2020_10_Jolt1(t *testing.T) {

	// load test data
	ints := goutils.Convert_strings_to_ints(strings.Split(TEST_10_DATA_01, "\n"))
	results := AOC_2020_10_part1_attempt1(ints)

	actual1, _ := results[1]
	actual3, _ := results[3]

	expected1 := 7
	expected3 := 5

	if actual1 != expected1 {
		t.Errorf("Count of 1s should be %v was %v.\n", expected1, actual1)
	}
	if actual3 != expected3 {
		t.Errorf("Count of 5s should be %v was %v.\n", expected3, actual3)
	}
	fmt.Printf("%v\n", results)
}

func Test_AOC2020_10_Jolt2(t *testing.T) {

	// load test data
	ints := goutils.Convert_strings_to_ints(strings.Split(TEST_10_DATA_02, "\n"))
	results := AOC_2020_10_part1_attempt1(ints)

	actual1, _ := results[1]
	actual3, _ := results[3]

	expected1 := 22
	expected3 := 10

	if actual1 != expected1 {
		t.Errorf("Count of 1s should be %v was %v.\n", expected1, actual1)
	}
	if actual3 != expected3 {
		t.Errorf("Count of 5s should be %v was %v.\n", expected3, actual3)
	}
	fmt.Printf("%v\n", results)
}
