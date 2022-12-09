package d09

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/goutils"
)

/*
--- Day 05:  ---

*/

type Grid struct {
	// the posiiton is the key (row_col) the value is the number of times visited
	data        map[string]int
	current_row int
	current_col int
	max_row     int
	max_col     int
}

func NewGrid() *Grid {
	data := make(map[string]int)
	g := Grid{data: data}
	g.Increment(0, 0)
	return &g
}

func (g *Grid) Increment(row int, col int) {
	key := fmt.Sprintf("%v_%v", row, col)
	value := g.data[key]
	value += 1
	g.data[key] = value
}

func (g *Grid) Get(row int, col int) int {
	key := fmt.Sprintf("%v_%v", row, col)
	return g.data[key]
}

func (g *Grid) Up() {
	g.current_row += 1
	g.Increment(g.current_row, g.current_col)
	g.max_row = goutils.Max(g.current_row, g.max_row)
}

func (g *Grid) UpRight() {
	g.current_row += 1
	g.current_col += 1
	g.Increment(g.current_row, g.current_col)
	g.max_row = goutils.Max(g.current_row, g.max_row)
}

func (g *Grid) UpLeft() {
	g.current_row += 1
	g.current_col -= 1
	g.Increment(g.current_row, g.current_col)
	g.max_row = goutils.Max(g.current_row, g.max_row)
}

func (g *Grid) Down() {
	g.current_row -= 1
	g.Increment(g.current_row, g.current_col)
	g.max_row = goutils.Max(g.current_row, g.max_row)
}

func (g *Grid) DownLeft() {
	g.current_row -= 1
	g.current_col -= 1
	g.Increment(g.current_row, g.current_col)
	g.max_row = goutils.Max(g.current_row, g.max_row)
}

func (g *Grid) DownRight() {
	g.current_row -= 1
	g.current_col += 1
	g.Increment(g.current_row, g.current_col)
	g.max_row = goutils.Max(g.current_row, g.max_row)
}
func (g *Grid) Left() {
	g.current_col -= 1
	g.Increment(g.current_row, g.current_col)
	g.max_col = goutils.Max(g.current_col, g.max_col)
}
func (g *Grid) Right() {
	g.current_col += 1
	g.Increment(g.current_row, g.current_col)
	g.max_col = goutils.Max(g.current_col, g.max_col)
}

type Board struct {
	Head         *Grid
	Tail         *Grid
	Instructions []string
	VERBOSE      bool
}

func NewBoard(input string, verbose bool) *Board {
	b := Board{}
	b.Head = NewGrid()
	b.Tail = NewGrid()
	b.VERBOSE = verbose
	b.Instructions = strings.Split(input, "\n")
	return &b
}

func (b *Board) Debug(rows int, cols int) string {
	result := ""
	for row := rows; row >= 0; row-- {
		line := ""
		for col := 0; col < cols; col++ {
			if b.Head.current_col == col && b.Head.current_row == row {
				line = fmt.Sprintf("%vH", line)
			} else if b.Tail.current_col == col && b.Tail.current_row == row {
				line = fmt.Sprintf("%vT", line)
			} else {
				line = fmt.Sprintf("%v.", line)
			}
		}
		result = fmt.Sprintf("%v\n%v", result, line)
	}
	return result
}

func (b *Board) RunInstructions(debug bool) {
	for _, instruction := range b.Instructions {
		b.RunInstruction(instruction, debug)
	}
}

func (b *Board) RunInstruction(instruction string, debug bool) {
	splits := strings.Split(instruction, " ")
	move := splits[0]
	value, _ := strconv.Atoi(splits[1])
	if b.VERBOSE {
		fmt.Printf("RunInstruction(%v) move=%v, value=%v\n", instruction, move, value)
	}
	if debug {
		fmt.Printf("%v --------------------\n", instruction)
	}
	for index := 0; index < value; index++ {

		if debug {
			fmt.Printf("%v %v/%v-----------------\n", instruction, index+1, value)
		}

		if move == "U" {
			b.Up()
		} else if move == "D" {
			b.Down()
		} else if move == "L" {
			b.Left()
		} else if move == "R" {
			b.Right()
		}

		if debug {
			fmt.Println(b.Debug(6, 6))
			fmt.Println("")
		}
	}
	if debug {
		fmt.Println("--------------------")
	}
}

func (b *Board) DistanceV() int {
	return b.Head.current_row - b.Tail.current_row
}

func (b *Board) DistanceH() int {
	return b.Head.current_col - b.Tail.current_col
}

func (b *Board) Up() {
	if b.VERBOSE {
		fmt.Println("b.Up()")
	}
	b.Head.Up()
	b.MoveTail()
}
func (b *Board) Down() {
	if b.VERBOSE {
		fmt.Println("b.Down()")
	}
	b.Head.Down()
	b.MoveTail()
}
func (b *Board) Left() {
	if b.VERBOSE {
		fmt.Println("b.Left()")
	}
	b.Head.Left()
	b.MoveTail()
}
func (b *Board) Right() {
	if b.VERBOSE {
		fmt.Println("b.Right()")
	}
	b.Head.Right()
	b.MoveTail()
}

func (b *Board) CountTailVisits() int {
	return len(b.Tail.data)
}

func (b *Board) MoveTail() {
	if b.Tail.current_col == b.Head.current_col && b.Tail.current_row == b.Head.current_row {
		if b.VERBOSE {
			fmt.Println("MoveTail() no need.")
		}
		return
	}
	distance := b.DistanceH() * b.DistanceV()
	if distance < 0 {
		// abs on ints, can't be bothered casting around
		distance *= -1
	}

	right := false
	up := false
	left := false
	down := false

	distanceH := b.DistanceH()
	distanceV := b.DistanceV()

	if distanceV > 1 {
		up = true
		if distanceH > 0 {
			right = true
		} else if distanceH < 0 {
			left = true
		}
	} else if distanceV < -1 {
		down = true
		if distanceH > 0 {
			right = true
		} else if distanceH < 0 {
			left = true
		}
	}

	if distanceH > 1 {
		right = true
		if distanceV > 0 {
			up = true
		} else if distanceV < 0 {
			down = true
		}
	} else if distanceH < -1 {
		left = true
		if distanceV > 0 {
			up = true
		} else if distanceV < 0 {
			down = true
		}
	}

	if left && up {
		b.Tail.UpLeft()
	} else if left && down {
		b.Tail.DownLeft()
	} else if left {
		b.Tail.Left()
	} else if right && up {
		b.Tail.UpRight()
	} else if right && down {
		b.Tail.DownRight()
	} else if right {
		b.Tail.Right()
	} else if up {
		b.Tail.Up()
	} else if down {
		b.Tail.Down()
	}
}
