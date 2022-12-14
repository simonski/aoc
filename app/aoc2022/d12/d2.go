package d12

import (
	"fmt"
	"math"
)

func dijkstra_v2(grid *Grid, VERBOSE bool) int {

	start := grid.Start
	destination := grid.Destination
	if VERBOSE {
		fmt.Printf("start: %v\n", start.Debug())
		fmt.Printf("dest : %v\n", destination.Debug())
	}

	stuck := 0
	stuck_max := 1000

	// 1. Mark all nodes unvisited. Create a set of all the unvisited nodes called the unvisited set.
	// visited := make(map[*XPoint]bool)

	/*
		2. Assign to every node a tentative distance value: set it to zero for our initial node and to infinity for all other nodes. The tentative distance of a node v is the length of the shortest path discovered so far between the node v and the starting node. Since initially no path is known to any other vertex than the source itself (which is a path of length zero), all other tentative distances are initially set to infinity. Set the initial node as current
	*/

	unvisited := make([]*Point, 0)
	for _, p := range grid.Points {
		// p := grid.Points[key]
		p.TentativeDistance = math.MaxInt
		p.Visited = false
		unvisited = append(unvisited, p)
	}
	start.TentativeDistance = 0
	current := start

	/*
		3. For the current node, consider all of its unvisited neighbors and calculate their tentative distances through the current node.

		Compare the newly calculated tentative distance to the current assigned value and assign the smaller one.

		For example, if the current node A is marked with a distance of 6, and the edge connecting it with a neighbor B has length 2, then the distance to B through A will be 6 + 2 = 8.

		If B was previously marked with a distance greater than 8 then change it to 8. Otherwise, the current value will be kept.
	*/
	// path := make([]*XPoint, 0)

	// visitCount := 0
	// lastPct := -1
	// size := len(grid.Points)
	// lastPct := -1

	// counter := 0
	// visitedCount := 0

	for {
		// counter += 1
		// pct := int((100.0 / size) * counter)
		// if pct != lastPct {
		// 	fmt.Printf("%v\n", pct)
		// 	lastPct = pct
		// }

		// if counter%1000 == 0 {
		// 	if VERBOSE {
		// 		fmt.Printf("Counter = %v, VisitedCount=%v\n", counter, visitedCount)
		// 	}
		// }

		// visitCount += 1
		// pct := (100.0 / len(visited) * visitCount)
		// // fmt.Println(pct)
		// if pct != lastPct {
		// 	lastPct = pct
		// 	fmt.Printf("%v\n", pct)
		// }

		neighbours := grid.UnvisitedNeighbours(current) //, visited)
		if VERBOSE {
			fmt.Printf("%v has %v unvisited neighbours. %v\n", current, len(neighbours), neighbours)
		}
		if len(neighbours) == 0 {
			stuck++
		} else {
			stuck = 0
		}
		if stuck > stuck_max {
			return 99999999999999
		}

		for _, pointB := range neighbours {
			if pointB.Visited {
				continue
			}
			newTentativeForPointB := current.TentativeDistance + 1 // pointB.Value
			if newTentativeForPointB < pointB.TentativeDistance {
				pointB.TentativeDistance = newTentativeForPointB
			}

		}

		/*
			4. When we are done considering all of the unvisited neighbors of the current node, mark the current node as visited and remove it from the unvisited set. A visited node will never be checked again.
		*/
		current.Visited = true // visited[current] = true
		// visitedCount += 1

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

	if VERBOSE {
		fmt.Printf("Score is %v\n", destination.TentativeDistance)
	}
	return destination.TentativeDistance
}
