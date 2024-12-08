package d8

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
Day 8: Resonant Collinearity
*/

type Puzzle struct {
	title string
	year  int
	day   int
	input string
	lines []string
}

type Grid struct {
	cells    map[string]*Cell
	antennas map[string][]*Cell
	width    int
	height   int
}

func (grid *Grid) isCellInGrid(cell *Cell) bool {
	return cell.x < grid.width && cell.x >= 0 && cell.y >= 0 && cell.y < grid.height
}

func (grid *Grid) getCell(x int, y int) *Cell {
	key := fmt.Sprintf("%v.%v", x, y)
	cell := grid.cells[key]
	if cell == nil {
		cell = NewCell(x, y, ".")
		grid.cells[key] = cell
	}
	return cell
}

func NewGrid(data []string) *Grid {
	g := Grid{}
	g.height = len(data)
	g.width = len(data[0])
	g.cells = make(map[string]*Cell)
	g.antennas = make(map[string][]*Cell)
	for row_index, row := range data {
		for col_index := range row {
			col := row[col_index : col_index+1]
			cell := NewCell(col_index, row_index, col)
			g.cells[cell.key] = cell
			if cell.isAntenna() {
				antennas := g.antennas[cell.cellType]
				if antennas == nil {
					antennas = make([]*Cell, 0)
					g.antennas[cell.cellType] = antennas
				}
				antennas = append(antennas, cell)
				g.antennas[cell.cellType] = antennas
			}
		}
	}
	return &g
}

type Cell struct {
	x            int
	y            int
	key          string
	cellType     string
	antinodeCell *Cell
}

func NewCell(x int, y int, cellType string) *Cell {
	c := Cell{}
	c.x = x
	c.y = y
	c.key = fmt.Sprintf("%v.%v", x, y)
	c.cellType = cellType
	return &c
}

func (c *Cell) isAntenna() bool {
	return c.cellType != "."
}

func (puzzle *Puzzle) GetSummary() utils.Summary {
	s := utils.Summary{Day: puzzle.day, Year: puzzle.year, Name: puzzle.title}
	s.ProgressP1 = utils.NotStarted
	s.ProgressP2 = utils.NotStarted
	s.DateStarted = "2024-12-08 09:20:58"
	return s
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("2024")
	iday, _ := strconv.Atoi("8")
	p := Puzzle{year: iyear, day: iday, title: "Day 8: Resonant Collinearity"}
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
	// puzzle.p1(TEST_DATA)
	puzzle.p1(REAL_DATA)
}

func (puzzle *Puzzle) p1(data string) {
	puzzle.Load(data)
	count := 0
	grid := NewGrid(puzzle.lines)

	keys := make(map[string]bool)

	for cellType, antennas := range grid.antennas {
		cellPairs := pairs(antennas)
		fmt.Printf("\nFor antenna type '%v', there are %v antennas, as %v pairs\n", cellType, len(antennas), len(cellPairs))
		for _, cp := range cellPairs {
			fmt.Printf("%v\n", cp.string())
		}

		for _, pair := range cellPairs {
			a1, a2 := calculate_antinodes(pair)
			cell1 := grid.getCell(a1.x, a1.y)
			if grid.isCellInGrid(cell1) {
				keys[cell1.key] = true
				count++
			}

			cell2 := grid.getCell(a2.x, a2.y)
			if grid.isCellInGrid(cell2) {
				keys[cell2.key] = true
				count++
			}
		}

	}
	fmt.Printf("\nThere are %v antinodes in the grid, %v total keys.\n", count, len(keys))
}

func (puzzle *Puzzle) Part2() {
	puzzle.Load(REAL_DATA)
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}

type CellPair struct {
	c1 *Cell
	c2 *Cell
}

func (cp *CellPair) string() string {
	c1 := cp.c1
	c2 := cp.c2
	return fmt.Sprintf("[(%v,%v),(%v,%v)] ", c1.x, c1.y, c2.x, c2.y)
}

func pairs(values []*Cell) []*CellPair {
	pairs := make([]*CellPair, 0)
	// sort.Slice(values, func(i int, j int) bool {
	// 	c1 := values[i]
	// 	c2 := values[j]
	// 	return c1.x < c2.x
	// })

	for index, cell_a := range values {
		for _, cell_b := range values[index:] {
			if cell_a != cell_b {
				pair := CellPair{c1: cell_a, c2: cell_b}
				pairs = append(pairs, &pair)
			}
		}
	}
	return pairs
}

func calculate_antinodes(pair *CellPair) (*Cell, *Cell) {
	c1 := pair.c1
	c2 := pair.c2
	x := c1.x - c2.x
	y := c1.y - c2.y
	if x < 0 {
		x *= -1
	}
	if y < 0 {
		y *= -1
	}

	min_x := utils.MinInt(c1.x, c2.x)
	min_y := utils.MinInt(c1.y, c2.y)
	max_x := utils.MaxInt(c1.x, c2.x)
	max_y := utils.MaxInt(c1.y, c2.y)

	a1_x := min_x - x
	a1_y := min_y - y
	a2_x := max_x + x
	a2_y := max_y + y

	a1 := NewCell(a1_x, a1_y, ".")
	a2 := NewCell(a2_x, a2_y, ".")

	return a1, a2

}
