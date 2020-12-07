package main

/*
Day 07 - Handy Haversacks
https://adventofcode.com/2020/day/7

Sample input

dim red bags contain 2 bright gold bags, 5 striped fuchsia bags.
dotted purple bags contain 5 bright olive bags, 3 faded maroon bags.
plaid chartreuse bags contain 1 vibrant olive bag, 5 bright black bags, 1 clear tomato bag.
wavy orange bags contain 4 dark lavender bags, 4 posh white bags.
light lavender bags contain 4 drab olive bags, 5 dark magenta bags.

Given: a shiny gold bag, how many bag colors can eventually contain at least one shiny gold bag?


etc.

Given

light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.

Then

In the above rules, the following options would be available to you:

A bright white bag, which can hold your shiny gold bag directly.
A muted yellow bag, which can hold your shiny gold bag directly, plus some other bags.
A dark orange bag, which can hold bright white and muted yellow bags, either of which could then hold your shiny gold bag.
A light red bag, which can hold bright white and muted yellow bags, either of which could then hold your shiny gold bag.
So, in this example, the number of bag colors that can eventually contain at least one shiny gold bag is 4.


*/
import (
	"fmt"
	"strings"

	goutils "github.com/simonski/goutils"
)

// AOC_2020_07 is the entrypoint to the various attempts for day six
func AOC_2020_07(cli *goutils.CLI) {
	AOC_2020_07_part1_attempt1(cli)

}

func AOC_2020_07_part1_attempt1(cli *goutils.CLI) {
	filename := cli.GetFileExistsOrDie("-input")
	g := NewBagGraphFromFilename(filename)
	g.Debug()
	fmt.Printf("There are %v possible combinations.\n", len(g.GetBagsThatCanContain("shiny gold")))
}

func NewBagGraphFromFilename(filename string) *BagGraph {
	lines := load_file_to_strings(filename)
	return NewBagGraphFromStrings(lines)
}

func NewBagGraphFromStrings(lines []string) *BagGraph {
	graph := &BagGraph{bags: make(map[string]*Bag)}
	for index := range lines {
		line := lines[index]

		// mirrored beige bags contain no other bags.
		// dotted silver bags contain 1 vibrant green bag.
		// light brown bags contain 1 shiny silver bag, 3 plaid olive bags, 1 clear tan bag.				// original
		// light brown  contain 1 shiny silver bag, 3 plaid olive , 1 clear tan bag 				        // remove bags
		// light brown  contain 1 shiny silver , 3 plaid olive , 1 clear tan  				                // remove bag

		// fmt.Printf("%v\n", line)
		if strings.TrimSpace(line) == "" {
			// nothing
		} else if strings.HasSuffix(line, "no other bags.") {
			// no children here
			splits := strings.Split(line, "contain")
			colour := strings.TrimSpace(strings.ReplaceAll(splits[0], "bags", ""))
			graph.GetOrCreate(colour) // add this in
		} else {
			// it has children

			line = strings.ReplaceAll(line, "bags", "")
			line = strings.ReplaceAll(line, "bag", "")
			line = strings.ReplaceAll(line, ".", "")
			splits := strings.Split(line, "contain") // [ "light brown", "1 shiny silver, 3 plain olive, 1 clear tan"]
			colour := strings.TrimSpace(splits[0])
			bag := graph.GetOrCreate(colour)
			children := strings.Split(splits[1], ",") // [ "1 shiny silver", "3 plain olive", "1 clear tan" ]
			for childIndex := range children {
				child := strings.Split(strings.TrimSpace(children[childIndex]), " ")
				// childCount, _ := strconv.Atoi(child[0])
				childColour := strings.Join(child[1:], " ")
				childBag := graph.GetOrCreate(childColour)
				bag.AddChild(childBag)
			}

		}
	}
	return graph
}

type BagGraph struct {
	bags map[string]*Bag
}

func (graph *BagGraph) Debug() {
	for key := range graph.bags {
		fmt.Printf("%v\n", key)
		bag := graph.GetOrCreate(key)
		for index := range bag.Children {
			child := bag.Children[index]
			fmt.Printf("  %v\n", child.Colour)
		}
		fmt.Printf("\n")
	}
}

func (graph *BagGraph) GetOrCreate(colour string) *Bag {
	value, exists := graph.bags[colour]
	if exists {
		return value
	}
	b := NewBag(colour)
	graph.bags[colour] = b
	return b
}

func (b *BagGraph) GetBagsThatCanContain(colour string) map[string]*Bag {
	bag := b.GetOrCreate(colour)
	// so walking 'up' any parent, adding each parent we have to the map
	p := make(map[string]*Bag)
	walkBag(bag, p)
	return p
}

func walkBag(bag *Bag, results map[string]*Bag) {
	_, exists := results[bag.Colour]
	if exists {
		return
	}
	for index := range bag.Parents {
		entry := bag.Parents[index]
		walkBag(entry, results)
		results[entry.Colour] = entry
	}

	// now walk 'up' the tree for each parent until it is empty and we have the total for that

}

// Bag is my own impl of a simple Tree I can walk later
type Bag struct {
	Colour   string
	Children []*Bag
	Parents  []*Bag
}

func (b *Bag) AddChild(child *Bag) {
	b.Children = append(b.Children, child)
	child.Parents = append(child.Parents, b)
}

func NewBag(line string) *Bag {
	children := make([]*Bag, 0)
	parents := make([]*Bag, 0)
	b := &Bag{Colour: line, Children: children, Parents: parents}
	return b
}
