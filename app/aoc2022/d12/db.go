package d12

import (
	"fmt"
	"strings"
)

type DB struct {
	NodesXY     map[string]*Node
	NodesID     map[int]*Node
	Start       *Node
	Destination *Node
	Cols        int
	Rows        int
	BestPath    []int
	Iteration   int
	DeadNodes   map[int]bool
}

func NewDB(data string) *DB {
	rows := strings.Split(data, "\n")
	db := &DB{NodesXY: make(map[string]*Node), NodesID: make(map[int]*Node)}
	width := 0
	height := 0
	id := 0
	for y, row := range rows {
		height += 1
		width = 0
		for x := 0; x < len(row); x++ {
			id += 1
			value := row[x : x+1] //, _ := strconv.Atoi(row[colIndex : colIndex+1])
			node := NewNode(id, value, x, y)
			if node.IsStart {
				db.Start = node
			}
			if node.IsDestination {
				db.Destination = node
			}
			db.Add(node)
			width++
		}
	}
	db.Rows = height
	db.Cols = width
	db.BestPath = make([]int, 0)
	db.DeadNodes = make(map[int]bool)
	return db
}

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

// println(color.White + "This is White" + color.Reset)

func (db *DB) Debug(showDot bool, path *Node) string {
	result := ""
	for y := 0; y < db.Rows; y++ {
		line := ""
		for x := 0; x < db.Cols; x++ {
			node := db.GetByXY(x, y)
			// if this node in grid is in the path, mark with an X
			if node.ID == path.Path[len(path.Path)-1] {
				line = fmt.Sprintf("%v%v", line, "*")
			} else if path.ContainsID(node, db) {
				line = fmt.Sprintf("%v%v", line, "x")
			} else {
				if showDot {
					line = fmt.Sprintf("%v%v", line, ".")
				} else {
					line = fmt.Sprintf("%v%v", line, node.Letter)
				}
			}
		}
		if y == 0 {
			result = fmt.Sprintf("%v", line)
		} else {
			result = fmt.Sprintf("%v\n%v", result, line)
		}
	}
	return result
}

func (db *DB) GetByXY(x int, y int) *Node {
	key := fmt.Sprintf("%v_%v", x, y)
	return db.NodesXY[key]
}

func (db *DB) GetByID(id int) *Node {
	return db.NodesID[id]
}

func (db *DB) Add(node *Node) {
	key := fmt.Sprintf("%v_%v", node.X, node.Y)
	db.NodesXY[key] = node
	db.NodesID[node.ID] = node
}

func (db *DB) AddDeadNode(node *Node) {
	db.DeadNodes[node.ID] = true
}

func (db *DB) IsDeadNode(node *Node) bool {
	return db.DeadNodes[node.ID]
}

// return N S E W neighbours of this node
// excludes visited or too-high neighbours
func (db *DB) Neighbours(node *Node) []*Node {
	nodes := make([]*Node, 0)
	up := db.GetByXY(node.X, node.Y-1)
	down := db.GetByXY(node.X, node.Y+1)
	left := db.GetByXY(node.X-1, node.Y)
	right := db.GetByXY(node.X+1, node.Y)
	if right != nil {
		if !db.IsDeadNode(right) && right.Value == node.Value || right.Value == node.Value+1 {
			// if right.Value-1 <= node.Value {
			nodes = append(nodes, right)
		} else {
			// fmt.Printf("neigbours - not adding right (source is %v, right is %v)\n", node.Letter, right.Letter)
		}
	}
	if up != nil {
		if !db.IsDeadNode(up) && up.Value == node.Value || up.Value == node.Value+1 {
			// if up.Value-1 <= node.Value {
			nodes = append(nodes, up)
		} else {
			// fmt.Printf("neigbours - not adding up (source is %v, up is %v)\n", node.Letter, up.Letter)
		}
	}
	if down != nil {
		if !db.IsDeadNode(down) && down.Value == node.Value || down.Value == node.Value+1 {
			// if down.Value-1 <= node.Value {
			nodes = append(nodes, down)
		} else {
			// fmt.Printf("neigbours - not adding down (source is %v, down is %v)\n", node.Letter, down.Letter)
		}
	}
	if left != nil {
		if !db.IsDeadNode(right) && left.Value == node.Value || left.Value == node.Value+1 {
			// if left.Value-1 <= node.Value {
			nodes = append(nodes, left)
		} else {
			// fmt.Printf("neigbours - not adding left (source is %v, left is %v)\n", node.Letter, left.Letter)
		}
	}
	// return nodes
	// // sort in preference of.. ascend, go right
	// sort.Slice(nodes, func(i int, j int) bool {

	// 	/// less func
	// 	a := nodes[i]
	// 	b := nodes[j]

	// 	if b.Value+1 == a.Value {
	// 		// higher
	// 		return true
	// 	}
	// 	return false
	// })

	return nodes
}

func (db *DB) GetParent(node *Node) *Node {
	index := len(node.Path) - 1
	id := node.Path[index]
	return db.NodesID[id]
}

type Node struct {
	ID             int
	X              int
	Y              int
	Letter         string
	Value          int
	IsStart        bool
	IsDestination  bool
	VisitDirection string
	Path           []int
}

func NewNode(id int, letter string, x int, y int) *Node {
	node := &Node{ID: id, Letter: letter, X: x, Y: y}
	runeLetter := letter[0]
	if letter == "S" {
		runeLetter = "a"[0]
	} else if letter == "E" {
		runeLetter = "z"[0]
	}
	node.IsStart = letter == "S"
	node.IsDestination = letter == "E"
	node.Letter = letter
	char := rune(runeLetter)
	ascii := int(char)
	node.Value = ascii
	node.VisitDirection = "."
	node.Path = make([]int, 0)
	// node.Path = append(node.Path, node.ID)
	return node
}

