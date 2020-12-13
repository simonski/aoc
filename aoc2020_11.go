package main

/*

https://adventofcode.com/2020/day/11

Advent of Code[About][Events][Shop][Settings][Log Out]simonski 20*
   $year=2020;[Calendar][AoC++][Sponsors][Leaderboard][Stats]
Our sponsors help make Advent of Code possible:
Replit - Use our free AoC templates. Code online with no setup. Easy!
--- Day 11: Seating System ---
Your plane lands with plenty of time to spare. The final leg of your journey is a ferry that goes directly to the tropical island where you can finally start your vacation. As you reach the waiting area to board the ferry, you realize you're so early, nobody else has even arrived yet!

By modeling the process people use to choose (or abandon) their seat in the waiting area, you're pretty sure you can predict the best place to sit. You make a quick map of the seat layout (your puzzle input).

The seat layout fits neatly on a grid. Each position is either floor (.), an empty seat (L), or an occupied seat (#). For example, the initial seat layout might look like this:

L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
Now, you just need to model the people who will be arriving shortly. Fortunately, people are entirely predictable and always follow a simple set of rules. All decisions are based on the number of occupied seats adjacent to a given seat (one of the eight positions immediately up, down, left, right, or diagonal from the seat). The following rules are applied to every seat simultaneously:

If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
Otherwise, the seat's state does not change.
Floor (.) never changes; seats don't move, and nobody sits on the floor.

After one round of these rules, every seat in the example layout becomes occupied:

#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##
After a second round, the seats with four or more occupied adjacent seats become empty again:

#.LL.L#.##
#LLLLLL.L#
L.L.L..L..
#LLL.LL.L#
#.LL.LL.LL
#.LLLL#.##
..L.L.....
#LLLLLLLL#
#.LLLLLL.L
#.#LLLL.##
This process continues for three more rounds:

#.##.L#.##
#L###LL.L#
L.#.#..#..
#L##.##.L#
#.##.LL.LL
#.###L#.##
..#.#.....
#L######L#
#.LL###L.L
#.#L###.##
#.#L.L#.##
#LLL#LL.L#
L.L.L..#..
#LLL.##.L#
#.LL.LL.LL
#.LL#L#.##
..L.L.....
#L#LLLL#L#
#.LLLLLL.L
#.#L#L#.##
#.#L.L#.##
#LLL#LL.L#
L.#.L..#..
#L##.##.L#
#.#L.LL.LL
#.#L#L#.##
..L.L.....
#L#L##L#L#
#.LLLLLL.L
#.#L#L#.##
At this point, something interesting happens: the chaos stabilizes and further applications of these rules cause no seats to change state! Once people stop moving around, you count 37 occupied seats.

Simulate your seating area by applying the seating rules repeatedly until no seats change state. How many seats end up occupied?

To begin, get your puzzle input.

Answer:


You can also [Share] this puzzle.

*/
import (
	"fmt"
	"strings"

	goutils "github.com/simonski/goutils"
)

// AOC_2020_11 is the entrypoint
func AOC_2020_11(cli *goutils.CLI) {
	AOC_2020_11_part1_attempt1(cli)
	AOC_2020_11_part2_attempt1(cli)
}

func AOC_2020_11_part1_attempt1(cli *goutils.CLI) {
	filename := cli.GetFileExistsOrDie("-input")
	tolerance := 4
	sp := NewSeatingPlanFromFile(filename, tolerance, false)
	for {
		if sp.Tick() == 0 {
			break
		}
	}
	fmt.Printf("Tick Count is     : %v\n", sp.TickCount)
	fmt.Printf("Occupied Count is : %v\n", sp.GetOccupiedCount())
}

func AOC_2020_11_part2_attempt1(cli *goutils.CLI) {
	filename := cli.GetFileExistsOrDie("-input")
	tolerance := 5
	searchFar := true
	sp := NewSeatingPlanFromFile(filename, tolerance, searchFar)
	for {
		if sp.Tick() == 0 {
			break
		}
	}
	fmt.Printf("Tick Count is     : %v\n", sp.TickCount)
	fmt.Printf("Occupied Count is : %v\n", sp.GetOccupiedCount())
}

