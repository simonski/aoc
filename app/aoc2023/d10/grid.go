package d10

import (
	"fmt"
	"strings"
)

func TERMINAL_GOLD(in string) string {
	return fmt.Sprintf("\033[33m%v\033[0m", in)
}

type PointType struct {
	North  bool
	South  bool
	East   bool
	West   bool
	letter string
}

func (pt *PointType) Render() string {
	if pt.North && pt.South {
		return "|"
	} else if pt.East && pt.West {
		return "-"
	} else if pt.East && pt.North {
		return "╰"
	} else if pt.West && pt.North {
		return "╯"
	} else if pt.South && pt.West {
		return "╮"
	} else if pt.South && pt.East {
		return "╭"
	} else {
		return "."
	}
}

func NewPointType(letter string) *PointType {
	pt := PointType{}
	pt.letter = letter
	if letter == "|" {
		pt.North = true
		pt.South = true
	} else if letter == "-" {
		pt.East = true
		pt.West = true
	} else if letter == "L" {
		pt.East = true
		pt.North = true
	} else if letter == "J" {
		pt.West = true
		pt.North = true
	} else if letter == "7" {
		pt.South = true
		pt.West = true
	} else if letter == "F" {
		pt.South = true
		pt.East = true
	}
	return &pt
}

type Point struct {
	x         int // 0...n left-right
	y         int //0 ... n top-bottom
	pt        *PointType
	connected bool
	visited   bool
	distance  int
	isStart   bool
}

func (p *Point) Debug() string {
	return fmt.Sprintf("(%v,%v) %v", p.x, p.y, p.pt.Render())
}

func NewPoint(x int, y int, pt *PointType) *Point {
	p := &Point{x: x, y: y, pt: pt}
	p.connected = false
	p.distance = -1
	return p
}

type Grid struct {
	data [][]*Point
	// current_s_pos *Point
	start_s_pos *Point
	s_x         int
	s_y         int
}

func NewGrid(data string) *Grid {
	g := Grid{}
	g.data = make([][]*Point, 0)
	rows := strings.Split(data, "\n")
	for rownum, row := range rows {
		rowpoints := make([]*Point, 0)
		for col := 0; col < len(row); col++ {
			colvalue := row[col : col+1]
			pt := NewPointType(colvalue)
			p := NewPoint(col, rownum, pt)
			rowpoints = append(rowpoints, p)
			if colvalue == "S" {
				fmt.Printf("S at %v\n", p.Debug())
				p.isStart = true
				g.s_x = col
				g.s_y = rownum
				g.start_s_pos = p
			}

		}
		g.data = append(g.data, rowpoints)
	}

	// we need to "fix" the s_point by workingout what's around it
	if g.start_s_pos != nil {
		// fmt.Printf("Found S at %v, calculating original type:", g.start_s_pos.Debug())
		above := g.Get(g.start_s_pos.x, g.start_s_pos.y-1)
		below := g.Get(g.start_s_pos.x, g.start_s_pos.y+1)
		left := g.Get(g.start_s_pos.x-1, g.start_s_pos.y)
		right := g.Get(g.start_s_pos.x+1, g.start_s_pos.y)

		// fmt.Printf("Above=%v, Below=%v, Left=%v, Right=%v\n", above.pt.Render(), below.pt.Render(), left.pt.Render(), right.pt.Render())

		g.start_s_pos.pt.North = above.pt.South
		g.start_s_pos.pt.South = below.pt.North
		g.start_s_pos.pt.East = right.pt.West
		g.start_s_pos.pt.West = left.pt.East

		// fmt.Printf("Starting S point is %v\n", g.start_s_pos.pt.Render())
		// fmt.Printf("CanNoveNorth : %v\n", g.start_s_pos.pt.North)
		// fmt.Printf("CanNoveSouth : %v\n", g.start_s_pos.pt.South)
		// fmt.Printf("CanNoveEast : %v\n", g.start_s_pos.pt.East)
		// fmt.Printf("CanNoveWest : %v\n", g.start_s_pos.pt.West)

	} else {
		// fmt.Println("No Starting S point found")

	}

	return &g
}

func (g *Grid) Get(x int, y int) *Point {
	if x < 0 || x >= len(g.data) {
		return NoPoint()
	}
	if y < 0 || y >= len(g.data[0]) {
		return NoPoint()
	}
	return g.data[y][x]
}

