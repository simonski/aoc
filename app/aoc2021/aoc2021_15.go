package aoc2021

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	// utils "github.com/simonski/aoc/utils"
)

/*
	// minTotal := 813 TOO HIGH
--- Day 15: Chiton ---
You've almost reached the exit of the cave, but the walls are getting closer together. Your submarine can barely still fit, though; the main problem is that the walls of the cave are covered in chitons, and it would be best not to bump any of them.

The cavern is large, but has a very low ceiling, restricting your motion to two dimensions. The shape of the cavern resembles a square; a quick scan of chiton density produces a map of risk level throughout the cave (your puzzle input). For example:

1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581
You start in the top left position, your destination is the bottom right position, and you cannot move diagonally. The number at each position is its risk level; to determine the total risk of an entire path, add up the risk levels of each position you enter (that is, don't count the risk level of your starting position unless you enter it; leaving it adds no risk to your total).

Your goal is to find a path with the lowest total risk. In this example, a path with the lowest total risk is highlighted here:

1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581
The total risk of this path is 40 (the starting position is never entered, so its risk is not counted).

What is the lowest total risk of any path from the top left to the bottom right?

*/

type PointGraph struct {
	Points map[string]*XPoint
	Width  int
	Height int
}

func (graph *PointGraph) Debug() string {
	line := ""
	for y := 0; y < graph.Height; y++ {
		for x := 0; x < graph.Width; x++ {
			point := graph.GetPoint(x, y)
			line = fmt.Sprintf("%v %v ", line, point.Value)
		}
		line += "\n"
	}
	return line
}

func growBy(graph *PointGraph, growBy int) *PointGraph {

	graphClone := graph.CopyAndIncrement(0)
	// col := 0
	// row := 0
	incrementBy := 0
	for tileY := 0; tileY < growBy; tileY++ {
		incrementBy = tileY
		for tileX := 0; tileX < growBy; tileX++ {
			// newValue := 8 + incrementBy
			// if newValue > 9 {
			// 	newValue = newValue % 9
			// }
			newGraph := graph.CopyAndIncrement(incrementBy)

			// we are currently at tile x, tile y
			for nx := 0; nx < newGraph.Width; nx++ {
				for ny := 0; ny < newGraph.Height; ny++ {
					new_x := nx + (tileX * graph.Width)
					new_y := ny + (tileY * graph.Height)
					point := newGraph.GetPoint(nx, ny)
					graphClone.SetPoint(new_x, new_y, point)
				}
			}
			// fmt.Printf("tile=%v,%v increment by %v, 8->%v\n", x, y, incrementBy, newValue)
			incrementBy += 1
		}
	}
	graphClone.Width *= growBy
	graphClone.Height *= growBy
	return graphClone
}

func (graph *PointGraph) CopyAndIncrement(toAdd int) *PointGraph {
	newGraph := &PointGraph{}
	newGraph.Width = graph.Width
	newGraph.Height = graph.Height
	newGraph.Points = make(map[string]*XPoint)
	for y := 0; y < graph.Height; y++ {
		for x := 0; x < graph.Width; x++ {
			value := graph.GetPoint(x, y).Value
			value += toAdd
			if value > 9 {
				value = value % 9
			}

			p := &XPoint{Value: value, X: x, Y: y, TentativeDistance: math.MaxInt}
			key := fmt.Sprintf("%v,%v", x, y)
			newGraph.Points[key] = p
		}
	}
	return newGraph
}

func (graph *PointGraph) GetPoint(x int, y int) *XPoint {
	key := fmt.Sprintf("%v,%v", x, y)
	return graph.Points[key]
}

func (graph *PointGraph) SetPoint(x int, y int, point *XPoint) {
	key := fmt.Sprintf("%v,%v", x, y)
	point.X = x
	point.Y = y
	graph.Points[key] = point
}

type XPoint struct {
	Value             int
	X                 int
	Y                 int
	TentativeDistance int
	Visited           bool
}

func (p *XPoint) Debug() string {
	return fmt.Sprintf("(%v,%v)=%v", p.X, p.Y, p.Value)
}

func NewPointGraph(data string) *PointGraph {
	points := make(map[string]*XPoint)
	rows := strings.Split(data, "\n")
	width := 0
	height := 0
	for rowIndex, row := range rows {
		height += 1
		width = 0
		for colIndex := 0; colIndex < len(row); colIndex++ {
			key := fmt.Sprintf("%v,%v", colIndex, rowIndex)
			value, _ := strconv.Atoi(row[colIndex : colIndex+1])
			point := &XPoint{Value: value, X: colIndex, Y: rowIndex, TentativeDistance: math.MaxInt}
			points[key] = point
			width++
			// fmt.Printf("Adding point%v\n", point)
		}
	}
	return &PointGraph{Points: points, Width: width, Height: height}
}

