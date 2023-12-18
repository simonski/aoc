package utils

import (
	"strconv"
	"strings"
)

type Q struct {
	data []*Node
}

func NewQ() *Q {
	return &Q{data: make([]*Node, 0)}
}

func (q *Q) append(node *Node) {
	q.data = append(q.data, node)
}

func (q *Q) append_children(node *Node) {
	q.data = append(q.data, node.Children...)
}

func (q *Q) prepend(node *Node) {
	//not very nice - not certain how to do quick copies
	x := make([]*Node, 0)
	x = append(x, node)
	x = append(x, q.data...)
	q.data = x
}

func (q *Q) prepend_children(node *Node) {
	//not very nice - not certain how to do quick copies
	x := make([]*Node, 0)
	x = append(x, node.Children...)
	x = append(x, q.data...)
	q.data = x
}

func (q *Q) take_first() *Node {
	if q.size() == 0 {
		return nil
	} else if q.size() == 1 {
		result := q.data[0]
		q.data = make([]*Node, 0)
		return result
	}
	result := q.data[0]
	data := q.data[1:]
	q.data = data
	return result
}

func (q *Q) take_last() *Node {
	if q.size() == 0 {
		return nil
	} else if q.size() == 1 {
		result := q.data[0]
		q.data = make([]*Node, 0)
		return result
	}
	result := q.data[0]
	data := q.data[1:]
	q.data = data
	return result
}

func (q *Q) size() int {
	return len(q.data)
}

func bfs_iterative(root *Node) {
	nodes_to_visit := NewQ()
	nodes_to_visit.append(root)
	for {
		currentnode := nodes_to_visit.take_first()
		nodes_to_visit.prepend_children(currentnode)
		//do something
		if nodes_to_visit.size() == 0 {
			break
		}
	}
}

func dfs_iterative(root *Node) {
	nodes_to_visit := NewQ()
	nodes_to_visit.append(root)
	for {
		currentnode := nodes_to_visit.take_first()
		nodes_to_visit.append_children(currentnode)
		//do something
		if nodes_to_visit.size() == 0 {
			break
		}
	}
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
