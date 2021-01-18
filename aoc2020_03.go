package main

import (
	"fmt"

	utils "github.com/simonski/aoc/utils"
	goutils "github.com/simonski/goutils"
)

// AOC_2020_03 is the entrypoint
func AOC_2020_03(cli *goutils.CLI) {
	verbose := cli.IndexOf("-v") > -1
	AOC_2020_03_part1_attempt1(cli, verbose)
	AOC_2020_03_part1_attempt2(cli, verbose)
}

// AOC_2020_01_part3_attempt1 this is part 1 of day 3, attempt 1
func AOC_2020_03_part1_attempt1(cli *goutils.CLI, verbose bool) {
	filename := cli.GetStringOrDie("-input")
	tm := NewTobogganMapFromFile(filename)
	tm.Debug()
	treeCount := tm.CountTreesWeEncounter(3, 1)
	fmt.Printf("For %v,%v We will encounter %v trees.\n", 3, 1, treeCount)
}

func AOC_2020_03_part1_attempt2(cli *goutils.CLI, verbose bool) {
	filename := cli.GetStringOrDie("-input")
	tm := NewTobogganMapFromFile(filename)
	treeCount1 := tm.CountTreesWeEncounter(1, 1)
	treeCount2 := tm.CountTreesWeEncounter(3, 1)
	treeCount3 := tm.CountTreesWeEncounter(5, 1)
	treeCount4 := tm.CountTreesWeEncounter(7, 1)
	treeCount5 := tm.CountTreesWeEncounter(1, 2)
	fmt.Printf("1x1: %v\n", treeCount1)
	fmt.Printf("3x1: %v\n", treeCount2)
	fmt.Printf("5x1: %v\n", treeCount3)
	fmt.Printf("7x1: %v\n", treeCount4)
	fmt.Printf("1x2: %v\n", treeCount5)
	fmt.Printf("Product of %v * %v * %v * %v * %v = %v\n", treeCount1, treeCount2, treeCount3, treeCount4, treeCount5, treeCount1*treeCount2*treeCount3*treeCount4*treeCount5)
}

// TobogganMap represents the whole snow field
type TobogganMap struct {
	rows []*TobogganMapRow
}

// NewTobogganMap constructs a new empty snowy field
func NewTobogganMap() *TobogganMap {
	tm := TobogganMap{}
	tm.rows = make([]*TobogganMapRow, 0)
	return &tm
}

// NewTobogganMapFromFile constructs and populates a snowy field from the passed filename
func NewTobogganMapFromFile(filename string) *TobogganMap {
	data := utils.Load_file_to_strings(filename)
	tm := NewTobogganMap()
	for index := 0; index < len(data); index++ {
		line := data[index]
		tm.Add(line)
	}
	return tm
}

// CountTreesWeEncounter a general purpose walk to the bottom and keep moving right
// where I moduluse the move right
func (tm *TobogganMap) CountTreesWeEncounter(moveColsBy int, moveRowsBy int) int {
	col := 0
	treeCount := 0
	for row := 0; row < tm.Size(); row += moveRowsBy {
		if tm.IsTree(col, row) {
			treeCount++
		}
		col += moveColsBy
	}
	return treeCount
}

// Debug prints to STDOUT the map
func (tm *TobogganMap) Debug() {
	for index := 0; index < tm.Size(); index++ {
		tm.GetRow(index).Debug()
	}
}

// Size tells me how deep the map is
func (tm *TobogganMap) Size() int {
	return len(tm.rows)
}

// Add appends a snowy line to the Map
func (tm *TobogganMap) Add(line string) int {
	row := TobogganMapRow{line: line}
	tm.rows = append(tm.rows, &row)
	return len(tm.rows)
}

// Get returns the entry at col/row
func (tm *TobogganMap) Get(col int, row int) string {
	return tm.rows[row].Get(col)
}

// IsTree indicates if we've landed on a tree
func (tm *TobogganMap) IsTree(col int, row int) bool {
	return tm.rows[row].IsTree(col)
}

// GetRow returns a specific 0-indexed row from the snowy field
func (tm *TobogganMap) GetRow(row int) *TobogganMapRow {
	return tm.rows[row]
}

// TobogganMapRow represents a single line of snow and trees
type TobogganMapRow struct {
	line string
}

// Get returns the thing that exists at position index; modulus the index as the row repeats
func (row *TobogganMapRow) Get(index int) string {
	searchIndex := index % len(row.line)
	return row.line[searchIndex : searchIndex+1]
}

// IsTree indicates if we have landed on a tree
func (row *TobogganMapRow) IsTree(col int) bool {
	value := row.Get(col)
	modCol := col % len(row.line)
	fmt.Printf("GetTree, line = '%v', col=%v (mod col %v), value=%v\n", row.line, col, modCol, value)
	return row.Get(col) == "#"
}

// Debug prints this snowly line out to STDOUT
func (row *TobogganMapRow) Debug() {
	fmt.Printf("%v\n", row.line)
}
