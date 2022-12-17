package d16

import (
	"fmt"
	"strconv"
	"strings"
)

/*
--- Day 05:  ---

*/

func LoadGraph(input string) map[string]*Node {
	splits := strings.Split(input, "\n")
	// Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
	nodes := make(map[string]*Node)
	for _, line := range splits {
		node := NewNode(line)
		nodes[node.ID] = node
	}

	for _, node := range nodes {
		for _, id := range node.ChildrenIDs {
			child := nodes[id]
			node.Children = append(node.Children, child)
		}
	}
	return nodes

}

type Node struct {
	ID             string
	Children       []*Node
	ChildrenIDs    []string
	IsOpen         bool
	Value          int
	OpenedAtMinute int
}

func NewNode(input string) *Node {
	line := strings.ReplaceAll(input, "Valve ", "")
	line = strings.ReplaceAll(line, "has flow rate=", "")
	line = strings.ReplaceAll(line, " tunnels lead to valves ", "")
	line = strings.ReplaceAll(line, " tunnel leads to valve ", "")
	line = strings.ReplaceAll(line, " tunnel lead to valves ", "")
	fmt.Printf("'%v' becomes '%v'\n", input, line)
	node_splits := strings.Split(line, ";")
	n1 := strings.Split(node_splits[0], " ")
	node_id := n1[0]
	flow_rate, _ := strconv.Atoi(n1[1])
	children_ids := strings.Split(strings.ReplaceAll(node_splits[1], " ", ""), ",")
	node := Node{ID: node_id, IsOpen: false, Value: flow_rate, ChildrenIDs: children_ids, Children: make([]*Node, 0)}
	return &node
}
