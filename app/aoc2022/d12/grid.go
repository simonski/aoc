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
	BestPath    *Path
	Iteration   int
}

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
					} else {
						point.VisitDirection = "v"
					}
				}

				if pathIndex == len(path.Points)-1 {
					point.VisitDirection = "*"
				}

				if pathIndex == -1 {
					point.VisitDirection = "." // point.Letter
				}

			}
			// fmt.Printf("GetPoint(%v,%v)=%v\n", col, row, point)
			line = fmt.Sprintf("%v%v", line, point.VisitDirection)
		}
		result = fmt.Sprintf("%v\n%v", result, line)
	}
	return result
}

// func (g *Grid) Clone() *Grid {
// 	g2 := &Grid{}
// 	g2.Rows = g.Rows
// 	g2.Cols = g.Cols
// 	g2.Points = make(map[string]*Point)
// 	for key, point := range g.Points {
// 		p2 := point.Clone()
// 		g2.Points[key] = p2
// 		if p2.IsStart {
// 			g2.Start = p2
// 		}
// 		if p2.IsDestination {
// 			g2.Destination = p2
// 		}
// 	}
// 	return g2
// }

type Path struct {
	Points []*Point
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
			key := fmt.Sprintf("%v_%v", colIndex, rowIndex)
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
	return &Grid{Points: points, Cols: width, Rows: height, Start: start, Destination: destination}
}

func (g *Grid) Neighbours(point *Point) []*Point {
	results := make([]*Point, 0)
	up := g.Get(point.Row-1, point.Col)
	down := g.Get(point.Row+1, point.Col)
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

// func (p *Point) Clone() *Point {
// 	p2 := &Point{}
// 	p2.Letter = p.Letter
// 	p2.VisitDirection = p.VisitDirection
// 	p2.Value = p.Value
// 	p2.Col = p.Col
// 	p2.Row = p.Row
// 	p2.TentativeDistance = p.TentativeDistance
// 	p2.Visited = p.Visited
// 	p2.IsStart = p.IsStart
// 	p2.IsDestination = p.IsDestination
// 	return p2
// }

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
	if up != nil && up != origin && !up.Visited { //!visited[up] {
		results = append(results, up)
	}
	if down != nil && down != origin && !down.Visited { //!visited[down] {
		results = append(results, down)
	}
	if left != nil && left != origin && !left.Visited { // !visited[left] {
		results = append(results, left)
	}
	if right != nil && right != origin && !right.Visited { // !visited[right] {
		results = append(results, right)
	}
	return results
}

func (graph *Grid) Get(row int, col int) *Point {
	key := fmt.Sprintf("%v_%v", col, row)
	return graph.Points[key]
}

func (graph *Grid) Put(row int, col int, point *Point) {
	key := fmt.Sprintf("%v_%v", col, row)
	graph.Points[key] = point
}

func (p *Point) Debug() string {
	return fmt.Sprintf("(%v,%v)=%v", p.Col, p.Row, p.Value)
}

// func dijkstra(graph *Grid, start *Point, destination *Point) {
// 	// start := graph.GetPoint(0, 0)
// 	fmt.Printf("start: %v\n", start.Debug())
// 	// destination := graph.GetPoint(graph.Width-1, graph.Height-1)
// 	fmt.Printf("dest : %v\n", destination.Debug())

// 	// 1. Mark all nodes unvisited. Create a set of all the unvisited nodes called the unvisited set.
// 	// visited := make(map[*XPoint]bool)

// 	/*
// 		2. Assign to every node a tentative distance value: set it to zero for our initial node and to infinity for all other nodes. The tentative distance of a node v is the length of the shortest path discovered so far between the node v and the starting node. Since initially no path is known to any other vertex than the source itself (which is a path of length zero), all other tentative distances are initially set to infinity. Set the initial node as current
// 	*/

// 	unvisited := make([]*Point, 0)
// 	for key := range graph.Points {
// 		p := graph.Points[key]
// 		if p != nil {
// 			p.TentativeDistance = math.MaxInt
// 			p.Visited = false
// 			// visited[p] = false
// 			unvisited = append(unvisited, p)
// 		}
// 	}
// 	start.TentativeDistance = 0
// 	current := start

// 	fmt.Printf("Marked unvisited\n")

// 	/*
// 		3. For the current node, consider all of its unvisited neighbors and calculate their tentative distances through the current node.

// 		Compare the newly calculated tentative distance to the current assigned value and assign the smaller one.

// 		For example, if the current node A is marked with a distance of 6, and the edge connecting it with a neighbor B has length 2, then the distance to B through A will be 6 + 2 = 8.

// 		If B was previously marked with a distance greater than 8 then change it to 8. Otherwise, the current value will be kept.
// 	*/
// 	// path := make([]*XPoint, 0)

// 	// visitCount := 0
// 	// lastPct := -1
// 	// size := len(graph.Points)
// 	// lastPct := -1

// 	counter := 0
// 	visitedCount := 0

// 	for {

// 		neighbours := graph.UnvisitedNeighbours(current)
// 		for _, pointB := range neighbours {
// 			if pointB.Visited {
// 				continue
// 			}
// 			newTentativeForPointB := current.TentativeDistance + pointB.Value
// 			if newTentativeForPointB < pointB.TentativeDistance {
// 				pointB.TentativeDistance = newTentativeForPointB
// 			}

// 		}

// 		/*
// 			4. When we are done considering all of the unvisited neighbors of the current node, mark the current node
// 			as visited and remove it from the unvisited set. A visited node will never be checked again.
// 		*/
// 		current.Visited = true // visited[current] = true
// 		visitedCount += 1

// 		/*
// 			5. If the destination node has been marked visited (when planning a route between two
// 				specific nodes) or if the smallest tentative distance among the nodes in the unvisited set is
// 				infinity (when planning a complete traversal; occurs when there is no connection between the
// 					initial node and remaining unvisited nodes), then stop. The algorithm has finished.
// 		*/
// 		if destination.Visited { // visited[destination] {
// 			// then stop
// 			break
// 		}

// 		/*
// 			6. Otherwise, select the unvisited node that is marked with the smallest tentative distance, set
// 			it as the new current node, and go back to step 3.
// 		*/

// 		minDistance := math.MaxInt
// 		for _, pointB := range unvisited {
// 			if pointB.Visited { // visited[pointB] {
// 				continue
// 			}
// 			if pointB.TentativeDistance < minDistance {
// 				a := current
// 				b := pointB
// 				fmt.Printf(" (%v,%v) moves to (%v,%v) (%v->%v)\n", a.Col, a.Row, b.Col, b.Row, a.Letter, b.Letter)
// 				current = pointB
// 				minDistance = current.TentativeDistance
// 			}
// 		}

// 		if counter%10000 == 0 {
// 			uv := make([]*Point, 0)
// 			for _, p := range unvisited {
// 				if !p.Visited {
// 					uv = append(uv, p)
// 				}
// 			}
// 			unvisited = uv

// 			sort.SliceStable(unvisited, func(i, j int) bool {
// 				return unvisited[i].TentativeDistance < unvisited[j].TentativeDistance
// 			})
// 		}

// 		// path = append(path, current)

// 	}

// 	fmt.Printf("Visited count is %v, Score is %v\n", visitedCount, destination.TentativeDistance)
// }

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
		b := neighbours[o]
		return a.Col < b.Col
	})

	return results
}
