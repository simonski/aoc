package d12

import (
	"fmt"
)

func HillClimb(debug bool, grid *Grid) *Path {
	path := NewPath()
	bestPath := NewPath()
	start := grid.Start
	return HillClimbRecursive(debug, 0, grid, start, path, bestPath)
}

func HillClimbRecursive(debug bool, iteration int, grid *Grid, point *Point, path *Path, bestCompletedPath *Path) *Path {
	// point.Visited = true
	// point.VisitDirection = "X"
	path.Add(point)
	if debug {
		fmt.Println(grid.Debug(path))
	}
	// recursion 101; get out first
	// fmt.Printf("point(%v,%v), bestPath=%v\n", point.X, point.Y, bestCompletedPath)
	if point == grid.Destination {
		if bestCompletedPath.Size() == 0 {
			// if we are at our destination AND we dont have a path to compare to
			// fmt.Printf("HillClimbRecursive[%v], steps=%v, first path = '%v'\n", path.Size(), iteration, path)
			return path
		} else if path.Size() < bestCompletedPath.Size() {
			// if we are at our destination AND this path is shortest
			fmt.Printf("HillClimbRecursive[%v], steps=%v, new best path = '%v'\n", path.Size(), iteration, path)
			return path
		} else {
			return bestCompletedPath
		}
	} else if path.Size() > bestCompletedPath.Size() && bestCompletedPath.Size() > 0 {
		// if our current path exceeds a completed path
		// fmt.Printf("HillClimbRecursive[%v], steps=%v, drop candidate path = '%v'\n", path.Size(), iteration, path)
		return bestCompletedPath
	}

	// or descend
	neighbours := grid.Neighbours(point)
	canWalk := false // we need at least one neighbour to check
	for _, neighbour := range neighbours {
		if path.IndexOf(neighbour) == -1 {
			// we haven't looked at this neighbour
			canWalk = true
		}
	}
	if !canWalk {
		bestCompletedPath.Pop()
		return bestCompletedPath
	}

	for _, neighbour := range neighbours {
		// if !neighbour.Visited {

		if !path.Contains(neighbour) { // not visited
			// fmt.Printf("%v not visited, current is %v\n", neighbour, point)
			if neighbour.Value <= point.Value || neighbour.Value == point.Value+1 {
				newPath := path.Clone()
				bestCompletedPath = HillClimbRecursive(debug, iteration+1, grid, neighbour, newPath, bestCompletedPath)
			}
		} else {
			// fmt.Printf("%v visited, ignoring.\n", neighbour)
		}
	}
	return bestCompletedPath

}
