package main

/*
--- Day 3: Perfectly Spherical Houses in a Vacuum ---
Santa is delivering presents to an infinite two-dimensional grid of houses.

He begins by delivering a present to the house at his starting location, and then an elf at the North Pole calls him via radio and tells him where to move next. Moves are always exactly one house to the north (^), south (v), east (>), or west (<). After each move, he delivers another present to the house at his new location.

However, the elf back at the north pole has had a little too much eggnog, and so his directions are a little off, and Santa ends up visiting some houses more than once. How many houses receive at least one present?

For example:

> delivers presents to 2 houses: one at the starting location, and one to the east.
^>v< delivers presents to 4 houses in a square, including twice to the house at his starting/ending location.
^v^v^v^v^v delivers a bunch of presents to some very lucky children at only 2 houses.*/

import (
	"fmt"

	goutils "github.com/simonski/goutils"
)

// AOC_2015_03 is the entrypoint
func AOC_2015_03(cli *goutils.CLI) {
	AOC_2015_03_part1_attempt1(cli)
	AOC_2015_03_part2_attempt1(cli)
}

func AOC_2015_03_part1_attempt1(cli *goutils.CLI) {
	grid := NewGrid201503()
	grid.Increment(0, 0)
	for index := 0; index < len(DAY_2015_03_DATA); index++ {
		c := DAY_2015_03_DATA[index : index+1]
		if c == "<" {
			grid.West()
		} else if c == ">" {
			grid.East()
		} else if c == "^" {
			grid.North()
		} else if c == "v" {
			grid.South()
		} else {
			panic("wtf")
		}
	}

	total_single_presents := 0
	for _, value := range grid.counter.data {
		if value == 1 {
			total_single_presents++
		}
	}

	fmt.Printf("Part1: Single present households: %v\n", total_single_presents)
	fmt.Printf("Part1: Total households: %v\n", len(grid.counter.data))
}

/*
--- Part Two ---
The next year, to speed up the process, Santa creates a robot version of himself, Robo-Santa, to deliver presents with him.

Santa and Robo-Santa start at the same location (delivering two presents to the same starting house), then take turns moving based on instructions from the elf, who is eggnoggedly reading from the same script as the previous year.

This year, how many houses receive at least one present?

For example:

^v delivers presents to 3 houses, because Santa goes north, and then Robo-Santa goes south.
^>v< now delivers presents to 3 houses, and Santa and Robo-Santa end up back where they started.
^v^v^v^v^v now delivers presents to 11 houses, with Santa going one direction and Robo-Santa going the other.
*/

func AOC_2015_03_part2_attempt1(cli *goutils.CLI) {
	grid := NewGrid201503()
	grid.Increment(0, 0)
	grid.Increment(0, 0)
	data := DAY_2015_03_DATA
	for index := 0; index < len(data); index += 2 {
		c1 := data[index : index+1]
		if c1 == "<" {
			grid.West()
		} else if c1 == ">" {
			grid.East()
		} else if c1 == "^" {
			grid.North()
		} else if c1 == "v" {
			grid.South()
		} else {
			panic("wtf")
		}

		idx := index + 1
		c2 := data[idx : idx+1]
		if c2 == "<" {
			grid.RoboWest()
		} else if c2 == ">" {
			grid.RoboEast()
		} else if c2 == "^" {
			grid.RoboNorth()
		} else if c2 == "v" {
			grid.RoboSouth()
		} else {
			panic("wtf")
		}

	}

	total_single_presents := 0
	for _, value := range grid.counter.data {
		if value == 1 {
			total_single_presents++
		}
	}

	fmt.Printf("Part2: Single present households: %v\n", total_single_presents)
	fmt.Printf("Part2: Total households: %v\n", len(grid.counter.data))
}

type Grid201503 struct {
	x       int
	y       int
	robo_x  int
	robo_y  int
	counter *Counter
}

func NewGrid201503() *Grid201503 {
	g := Grid201503{x: 0, y: 0, counter: NewCounter()}
	return &g
}

func (grid *Grid201503) Increment(x int, y int) {
	key := fmt.Sprintf("%v,%v", x, y)
	grid.counter.Increment(key)
}

func (grid *Grid201503) North() {
	y := grid.y + 1
	x := grid.x
	grid.Increment(x, y)
	grid.x = x
	grid.y = y
}

func (grid *Grid201503) South() {
	y := grid.y - 1
	x := grid.x
	grid.Increment(x, y)
	grid.x = x
	grid.y = y
}

func (grid *Grid201503) East() {
	y := grid.y
	x := grid.x + 1
	grid.Increment(x, y)
	grid.x = x
	grid.y = y
}

func (grid *Grid201503) West() {
	y := grid.y
	x := grid.x - 1
	grid.Increment(x, y)
	grid.x = x
	grid.y = y
}

func (grid *Grid201503) RoboNorth() {
	y := grid.robo_y + 1
	x := grid.robo_x
	grid.Increment(x, y)
	grid.robo_x = x
	grid.robo_y = y
}

func (grid *Grid201503) RoboSouth() {
	y := grid.robo_y - 1
	x := grid.robo_x
	grid.Increment(x, y)
	grid.robo_x = x
	grid.robo_y = y
}

func (grid *Grid201503) RoboEast() {
	y := grid.robo_y
	x := grid.robo_x + 1
	grid.Increment(x, y)
	grid.robo_x = x
	grid.robo_y = y
}

func (grid *Grid201503) RoboWest() {
	y := grid.robo_y
	x := grid.robo_x - 1
	grid.Increment(x, y)
	grid.robo_x = x
	grid.robo_y = y
}
