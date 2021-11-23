package aoc2018

import (
	"fmt"
)

type Grid struct {
	min_x    int
	min_y    int
	max_x    int
	max_y    int
	points   []*Point
	width    int
	height   int
	pointMap map[string]*Point
}

func NewGrid() *Grid {
	g := &Grid{}
	return g
}

func (g *Grid) Bounds() (int, int, int, int, int, int) {
	min_x := g.points[0].position_x
	min_y := g.points[0].position_y
	max_x := min_x
	max_y := min_y

	for _, point := range g.points {
		max_x = Max(max_x, point.position_x)
		max_y = Max(max_y, point.position_y)
		min_x = Min(min_x, point.position_x)
		min_y = Min(min_y, point.position_y)
	}
	width := max_x - min_x
	height := max_y - min_y

	return min_x, min_y, max_x, max_y, width, height
}

func (g *Grid) Snapshot(x int, y int, width int, height int) []string {
	var snapshot []string
	for yindex := y; yindex < y+height; yindex++ {
		line := ""
		for xindex := x; xindex < x+width; xindex++ {
			point := g.Get(xindex, yindex)
			if point == nil {
				line += "."
			} else {
				line += "#"
			}
		}
		snapshot = append(snapshot, line)
	}
	return snapshot
}

// indicates if the position x,y matches the passed search expression letter
func (g *Grid) Compare(letter []string, x int, y int) (bool, []string) {
	// in this case we need to match each character in letter
	// with each character in the grid at that point
	width := len(letter[0])
	height := len(letter)
	snapshot := g.Snapshot(x, y, width-1, height-1)

	for index := 0; index < len(letter); index++ {
		if letter[index] != snapshot[index] {
			return false, snapshot
		}
	}
	return true, snapshot
}

func (g *Grid) Draw() {
	min_x, min_y, max_x, max_y, _, _ := g.Bounds()

	for y := min_y; y <= max_y; y++ {
		for x := min_x; x <= max_x; x++ {

			// for y := g.min_y; y <= g.max_y; y++ {
			// 	for x := g.min_x; x <= g.max_x; x++ {

			p := g.Get(x, y)
			if p != nil {
				// if p.isTopLeft {
				// 	fmt.Print("T")
				// } else {
				fmt.Print("#")
				// }
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func (g *Grid) IsH(point *Point) bool {
	// return false
	// return false
	// #...#
	// #...#
	// #...#
	// #####
	// #...#
	// #...#
	// #...#
	// #...#

	// so an H has a point which has a sibling at
	// right vertical
	// x+4, y
	// x+4, y+1
	// x+4, y+2
	// x+4, y+3
	// x+4, y+4
	// x+4, y+5
	// x+4, y+6
	// x+4, y+7
	x := point.position_x
	y := point.position_y

	for index := 0; index <= 4; index++ {
		if g.Get(x+index, y+3) == nil {
			// bar not present
			// fmt.Printf("!got the horiz\n")
			return false
		}
	}
	fmt.Printf("IsH(); got a horiz line horiz\n")
	fmt.Printf("\n")
	x, y, maxx, maxy, width, height := g.Bounds()
	fmt.Printf("x,y,maxx,maxy,width,height=%v,%v %v,%v %v,%v\n", x, y, maxx, maxy, width, height)
	g.Draw()
	fmt.Printf("\n")
	return true

	for index := 0; index <= 7; index++ {
		if g.Get(x+4, y+index) == nil {
			// right bar not present
			return false
		}
	}
	fmt.Printf("IsH(); found a right bar\n")
	for index := 0; index <= 7; index++ {
		if g.Get(x, y+index) == nil {
			// left bar not present
			return false
		}
	}
	fmt.Printf("IsH(); found a left bar\n")
	// horizontal line
	// x, y+3
	// x+1, y+3
	// x+2, y+3
	// x+3, y+3
	// x+4, y+3
	return true

	// left and right must both be present

	// left vertical
	// x, y+1
	// x, y+2
	// x, y+3
	// x, y+4
	// x, y+5
	// x, y+6
	// x, y+7

}

func (g *Grid) Load(lines []string) {
	for _, line := range lines {
		p := NewPoint(line)
		g.AddPoint(p)
	}

	if g.min_x < 0 {
		g.width = Abs(g.min_x) + g.max_x
	} else {
		g.width = g.max_x
	}

	if g.min_y < 0 {
		g.height = Abs(g.min_y) + g.max_y
	} else {
		g.height = g.max_y
	}
}

func (g *Grid) Remap() {
	// g.pointMap = make(map[string]*Point)
	// // everything cycles to where it needs to be
	// for _, p := range g.points {
	// 	p.Remap()
	// 	g.pointMap[p.Key] = p
	// }
}

func (g *Grid) AddPoint(p *Point) {
	g.min_x = Min(p.position_x, g.min_x)
	g.max_x = Max(p.position_x, g.max_x)
	g.min_y = Min(p.position_y, g.min_y)
	g.max_y = Max(p.position_y, g.max_y)
	g.points = append(g.points, p)
}

func (g *Grid) Get(x int, y int) *Point {
	key := fmt.Sprintf("%v.%v", x, y)
	return g.pointMap[key]
}

func (g *Grid) Step(remap bool) {
	// g.step += 1
	if remap {
		g.pointMap = make(map[string]*Point)
	}
	// everything cycles to where it needs to be
	for _, p := range g.points {
		p.Step(remap)
		if remap {
			g.pointMap[p.Key] = p
		}
	}
}

func (g *Grid) CountLetters() int {
	h_counter := 0
	for _, p := range g.points {
		if g.IsH(p) {
			h_counter += 1
		}
	}
	return h_counter
}

func (g *Grid) Debug() {
	fmt.Printf("Grid.debug() min_x=%v, min_y=%v, max_x=%v, max_y=%v, width=%v, height=%v\n", g.min_x, g.min_y, g.max_x, g.max_y, g.width, g.height)
}

// returns the outermost points moving in their "opposite" direction
func (g *Grid) FindMovingBoundaryPoints() (*Point, *Point, *Point, *Point) {
	var leftMostMovingRight *Point
	var rightMostMovingLeft *Point
	var topMostMovingDown *Point
	var bottomMostMovingUp *Point
	for _, p := range g.points {
		if p.velocity_x > 0 {
			// moving right
			if leftMostMovingRight == nil || p.position_x < leftMostMovingRight.position_x {
				leftMostMovingRight = p
			}
		} else if p.velocity_x < 0 {
			// moving left
			if rightMostMovingLeft == nil || p.position_x > rightMostMovingLeft.position_x {
				rightMostMovingLeft = p
			}
		}

		if p.velocity_y > 0 {
			// moving down
			if topMostMovingDown == nil || p.position_y < topMostMovingDown.position_y {
				topMostMovingDown = p
			}
		} else if p.velocity_y < 0 {
			// moving up
			if bottomMostMovingUp == nil || p.position_y > bottomMostMovingUp.position_y {
				bottomMostMovingUp = p
			}
		}
	}
	return leftMostMovingRight, rightMostMovingLeft, topMostMovingDown, bottomMostMovingUp
}
