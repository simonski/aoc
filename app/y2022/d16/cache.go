package d16

import (
	"fmt"
	"sort"
)

type Cache struct {
	Graph       *Graph
	nodeIndexes map[string]int
	hits        int
	misses      int
	keys        map[string]bool
	data        map[string]int
	max_value   int
	path        *Path
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
	cache.keys = make(map[string]bool)
	cache.data = make(map[string]int)
	cache.Graph = g
	return &cache
}

func (c *Cache) GetNodeIndex(node *Node) int {
	return c.nodeIndexes[node.ID]
}

func (g *Graph) CreateKey(node *Node, path *Path, time int) string {
	key := fmt.Sprintf("%v|%v|%v", node.ID, path.Key(g), time)
	// key := fmt.Sprintf("%v_%v", node.ID, path.Key(g)) //, time)
	return key
}

func (c *Cache) Get(node *Node, path *Path, time int) (int, bool) {
	key := c.Graph.CreateKey(node, path, time)
	if c.keys[key] {
		// fmt.Printf("Cache(%v) HIT, value=%v\n", key, c.data[key])
		c.hits += 1
		return c.data[key], true
	} else {
		// fmt.Printf("Cache(%v) MISS, value=%v\n", key, c.data[key])
		c.misses += 1
		return 0, false
	}
}

func (c *Cache) Put(node *Node, path *Path, time int, value int) {
	key := c.Graph.CreateKey(node, path, time)
	// fmt.Printf("Cache(%v) PUT, value=%v\n", key, c.data[key])
	c.keys[key] = true
	c.data[key] = value

	if value > c.max_value {
		c.max_value = value
		c.path = path
	}
}
