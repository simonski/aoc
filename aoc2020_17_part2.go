package main

/*

--- Day 17: Conway Cubes ---
Part 2

For some reason, your simulated results don't match what the experimental energy source engineers expected. Apparently, the pocket dimension actually has four spatial dimensions, not three.

The pocket dimension contains an infinite 4-dimensional grid. At every integer 4-dimensional coordinate (x,y,z,w), there exists a single cube (really, a hypercube) which is still either active or inactive.

Each cube only ever considers its neighbors: any of the 80 other cubes where any of their coordinates differ by at most 1. For example, given the cube at x=1,y=2,z=3,w=4, its neighbors include the cube at x=2,y=2,z=3,w=3, the cube at x=0,y=2,z=3,w=4, and so on.

The initial state of the pocket dimension still consists of a small flat region of cubes. Furthermore, the same rules for cycle updating still apply: during each cycle, consider the number of active neighbors of each cube.

For example, consider the same initial state as in the example above. Even though the pocket dimension is 4-dimensional, this initial state represents a small 2-dimensional slice of it. (In particular, this initial state defines a 3x3x1x1 region of the 4-dimensional space.)

Simulating a few cycles from this initial state produces the following configurations, where the result of each cycle is shown layer-by-layer at each given z and w coordinate:
*/

import (
	"fmt"
	"strconv"
	"strings"

	utils "github.com/simonski/aoc/utils"
	goutils "github.com/simonski/goutils"
)

func AOC_2020_17_part2_attempt1(cli *goutils.CLI) {

	g := NewGrid4D(DAY_17_INPUT)
	g.Cycle()
	g.Cycle()
	g.Cycle()
	g.Cycle()
	g.Cycle()
	g.Cycle()
	fmt.Printf("Part 2 Active Count is %v\n", g.CountActiveTotal())
}

type Grid4D struct {
	data map[string]string
}

func NewGrid4D(input string) *Grid4D {

	data := make(map[string]string)
	g := Grid4D{data: data}

	lines := strings.Split(input, "\n")

	y := 0
	z := 0
	w := 0
	for _, line := range lines {
		for x, _ := range line {
			value := line[x : x+1]
			key := fmt.Sprintf("%v.%v.%v.%v", x, y, z, w)
			g.Set(key, string(value))
		}
		y++
	}

	return &g
}

func (g *Grid4D) Get(key string) string {
	result, exists := g.data[key]
	if exists {
		return result
	} else {
		return INACTIVE
	}
}

func (g *Grid4D) Set(key string, value string) {
	if value == INACTIVE {
		delete(g.data, key)
	} else {
		g.data[key] = value
	}
}

func (g *Grid4D) Neighbours(parentKey string) []string {
	x, y, z, w := g.ParseKey(parentKey)
	keys := make([]string, 0)
	// fmt.Printf("Neighbours of %v\n", key)
	for wpos := w - 1; wpos <= w+1; wpos++ {
		for zpos := z - 1; zpos <= z+1; zpos++ {
			for xpos := x - 1; xpos <= x+1; xpos++ {
				for ypos := y - 1; ypos <= y+1; ypos++ {
					if xpos == x && ypos == y && zpos == z && wpos == w {
						continue
					}
					key := fmt.Sprintf("%v.%v.%v.%v", xpos, ypos, zpos, wpos)
					keys = append(keys, key)
					// fmt.Printf("Neighbour of (%v) = %v\n", parentKey, key)
				}
			}
		}
	}
	return keys
}

func (g *Grid4D) ParseKey(key string) (int, int, int, int) {
	splits := strings.Split(key, ".")
	x, _ := strconv.Atoi(splits[0])
	y, _ := strconv.Atoi(splits[1])
	z, _ := strconv.Atoi(splits[2])
	w, _ := strconv.Atoi(splits[3])
	return x, y, z, w
}

func (g *Grid4D) CountActiveNeighbours(parentKey string) int {
	keys := g.Neighbours(parentKey)
	active := 0
	for _, key := range keys {
		value := g.Get(key)
		if value == ACTIVE {
			active++
		}
	}
	return active
}

func (g *Grid4D) CountActiveTotal() int {
	activeCount := 0
	for _, value := range g.data {
		if value == ACTIVE {
			activeCount++
		}
	}
	return activeCount
}

func (g *Grid4D) Cycle() {
	data := make(map[string]string)

	minp, maxp := g.Dimensions()
	for w := minp.W - 1; w <= maxp.W+1; w++ {
		for z := minp.Z - 1; z <= maxp.Z+1; z++ {
			for y := minp.Y - 1; y <= maxp.Y+1; y++ {
				for x := minp.X - 1; x <= maxp.X+1; x++ {
					key := fmt.Sprintf("%v.%v.%v.%v", x, y, z, w)
					originalValue := g.Get(key)
					// fmt.Printf("Cycle(); key=%v, state=%v\n", key, originalValue)
					newValue := originalValue
					activeNeighbours := g.CountActiveNeighbours(key)
					if originalValue == ACTIVE {

						if activeNeighbours == 2 || activeNeighbours == 3 {
							// remain active
							newValue = ACTIVE
						} else {
							newValue = INACTIVE
						}

					} else if activeNeighbours == 3 {
						newValue = ACTIVE
					} else {
						newValue = originalValue
					}
					if newValue == ACTIVE {
						data[key] = newValue
					}
				}
			}
		}
	}
	g.data = data
}

// Dimensions returns the min/max points that exist
func (g *Grid4D) Dimensions() (utils.Point4D, utils.Point4D) {
	minp := utils.Point4D{X: 10000, Y: 10000, Z: 10000, W: 10000}
	maxp := utils.Point4D{X: -10000, Y: -10000, Z: -10000, W: -10000}
	for key := range g.data {
		x, y, z, w := g.ParseKey(key)
		minp.X = utils.Min(x, minp.X)
		maxp.X = utils.Max(x, maxp.X)
		minp.Y = utils.Min(y, minp.Y)
		maxp.Y = utils.Max(y, maxp.Y)
		minp.Z = utils.Min(z, minp.Z)
		maxp.Z = utils.Max(z, maxp.Z)
		minp.W = utils.Min(w, minp.W)
		maxp.W = utils.Max(w, maxp.W)
	}
	return minp, maxp
}

func (g *Grid4D) DebugZ(z int) string {
	minp, maxp := g.Dimensions()
	min_x := minp.X
	max_x := maxp.X
	min_y := minp.Y
	max_y := maxp.Y
	// min_w := minp.w
	// max_w := maxp.w
	line := ""
	for y := min_y; y <= max_y; y++ {
		for x := min_x; x <= max_x; x++ {
			key := fmt.Sprintf("%v.%v.%v", x, y, z)
			value := g.Get(key)
			line += value
		}
		line += "\n"
	}
	return line
}
