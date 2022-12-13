package d12

import (
	"fmt"
	"testing"
)

func Test_Grid1(t *testing.T) {
	g := NewGrid(TEST_DATA)
	if g.Cols != 8 {
		t.Fatalf("Expected 8 cols, got %v\n", g.Cols)
	}
}

func Test_Grid2(t *testing.T) {
	g := NewGrid(TEST_DATA)
	if g.Rows != 5 {
		t.Fatalf("Expected 5 rows, got %v\n", g.Rows)
	}
}

// func Test_Dijkstra1(t *testing.T) {
// 	g := NewGrid(TEST_DATA)
// 	var start *Point
// 	var destination *Point
// 	for _, node := range g.Points {
// 		if node.IsStart {
// 			start = node
// 		}
// 		if node.IsDestination {
// 			destination = node
// 		}
// 	}
// 	fmt.Println("")
// 	fmt.Println(g.DebugVisitDirection())
// 	fmt.Println("")
// 	dijkstra(g, start, destination)
// 	fmt.Println("")
// 	fmt.Println(g.DebugValue())
// 	fmt.Println("")
// 	fmt.Println(g.DebugVisitDirection())
// 	fmt.Println("")

// 	fmt.Println(g.DebugLetter())

// 	visited := 0
// 	not_visited := 0
// 	for _, node := range g.Points {
// 		if node.Visited {
// 			visited++
// 		} else {
// 			not_visited++
// 		}
// 	}
// 	fmt.Printf("visited %v, not visited %v\n", visited, not_visited)

// 	if g.Rows != 1 {
// 		t.Fatalf("Expected 5 rows, got %v\n", g.Rows)
// 	}
// }

func Test_HillClimb(t *testing.T) {
	grid := NewGrid(TEST_DATA)
	path := HillClimb(true, grid)
	fmt.Printf("%v\n", path)

	fmt.Printf("\n\nBest path size is %v\n\n", path.Size())

	if grid.Rows != 1 {
		t.Fatalf("Expected 5 rows, got %v\n", grid.Rows)
	}
}
