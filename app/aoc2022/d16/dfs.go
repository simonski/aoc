package d16

import "fmt"

// there is something about "knowing" from your current path the best achieveable score
// so we need access to the current "best" scores.  If our route *cannot* give us a better score
// then we can quit out of it quickly - avoiding all permutations and therefore reducing the search space
// ABOVE NEEDS INTEGRATING
// CODE NEEDS COMMITTING

// from SOURCE, list available path
func (g *Graph) dfs(source *Node, current_path *Path, time int, VERBOSE bool) int {

	// not 5336 - too high
	result, hit := g.Cache.Get(source, current_path, time)
	if hit {
		return result
	}

	// fmt.Printf("dfs[t=%v] path %v\n", time, current_path)

	// get all closed nodes to find their paths from here
	available_nodes := make([]*Node, 0)
	for _, node := range g.GetScoredNodes() {
		if !current_path.Contains(node) && node != source {
			available_nodes = append(available_nodes, node)
		}
	}

	best := 0 // current_path.Score()
	// if len(available_nodes) == 0 {

	// var best_path *Path
	// }
	for _, destination := range available_nodes {

		time_remaining := time
		// we check each one - rather than have some heuristics, we just walk the path
		// th}e "winner" (which should be a complete path) is the best score
		subpath := current_path.Clone()
		_, steps := g.Graph2.getPath(source.ID, destination.ID)

		// fmt.Printf("source=%v, destination=%v, steps=%v\n", source, destination, steps)
		// steps is the sequence of steps required to get to this destination
		// so we will add each entry into our new subpath
		for index := 1; index < len(steps); index++ { // first step is origin
			node_id := steps[index]
			nodeToMoveTo := g.Get(node_id)
			subpath.Move(nodeToMoveTo)
		}
		time_remaining = time - (len(steps) - 1) - 1
		if time_remaining <= 0 {
			continue
		}
		subpath.Open(destination)

		if len(available_nodes) == 1 {
			fmt.Println(subpath)
		}
		// open NODE
		// assign NODE to results (time, value, node)
		result := g.dfs(destination, subpath, time_remaining, VERBOSE) + destination.Value*time_remaining
		if result > best {
			best = result
			// g.Cache.Put(source, subpath, this_time, best)
			// best_path = subpath.Clone()
		}

	}

	g.Cache.Put(source, current_path, time, best)
	// if best_path != nil {
	// g.Cache.Put(source, best_path, time, best)
	// }

	return best

}
