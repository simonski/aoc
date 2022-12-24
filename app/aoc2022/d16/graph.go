package d16

import (
	"sort"
	"strconv"
	"strings"
)

/*
--- Day 05:  ---

*/

func NewGraph(input string) *Graph {
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

	graph2 := LoadGraph2(input)

	g := &Graph{NodeMap: nodeMap, Nodes: nodes, Graph2: graph2}
	cache := NewCache(g)
	g.Cache = cache
	g.ScoredNodes = g._GetScoredNodes()
	g.PathCache = make(map[int64]*Path)
	return g
}

type Graph struct {
	NodeMap     map[string]*Node
	Nodes       []*Node
	Time        int
	Graph2      *graph
	Cache       *Cache
	ScoredNodes []*Node
	PathCache   map[int64]*Path
}

// returns all nodes with a Value > 0 sorted by ID
func (g *Graph) GetScoredNodes() []*Node {
	return g.ScoredNodes
}

func (g *Graph) _GetScoredNodes() []*Node {

	nodes := make([]*Node, 0)
	for _, node := range g.NodeMap {
		if node.Value > 0 {
			nodes = append(nodes, node)
		}
	}

	sort.Slice(nodes, func(i int, j int) bool {
		n1 := nodes[i]
		n2 := nodes[j]
		return n1.ID < n2.ID
	})
	return nodes
}

func (g *Graph) Get(id string) *Node {
	return g.NodeMap[id]
}

func (g *Graph) Size() int {
	return len(g.NodeMap)
}

func (g *Graph) NewPathAllOn() *Path {
	scored_nodes := g.GetScoredNodes()
	p := NewPath()
	for _, node := range scored_nodes {
		p.Open(node)
	}
	return p
}

func (g *Graph) NewPathFromInt(value int64) *Path {
	p := g.PathCache[value]
	if p != nil {
		return p.Clone()
	}
	path := NewPath()
	binary_value := strconv.FormatInt(value, 2) // 1111011
	scored_nodes := g.GetScoredNodes()
	for index := 0; len(binary_value) < len(scored_nodes); index++ {
		binary_value = "0" + binary_value
	}

	for index := 0; index < len(binary_value); index++ {
		value := binary_value[index : index+1]
		if value == "1" {
			node := scored_nodes[index]
			path.Open(node)
		}
	}
	g.PathCache[value] = path

	return path
}
