package d14

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/goutils"
)

/*
--- Day 05:  ---

*/

type BlockType string

const ROCK BlockType = "#"
const AIR BlockType = "."
const SAND BlockType = "o"
const SOURCE BlockType = "+"
const ORIGIN BlockType = "+"

type Block struct {
	BlockType BlockType
	X         int
	Y         int
}

func (b *Block) String() string {
	return string(b.BlockType)
}

func NewBlock(x int, y int, blockType BlockType) *Block {
	b := Block{X: x, Y: y, BlockType: blockType}
	return &b
}

type Puzzle struct {
	title    string
	year     string
	day      string
	input    string
	Lines    []string
	Rows     int
	Cols     int
	Grid     map[string]*Block
	Origin   *Block
	HasFloor bool
	Floor    int
}

func NewPuzzleWithData(VERBOSE bool, input string) *Puzzle {
	p := Puzzle{year: "2022", day: "14", title: "Regolith Reservoir", Cols: 0, Rows: 0, Grid: make(map[string]*Block)}
	p.Load(VERBOSE, input)
	p.Origin = p.Put(500, 0, ORIGIN)
	return &p
}

func (puzzle *Puzzle) SetFloor(floor int) {
	puzzle.HasFloor = true
	puzzle.Floor = floor
}

func NewPuzzle() *Puzzle {
	return NewPuzzleWithData(false, REAL_DATA)
}

func (p *Puzzle) Put(x int, y int, value BlockType) *Block {
	b := NewBlock(x, y, value)
	key := fmt.Sprintf("%v_%v", x, y)
	p.Grid[key] = b
	fmt.Printf("PUT (%v,%v)\n", x, y)
	return b
}

func (p *Puzzle) Get(x int, y int) *Block {
	key := fmt.Sprintf("%v_%v", x, y)
	if p.HasFloor && y >= p.Floor {
		return NewBlock(x, p.Floor, ROCK)
	}
	return p.Grid[key]
}

func (p *Puzzle) Contains(x int, y int) bool {
	key := fmt.Sprintf("%v_%v", x, y)
	if p.HasFloor && y >= p.Floor {
		return true
	}
	return p.Grid[key] != nil
}

func (p *Puzzle) Debug() string {
	return p.DebugFrame(0, 0, p.Cols, p.Rows)
}

func (p *Puzzle) Bounds() (int, int, int, int) {
	min_x := 1000000
	min_y := 1000000
	max_x := 0
	max_y := 0
	for _, entry := range p.Grid {
		min_x = goutils.Min(min_x, entry.X)
		min_y = goutils.Min(min_y, entry.Y)
		max_x = goutils.Max(max_x, entry.X)
		max_y = goutils.Max(max_y, entry.Y)
	}
	return min_x, min_y, max_x, max_y
}

func (p *Puzzle) DebugFrame(min_x int, min_y int, max_x int, max_y int) string {
	result := ""
	for y := min_y; y < max_y; y++ {
		row := ""
		for x := min_x; x < max_x; x++ {
			block := p.Get(x, y)
			if block == nil {
				row = fmt.Sprintf("%v%v", row, AIR)
			} else {
				row = fmt.Sprintf("%v%v", row, block)
			}
		}
		if y > 0 {
			result = fmt.Sprintf("%v\n%v", result, row)
		} else {
			result = fmt.Sprintf("%v%v", result, row)
		}
	}
	return result
}

func (p *Puzzle) Size() int {
	return len(p.Grid)

}

