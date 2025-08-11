package d12

import (
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

func Test_Map(t *testing.T) {
	m := make(map[string]*Point)
	m["foo"] = &Point{Col: 4, Row: 9}

	testPoint := m["bar"]
	testPoint2 := m["foo"]
	if testPoint != nil {
		t.Fatal("a")
	}

	if testPoint2 == nil {
		t.Fatal("b")
	}

}