// Tick performs one round of the seating change logic, returning the
// number of changes made in this round
func (sp *SeatingPlan) Tick() int {
	changes := 0
	newdata := make([]int, len(sp.data))
	for index, _ := range sp.data {
		newdata[index] = sp.ConvertAtIndex(index)
		if newdata[index] != sp.data[index] {
			changes++
		}

	}
	if changes > 0 {
		sp.data = newdata
	}
	sp.TickCount++
	return changes
}

func (sp *SeatingPlan) ConvertAtIndex(index int) int {
	seat := sp.data[index]
	if seat == FLOOR {
		return FLOOR
	}
	if seat == EMPTY && sp.CountAdjacentOccupied(index) == 0 {
		return OCCUPIED
	}
	if seat == OCCUPIED && sp.CountAdjacentOccupied(index) >= sp.tolerance {
		return EMPTY
	}
	return seat
}

func (sp *SeatingPlan) GetOccupiedCount() int {
	count := 0
	for _, value := range sp.data {
		if value == OCCUPIED {
			count++
		}
	}
	return count
}

// CountOccupied checks occupancy in all directions
func (sp *SeatingPlan) CountAdjacentOccupied(index int) int {
	count := 0

	col, row := sp.ColRowFromIndex(index)
	if !sp.searchFar {
		// then only search 1
		if sp.Get(col-1, row) == OCCUPIED {
			count++
		}
		if sp.Get(col+1, row) == OCCUPIED {
			count++
		}
		if sp.Get(col, row+1) == OCCUPIED {
			count++
		}
		if sp.Get(col, row-1) == OCCUPIED {
			count++
		}

		if sp.Get(col-1, row-1) == OCCUPIED {
			count++
		}
		if sp.Get(col+1, row-1) == OCCUPIED {
			count++
		}
		if sp.Get(col-1, row+1) == OCCUPIED {
			count++
		}
		if sp.Get(col+1, row+1) == OCCUPIED {
			count++
		}

	} else {

		// left
		// first you can see left is
		for test_col := col - 1; test_col >= 0; test_col-- {
			value := sp.Get(test_col, row)
			if value == OCCUPIED {
				count++
				break
			} else if value == EMPTY {
				break
			}
		}

		// right
		for test_col := col + 1; test_col < sp.width; test_col++ {
			value := sp.Get(test_col, row)
			if value == OCCUPIED {
				count++
				break
			} else if value == EMPTY {
				break
			}
		}

		// up
		for test_row := row - 1; test_row >= 0; test_row-- {
			value := sp.Get(col, test_row)
			if value == OCCUPIED {
				count++
				break
			} else if value == EMPTY {
				break
			}
		}

		// down
		for test_row := row + 1; test_row < sp.depth; test_row++ {
			value := sp.Get(col, test_row)
			if value == OCCUPIED {
				count++
				break
			} else if value == EMPTY {
				break
			}
		}

		// up left
		for distance := 1; ; distance++ {
			test_col := col - distance
			test_row := row - distance
			value := sp.Get(test_col, test_row)
			if value == OCCUPIED {
				count++
				break
			} else if value == NONE {
				break
			} else if value == EMPTY {
				break
			}
		}

		// up right
		for distance := 1; ; distance++ {
			test_col := col + distance
			test_row := row - distance
			value := sp.Get(test_col, test_row)
			if value == OCCUPIED {
				count++
				break
			} else if value == NONE {
				break
			} else if value == EMPTY {
				break
			}
		}

		// down left
		for distance := 1; ; distance++ {
			test_col := col - distance
			test_row := row + distance
			value := sp.Get(test_col, test_row)
			if value == OCCUPIED {
				count++
				break
			} else if value == NONE {
				break
			} else if value == EMPTY {
				break
			}
		}

		// down right
		for distance := 1; ; distance++ {
			test_col := col + distance
			test_row := row + distance
			value := sp.Get(test_col, test_row)
			if value == OCCUPIED {
				count++
				break
			} else if value == NONE {
				break
			} else if value == EMPTY {
				break
			}
		}
	}
	// fmt.Printf("CountAdjacent(%v,%v), distance=%v, count=%v\n", col, row, sp.maxDistance, count)

	return count
}

