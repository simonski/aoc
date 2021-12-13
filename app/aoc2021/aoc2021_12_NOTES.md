    start
    /   \
c--A-----b--d
    \   /
    end

Your goal is to find the number of distinct paths that start at start, end at end, and don't visit small caves more than once. 

There are two types of caves: big caves (written in uppercase, like A) and small caves (written in lowercase, like b). 

So, all paths you find should visit small caves at most once, and can visit big caves any number of times.

There are 10 paths

start,A,b,A,c,A,end
start,A,b,A,end
start,A,b,end
start,A,c,A,b,A,end
start,A,c,A,b,end
start,A,c,A,end
start,A,end
start,b,A,c,A,end
start,b,A,end
start,b,end

PATHFINDING PSEUDOCODE

path:
    node[]
    isComplete:
        return node[-].isEnd

node:
    canVisit():
        if isStart():
            return false
        elif isSmall() && Visited > 1:
            return false 
        else:
            return true



    start
    /   \
c--A-----b--d
    \   /
    end


run until END
paths = []

foreach child in start.children():
    path = (start+child)
    paths.add(path)
    walk(path, paths)

START[children: A, b]
PATH1=[start, A)
PATHS=[PATH1]
WALK(PATH, PATHS)


----------------- 2

WALK:    [start,A,c]
    NODE=PATH.NODE[-1]
    NODE.IsEnd?  NO
        return false, NODE = "c"
    NODE.IsSmall(O) && Node.CanVisit()? YES [visited == 1]
        return false, NODE = "A"

    NODE.Visited++
    Node.Visited = 2  ( no more visits allowed )

    # WALK

    FOR CHILD in NODE.CHILDREN  [ A ]
        CHILD1 = A
        CHILD1.canVisit? YES
        NEW_PATH = COPY_PATH   ( start, A, c, A )
        NEW_PATH.add(c)        ( start, A, c )
        PATHS.ADD(NEW_PATH)
        WALK(NEW_PATH, PATHS)


----------------- 1
WALK:    [start,A]
    NODE=PATH.NODE[-1]
    NODE.IsEnd?  NO
        false, NODE = "A"
    NODE.IsSmall(O) && Node.IsVisibted()? NO
        false, NODE = "A"

    NODE.Visited++
    # Node.Visited = 1

    # WALK

    FOR CHILD in NODE.CHILDREN  [ c, B, end ]:
        CHILD1 = c
        CHILD1.canVisit? YES
        NEW_PATH = COPY_PATH   ( start, A )
        NEW_PATH.add(c)        ( start, A, c )
        PATHS.ADD(NEW_PATH)
        WALK(NEW_PATH, PATHS)

       

def walk(path, paths):
    node = path[-1]
    if node.IsEnd():
        # then this path is done
        path.add(node)
        return
    else if node.IsSmall() && node.IsVisited():
        # then we have reached the end of this trip
        path.add(node)
        return

    # so we are NOT at the end
    # we are NOT in a small cave already visited

    node.Visited++

    # okay, lets look at any children and see what we can do
    for child in node.children:
        # this might be a new path to traverse
        if child.canVisit():
            # then we need a 'new' path at this point to start walking
            newpath = copypath(path)
            paths.add(newpath)
            walk(node, newpath, paths)
        

start->A->c->A->end


first node, start 
fo
