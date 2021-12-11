package aoc2021

import (
	"fmt"
	"strconv"
	"strings"
)

/*
--- Day 11: Dumbo Octopus ---
You enter a large cavern full of rare bioluminescent dumbo octopuses! They seem to not like the Christmas lights on your submarine, so you turn them off for now.

There are 100 octopuses arranged neatly in a 10 by 10 grid. Each octopus slowly gains energy over time and flashes brightly for a moment when its energy is full. Although your lights are off, maybe you could navigate through the cave without disturbing the octopuses if you could predict when the flashes of light will happen.

Each octopus has an energy level - your submarine can remotely measure the energy level of each octopus (your puzzle input). For example:

5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526
The energy level of each octopus is a value between 0 and 9. Here, the top-left octopus has an energy level of 5, the bottom-right one has an energy level of 6, and so on.

You can model the energy levels and flashes of light in steps. During a single step, the following occurs:

First, the energy level of each octopus increases by 1.
Then, any octopus with an energy level greater than 9 flashes. This increases the energy level of all adjacent octopuses by 1, including octopuses that are diagonally adjacent. If this causes an octopus to have an energy level greater than 9, it also flashes. This process continues as long as new octopuses keep having their energy level increased beyond 9. (An octopus can only flash at most once per step.)
Finally, any octopus that flashed during this step has its energy level set to 0, as it used all of its energy to flash.
Adjacent flashes can cause an octopus to flash on a step even if it begins that step with very little energy. Consider the middle octopus with 1 energy in this situation:

Before any steps:
11111
19991
19191
19991
11111

After step 1:
34543
40004
50005
40004
34543

After step 2:
45654
51115
61116
51115
45654
An octopus is highlighted when it flashed during the given step.

Here is how the larger example above progresses:

Before any steps:
5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526

After step 1:
6594254334
3856965822
6375667284
7252447257
7468496589
5278635756
3287952832
7993992245
5957959665
6394862637

After step 2:
8807476555
5089087054
8597889608
8485769600
8700908800
6600088989
6800005943
0000007456
9000000876
8700006848

After step 3:
0050900866
8500800575
9900000039
9700000041
9935080063
7712300000
7911250009
2211130000
0421125000
0021119000

After step 4:
2263031977
0923031697
0032221150
0041111163
0076191174
0053411122
0042361120
5532241122
1532247211
1132230211

After step 5:
4484144000
2044144000
2253333493
1152333274
1187303285
1164633233
1153472231
6643352233
2643358322
2243341322

After step 6:
5595255111
3155255222
3364444605
2263444496
2298414396
2275744344
2264583342
7754463344
3754469433
3354452433

After step 7:
6707366222
4377366333
4475555827
3496655709
3500625609
3509955566
3486694453
8865585555
4865580644
4465574644

After step 8:
7818477333
5488477444
5697666949
4608766830
4734946730
4740097688
6900007564
0000009666
8000004755
6800007755

After step 9:
9060000644
7800000976
6900000080
5840000082
5858000093
6962400000
8021250009
2221130009
9111128097
7911119976

After step 10:
0481112976
0031112009
0041112504
0081111406
0099111306
0093511233
0442361130
5532252350
0532250600
0032240000
After step 10, there have been a total of 204 flashes. Fast forwarding, here is the same configuration every 10 steps:

After step 20:
3936556452
5686556806
4496555690
4448655580
4456865570
5680086577
7000009896
0000000344
6000000364
4600009543

After step 30:
0643334118
4253334611
3374333458
2225333337
2229333338
2276733333
2754574565
5544458511
9444447111
7944446119

After step 40:
6211111981
0421111119
0042111115
0003111115
0003111116
0065611111
0532351111
3322234597
2222222976
2222222762

After step 50:
9655556447
4865556805
4486555690
4458655580
4574865570
5700086566
6000009887
8000000533
6800000633
5680000538

After step 60:
2533334200
2743334640
2264333458
2225333337
2225333338
2287833333
3854573455
1854458611
1175447111
1115446111

After step 70:
8211111164
0421111166
0042111114
0004211115
0000211116
0065611111
0532351111
7322235117
5722223475
4572222754

After step 80:
1755555697
5965555609
4486555680
4458655580
4570865570
5700086566
7000008666
0000000990
0000000800
0000000000

After step 90:
7433333522
2643333522
2264333458
2226433337
2222433338
2287833333
2854573333
4854458333
3387779333
3333333333

After step 100:
0397666866
0749766918
0053976933
0004297822
0004229892
0053222877
0532222966
9322228966
7922286866
6789998766
After 100 steps, there have been a total of 1656 flashes.

Given the starting energy levels of the dumbo octopuses in your cavern, simulate 100 steps. How many total flashes are there after 100 steps?
*/

type Day11Point struct {
	X       int
	Y       int
	Value   int
	Grid    *Day11Grid
	Checked bool
}

func (p *Day11Point) Increment() bool {
	p.Value += 1
	return p.Value == 10
}

func (p *Day11Point) IsFlashing() bool {
	return p.Value > 9
}

type Day11Grid struct {
	Points    [][]*Day11Point
	Width     int
	Height    int
	StepCount int
	Flashing  []*Day11Point
}