func (puzzle *Puzzle) Load(VERBOSE bool, input string) {
	lines := strings.Split(input, "\n")
	puzzle.input = input
	paths := make([]string, 0)
	for _, line := range lines {
		if VERBOSE {
			fmt.Printf("line: '%v'\n", line)
		}
		splits := strings.Split(line, " -> ")
		if VERBOSE {
			fmt.Printf("splits: '%v'\n", splits)
		}
		start := splits[0]
		end := ""
		for index := 1; index < len(splits); index++ {
			// make the blocks
			end = splits[index]
			if VERBOSE {
				fmt.Printf("[%v] start=%v, end=%v\n", index, start, end)
			}
			start_coordinates := strings.Split(start, ",")
			end_coordinates := strings.Split(end, ",")
			if VERBOSE {
				fmt.Printf("start_coords: %v\n", start_coordinates)
				fmt.Printf("end_coords  : %v\n", end_coordinates)
			}
			s_x, _ := strconv.Atoi(start_coordinates[0])
			s_y, _ := strconv.Atoi(start_coordinates[1])
			e_x, _ := strconv.Atoi(end_coordinates[0])
			e_y, _ := strconv.Atoi(end_coordinates[1])
			if VERBOSE {
				fmt.Printf("s_x: %v, s_y: %v, e_x: %v, e_y: %v\n", s_x, s_y, e_x, e_y)
			}

			x_diff := 0
			y_diff := 0
			if s_x < e_x {
				x_diff = 1
			} else if s_x > e_x {
				x_diff = -1
			}

			if s_y < e_y {
				y_diff = 1
			} else if s_y > e_y {
				y_diff = -1
			}

			if VERBOSE {
				fmt.Printf("x: %v -> %v by way of %v\n", s_x, e_x, x_diff)
				fmt.Printf("y: %v -> %v by way of %v\n", s_y, e_y, y_diff)
			}
			xquit := false
			for x := s_x; !xquit; x += x_diff {
				if VERBOSE {
					fmt.Printf("x=%v\n", x)
				}
				yquit := false
				for y := s_y; !yquit; y += y_diff {
					if VERBOSE {
						fmt.Printf("y=%v\n", y)
						fmt.Printf("PUT(%v,%v) %v\n", x, y, ROCK)
					}
					puzzle.Put(x, y, ROCK)
					yquit = y == e_y
				}
				xquit = x == e_x
			}

			puzzle.Rows = goutils.Max(puzzle.Rows, s_y)
			puzzle.Rows = goutils.Max(puzzle.Rows, e_y)
			puzzle.Cols = goutils.Max(puzzle.Cols, s_x)
			puzzle.Cols = goutils.Max(puzzle.Cols, e_x)
			start = end

		}
	}
	puzzle.Lines = paths
}

func (puzzle *Puzzle) Part1() {
	puzzle.Load(false, REAL_DATA)
}

func (puzzle *Puzzle) Part2() {
	puzzle.Load(false, REAL_DATA)
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}

func (puzzle *Puzzle) AddSand(max_y int) (bool, *Block, int) {
	origin := puzzle.Origin
	x := origin.X
	y := origin.Y

	sand := NewBlock(x, y, SAND)

	// if we can move down, move down
	steps := 0
	for {
		steps += 1
		cx := sand.X
		cy := sand.Y + 1
		if !puzzle.Contains(cx, cy) {
			// drop down one
			fmt.Printf("[%v] DOWN (%v,%v)->(%v,%v)\n", steps, sand.X, sand.Y, cx, cy)
			sand.X = cx
			sand.Y = cy
		} else if !puzzle.Contains(cx-1, cy) {
			// drop down one left
			fmt.Printf("[%v] LEFT (%v,%v)->(%v,%v)\n", steps, sand.X, sand.Y, cx-1, cy)
			sand.X = cx - 1
			sand.Y = cy
		} else if !puzzle.Contains(cx+1, cy) {
			// drop down one right
			fmt.Printf("[%v] RGHT (%v,%v)->(%v,%v)\n", steps, sand.X, sand.Y, cx+1, cy)
			sand.X = cx + 1
			sand.Y = cy

		} else if cy > max_y && puzzle.HasFloor && cy > puzzle.Floor {

			fmt.Printf("[%v] cy>max_y (%v,%v), floor=%v\n", steps, cx, max_y, puzzle.Floor)
			return false, sand, steps

		} else {
			// cannot drop in any way
			fmt.Printf("[%v] STAY (%v,%v)\n", steps, sand.X, sand.Y)
			puzzle.Put(sand.X, sand.Y, SAND)
			break
		}

		if steps == 1000 {
			break
		}
	}

	return true, sand, steps

	// if we cannot, move down and left
	// if we cannot, move down and right
}
