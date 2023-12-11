package d11

import (
	"fmt"
	"math"
	"strings"

	"github.com/simonski/aoc/utils"
)

const CELL_EMPTY = "."
const CELL_GALAXY = "#"

type Universe struct {
	data [][]*Cell
}

func NewUniverse(input string) *Universe {
	u := &Universe{}
	u.Load(input)
	return u
}

func (u *Universe) Load(input string) {
	lines := strings.Split(input, "\n")
	rows := make([][]*Cell, 0)
	for _, line := range lines {
		row := make([]*Cell, 0)
		for colNum := 0; colNum < len(line); colNum++ {
			value := line[colNum : colNum+1]
			cell := NewCell(value)
			row = append(row, cell)
		}
		rows = append(rows, row)
	}
	u.data = rows

}

func (u *Universe) Debug() string {
	result := ""
	for rowNum, row := range u.data {
		for colNum := range row {
			cell := u.Get(colNum, rowNum)
			result = fmt.Sprintf("%v%v", result, cell.cellType)
		}
		result = fmt.Sprintf("%v\n", result)
	}
	return result
}

func (u *Universe) Expand() {
	cols := len(u.data[0])
	rows := len(u.data)

	for col := cols - 1; col >= 0; col-- {
		if u.IsColEmpty(col) {
			u.AddColumn(col)
		}
	}

	for row := rows - 1; row >= 0; row-- {
		if u.IsRowEmpty(row) {
			u.AddRow(row)
		}
	}
}

func (u *Universe) Get(x int, y int) *Cell {
	return u.data[y][x]
}

func (u *Universe) IsRowEmpty(row int) bool {
	cols := len(u.data[0])
	for col := 0; col < cols; col++ {
		if u.Get(col, row).cellType != CELL_EMPTY {
			return false
		}

	}
	return true
}

func (u *Universe) IsColEmpty(col int) bool {
	rows := len(u.data)
	for row := 0; row < rows; row++ {
		if u.Get(col, row).cellType != CELL_EMPTY {
			return false
		}

	}
	return true
}

func (u *Universe) AddRow(indexToInsert int) {
	rows := u.data
	cols := len(u.data[0])
	new_row := make([]*Cell, 0)
	for col := 0; col < cols; col++ {
		new_row = append(new_row, NewCell("."))
	}
	new_rows := make([][]*Cell, 0)
	for index, row := range rows {
		if index == indexToInsert {
			new_rows = append(new_rows, new_row)
		}
		new_rows = append(new_rows, row)
	}
	u.data = new_rows
}

func (u *Universe) AddColumn(index int) {
	rows := len(u.data)
	cols := len(u.data[0])

	data := make([][]*Cell, 0)
	for rowNum := 0; rowNum < rows; rowNum++ {
		new_row := make([]*Cell, 0)
		cell := NewCell(".")
		for colNum := 0; colNum < cols; colNum++ {
			new_row = append(new_row, u.Get(colNum, rowNum))
			if colNum == index {
				new_row = append(new_row, cell)
			}
		}
		data = append(data, new_row)
	}
	u.data = data
}

func (u *Universe) Pairs() [][]*Cell {
	// find galaxies
	cells := make([]*Cell, 0)
	for rowNum, row := range u.data {
		for colNum, _ := range row {
			cell := u.Get(colNum, rowNum)
			if cell.cellType == CELL_GALAXY {
				cell.x = colNum
				cell.y = rowNum
				cells = append(cells, cell)
			}
		}
	}

	cache := make(map[string][]*Cell)
	for index1 := 0; index1 < len(cells); index1++ {
		for index2 := 0; index2 < len(cells); index2++ {
			if index1 == index2 {
				continue
			}
			key := ""
			if index1 < index2 {
				key = fmt.Sprintf("%v.%v.", index1, index2)
			} else {
				key = fmt.Sprintf("%v.%v.", index2, index1)
			}
			_, exists := cache[key]
			if exists {
				continue
			}

			cell1 := cells[index1]
			cell2 := cells[index2]
			pair := make([]*Cell, 0)
			pair = append(pair, cell1)
			pair = append(pair, cell2)
			cache[key] = pair

		}
	}

	pairs := make([][]*Cell, 0)
	for _, v := range cache {
		pairs = append(pairs, v)
	}
	return pairs

}

type Cell struct {
	cellType string
	x        int
	y        int
}

func NewCell(cellType string) *Cell {
	return &Cell{cellType: cellType}
}

func (c *Cell) Distance(o *Cell) int {
	return int(math.Abs(float64(c.x-o.x)) + math.Abs(float64(c.y-o.y)))
}

func (c *Cell) DistanceP2(u *Universe, o *Cell, sizeEmpty int) int {

	min_x := utils.MinInt(c.x, o.x)
	max_x := utils.MaxInt(c.x, o.x)
	min_y := utils.MinInt(c.y, o.y)
	max_y := utils.MaxInt(c.y, o.y)

	emptyCols := 0
	emptyRows := 0
	for x := min_x; x <= max_x; x++ {
		if u.IsColEmpty(x) {
			emptyCols++
		}
	}
	for y := min_y; y <= max_y; y++ {
		if u.IsRowEmpty(y) {
			emptyRows++
		}
	}

	max_x = max_x - (emptyCols) + (emptyCols * sizeEmpty)
	max_y = max_y - (emptyRows) + (emptyRows * sizeEmpty)
	xdistance := max_x - min_x
	ydistance := max_y - min_y
	return xdistance + ydistance

	// x1 := c.x
	// x2 := o.x
	// y1 := c.y
	// y2 := o.y

	// million := 1000000
	// xdiff := emptyCols * million
	// ydiff := emptyRows * million

	// return int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
}