func NewDay11Grid(data string) *Day11Grid {
	lines := strings.Split(data, "\n")
	y := 0
	grid := &Day11Grid{}
	for _, line := range lines {
		line = strings.Trim(line, " ")
		row := make([]*Day11Point, 0)
		for x := 0; x < len(line); x++ {
			value, _ := strconv.Atoi(line[x : x+1])
			point := &Day11Point{X: x, Y: y, Grid: grid, Value: value}
			row = append(row, point)
		}
		grid.Points = append(grid.Points, row)
		y += 1
	}
	grid.Height = y
	grid.Width = len(lines[0])
	return grid
}

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func (g *Day11Grid) Debug() string {
	line := ""
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			p := g.GetPoint(x, y)
			if p.IsFlashing() {
				// line = fmt.Sprintf("%v %v%v%v", line, White, 0, Reset)
				line = fmt.Sprintf("%v %v", line, 0)
			} else {
				line = fmt.Sprintf("%v %v", line, p.Value)
			}
		}
		line += "\n"
	}
	return line
}

func (g *Day11Grid) NonFlashingNeighbours(p *Day11Point) []*Day11Point {
	neighbours := make([]*Day11Point, 0)
	for nx := p.X - 1; nx <= p.X+1; nx++ {
		for ny := p.Y - 1; ny <= p.Y+1; ny++ {
			np := g.GetPoint(nx, ny)
			if np == nil {
				continue
			}

			if np == p {
				continue
			}
			if np.IsFlashing() {
				continue
			}
			neighbours = append(neighbours, np)

		}
	}
	return neighbours
}

func (g *Day11Grid) Reset() {
	for _, p := range g.Flashing {
		p.Value = 0
	}

	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			p := g.GetPoint(x, y)
			p.Checked = false
			if p.Value > 9 {
				p.Value = 0
			}
		}
	}
	g.Flashing = make([]*Day11Point, 0)
}

func (g *Day11Grid) GetPoint(x int, y int) *Day11Point {
	if x < 0 || x >= g.Width {
		return nil
	}
	if y < 0 || y >= g.Height {
		return nil
	}
	return g.Points[y][x]
}

// Step should

// 1. icnrement all
// 2. flashwalk
// 3. return number of flashes
func (g *Day11Grid) Step(DEBUG bool) int {
	g.StepCount += 1

	// first, mark all as not flashing
	g.Reset()

	// first, increment everything
	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			p := g.GetPoint(x, y)
			p.Increment()
		}
	}

	// now everything is in incremented, we walk each entry and see if it is flashing; if it is,
	// we increment all neighbours if we can and see if they flash
	flashes := 0
	depth := 0
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			p := g.GetPoint(x, y)
			if !p.Checked && p.IsFlashing() {
				flashes += g.FlashWalk(g.StepCount, depth+1, p, p)
			}
			p.Checked = true
		}
	}

	if DEBUG {
		fmt.Printf("Step[%v] (post flashwalk) total flashes=%v\n%v\n", g.StepCount, flashes, g.Debug())
	}

	return flashes
}

func (g *Day11Grid) FlashWalk(step int, depth int, p *Day11Point, originPoint *Day11Point) int {
	if !p.IsFlashing() {
		return 0
	}
	p.Checked = true

	count := 1
	neighbours := g.NonFlashingNeighbours(p)
	for _, np := range neighbours {
		if np == p || np == originPoint {
			continue
		}
		causedFlash := np.Increment()
		if causedFlash {
			count += g.FlashWalk(step, depth+1, np, originPoint)
		}
	}
	return count

}

// rename this to the year and day in question
func (app *Application) Y2021D11P1() {
	DEBUG := true
	data := DAY_2021_11_TEST_DATA_1
	grid := NewDay11Grid(data)
	fmt.Printf("[0]\n%v\n", grid.Debug())
	grid.Step(DEBUG)
	fmt.Printf("[1]\n%v\n", grid.Debug())
	grid.Step(DEBUG)
	fmt.Printf("[2]\n%v\n", grid.Debug())

	data1 := DAY_2021_11_TEST_DATA
	grid1 := NewDay11Grid(data1)
	fmt.Printf("[0]\n%v\n", grid1.Debug())
	total1 := 0
	for x := 1; x <= 100; x++ {
		total1 += grid1.Step(DEBUG)
		// fmt.Printf("[%v]\n%v\n", x, grid1.Debug())
	}
	fmt.Printf("Total of %v flashes.\n", total1)

	data2 := DAY_2021_11_DATA
	grid2 := NewDay11Grid(data2)
	fmt.Printf("[0]\n%v\n", grid2.Debug())
	total2 := 0
	for x := 1; x <= 100; x++ {
		total2 += grid2.Step(DEBUG)
	}
	fmt.Printf("Total of %v flashes.\n", total2)

}

/*
--- Part Two ---
It seems like the individual flashes aren't bright enough to navigate. However, you might have a better option: the flashes seem to be synchronizing!

In the example above, the first time all octopuses flash simultaneously is step 195:


*/
func (app *Application) Y2021D11P2() {

	data1 := DAY_2021_11_TEST_DATA
	grid1 := NewDay11Grid(data1)
	fmt.Printf("[0]\n%v\n", grid1.Debug())
	for x := 1; x <= 192; x++ {
		grid1.Step(false)
	}

	grid1.Step(true)
	grid1.Step(true)
	grid1.Step(true)

	data2 := DAY_2021_11_DATA
	grid2 := NewDay11Grid(data2)
	for {
		flashes := grid2.Step(false)
		if flashes == 100 {
			fmt.Printf("100 flashes at step %v.\n", grid2.StepCount)
			break
		}
	}

}

// rename and uncomment this to the year and day in question once complete for a gold star!
// func (app *Application) Y20XXDXXP1Render() {
// }

// rename and uncomment this to the year and day in question once complete for a gold star!
// func (app *Application) Y20XXDXXP2Render() {
// }

// this is what we will reflect and call - so both parts with run. It's up to you to make it print nicely etc.
// The app reference has a CLI for logging.
// func (app *Application) Y2021D11() {
// 	app.Y2021D11P1()
// 	app.Y2021D11P2()
// }
