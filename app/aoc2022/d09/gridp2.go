package d09

import (
	"fmt"
	"strconv"
	"strings"
)

/*
--- Day 05:  ---

*/

type BoardP2 struct {
	Knots        []*Grid
	Instructions []string
	VERBOSE      bool
	rows         int
	cols         int
}

func NewBoardP2(input string, verbose bool, size int, rows int, cols int, start_row int, start_col int) *BoardP2 {
	b := BoardP2{}
	b.Knots = make([]*Grid, size)
	for index := 0; index < size; index++ {
		b.Knots[index] = NewGrid(start_row, start_col)
	}
	b.VERBOSE = verbose
	b.Instructions = strings.Split(input, "\n")
	b.rows = rows
	b.cols = cols
	return &b
}

func (b *BoardP2) Debug(rows int, cols int) string {
	result := ""
	for row := rows; row >= 0; row-- {
		line := ""
		for col := 0; col < cols; col++ {
			found := false
			for index := 0; index < len(b.Knots); index++ {
				knot := b.Knots[index]
				if knot.current_col == col && knot.current_row == row {
					if index == 0 {
						line = fmt.Sprintf("%v%v", line, "H")
					} else {
						line = fmt.Sprintf("%v%v", line, index)
					}
					found = true
					break
				}
			}
			if !found {
				line = fmt.Sprintf("%v.", line)
			}
		}
		result = fmt.Sprintf("%v\n%v", result, line)
	}
	return result
}

func (b *BoardP2) RunInstructions(debug bool) {
	for _, instruction := range b.Instructions {
		b.RunInstruction(instruction, debug)
	}
}

func (b *BoardP2) RunInstruction(instruction string, debug bool) {
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
			fmt.Println(b.Debug(b.rows, b.cols))
			fmt.Println("")
		}
	}
	if debug {
		fmt.Println("--------------------")
	}
}

func (b *BoardP2) DistanceV(index int) int {
	head := b.Knots[index]
	tail := b.Knots[index+1]
	return head.current_row - tail.current_row
}

func (b *BoardP2) DistanceH(index int) int {
	head := b.Knots[index]
	tail := b.Knots[index+1]
	return head.current_col - tail.current_col
}

func (b *BoardP2) Up() {
	if b.VERBOSE {
		fmt.Println("b.Up()")
	}
	b.Knots[0].Up()
	b.MoveTail()
}
func (b *BoardP2) Down() {
	if b.VERBOSE {
		fmt.Println("b.Down()")
	}
	b.Knots[0].Down()
	b.MoveTail()
}
func (b *BoardP2) Left() {
	if b.VERBOSE {
		fmt.Println("b.Left()")
	}
	b.Knots[0].Left()
	b.MoveTail()
}
func (b *BoardP2) Right() {
	if b.VERBOSE {
		fmt.Println("b.Right()")
	}
	b.Knots[0].Right()
	b.MoveTail()
}

func (b *BoardP2) CountTailVisits() int {
	lastIndex := len(b.Knots) - 1
	tail := b.Knots[lastIndex]
	return len(tail.data)
}

func (b *BoardP2) MoveTail() {
	for i := 0; i < len(b.Knots)-1; i++ {
		head := b.Knots[i]
		tail := b.Knots[i+1]

		if tail.current_col == head.current_col && tail.current_row == head.current_row {
			if b.VERBOSE {
				fmt.Println("MoveTail() no need.")
			}
			return
		}
		distance := b.DistanceH(i) * b.DistanceV(i)
		if distance < 0 {
			// abs on ints, can't be bothered casting around
			distance *= -1
		}

		right := false
		up := false
		left := false
		down := false

		distanceH := b.DistanceH(i)
		distanceV := b.DistanceV(i)

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
			tail.UpLeft()
		} else if left && down {
			tail.DownLeft()
		} else if left {
			tail.Left()
		} else if right && up {
			tail.UpRight()
		} else if right && down {
			tail.DownRight()
		} else if right {
			tail.Right()
		} else if up {
			tail.Up()
		} else if down {
			tail.Down()
		}
	}
}
