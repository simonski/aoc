package d16

import "fmt"

type PathEntry struct {
	Action string
	Node   *Node
}

func (pe *PathEntry) Clone() *PathEntry {
	return NewPathEntry(pe.Action, pe.Node)
}

func NewPathEntry(action string, node *Node) *PathEntry {
	return &PathEntry{Action: action, Node: node}
}

type Path struct {
	Actions []*PathEntry
	Nodes   map[string]*Node
}

func (p *Path) Key(g *Graph) string {
	scored_nodes := g.GetScoredNodes()
	key := ""
	for _, node := range scored_nodes {
		if p.Nodes[node.ID] != nil {
			key = fmt.Sprintf("%v1", key)
		} else {
			key = fmt.Sprintf("%v0", key)
		}
	}
	return key
}

func (p *Path) String() string {
	line := ""
	for index, action := range p.Actions {
		if index > 0 {
			if action.Action == "MOVE" {
				line = fmt.Sprintf("%v -> %v", line, action.Node.ID)
			} else {
				line = fmt.Sprintf("%v (%v)", line, action.Node.ID)
			}

		} else {
			line = fmt.Sprintf(action.Node.ID)
		}
	}
	line = fmt.Sprintf("%v (score=%v)", line, p.Score())
	line = fmt.Sprintf("%v, (actions=%v)", line, len(p.Actions))
	return line
}

func (p *Path) Move(node *Node) {
	pe := NewPathEntry("MOVE", node)
	p.Actions = append(p.Actions, pe)
}

func (p *Path) Open(node *Node) {
	pe := NewPathEntry("OPEN", node)
	p.Actions = append(p.Actions, pe)
	p.Nodes[node.ID] = node
}

func (p *Path) Score() int {
	total := 0
	time := 30
	for index := 0; index < len(p.Actions); index++ {
		time -= 1
		action := p.Actions[index]
		if action.Action == "OPEN" {
			value := time * action.Node.Value
			total += value
		}
	}
	return total
}

func (p *Path) Contains(node *Node) bool {
	return p.Nodes[node.ID] != nil
}

func (p *Path) Size() int {
	return len(p.Actions)
}

func NewPath() *Path {
	p := Path{Actions: make([]*PathEntry, 0), Nodes: make(map[string]*Node)}
	return &p
}

func (p *Path) Clone() *Path {
	p2 := NewPath()
	for _, e := range p.Actions {
		p2.Actions = append(p2.Actions, e.Clone())
	}

	for k, n := range p.Nodes {
		p2.Nodes[k] = n
	}
	return p2
}
