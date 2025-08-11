package d16

import "fmt"

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

func bfs_and_bfs(root *Node) {

	// bfs is me, I go first.
	// dfs is the elephant, they go second
	fmt.Println("bds_and_dfs")

	bfs_nodes_to_visit := NewQ()
	bfs_nodes_to_visit.append(root)

	// dfs
	dfs_nodes_to_visit := NewQ()
	dfs_nodes_to_visit.append(root)

	for time := 0; time < 26; time++ {

		// bfs
		bfs_currentnode := bfs_nodes_to_visit.take_first()
		bfs_nodes_to_visit.prepend_children(bfs_currentnode)
		//do something
		fmt.Printf("[%v] BFS (me) visited %v\n", time, bfs_currentnode)
		if bfs_nodes_to_visit.size() == 0 {
			break
		}

		dfs_currentnode := dfs_nodes_to_visit.take_first()
		dfs_nodes_to_visit.append_children(dfs_currentnode)
		fmt.Printf("[%v] DFS (trunky) visited %v\n", time, dfs_currentnode)
		//do something
		if dfs_nodes_to_visit.size() == 0 {
			break
		}

	}

}
