package d16

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

/*
--- Day 05:  ---

*/

type Puzzle struct {
	title string
	year  string
	day   string
	input string
	lines []string
}

func NewPuzzleWithData(input string) *Puzzle {
	p := Puzzle{year: "2022", day: "16", title: "Proboscidea Volcanium"}
	p.Load(input)
	return &p
}

func NewPuzzle() *Puzzle {
	return NewPuzzleWithData(REAL_DATA)
}

func (puzzle *Puzzle) Load(input string) {
	lines := strings.Split(input, "\n")
	puzzle.input = input
	puzzle.lines = lines
}

func (puzzle *Puzzle) Part1() {
	graph2 := LoadGraph2(TEST_DATA)
	n1 := os.Args[4]
	n2 := os.Args[5]
	fmt.Println(graph2.getPath(n1, n2))

	puzzle.Part1Attempt1()

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("xxxxxxxxxxxxxxxx")
	fmt.Println("xxxxxxxxxxxxxxxx")
	fmt.Println("xxxxxxxxxxxxxxxx")
	fmt.Println("xxxxxxxxxxxxxxxx")
	fmt.Println("xxxxxxxxxxxxxxxx")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	// graph := LoadGraph(TEST_DATA)
	// aa := graph.Get("AA")
	// path := NewPathX()
	// time := 0
	// MAX_TIME := 30
	// VERBOSE := true
	// best_path := graph.walkies(aa, path, time, MAX_TIME, VERBOSE)
	// fmt.Printf("\nBest=\n%v\n", best_path)

	graph := LoadGraph(REAL_DATA)
	aa := graph.Get("AA")
	path := NewPathX()
	time := 0
	MAX_TIME := 30
	VERBOSE := true
	best_path := graph.walkies(aa, path, time, MAX_TIME, VERBOSE)
	fmt.Printf("\nBest=\n%v\n", best_path)

}

func NewPath(path []int) []int {
	p := make([]int, len(path))
	for index := 0; index < len(path); index++ {
		p[index] = path[index]
	}
	return p
}

// func (puzzle *Puzzle) Part1Attempt2() {
// 	graph := LoadGraph(TEST_DATA)
// 	value := 0
// 	path := make([]*Node, 0)
// 	aa := graph.Get("AA")
// 	path = append(path, aa)
// 	VERBOSE := false
// 	for time := 30; time > 0; time-- {
// 		value += Walk(aa, time, VERBOSE)
// 	}
// }
// func Walk(node *Node, time int, VERBOSE bool) int {
// 	time--
// 	if time == 0 {
// 		return 0
// 	}
// 	for _, child := range node.Children {
// 		t := time
// 		/Walk(child, t, VERBOSE)

// 	}
// }

/*
a node has a weight
visit each weighted node if possible
in best order of weights
*/

func (puzzle *Puzzle) Part1Attempt1() {
	graph := LoadGraph(TEST_DATA)

	// while time < 0
	// 1. build a list of paths to closed values with a score
	// 2. sort them by their best score at this time (score = function of values * time remaining)
	// 3. exclude any that take "too long" to open
	// 4. choose the 1st path

	// 1. build a list of paths that are closed with a score
	aa := graph.Get("AA")
	VERBOSE := true

	// fmt.Println(">>>>>")
	// data := make([]string, 0)
	// for _, node := range graph.Nodes {
	// 	node_id := node.ID
	// 	data = append(data, node_id)
	// }
	// fmt.Println(data)

	// for v := range itertools.PermutationsStr(data, len(data)) {
	// 	fmt.Println(v)
	// }
	// fmt.Println(">>>>>")
	// time.Sleep(1000)
	value := graph.GetAvailablePaths(aa, 29, VERBOSE)
	fmt.Printf("Value is %v\n", value)

}

// returns the closed nodes with a non zero value
func (g *Graph) GetClosedNonZeroNodes(current *Node) []*Node {
	result := make([]*Node, 0)
	for _, node := range g.Nodes {
		if node != current && !node.IsOpen && node.Value > 0 {
			result = append(result, node)
		}
	}
	return result
}

