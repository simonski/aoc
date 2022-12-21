package d16

import "fmt"

// there is something about "knowing" from your current path the best achieveable score
// so we need access to the current "best" scores.  If our route *cannot* give us a better score
// then we can quit out of it quickly - avoiding all permutations and therefore reducing the search space
// ABOVE NEEDS INTEGRATING
// CODE NEEDS COMMITTING

func (g *Graph) dfs(source *Node, current_path *Path, time int, MAX_TIME int, VERBOSE bool) *Path {
	// from SOURCE, list available path

	result, hit := g.Cache.Get(source, current_path, time)
	if hit {
		return result
	}

	if time == MAX_TIME {
		if VERBOSE {
			fmt.Printf("dfs[t=%v] exiting with path %v\n", time, current_path)
		}
		return current_path
	}
	fmt.Printf("dfs[t=%v] path %v\n", time, current_path)

	// get all closed nodes to find their paths from here
	available_nodes := make([]*Node, 0)
	for _, node := range g.GetScoredNodes() {
		if !current_path.Contains(node) && node != source {
			available_nodes = append(available_nodes, node)
		}
	}

	if len(available_nodes) == 0 {
		// then there are no more paths
		// visited is the results (the path chosen in sequence)
		return current_path
	}

	var best_path *Path

	for _, destination := range available_nodes {

		this_time := time
		// we check each one - rather than have some heuristics, we just walk the path
		// th}e "winner" (which should be a complete path) is the best score
		subpath := current_path.Clone()
		_, steps := g.Graph2.getPath(source.ID, destination.ID)

		// steps is the sequence of steps required to get to this destination
		// so we will add each entry into our new subpath
		// fmt.Printf("source=%v, destination=%v, path=%v\n", source.ID, destination.ID, steps)
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
		if best_path == nil {
			best_path = result
		}
		if result.Score() > best_path.Score() {
			best_path = result
		}

		if this_time == MAX_TIME {
			break
		}

	}

	g.Cache.Put(source, best_path, time, best_path)

	return best_path

}

// // A>D opcost is 1, value is 28*20 = 560
//   // not A-I->J because opcost is 2, value is 27*21   567
//   // actually opcost is 4 because this path is closed so we have to pop the stack to get out
//   // not A->B because opcost is same move but lower return

//   look at dfs and bfs
//   dfs visits children before siblings

//   permutations

//   things we know at any given time
//   	all nodes (with total Value, their opportunity cost (travel+turn on))
// 	the routes from here to a scored node
// 		if the route contains scoring nodes, then effectively it is a route we could look at

// visited = stack()
// 	push(move_a)
// 	push(open_b)

// walk(A, visited)
// def walk(source, graph, stack):
// 	from SOURCE, list available path
// 	if available_paths == 0
// 		// then there are no more paths
// 		// visited is the results (the path chosen in sequence)
// 		return visited

// 	best_result = None
// 	best_score = 0
// 	for each path from SOURCE
// 		// we check each one
// 		// the "winner" (which should be a complete path) is the best score
// 		new_visited_paths = copy(visited_paths)
// 		walk to NODE
// 		open NODE
// 		assign NODE to results (time, value, node)
// 		result = walk(NODE, graph, new_visited)
// 		if result.Score() > best_score {
// 			best_result = result
// 		}

// 	}

// 	return best_result

// // func foo() {
// // 	g := LoadGraph(TEST_DATA)

// // 	// walk the graph and choose based on some sort of weird logic
// // 	aa := g.Get("AA")

// // 	// a,d,open-d,c,b,open-b,a,i,j,open-j,i,a,d,e,f,g,h,open-h,g,f,e,open-e,d,c,open-c

// // 	// get list of scored nodes
// // 		B(13)-1
// // 		D(20)-1
// // 		J(21)-2
// // 		C(2)-2
// // 		E(3)-2
// // 		H(22)-5

// // 	// walk the tree, all permutations

// // }
