package aoc2018

import (
	"fmt"
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

After 3 seconds, the message appeared briefly: HI. Of course, your message will be much longer and will take many more seconds to appear.

What message will eventually appear in the sky?

*/

// rename this to the year and day in question
func (app *Application) Y2018D10P1() {
	lines := strings.Split(DAY_2018_10_DATA_TEST, "\n")
	g := NewGrid()
	g.Load(lines)
	g.Remap()

	// let's see if we can find the lowest common factors for lining up many vertical
	// bars of size 8.  this is because I assume a letter is a constant height (8)
	// will generall have at least 1 vertical bar.  So, if I find a number of them, this
	// is a signal that this timeperiod is a set of words.
	// what I expect is a timecode; this timecode is what I then move to and *should*
	// find myself able to print the line

	// (x1,y1)+(vx1+vx1)  (x2,y2)+(vx2+vy2)
	// x1+(n*vx1)   common n
	// x2+(n*vx2)

	// 6 + 1
	// 3 + 2   << never meet  because this cannot catch  up; so never compare

	// 2+1
	// -1000 + 2

	// know they will intersect
	// 	IF they are going in the same direction
	// 	AND the LOWER is travelling *faster* in that direction
	// 	THEN calculate number of steps to intersect
	// 		faster is faster by a ratio
	// 		above
	// 		500steps gets to -1000 0
	// 		500 steps gets 2 to 502
	// 		now calculate, how many steps to catch up with 2+1
	// 		so second is 2x faster
	// 		difference is 500
	// 		log(2) 500
	// 		250steps, we are now at 1502 and 752
	// 		difference is now 250
	// 		in 125 steps we add on to be
	// 		125

	// ox1 := -1753
	// v1 := 7
	// ox2 := 1313
	// v2 := 6

	// x1 := ox1
	// x2 := ox2

	// step := 0
	// for {
	// 	step += 1
	// 	x1 = ox1 + (v1 * step)
	// 	x2 = ox2 + (v2 * step)
	// 	if x1 == x2 {
	// 		// fmt.Printf("x1==x2 (%v,%v), step=%v\n", x1, x2, step)
	// 		break
	// 	} else if x1 > x2 {
	// 		// fmt.Printf("x1>=x2 (%v, %v), step=%v\n", x1, x2, step)
	// 		break
	// 	} else {
	// 		// fmt.Printf("x1!=x2 (%v,%v), step=%v\n", x1, x2, step)
	// 	}
	// }

	// where x1 = -1000 v1 = 2
	//       x2 = 0, v1 = 1

	// steps = 1000 to get to same
	// difference = 1000
	// 500 steps to close difference gives new difference of 500
	// x2 nw at 500, so 500 steps halved the distsance

	// on x1,500 steps creates 1000, 0+1000=1000
	// 500 gives 500, 500+500 = 100
	// so 1000 steps

	// -1643

	// return

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

	g.Step(true)
	h = g.CountLetters()
	fmt.Printf("There is %v H.\n", h)
	g.Draw()

}

// rename this to the year and day in question
func (app *Application) Y2018D10P2() {

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
