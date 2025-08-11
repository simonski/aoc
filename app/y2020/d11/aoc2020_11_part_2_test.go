package aoc2020

import (
	"fmt"
	"strings"
	"testing"
)

const TEST_11_2_DATA_01 = `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

const TEST_11_2_DATA_02 = `#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`

const TEST_11_2_DATA_03 = `#.LL.LL.L#
#LLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLLL.L
#.LLLLL.L#`

const TEST_11_2_DATA_04 = `#.L#.##.L#
#L#####.LL
L.#.#..#..
##L#.##.##
#.##.#L.##
#.#####.#L
..#.#.....
LLL####LL#
#.L#####.L
#.L####.L#`

const TEST_11_2_DATA_05 = `#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##LL.LL.L#
L.LL.LL.L#
#.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLL#.L
#.L#LL#.L#`

const TEST_11_2_DATA_06 = `#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.#L.L#
#.L####.LL
..#.#.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#`

const TEST_11_2_DATA_07 = `#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.LL.L#
#.LLLL#.LL
..#.L.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#`

const TEST_ADJACENT_0 = `LLLLLLLLLL
LLLLLLLLLL
LLLLLLLLLL
LLLLL#LLLL
LLLLLLLLLL
LLLLLLLLLL
LLLLLLLLLL`

const TEST_ADJACENT_1 = `LLLLL#LLLL
LLLLLLLLLL
LLLLLLLLLL
LLLLL#LLLL
LLLLLLLLLL
LLLLLLLLLL
LLLLLLLLLL`

func Test_AOC2020_11_day2_Ticks(t *testing.T) {
	tolerance := 5
	searchFar := true
	sp := NewSeatingPlanFromStrings(strings.Split(TEST_11_2_DATA_01, "\n"), tolerance, searchFar)
	actualTick0 := sp.Debug()
	if actualTick0 != TEST_11_2_DATA_01 {
		t.Errorf("Test_AOC2020_11_day2_Ticks: Loaded data does not conform\nExpected\n%v\n\nActual\n%v\n\n", TEST_11_2_DATA_01, actualTick0)
	} else {
		fmt.Printf("Test_AOC2020_11_day2_Ticks: Loaded data conforms to expected format.\n")
	}

	changes := sp.Tick()
	fmt.Printf("Tick(1), changes=%v\n", changes)
	actualTick1 := sp.Debug()
	if actualTick1 != TEST_11_2_DATA_02 {
		t.Errorf("Test_AOC2020_11_day2_Ticks: Tick(1) not conform\nExpected\n%v\n\nActual\n%v\n\n", TEST_11_2_DATA_02, actualTick1)
	} else {
		fmt.Printf("OK Test_AOC2020_11_day2_Ticks: Tick(1) conforms.\n")
	}

	changes = sp.Tick()
	fmt.Printf("Tick(2), changes=%v\n", changes)
	actualTick2 := sp.Debug()
	if actualTick2 != TEST_11_2_DATA_03 {
		t.Errorf("Test_AOC2020_11_day2_Ticks: Tick(2) not conform\nExpected\n%v\n\nActual\n%v\n\n", TEST_11_2_DATA_03, actualTick2)
	} else {
		fmt.Printf("OK Test_AOC2020_11_day2_Ticks: Tick(2) conforms.\n")
	}

	changes = sp.Tick()
	fmt.Printf("Tick(3), changes=%v\n", changes)
	actualTick3 := sp.Debug()
	if actualTick3 != TEST_11_2_DATA_04 {
		t.Errorf("Test_AOC2020_11_day2_Ticks: Tick(3) changed %v, does not conform\n\nWas previously\n\n%v\n\nExpected\n%v\n\nActual\n%v\n\n", changes, actualTick2, TEST_11_2_DATA_04, actualTick3)
	} else {
		fmt.Printf("OK Test_AOC2020_11_day2_Ticks: Tick(3).\n")
	}

	changes = sp.Tick()
	actualTick4 := sp.Debug()
	if actualTick4 != TEST_11_2_DATA_05 {
		t.Errorf("Test_AOC2020_11_day2_Ticks: Tick(4)  not conform\nExpected\n%v\n\nActual\n%v\n\n", TEST_11_2_DATA_05, actualTick4)
	} else {
		fmt.Printf("OK Test_AOC2020_11_day2_Ticks: Tick(4).\n")
	}

	changes = sp.Tick()
	actualTick5 := sp.Debug()
	if actualTick5 != TEST_11_2_DATA_06 {
		t.Errorf("Test_AOC2020_11_day2_Ticks: Tick(5) not  conform\nExpected\n%v\n\nActual\n%v\n\n", TEST_11_2_DATA_06, actualTick5)
	} else {
		fmt.Printf("OK Test_AOC2020_11_day2_Ticks: Tick(5).\n")
	}

	changes = sp.Tick()
	actualTick6 := sp.Debug()
	if actualTick6 != TEST_11_2_DATA_07 {
		t.Errorf("Test_AOC2020_11_day2_Ticks: Tick(6) not conform\nExpected\n%v\n\nActual\n%v\n\n", TEST_11_2_DATA_07, actualTick6)
	} else {
		fmt.Printf("OK Test_AOC2020_11_day2_Ticks: Tick(6).\n")
	}

	fmt.Printf("Changes %v\n", changes)

}
