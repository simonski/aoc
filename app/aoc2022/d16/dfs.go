package d16

import "fmt"

// there is something about "knowing" from your current path the best achieveable score
// so we need access to the current "best" scores.  If our route *cannot* give us a better score
// then we can quit out of it quickly - avoiding all permutations and therefore reducing the search space
// ABOVE NEEDS INTEGRATING
// CODE NEEDS COMMITTING

func (g *Graph) dfs(source *Node, current_path *Path, time int, MAX_TIME int, VERBOSE bool) int {
	// from SOURCE, list available path

	result, hit := g.Cache.Get(source, current_path, time)
	if hit {
		return result
	}

	fmt.Printf("dfs[t=%v] path %v\n", time, current_path)

	// get all closed nodes to find their paths from here
	available_nodes := make([]*Node, 0)
	for _, node := range g.GetScoredNodes() {
		if !current_path.Contains(node) && node != source {
			available_nodes = append(available_nodes, node)
		}
	}

	best := current_path.Score()
	for _, destination := range available_nodes {

		this_time := time
		// we check each one - rather than have some heuristics, we just walk the path
		// th}e "winner" (which should be a complete path) is the best score
		subpath := current_path.Clone()
		_, steps := g.Graph2.getPath(source.ID, destination.ID)

		// steps is the sequence of steps required to get to this destination
		// so we will add each entry into our new subpath
		for index := 1; index < len(steps); index++ { // first step is origin
			node_id := steps[index]
			this_time += 1
			nodeToMoveTo := g.Get(node_id)
			subpath.Move(nodeToMoveTo)
			if this_time == MAX_TIME {
				break
			}
		}
		if this_time == MAX_TIME {
			continue
		}

		this_time += 1
		subpath.Open(destination)

		if this_time == MAX_TIME {
			continue
		}

		// open NODE
		// assign NODE to results (time, value, node)

		result := g.dfs(destination, subpath, this_time, MAX_TIME, VERBOSE)
		if result > best {
			best = result
		}

	}

	g.Cache.Put(source, current_path, time, best)

	return best

}
