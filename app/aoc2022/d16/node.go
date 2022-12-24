package d16

import (
	"strconv"
	"strings"
)

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
