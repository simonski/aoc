package aoc2021

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
--- Day 9: Smoke Basin ---
These caves seem to be lava tubes. Parts are even still volcanically active; small hydrothermal vents release smoke into the caves that slowly settles like rain.

If you can model how the smoke flows through the caves, you might be able to avoid it and be that much safer. The submarine generates a heightmap of the floor of the nearby caves for you (your puzzle input).

Smoke flows to the lowest point of the area it's in. For example, consider the following heightmap:

2199943210
3987894921
9856789892
8767896789
9899965678
Each number corresponds to the height of a particular location, where 9 is the highest and 0 is the lowest a location can be.

Your first goal is to find the low points - the locations that are lower than any of its adjacent locations. Most locations have four adjacent locations (up, down, left, and right); locations on the edge or corner of the map have three or two adjacent locations, respectively. (Diagonal locations do not count as adjacent.)

In the above example, there are four low points, all highlighted: two are in the first row (a 1 and a 0), one is in the third row (a 5), and one is in the bottom row (also a 5). All other locations on the heightmap have some lower adjacent location, and so are not low points.

The risk level of a low point is 1 plus its height. In the above example, the risk levels of the low points are 2, 1, 6, and 6. The sum of the risk levels of all low points in the heightmap is therefore 15.

Find all of the low points on your heightmap. What is the sum of the risk levels of all low points on your heightmap?
*/

type Point struct {
	X     int
	Y     int
	Value int
	Taken bool
}

type Day9Grid struct {
	Data   [][]*Point
	Height int
	Width  int
}

func (grid *Day9Grid) Up(p *Point) *Point {
	if p.Y > 0 {
		return grid.Data[p.Y-1][p.X]
	} else {
		return nil
	}
}

func (grid *Day9Grid) Down(p *Point) *Point {
	if p.Y+1 < grid.Height {
		return grid.Data[p.Y+1][p.X]
	} else {
		return nil
	}
}

func (grid *Day9Grid) Left(p *Point) *Point {
	if p.X > 0 {
		return grid.Data[p.Y][p.X-1]
	} else {
		return nil
	}
}

func (grid *Day9Grid) Right(p *Point) *Point {
	if p.X+1 < grid.Width {
		return grid.Data[p.Y][p.X+1]
	} else {
		return nil
	}
}

func (g *Day9Grid) Value(col int, row int) int {
	return g.Data[row][col].Value
}

func (g *Day9Grid) IsLowest(col int, row int) bool {
	sample_row := g.Data[0]
	MAX_COLS := len(sample_row) - 1
	MAX_ROWS := len(g.Data) - 1
	value := g.Value(col, row)
	if row-1 >= 0 {
		up := g.Value(col, row-1)
		if value >= up {
			return false
		}
	}
	if row+1 <= MAX_ROWS {
		down := g.Value(col, row+1)
		if value >= down {
			return false
		}
	}
	if col-1 >= 0 {
		left := g.Value(col-1, row)
		if value >= left {
			return false
		}
	}
	if col+1 <= MAX_COLS {
		right := g.Value(col+1, row)
		if value >= right {
			return false
		}
	}
	// if col+1 < MAX_COLS && row-1 >= 0 {
	// 	upright := g.Value(col+1, row-1)
	// 	if value >= upright {
	// 		return false
	// 	}
	// }
	// if col-1 >= 0 && row-1 >= 0 {
	// 	upleft := g.Value(col-1, row-1)
	// 	if value >= upleft {
	// 		return false
	// 	}
	// }
	// if col+1 < MAX_COLS && row+1 < MAX_ROWS {
	// 	downright := g.Value(col+1, row+1)
	// 	if value >= downright {
	// 		return false
	// 	}
	// }
	// if col-1 >= 0 && row+1 <= MAX_ROWS {
	// 	downleft := g.Value(col-1, row+1)
	// 	if value >= downleft {
	// 		return false
	// 	}
	// }
	return true
}

func (g *Day9Grid) Debug() string {
	line := ""
	for rowNum := 0; rowNum < len(g.Data); rowNum++ {
		row := g.Data[rowNum]
		for colNum := 0; colNum < len(row); colNum++ {
			value := row[colNum].Value
			line += fmt.Sprintf("%v ", value)
		}
		line += "\n"
	}
	return line
}

func (g *Day9Grid) DebugLowest() (string, int) {
	line := ""
	total := 0
	for rowNum := 0; rowNum < len(g.Data); rowNum++ {
		row := g.Data[rowNum]
		for colNum := 0; colNum < len(row); colNum++ {
			value := row[colNum].Value
			if g.IsLowest(colNum, rowNum) {
				total += g.Value(colNum, rowNum)
				total += 1
				line += fmt.Sprintf("*%v ", value)
			} else {
				line += fmt.Sprintf(" %v ", value)

			}
		}
		line += "\n"
	}
	return line, total
}

