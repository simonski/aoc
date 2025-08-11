package d15

import (
	"fmt"
	"strings"
)

/*
Day 15: Warehouse Woes
*/

func ParseP2(data1, data2 string) (*Grid2, string) {
	g := NewGrid2(data1)
	splits := strings.Split(data2, "\n")
	line := ""
	for _, l := range splits {
		line += l
	}
	return g, line

}

type Grid2 struct {
	data   map[string]*Cell2
	width  int
	height int
	robot  *Cell2
}

func NewGrid2(data string) *Grid2 {
	gdata := make(map[string]*Cell2)
	rows := strings.Split(data, "\n")
	var robot *Cell2
	for y, row := range rows {
		newX := 0
		for x := 0; x < len(row); x++ {
			newX = x * 2
			cellType := row[x : x+1]
			if cellType == "#" {
				cell1 := &Cell2{x: newX, y: y, cellType: cellType}
				cell2 := &Cell2{x: newX + 1, y: y, cellType: cellType}
				gdata[cell1.key()] = cell1
				gdata[cell2.key()] = cell2
			} else if cellType == "." {
				cell1 := &Cell2{x: newX, y: y, cellType: cellType}
				cell2 := &Cell2{x: newX + 1, y: y, cellType: cellType}
				gdata[cell1.key()] = cell1
				gdata[cell2.key()] = cell2
			} else if cellType == "O" {
				cell1 := &Cell2{x: newX, y: y, cellType: "["}
				cell2 := &Cell2{x: newX + 1, y: y, cellType: "]"}
				gdata[cell1.key()] = cell1
				gdata[cell2.key()] = cell2
			} else if cellType == "@" {
				cell1 := &Cell2{x: newX, y: y, cellType: "@"}
				cell2 := &Cell2{x: newX + 1, y: y, cellType: "."}
				gdata[cell1.key()] = cell1
				gdata[cell2.key()] = cell2
				robot = cell1
			}

		}
	}

	height := len(rows)
	width := len(rows[0]) * 2
	g := Grid2{data: gdata, width: width, height: height, robot: robot}
	return &g
}

func (g *Grid2) debug() string {
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

func (g *Grid2) get(x int, y int) *Cell2 {
	key := fmt.Sprintf("%v.%v", x, y)
	cell := g.data[key]
	if cell == nil {
		return &Cell2{x: x, y: y, cellType: "."}
	} else {
		return cell
	}
}

func (g *Grid2) up(c *Cell2) {
	delete(g.data, c.key())
	c.y -= 1
	g.data[c.key()] = c
}

func (g *Grid2) down(c *Cell2) {
	delete(g.data, c.key())
	c.y += 1
	g.data[c.key()] = c
}

func (g *Grid2) left(c *Cell2) {
	delete(g.data, c.key())
	c.x -= 1
	g.data[c.key()] = c
}

func (g *Grid2) right(c *Cell2) {
	delete(g.data, c.key())
	c.x += 1
	g.data[c.key()] = c
}

type Cell2 struct {
	x        int
	y        int
	cellType string
}

func (c *Cell2) key() string {
	return fmt.Sprintf("%v.%v", c.x, c.y)
}

func (c *Cell2) isEmpty() bool {
	return c.cellType == "."
}

func (c *Cell2) isWall() bool {
	return c.cellType == "#"
}

func (c *Cell2) isRobot() bool {
	return c.cellType == "@"
}

func (c *Cell2) isBox() bool {
	return c.cellType == "[" || c.cellType == "]"
}

func (c *Cell2) isBoxL() bool {
	return c.cellType == "["
}

func (c *Cell2) isBoxR() bool {
	return c.cellType == "]"
}

func (g *Grid2) execute(instruction string) {
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

func (g *Grid2) moveLeft(robot *Cell2) {
	cells := make([]*Cell2, 0)
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

func (g *Grid2) moveRight(robot *Cell2) {
	cells := make([]*Cell2, 0)
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

func (g *Grid2) moveUp(robot *Cell2) {
	cells := make([]*Cell2, 0)
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

func (g *Grid2) moveDown(robot *Cell2) {
	cells := make([]*Cell2, 0)
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

func (g *Grid2) gps() int {
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
