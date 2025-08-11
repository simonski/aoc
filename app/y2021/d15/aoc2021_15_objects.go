package aoc2021

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"

// 	utils "github.com/simonski/aoc/utils"
// )

// type Point_D15 struct {
// 	Value int
// 	X     int
// 	Y     int
// 	Grid  *Grid_D15
// }

// func (p *Point_D15) IsEnd() bool {
// 	return p.X == p.Grid.Width-1 && p.Y == p.Grid.Height-1
// }

// func (p *Point_D15) IsStart() bool {
// 	return p.X == 0 && p.Y == 0
// }

// type Grid_D15 struct {
// 	Points map[string]*Point_D15
// 	Width  int
// 	Height int
// }

// func (g *Grid_D15) GetPoint(x int, y int) *Point_D15 {
// 	key := fmt.Sprintf("%v.%v", x, y)
// 	return g.Points[key]
// }

// func (g *Grid_D15) PathCost(p *Point_D15, path []int) int {
// 	total := 0
// 	for _, instruction := range path {
// 		if instruction == 0 {
// 			p = g.GetPoint(p.X, p.Y+1)
// 			total += p.Value
// 		} else {
// 			p = g.GetPoint(p.X+1, p.Y)
// 			total += p.Value
// 		}
// 	}
// 	return total

// }

// func (g *Grid_D15) PathCostStr(p *Point_D15, path string) int {
// 	total := 0
// 	for index := 0; index < len(path); index++ {
// 		instruction := path[index : index+1]

// 		// path = strings.Replace(path, "a", "0", 1)
// 		// path = strings.Replace(path, "b", "0", 1)
// 		// path = strings.Replace(path, "c", "0", 1)
// 		// path = strings.Replace(path, "d", "0", 1)
// 		// path = strings.Replace(path, "e", "0", 1)
// 		// path = strings.Replace(path, "f", "0", 1)
// 		// path = strings.Replace(path, "g", "0", 1)
// 		// path = strings.Replace(path, "h", "0", 1)
// 		// path = strings.Replace(path, "i", "0", 1)

// 		// path = strings.Replace(path, "J", "1", 1)
// 		// path = strings.Replace(path, "K", "1", 1)
// 		// path = strings.Replace(path, "L", "1", 1)
// 		// path = strings.Replace(path, "M", "1", 1)
// 		// path = strings.Replace(path, "N", "1", 1)
// 		// path = strings.Replace(path, "O", "1", 1)
// 		// path = strings.Replace(path, "P", "1", 1)
// 		// path = strings.Replace(path, "Q", "1", 1)
// 		// path = strings.Replace(path, "R", "1", 1)

// 		if instruction == "a" || instruction == "b" || instruction == "c" || instruction == "d" || instruction == "e" || instruction == "f" || instruction == "g" || instruction == "h" || instruction == "i" {
// 			p = g.GetPoint(p.X, p.Y+1)
// 			total += p.Value
// 		} else {
// 			p = g.GetPoint(p.X+1, p.Y)
// 			total += p.Value
// 		}
// 	}
// 	return total

// }

// func (g *Grid_D15) GetDistanceCost(op *Point_D15) (map[int]int, map[int]int, int) {
// 	xtotal := 0
// 	ytotal := 0
// 	xcosts := make(map[int]int)
// 	ycosts := make(map[int]int)
// 	for x := op.X + 1; x < g.Width; x++ {
// 		p := g.GetPoint(x, op.Y)
// 		xtotal += p.Value
// 		distance := p.X - op.X
// 		xcosts[distance] = xtotal
// 	}
// 	for y := op.Y + 1; y < g.Height; y++ {
// 		p := g.GetPoint(op.X, y)
// 		ytotal += p.Value
// 		distance := p.Y - op.Y
// 		ycosts[distance] = ytotal
// 	}

// 	l := 0
// 	l = utils.Max(l, len(xcosts))
// 	l = utils.Max(l, len(ycosts))
// 	return xcosts, ycosts, l
// }

// func NewGrid_D15(data string) *Grid_D15 {
// 	rows := strings.Split(data, "\n")
// 	Height := len(rows)
// 	Width := 0
// 	Points := make(map[string]*Point_D15)
// 	g := &Grid_D15{}
// 	for row := 0; row < len(rows); row++ {
// 		line := rows[row]
// 		Width = len(line)
// 		for col := 0; col < len(line); col++ {
// 			Value, _ := strconv.Atoi(line[col : col+1])
// 			p := &Point_D15{X: col, Y: row, Value: Value, Grid: g}
// 			key := fmt.Sprintf("%v.%v", col, row)
// 			Points[key] = p
// 		}
// 	}
// 	g.Points = Points
// 	g.Width = Width
// 	g.Height = Height
// 	return g
// }

// func (g *Grid_D15) RowTotal(row int) int {
// 	total := 0
// 	for col := 0; col < g.Width; col++ {
// 		p := g.GetPoint(col, row)
// 		total += p.Value
// 	}
// 	return total
// }

// func (g *Grid_D15) ColTotal(col int) int {
// 	total := 0
// 	for row := 0; row < g.Width; row++ {
// 		p := g.GetPoint(col, row)
// 		total += p.Value
// 	}
// 	return total
// }

// func (g *Grid_D15) Debug() string {
// 	line := ""
// 	for row := 0; row < g.Height; row++ {
// 		for col := 0; col < g.Width; col++ {
// 			p := g.GetPoint(col, row)
// 			line = fmt.Sprintf("%v %v", line, p.Value)
// 		}
// 		line = fmt.Sprintf("%v (%v)\n", line, g.RowTotal(row))

// 	}
// 	for col := 0; col < g.Width; col++ {
// 		line = fmt.Sprintf("%v (%v)", line, g.ColTotal(col))
// 	}
// 	line += "\n"
// 	return line
// }

// func (g *Grid_D15) Cost(x1 int, y1 int, x2 int, y2 int) string {
// 	line := ""
// 	for row := 0; row < g.Height; row++ {
// 		for col := 0; col < g.Width; col++ {
// 			p := g.GetPoint(col, row)
// 			line = fmt.Sprintf("%v %v", line, p.Value)
// 		}
// 		line = fmt.Sprintf("%v (%v)\n", line, g.RowTotal(row))

// 	}
// 	for col := 0; col < g.Width; col++ {
// 		line = fmt.Sprintf("%v (%v)", line, g.ColTotal(col))
// 	}
// 	line += "\n"
// 	return line
// }

// func (g *Grid_D15) FrequencyDistribution() map[int]int {
// 	results := make(map[int]int)
// 	for _, p := range g.Points {
// 		key := p.Value
// 		results[key] += 1
// 	}
// 	return results
// }
