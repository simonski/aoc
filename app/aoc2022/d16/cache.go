package d16

import "sort"

type Cache struct {
	nodeIndexes map[string]int
}

func NewCache(g *Graph) *Cache {
	nodes := g.GetScoredNodes()
	sort.Slice(nodes, func(i int, j int) bool {
		n1 := nodes[i]
		n2 := nodes[j]
		return n1.ID < n2.ID
	})

	nodeIndexes := make(map[string]int)
	for index := 0; index < len(nodes); index++ {
		nodeIndexes[nodes[index].ID] = index
	}

	cache := Cache{}
	cache.nodeIndexes = nodeIndexes
	return &cache
}

func (c *Cache) GetNodeIndex(node *Node) int {
	return c.nodeIndexes[node.ID]
}

func (c *Cache) Get(node *Node, path *Path, time int) (*Path, bool) {
	return nil, false
}

func (c *Cache) Put(node *Node, path *Path, time int, bestPath *Path) {

}