func (g *PointGraph) UnvisitedNeighbours(origin *XPoint) []*XPoint { //, visited map[*XPoint]bool) []*XPoint {
	results := make([]*XPoint, 0)
	x := origin.X
	y := origin.Y
	up := g.GetPoint(x, y-1)
	down := g.GetPoint(x, y+1)
	left := g.GetPoint(x-1, y)
	right := g.GetPoint(x+1, y)
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

func dijkstra(prefix string, graph *PointGraph) {
	start := graph.GetPoint(0, 0)
	fmt.Printf("%v start: %v\n", prefix, start.Debug())
	destination := graph.GetPoint(graph.Width-1, graph.Height-1)
	fmt.Printf("%v dest : %v\n", prefix, destination.Debug())

	// 1. Mark all nodes unvisited. Create a set of all the unvisited nodes called the unvisited set.
	// visited := make(map[*XPoint]bool)

	/*
		2. Assign to every node a tentative distance value: set it to zero for our initial node and to infinity for all other nodes. The tentative distance of a node v is the length of the shortest path discovered so far between the node v and the starting node. Since initially no path is known to any other vertex than the source itself (which is a path of length zero), all other tentative distances are initially set to infinity. Set the initial node as current
	*/

	unvisited := make([]*XPoint, 0)
	for key := range graph.Points {
		p := graph.Points[key]
		p.TentativeDistance = math.MaxInt
		p.Visited = false
		// visited[p] = false
		unvisited = append(unvisited, p)
	}
	start.TentativeDistance = 0
	current := start

	fmt.Printf("Marked unvisited\n")

	/*
		3. For the current node, consider all of its unvisited neighbors and calculate their tentative distances through the current node.

		Compare the newly calculated tentative distance to the current assigned value and assign the smaller one.

		For example, if the current node A is marked with a distance of 6, and the edge connecting it with a neighbor B has length 2, then the distance to B through A will be 6 + 2 = 8.

		If B was previously marked with a distance greater than 8 then change it to 8. Otherwise, the current value will be kept.
	*/
	// path := make([]*XPoint, 0)

	// visitCount := 0
	// lastPct := -1
	size := len(graph.Points)
	lastPct := -1

	counter := 0
	visitedCount := 0

	for {
		counter += 1
		pct := int((100.0 / size) * counter)
		if pct != lastPct {
			fmt.Printf("%v\n", pct)
			lastPct = pct
		}

		if counter%1000 == 0 {
			fmt.Printf("Counter = %v, VisitedCount=%v\n", counter, visitedCount)
		}

		// visitCount += 1
		// pct := (100.0 / len(visited) * visitCount)
		// // fmt.Println(pct)
		// if pct != lastPct {
		// 	lastPct = pct
		// 	fmt.Printf("%v\n", pct)
		// }

		neighbours := graph.UnvisitedNeighbours(current) //, visited)
		// fmt.Printf("neighbours of %v:\n", current.Debug())
		// for _, neigh := range neighbours {
		// 	fmt.Printf("   %v\n", neigh.Debug())
		// }
		// fmt.Println()
		for _, pointB := range neighbours {
			if pointB.Visited {
				continue
			}
			newTentativeForPointB := current.TentativeDistance + pointB.Value
			if newTentativeForPointB < pointB.TentativeDistance {
				pointB.TentativeDistance = newTentativeForPointB
			}

			// if current.TentativeDistance+pointB.Value < pointB.TentativeDistance {
			// 	pointB.TentativeDistance = current.TentativeDistance + pointB.Value
			// }
		}

		/*
			4. When we are done considering all of the unvisited neighbors of the current node, mark the current node as visited and remove it from the unvisited set. A visited node will never be checked again.
		*/
		current.Visited = true // visited[current] = true
		visitedCount += 1

		/*
			5. If the destination node has been marked visited (when planning a route between two specific nodes) or if the smallest tentative distance among the nodes in the unvisited set is infinity (when planning a complete traversal; occurs when there is no connection between the initial node and remaining unvisited nodes), then stop. The algorithm has finished.
		*/
		if destination.Visited { // visited[destination] {
			// then stop
			break
		}

		/*
			6. Otherwise, select the unvisited node that is marked with the smallest tentative distance, set it as the new current node, and go back to step 3.
		*/

		minDistance := math.MaxInt
		for _, pointB := range unvisited {
			if pointB.Visited { // visited[pointB] {
				continue
			}
			if pointB.TentativeDistance < minDistance {
				current = pointB
				minDistance = current.TentativeDistance
			}
		}

		if counter%10000 == 0 {
			uv := make([]*XPoint, 0)
			for _, p := range unvisited {
				if !p.Visited {
					uv = append(uv, p)
				}
			}
			unvisited = uv

			sort.SliceStable(unvisited, func(i, j int) bool {
				return unvisited[i].TentativeDistance < unvisited[j].TentativeDistance
			})
		}

		// path = append(path, current)

	}

	fmt.Printf("Score is %v\n", destination.TentativeDistance)
}

func debug(graph *PointGraph, visited map[*XPoint]bool, current *XPoint) string {
	line := ""
	for y := 0; y < graph.Height; y++ {
		for x := 0; x < graph.Width; x++ {
			p := graph.GetPoint(x, y)
			// v := visited[p]
			var output string
			if p == current {
				output = fmt.Sprintf("C(%v)", p.TentativeDistance)
			} else {
				distance := p.TentativeDistance
				if visited[p] {
					output = fmt.Sprintf("%v", distance)
					// output = "V"
				} else {
					if distance == math.MaxInt {
						output = "X"
					} else {
						output = fmt.Sprintf("%v", distance)
					}
				}
			}
			line += output
			line += " "
		}
		line += "\n"
	}
	return line
}

// rename this to the year and day in question
func (app *Application) Y2021D15P1() {
	// graph := NewPointGraph(DAY_2021_15_DATA)
	// dijkstra("D15Part1", graph)
}

func (app *Application) Y2021D15P2() {
	graph := NewPointGraph(DAY_2021_15_DATA)
	bigGraph := growBy(graph, 5)
	fmt.Printf("%v, %v, %v points.\n", bigGraph.Width, bigGraph.Height, len(bigGraph.Points))
	// fmt.Printf("%v\n", bigGraph.Debug())
	dijkstra("D15Part2", bigGraph)
}

// func (app *Application) Y2021D15P1() {
// 	growBy5(DAY_2021_15_DATA)
// 	dijkstra(DAY_2021_15_DATA)
// }
