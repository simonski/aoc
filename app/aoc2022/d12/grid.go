package d12

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type Grid struct {
	Points      map[string]*Point
	Start       *Point
	Destination *Point
	Cols        int
	Rows        int
	// BestPath    *Path
	Iteration int
	// ClosedPaths map[string]bool
	// BigPaths    map[string]bool
}

func NewGrid(data string) *Grid {
	points := make(map[string]*Point)
	rows := strings.Split(data, "\n")
	width := 0
	height := 0
	var start *Point
	var destination *Point

	for rowIndex, row := range rows {
		height += 1
		width = 0
		for colIndex := 0; colIndex < len(row); colIndex++ {
			key := fmt.Sprintf("%v_%v", rowIndex, colIndex)
			value := row[colIndex : colIndex+1] //, _ := strconv.Atoi(row[colIndex : colIndex+1])
			point := NewPoint(value, rowIndex, colIndex)
			if point.IsStart {
				start = point
			}
			if point.IsDestination {
				destination = point
			}
			points[key] = point
			width++
			// fmt.Printf("Adding point%v\n", point)
		}
	}
	// closedPaths := make(map[string]bool)
	// bigPaths := make(map[string]bool)
	return &Grid{Points: points, Cols: width, Rows: height, Start: start, Destination: destination} //, ClosedPaths: closedPaths, BigPaths: bigPaths}
}

// func (g *Grid) AddBigPath(path *Path) {
// 	key := path.Key()
// 	g.BigPaths[key] = true
// }

// func (g *Grid) AddClosedPath(path *Path) {
// 	key := path.Key()
// 	g.ClosedPaths[key] = true
// }

// func (g *Grid) HasCheckedPath(path *Path) bool {
// 	s := path.Key()
// 	if g.ClosedPaths[s] == true || g.BigPaths[s] == true {
// 		return true
// 	}
// 	return false
// }

func (g *Grid) Debug(path *Path) string {
	// in this debug I build the direction by zooming over hte grid, then
	// seeing if the point in question exists in the path, in which case
	// I figre outthe direction from the prior point in the path
	// note: I make the "current" tail a * to stand out
	result := ""
	for row := 0; row < g.Rows; row++ {
		line := ""
		for col := 0; col < g.Cols; col++ {
			point := g.Get(row, col)
			if point.IsStart {
				point.VisitDirection = "S"
			} else if point.IsDestination {
				point.VisitDirection = "E"
			} else {
				pathIndex := path.IndexOf(point)
				if pathIndex > -1 && pathIndex+1 < len(path.Points) {
					nextPoint := path.Points[pathIndex+1]
					if nextPoint.IsAbove(point) {
						point.VisitDirection = "^"
					} else if nextPoint.IsLeft(point) {
						point.VisitDirection = "<"
					} else if nextPoint.IsRight(point) {
						point.VisitDirection = ">"
					} else if nextPoint.IsBelow(point) {
						point.VisitDirection = "v"
					}
				}

				if pathIndex == len(path.Points)-1 {
					point.VisitDirection = "*"
				}

				if pathIndex == -1 {
					point.VisitDirection = "." // point.Letter // "."
				}

			}
			// fmt.Printf("GetPoint(%v,%v)=%v\n", col, row, point)
			line = fmt.Sprintf("%v%v", line, point.VisitDirection)
		}
		result = fmt.Sprintf("%v\n%v", result, line)
	}
	return result
}

func (g *Grid) Neighbours(point *Point) []*Point {
	results := make([]*Point, 0)
	up := g.Get(point.Row+1, point.Col)
	down := g.Get(point.Row-1, point.Col)
	left := g.Get(point.Row, point.Col-1)
	right := g.Get(point.Row, point.Col+1)
	if up != nil {
		results = append(results, up)
	}
	if down != nil {
		results = append(results, down)
	}
	if left != nil {
		results = append(results, left)
	}
	if right != nil {
		results = append(results, right)
	}
	return results
}

type Path struct {
	Points []*Point
}

func (p *Path) Key() string {
	key := ""
	for _, point := range p.Points {
		key = fmt.Sprintf("%v->(%v,%v)", key, point.Col, point.Row)
	}
	return key
}

func (p *Path) Pop() *Point {
	size := p.Size()
	if size > 0 {
		point := p.Points[size-1]
		p.Points = p.Points[0 : size-1]
		return point
	} else {
		return nil
	}
}

func (p *Path) IsComplete() bool {
	index := len(p.Points) - 1
	return p.Points[index].IsDestination
}

func (p *Path) IndexOf(point *Point) int {
	for index, candidate := range p.Points {
		if candidate.Row == point.Row && candidate.Col == point.Col {
			return index
		}
	}
	return -1
}

func (p *Path) Contains(point *Point) bool {
	return p.IndexOf(point) > -1
}

func (p *Path) String() string {
	result := ""
	for _, point := range p.Points {
		result = fmt.Sprintf("%v -> %v (%v,%v)", result, point.Letter, point.Col, point.Row)
	}
	return result
}

func NewPath() *Path {
	p := Path{Points: make([]*Point, 0)}
	return &p
}

func (p *Path) Add(point *Point) {
	p.Points = append(p.Points, point)
}

func (p *Path) Size() int {
	return len(p.Points)
}

func (p *Path) Clone() *Path {
	p2 := NewPath()
	for _, point := range p.Points {
		p2.Add(point)
	}
	return p2
}

type Point struct {
	Letter            string
	VisitDirection    string
	Value             int
	Col               int
	Row               int
	TentativeDistance int
	Visited           bool
	IsStart           bool
	IsDestination     bool
}

func (p *Point) IsAbove(other *Point) bool {
	return p.Col == other.Col && p.Row+1 == other.Row
}

