package d15

import (
	"fmt"
	"os"
	"strings"
)

/*
--- Day 05:  ---

*/

type Puzzle struct {
	title string
	year  string
	day   string
	input string
	lines []string
}

func NewPuzzleWithData(input string) *Puzzle {
	p := Puzzle{year: "2022", day: "15", title: "Beacon Exclusion Zone"}
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

func (puzzle *Puzzle) Part1() {
	// not 5129731 (too low)
	//     5129730
	//     5129730
	//     5129730

	//     5870799
	// 5870800
	// 174,580,378
	// data := REAL_DATA
	// row := 2000000
	data := TEST_DATA
	row := 11

	g := NewGrid(data)
	// minx, miny, maxx, maxy := g.Bounds()
	// width := g.Width()
	// height := g.Height()
	// fmt.Printf("Grid bounds (%v,%v,%v,%v)\nHeight=%v\nWidth=%v\nVolume=%v\nSensors=%v\nBeacons=%v\n", minx, miny, maxx, maxy, height, width, height*width, len(g.Sensors), len(g.Beacons))
	could_be_beacon, could_not_be_beacon, result, min_x := g.CountCannotsForRow_V2(row)
	debug := os.Args[4] == "Y"
	if debug {
		line := ""
		for col := 0; col < len(result); col++ {
			if result[col] {
				line += "#"
			} else {
				line += "."
			}
		}
		fmt.Println(line)
	}
	fmt.Printf("min_x=%v\n", min_x)
	fmt.Printf("2022_15, Part1: could_be_beacon=%v, could_not_be_beacon=%v\n", could_be_beacon,
		could_not_be_beacon+1)

	// at row 2000000, answer is , 5870800
}

func (puzzle *Puzzle) Part2() {
	//	NOT resultrow=2263076, result=[2571519] (giving) 10286078263076
	// 4000000 * 2571519 + 2263076 =

	// 2727057,2916597

	// 2727057 * 4000000 + 2916597 10908230916597

	g := NewGrid(REAL_DATA)
	result := make([]int, 0)
	resultrow := 0
	for row := 0; row < 400000; row++ {
		if row%10000 == 0 {
			fmt.Println(row)
		}
		missing := g.CountMissing(false, row)
		if len(missing) == 1 {
			result = missing
			resultrow = row
			break
		}
	}
	fmt.Printf("resultrow=%v, result=%v\n", resultrow, result)

}

func (puzzle *Puzzle) Part3() {
	//	NOT resultrow=2263076, result=[2571519] (giving) 10286078263076
	// 4000000 * 2571519 + 2263076 =

	g := NewGrid(TEST_DATA)
	result := make([]int, 0)
	resultrow := 0
	for row := -2; row < 22; row++ {
		if row%10000 == 0 {
			fmt.Println(row)
		}
		missing := g.CountMissing(true, row)
		if len(missing) == 1 {
			fmt.Printf("Breaking on row %v\n", row)
			result = missing
			resultrow = row
			break
		}
	}
	fmt.Printf("resultrow=%v, result=%v\n", resultrow, result)

}

func (puzzle *Puzzle) Part1_x() {
	g := NewGrid(TEST_DATA)
	VERBOSE := false
	// segments10 := g.GetSegments(VERBOSE, 10)
	cannot := g.CountCannotBePresent(VERBOSE, 10)
	fmt.Printf("Segments for row %v: %v\n", 10, cannot)

	// cannot = g.CountCannotBePresent(VERBOSE, 11)
	// fmt.Printf("Segments for row %v: %v\n", 11, cannot)

}

func (puzzle *Puzzle) Part1_y() {
	g := NewGrid(REAL_DATA)
	VERBOSE := false
	// segments10 := g.GetSegments(VERBOSE, 10)
	cannot := g.CountCannotBePresent(VERBOSE, 2000000)
	fmt.Printf("Segments for row %v: %v\n", 2000000, cannot)

	// cannot = g.CountCannotBePresent(VERBOSE, 11)
	// fmt.Printf("Segments for row %v: %v\n", 11, cannot)

}

func (puzzle *Puzzle) Part2_x() {
	g := NewGrid(TEST_DATA)
	VERBOSE := true
	// segments10 := g.GetSegments(VERBOSE, 10)
	answer_row := -1
	answer_col := -1
	VERBOSE = false
	for row := 0; row < 20; row++ {
		gaps, col := g.CountGaps(VERBOSE, row)
		fmt.Printf("Gaps for row[%v]=%v, col=%v\n", row, gaps, col)
		if gaps == 1 && col > -1 {
			if g.Get(col, row) == nil {
				answer_row = row
				answer_col = col
			}
		}
	}
	fmt.Printf("(%v,%v)", answer_col, answer_row)

	// cannot = g.CountCannotBePresent(VERBOSE, 11)
	// fmt.Printf("Segments for row %v: %v\n", 11, cannot)

}

func (puzzle *Puzzle) Part2_y() {
	g := NewGrid(REAL_DATA)
	VERBOSE := true
	// segments10 := g.GetSegments(VERBOSE, 10)
	answer_row := -1
	answer_col := -1
	VERBOSE = false
	for row := 0; row < 4000000; row++ {
		gaps, col := g.CountGaps(VERBOSE, row)
		fmt.Printf("Gaps for row[%v]=%v, col=%v\n", row, gaps, col)
		if gaps == 1 && col > -1 {
			if g.Get(col, row) == nil {
				answer_row = row
				answer_col = col
			}
		}
	}
	fmt.Printf("(%v,%v)", answer_col, answer_row)

	// cannot = g.CountCannotBePresent(VERBOSE, 11)
	// fmt.Printf("Segments for row %v: %v\n", 11, cannot)

}

func (puzzle *Puzzle) Run() {
	// puzzle.Part1()
	// puzzle.Part2()
	// puzzle.Part3()
	// puzzle.Part1_x()
	// puzzle.Part1_y()

	puzzle.Part2_x()
	puzzle.Part2_y()

}
