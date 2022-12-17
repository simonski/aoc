package d16

import (
	"fmt"
	"testing"
)

func Test_1(t *testing.T) {
	graph := LoadGraph(TEST_DATA)
	fmt.Printf("There are %v nodes.\n", len(graph))
	t.Fatalf("mm")
}

func Test_Relationships(t *testing.T) {
	graph := LoadGraph(TEST_DATA)
	nodeAA := graph["AA"]
	if nodeAA == nil {
		t.Fatalf("AA should not be nil.")
	}
	if nodeAA.Value != 0 {
		t.Fatalf("AA value should be 0, was %v\n", nodeAA.Value)
	}
	if len(nodeAA.Children) != 3 {
		t.Fatalf("AA chlidren shoudl be 3, was %v\n", len(nodeAA.Children))
	}
	nodeABC := graph["ABA"]
	if nodeABC != nil {
		t.Fatalf("ABC should be nil.")
	}

	nodeDD := graph["DD"]
	if nodeDD.Value != 20 {
		t.Fatalf("DD value should be 20, was %v\n", nodeDD.Value)
	}
	if len(nodeDD.Children) != 3 {
		t.Fatalf("DD chlidren shoudl be 3, was %v\n", len(nodeDD.Children))
	}

	nodeII := graph["II"]
	if nodeII.Value != 0 {
		t.Fatalf("II value should be 0, was %v\n", nodeII.Value)
	}
	if len(nodeII.Children) != 2 {
		t.Fatalf("II chlidren shoudl be 2, was %v\n", len(nodeII.Children))
	}

	nodeBB := graph["BB"]
	if nodeBB.Value != 13 {
		t.Fatalf("BB value should be 13, was %v\n", nodeBB.Value)
	}

	if nodeABC != nil {
		t.Fatalf("ABC should be nil.")
	}

	nodeCC := graph["CC"]
	if nodeCC.Value != 2 {
		t.Fatalf("CC value should be 2, was %v\n", nodeCC.Value)
	}

}
