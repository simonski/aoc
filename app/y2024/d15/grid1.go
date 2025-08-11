package d15

import (
	"fmt"
	"strings"
)

/*
Day 15: Warehouse Woes
*/

func Parse(data1, data2 string) (*Grid, string) {
	g := NewGrid(data1)
	splits := strings.Split(data2, "\n")
	line := ""
	for _, l := range splits {
		line += l
	}
	return g, line

}

type Grid struct {
	data   map[string]*Cell
	width  int
	height int
	robot  *Cell
}

func NewGrid(data string) *Grid {
	gdata := make(map[string]*Cell)
	rows := strings.Split(data, "\n")
	var robot *Cell
	for y, row := range rows {
		for x := 0; x < len(row); x++ {
			cellType := row[x : x+1]
			cell := &Cell{x: x, y: y, cellType: cellType}
			if !cell.isEmpty() {
				gdata[cell.key()] = cell
			}
			if cell.isRobot() {
				robot = cell
			}
		}
	}

	height := len(rows)
	width := len(rows[0])
	g := Grid{data: gdata, width: width, height: height, robot: robot}
	return &g
}

func (g *Grid) debug() string {
	line := ""
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			c := g.get(x, y)
			if c != nil {
				line += c.cellType
			} else {
				line += "."
			}
		}
		line += "\n"
	}
	return line
}

func (g *Grid) get(x int, y int) *Cell {
	key := fmt.Sprintf("%v.%v", x, y)
	cell := g.data[key]
	if cell == nil {
		return &Cell{x: x, y: y, cellType: "."}
	} else {
		return cell
	}
}

func (g *Grid) up(c *Cell) {
	delete(g.data, c.key())
	c.y -= 1
	g.data[c.key()] = c
}

func (g *Grid) down(c *Cell) {
	delete(g.data, c.key())
	c.y += 1
	g.data[c.key()] = c
}

func (g *Grid) left(c *Cell) {
	delete(g.data, c.key())
	c.x -= 1
	g.data[c.key()] = c
}

func (g *Grid) right(c *Cell) {
	delete(g.data, c.key())
	c.x += 1
	g.data[c.key()] = c
}

type Cell struct {
	x        int
	y        int
	cellType string
}

func (c *Cell) key() string {
	return fmt.Sprintf("%v.%v", c.x, c.y)
}

func (c *Cell) isEmpty() bool {
	return c.cellType == "."
}

func (c *Cell) isWall() bool {
	return c.cellType == "#"
}

func (c *Cell) isRobot() bool {
	return c.cellType == "@"
}

func (c *Cell) isBox() bool {
	return c.cellType == "O"
}

func (g *Grid) execute(instruction string) {
	if instruction == "^" {
		g.moveUp(g.robot)
	} else if instruction == "v" {
		g.moveDown(g.robot)
	} else if instruction == "<" {
		g.moveLeft(g.robot)
	} else if instruction == ">" {
		g.moveRight(g.robot)
	}
}

func (g *Grid) moveLeft(robot *Cell) {
	cells := make([]*Cell, 0)
	canPush := false
	cells = append(cells, robot)
	for x := robot.x - 1; x >= 0; x-- {
		cell := g.get(x, robot.y)
		if cell.isBox() {
			// it can be pushed
			cells = append(cells, g.get(x, robot.y))
		} else if cell.isEmpty() {
			// this is the end
			canPush = true
			break
		} else if cell.isWall() {
			break
		}
	}

	if canPush {
		for index := len(cells) - 1; index >= 0; index-- {
			cell := cells[index]
			g.left(cell)
		}
	}
}

func (g *Grid) moveRight(robot *Cell) {
	cells := make([]*Cell, 0)
	canPush := false
	cells = append(cells, robot)
	for x := robot.x + 1; x < g.width; x++ {
		cell := g.get(x, robot.y)
		if cell.isBox() {
			// it can be pushed
			cells = append(cells, g.get(x, robot.y))
		} else if cell.isEmpty() {
			// this is the end
			canPush = true
			break
		} else if cell.isWall() {
			break
		}
	}

	if canPush {
		for index := len(cells) - 1; index >= 0; index-- {
			cell := cells[index]
			g.right(cell)
		}
	}
}

func (g *Grid) moveUp(robot *Cell) {
	cells := make([]*Cell, 0)
	canPush := false
	cells = append(cells, robot)
	for y := robot.y - 1; y >= 0; y-- {
		cell := g.get(robot.x, y)
		if cell.isBox() {
			// it can be pushed
			cells = append(cells, g.get(robot.x, y))
		} else if cell.isEmpty() {
			// this is the end
			canPush = true
			break
		} else if cell.isWall() {
			break
		}
	}

	if canPush {
		for index := len(cells) - 1; index >= 0; index-- {
			cell := cells[index]
			g.up(cell)
		}
	}
}

func (g *Grid) moveDown(robot *Cell) {
	cells := make([]*Cell, 0)
	canPush := false
	cells = append(cells, robot)
	for y := robot.y + 1; y < g.height; y++ {
		cell := g.get(robot.x, y)
		if cell.isBox() {
			// it can be pushed
			cells = append(cells, g.get(robot.x, y))
		} else if cell.isEmpty() {
			// this is the end
			canPush = true
			break
		} else if cell.isWall() {
			break
		}
	}

	if canPush {
		for index := len(cells) - 1; index >= 0; index-- {
			cell := cells[index]
			g.down(cell)
		}
	}
}

func (g *Grid) gps() int {
	total := 0
	for _, cell := range g.data {
		if cell.isBox() {
			x := cell.x
			y := cell.y
			value := 100*y + x
			total += value
		}
	}
	return total
}
