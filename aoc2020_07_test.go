package main

import (
	"fmt"
	"strings"
	"testing"
)

const DAY_07_TEST_INPUT = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.
`

func Test_AOC2020_07_BagParser(t *testing.T) {
	graph := NewBagGraphFromStrings(strings.Split(DAY_07_TEST_INPUT, "\n"))
	// graph.Debug()

	actual := graph.GetBagsThatCanContain("shiny gold")
	expected := 4
	if len(actual) != expected {
		fmt.Printf("\nWalked tree and received the following: %v\n\n", actual)
		fmt.Printf("\n\n")
		graph.Debug()
		fmt.Printf("\n\n")
		t.Errorf("Expected %v, got %v.\n", expected, len(actual))
	}

	fmt.Printf("\n\n")
	graph.Debug()
	fmt.Printf("\nWalked tree and received the following: %v\n\n", actual)
}

func Test_AOC2020_07_BagTotals(t *testing.T) {
	graph := NewBagGraphFromStrings(strings.Split(DAY_07_TEST_INPUT, "\n"))
	// graph.Debug()

	actual := graph.GetTotalBagsContainedBy("shiny gold")
	expected := 32
	if actual != expected {
		t.Errorf("\nWalked tree and received the following: %v\n\n", actual)
	}
	fmt.Printf("\nWalked tree expected %v and received the following: %v\n\n", expected, actual)

}