// returns the paths to the closed nodes with a non zero value, sorted by their value over time descending
func (g *Graph) GetAvailablePaths(source *Node, time int, VERBOSE bool) int {
	if time == 0 {
		return 0
	}
	closedNonZeroNodes := g.GetClosedNonZeroNodes(source)
	// highest value first
	// sort.Slice(availableNodesFromHere, func(i int, j int) bool {
	// 	return availableNodesFromHere[i].Value > availableNodesFromHere[j].Value
	// })
	// return result

	// closest closed node
	sort.Slice(closedNonZeroNodes, func(i int, j int) bool {
		n1 := closedNonZeroNodes[i]
		n2 := closedNonZeroNodes[j]
		d1, _ := g.Graph2.getPath(source.ID, n1.ID)
		d2, _ := g.Graph2.getPath(source.ID, n2.ID)
		return d1 < d2
	})

	// closest highest value node
	sort.Slice(closedNonZeroNodes, func(i int, j int) bool {
		n1 := closedNonZeroNodes[i]
		n2 := closedNonZeroNodes[j]
		d1, _ := g.Graph2.getPath(source.ID, n1.ID)
		d2, _ := g.Graph2.getPath(source.ID, n2.ID)

		if d1 == d2 {
			return n1.Value > n2.Value
		}

		return d1 < d2
	})

	// sort by opportunity
	sort.Slice(closedNonZeroNodes, func(i int, j int) bool {
		n1 := closedNonZeroNodes[i]
		n2 := closedNonZeroNodes[j]
		d1, _ := g.Graph2.getPath(source.ID, n1.ID)
		d2, _ := g.Graph2.getPath(source.ID, n2.ID)

		n1_timeCost := d1 + 1
		n1_opportunity := (time - n1_timeCost) * n1.Value

		n2_timeCost := d2 + 1
		n2_opportunity := (time - n2_timeCost) * n2.Value

		return n1_opportunity > n2_opportunity

	})

	// is the opportunity cost of the nearest/highest value better than the opportunity cost
	// of the "next" nearest, high value?

	// // maximise the opportunity cost
	// sort.Slice(closedNonZeroNodes, func(i int, j int) bool {
	// 	n1 := closedNonZeroNodes[i]
	// 	n2 := closedNonZeroNodes[j]
	// 	d1, _ := g.Graph2.getPath(source.ID, n1.ID)
	// 	d2, _ := g.Graph2.getPath(source.ID, n2.ID)

	// 	o1 := (time - d1 + 1) * n1.Value
	// 	o2 := (time - d2 + 1) * n2.Value

	// 	return o1 > o2
	// })

	if VERBOSE {
		fmt.Printf("[%v], time=%v, closedNonZeroNodes=%v\n", source.ID, time, len(closedNonZeroNodes))
		for _, destination := range closedNonZeroNodes {
			_, path := g.Graph2.getPath(source.ID, destination.ID)
			depth := len(path) - 1
			timeCost := depth + 1
			opportunity := (time - timeCost) * destination.Value
			fmt.Printf("  %v - depth is %v (path %v), cost to open is %v, value is %v, opportunity is %v * %v = %v\n", destination.ID, depth, path, timeCost, destination.Value, time-timeCost, destination.Value, opportunity)
		}
	}

	if len(closedNonZeroNodes) == 0 {
		return 0
	}

	for _, destination := range closedNonZeroNodes {
		// 1. build a path through the graph to each currently open node
		_, path := g.Graph2.getPath(source.ID, destination.ID)
		// 2. Node score that based on its length and the value it can provide over time
		depth := len(path) - 1
		moves := depth
		opens := 1
		timeCost := moves + opens
		remainingTime := time - timeCost
		fmt.Printf("  %v move+open cost is %v, remaining time will be %v\n", destination.ID, timeCost, remainingTime)
		pathValue := destination.Value * remainingTime
		if remainingTime == 0 {
			return pathValue
		} else {
			if VERBOSE {
				opportunity := remainingTime * destination.Value
				fmt.Printf(">>> OPEN %v, will provide %v starting at time %v (total %v)\n", destination.ID, destination.Value, remainingTime, opportunity)
			}
			destination.IsOpen = true
			return pathValue + g.GetAvailablePaths(destination, remainingTime, VERBOSE)
		}
	}
	return 0
}

