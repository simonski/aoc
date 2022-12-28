package d19

import (
	"strings"
	"testing"
)

func Test_LoadBluePrint1(t *testing.T) {
	inputs := strings.Split(TEST_DATA, "\n")

	bp := NewBlueprint(inputs[0])
	if bp.id != 1 {
		t.Fatalf("BP id should be 1, was %v\n", bp.id)

	}
	if len(bp.robots) != 4 {
		t.Fatalf("There should be 4 robots, there were %v\n", len(bp.robots))
	}

	robot := bp.robots[ORE]
	if robot.Type != ORE {
		t.Fatalf("1st robot should be type ORE, was %v\n", robot.Type)
	}
	if robot.OreCost != 4 || robot.ClayCost != 0 || robot.ObsidianCost != 0 {
		t.Fatalf("1st robot OreCost/ClayCost/ObsidianCost should be type 4/0/0, was %v/%v/%v\n", robot.OreCost, robot.ClayCost, robot.ObsidianCost)
	}

	robot = bp.robots[CLAY]
	if robot.Type != CLAY {
		t.Fatalf("2nd robot should be type CLAY, was %v\n", robot.Type)
	}
	if robot.OreCost != 2 || robot.ClayCost != 0 || robot.ObsidianCost != 0 {
		t.Fatalf("2nd robot OreCost/ClayCost/ObsidianCost should be type 2/0/0, was %v/%v/%v\n", robot.OreCost, robot.ClayCost, robot.ObsidianCost)
	}

	robot = bp.robots[OBSIDIAN]
	if robot.Type != OBSIDIAN {
		t.Fatalf("3rd robot should be type OBSIDIAN, was %v\n", robot.Type)
	}
	if robot.OreCost != 3 || robot.ClayCost != 14 || robot.ObsidianCost != 0 {
		t.Fatalf("3rd robot OreCost/ClayCost/ObsidianCost should be type 3/14/0, was %v/%v/%v\n", robot.OreCost, robot.ClayCost, robot.ObsidianCost)
	}

	robot = bp.robots[GEODE]
	if robot.Type != GEODE {
		t.Fatalf("4th robot should be type GEODE, was %v\n", robot.Type)
	}
	if robot.OreCost != 2 || robot.ClayCost != 0 || robot.ObsidianCost != 7 {
		t.Fatalf("4th robot OreCost/ClayCost/ObsidianCost should be type 2/0/7, was %v/%v/%v\n", robot.OreCost, robot.ClayCost, robot.ObsidianCost)
	}

	bp2 := NewBlueprint(inputs[1])
	if bp2.id != 2 {
		t.Fatalf("BP2 id should be 2, was %v\n", bp.id)

	}
	if len(bp2.robots) != 4 {
		t.Fatalf("There should be 4 robots, there were %v\n", len(bp.robots))
	}

}

func Test_LoadBluePrint2(t *testing.T) {
	inputs := strings.Split(TEST_DATA, "\n")

	bp := NewBlueprint(inputs[1])
	if bp.id != 2 {
		t.Fatalf("BP id should be 2, was %v\n", bp.id)

	}
	if len(bp.robots) != 4 {
		t.Fatalf("There should be 4 robots, there were %v\n", len(bp.robots))
	}

	robot := bp.robots[ORE]
	if robot.Type != ORE {
		t.Fatalf("1st robot should be type ORE, was %v\n", robot.Type)
	}
	if robot.OreCost != 2 || robot.ClayCost != 0 || robot.ObsidianCost != 0 {
		t.Fatalf("1st robot OreCost/ClayCost/ObsidianCost should be type 2/0/0, was %v/%v/%v\n", robot.OreCost, robot.ClayCost, robot.ObsidianCost)
	}

	robot = bp.robots[CLAY]
	if robot.Type != CLAY {
		t.Fatalf("2nd robot should be type CLAY, was %v\n", robot.Type)
	}
	if robot.OreCost != 3 || robot.ClayCost != 0 || robot.ObsidianCost != 0 {
		t.Fatalf("2nd robot OreCost/ClayCost/ObsidianCost should be type 3/0/0, was %v/%v/%v\n", robot.OreCost, robot.ClayCost, robot.ObsidianCost)
	}

	robot = bp.robots[OBSIDIAN]
	if robot.Type != OBSIDIAN {
		t.Fatalf("3rd robot should be type OBSIDIAN, was %v\n", robot.Type)
	}
	if robot.OreCost != 3 || robot.ClayCost != 8 || robot.ObsidianCost != 0 {
		t.Fatalf("3rd robot OreCost/ClayCost/ObsidianCost should be type 3/8/0, was %v/%v/%v\n", robot.OreCost, robot.ClayCost, robot.ObsidianCost)
	}

	robot = bp.robots[GEODE]
	if robot.Type != GEODE {
		t.Fatalf("4th robot should be type GEODE, was %v\n", robot.Type)
	}
	if robot.OreCost != 3 || robot.ClayCost != 0 || robot.ObsidianCost != 12 {
		t.Fatalf("4th robot OreCost/ClayCost/ObsidianCost should be type 3/0/12, was %v/%v/%v\n", robot.OreCost, robot.ClayCost, robot.ObsidianCost)
	}

	bp2 := NewBlueprint(inputs[1])
	if bp2.id != 2 {
		t.Fatalf("BP2 id should be 2, was %v\n", bp.id)

	}
	if len(bp2.robots) != 4 {
		t.Fatalf("There should be 4 robots, there were %v\n", len(bp.robots))
	}

}

func Test_LoadBluePrints(t *testing.T) {
	bps := LoadBlueprints(TEST_DATA)
	if len(bps) != 2 {
		t.Fatalf("There shoudl be test 2 blueprints, there were %v\n", len(bps))
	}

	bps_real := LoadBlueprints(REAL_DATA)
	if len(bps_real) != 30 {
		t.Fatalf("There shoudl be test 30 blueprints, there were %v\n", len(bps_real))
	}
}

func Test_All(t *testing.T) {
	Test_LoadBluePrint1(t)
	Test_LoadBluePrint2(t)
	Test_LoadBluePrints(t)
}
