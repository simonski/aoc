package aoc2021

import (
	"fmt"
	"strconv"
	"strings"

	utils "github.com/simonski/aoc/utils"
)

/*
--- Day 13: Transparent Origami ---
You reach another volcanically active part of the cave. It would be nice if you could do some kind of thermal imaging so you could tell ahead of time which caves are too hot to safely enter.

Fortunately, the submarine seems to be equipped with a thermal camera! When you activate it, you are greeted with:

Congratulations on your purchase! To activate this infrared thermal imaging
camera system, please enter the code found on page 1 of the manual.
Apparently, the Elves have never used this feature. To your surprise, you manage to find the manual; as you go to open it, page 1 falls out. It's a large sheet of transparent paper! The transparent paper is marked with random dots and includes instructions on how to fold it up (your puzzle input). For example:

6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5
The first section is a list of dots on the transparent paper. 0,0 represents the top-left coordinate. The first value, x, increases to the right. The second value, y, increases downward. So, the coordinate 3,0 is to the right of 0,0, and the coordinate 0,7 is below 0,0. The coordinates in this example form the following pattern, where # is a dot on the paper and . is an empty, unmarked position:

...#..#..#.
....#......
...........
#..........
...#....#.#
...........
...........
...........
...........
...........
.#....#.##.
....#......
......#...#
#..........
#.#........
Then, there is a list of fold instructions. Each instruction indicates a line on the transparent paper and wants you to fold the paper up (for horizontal y=... lines) or left (for vertical x=... lines). In this example, the first fold instruction is fold along y=7, which designates the line formed by all of the positions where y is 7 (marked here with -):

...#..#..#.
....#......
...........
#..........
...#....#.#
...........
...........
-----------
...........
...........
.#....#.##.
....#......
......#...#
#..........
#.#........
Because this is a horizontal line, fold the bottom half up. Some of the dots might end up overlapping after the fold is complete, but dots will never appear exactly on a fold line. The result of doing this fold looks like this:

#.##..#..#.
#...#......
......#...#
#...#......
.#.#..#.###
...........
...........
Now, only 17 dots are visible.

Notice, for example, the two dots in the bottom left corner before the transparent paper is folded; after the fold is complete, those dots appear in the top left corner (at 0,0 and 0,1). Because the paper is transparent, the dot just below them in the result (at 0,3) remains visible, as it can be seen through the transparent paper.

Also notice that some dots can end up overlapping; in this case, the dots merge together and become a single dot.

The second fold instruction is fold along x=5, which indicates this line:

#.##.|#..#.
#...#|.....
.....|#...#
#...#|.....
.#.#.|#.###
.....|.....
.....|.....
Because this is a vertical line, fold left:

#####
#...#
#...#
#...#
#####
.....
.....
The instructions made a square!

The transparent paper is pretty big, so for now, focus on just completing the first fold. After the first fold in the example above, 17 dots are visible - dots that end up overlapping after the fold is completed count as a single dot.

How many dots are visible after completing just the first fold instruction on your transparent paper?
*/

type Day13Point struct {
	X int
	Y int
}

func (p *Day13Point) Key() string {
	return fmt.Sprintf("%v.%v", p.X, p.Y)
}

type Day13Fold struct {
	Axis  string
	Value int
}

type Day13Grid struct {
	Width     int
	Height    int
	Points    []*Day13Point
	PointsMap map[string]*Day13Point
	Folds     []*Day13Fold
}

func (g *Day13Grid) GetPoint(x int, y int) *Day13Point {
	key := fmt.Sprintf("%v.%v", x, y)
	return g.PointsMap[key]
}

