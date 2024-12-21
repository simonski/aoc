package d14

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
Day 14: Restroom Redoubt
*/

type Puzzle struct {
	title string
	year  int
	day   int
	input string
	lines []string
}

type Robot struct {
	x  int
	y  int
	vx int
	vy int
}

func (r *Robot) key() string {
	return fmt.Sprintf("%v.%v", r.x, r.y)
}

type Grid struct {
	width  int
	height int
	index  map[string][]*Robot
	robots []*Robot
}

func (g *Grid) add(r *Robot) {
	g.robots = append(g.robots, r)
	bots, exists := g.index[r.key()]
	if exists {
		bots = append(bots, r)
		g.index[r.key()] = bots
	} else {
		bots = make([]*Robot, 0)
		bots = append(bots, r)
		g.index[r.key()] = bots
	}
}

func (g *Grid) safetyFactor() int {
	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0

	for _, r := range g.robots {
		col := r.x
		row := r.y

		if col == g.width/2 || row == g.height/2 {
			continue
		}

		if col < 0 || row < 0 || col > g.width || row > g.height {
			continue
		}

		if row < g.height/2 && col < g.width/2 {
			q1 += 1
		} else if row < g.height/2 && col > g.width/2 {
			q2 += 1
		} else if row > g.height/2 && col < g.width/2 {
			q3 += 1
		} else if row > g.height/2 && col > g.width/2 {
			q4 += 1
		}

	}
	fmt.Printf("%v,%v,%v,%v\n", q1, q2, q3, q4)
	return q1 * q2 * q3 * q4
}

func (g *Grid) debug() string {
	line := ""
	for row := 0; row < g.height; row++ {
		for col := 0; col < g.width; col++ {
			robots := g.get(col, row)

			if row == g.height/2 {
				line += "-"
			} else if col == g.width/2 {
				line += "|"
			} else if robots == nil {
				line += "."
			} else {
				line += fmt.Sprintf("%v", len(robots))
			}
		}
		line += "\n"
	}
	return line
}

func NewGrid(width int, height int, data string) *Grid {
	lines := strings.Split(data, "\n")
	g := &Grid{width: width, height: height, robots: make([]*Robot, 0), index: make(map[string][]*Robot)}
	for _, line := range lines {
		// p=0,4 v=3,-3
		line = strings.ReplaceAll(line, "p=", "")
		line = strings.ReplaceAll(line, "v=", "")
		line = strings.ReplaceAll(line, " ", ",")
		splits := utils.SplitDataToListOfInts(line, ",")
		r := &Robot{x: splits[0], y: splits[1], vx: splits[2], vy: splits[3]}
		g.add(r)
	}
	return g

}

func (r *Robot) debug() string {
	return fmt.Sprintf("x=%v, y=%v, vx=%v, vy=%v\n", r.x, r.y, r.vx, r.vy)
}

func (g *Grid) get(col int, row int) []*Robot {
	key := fmt.Sprintf("%v.%v", col, row)
	return g.index[key]
}

func (g *Grid) step() {
	index := make(map[string][]*Robot)
	// fmt.Printf("There are %v robots.\n", len(g.robots))
	for _, r := range g.robots {

		// fmt.Println("before")
		// fmt.Println(r.debug())
		r.x += r.vx
		r.y += r.vy

		if r.x >= g.width {
			r.x = r.x - g.width
		} else if r.x < 0 {
			r.x = g.width + r.x
		}

		if r.y >= g.height {
			r.y = r.y - g.height
		} else if r.y < 0 {
			r.y = g.height + r.y
		}
		// fmt.Println("after")
		// fmt.Println(r.debug())
		// fmt.Println("")

		bots, exists := index[r.key()]
		if exists {
			bots = append(bots, r)
			index[r.key()] = bots
		} else {
			bots = make([]*Robot, 0)
			bots = append(bots, r)
			index[r.key()] = bots
		}
	}
	g.index = index
}

func (puzzle *Puzzle) GetSummary() utils.Summary {
	s := utils.Summary{Day: puzzle.day, Year: puzzle.year, Name: puzzle.title}
	s.ProgressP1 = utils.NotStarted
	s.ProgressP2 = utils.NotStarted
	s.DateStarted = "2024-12-21 16:02:05"
	return s
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("2024")
	iday, _ := strconv.Atoi("14")
	p := Puzzle{year: iyear, day: iday, title: "Day 14: Restroom Redoubt"}
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
	// g := NewGrid(11, 7, TEST_DATA)
	g := NewGrid(101, 103, REAL_DATA)
	for count := 0; count < 100; count++ {
		g.step()
		fmt.Printf("After %v seconds.\n", count+1)
		fmt.Println(g.debug())
		fmt.Println()
	}
	// 223319536 too high
	// 218619120
	fmt.Printf("Safety factor is %v\n", g.safetyFactor())

}

func (g *Grid) find(distance int) (bool, int) {
	col1 := g.width/2 - distance
	col2 := g.width/2 + distance
	for row := 0; row < g.height; row++ {
		r1 := g.get(col1, row)
		r2 := g.get(col2, row)
		if len(r1) > 0 && len(r2) > 0 {
			fmt.Printf("%v,%v && %v,%v = true\n", col1, row, col2, row)
			return true, row
		}
	}
	return false, -1
}

func (puzzle *Puzzle) Part2X() {
	// g := NewGrid(11, 7, TEST_DATA)
	g := NewGrid(101, 103, REAL_DATA)
	count := 0
	seconds := 0
	for {
		g.step()

		ok1, r1 := g.find(2)
		ok2, r2 := g.find(3)
		ok3, r3 := g.find(4)
		ok4, r4 := g.find(5)

		if ok1 && ok2 && ok3 && ok4 && r1 < r2 && r2 < r3 && r3 < r4 {

			// if g.find(1) && g.find(2) && g.find(3) && g.find(4) && g.find(5) {
			fmt.Printf("After %v seconds.\n", seconds+1)
			fmt.Println(g.debug())
			fmt.Println()
			count++
		}
		if count == 10 {
			break
		}
		seconds += 1
	}
	// 223319536 too high
	// 218619120
	// fmt.Printf("Safety factor is %v\n", g.safetyFactor())

}

func (g *Grid) countRobotsInRow(row int) int {
	count := 0
	for col := 0; col < g.width; col++ {
		l := g.get(col, row)
		if len(l) > 0 {
			count += 1
		}
	}
	return count
}
func (puzzle *Puzzle) Part2() {
	// g := NewGrid(11, 7, TEST_DATA)
	g := NewGrid(101, 103, REAL_DATA)
	count := 0
	seconds := 0
	for {
		g.step()

		for row := 0; row < g.height; row++ {
			rows := g.countRobotsInRow(row)
			if rows > 30 {

				// if g.find(1) && g.find(2) && g.find(3) && g.find(4) && g.find(5) {
				fmt.Printf("After %v seconds.\n", seconds+1)
				fmt.Println(g.debug())
				fmt.Println()
				count += 1
				break
			}
		}
		if count == 10 {
			break
		}
		seconds += 1
	}
	// 223319536 too high
	// 218619120
	// fmt.Printf("Safety factor is %v\n", g.safetyFactor())

}

func (puzzle *Puzzle) Run() {
	// puzzle.Part1()
	puzzle.Part2()
}
