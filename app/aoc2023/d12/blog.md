Puzzle: Day 12: Hot Springs
Year: 2023
Day: 12

P1 you can do combinations, so  I did.
P2, you can't. So it's a graph search space.  There are some complicating factors.

I think it is a depth first search,  I think at each step I can work out if it has
passed a regex so far.  If it has, continue, if it has not, stop.  I think.

I think each choice at each step will create a node

P2: OK some sort of recursive DFS where at each leaf
    the criteria of the leaf should "know" 
        - position from top
        - how much damage
        - sequence of damage so far as being rule-passing
        - remaining damage
        - remaining positions to depth
    so some sort of backtracking where I mark a leaf as "you shall not pass"

DFS with backtracking to get out when the conditoin explicitly fails.

as I drop down a bitset might work on the damage count
the rules make it hard but I can pass on a true/false style
the point at which a single rule passes in sequence I should be able to mark in the information being carried
the point at which a rule can no longer pass I should be able to terminate the tree at that point.

recursive DFS from wikipedia 

procedure DFS(G, v) is
    label v as discovered
    for all directed edges from v to w that are in G.adjacentEdges(v) do
        if vertex w is not labeled as discovered then
            recursively call DFS(G, w)



iterative DFS from wikipedia

procedure DFS_iterative(G, v) is
    let S be a stack
    S.push(v)
    while S is not empty do
        v = S.pop()
        if v is not labeled as discovered then
            label v as discovered
            for all edges from v to w in G.adjacentEdges(v) do 
                S.push(w)

iterative DFS v2 from wikipedia

procedure DFS_iterative(G, v) is
    let S be a stack
    label v as discovered
    S.push(iterator of G.adjacentEdges(v))
    while S is not empty do
        if S.peek().hasNext() then
            w = S.peek().next()
            if w is not labeled as discovered then
                label w as discovered
                S.push(iterator of G.adjacentEdges(w))
        else
            S.pop()                