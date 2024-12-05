package d4

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
Day 4: Ceres Search
*/

type Puzzle struct {
	title  string
	year   int
	day    int
	input  string
	lines  []string
	cache  map[string]bool
	cache2 map[string]bool
}

func (puzzle *Puzzle) GetSummary() utils.Summary {
	s := utils.Summary{Day: puzzle.day, Year: puzzle.year, Name: puzzle.title}
	s.ProgressP1 = utils.NotStarted
	s.ProgressP2 = utils.NotStarted
	s.DateStarted = "2024-12-04 06:28:26"
	return s
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("2024")
	iday, _ := strconv.Atoi("4")
	p := Puzzle{year: iyear, day: iday, title: "Day 4: Ceres Search"}
	p.Load(input)
	p.cache = make(map[string]bool)
	p.cache2 = make(map[string]bool)
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
	puzzle.Load(TEST_DATA)

	/*
		for col in cols
			for row in rows
			if x
				search up
				search down
				search left
				search right
				search diagonal up-left
				search diagonal up-right
				search diagonal down-left-left
				search diagonal down-right

	*/

	count := 0
	rows := len(puzzle.lines)
	cols := len(puzzle.lines[0])

	fmt.Printf("rows=%v, cols=%v\n", rows, cols)
	for col_index := 0; col_index < cols; col_index++ {
		for row_index := 0; row_index < rows; row_index++ {
			cell := puzzle.lines[row_index][col_index : col_index+1]
			// fmt.Printf("(%v, %v)=%v]\n", col_index, row_index, cell)
			if cell == "X" {
				// is x
				// search up

				count += puzzle.searchUp(col_index, row_index)
				count += puzzle.searchDown(col_index, row_index)
				count += puzzle.searchLeft(col_index, row_index)
				count += puzzle.searchRight(col_index, row_index)
				count += puzzle.searchUpLeft(col_index, row_index)
				count += puzzle.searchUpRight(col_index, row_index)
				count += puzzle.searchDownLeft(col_index, row_index)
				count += puzzle.searchDownRight(col_index, row_index)
			}
		}
	}

	fmt.Println(count)
	fmt.Printf("cache size %v\n", len(puzzle.cache))
	keys := make([]string, 0)
	for k := range puzzle.cache {
		// 	fmt.Printf("%v: %v\n", k, v)
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k)
	}

	fmt.Print("\n")
	for row_index := 0; row_index < rows; row_index++ {
		for col_index := 0; col_index < cols; col_index++ {
			key := fmt.Sprintf("%v,%v", col_index, row_index)
			if puzzle.cache2[key] {
				fmt.Printf(puzzle.lines[row_index][col_index : col_index+1])
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}

}

func (puzzle *Puzzle) searchRight(col_index int, row_index int) int {
	word := "X"
	need_word := "XMAS"
	need := "M"
	location := make([]string, 0)
	pos := fmt.Sprintf("%v,%v", col_index, row_index)
	location = append(location, pos)

	row := puzzle.lines[row_index]

	for col_index2 := col_index + 1; col_index2 < len(puzzle.lines[0]); col_index2++ {
		cell := row[col_index2 : col_index2+1]
		pos := fmt.Sprintf("%v,%v", col_index2, row_index)
		if cell == need {
			location = append(location, pos)
			word += cell
			if word == need_word {
				// now check - do we already have this?
				sort.Strings(location)
				key := fmt.Sprintf("%v", location)
				_, exists := puzzle.cache[key]
				if !exists {
					for _, k := range location {
						fmt.Printf("cache2 (%v)\n", k)
						puzzle.cache2[k] = true
					}
					puzzle.cache[key] = true
					return 1
				} else {
					return 0
				}
			}
			need = need_word[len(word) : len(word)+1]
		} else {
			break
		}
	}
	return 0
}

func (puzzle *Puzzle) searchLeft(col_index int, row_index int) int {
	word := "X"
	need_word := "XMAS"
	need := "M"
	location := make([]string, 0)
	pos := fmt.Sprintf("%v,%v", col_index, row_index)
	location = append(location, pos)

	row := puzzle.lines[row_index]
	// cell := row[col_index : col_index+1]
	// fmt.Printf("start %v, value=%v, need=%v, word=%v\n", pos, cell, need, word)

	for col_index2 := col_index - 1; col_index2 >= 0; col_index2-- {
		cell := row[col_index2 : col_index2+1]
		pos := fmt.Sprintf("%v,%v", col_index2, row_index)
		// fmt.Printf("    %v, value=%v, need=%v, word=%v\n", pos, cell, need, word)
		if cell == need {
			location = append(location, pos)
			word += cell
			if word == need_word {
				// now check - do we already have this?
				sort.Strings(location)
				key := fmt.Sprintf("%v", location)
				_, exists := puzzle.cache[key]
				if !exists {
					for _, k := range location {
						fmt.Printf("cache2 (%v)\n", k)
						puzzle.cache2[k] = true
					}
					puzzle.cache[key] = true
					return 1
				} else {
					return 0
				}
			}
			need = need_word[len(word) : len(word)+1]
		} else {
			break
		}
	}
	return 0
}

func (puzzle *Puzzle) searchUp(col_index int, row_index int) int {
	word := "X"
	need_word := "XMAS"
	need := "M"
	location := make([]string, 0)
	pos := fmt.Sprintf("%v,%v", col_index, row_index)
	location = append(location, pos)

	row := puzzle.lines[row_index]
	cell := row[col_index : col_index+1]
	fmt.Printf("start %v, value=%v, need=%v, word=%v\n", pos, cell, need, word)

	for row_index2 := row_index - 1; row_index2 >= 0; row_index2-- {
		row := puzzle.lines[row_index2]
		cell := row[col_index : col_index+1]
		pos := fmt.Sprintf("%v,%v", col_index, row_index2)
		fmt.Printf("    %v, value=%v, need=%v, word=%v\n", pos, cell, need, word)
		if cell == need {
			location = append(location, pos)
			word += cell
			if word == need_word {
				// now check - do we already have this?
				sort.Strings(location)
				key := fmt.Sprintf("%v", location)
				_, exists := puzzle.cache[key]
				if !exists {
					for _, k := range location {
						fmt.Printf("cache2 (%v)\n", k)
						puzzle.cache2[k] = true
					}
					puzzle.cache[key] = true
					return 1
				} else {
					return 0
				}
			}
			need = need_word[len(word) : len(word)+1]
		} else {
			break
		}
	}
	return 0
}

func (puzzle *Puzzle) searchDown(col_index int, row_index int) int {
	word := "X"
	need_word := "XMAS"
	need := "M"
	location := make([]string, 0)
	pos := fmt.Sprintf("%v,%v", col_index, row_index)
	location = append(location, pos)

	row := puzzle.lines[row_index]
	cell := row[col_index : col_index+1]
	fmt.Printf("start %v, value=%v, need=%v, word=%v\n", pos, cell, need, word)

	for row_index2 := row_index + 1; row_index2 < len(puzzle.lines); row_index2++ {
		row := puzzle.lines[row_index2]
		cell := row[col_index : col_index+1]
		pos := fmt.Sprintf("%v,%v", col_index, row_index2)
		fmt.Printf("    %v, value=%v, need=%v, word=%v\n", pos, cell, need, word)
		if cell == need {
			location = append(location, pos)
			word += cell
			if word == need_word {
				// now check - do we already have this?
				sort.Strings(location)
				key := fmt.Sprintf("%v", location)
				_, exists := puzzle.cache[key]
				if !exists {
					for _, k := range location {
						fmt.Printf("cache2 (%v)\n", k)
						puzzle.cache2[k] = true
					}
					puzzle.cache[key] = true
					return 1
				} else {
					return 0
				}
			}
			need = need_word[len(word) : len(word)+1]
		} else {
			break
		}
	}
	return 0
}

// func (puzzle *Puzzle) searchUp(row_index int, col_index int) int {
// 	word := "X"
// 	need_word := "XMAS"
// 	need := "M"
// 	cache := puzzle.cache
// 	location := make([]string, 0)
// 	pos := fmt.Sprintf("%v,%v", row_index, col_index)
// 	location = append(location, pos)
// 	for row_index2 := row_index - 1; row_index2 >= 0; row_index2-- {
// 		cell := puzzle.lines[row_index2][col_index : col_index+1]
// 		pos := fmt.Sprintf("%v,%v", row_index2, col_index)
// 		if cell == need {
// 			location = append(location, pos)
// 			word += cell
// 			if word == need_word {
// 				// now check - do we already have this?
// 				sort.Strings(location)
// 				key := fmt.Sprintf("%v", location)
// 				_, exists := cache[key]
// 				if !exists {
// 					for _, k := range location {
// 						puzzle.cache2[k] = true
// 					}
// 					cache[key] = true
// 					return 1
// 				} else {
// 					return 0
// 				}
// 			}
// 			need = need_word[len(word) : len(word)+1]
// 		}
// 	}
// 	return 0
// }

func (puzzle *Puzzle) searchUpLeft(col_index int, row_index int) int {
	word := "X"
	need_word := "XMAS"
	need := "M"
	cache := puzzle.cache
	location := make([]string, 0)
	row_index2 := row_index
	col_index2 := col_index
	pos := fmt.Sprintf("%v,%v", col_index, row_index)
	location = append(location, pos)
	for offset := 1; offset < 100; offset++ {
		row_index2 = row_index - offset
		col_index2 = col_index - offset
		if row_index2 < 0 || col_index2 < 0 {
			return 0
		}
		cell := puzzle.lines[row_index2][col_index2 : col_index2+1]
		if cell == need {
			pos := fmt.Sprintf("%v,%v", col_index2, row_index2)
			location = append(location, pos)
			word += cell
			if word == need_word {
				// now check - do we already have this?
				sort.Strings(location)
				key := fmt.Sprintf("%v", location)
				_, exists := cache[key]
				if !exists {
					for _, k := range location {
						puzzle.cache2[k] = true
					}
					cache[key] = true
					return 1
				} else {
					return 0
				}
			}
			need = need_word[len(word) : len(word)+1]
		} else {
			break
		}
	}
	return 0
}

func (puzzle *Puzzle) searchUpRight(col_index int, row_index int) int {
	word := "X"
	need_word := "XMAS"
	need := "M"
	cache := puzzle.cache
	location := make([]string, 0)
	row_index2 := row_index
	col_index2 := col_index
	max_cols := len(puzzle.lines[0]) - 1
	pos := fmt.Sprintf("%v,%v", col_index, row_index)
	location = append(location, pos)
	cell := puzzle.lines[row_index][col_index : col_index+1]
	fmt.Printf("start %v, value=%v, need=%v, word=%v\n", pos, cell, need, word)

	for offset := 1; offset < 100; offset++ {
		row_index2 = row_index - offset
		col_index2 = col_index + offset
		if row_index2 < 0 || col_index2 > max_cols {
			return 0
		}
		cell := puzzle.lines[row_index2][col_index2 : col_index2+1]
		fmt.Printf("    %v, value=%v, need=%v, word=%v\n", pos, cell, need, word)
		if cell == need {
			pos := fmt.Sprintf("%v,%v", col_index2, row_index2)
			location = append(location, pos)
			word += cell
			if word == need_word {
				// now check - do we already have this?
				sort.Strings(location)
				key := fmt.Sprintf("%v", location)
				_, exists := cache[key]
				if !exists {
					for _, k := range location {
						puzzle.cache2[k] = true
					}
					cache[key] = true
					return 1
				} else {
					return 0
				}
			}
			need = need_word[len(word) : len(word)+1]
		} else {
			break
		}
	}
	return 0
}

func (puzzle *Puzzle) searchDownLeft(col_index int, row_index int) int {
	word := "X"
	need_word := "XMAS"
	need := "M"
	row_index2 := row_index
	col_index2 := col_index
	max_rows := len(puzzle.lines) - 1
	cache := puzzle.cache
	location := make([]string, 0)
	pos := fmt.Sprintf("%v,%v", col_index, row_index)
	location = append(location, pos)

	for offset := 1; offset < 100; offset++ {
		row_index2 = row_index + offset
		col_index2 = col_index - offset
		if row_index2 > max_rows || col_index2 < 0 {
			return 0
		}
		cell := puzzle.lines[row_index2][col_index2 : col_index2+1]
		if cell == need {
			pos := fmt.Sprintf("%v,%v", col_index2, row_index2)
			location = append(location, pos)
			word += cell
			if word == need_word {
				// now check - do we already have this?
				sort.Strings(location)
				key := fmt.Sprintf("%v", location)
				_, exists := cache[key]
				if !exists {
					for _, k := range location {
						puzzle.cache2[k] = true
					}
					cache[key] = true
					return 1
				} else {
					return 0
				}
			}
			need = need_word[len(word) : len(word)+1]
		} else {
			break
		}
	}
	return 0
}

func (puzzle *Puzzle) searchDownRight(col_index int, row_index int) int {
	word := "X"
	need_word := "XMAS"
	need := "M"
	row_index2 := row_index
	col_index2 := col_index
	max_rows := len(puzzle.lines) - 1
	max_cols := len(puzzle.lines[0]) - 1
	cache := puzzle.cache
	location := make([]string, 0)
	pos := fmt.Sprintf("%v,%v", col_index, row_index)
	location = append(location, pos)

	for offset := 1; offset < 100; offset++ {
		row_index2 = row_index + offset
		col_index2 = col_index + offset
		if row_index2 > max_rows || col_index2 > max_cols {
			return 0
		}
		cell := puzzle.lines[row_index2][col_index2 : col_index2+1]
		if cell == need {
			pos := fmt.Sprintf("%v,%v", col_index2, row_index2)
			location = append(location, pos)
			word += cell
			if word == need_word {
				// now check - do we already have this?
				sort.Strings(location)
				key := fmt.Sprintf("%v", location)
				_, exists := cache[key]
				if !exists {
					for _, k := range location {
						puzzle.cache2[k] = true
					}
					cache[key] = true
					return 1
				} else {
					return 0
				}
			}
			need = need_word[len(word) : len(word)+1]
		} else {
			break
		}
	}
	return 0
}

// func (puzzle *Puzzle) searchDown(row_index int, col_index int) int {
// 	// cell := puzzle.lines[row_index][col_index:1]
// 	word := "X"
// 	need_word := "XMAS"
// 	need := "M"
// 	cache := puzzle.cache
// 	location := make([]string, 0)
// 	pos := fmt.Sprintf("%v,%v", row_index, col_index)
// 	location = append(location, pos)

// 	for row_index2 := row_index + 1; row_index2 < len(puzzle.lines); row_index2++ {
// 		cell := puzzle.lines[row_index2][col_index : col_index+1]
// 		if cell == need {
// 			pos := fmt.Sprintf("%v,%v", row_index2, col_index)
// 			location = append(location, pos)
// 			word += cell
// 			if word == need_word {
// 				// now check - do we already have this?
// 				sort.Strings(location)
// 				key := fmt.Sprintf("%v", location)
// 				_, exists := cache[key]
// 				if !exists {
// 					for _, k := range location {
// 						puzzle.cache2[k] = true
// 					}
// 					cache[key] = true
// 					return 1
// 				} else {
// 					return 0
// 				}
// 			}
// 			need = need_word[len(word) : len(word)+1]
// 		}
// 	}
// 	return 0
// }

// func (puzzle *Puzzle) searchLeft(row_index int, col_index int) int {
// 	// cell := puzzle.lines[row_index][col_index:1]
// 	word := "X"
// 	need_word := "XMAS"
// 	need := "M"
// 	cache := puzzle.cache
// 	location := make([]string, 0)
// 	pos := fmt.Sprintf("%v,%v", row_index, col_index)
// 	location = append(location, pos)

// 	for col_index2 := col_index - 1; col_index2 >= 0; col_index2-- {
// 		cell := puzzle.lines[row_index][col_index2 : col_index2+1]
// 		if cell == need {
// 			pos := fmt.Sprintf("%v,%v", row_index, col_index2)
// 			location = append(location, pos)
// 			word += cell
// 			if word == need_word {
// 				// now check - do we already have this?
// 				sort.Strings(location)
// 				key := fmt.Sprintf("%v", location)
// 				_, exists := cache[key]
// 				if !exists {
// 					for _, k := range location {
// 						puzzle.cache2[k] = true
// 					}
// 					cache[key] = true
// 					return 1
// 				} else {
// 					return 0
// 				}
// 			}
// 			need = need_word[len(word) : len(word)+1]
// 		}
// 	}
// 	return 0
// }

func (puzzle *Puzzle) Part2() {
	puzzle.Load(REAL_DATA)
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}
