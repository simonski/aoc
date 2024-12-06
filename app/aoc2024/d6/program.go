package d6

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
Day 6: Guard Gallivant
*/

type Puzzle struct {
	title string
	year  int
	day   int
	input string
	lines []string
}

type Grid struct {
	cells            map[string]*Cell
	width            int
	height           int
	guardX           int
	guardY           int
	guardOrientation string
}

type Cell struct {
	x          int
	y          int
	cellType   string
	visitCount int
	visited    bool
}

const NORTH string = "^"
const SOUTH string = "v"
const EAST string = ">"
const WEST string = "<"
const FURNIURE string = "#"
const OTHER string = "."

func NewGrid(lines []string) *Grid {
	grid := Grid{}
	grid.height = len(lines)
	grid.width = len(lines[0])
	grid.cells = make(map[string]*Cell)
	for y, row := range lines {
		for x := 0; x < len(row); x++ {
			cellType := row[x : x+1]
			cell := NewCell(x, y, cellType)
			key := fmt.Sprintf("%v.%v", x, y)
			grid.cells[key] = cell
			if isGuard(cellType) {
				grid.guardX = x
				grid.guardY = y
				grid.guardOrientation = cellType
				cell.cellType = OTHER
			}
		}
	}
	return &grid
}

func (grid *Grid) Debug(showVisited bool) string {
	result := ""
	for y := 0; y < grid.height; y++ {
		for x := 0; x < grid.width; x++ {
			key := fmt.Sprintf("%v.%v", x, y)
			cell := grid.cells[key]
			if grid.guardX == x && grid.guardY == y {
				result += grid.guardOrientation
			} else {
				if showVisited && cell.visited {
					result += "X"
				} else {
					result += cell.cellType

				}
			}
		}
		result += "\n"
	}
	return result
}

func (grid *Grid) Tick() bool {
	if grid.guardX < 0 || grid.guardX+1 == grid.width || grid.guardY < 0 || grid.guardY+1 == grid.height {
		return false
	}
	if isNorth(grid.guardOrientation) {
		key := fmt.Sprintf("%v.%v", grid.guardX, grid.guardY-1)
		targetCell := grid.cells[key]
		if isFurniture(targetCell.cellType) {
			grid.guardOrientation = EAST
		} else {
			grid.guardY = targetCell.y
			targetCell.visitCount += 1
			targetCell.visited = true
		}
	} else if isSouth(grid.guardOrientation) {
		key := fmt.Sprintf("%v.%v", grid.guardX, grid.guardY+1)
		targetCell := grid.cells[key]
		if isFurniture(targetCell.cellType) {
			grid.guardOrientation = WEST
		} else {
			grid.guardY = targetCell.y
			targetCell.visitCount += 1
			targetCell.visited = true
		}
	} else if isEast(grid.guardOrientation) {
		key := fmt.Sprintf("%v.%v", grid.guardX+1, grid.guardY)
		targetCell := grid.cells[key]
		if isFurniture(targetCell.cellType) {
			grid.guardOrientation = SOUTH
		} else {
			grid.guardX = targetCell.x
			targetCell.visitCount += 1
			targetCell.visited = true
		}
	} else if isWest(grid.guardOrientation) {
		key := fmt.Sprintf("%v.%v", grid.guardX-1, grid.guardY)
		targetCell := grid.cells[key]
		if isFurniture(targetCell.cellType) {
			grid.guardOrientation = NORTH
		} else {
			grid.guardX = targetCell.x
			targetCell.visitCount += 1
			targetCell.visited = true
		}
	}
	return true
}

func isGuard(s string) bool {
	return isNorth(s) || isSouth(s) || isEast(s) || isWest(s)
}

func isFurniture(s string) bool {
	return s == FURNIURE
}

func isNorth(s string) bool {
	return s == NORTH
}

func isSouth(s string) bool {
	return s == SOUTH
}

func isWest(s string) bool {
	return s == WEST
}

func isEast(s string) bool {
	return s == EAST
}

func NewCell(x int, y int, cellType string) *Cell {
	c := Cell{}
	c.x = x
	c.y = y
	c.cellType = cellType
	return &c
}

func (puzzle *Puzzle) GetSummary() utils.Summary {
	s := utils.Summary{Day: puzzle.day, Year: puzzle.year, Name: puzzle.title}
	s.ProgressP1 = utils.NotStarted
	s.ProgressP2 = utils.NotStarted
	s.DateStarted = "2024-12-06 08:15:03"
	return s
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("2024")
	iday, _ := strconv.Atoi("6")
	p := Puzzle{year: iyear, day: iday, title: "Day 6: Guard Gallivant"}
	p.Load(input)
	return &p
}

func NewPuzzle() *Puzzle {
	return NewPuzzleWithData(REAL_DATA)
}

func (puzzle *Puzzle) Load(input string) {
	lines := strings.Split(input, "\n")
	puzzle.input = input
	puzzle.lines = lines
}

func (puzzle *Puzzle) Part1() {
	// puzzle.Load(TEST_DATA_1)
	puzzle.Load(REAL_DATA)
	grid := NewGrid(puzzle.lines)

	key := fmt.Sprintf("%v.%v", grid.guardX, grid.guardY)
	startCell := grid.cells[key]
	startCell.visited = true
	startCell.visitCount = 1

	for {
		// fmt.Println(grid.Debug())
		// fmt.Println("")

		if !grid.Tick() {
			break
		}
	}
	total := 0
	distinct := 0
	for _, cell := range grid.cells {
		if cell.visited {
			distinct += 1
		}
		total += cell.visitCount
	}
	fmt.Println(grid.Debug(true))
	fmt.Printf("%v visits, %v distinct visits\n", total, distinct)
}

func (puzzle *Puzzle) Part2() {
	puzzle.Load(REAL_DATA)
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}
