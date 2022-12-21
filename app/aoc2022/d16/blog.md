Puzzle: Proboscidea Volcanium
Year: 2022
Day: 16

---
Saturday 17th 0930

Before I read it, I want to say to myself - I am enjoying this year much more.  The first ten(?) days were I think toy problems to stretch the fingers I suppose.  I could have used those to learn some rust or, - well, whateger.

As we got past maybe D12, it got more involved - the demon on my shoulder telling me I'm no good was beginning to be fed, so I had a word with myself and said: why are you doing this? it's supposed to be fun, not a pressured win/lose.  

This means that the days 14/15 which took > 24hours, I do still mind and feel a failure, but I don't mind so much.  so what that it took my more than eleventy seconds. The point is I aimed to solve them and I did. I think Cryptonomicon has a mention of Newton saying to Enoch something along the lines of - "what matters is that you did understand the calculus, not that you didn't discover it."   So, I could be Enoch and happy.  At least - was that Enoch? or maybe it was the other guy.  Time to reread Cryptonomicon - it is Christmas break after all (in a week).

Also - blogging (or writing) might be a good thing.

-----

Saturday 0943 

1 coffee later and a read of the problem.

I like this problem.  I'm going to attempt to blog my thoughts along the way - I haven't written the blogging engine but why shoudl that stop me from writing the content - no software required for that :).   It's' *really* cold outside minus 4 I think.  

Ok so I think this problem is a tree, the paths are a set of values.  I think I 
- build the tree
- walk the paths to find the "best". 

So if my aim is to 1. solve and 2. *undertand*, then I think that I should do it in two parts

1. write a tree (yes - I write it all from scratch - this is actuallt the *point*)
2. write a tree walker 

I can do 1. mostly.  I cannot remember tree walking properly, so I will consult Donald *before* I attempt the walking.

---

A tree is a graph.  

So we have a nubmer of things

## Recursion

The base case is around time.

The Score - given a stack of actions, move, open, we can look at the stack depth - or the path in the graph, and calculate the cost.

The other thing is the path is a tree; you can calculate the score of a path of N moves if we walk the path fully to the max depth; some nodes we should be able to calcualte the maximum score and exclude them - e.g anything with a zero score, we don't walk.

Ok so again I think I tried brute forcing it by accident.  So now I look and say

some are scored.   Build all combos of sequence then attempt to walk the path to each one

- build the list of scored entries, sort descending as preference
generate the set of all combinations of paths
- for each entry, build the path that gets you there
- walk from each to the next
- once all scored are touched or time is at max, quit




- there are paths that just can't be correct - nonscored paths.
- so we prefer to land scored paths quickly
- so a sorted scored path, then the variations between, then traverse the route from 
    scored vertex1 to scored vertex2

Ok I've dijkstra'ed up my graph. Next up I will build the paths between each scored node and see
what the best paths are.

Each path could be ranked based on its proximity to other scored values  E.g. "knowing" a score is X moves away
gives me its value.   at that point I can build a tree of potential scores.


-----

Sunday 10:10 - a tree is a graph, but this graph is a tree!. It's a depth-first walk of a tree, that's all.

So.. I'll from-scratch rewrite a tree, retain my graph with my Dijkstra and attempt to learn to *read the problem first*.

-----

Sunday 10:45: It's not a tree.  I did read it, it's a graph.  I think I was maybe correct on my djikstra.  But... perhaps the notes on the walking of the graph hve some hints.   So, perhaps my walker needs a stack I can inspect for loops.

Sunday 15:43 Ok am back and am going to try something like

    from current node,
    1. build list of paths that are a) unexplored and b) score > 0 
    2. shortest path to each scorable
    3. "score" each path given a time cost and time position 
        cost=cost to traverse and switch on 
        score=total cost from a time value X down to 0 this path gives

    4. repeat the above for the remaining paths "from" the new node position


-----

Wed 18:02: Part 1 done. I had this tension between "I can't brute force it" and then got to caching fragments.  
The annoying thing was I had the TEST_DATA loading into the graph, so I probably spent and hour running and being
confused that it wasn't working.