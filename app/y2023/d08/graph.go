package d8

import (
	"strings"
)

type Graph struct {
	Nodes map[string]*Node
}

func (g *Graph) FindA() []*Node {
	results := make([]*Node, 0)
	for _, node := range g.Nodes {
		if node.A {
			results = append(results, node)
		}
	}
	return results
}

func LoadGraph(input string) (*Graph, *Instruction) {

	lines := strings.Split(input, "\n")
	instruction := NewInstruction(lines[0])
	graph_input := lines[2:]

	g := Graph{}
	g.Nodes = make(map[string]*Node)
	for _, line := range graph_input {
		node := NewNode(line)
		g.Nodes[node.Name] = node
	}

	for _, node := range g.Nodes {
		left := g.Nodes[node.LeftName]
		node.Left = left
		right := g.Nodes[node.RightName]
		node.Right = right
	}

	return &g, instruction

}

func (g *Graph) FindZ(startingNode string, instructions *Instruction) (*Node, int) {
	node := g.Nodes[startingNode]
	count := 0
	for {
		count += 1
		if instructions.Next() == "L" {
			node = node.Left
		} else {
			node = node.Right
		}
		if node.Z {
			return node, count
		}
	}
}

type Node struct {
	Line      string
	Name      string
	Left      *Node
	Right     *Node
	LeftName  string
	RightName string
	Z         bool
	A         bool
}

func NewNode(line string) *Node {
	// MPX = (DHN, NCX)
	node := Node{}
	node.Line = line
	splits := strings.Split(line, " = ")
	node.Name = splits[0]
	node.A = node.Name[2:3] == "A"
	node.Z = node.Name[2:3] == "Z"

	splits = strings.Split(splits[1], ", ")
	node.LeftName = strings.ReplaceAll(splits[0], "(", "")
	node.RightName = strings.ReplaceAll(splits[1], ")", "")
	return &node
}