func (sp *SeatingPlan) Get(col int, row int) int {
	index := sp.IndexFromColRow(col, row)
	if index == -1 || index >= len(sp.data) {
		return NONE
	}
	return sp.data[index]
}

func (sp *SeatingPlan) Put(position int, value int) {
	sp.data[position] = value
}

func (sp *SeatingPlan) ColRowFromIndex(index int) (int, int) {
	if index < 0 || index > len(sp.data) {
		return -1, -1
	}
	// index = (row * sp.width) + col
	row := index / sp.width
	col := index % sp.width
	return col, row
}

func (sp *SeatingPlan) IndexFromColRow(col int, row int) int {
	if row < 0 || row > len(sp.data)/sp.width {
		return -1
	}
	if col < 0 || col >= sp.width {
		return -1
	}
	return row*sp.width + col
}

// func (sp *SeatingPlan) left(position int) int {
// 	col, row := sp.ColRowFromIndex(position)
// 	return sp.Get(col-1, row)
// }

// func (sp *SeatingPlan) right(position int) int {
// 	col, row := sp.ColRowFromIndex(position)
// 	return sp.Get(col+1, row)
// }

// func (sp *SeatingPlan) up(position int) int {
// 	col, row := sp.ColRowFromIndex(position)
// 	return sp.Get(col, row-1)
// }

// func (sp *SeatingPlan) down(position int) int {
// 	col, row := sp.ColRowFromIndex(position)
// 	return sp.Get(col, row+1)
// }

// func (sp *SeatingPlan) upperleft(position int) int {
// 	col, row := sp.ColRowFromIndex(position)
// 	return sp.Get(col-1, row-1)
// }

// func (sp *SeatingPlan) upperright(position int) int {
// 	col, row := sp.ColRowFromIndex(position)
// 	return sp.Get(col+1, row-1)
// }

// func (sp *SeatingPlan) lowerleft(position int) int {
// 	col, row := sp.ColRowFromIndex(position)
// 	return sp.Get(col-1, row+1)
// }

// func (sp *SeatingPlan) lowerright(position int) int {
// 	col, row := sp.ColRowFromIndex(position)
// 	return sp.Get(col+1, row+1)
// }

const NONE = -1    // ?
const OCCUPIED = 0 // #
const EMPTY = 1    // L
const FLOOR = 2    // .

type SeatingPlan struct {
	data      []int
	width     int
	depth     int
	tolerance int
	searchFar bool // indicates if we search as far as we can see
	TickCount int  // the number of ticks that have happened
}

func NewSeatingPlanFromFile(filename string, tolerance int, searchFar bool) *SeatingPlan {
	lines := load_file_to_strings(filename)
	return NewSeatingPlanFromStrings(lines, tolerance, searchFar)
}

func NewSeatingPlanFromStrings(lines []string, tolerance int, searchFar bool) *SeatingPlan {
	data := make([]int, 0)
	for _, line := range lines {
		for index := 0; index < len(line); index++ {
			entry := line[index : index+1]
			if entry == "#" {
				data = append(data, OCCUPIED)
			} else if entry == "L" {
				data = append(data, EMPTY)
			} else if entry == "." {
				data = append(data, FLOOR)
			}
		}
	}
	width := len(lines[0])
	depth := len(lines)
	sp := SeatingPlan{data: data, width: width, depth: depth, tolerance: tolerance, searchFar: searchFar}
	return &sp
}

func (sp *SeatingPlan) Translate(value int) string {
	if value == FLOOR {
		return "."
	} else if value == OCCUPIED {
		return "#"
	} else if value == EMPTY {
		return "L"
	} else {
		return "?"
	}
}

func (sp *SeatingPlan) Debug() string {
	output := ""
	for index, value := range sp.data {
		translated := sp.Translate(value)
		output += translated
		if index > 0 && (index+1)%sp.width == 0 {
			output += "\n"
		}
	}
	return strings.TrimSpace(output)
}
