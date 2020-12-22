package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test_AOC2020_17_Test1(t *testing.T) {

	input := `.#.
..#
###`
	g := NewGrid3D(input)

	minp, maxp := g.Dimensions()
	fmt.Printf("Dimensions %v %v\n", minp, maxp)
	fmt.Printf(">>>>>>>>>\n")
	fmt.Printf("Z=-1\n")
	fmt.Printf(g.DebugZ(-1))
	fmt.Printf(">>>>>>>>>\n\n")

	fmt.Printf(">>>>>>>>>\n")
	fmt.Printf("Z=0\n")
	fmt.Printf(g.DebugZ(0))
	fmt.Printf(">>>>>>>>>\n\n")

	fmt.Printf(">>>>>>>>>\n")
	fmt.Printf("Z=1\n")
	fmt.Printf(g.DebugZ(1))
	fmt.Printf(">>>>>>>>>\n\n")

	verifyDebugZ(0, g, input, t)
	// if strings.TrimSpace(g.DebugZ(0)) != strings.TrimSpace(input) {
	// 	t.Errorf("Expected grid input z=0 to be\n%v\nbut was\n%v\n", input, g.DebugZ(0))
	// }

	verifyCubeNeighbours("0.0.0", 26, g, t)

	// // z=0
	// // .#.
	// // ..#
	// // ###

	verifyCube("0.0.0", INACTIVE, g, t)
	verifyCube("1.0.0", ACTIVE, g, t)
	verifyCube("2.0.0", INACTIVE, g, t)
	verifyCube("0.1.0", INACTIVE, g, t)
	verifyCube("1.1.0", INACTIVE, g, t)
	verifyCube("2.1.0", ACTIVE, g, t)
	verifyCube("0.2.0", ACTIVE, g, t)
	verifyCube("1.2.0", ACTIVE, g, t)
	verifyCube("2.2.0", ACTIVE, g, t)
	verifyCubeActiveNeighbours("0.0.0", 1, g, t)

	verifyCube("1.1.0", INACTIVE, g, t)
	verifyCubeActiveNeighbours("1.1.0", 5, g, t)
	verifyCubeActiveNeighbours("1.1.-1", 5, g, t)
	verifyCubeTotalActive(5, g, t)

	// // before we cycle once, the z-1 should be inactive
	verifyParseKey("0.0.-1", 0, 0, -1, g, t)
	verifyCube("0.0.-1", INACTIVE, g, t)

	// 	After 1 cycle:

	verifyCubeTotalActive(5, g, t)
	g.Cycle()
	verifyCubeTotalActive(11, g, t)

	g.Cycle()
	verifyCubeTotalActive(21, g, t)

	g.Cycle()
	verifyCubeTotalActive(38, g, t)

	g.Cycle()
	verifyCubeTotalActive(58, g, t)

	g.Cycle()
	verifyCubeTotalActive(101, g, t)

	g.Cycle()
	verifyCubeTotalActive(112, g, t)

	// 	minp, maxp = g.Dimensions()
	// 	for z := minp.z; z <= maxp.z; z++ {
	// 		fmt.Printf("Dimensions %v %v\n", minp, maxp)
	// 		fmt.Printf(">>>>>>>>>\n")
	// 		fmt.Printf("Z=%v\n", minp.z)
	// 		fmt.Printf(g.DebugZ(z))
	// 		fmt.Printf(">>>>>>>>>\n\n")
	// 	}

	// 	expectedCycle1ZM1 := `#..
	// ..#
	// .#.`

	// 	expectedCycle1Z0 := `#.#
	// .##
	// .#.`

	// 	expectedCycle1Z1 := `#..
	// ..#
	// .#.`

	// 	verifyDebugZ(-1, g, expectedCycle1ZM1, t)
	// 	verifyDebugZ(0, g, expectedCycle1Z0, t)
	// 	verifyDebugZ(1, g, expectedCycle1Z1, t)

	// z=-1
	// #..
	// ..#
	// .#.

	// z=0
	// #.#
	// .##
	// .#.
	// verifyCube("0.0.0", ACTIVE, g, t)
	// verifyCube("1.0.0", INACTIVE, g, t)
	// verifyCube("2.0.0", ACTIVE, g, t)

	// verifyCube("0.1.0", INACTIVE, g, t)
	// verifyCube("1.1.0", ACTIVE, g, t)
	// verifyCube("2.1.0", ACTIVE, g, t)

	// verifyCube("0.2.0", INACTIVE, g, t)
	// verifyCube("1.2.0", ACTIVE, g, t)
	// verifyCube("2.2.0", INACTIVE, g, t)

	// z=1
	// #..
	// ..#
	// .#.

	// once we cycle the z -1 should be active
	// verifyCube("0.0.-1", ACTIVE, g, t)

	// verifyCubeTotalActive(11, g, t)

	// g.Cycle()
	// verifyCubeTotalActive(21, g, t)
	// g.Cycle()
	// verifyCubeTotalActive(38, g, t)
	// g.Cycle()
	// g.Cycle()
	// g.Cycle()
	// verifyCubeTotalActive(112, g, t)
	/*
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
	*/

}

func verifyCube(key string, expected string, g *Grid3D, t *testing.T) {
	actual := g.Get(key)
	if actual != expected {
		t.Errorf("verifyCube(%v) expects %v actual %v\n", key, expected, actual)
	}
}

func verifyCubeTotalActive(expected int, g *Grid3D, t *testing.T) {
	actual := g.CountActiveTotal()
	if actual != expected {
		t.Errorf("verifyCubeTotalActive(): expects %v actual %v\n", expected, actual)
	}
}

func verifyCubeActiveNeighbours(key string, expected int, g *Grid3D, t *testing.T) {
	actual := g.CountActiveNeighbours(key)
	if actual != expected {
		t.Errorf("verifyCubeActiveNeigbours(%v): expects %v actual %v\n", key, expected, actual)
	}
}

func verifyCubeNeighbours(key string, expected int, g *Grid3D, t *testing.T) {
	actual := g.Neighbours(key)
	if len(actual) != expected {
		t.Errorf("verifyCubeNeighbours(%v): expects %v actual %v\n\n", key, expected, len(actual))
	}

	// for _, n := range actual {
	// 	fmt.Printf("%v\n", n)
	// }

}

func verifyParseKey(key string, expected_x int, expected_y int, expected_z int, g *Grid3D, t *testing.T) {
	x, y, z := g.ParseKey(key)
	if x != expected_x {
		t.Errorf("verifyParseKey(%v) expected x %v actual x %v\n", key, expected_x, x)
	}
	if y != expected_y {
		t.Errorf("verifyParseKey(%v) expected x %v actual x %v\n", key, expected_y, y)
	}
	if z != expected_z {
		t.Errorf("verifyParseKey(%v) expected x %v actual x %v\n", key, expected_z, z)
	}
}

func verifyDebugZ(layer int, g *Grid3D, expected string, t *testing.T) {
	actual := g.DebugZ(layer)
	if strings.TrimSpace(actual) != strings.TrimSpace(expected) {
		t.Errorf("verifyDebugZ(%v)\nexpected\n%v\nactual\n%v\n", layer, expected, actual)
	}
}