func NoPoint() *Point {
	p := NewPoint(-1, -1, NewPointType(""))
	return p
}

func (g *Grid) Debug() string {
	showS := true
	result := ""
	for rowNum, row := range g.data {
		for colNum, col := range row {
			if showS && g.s_x == colNum && g.s_y == rowNum {
				result = fmt.Sprintf("%v%v", result, TERMINAL_GOLD("S"))
			} else {
				result = fmt.Sprintf("%v%v", result, col.pt.Render())
			}
		}
		result = fmt.Sprintf("%v\n", result)
	}
	return result
}

func (g *Grid) DebugDistance() string {
	result := ""
	for _, row := range g.data {
		for _, col := range row {
			if col.visited {
				result = fmt.Sprintf("%v%v", result, col.distance)
			} else {
				result = fmt.Sprintf("%v%v", result, col.pt.Render())
			}
		}
		result = fmt.Sprintf("%v\n", result)
	}
	return result
}

func (g *Grid) East(p *Point) (bool, *Point) {
	if p.pt.East {
		ep := g.Get(p.x+1, p.y)
		return true, ep
	} else {
		// fmt.Printf("g.East(%v) %v - Cannot move east\n", p, p.pt.Render())
		return false, nil
	}
}

func (g *Grid) West(p *Point) (bool, *Point) {
	if p.pt.West {
		ep := g.Get(p.x-1, p.y)
		return true, ep
	} else {
		return false, nil

	}
}

func (g *Grid) North(p *Point) (bool, *Point) {
	if p.pt.North {
		ep := g.Get(p.x, p.y-1)
		return true, ep
	} else {
		return false, nil
	}
}

func (g *Grid) South(p *Point) (bool, *Point) {
	if p.pt.South {
		ep := g.Get(p.x, p.y+1)
		return true, ep
	} else {
		return false, nil
	}
}

func (g *Grid) Walk(p *Point) int {

	distance := 0
	p.distance = distance
	// walk in two directions from this point until we meet
	points := g.PointsFrom(p)
	// fmt.Printf("Walk[%v], childPoints (%v): \n", distance, len(points))
	for _, p := range points {
		fmt.Printf("    (%v,%v) %v\n", p.x, p.y, p.pt.letter)
	}
	for {
		distance += 1
		next_points := make([]*Point, 0)
		for _, subPoint := range points {
			if !subPoint.visited {
				// fmt.Printf("Moving to subPoint: (%v,%v)\n", subPoint.x, subPoint.y)
				subPoint.visited = true
				subPoint.distance = distance

				points := g.PointsFrom(subPoint)
				next_points = append(next_points, points...)
			}
		}
		points = next_points
		// fmt.Printf("Walk[%v], childPoints to visit next: %v\n", distance, points)
		for _, p := range points {
			fmt.Printf("    (%v,%v) %v\n", p.x, p.y, p.pt.letter)
		}
		if len(points) == 0 {
			// fmt.Printf("No points left to visit, returning max distance %v\n", distance-1)
			return distance - 1
		}

	}

}

func (g *Grid) IfEast(p *Point) (*Point, bool) {
	cp := g.Get(p.x+1, p.y)
	if cp != nil && p.pt.East && cp.pt.West {
		return cp, true
	}
	return nil, false
}
func (g *Grid) IfWest(p *Point) (*Point, bool) {
	cp := g.Get(p.x-1, p.y)
	if cp != nil && p.pt.West && cp.pt.East {
		return cp, true
	}
	return nil, false
}
func (g *Grid) IfNorth(p *Point) (*Point, bool) {
	cp := g.Get(p.x, p.y-1)
	if cp != nil && p.pt.North && cp.pt.South {
		return cp, true
	}
	return nil, false
}
func (g *Grid) IfSouth(p *Point) (*Point, bool) {
	cp := g.Get(p.x, p.y+1)
	if cp != nil && p.pt.South && cp.pt.North {
		return cp, true
	}
	return nil, false
}

func (g *Grid) PointsFrom(p *Point) []*Point {
	east, east_ok := g.IfEast(p)
	west, west_ok := g.IfWest(p)
	north, north_ok := g.IfNorth(p)
	south, south_ok := g.IfSouth(p)

	results := make([]*Point, 0)
	if east_ok {
		results = append(results, east)
	}
	if west_ok {
		results = append(results, west)
	}
	if south_ok {
		results = append(results, south)
	}
	if north_ok {
		results = append(results, north)
	}
	return results
}