func (n *Node) String() string {
	return fmt.Sprintf("(%v,%v)[%v]", n.X, n.Y, n.Letter)
}

func (n *Node) Clone(path []int) *Node {
	c := &Node{}
	c.X = n.X
	c.Y = n.Y
	c.ID = n.ID
	c.Value = n.Value
	c.Letter = n.Letter
	c.VisitDirection = n.VisitDirection
	c.Path = make([]int, 0)
	for _, value := range path {
		c.Path = append(c.Path, value)
	}
	return c
}

// func (n *Node) SetPath(otherPath []int) {
// 	n.Path = make([]int, 0)
// 	n.Path = append(n.Path, n.ID)
// 	for _, id := range otherPath {
// 		n.Path = append(n.Path, id)
// 	}
// }

func (n *Node) Equals(other *Node) bool {
	return n.X == other.X && n.Y == other.Y
}

func (n *Node) ContainsID(other *Node, db *DB) bool {
	// do anty node children equal this candidate
	// is it in the path (the sequence of parent ids)
	for _, id := range n.Path {
		if id == other.ID {
			return true
		}
	}
	return false
}

func (n *Node) Size(db *DB) int {
	return len(n.Path)
}

func (n *Node) NicePath(db *DB) string {
	nicePath := ""
	for index, id := range n.Path {
		n := db.GetByID(id)
		if index == 0 {
			nicePath = fmt.Sprintf("%v", n)
		} else {
			nicePath = fmt.Sprintf("%v <- %v ", nicePath, n)
		}
	}
	return nicePath
}

func (db *DB) Walk(mod int, verbose bool, showDot bool, debug bool) {
	node := db.Start
	db.walk(mod, verbose, showDot, debug, node)
}

func (db *DB) walk(mod int, VERBOSE bool, showDot, debug bool, node *Node) {
	node.Path = append(node.Path, node.ID)
	if VERBOSE {
		fmt.Printf("walk[%v] %v : the path size is %v, path is %v\n", db.Iteration, node, len(node.Path), node.NicePath(db))
	}

	db.Iteration++
	if debug && db.Iteration%mod == 0 {
		fmt.Printf("Iteration %v\n", db.Iteration)
		fmt.Println(db.Debug(showDot, node))
		fmt.Println("")
	}

	// get the N S W E
	if node.IsDestination {
		if len(db.BestPath) == 0 {
			// this path is the first, use it
			if VERBOSE {
				fmt.Println("walk: have a destination, BestPath is nill, automatically set this one.")
			}
			db.BestPath = node.Path
			return
		} else if len(db.BestPath) > len(node.Path) {
			// this path is the best, use it
			if VERBOSE {
				fmt.Println("walk: new path is smaller")
			}
			db.BestPath = node.Path
			return
		} else if len(node.Path) > len(db.BestPath) {
			// this path is too big, exit
			if VERBOSE {
				fmt.Println("walk: this path is too big")
			}
			return
		}
	} else if len(db.BestPath) > 0 {
		if len(node.Path) > len(db.BestPath) {
			// this path is too big
			return
		}
	}

	// nicePath := node.NicePath(db)

	// current position cX, cY
	// a path is a box fo ra given neighbour if the choice means you cannot get out

	// A box is four sides walled in
	// 	so is there a line x[a...n...b] where N is >= a and <= b
	// 	is there a line where y[a...b..b] where M >= a and <= b

	// }

	neighbours := db.Neighbours(node)
	if len(neighbours) == 0 {
		db.AddDeadNode(node)
		if VERBOSE {
			fmt.Printf("walk[%v] %v has no unused neighbours - time to unwind.", db.Iteration, node)
		}
		return
	}

	// if VERBOSE {
	// fmt.Printf("walk[%v] %v has %v neighbours %v, current path len '%v'\n", db.Iteration, node, len(neighbours), neighbours, len(node.Path)) // nicePath
	// }
	for _, neighbour := range neighbours {
		if !node.ContainsID(neighbour, db) {
			// if neighbour.Value-1 <= node.Value {
			if VERBOSE {
				fmt.Printf("walk[%v] %v : about moving to %v, not in path, height OK (%v->%v)\n", db.Iteration, node, neighbour, node.Letter, neighbour.Letter)
			}
			clonedNeighbour := neighbour.Clone(node.Path)
			if VERBOSE {
				fmt.Printf("walk[%v] %v : cloned the neighbour %v, about to walk, the neighbour path is %v (%v)\n", db.Iteration, node, clonedNeighbour, len(clonedNeighbour.Path), clonedNeighbour.NicePath(db))
			}
			// fmt.Printf("walk[%v] %v : added the curent path to the neighbour %v we cloned, the new path is %v\n", db.Iteration, node, clonedNeighbour, clonedNeighbour.NicePath(db))
			db.walk(mod, VERBOSE, showDot, debug, clonedNeighbour)
			// } else {
			// 	// ascent too big
			// 	if VERBOSE {
			// 		fmt.Printf("walk[%v] %v : canot move to %v, ascent too big (%v->%v)\n", db.Iteration, node, neighbour, node.Letter, neighbour.Letter)
			// 	}
			// }
		} else {
			if VERBOSE {
				fmt.Printf("walk[%v] %v : canot move to %v, already in path\n", db.Iteration, node, neighbour)
			}
		}
	}

}
