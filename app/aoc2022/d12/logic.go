package d12

import (
	"fmt"
)

func HillClimb(debug bool, grid *Grid) *Path {
	path := NewPath()
	grid.BestPath = NewPath()
	start := grid.Start
	HillClimbRecursive(debug, grid, start, path)
	return grid.BestPath
}

func HillClimbRecursive(debug bool, grid *Grid, point *Point, path *Path) {
	// point.Visited = true
	grid.Iteration += 1
	// point.VisitDirection = "X"
	path.Add(point)
	if debug && grid.Iteration%10000 == 0 {
		fmt.Println(grid.Iteration)
		fmt.Println(grid.Debug(path))
	}
	// recursion 101; get out first
	bestPath := grid.BestPath
	// fmt.Printf("point(%v,%v), bestPath=%v\n", point.X, point.Y, bestCompletedPath)
	if point == grid.Destination {
		if bestPath.Size() == 0 {
			// if we are at our destination AND we dont have a path to compare to
			grid.BestPath = path
			fmt.Printf("HillClimbRecursive[%v], steps=%v, first path = '%v'\n", path.Size(), path)
			return
		} else if path.Size() < bestPath.Size() {
			// if we are at our destination AND this path is shortest
			fmt.Printf("HillClimbRecursive[%v], steps=%v, new best path = '%v'\n", path.Size(), path)
			grid.BestPath = path
			return
		} else {
			return
		}
	} else if path.Size() > bestPath.Size() && bestPath.Size() > 0 {
		// if our current path exceeds a completed path
		// fmt.Printf("HillClimbRecursive[%v], steps=%v, drop candidate path = '%v'\n", path.Size(), iteration, path)
		return
	}

	// or descend
	// do we have any neighbours we can visit?
	// if we don't, we need to return so that the pathfinding can continue earlier
	neighbours := grid.NeighboursNotInPath(point, path)
	if len(neighbours) == 0 {
		// fmt.Printf(">>> returning with no neighbours: %v\n", path)
		return
	}

	for _, neighbour := range neighbours {
		// if !neighbour.Visited {
		if neighbour.IsStart {
			continue
		}
		newPath := path.Clone()
		HillClimbRecursive(debug, grid, neighbour, newPath)
	}

}
