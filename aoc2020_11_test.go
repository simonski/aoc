package main

import (
	"fmt"
	"strings"
	"testing"
)

const TEST_11_DATA_01 = `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

const TEST_11_DATA_02 = `#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`

const TEST_11_DATA_03 = `#.LL.L#.##
#LLLLLL.L#
L.L.L..L..
#LLL.LL.L#
#.LL.LL.LL
#.LLLL#.##
..L.L.....
#LLLLLLLL#
#.LLLLLL.L
#.#LLLL.##`

const TEST_11_DATA_04 = `#.##.L#.##
#L###LL.L#
L.#.#..#..
#L##.##.L#
#.##.LL.LL
#.###L#.##
..#.#.....
#L######L#
#.LL###L.L
#.#L###.##`

const TEST_11_DATA_05 = `#.#L.L#.##
#LLL#LL.L#
L.L.L..#..
#LLL.##.L#
#.LL.LL.LL
#.LL#L#.##
..L.L.....
#L#LLLL#L#
#.LLLLLL.L
#.#L#L#.##`

const TEST_11_DATA_06 = `#.#L.L#.##
#LLL#LL.L#
L.#.L..#..
#L##.##.L#
#.#L.LL.LL
#.#L#L#.##
..L.L.....
#L#L##L#L#
#.LLLLLL.L
#.#L#L#.##`

// L.LL.LL.LL
// LLLLLLL.LL
// L.L.L..L..
// LLLL.LL.LL
// L.LL.LL.LL
// L.LLLLL.LL
// ..L.L.....
// LLLLLLLLLL
// L.LLLLLL.L
// L.LLLLL.LL
func Test_AOC2020_11_day1_IndexFromColRow(t *testing.T) {
	width := 10
	sp := NewSeatingPlanFromStrings(strings.Split(TEST_11_DATA_01, "\n"), width)
	verifyIndexFromColRow(0, 0, 0, sp, t)
	verifyIndexFromColRow(1, 1, 0, sp, t)
	verifyIndexFromColRow(2, 2, 0, sp, t)
	verifyIndexFromColRow(3, 3, 0, sp, t)
	verifyIndexFromColRow(-1, 10, 0, sp, t)

	verifyIndexFromColRow(-1, -10, 0, sp, t)
	verifyIndexFromColRow(-1, -10, -10, sp, t)
	verifyIndexFromColRow(-1, 10, 0, sp, t)
	verifyIndexFromColRow(-1, 5, -1, sp, t)
	verifyIndexFromColRow(-1, 3, 100, sp, t)

	verifyIndexFromColRow(55, 5, 5, sp, t)
}

func verifyIndexFromColRow(expectedIndex int, col int, row int, sp *SeatingPlan, t *testing.T) {
	actual := sp.IndexFromColRow(col, row)
	if actual != expectedIndex {
		t.Errorf("Test_AOC2020_11_day1_IndexFromColRow(%v,%v)=%v, expected %v.\n", col, row, actual, expectedIndex)
	} else {
		fmt.Printf("OK Test_AOC2020_11_day1_IndexFromColRow(%v,%v)=%v, expected %v.\n", col, row, actual, expectedIndex)
	}
}

func verifyColRowFromIndex(index int, expectedCol int, expectedRow int, sp *SeatingPlan, t *testing.T) {
	actualCol, actualRow := sp.ColRowFromIndex(index)
	if actualCol != expectedCol || actualRow != expectedRow {
		t.Errorf("verifyColRowFromIndex (%v), expected (%v,%v) actual (%v,%v).\n", index, expectedCol, expectedRow, actualCol, actualRow)
	} else {
		fmt.Printf("OK verifyColRowFromIndex (%v), expected (%v,%v) actual (%v,%v).\n", index, expectedCol, expectedRow, actualCol, actualRow)
	}
}

// L.LL.LL.LL
// LLLLLLL.LL
// L.L.L..L..
// LLLL.LL.LL
// L.LL.LL.LL
// L.LLLLL.LL
// ..L.L.....
// LLLLLLLLLL
// L.LLLLLL.L
// L.LLLLL.LL
func Test_AOC2020_11_day1_loading_and_debugs(t *testing.T) {

	width := 10
	sp := NewSeatingPlanFromStrings(strings.Split(TEST_11_DATA_01, "\n"), width)
	verifyContentsAtColRow(0, 0, EMPTY, sp, t)
	verifyContentsAtColRow(1, 0, FLOOR, sp, t)
	verifyContentsAtColRow(2, 0, EMPTY, sp, t)
	verifyContentsAtColRow(-2, 0, NONE, sp, t)
	verifyContentsAtColRow(4, 0, FLOOR, sp, t)
	verifyContentsAtColRow(4, 3, FLOOR, sp, t)
	verifyContentsAtColRow(54, 3, NONE, sp, t)
	verifyContentsAtColRow(-54, 3, NONE, sp, t)
	verifyContentsAtColRow(54, -3, NONE, sp, t)
	verifyContentsAtColRow(-54, -3, NONE, sp, t)
	verifyContentsAtColRow(54, 999, NONE, sp, t)
	verifyContentsAtColRow(-54, -999, NONE, sp, t)

}

