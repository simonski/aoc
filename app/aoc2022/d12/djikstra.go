package d12

import (
	"fmt"
	"math"
)

func Djikstra(grid *Grid) []*Point {
	// start := graph.GetPoint(0, 0)
	path := make([]*Point, 0)
	start := grid.Start
	destination := grid.Destination

	fmt.Printf("Djikstra: start: %v\n", start.Debug())
	// destination := graph.GetPoint(graph.Width-1, graph.Height-1)
	fmt.Printf("Djikstra: dest : %v\n", destination.Debug())

	/*
		1. Mark all nodes unvisited. Create a set of all the unvisited nodes called the unvisited set.
		2. Assign to every node a tentative distance value: set it to zero for our initial node and to infinity for all other nodes. The tentative distance of a node v is the length of the shortest path discovered so far between the node v and the starting node. Since initially no path is known to any other vertex than the source itself (which is a path of length zero), all other tentative distances are initially set to infinity.
		Set the initial node as current
	*/

	unvisited := make([]*Point, 0)
	for key := range grid.Points {
		p := grid.Points[key]
		// if p != nil {
		p.TentativeDistance = math.MaxInt
		p.Visited = false
		unvisited = append(unvisited, p)
	}
	start.TentativeDistance = 0
	current := start
	path = append(path, current)

	/*
		3. For the current node, consider all of its unvisited neighbors and calculate their tentative distances through the current node.
		Compare the newly calculated tentative distance to the current assigned value and assign the smaller one.
		For example, if the current node A is marked with a distance of 6, and the edge connecting it with a neighbor B has length 2, then the distance to B through A will be 6 + 2 = 8.

		If B was previously marked with a distance greater than 8 then change it to 8. Otherwise, the current value will be kept.
	*/
	visitedCount := 0

	for {

		neighbours := grid.UnvisitedNeighbours(current)
		for _, pointB := range neighbours {
			// if pointB.Visited {
			// 	continue
			// }

			newTentativeForPointB := current.TentativeDistance + 1 // all distances are 1 as long as the neighbour is valid pointB.Value
			if newTentativeForPointB < pointB.TentativeDistance {
				pointB.TentativeDistance = newTentativeForPointB
			}

		}

		/*
			4. When we are done considering all of the unvisited neighbors of the current node, mark the current node
			as visited and remove it from the unvisited set. A visited node will never be checked again.
		*/
		current.Visited = true // visited[current] = true
		visitedCount += 1

		/*
			5. If the destination node has been marked visited (when planning a route between two
				specific nodes) or if the smallest tentative distance among the nodes in the unvisited set is
				infinity (when planning a complete traversal; occurs when there is no connection between the
					initial node and remaining unvisited nodes), then stop. The algorithm has finished.
		*/
		if destination.Visited { // visited[destination] {
			// then stop
			break
		}

		/*
			6. Otherwise, select the unvisited node that is marked with the smallest tentative distance, set
			it as the new current node, and go back to step 3.
		*/

		minDistance := math.MaxInt
		for _, pointB := range unvisited {
			if pointB.Visited { // visited[pointB] {
				continue
			}
			// if !isNeighbour(current, pointB) {
			// 	continue
			// }
			// if !isHeightOk(current, pointB) {
			// 	continue
			// }
			fmt.Printf(" current is (%v,%v)[%v], pointB is (%v,%v)[%v]\n", current.Col, current.Row, current.Letter, pointB.Col, pointB.Row, pointB.Letter)
			if pointB.TentativeDistance < minDistance {
				// if pointB.Value-1 == current.Value || pointB.Value < current.Value {
				a := current
				b := pointB
				fmt.Printf(" (%v,%v) moves to (%v,%v) (%v->%v)\n", a.Col, a.Row, b.Col, b.Row, a.Letter, b.Letter)
				path = append(path, pointB)
				current = pointB
				minDistance = current.TentativeDistance
				// }
			}
		}

		// if counter%10000 == 0 {
		// 	uv := make([]*Point, 0)
		// 	for _, p := range unvisited {
		// 		if !p.Visited {
		// 			uv = append(uv, p)
		// 		}
		// 	}
		// 	unvisited = uv

		// 	sort.SliceStable(unvisited, func(i, j int) bool {
		// 		return unvisited[i].TentativeDistance < unvisited[j].TentativeDistance
		// 	})
		// }

		// path = append(path, current)

	}
	fmt.Println(visitedCount)

	// 	fmt.Printf("Visited count is %v, Score is %v\n", visitedCount, destination.TentativeDistance)
	// }

	return path
}

// 4521 too high
// 4921
// Dean, Emily, John Mason

func isNeighbour(a *Point, b *Point) bool {
	if a.Row+1 == b.Row && a.Col == b.Col {
		return true // up
	} else if a.Row-1 == b.Row && a.Col == b.Col {
		return true // down
	} else if a.Col-1 == b.Col && a.Row == b.Row {
		return true // left
	} else if a.Col+1 == b.Col && a.Row == b.Row {
		return true // right
	}

	return false
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
