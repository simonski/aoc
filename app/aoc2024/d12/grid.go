package d12

import (
	"fmt"
	"strings"
)

/*
Day 12: Garden Groups
*/

type Grid struct {
	data    map[string]*Cell
	width   int
	height  int
	region  int
	regions []*Region
}

func NewGrid(data string) *Grid {
	grid := Grid{data: make(map[string]*Cell)}
	rows := strings.Split(data, "\n")
	for row_index := 0; row_index < len(rows); row_index++ {
		row := rows[row_index]
		for col_index := 0; col_index < len(row); col_index++ {
			value := row[col_index : col_index+1]
			cell := NewCell(col_index, row_index, value)
			grid.add(cell)
		}
	}
	grid.width = len(rows)
	grid.height = len(rows[0])
	grid.regions = make([]*Region, 0)
	grid.findRegions()
	return &grid
}

func (g *Grid) add(c *Cell) {
	g.data[c.key] = c
}

func (g *Grid) get(x int, y int) *Cell {
	if x < 0 || x >= g.width || y < 0 || y >= g.height {
		return nil
	}
	key := fmt.Sprintf("%v.%v", x, y)
	cell := g.data[key]
	return cell
}

func (g *Grid) area(r *Region) int {
	return 0
}

func (g *Grid) perimeter(r *Region) int {
	return 0
}

func (g *Grid) neighboursSameType(c *Cell) []*Cell {
	results := make([]*Cell, 0)
	up := g.get(c.x, c.y-1)
	down := g.get(c.x, c.y+1)
	left := g.get(c.x-1, c.y)
	right := g.get(c.x+1, c.y)
	if up != nil && up.value == c.value {
		results = append(results, up)
	}
	if down != nil && down.value == c.value {
		results = append(results, down)
	}
	if left != nil && left.value == c.value {
		results = append(results, left)
	}
	if right != nil && right.value == c.value {
		results = append(results, right)
	}
	return results
}

func (grid *Grid) addRegion(r *Region) {
	grid.regions = append(grid.regions, r)
}

func (grid *Grid) findRegions() {
	for y := 0; y < grid.height; y++ {
		for x := 0; x < grid.width; x++ {
			cell := grid.get(x, y)
			if cell.region == -1 {
				grid.region += 1
				cell.region = grid.region
				region := NewRegion(grid.region)
				region.add(cell)
				grid.addRegion(region)
				// fmt.Printf("region %v (%v, %v)\n", grid.region, cell.x, cell.y)
				grid.floodFindRegions(region, cell)
			}
		}
	}

}

func (grid *Grid) floodFindRegions(region *Region, cell *Cell) {
	// any neighbours get walked and region assigned
	neighbours := grid.neighboursSameType(cell)
	for _, n := range neighbours {
		if n.region == -1 {
			n.region = cell.region
			region.add(n)
			grid.floodFindRegions(region, n)
		}
	}
}

func (grid *Grid) countCorners(cell *Cell) int {

	//  00000XXX000000
	//  00XXXXXXXXX000
	//  0XXXXXXXXXXXX0
	//  00000X00XXXXXX

	// a corner cell tells me the edge
	isNotNeighbour := func(candidate *Cell, cell *Cell) bool {
		if candidate == nil {
			return true
		}
		return candidate.region != cell.region
	}

	isNeighbour := func(candidate *Cell, cell *Cell) bool {
		return candidate != nil && candidate.region == cell.region
	}

	left := grid.left(cell)
	up := grid.up(cell)
	right := grid.right(cell)
	down := grid.down(cell)

	upRight := grid.upRight(cell)
	downLeft := grid.downLeft(cell)
	upLeft := grid.upLeft(cell)

	count := 0
	// _. outside top left   (bottom-right is origin)
	// .X
	if isNotNeighbour(up, cell) && isNotNeighbour(left, cell) {
		count++
	}

	// _. outside top right   (bottom-right is origin)
	// X.
	if isNotNeighbour(up, cell) && isNotNeighbour(right, cell) {
		count++
	}

	// .X bottom left
	// _.
	if isNotNeighbour(left, cell) && isNotNeighbour(down, cell) {
		count++
	}

	// X. bottom right
	// ._
	if isNotNeighbour(right, cell) && isNotNeighbour(down, cell) {
		count++
	}

	// interior

	// XX top left   (origin is bottom left)
	// X.
	if isNeighbour(up, cell) && isNeighbour(upRight, cell) && isNotNeighbour(right, cell) {
		count++
	}

	// XX top right
	// .X
	if isNeighbour(left, cell) && isNeighbour(down, cell) && isNotNeighbour(downLeft, cell) {
		count++
	}

	// X. bottom left (origin is bottom left)
	// XX
	if isNeighbour(up, cell) && isNeighbour(right, cell) && isNotNeighbour(upRight, cell) {
		count++
	}

	// .X bottom right  origin is bottom right
	// XX
	if isNeighbour(left, cell) && isNeighbour(up, cell) && isNotNeighbour(upLeft, cell) {
		count++
	}

	return count
}

func (g *Grid) left(cell *Cell) *Cell {
	return g.get(cell.x-1, cell.y)
}

func (g *Grid) right(cell *Cell) *Cell {
	return g.get(cell.x+1, cell.y)
}

func (g *Grid) up(cell *Cell) *Cell {
	return g.get(cell.x, cell.y-1)
}

func (g *Grid) down(cell *Cell) *Cell {
	return g.get(cell.x, cell.y+1)
}

func (g *Grid) downLeft(cell *Cell) *Cell {
	return g.get(cell.x-1, cell.y+1)
}

func (g *Grid) upLeft(cell *Cell) *Cell {
	return g.get(cell.x-1, cell.y-1)
}

func (g *Grid) upRight(cell *Cell) *Cell {
	return g.get(cell.x+1, cell.y-1)
}

type Cell struct {
	x      int
	y      int
	value  string
	key    string
	region int
}

func NewCell(x int, y int, value string) *Cell {
	key := fmt.Sprintf("%v.%v", x, y)
	c := Cell{x: x, y: y, key: key, value: value, region: -1}
	return &c
}

type Region struct {
	id    int
	cells []*Cell
}

func NewRegion(id int) *Region {
	r := Region{id: id, cells: make([]*Cell, 0)}
	return &r
}

func (r *Region) sides(grid *Grid) int {
	// the number of sides == the number of corners
	count := 0
	for _, c := range r.cells {
		count += grid.countCorners(c)
	}
	return count
}

func (r *Region) add(cell *Cell) {
	r.cells = append(r.cells, cell)
}

func (r *Region) area() int {
	return len(r.cells)
}

func (r *Region) perimeter(grid *Grid) int {
	p := 0
	for _, cell := range r.cells {
		p += 4 - len(grid.neighboursSameType(cell))
	}
	return p
}

func (r *Region) price(grid *Grid) int {
	return r.area() * r.perimeter(grid)
}

func (r *Region) new_price(grid *Grid) int {
	return r.area() * r.sides(grid)
}