func verifyContentsAtColRow(col int, row int, expected int, sp *SeatingPlan, t *testing.T) {
	actual := sp.Get(col, row)
	if actual != expected {
		t.Errorf("verifyContentsAtColRow(%v,%v) expected %v actual %v\n", col, row, sp.Translate(actual), sp.Translate(expected))
	} else {
		fmt.Printf("OK verifyContentsAtColRow(%v,%v) expected %v actual %v\n", col, row, sp.Translate(actual), sp.Translate(expected))
	}
}

// // top

// // #.##.##.##
// // #######.##
func Test_AOC2020_11_conversion_scenarios(t *testing.T) {
	testdata := strings.Split("#.##.##.##\n#######.##", "\n")
	sp := NewSeatingPlanFromStrings(testdata, 10)

	// top left
	checkValueMatches(sp, "left", sp.left(0), NONE, t)
	checkValueMatches(sp, "right", sp.right(0), FLOOR, t)
	checkValueMatches(sp, "up", sp.up(0), NONE, t)
	checkValueMatches(sp, "down", sp.down(0), OCCUPIED, t)
	checkValueMatches(sp, "upperleft", sp.upperleft(0), NONE, t)
	checkValueMatches(sp, "uperright", sp.upperright(0), NONE, t)
	checkValueMatches(sp, "lowerleft", sp.lowerleft(0), NONE, t)
	checkValueMatches(sp, "loweright", sp.lowerright(0), OCCUPIED, t)
	checkValueBeforeAfter(sp, 0, 2, OCCUPIED, OCCUPIED, t)

	// first left
	index := 1
	col := 1
	row := 0
	col, row = sp.ColRowFromIndex(index)
	verifyColRowFromIndex(index, col, row, sp, t)
	verifyIndexFromColRow(index, col, row, sp, t)
	verifyContentsAtColRow(col, row, FLOOR, sp, t)

	checkValueMatches(sp, "left", sp.left(1), OCCUPIED, t)

	checkValueMatches(sp, "right", sp.right(1), OCCUPIED, t)
	checkValueMatches(sp, "up", sp.up(1), NONE, t)
	checkValueMatches(sp, "down", sp.down(1), OCCUPIED, t)
	checkValueMatches(sp, "upperleft", sp.upperleft(1), NONE, t)
	checkValueMatches(sp, "uperright", sp.upperright(1), NONE, t)
	checkValueMatches(sp, "lowerleft", sp.lowerleft(1), OCCUPIED, t)
	checkValueMatches(sp, "loweright", sp.lowerright(1), OCCUPIED, t)
	checkValueBeforeAfter(sp, 1, 5, FLOOR, FLOOR, t)

	// top right
	checkValueMatches(sp, "left", sp.left(9), OCCUPIED, t)
	checkValueMatches(sp, "right", sp.right(9), NONE, t)
	checkValueMatches(sp, "up", sp.up(9), NONE, t)
	checkValueMatches(sp, "down", sp.down(9), OCCUPIED, t)
	checkValueMatches(sp, "upperleft", sp.upperleft(9), NONE, t)
	checkValueMatches(sp, "uperright", sp.upperright(9), NONE, t)
	checkValueMatches(sp, "lowerleft", sp.lowerleft(9), OCCUPIED, t)
	checkValueMatches(sp, "loweright", sp.lowerright(9), NONE, t)
	checkValueBeforeAfter(sp, 9, 3, OCCUPIED, OCCUPIED, t)

	// 2nd row left
	checkValueMatches(sp, "left", sp.left(10), NONE, t)
	checkValueMatches(sp, "right", sp.right(10), OCCUPIED, t)
	checkValueMatches(sp, "up", sp.up(10), OCCUPIED, t)
	checkValueMatches(sp, "down", sp.down(10), NONE, t)
	checkValueMatches(sp, "upperleft", sp.upperleft(10), NONE, t)
	checkValueMatches(sp, "uperright", sp.upperright(10), FLOOR, t)
	checkValueMatches(sp, "lowerleft", sp.lowerleft(10), NONE, t)
	checkValueMatches(sp, "loweright", sp.lowerright(10), NONE, t)
	checkValueBeforeAfter(sp, 10, 2, OCCUPIED, OCCUPIED, t)
}

func checkValueMatches(sp *SeatingPlan, description string, actualValue int, expectedValue int, t *testing.T) {
	if actualValue != expectedValue {
		t.Errorf("Value %v='%v' is not correct to start with (%v).\n", description, sp.Translate(actualValue), sp.Translate(expectedValue))
	}
}