func (p *Point) IsBelow(other *Point) bool {
	return p.Col == other.Col && p.Row-1 == other.Row
}

func (p *Point) IsLeft(other *Point) bool {
	return p.Row == other.Row && p.Col+1 == other.Col
}

func (p *Point) IsRight(other *Point) bool {
	return p.Row == other.Row && p.Col-1 == other.Col
}

func NewPoint(letter string, row int, col int) *Point {
	point := &Point{Letter: letter, Col: col, Row: row, TentativeDistance: math.MaxInt}
	runeLetter := letter[0]
	if letter == "S" {
		runeLetter = "a"[0]
	} else if letter == "E" {
		runeLetter = "z"[0]
	}
	point.IsStart = letter == "S"
	point.IsDestination = letter == "E"
	point.Letter = letter
	char := rune(runeLetter)
	ascii := int(char)
	point.Value = ascii
	point.VisitDirection = "."
	return point
}

func (p *Point) String() string {
	return fmt.Sprintf("(%v,%v) (%v)=%v", p.Col, p.Row, p.Letter, p.Value)
}

func (g *Grid) DebugVisitDirection() string {
	result := ""
	for rowNum := 0; rowNum < g.Rows; rowNum++ {
		for colNum := 0; colNum < g.Cols; colNum++ {
			result = fmt.Sprintf("%v%v", result, g.Get(rowNum, colNum).VisitDirection)
		}
		result = fmt.Sprintf("%v\n", result)
	}
	return result
}

func (g *Grid) DebugValue() string {
	result := ""
	for rowNum := 0; rowNum < g.Rows; rowNum++ {
		for colNum := 0; colNum < g.Cols; colNum++ {
			result = fmt.Sprintf("%v %v", result, g.Get(rowNum, colNum).Value)
		}
		result = fmt.Sprintf("%v\n", result)
	}
	return result
}

func (g *Grid) DebugLetter() string {
	result := ""
	for rowNum := 0; rowNum < g.Rows; rowNum++ {
		for colNum := 0; colNum < g.Cols; colNum++ {
			result = fmt.Sprintf("%v%v", result, g.Get(rowNum, colNum).Letter)
		}
		result = fmt.Sprintf("%v\n", result)
	}
	return result
}

// // func NewNode(letter string) *Node {
// 	node := Node{}
// 	node.IsStart = letter == "S"
// 	if node.IsStart {
// 		letter = "a"
// 	}
// 	node.IsDestination = letter == "E"
// 	if node.IsDestination {
// 		letter = "z"
// 	}
// 	node.Letter = letter
// 	char := rune(letter[0])
// 	ascii := int(char)
// 	node.Value = ascii
// 	node.Visit = "."
// 	return &node
// }

func (g *Grid) UnvisitedNeighbours(origin *Point) []*Point { //, visited map[*XPoint]bool) []*XPoint {
	results := make([]*Point, 0)
	x := origin.Col
	y := origin.Row
	up := g.Get(y-1, x)
	down := g.Get(y+1, x)
	left := g.Get(y, x-1)
	right := g.Get(y, x+1)

	// fmt.Printf("origin=%v[%v], value=%v\n", origin, origin.Letter, origin.Value)

	if right != nil && right != origin && !right.Visited { // !visited[right] {
		keep := isHeightOk(origin, right)
		// fmt.Printf("    left=%v[%v], value=%v, heighOk=%v\n", right, right.Letter, right.Value, keep)
		if keep {
			results = append(results, right)
		}
	}
	if up != nil && up != origin && !up.Visited { //!visited[up] {
		keep := isHeightOk(origin, up)
		// fmt.Printf("    up=%v[%v], value=%v, heighOk=%v\n", up, up.Letter, up.Value, keep)
		if keep {
			results = append(results, up)
		}
	}
	if down != nil && down != origin && !down.Visited { //!visited[down] {
		keep := isHeightOk(origin, down)
		// fmt.Printf("    down=%v[%v], value=%v, heighOk=%v\n", down, down.Letter, down.Value, keep)
		if keep {
			results = append(results, down)
		}
	}
	if left != nil && left != origin && !left.Visited { // !visited[left] {
		keep := isHeightOk(origin, left)
		// fmt.Printf("    left=%v[%v], value=%v, heighOk=%v\n", left, left.Letter, left.Value, keep)
		if keep {
			results = append(results, left)
		}
	}
	return results
}

func (grid *Grid) Get(row int, col int) *Point {
	key := fmt.Sprintf("%v_%v", row, col)
	return grid.Points[key]
}

func (grid *Grid) Put(row int, col int, point *Point) {
	key := fmt.Sprintf("%v_%v", row, col)
	grid.Points[key] = point
}

func (p *Point) Debug() string {
	return fmt.Sprintf("(%v,%v)=%v", p.Col, p.Row, p.Value)
}

func (g *Grid) NeighboursNotInPath(point *Point, path *Path) []*Point {
	neighbours := g.Neighbours(point)
	results := make([]*Point, 0)
	for _, n := range neighbours {
		if path.IndexOf(n) == -1 {
			if n.Value <= point.Value || n.Value == point.Value+1 {
				results = append(results, n)
			}
		}
	}
	sort.Slice(results, func(i int, o int) bool {
		a := neighbours[i]
		// b := neighbours[o]
		if a.Col == point.Col {
			return false
		}
		return true
	})

	return results
}

func isHeightOk(a *Point, b *Point) bool {
	if a.Value > b.Value { // e.g. c to a
		return true
	}
	if a.Value == b.Value { // sameies ok
		return true
	}
	if a.Value+1 == b.Value { // a to b
		return true
	}
	return false
}