// Depth first walk turning on

func (g *Graph) walk(node *Node, path []int, MAX_SIZE int) {
	path = append(path, 0) // we have just walked to this path
	fmt.Printf("walk(%v), time=%v, path=%v\n", node.ID, MAX_SIZE-len(path), path)
	if len(path) >= MAX_SIZE {
		total := 0
		for index := 0; index < MAX_SIZE; index++ {
			time := MAX_SIZE - index
			value := path[index] * time
			total += value
		}
		fmt.Println(total)
		return
		// return total
	}

	// either it is open or closed
	// either it has a value or it does not

	// if closed + 0
	// always pass by
	/// else if closed + > 0
	// 1: walk by
	// 2: open, then choose child
	// else if open
	// for each child
	// walk to child

	if node.Value > 0 && !node.IsOpen {
		// open this node
		path = append(path, node.Value) // we have arrived a this path, now we open it
		if len(path) == MAX_SIZE {
			return
		}

	} else if node.UnopenedValue() > 0 {
		// desend to open a child		for _
		for index := 0; index < len(node.Children); index++ {
			child := node.Children[index]
			if !child.IsOpen {
				// p := NewPath(path)
				g.walk(child, path, MAX_SIZE)
				if len(path) == MAX_SIZE {
					return
				}
			}
		}

		// p := NewPath(path)
		// child := node.Children[index]
		// g.walk(child, p)
	} else {
		// move up

	}
	if node.Value == 0 {
		// closed and value 0 - walk past this into any children
		for index := 0; index < len(node.Children); index++ {
			// p := NewPath(path)
			child := node.Children[index]
			g.walk(child, path, MAX_SIZE)
			if len(path) >= MAX_SIZE {
				return
			}
		}
	} else if !node.IsOpen {
		p := NewPath(path)
		node.IsOpen = true
		p = append(p, node.Value)
		for index := 0; index < len(node.Children); index++ {
			// p := NewPath(p)
			child := node.Children[index]
			g.walk(child, path, MAX_SIZE)
			if len(path) >= MAX_SIZE {
				return
			}
		}
	} else if node.IsOpen {
		// then we walk to a child
		for index := 0; index < len(node.Children); index++ {
			// p := NewPath(path)
			child := node.Children[index]
			g.walk(child, path, MAX_SIZE)
			if len(path) >= MAX_SIZE {
				return
			}
		}

	}

}

// func (graph *Graph) TryWalk(time int, path []int) int {
// 	source := graph.Get("AA")
// 	nodes := graph.GetScoredNodes()
// 	for _, destination := range nodes {
// 		path := graph.Dijkstra(source, destination, true)
// 		fmt.Printf("%v->%v, path=%v\n", source, destination, path)
// 	}
// 	return 0
// }

func (g *Graph) Walk(node *Node) {

}

func (puzzle *Puzzle) Part2() {
	puzzle.Load(REAL_DATA)
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}

// func main() {
// 	fmt.Println("Dijkstra")
// 	// Example
// 	graph := newGraph()
// 	graph.addEdge("S", "B", 4)
// 	graph.addEdge("S", "C", 2)
// 	graph.addEdge("B", "C", 1)
// 	graph.addEdge("B", "D", 5)
// 	graph.addEdge("C", "D", 8)
// 	graph.addEdge("C", "E", 10)
// 	graph.addEdge("D", "E", 2)
// 	graph.addEdge("D", "T", 6)
// 	graph.addEdge("E", "T", 2)
// 	fmt.Println(graph.getPath("S", "T"))
// }
