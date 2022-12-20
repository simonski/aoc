package d16

import (
	"fmt"
	"strconv"
	"strings"
)

/*
--- Day 05:  ---

*/

func LoadGraph(input string) *Graph {
	splits := strings.Split(input, "\n")
	// Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
	nodeMap := make(map[string]*Node)
	nodes := make([]*Node, 0)
	for _, line := range splits {
		node := NewNode(line)
		nodeMap[node.ID] = node
		nodes = append(nodes, node)
	}

	for _, node := range nodes {
		for _, id := range node.ChildrenIDs {
			child := nodeMap[id]
			node.Children = append(node.Children, child)
			child.Parent = node
		}
	}

	graph2 := LoadGraph2(TEST_DATA)

	return &Graph{NodeMap: nodeMap, Nodes: nodes, Graph2: graph2}

}

type Graph struct {
	NodeMap map[string]*Node
	Nodes   []*Node
	Time    int
	Graph2  *graph
}

func (g *Graph) CalculatePathScore(path []*Node) int {
	return 0
}

// returns all nodes with a Value > 0
func (g *Graph) GetScoredNodes() []*Node {

	nodes := make([]*Node, 0)
	for _, node := range g.NodeMap {
		if node.Value > 0 {
			nodes = append(nodes, node)
		}
	}
	return nodes
}

func (g *Graph) Get(id string) *Node {
	return g.NodeMap[id]
}

func (g *Graph) IsOpen() bool {
	return true
}

func (g *Graph) Size() int {
	return len(g.NodeMap)
}

func (g *Graph) Tick() int {
	return 0
}

type Node struct {
	ID                         string
	Children                   []*Node
	ChildrenIDs                []string
	IsOpen                     bool
	Value                      int
	OpenedAtMinute             int
	Parent                     *Node
	Dijkstra_TentativeDistance int
	Dijkstra_Visited           bool
}

func (n *Node) String() string {
	return n.ID
}

func (n *Node) Djikstra_UnvisitedNeighbours() []*Node {
	results := make([]*Node, 0)
	for _, node := range n.Children {
		if !node.Dijkstra_Visited {
			results = append(results, node)
		}
	}
	return results
}

func NewNode(input string) *Node {
	line := strings.ReplaceAll(input, "Valve ", "")
	line = strings.ReplaceAll(line, "has flow rate=", "")
	line = strings.ReplaceAll(line, " tunnels lead to valves ", "")
	line = strings.ReplaceAll(line, " tunnel leads to valve ", "")
	line = strings.ReplaceAll(line, " tunnel lead to valves ", "")
	// fmt.Printf("'%v' becomes '%v'\n", input, line)
	node_splits := strings.Split(line, ";")
	n1 := strings.Split(node_splits[0], " ")
	node_id := n1[0]
	flow_rate, _ := strconv.Atoi(n1[1])
	children_ids := strings.Split(strings.ReplaceAll(node_splits[1], " ", ""), ",")
	node := Node{ID: node_id, IsOpen: false, Value: flow_rate, ChildrenIDs: children_ids, Children: make([]*Node, 0)}
	return &node
}

// returns the total unopened value when walking this path in the tree
func (n *Node) UnopenedValue() int {
	fmt.Printf("n.UnopenedValue(%v)\n", n.ID)
	if len(n.Children) == 0 {
		if n.IsOpen {
			return 0
		} else {
			return n.Value
		}
	} else {
		value := 0
		if !n.IsOpen {
			value += n.Value
		}
		for _, child := range n.Children {
			if child.Parent != n {
				value += child.UnopenedValue()
			}
		}
		return value
	}
}

func (n *Node) CountUnopenedChildren() int {
	if len(n.Children) == 0 {
		return 0
	} else {
		value := 0
		if !n.IsOpen {
			value += 1
		}
		for _, child := range n.Children {
			if child.Parent != n {
				value += child.CountUnopenedChildren()
			}
		}
		return value
	}
}
