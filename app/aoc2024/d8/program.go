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
	cells        map[string]*Cell
	antennas     map[string][]*Cell
	width        int
	height       int
	antennaCount int
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
				g.antennaCount += 1
			}
		}
	}
	return &g
}

type Cell struct {
	x        int
	y        int
	key      string
	cellType string
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

func (puzzle *Puzzle) p2(data string) {
	puzzle.Load(data)
	count := 0
	grid := NewGrid(puzzle.lines)

	// keys := make(map[string]bool)

	antinodes := make(map[string]*Cell)
	for _, antennas := range grid.antennas {
		cellPairs := pairs(antennas)
		// fmt.Printf("\nFor antenna type '%v', there are %v antennas, as %v pairs\n", cellType, len(antennas), len(cellPairs))
		// for _, cp := range cellPairs {
		// 	fmt.Printf("%v\n", cp.string())
		// }

		for _, pair := range cellPairs {
			results := calculate_antinodes2(grid.width, grid.height, pair)
			// fmt.Printf("  pair %v has %v antinodes.\n", pair.string(), len(results))
			count += len(results)
			antinodes[pair.c1.key] = pair.c1
			antinodes[pair.c2.key] = pair.c2
			for _, antinode := range results {
				antinodes[antinode.key] = antinode
			}
		}

	}

	for row := 0; row < grid.height; row++ {
		for col := 0; col < grid.width; col++ {
			key := fmt.Sprintf("%v.%v", col, row)
			cell, exists := antinodes[key]
			if exists {
				cell_antenna, antenna_exists := grid.cells[key]
				if antenna_exists && cell_antenna.isAntenna() {
					fmt.Print(cell_antenna.cellType)
				} else {
					fmt.Print(cell.cellType)
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	// for key, v := range antinodes {
	// 	fmt.Printf("%v, %v\n", key, v)
	// }
	fmt.Printf("There are %v antinodes in the grid, %v total antinodes, %v count.\n", len(antinodes), len(antinodes)+len(grid.antennas), count)
	// fmt.Printf("Count=%v, antennas=%v\n", count, grid.antennaCount)
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

func calculate_antinodes2(width int, height int, pair *CellPair) []*Cell {

	c1 := pair.c1
	c2 := pair.c2
	x_offset := c1.x - c2.x
	y_offset := c1.y - c2.y

	results := make([]*Cell, 0)

	c1x := c1.x
	c1y := c1.y
	c2x := c2.x
	c2y := c2.y
	for {

		c1x = c1x + x_offset
		c1y = c1y + y_offset
		c2x = c2x - x_offset
		c2y = c2y - y_offset

		a1 := NewCell(c1x, c1y, "#")
		a2 := NewCell(c2x, c2y, "#")

		if a1.x < 0 || a1.x >= width || a1.y < 0 || a1.y >= height {
			a1 = nil
		} else {
			results = append(results, a1)
		}

		if a2.x < 0 || a2.x >= width || a2.y < 0 || a2.y >= height {
			a2 = nil
		} else {
			results = append(results, a2)
		}

		if a1 == nil && a2 == nil {
			break
		}
	}

	return results

}

func (puzzle *Puzzle) Part1() {
	puzzle.p1(REAL_DATA)
}

func (puzzle *Puzzle) Part2() {
	fmt.Println("TEST_DATA: ")
	puzzle.p2(TEST_DATA)

	fmt.Println("TEST_DATA_P2_0: ")
	puzzle.p2(TEST_DATA_P2_O)

	fmt.Println("REAL_DATA: ")
	puzzle.p2(REAL_DATA)

}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}
