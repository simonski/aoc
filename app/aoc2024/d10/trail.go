package d10

import (
	"fmt"
	"strconv"
	"strings"
)

type TrailMap struct {
	cells  map[string]*Cell
	starts []*Cell
	ends   []*Cell
	paths  [][]*Cell
}

type Cell struct {
	x     int
	y     int
	value int
	key   string
}

func NewCell(x int, y int, value int) *Cell {
	c := Cell{x: x, y: y, value: value, key: fmt.Sprintf("%v.%v", x, y)}
	return &c
}
func (m *TrailMap) GetCell(x int, y int) *Cell {
	key := fmt.Sprintf("%v.%v", x, y)
	return m.cells[key]
}

func NewTrailMap(data string) *TrailMap {
	rows := strings.Split(data, "\n")
	cells := make(map[string]*Cell)
	starts := make([]*Cell, 0)
	ends := make([]*Cell, 0)
	for row_index, row := range rows {
		for col_index := 0; col_index < len(row); col_index++ {
			value, _ := strconv.Atoi(row[col_index : col_index+1])
			cell := NewCell(col_index, row_index, value)
			// fmt.Printf("cell(%v, %v=%v)\n", col_index, row_index, value)
			cells[cell.key] = cell
			if cell.value == 0 {
				starts = append(starts, cell)
			} else if cell.value == 9 {
				ends = append(ends, cell)
			}
		}
	}
	paths := make([][]*Cell, 0)
	trail := TrailMap{cells: cells, starts: starts, ends: ends, paths: paths}
	return &trail
}

func (tm *TrailMap) neighbours(c *Cell) []*Cell {
	left := tm.GetCell(c.x-1, c.y)
	right := tm.GetCell(c.x+1, c.y)
	up := tm.GetCell(c.x, c.y-1)
	down := tm.GetCell(c.x, c.y+1)
	results := make([]*Cell, 0)
	if left != nil && left.value == c.value+1 {
		results = append(results, left)
	}
	if right != nil && right.value == c.value+1 {
		results = append(results, right)
	}
	if up != nil && up.value == c.value+1 {
		results = append(results, up)
	}
	if down != nil && down.value == c.value+1 {
		results = append(results, down)
	}
	return results
}

func (tm *TrailMap) walk(cell *Cell, path []*Cell) {
	path = append(path, cell)
	// indent := strings.Repeat(" ", len(path))
	// fmt.Printf("%v%v \n", indent, debugPath(path))
	neighbours := tm.neighbours(cell)
	for _, n := range neighbours {
		if n.value == 9 {
			// start := path[0]
			// fmt.Printf("%v (%v,%v)->(%v,%v)=%v - end of trail.\n", indent, start.x, start.y, n.x, n.y, n.value)
			newPath := make([]*Cell, 0)
			newPath = append(newPath, path...)
			newPath = append(newPath, n)
			tm.paths = append(tm.paths, newPath)
			// fmt.Printf("There are now %v paths.\n", len(tm.paths))
			// fmt.Println(debugPath(newPath))
			// fmt.Println()

		} else {
			newPath := make([]*Cell, 0)
			newPath = append(newPath, path...)
			// fmt.Printf("%v %v - next move.\n", indent, debugPath(newPath))
			tm.walk(n, newPath)
		}
	}
}

func debugPath(path []*Cell) string {
	s := ""
	for _, c := range path {
		s += " -> " + fmt.Sprintf("(%v,%v) [%v]", c.x, c.y, c.value)
	}
	return s
}
