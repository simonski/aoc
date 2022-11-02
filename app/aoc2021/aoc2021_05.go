package aoc2021

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
	goutils "github.com/simonski/goutils"
)

/*
--- Day 5: Hydrothermal Venture ---
You come across a field of hydrothermal vents on the ocean floor! These vents constantly produce large, opaque clouds, so it would be best to avoid them if possible.

They tend to form in lines; the submarine helpfully produces a list of nearby lines of vents (your puzzle input) for you to review. For example:

0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2
Each line of vents is given as a line segment in the format x1,y1 -> x2,y2 where x1,y1 are the coordinates of one end the line segment and x2,y2 are the coordinates of the other end. These line segments include the points at both ends. In other words:

An entry like 1,1 -> 1,3 covers points 1,1, 1,2, and 1,3.
An entry like 9,7 -> 7,7 covers points 9,7, 8,7, and 7,7.
For now, only consider horizontal and vertical lines: lines where either x1 = x2 or y1 = y2.

So, the horizontal and vertical lines from the above list would produce the following diagram:

.......1..
..1....1..
..1....1..
.......1..
.112111211
..........
..........
..........
..........
222111....
In this diagram, the top left corner is 0,0 and the bottom right corner is 9,9. Each position is shown as the number of lines which cover that point or . if no line covers that point. The top-left pair of 1s, for example, comes from 2,2 -> 2,1; the very bottom row is formed by the overlapping lines 0,9 -> 5,9 and 0,9 -> 2,9.

To avoid the most dangerous areas, you need to determine the number of points where at least two lines overlap. In the above example, this is anywhere in the diagram with a 2 or larger - a total of 5 points.

Consider only horizontal and vertical lines. At how many points do at least two lines overlap?

*/

func (app *Application) Y2021D05_Summary() *utils.Summary {
	s := utils.NewSummary(2021, 5)
	s.Name = "Hydrothermal Venture"
	s.ProgressP1 = utils.Completed
	s.ProgressP2 = utils.Completed
	return s
}

type Grid struct {
	Width  int     `json:"width"`
	Height int     `json:"height"`
	Lines  []*Line `json:"lines"`
}

func NewGrid(data string) *Grid {
	splits := strings.Split(data, "\n")
	minx := 100
	miny := 100
	maxx := -100
	maxy := -100
	lines := make([]*Line, 0)

	for _, line_s := range splits {
		line := NewLine(line_s)
		lines = append(lines, line)
		// fmt.Print(line.Debug())
		minx = goutils.Min(minx, line.X1)
		minx = goutils.Min(minx, line.X2)
		maxx = goutils.Max(maxx, line.X1)
		maxx = goutils.Max(maxx, line.X2)

		miny = goutils.Min(miny, line.Y1)
		miny = goutils.Min(miny, line.Y2)
		maxy = goutils.Max(maxy, line.Y1)
		maxy = goutils.Max(maxy, line.Y2)
	}
	// fmt.Printf("minx=%v, miny=%v, maxx=%v, maxy=%v\n", minx, miny, maxx, maxy)
	width := maxx - minx
	height := maxx - miny
	return &Grid{Width: width, Height: height, Lines: lines}
}

func (grid *Grid) PlayPartOne() int {
	// the x * y grid containing all the occupants
	counts := make(map[string]int)
	for _, line := range grid.Lines {
		// fmt.Printf("Line[%v] %v", index, line.Debug())
		if line.IsHorizontal() {
			min_x := goutils.Min(line.X1, line.X2)
			max_x := goutils.Max(line.X1, line.X2)

			for x := min_x; x <= max_x; x++ {
				key := fmt.Sprintf("%v,%v", x, line.Y1)
				value := counts[key]
				value++
				counts[key] = value
				// if value > 1 {
				// 	fmt.Printf("%v=%v\n", key, value)
				// }
			}
		} else if line.IsVertical() {
			min_y := goutils.Min(line.Y1, line.Y2)
			max_y := goutils.Max(line.Y1, line.Y2)
			for y := min_y; y <= max_y; y++ {
				key := fmt.Sprintf("%v,%v", line.X1, y)
				value := counts[key]
				value++
				counts[key] = value
				// if value > 1 {
				// 	fmt.Printf("%v=%v\n", key, value)
				// }
			}
		}
		fmt.Printf("%v", DebugCounts(line, counts))
		fmt.Println()
	}

	counter := 0
	for _, value := range counts {
		if value > 1 {
			counter += 1
		}
	}
	return counter
}

