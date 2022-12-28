package d19

import (
	"strconv"
	"strings"
)

/*
--- Day 05:  ---

*/

type RobotType int

const ORE RobotType = 0
const CLAY RobotType = 1
const OBSIDIAN RobotType = 2
const GEODE RobotType = 3

type Blueprint struct {
	id     int
	robots []*Robot
}

func LoadBlueprints(input string) []*Blueprint {
	results := make([]*Blueprint, 0)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		bp := NewBlueprint(line)
		results = append(results, bp)
	}
	return results
}

func NewBlueprint(input string) *Blueprint {
	// Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 4 ore. Each obsidian robot costs 4 ore and 17 clay. Each geode robot costs 4 ore and 16 obsidian.
	input = strings.ReplaceAll(input, "Blueprint ", "")
	input = strings.ReplaceAll(input, "Each ore robot costs", "")
	input = strings.ReplaceAll(input, "Each clay robot costs", "")
	input = strings.ReplaceAll(input, "Each obsidian robot costs", "")
	input = strings.ReplaceAll(input, "Each geode robot costs", "")
	splits := strings.Split(input, ":")

	blueprint_id, _ := strconv.Atoi(splits[0])
	robots_input := strings.Split(splits[1], ".")
	robots := make([]*Robot, 0)
	for index, robot_input := range robots_input {
		robot_input = strings.Trim(robot_input, " ")
		if index == int(ORE) {
			// Each ore robot costs "4 ore"
			robot_input = strings.ReplaceAll(robot_input, " ore", "") // N
			oreCost, _ := strconv.Atoi(robot_input)
			robot := &Robot{Type: ORE, OreCost: oreCost}
			robots = append(robots, robot)
		} else if index == int(CLAY) {
			// Each clay robot costs "4 ore"
			robot_input = strings.ReplaceAll(robot_input, " ore", "") // N
			oreCost, _ := strconv.Atoi(robot_input)
			robot := &Robot{Type: CLAY, OreCost: oreCost}
			robots = append(robots, robot)
		} else if index == int(OBSIDIAN) {
			// Each obsidian robot costs "4 ore and 17 clay"
			robot_input = strings.ReplaceAll(robot_input, " ore and", "") // N
			robot_input = strings.ReplaceAll(robot_input, " clay", "")    // N N
			splits = strings.Split(robot_input, " ")
			oreCost, _ := strconv.Atoi(splits[0])
			clayCost, _ := strconv.Atoi(splits[1])
			robot := &Robot{Type: OBSIDIAN, OreCost: oreCost, ClayCost: clayCost}
			robots = append(robots, robot)

		} else if index == int(GEODE) {
			// Each geode robot costs "4 ore and 16 obsidian"
			robot_input = strings.ReplaceAll(robot_input, " ore and", "")  // N
			robot_input = strings.ReplaceAll(robot_input, " obsidian", "") // N N
			splits = strings.Split(robot_input, " ")
			oreCost, _ := strconv.Atoi(splits[0])
			obsidianCost, _ := strconv.Atoi(splits[1])
			robot := &Robot{Type: GEODE, OreCost: oreCost, ObsidianCost: obsidianCost}
			robots = append(robots, robot)

		}
	}
	bp := Blueprint{id: blueprint_id, robots: robots}
	return &bp
}

type Robot struct {
	Type         RobotType
	OreCost      int
	ClayCost     int
	ObsidianCost int
}
