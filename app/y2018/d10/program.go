package d10

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
	"github.com/simonski/cli"
)

/*
The Stars Align
*/

type Puzzle struct {
	title     string
	year      int
	day       int
	input     string
	lines     []string
	HasRender bool
	HasWWW    bool
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("2018")
	iday, _ := strconv.Atoi("10")
	p := Puzzle{year: iyear, day: iday, title: "The Stars Align"}
	p.Load(input)
	return &p
}

func NewPuzzle() *Puzzle {
	return NewPuzzleWithData(REAL_DATA)
}

func (puzzle *Puzzle) Load(input string) {
	lines := strings.Split(input, "\n")
	puzzle.input = input
	puzzle.lines = lines
}

func (puzzle *Puzzle) GetSummary() utils.Summary {
	s := utils.Summary{Day: puzzle.day, Year: puzzle.year, Name: puzzle.title}
	s.ProgressP1 = utils.Completed
	s.ProgressP2 = utils.Completed
	s.DateStarted = "d10"
	return s
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}

func (puzzle *Puzzle) RunWWW() string {
	return ""
}

func (puzzle *Puzzle) Render(cli *cli.CLI) {
}

// rename this to the year and day in question
func (puzzle *Puzzle) Part1() {
	lines := strings.Split(REAL_DATA, "\n")
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
func (puzzle *Puzzle) Part2() {

}