func (grid *Grid) PlayPartTwo() int {
	// the x * y grid containing all the occupants
	counts := make(map[string]int)
	for _, line := range grid.Lines {
		// fmt.Printf("Line[%v] %v", index, line.Debug())
		if line.IsHorizontal() {
			min_x := goutils.Min(line.X1, line.X2)
			max_x := goutils.Max(line.X1, line.X2)

			for x := min_x; x <= max_x; x++ {
				key := fmt.Sprintf("%v,%v", x, line.Y1)
				value := counts[key]
				value++
				counts[key] = value
				// if value > 1 {
				// 	fmt.Printf("%v=%v\n", key, value)
				// }
			}

		} else if line.IsVertical() {
			min_y := goutils.Min(line.Y1, line.Y2)
			max_y := goutils.Max(line.Y1, line.Y2)
			for y := min_y; y <= max_y; y++ {
				key := fmt.Sprintf("%v,%v", line.X1, y)
				value := counts[key]
				value++
				counts[key] = value
				// if value > 1 {
				// 	fmt.Printf("%v=%v\n", key, value)
				// }
			}
		} else {
			// It MUST be diagonal - I don't need to detect it cos the rules
			// is it a left-to-right or right-to-left
			if line.X1 < line.X2 {
				y := line.Y1
				y_diff := 0
				if line.Y1 < line.Y2 {
					// it is a downy angle "\"
					y_diff = 1
				} else {
					// it is an uppy angle "/"
					y_diff = -1
				}
				for x := line.X1; x <= line.X2; x++ {
					// is it an uppy angle or a downy angle
					// we will go left to right, but will it be up or down
					key := fmt.Sprintf("%v,%v", x, y)
					value := counts[key]
					value++
					counts[key] = value
					y += y_diff
				}
			} else {
				y := line.Y1
				y_diff := 0
				if line.Y1 < line.Y2 {
					// it is a downy angle "\"
					y_diff = 1
				} else {
					// it is an uppy angle "/"
					y_diff = -1
				}
				for x := line.X1; x >= line.X2; x-- {
					key := fmt.Sprintf("%v,%v", x, y)
					value := counts[key]
					value++
					counts[key] = value
					y += y_diff
				}
			}
		}
		fmt.Printf("%v", DebugCounts(line, counts))
		fmt.Println()
	}

	counter := 0
	for _, value := range counts {
		if value > 1 {
			counter += 1
		}
	}
	return counter
}

func DebugCounts(line *Line, counts map[string]int) string {
	result := ""
	result += fmt.Sprint(line.Debug())
	for y := 0; y < 10; y++ {
		gridline := ""
		for x := 0; x < 10; x++ {
			key := fmt.Sprintf("%v,%v", x, y)
			value := counts[key]
			if value == 0 {
				gridline += "."
			} else {
				gridline += fmt.Sprintf("%v", value)
			}
		}
		result += gridline
		result += "\n"

	}
	return result
}

type Line struct {
	X1 int `json:"x1"`
	Y1 int `json:"y1"`
	X2 int `json:"x2"`
	Y2 int `json:"y2"`
}

func (line *Line) Debug() string {
	linetype := "?"
	if line.IsHorizontal() {
		linetype = "Horizontal"
	} else if line.IsVertical() {
		linetype = "Vertical"
	}
	return fmt.Sprintf("(%v,%v)->(%v,%v) [%v]\n", line.X1, line.Y1, line.X2, line.Y2, linetype)
}

func NewLine(line_s string) *Line {
	// 11,22 -> 935,946
	line_s = strings.ReplaceAll(line_s, " -> ", ",")
	points := strings.Split(line_s, ",")
	x1, _ := strconv.Atoi(points[0])
	y1, _ := strconv.Atoi(points[1])
	x2, _ := strconv.Atoi(points[2])
	y2, _ := strconv.Atoi(points[3])

	line := Line{X1: x1, X2: x2, Y1: y1, Y2: y2}
	return &line
}

func (line *Line) IsHorizontal() bool {
	return line.Y1 == line.Y2
}

func (line *Line) IsVertical() bool {
	return line.X1 == line.X2
}

func (line *Line) IsDiagonally() bool {
	return line.X1 == line.X2
}

func Part1D5(data string) {
	grid := NewGrid(data)
	result := grid.PlayPartOne()
	fmt.Printf("Part 1: %v\n", result)
}

func Part2D5(data string) {
	grid := NewGrid(data)
	result := grid.PlayPartTwo()
	fmt.Printf("Part 2: %v\n", result)
}

// rename this to the year and day in question
func (app *Application) Y2021D05P1() {
	Part1D5(DAY_2021_05_TEST_DATA)
	// Part1D5(DAY_2021_05_DATA)
}

/*
--- Part Two ---
Unfortunately, considering only horizontal and vertical lines doesn't give you the full picture; you need to also consider diagonal lines.

Because of the limits of the hydrothermal vent mapping system, the lines in your list will only ever be horizontal, vertical, or a diagonal line at exactly 45 degrees. In other words:

An entry like 1,1 -> 3,3 covers points 1,1, 2,2, and 3,3.
An entry like 9,7 -> 7,9 covers points 9,7, 8,8, and 7,9.
Considering all lines from the above example would now produce the following diagram:

1.1....11.
.111...2..
..2.1.111.
...1.2.2..
.112313211
...1.2....
..1...1...
.1.....1..
1.......1.
222111....
You still need to determine the number of points where at least two lines overlap. In the above example, this is still anywhere in the diagram with a 2 or larger - now a total of 12 points.

Consider all of the lines. At how many points do at least two lines overlap?
*/
func (app *Application) Y2021D05P2() {
	Part2D5(DAY_2021_05_TEST_DATA)
	Part2D5(DAY_2021_05_DATA)
}

// rename and uncomment this to the year and day in question once complete for a gold star!
// func (app *Application) Y20XXDXXP1Render() {
// }

// rename and uncomment this to the year and day in question once complete for a gold star!
// func (app *Application) Y20XXDXXP2Render() {
// }

// this is what we will reflect and call - so both parts with run. It's up to you to make it print nicely etc.
// The app reference has a CLI for logging.
func (app *Application) Y2021D05() {
	app.Y2021D05P1()
	app.Y2021D05P2()
}
