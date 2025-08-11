package d8

import (
	"fmt"
	"strconv"
	"strings"
)

/*
--- Day 05:  ---

*/

type Grid struct {
	data map[string]int
	rows int
	cols int
}

func NewGrid(input string) *Grid {
	rows := strings.Split(input, "\n")
	size := len(rows)
	data := make(map[string]int)
	g := Grid{data: data, rows: size, cols: size}
	for rownum := 0; rownum < len(rows); rownum++ {
		row := rows[rownum]
		for colnum := 0; colnum < len(row); colnum++ {
			tree := row[colnum : colnum+1]
			treeHeight, _ := strconv.Atoi(tree)
			g.Put(rownum, colnum, treeHeight)
		}
	}
	return &g

}

func (g *Grid) Put(row int, col int, height int) {
	key := fmt.Sprintf("%v_%v", row, col)
	g.data[key] = height
}

func (g *Grid) Get(row int, col int) int {
	key := fmt.Sprintf("%v_%v", row, col)
	return g.data[key]
}

func (g *Grid) IsVisible(treeRow int, treeCol int) (bool, bool, bool, bool) {
	// if all trees between it and edge of grid are shorter (it is higher)
	top := g.IsVisibleFromTop(treeRow, treeCol)
	bottom := g.IsVisibleFromBottom(treeRow, treeCol)
	left := g.IsVisibleFromLeft(treeRow, treeCol)
	right := g.IsVisibleFromRight(treeRow, treeCol)
	return top, bottom, left, right

}

func (g *Grid) IsVisibleFromTop(treeRow int, treeCol int) bool {
	treeHeight := g.Get(treeRow, treeCol)
	if treeRow > 0 {
		// walk "up" to 0
		for r := treeRow - 1; r >= 0; r-- {
			candidate := g.Get(r, treeCol)
			if candidate >= treeHeight {
				return false
			}
		}
	}
	return true
}

func (g *Grid) IsVisibleFromBottom(treeRow int, treeCol int) bool {
	treeHeight := g.Get(treeRow, treeCol)
	// check "down" the rows - this checks walks down each row verifying the column
	if treeRow < g.rows {
		for r := treeRow + 1; r < g.rows; r++ {
			candidate := g.Get(r, treeCol)
			if candidate >= treeHeight {
				return false
			}
		}
	}
	return true
}

func (g *Grid) IsVisibleFromRight(treeRow int, treeCol int) bool {
	treeHeight := g.Get(treeRow, treeCol)
	if treeCol < g.cols {
		for c := treeCol + 1; c < g.cols; c++ {
			candidate := g.Get(treeRow, c)
			if candidate >= treeHeight {
				return false
			}
		}
	}
	return true
}

func (g *Grid) IsVisibleFromLeft(treeRow int, treeCol int) bool {
	treeHeight := g.Get(treeRow, treeCol)
	if treeCol > 0 {
		// walk "up" to 0
		for c := treeCol - 1; c >= 0; c-- {
			candidate := g.Get(treeRow, c)
			if candidate >= treeHeight {
				return false
			}
		}
	}
	return true
}

func (g *Grid) CountVisibleBottom(treeRow int, treeCol int) int {
	treeHeight := g.Get(treeRow, treeCol)
	count := 0
	for row := treeRow + 1; row < g.rows; row++ {
		height := g.Get(row, treeCol)
		if height < treeHeight {
			fmt.Printf("bottom: my height (%v,%v) %v, this height (%v,%v) %v, adding.\n", treeRow, treeCol, treeHeight, row, treeCol, height)
			count++
		} else {
			fmt.Printf("bottom: my height (%v,%v) %v, this height (%v,%v) %v, adding and quitting.\n", treeRow, treeCol, treeHeight, row, treeCol, height)
			count++
			break
		}
	}
	return count
}

func (g *Grid) CountVisibleTop(treeRow int, treeCol int) int {
	treeHeight := g.Get(treeRow, treeCol)
	count := 0
	for row := treeRow - 1; row >= 0; row-- {
		height := g.Get(row, treeCol)
		if height < treeHeight {
			fmt.Printf("top: my height (%v,%v) %v, this height (%v,%v) %v, adding.\n", treeRow, treeCol, treeHeight, row, treeCol, height)
			count++
		} else {
			fmt.Printf("top: my height (%v,%v) %v, this height (%v,%v) %v, adding and stopping.\n", treeRow, treeCol, treeHeight, row, treeCol, height)
			count++
			break
		}
	}
	return count
}

func (g *Grid) CountVisibleLeft(treeRow int, treeCol int) int {
	treeHeight := g.Get(treeRow, treeCol)
	count := 0
	for col := treeCol - 1; col >= 0; col-- {
		height := g.Get(treeRow, col)
		if height < treeHeight {
			fmt.Printf("left: my height (%v,%v) %v, this height (%v,%v) %v, adding.\n", treeRow, treeCol, treeHeight, treeRow, col, height)
			count++
		} else {
			count++
			fmt.Printf("left: my height (%v,%v) %v, this height (%v,%v) %v, adding and quitting.\n", treeRow, treeCol, treeHeight, treeRow, col, height)
			break
		}
	}
	return count
}

func (g *Grid) CountVisibleRight(treeRow int, treeCol int) int {
	treeHeight := g.Get(treeRow, treeCol)
	count := 0
	for col := treeCol + 1; col < g.cols; col++ {
		height := g.Get(treeRow, col)
		if height < treeHeight {
			fmt.Printf("right: my height (%v,%v) %v, this height (%v,%v) %v, adding.\n", treeRow, treeCol, treeHeight, treeRow, col, height)
			count++
		} else {
			fmt.Printf("right: my height (%v,%v) %v, this height (%v,%v) %v, adding and quitting.\n", treeRow, treeCol, treeHeight, treeRow, col, height)
			count++
			break
		}
	}
	return count
}