func checkValueBeforeAfter(sp *SeatingPlan, index int, expectedOccupiedCount int, expectedBefore int, expectedAfter int, t *testing.T) {
	col, row := sp.ColRowFromIndex(index)
	if sp.Get(col, row) != expectedBefore {
		t.Errorf("ConvertValue: Value at index %v %v is not correct to start with.\n", index, sp.Translate(expectedBefore))
	} else {
		fmt.Printf("ConvertValue: Value at index %v %v is correct to start with.\n", index, sp.Translate(expectedBefore))
	}

	actualOccupiedCount := sp.CountOccupied(index)
	if actualOccupiedCount != expectedOccupiedCount {
		t.Errorf("ConvertValue: Occupied count at index %v is expected to be %v but was %v.\n", index, expectedOccupiedCount, actualOccupiedCount)
	} else {
		fmt.Printf("ConvertValue: Occupied count is correct at %v.\n", expectedOccupiedCount)
	}

	actualAfter := sp.ConvertAtIndex(index)
	if actualAfter != expectedAfter {
		t.Errorf("ConvertValue: at index %v %v -> %v is not correct, converted to %v.\n", index, sp.Translate(expectedBefore), sp.Translate(expectedAfter), sp.Translate(actualAfter))
	} else {
		fmt.Printf("ConvertValue: at index %v %v -> %v is correct.\n", index, sp.Translate(expectedBefore), sp.Translate(expectedAfter))
	}
}

func Test_AOC2020_11_day_ticks(t *testing.T) {

	width := 10
	sp := NewSeatingPlanFromStrings(strings.Split(TEST_11_DATA_01, "\n"), width)
	actualTick0 := sp.Debug()
	if actualTick0 != TEST_11_DATA_01 {
		t.Errorf("Loaded data does not conform\nExpected\n%v\n\nActual\n%v\n\n", TEST_11_DATA_01, actualTick0)
	} else {
		fmt.Printf("Day 11: Loaded data conforms to expected format.\n")
	}

	// Tick 1 -> 2
	changes := sp.Tick()
	fmt.Printf("Tick(1): increments tick count to %v, changes %v.\n", sp.TickCount, changes)
	actualTick1 := sp.Debug()
	if actualTick1 != TEST_11_DATA_02 {
		t.Errorf("Tick(1): does not conform\nExpected\n%v\n\nActual\n%v\n\n", TEST_11_DATA_02, actualTick1)
	} else {
		fmt.Printf("Tick(1): output conforms as expected.\n%v\n\n", actualTick1)
	}

	// Tick 2 -> 3
	changes = sp.Tick()
	fmt.Printf("Tick(2): increments tick count to %v, changes %v.\n", sp.TickCount, changes)
	actualTick2 := sp.Debug()
	if actualTick2 != TEST_11_DATA_03 {
		t.Errorf("Tick(2): does not conform\n\nGave Input\n%v\n\nExpected Output\n%v\n\nActual\n%v\n\n", actualTick1, TEST_11_DATA_03, actualTick2)
	} else {
		fmt.Printf("Tick(2): output conforms as expected.\n%v\n\n", actualTick2)
	}

	changes = sp.Tick()
	actualTick3 := sp.Debug()
	if actualTick3 != TEST_11_DATA_04 {
		t.Errorf("Tick(3): does not conform\n\nGave Input\n%v\n\nExpected Output\n%v\n\nActual\n%v\n\n", actualTick2, TEST_11_DATA_04, actualTick3)
	} else {
		fmt.Printf("Tick(3): output conforms as expected.\n%v\n\n", actualTick3)
	}

	changes = sp.Tick()
	actualTick4 := sp.Debug()
	if actualTick4 != TEST_11_DATA_05 {
		t.Errorf("Tick(4): does not conform\n\nGave Input\n%v\n\nExpected Output\n%v\n\nActual\n%v\n\n", actualTick3, TEST_11_DATA_05, actualTick4)
	} else {
		fmt.Printf("Tick(4): output conforms as expected.\n%v\n\n", actualTick4)
	}

	changes = sp.Tick()
	actualTick5 := sp.Debug()
	if actualTick5 != TEST_11_DATA_06 {
		t.Errorf("Tick(5): does not conform\n\nGave Input\n%v\n\nExpected Output\n%v\n\nActual\n%v\n\n", actualTick4, TEST_11_DATA_06, actualTick5)
	} else {
		fmt.Printf("Tick(5): output conforms as expected.\n%v\n\n", actualTick5)
	}

	changes = sp.Tick()
	actualTick6 := sp.Debug()
	if actualTick6 != TEST_11_DATA_06 {
		t.Errorf("Tick(6): does not conform\n\nGave Input\n%v\n\nExpected Output\n%v\n\nActual\n%v\n\n", actualTick5, TEST_11_DATA_06, actualTick6)
	} else {
		fmt.Printf("Tick(6): output conforms as expected.\n%v\n\n", actualTick6)
	}
	if changes > 0 {
		t.Errorf("Tick(6) should not change anything, actual=%v\n", changes)
	} else {
		fmt.Printf("Tick(6): has frozen.\n")
	}
	fmt.Printf("Occupied Count is %v\n", sp.GetOccupiedCount())

}