// build a [][]grid I can reference to get the value
func NewDay9Grid(data string) *Day9Grid {
	lines := strings.Split(data, "\n")
	rows := make([][]*Point, 0)
	for rowNum, line := range lines {
		line = strings.Trim(line, " ")
		row := make([]*Point, 0)
		for colNum := 0; colNum < len(line); colNum++ {
			c := line[colNum : colNum+1]
			i, _ := strconv.Atoi(c)
			row = append(row, &Point{Value: i, X: colNum, Y: rowNum})
		}
		rows = append(rows, row)
	}
	width := len(rows[0])
	height := len(rows)
	return &Day9Grid{Data: rows, Width: width, Height: height}
}

// func (g *Day9Grid) Basins() {
// 	point{x: int, y: int, value: int, taken: bool}
// }

type Basin struct {
	Points []*Point
}

func (b *Basin) Size() int {
	return len(b.Points)
}

func (b *Basin) take(p *Point) {
	b.Points = append(b.Points, p)
	p.Taken = true
}

func (b *Basin) fillUp(grid *Day9Grid, point *Point) {
	up := grid.Up(point)
	right := grid.Right(point)
	down := grid.Down(point)
	left := grid.Left(point)
	if up != nil && !up.Taken && up.Value != 9 { // } point.Value {
		// take it
		b.take(up)
		b.fillUp(grid, up)
	}
	if right != nil && !right.Taken && right.Value != 9 { // Value < point.Value {
		// take it
		b.take(right)
		b.fillUp(grid, right)
	}
	if down != nil && !down.Taken && down.Value != 9 { //  < point.Value {
		// take it
		b.take(down)
		b.fillUp(grid, down)
	}
	if left != nil && !left.Taken && left.Value != 9 { //  < point.Value {
		// take it
		b.take(left)
		b.fillUp(grid, left)
	}

}

// rename this to the year and day in question
func (app *Application) Y2021D09P1() {
	g := NewDay9Grid(DAY_2021_09_TEST_DATA)
	fmt.Printf("%v\n", g.Debug())
	fmt.Println()
	_, count := g.DebugLowest()
	// fmt.Printf("%v\n, count=%v\n", view, count)
	fmt.Printf("count=%v, width=%v, height=%v\n", count, g.Width, g.Height)
	fmt.Println()
	fmt.Println()

	g = NewDay9Grid(DAY_2021_09_DATA)
	fmt.Printf("%v\n", g.Debug())
	fmt.Println()
	_, count = g.DebugLowest()
	// fmt.Printf("%v\n, count=%v\n", view, count)
	fmt.Printf("count=%v, width=%v, height=%v\n", count, g.Width, g.Height)

}

// rename this to the year and day in question
func (app *Application) Y2021D09P2() {
	g := NewDay9Grid(DAY_2021_09_DATA)
	fmt.Printf("loaded grid, width=%v, height=%v\n", g.Width, g.Height)
	basins := make([]*Basin, 0)
	for rowNum := 0; rowNum < len(g.Data); rowNum++ {
		row := g.Data[rowNum]
		for colNum := 0; colNum < len(row); colNum++ {
			point := row[colNum]
			fmt.Printf("scanning[%v,%v], point=(%v,%v=%v)\n", colNum, rowNum, point.X, point.Y, point.Value)
			if point.Value == 9 {
				fmt.Printf("[%v][%v] == %v, ignore, (%v, %v)\n", colNum, rowNum, point.Value, point.X, point.Y)
				continue
			}
			if point.Taken {
				fmt.Printf("[%v][%v] == %v is taken, ignore, (%v, %v)\n", colNum, rowNum, point.Value, point.X, point.Y)
				continue
			}
			// ok so now wind in a clock recursively asking
			// am I flowing downwards to my neighbour?
			// if yes, move to the neighbour and "take" it
			basin := &Basin{}
			basin.fillUp(g, point)
			// basin.take(point)
			if basin.Size() > 1 {
				// basin.take(point)
				fmt.Printf("[%v][%v] basin is size %v\n", colNum, rowNum, basin.Size())
				basins = append(basins, basin)
			} else {
				fmt.Printf("[%v][%v] basin is size %v, ignore\n", colNum, rowNum, basin.Size())
			}
			fmt.Println()
		}
	}

	sizes := make([]int, 0)
	for _, basin := range basins {
		size := basin.Size()
		sizes = append(sizes, size)
	}
	index1 := len(sizes) - 1
	index2 := len(sizes) - 2
	index3 := len(sizes) - 3

	sort.Ints(sizes)

	value := sizes[index1] * sizes[index2] * sizes[index3]
	fmt.Printf("%v, value=%v\n", sizes, value)

}
