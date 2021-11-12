package app

/*

--- Day 17: Conway Cubes ---
As your flight slowly drifts through the sky, the Elves at the Mythical Information Bureau at the North Pole contact you. They'd like some help debugging a malfunctioning experimental energy source aboard one of their super-secret imaging satellites.

The experimental energy source is based on cutting-edge technology: a set of Conway Cubes contained in a pocket dimension! When you hear it's having problems, you can't help but agree to take a look.

The pocket dimension contains an infinite 3-dimensional grid. At every integer 3-dimensional coordinate (x,y,z), there exists a single cube which is either active or inactive.

In the initial state of the pocket dimension, almost all cubes start inactive. The only exception to this is a small flat region of cubes (your puzzle input); the cubes in this region start in the specified active (#) or inactive (.) state.

The energy source then proceeds to boot up by executing six cycles.

Each cube only ever considers its neighbors: any of the 26 other cubes where any of their coordinates differ by at most 1. For example, given the cube at x=1,y=2,z=3, its neighbors include the cube at x=2,y=2,z=2, the cube at x=0,y=2,z=3, and so on.

During a cycle, all cubes simultaneously change their state according to the following rules:

If a cube is active and exactly 2 or 3 of its neighbors are also active, the cube remains active. Otherwise, the cube becomes inactive.
If a cube is inactive but exactly 3 of its neighbors are active, the cube becomes active. Otherwise, the cube remains inactive.
The engineers responsible for this experimental energy source would like you to simulate the pocket dimension and determine what the configuration of cubes should be at the end of the six-cycle boot process.

For example, consider the following initial state:

.#.
..#
###
Even though the pocket dimension is 3-dimensional, this initial state represents a small 2-dimensional slice of it. (In particular, this initial state defines a 3x3x1 region of the 3-dimensional space.)

Simulating a few cycles from this initial state produces the following configurations, where the result of each cycle is shown layer-by-layer at each given z coordinate (and the frame of view follows the active cells in each cycle):

Before any cycles:

z=0
.#.
..#
###


After 1 cycle:

z=-1
#..
..#
.#.

z=0
#.#
.##
.#.

z=1
#..
..#
.#.


After 2 cycles:

z=-2
.....
.....
..#..
.....
.....

z=-1
..#..
.#..#
....#
.#...
.....

z=0
##...
##...
#....
....#
.###.

z=1
..#..
.#..#
....#
.#...
.....

z=2
.....
.....
..#..
.....
.....


After 3 cycles:

z=-2
.......
.......
..##...
..###..
.......
.......
.......

z=-1
..#....
...#...
#......
.....##
.#...#.
..#.#..
...#...

z=0
...#...
.......
#......
.......
.....##
.##.#..
...#...

z=1
..#....
...#...
#......
.....##
.#...#.
..#.#..
...#...

z=2
.......
.......
..##...
..###..
.......
.......
.......
After the full six-cycle boot process completes, 112 cubes are left in the active state.

Starting with your given initial configuration, simulate six cycles. How many cubes are left in the active state after the sixth cycle?
*/

import (
	"fmt"
	"strconv"
	"strings"

	goutils "github.com/simonski/goutils"
)

const DAY_17_INPUT = `.#.####.
.#...##.
..###.##
#..#.#.#
#..#....
#.####..
##.##..#
#.#.#..#`

const DAY_17_TEST_INPUT = `.#.
..#
###`

const ACTIVE = "#"
const INACTIVE = "."

// AOC_2020_17 is the entrypoint
func AOC_2020_17(cli *goutils.CLI) {
	AOC_2020_17_part1_attempt1(cli)
	AOC_2020_17_part2_attempt1(cli)
}

func AOC_2020_17_part1_attempt1(cli *goutils.CLI) {

	g := NewGrid3D(DAY_17_INPUT)
	g.Cycle()
	g.Cycle()
	g.Cycle()
	g.Cycle()
	g.Cycle()
	g.Cycle()
	fmt.Printf("Part 1 Active Count is %v\n", g.CountActiveTotal())
}

type Grid3D struct {
	data map[string]string
}

func NewGrid3D(input string) *Grid3D {

	data := make(map[string]string)
	g := Grid3D{data: data}

	lines := strings.Split(input, "\n")

	y := 0
	z := 0
	for _, line := range lines {
		for x, _ := range line {
			value := line[x : x+1]
			key := fmt.Sprintf("%v.%v.%v", x, y, z)
			g.Set(key, string(value))
		}
		y++
	}

	return &g
}

func (g *Grid3D) Get(key string) string {
	result, exists := g.data[key]
	if exists {
		return result
	} else {
		return INACTIVE
	}
}

func (g *Grid3D) Set(key string, value string) {
	if value == INACTIVE {
		delete(g.data, key)
	} else {
		g.data[key] = value
	}
}

func (g *Grid3D) Neighbours(parentKey string) []string {
	x, y, z := g.ParseKey(parentKey)
	keys := make([]string, 0)
	// fmt.Printf("Neighbours of %v\n", key)
	for zpos := z - 1; zpos <= z+1; zpos++ {
		for xpos := x - 1; xpos <= x+1; xpos++ {
			for ypos := y - 1; ypos <= y+1; ypos++ {
				if xpos == x && ypos == y && zpos == z {
					continue
				}
				key := fmt.Sprintf("%v.%v.%v", xpos, ypos, zpos)
				keys = append(keys, key)
				// fmt.Printf("Neighbour of (%v) = %v\n", parentKey, key)
			}
		}
	}
	return keys
}

func (g *Grid3D) ParseKey(key string) (int, int, int) {
	splits := strings.Split(key, ".")
	x, _ := strconv.Atoi(splits[0])
	y, _ := strconv.Atoi(splits[1])
	z, _ := strconv.Atoi(splits[2])
	return x, y, z
}

func (g *Grid3D) CountActiveNeighbours(parentKey string) int {
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

func (g *Grid3D) CountActiveTotal() int {
	activeCount := 0
	for _, value := range g.data {
		if value == ACTIVE {
			activeCount++
		}
	}
	return activeCount
}

func (g *Grid3D) Cycle() {
	data := make(map[string]string)

	minp, maxp := g.Dimensions()
	for z := minp.Z - 1; z <= maxp.Z+1; z++ {
		for y := minp.Y - 1; y <= maxp.Y+1; y++ {
			for x := minp.X - 1; x <= maxp.X+1; x++ {
				key := fmt.Sprintf("%v.%v.%v", x, y, z)
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
	g.data = data
}

// Dimensions returns the min/max points that exist
func (g *Grid3D) Dimensions() (goutils.Point3D, goutils.Point3D) {
	minp := goutils.Point3D{X: 10000, Y: 10000, Z: 10000}
	maxp := goutils.Point3D{X: -10000, Y: -10000, Z: -10000}
	for key := range g.data {
		x, y, z := g.ParseKey(key)
		minp.X = goutils.Min(x, minp.X)
		maxp.X = goutils.Max(x, maxp.X)
		minp.Y = goutils.Min(y, minp.Y)
		maxp.Y = goutils.Max(y, maxp.Y)
		minp.Z = goutils.Min(z, minp.Z)
		maxp.Z = goutils.Max(z, maxp.Z)
	}
	return minp, maxp
}

func (g *Grid3D) DebugZ(z int) string {
	minp, maxp := g.Dimensions()
	min_x := minp.X
	max_x := maxp.X
	min_y := minp.Y
	max_y := maxp.Y
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
