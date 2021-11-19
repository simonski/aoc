package aoc2018

import (
	"fmt"
	"strconv"
	"strings"
)

/*
--- Day 10: The Stars Align Description ---

It's no use; your navigation system simply isn't capable of providing walking directions in the arctic circle, and certainly not in 1018.

The Elves suggest an alternative.

In times like these, North Pole rescue operations will arrange points of light in the sky to guide missing Elves back to base. Unfortunately, the message is easy to miss: the points move slowly enough that it takes hours to align them, but have so much momentum that they only stay aligned for a second. If you blink at the wrong time, it might be hours before another message appears.

You can see these points of light floating in the distance, and record their position in the sky and their velocity, the relative change in position per second (your puzzle input). The coordinates are all given from your perspective; given enough time, those positions and velocities will move the points into a cohesive message!

Rather than wait, you decide to fast-forward the process and calculate what the points will eventually spell.

For example, suppose you note the following points:

position=< 9,  1> velocity=< 0,  2>
position=< 7,  0> velocity=<-1,  0>
position=< 3, -2> velocity=<-1,  1>
position=< 6, 10> velocity=<-2, -1>
position=< 2, -4> velocity=< 2,  2>
position=<-6, 10> velocity=< 2, -2>
position=< 1,  8> velocity=< 1, -1>
position=< 1,  7> velocity=< 1,  0>
position=<-3, 11> velocity=< 1, -2>
position=< 7,  6> velocity=<-1, -1>
position=<-2,  3> velocity=< 1,  0>
position=<-4,  3> velocity=< 2,  0>
position=<10, -3> velocity=<-1,  1>
position=< 5, 11> velocity=< 1, -2>
position=< 4,  7> velocity=< 0, -1>
position=< 8, -2> velocity=< 0,  1>
position=<15,  0> velocity=<-2,  0>
position=< 1,  6> velocity=< 1,  0>
position=< 8,  9> velocity=< 0, -1>
position=< 3,  3> velocity=<-1,  1>
position=< 0,  5> velocity=< 0, -1>
position=<-2,  2> velocity=< 2,  0>
position=< 5, -2> velocity=< 1,  2>
position=< 1,  4> velocity=< 2,  1>
position=<-2,  7> velocity=< 2, -2>
position=< 3,  6> velocity=<-1, -1>
position=< 5,  0> velocity=< 1,  0>
position=<-6,  0> velocity=< 2,  0>
position=< 5,  9> velocity=< 1, -2>
position=<14,  7> velocity=<-2,  0>
position=<-3,  6> velocity=< 2, -1>


22
0                    21
-6                   15
........#.............

Each line represents one point. Positions are given as <X, Y> pairs: X represents how far left (negative) or right (positive) the point appears, while Y represents how far up (negative) or down (positive) the point appears.

At 0 seconds, each point has the position given. Each second, each point's velocity is added to its position. So, a point with velocity <1, -2> is moving to the right, but is moving upward twice as quickly. If this point's initial position were <3, 9>, after 3 seconds, its position would become <6, 3>.

Over time, the points listed above would move like this:

Initially:
0123456789012345678901
........#............. 0
................#..... 1
.........#.#..#....... 2
...................... 3
#..........#.#.......# 4
...............#...... 5
....#................. 6
..#.#....#............ 7
.......#.............. 8
......#............... 9
...#...#.#...#........ 10
....#..#..#.........#. 11
.......#.............. 12
...........#..#....... 13
#...........#......... 14
...#.......#.......... 15


After 1 second:
......................
......................
..........#....#......
........#.....#.......
..#.........#......#..
......................
......#...............
....##.........#......
......#.#.............
.....##.##..#.........
........#.#...........
........#...#.....#...
..#...........#.......
....#.....#.#.........
......................
......................


After 2 seconds:
......................
......................
......................
..............#.......
....#..#...####..#....
......................
........#....#........
......#.#.............
.......#...#..........
.......#..#..#.#......
....#....#.#..........
.....#...#...##.#.....
........#.............
......................
......................
......................

After 3 seconds:
......................
......................
......................
......................
......#...#..###......
......#...#...#.......
......#...#...#.......
......#####...#.......
......#...#...#.......
......#...#...#.......
......#...#...#.......
......#...#..###......
......................
......................
......................
......................


After 4 seconds:
......................
......................
......................
............#.........
........##...#.#......
......#.....#..#......
.....#..##.##.#.......
.......##.#....#......
...........#....#.....
..............#.......
....#......#...#......
.....#.....##.........
...............#......
...............#......
......................
......................


*/

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func NewH() []string {
	// an H is a 5x8 grid
	// #...#
	// #...#
	// #...#
	// #####
	// #...#
	// #...#
	// #...#
	// #...#
	var letter []string
	letter = append(letter, "#...#")
	letter = append(letter, "#...#")
	letter = append(letter, "#...#")
	letter = append(letter, "#####")
	letter = append(letter, "#...#")
	letter = append(letter, "#...#")
	letter = append(letter, "#...#")
	letter = append(letter, "#...#")
	return letter
}