func (g *Day13Grid) Fold(f *Day13Fold) {
	points := make([]*Day13Point, 0)
	if f.Axis == "y" {
		// then y
		for _, p := range g.Points {
			if p.Y > g.Height/2 {
				newY := g.Height - p.Y
				p.Y = newY
			}
			points = append(points, p)
		}

	} else {
		// then x
		for _, p := range g.Points {
			if p.X > g.Width/2 {
				newX := g.Width - p.X
				p.X = newX
			}
			points = append(points, p)
		}
		// then x
	}

	g.Points = points
	pointsMap := make(map[string]*Day13Point)
	width := 0
	height := 0
	for _, p := range g.Points {
		pointsMap[p.Key()] = p
		width = utils.Max(width, p.X)
		height = utils.Max(height, p.Y)
	}
	g.PointsMap = pointsMap
	g.Width = width
	g.Height = height

}

func (g *Day13Grid) Debug() string {
	line := ""
	for y := 0; y <= g.Height; y++ {
		for x := 0; x <= g.Width; x++ {
			p := g.GetPoint(x, y)
			if p != nil {
				line += "#"
			} else {
				line += "."
			}
		}
		line += "\n"
	}
	return line
}

func NewDay13Grid(data string) *Day13Grid {
	lines := strings.Split(data, "\n")
	points := make([]*Day13Point, 0)
	pointsMap := make(map[string]*Day13Point)
	folds := make([]*Day13Fold, 0)
	isFold := false
	width := 0
	height := 0
	for _, line := range lines {
		if line == "" {
			// move to the foldalong reader
			isFold = true
			continue
		}
		if !isFold {
			// 0,14
			splits := strings.Split(line, ",")
			x, _ := strconv.Atoi(splits[0])
			y, _ := strconv.Atoi(splits[1])
			p := &Day13Point{X: x, Y: y}
			points = append(points, p)
			pointsMap[p.Key()] = p
			width = utils.Max(width, x)
			height = utils.Max(height, y)

		} else {
			// fold along y=7
			line = strings.ReplaceAll(line, "fold along ", "")
			splits := strings.Split(line, "=")
			axis := splits[0]
			value, _ := strconv.Atoi(splits[1])
			fold := &Day13Fold{Axis: axis, Value: value}
			folds = append(folds, fold)

		}
	}
	return &Day13Grid{Points: points, PointsMap: pointsMap, Folds: folds, Width: width, Height: height}
}

// rename this to the year and day in question
func (app *Application) Y2021D13P1() {
	g := NewDay13Grid(DAY_2021_13_TEST_DATA)
	fmt.Println(g.Debug())
	for _, fold := range g.Folds {
		fmt.Println()
		g.Fold(fold)
		fmt.Println(g.Debug())
	}

	g2 := NewDay13Grid(DAY_2021_13_DATA)
	// fmt.Println(g2.Debug())
	fold := g2.Folds[0]
	g2.Fold(fold)
	fmt.Printf("There are %v dots visible.\n", len(g2.PointsMap))

	g3 := NewDay13Grid(DAY_2021_13_DATA)
	// fmt.Println(g2.Debug())
	count := 0
	for _, f := range g3.Folds {
		g3.Fold(f)
		count++
		fmt.Printf("%v,%v There are %v dots visible.\n", g3.Width, g3.Height, len(g3.PointsMap))
		if count > 5 {
			fmt.Println(g3.Debug())
		}
	}
	// g2.Debug()
	// for _, fold := range g.Folds {
	// 	fmt.Println()
	// 	g.Fold(fold)
	// 	fmt.Println(g.Debug())
	// }
}

// // rename this to the year and day in question
// func (app *Application) Y2021D13P2() {
// }

// rename and uncomment this to the year and day in question once complete for a gold star!
// func (app *Application) Y20XXDXXP1Render() {
// }

// rename and uncomment this to the year and day in question once complete for a gold star!
// func (app *Application) Y20XXDXXP2Render() {
// }

// this is what we will reflect and call - so both parts with run. It's up to you to make it print nicely etc.
// The app reference has a CLI for logging.
// func (app *Application) Y2021D13() {
// 	app.Y2021D13P1()
// 	app.Y2021D13P2()
// }
