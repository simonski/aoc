package d11

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
Day 11: Plutonian Pebbles
*/

type Puzzle struct {
	title string
	year  int
	day   int
	input string
	lines []string
}

func (puzzle *Puzzle) GetSummary() utils.Summary {
	s := utils.Summary{Day: puzzle.day, Year: puzzle.year, Name: puzzle.title}
	s.ProgressP1 = utils.NotStarted
	s.ProgressP2 = utils.NotStarted
	s.DateStarted = "2024-12-11 07:31:41"
	return s
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("2024")
	iday, _ := strconv.Atoi("11")
	p := Puzzle{year: iyear, day: iday, title: "Day 11: Plutonian Pebbles"}
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
	cache := NewCache()
	fmt.Println(cache.calculateTotalStones(TEST_DATA_1, 25))
}

func (puzzle *Puzzle) Part2() {
	cache := NewCache()
	fmt.Println(cache.calculateTotalStones(TEST_DATA_2, 25))

	cache = NewCache()
	fmt.Println(cache.calculateTotalStones(REAL_DATA, 75))

}

func (c *Cache) calculateTotalStones(values_str string, blinks int) int {
	values := utils.SplitDataToListOfInts(values_str, " ")
	total := 0
	for _, v := range values {
		value := c.count(v, blinks)
		total += value
	}
	return total
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}

type Stone struct {
	id     int
	blinks int
}

type Cache struct {
	cache map[*Stone]int
}

func NewCache() *Cache {
	data := make(map[*Stone]int)
	return &Cache{cache: data}
}

func (c *Cache) put(stone *Stone, value int) {
	c.cache[stone] = value
}

func (c *Cache) get(id int, blinks int) int {
	for k, v := range c.cache {
		if k.id == id && k.blinks == blinks {
			return v
		}
	}
	return -1
}

func (c *Cache) count(id int, blinks int) int {
	stone := &Stone{id, blinks}
	result := c.get(id, blinks)
	if result == -1 {
		result = c.walk(id, blinks)
		c.put(stone, result)
	}
	return result
}

func (c *Cache) walk(value int, blinks int) int {
	if blinks == 0 {
		return 1
	}

	value_str := fmt.Sprintf("%v", value)
	if value == 0 {
		return c.walk(1, blinks-1)
	} else if len(value_str)%2 == 0 {
		left_str := value_str[0 : len(value_str)/2]
		right_str := value_str[len(value_str)/2:]
		left, _ := strconv.Atoi(left_str)
		right, _ := strconv.Atoi(right_str)
		return c.count(left, blinks-1) + c.count(right, blinks-1)
	} else {
		return c.count(value*2024, blinks-1)
	}

}
