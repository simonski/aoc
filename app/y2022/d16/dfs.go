package d16

import (
	"fmt"
	"strconv"

	"github.com/simonski/goutils"
)

// CACHE time on the DFS brute force -
// there is something about "knowing" from your current path the best achieveable score
// so we need access to the current "best" scores.  If our route *cannot* give us a better score
// then we can quit out of it quickly - avoiding all permutations and therefore reducing the search space

// from SOURCE, list available path
func (g *Graph) dfs(source *Node, current_path *Path, time int, VERBOSE bool) int {

	// not 5336 - too high (this was because the graph for node distance was hardcoded witht he
	// test data
	result, hit := g.Cache.Get(source, current_path, time)
	if hit {
		return result
	}

	// get all closed nodes to find their paths from here
	available_nodes := make([]*Node, 0)
	for _, node := range g.GetScoredNodes() {
		if !current_path.Contains(node) && node != source {
			available_nodes = append(available_nodes, node)
		}
	}

	best := 0
	for _, destination := range available_nodes {

		time_remaining := time
		// we check each one - rather than have some heuristics, we just walk the path
		// the "winner" (which should be a complete path) is the best score
		subpath := current_path.Clone()
		_, steps := g.Graph2.getPath(source.ID, destination.ID)

		// fmt.Printf("source=%v, destination=%v, steps=%v\n", source, destination, steps)
		// steps is the sequence of steps required to get to this destination
		// so we will add each entry into our new subpath
		for index := 1; index < len(steps); index++ { // first step [0] is origin so drop it
			node_id := steps[index]
			nodeToMoveTo := g.Get(node_id)
			subpath.Move(nodeToMoveTo)
		}
		time_remaining = time - (len(steps) - 1) - 1
		if time_remaining <= 0 {
			continue
		}
		subpath.Open(destination)

		result := g.dfs(destination, subpath, time_remaining, VERBOSE) + destination.Value*time_remaining
		if result > best {
			best = result
		}

	}

	g.Cache.Put(source, current_path, time, best)

	return best

}

func (g *Graph) dfs2(source *Node, current_path *Path, time int, VERBOSE bool) int {

	all_on := g.NewPathAllOn()
	imax, _ := strconv.ParseInt(all_on.Key(g), 2, 64)

	m := 0
	b := g.NewPathAllOn()
	ib := b.IntValue(g)

	aa := g.Get("AA")
	maxvalue := (imax + 1) / 2
	var i int64
	for i = 0; i < maxvalue; i++ {
		path1 := g.NewPathFromInt(i)
		path2 := g.NewPathFromInt(ib ^ i)

		fmt.Printf("path1=%v, i=%v, intval=%v\n", path1.Key(g), i, path1.IntValue(g))
		fmt.Printf("path2=%v, ib^i=%v, intval=%v\n", path2.Key(g), ib^i, path2.IntValue(g))

		m = goutils.Max(m, g.dfs(aa, path1, 26, VERBOSE)+g.dfs(aa, path2, 26, VERBOSE))
		fmt.Printf("m=%v, i=%v, max=%v\n", m, i, maxvalue)
	}

	return m
}