func NewI() []string {
	// an I is a 3x8 grid
	// ###
	// .#.
	// .#.
	// .#.
	// .#.
	// .#.
	// .#.
	// ###
	var letter []string
	letter = append(letter, "###")
	letter = append(letter, ".#.")
	letter = append(letter, ".#.")
	letter = append(letter, ".#.")
	letter = append(letter, ".#.")
	letter = append(letter, ".#.")
	letter = append(letter, ".#.")
	letter = append(letter, "###")
	return letter

}

type Point struct {
	line       string
	position_x int
	position_y int
	velocity_x int
	velocity_y int
	Key        string
}

type Grid struct {
	min_x    int
	min_y    int
	max_x    int
	max_y    int
	points   []*Point
	width    int
	height   int
	pointMap map[string]*Point
}

func NewGrid() *Grid {
	g := &Grid{}
	return g
}

func (g *Grid) Draw() {
	for y := g.min_y; y <= g.max_y; y++ {
		for x := g.min_x; x <= g.max_x; x++ {
			p := g.Get(x, y)
			if p != nil {
				// if p.isTopLeft {
				// 	fmt.Print("T")
				// } else {
				fmt.Print("#")
				// }
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func (g *Grid) IsH(point *Point) bool {
	// #...#
	// #...#
	// #...#
	// #####
	// #...#
	// #...#
	// #...#
	// #...#

	// so an H has a point which has a sibling at
	// right vertical
	// x+4, y
	// x+4, y+1
	// x+4, y+2
	// x+4, y+3
	// x+4, y+4
	// x+4, y+5
	// x+4, y+6
	// x+4, y+7
	x := point.position_x
	y := point.position_y
	for index := 0; index <= 7; index++ {
		if g.Get(x+4, y+index) == nil {
			// right bar not present
			return false
		}
		if g.Get(x, y+index) == nil {
			// left bar not present
			return false
		}
	}

	// horizontal line
	// x, y+3
	// x+1, y+3
	// x+2, y+3
	// x+3, y+3
	// x+4, y+3
	for index := 0; index <= 4; index++ {
		if g.Get(x+index, y+3) == nil {
			// bar not present
			return false
		}
	}

	return true

	// left and right must both be present

	// left vertical
	// x, y+1
	// x, y+2
	// x, y+3
	// x, y+4
	// x, y+5
	// x, y+6
	// x, y+7

}

func (g *Grid) Snapshot(x int, y int, width int, height int) []string {
	var snapshot []string
	for yindex := y; yindex < y+height; yindex++ {
		line := ""
		for xindex := x; xindex < x+width; xindex++ {
			point := g.Get(xindex, yindex)
			if point == nil {
				line += "."
			} else {
				line += "#"
			}
		}
		snapshot = append(snapshot, line)
	}
	return snapshot
}

func DebugLetter(letter []string) {
	for index := 0; index < len(letter); index++ {
		line := letter[index]
		fmt.Printf("%v\n", line)
	}
}

// indicates if the position x,y matches the passed search expression letter
func (g *Grid) Compare(letter []string, x int, y int) (bool, []string) {
	// in this case we need to match each character in letter
	// with each character in the grid at that point
	width := len(letter[0])
	height := len(letter)
	snapshot := g.Snapshot(x, y, width-1, height-1)

	for index := 0; index < len(letter); index++ {
		if letter[index] != snapshot[index] {
			return false, snapshot
		}
	}
	return true, snapshot
}

func (g *Grid) Load(lines []string) {
	for _, line := range lines {
		p := NewPoint(line)
		g.AddPoint(p)
	}

	if g.min_x < 0 {
		g.width = Abs(g.min_x) + g.max_x
	} else {
		g.width = g.max_x
	}

	if g.min_y < 0 {
		g.height = Abs(g.min_y) + g.max_y
	} else {
		g.height = g.max_y
	}
}

func (g *Grid) AddPoint(p *Point) {
	g.min_x = Min(p.position_x, g.min_x)
	g.max_x = Max(p.position_x, g.max_x)
	g.min_y = Min(p.position_y, g.min_y)
	g.max_y = Max(p.position_y, g.max_y)
	g.points = append(g.points, p)
}

func (g *Grid) Get(x int, y int) *Point {
	key := fmt.Sprintf("%v.%v", x, y)
	return g.pointMap[key]
}

func (g *Grid) Step(remap bool) {
	// g.step += 1
	if remap {
		g.pointMap = make(map[string]*Point)
	}
	// everything cycles to where it needs to be
	for _, p := range g.points {
		p.Step(remap)
		if remap {
			g.pointMap[p.Key] = p
		}
	}
}

func (g *Grid) CountLetters() int {
	h_counter := 0
	for _, p := range g.points {
		if g.IsH(p) {
			h_counter += 1
		}
	}
	return h_counter
}

func (g *Grid) Debug() {
	fmt.Printf("Grid.debug() min_x=%v, min_y=%v, max_x=%v, max_y=%v, width=%v, height=%v\n", g.min_x, g.min_y, g.max_x, g.max_y, g.width, g.height)
}

func (p *Point) Debug() {
	fmt.Printf("line=%v, px=%v, py=%v, vx=%v, vy=%v\n", p.line, p.position_x, p.position_y, p.velocity_x, p.velocity_y)
}

func (p *Point) Step(remap bool) {
	// p.step += 1
	p.position_x += p.velocity_x
	p.position_y += p.velocity_y
	if remap {
		p.Key = fmt.Sprintf("%v.%v", p.position_x, p.position_y)
	}

	// p.position_x = applyrange(p.position_x, p.velocity_x, g.min_x, g.max_x)
	// p.position_y = applyrange(p.position_y, p.velocity_y, g.min_y, g.max_y)
}

func NewPoint(line string) *Point {
	// position=< 9,  1> velocity=< 0,  2>
	splits := strings.Split(line, "velocity")
	position := strings.Split(splits[0], "=")[1]
	position = strings.Replace(position, "position=", "", -1)
	position = strings.Replace(position, "<", "", -1)
	position = strings.Replace(position, ">", "", -1)
	position = strings.Replace(position, " ", "", -1)
	p := strings.Split(position, ",")
	pos_x, _ := strconv.Atoi(p[0])
	pos_y, _ := strconv.Atoi(p[1])

	velocity := strings.Split(splits[1], "=")[1]
	velocity = strings.Replace(velocity, "velocity=", "", -1)
	velocity = strings.Replace(velocity, "<", "", -1)
	velocity = strings.Replace(velocity, ">", "", -1)
	velocity = strings.Replace(velocity, " ", "", -1)
	v := strings.Split(velocity, ",")
	v_x, _ := strconv.Atoi(v[0])
	v_y, _ := strconv.Atoi(v[1])

	return &Point{line: line, position_x: pos_x, position_y: pos_y, velocity_x: v_x, velocity_y: v_y}
}

func applyrange(value int, changeby int, lowerbound int, upperbound int) int {
	value += changeby
	if value < lowerbound {
		// we went under the lowerbound, wrap around to the upperbound
		diff := Abs(value - lowerbound)
		value = upperbound - diff
	} else if value > upperbound {
		// we went over the upperbound, wrap aroudn to the lowerbound
		diff := Abs(value - upperbound)
		value = lowerbound + diff
	}
	return value
}

// rename this to the year and day in question
func (app *Application) Y2018D10P1() {
	lines := strings.Split(DAY_2018_10_DATA_TEST, "\n")
	g := NewGrid()
	g.Load(lines)

	h := g.CountLetters()
	g.Draw()

	g.Step(true)
	h = g.CountLetters()
	fmt.Printf("There is %v H.\n", h)
	g.Draw()

	g.Step(true)
	h = g.CountLetters()
	fmt.Printf("There is %v H.\n", h)
	g.Draw()

	g.Step(true)
	h = g.CountLetters()
	fmt.Printf("There is %v H.\n", h)
	g.Draw()

}

func DrawLetter(l []string) {
	for _, line := range l {
		fmt.Printf("%v\n", line)
	}
}

// rename this to the year and day in question
func (app *Application) Y2018D10P2() {
	// return
	lines := strings.Split(DAY_2018_10_DATA, "\n")
	g := NewGrid()
	g.Load(lines)
	fmt.Printf("There are %v stars.\n", len(g.points))

	step := 0
	for index := 0; index < 1000000; index++ {
		step += 1
		g.Step(true)
		// g.DrawLimited(50, 50)
		h := g.CountLetters()
		// h, i := 0, 0
		if step%1000 == 0 {
			fmt.Printf("[%v] %v H\n", step, h)
		}
		if h > 0 {
			fmt.Printf("Step %v, H %v\n", step, h)
			break
		}

	}

}

// rename and uncomment this to the year and day in question once complete for a gold star!
// func (app *Application) Y202018D10P1Render() {
// }

// rename and uncomment this to the year and day in question once complete for a gold star!
// func (app *Application) Y202018D10P2Render() {
// }

// this is what we will reflect and call - so both parts with run. It's up to you to make it print nicely etc.
// The app reference has a CLI for logging.
func (app *Application) Y202018D10() {
	app.Y2018D10P1()
	app.Y2018D10P2()
}
