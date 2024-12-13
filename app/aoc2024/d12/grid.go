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
