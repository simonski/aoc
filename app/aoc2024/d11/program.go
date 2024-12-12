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
	puzzle.Load(REAL_DATA)
	puzzle.p1(REAL_DATA)
	// puzzle.p1(REAL_DATA)
}

func (puzzle *Puzzle) p1(data string) {
	values := strings.Split(data, " ")
	fmt.Printf("%v: %v %v\n", 0, len(values), values)
	for count := 0; count < 25; count++ {
		values = p1_blink(values)
		if len(values) < 10 {
			fmt.Printf("%v: %v %v\n", count+1, len(values), values)
		} else {
			fmt.Printf("%v: %v\n", count+1, len(values))
		}
	}
}

func (puzzle *Puzzle) p2(debug bool, data string, rounds int) {
	// each number will be run 75 times
	// some will get there no problem
	// some will generate a split
	// some might not
	// if it generates two numbers
	// do we have that number already
	// if so we can use it

	values := strings.Split(data, " ")

	// walk a number until I get to a split
	// walk the split until you get to a split
	cache := NewCache()
	max_depth := 99999000000
	for _, v := range values {
		entry := cache.get(v)
		cache.walk(entry, debug, 0, max_depth)
	}
	size := cache.size(values, rounds)
	fmt.Printf("Ok, after %v rounds the number of stones will be %v\n", rounds, size)

}

type Entry struct {
	value       string
	childrenMap map[string]*Entry
	children    []*Entry
}

func (e *Entry) addChild(c *Entry) bool {
	if e.hasChild(c) {
		return false
	}
	if c != nil {
		// if c.parent != nil {
		// 	panic("nop")
		// }
		// fmt.Printf("%v.add(%v)\n", e.value, c.value)
		// c.parent = e
		e.childrenMap[c.value] = c
		e.children = append(e.children, c)
		// c.parent = e
		return true
	} else {
		return false
	}
}

func (e *Entry) hasDescendent(candidate *Entry) bool {
	_, exists := e.childrenMap[candidate.value]
	if exists {
		return true
	}
	for _, child := range e.children {
		desc := child.hasDescendent(candidate)
		if desc {
			return true
		}
	}
	return false
}

func (e *Entry) hasChild(candidate *Entry) bool {
	_, exists := e.childrenMap[candidate.value]
	return exists
}

func (e *Entry) contains(c *Entry) bool {
	_, exists := e.childrenMap[c.value]
	return exists
}

type Cache struct {
	cache map[string]*Entry
}

func NewCache() *Cache {
	c := &Cache{}
	c.cache = make(map[string]*Entry)
	return c
}

func (c *Cache) get(key string) *Entry {
	e, exists := c.cache[key]
	if exists {
		return e
	} else {
		e := &Entry{value: key, childrenMap: make(map[string]*Entry), children: make([]*Entry, 0)}
		c.cache[key] = e
		return e
	}
}

func (c *Cache) blink(entry *Entry) (*Entry, *Entry) {
	value := entry.value
	var left *Entry
	var right *Entry
	if value == "0" {
		// rule 1
		value = strings.ReplaceAll(value, "0", "1")
		left = c.get(value)
		// left.parent = entry
	} else if len(value)%2 == 0 {
		// rule 2
		lpart := value[0 : len(value)/2]
		rpart := value[len(value)/2:]
		lpart_int, _ := strconv.Atoi(lpart)
		rpart_int, _ := strconv.Atoi(rpart)
		lpart = fmt.Sprintf("%v", lpart_int)
		rpart = fmt.Sprintf("%v", rpart_int)

		left = c.get(lpart)
		right = c.get(rpart)

	} else {
		// rule 3
		value_int, _ := strconv.Atoi(value)
		value_int *= 2024
		value = fmt.Sprintf("%v", value_int)
		left = c.get(value)
	}
	return left, right
}

func (cache *Cache) size(original_values []string, rounds int) int {
	// so the cache is now populated but we need to
	// the cache holds an entry with a number of children.
	// for each round we can do a single check by iterating once.

	values := make([]string, 0)
	values = append(values, original_values...)
	count := len(values)
	// fmt.Println(values)

	for round := 0; round < rounds; round++ {
		new_values := make([]string, 0)
		for _, value := range values {
			entry := cache.get(value)
			count += len(entry.children) - 1
			for _, child := range entry.children {
				new_values = append(new_values, child.value)
			}
		}
		values = new_values
		// fmt.Println(values)

	}
	return count
}

// the purpose is to learn the size of a number until it repeats
func (c *Cache) walk(entry *Entry, debug bool, depth int, max_depth int) {
	if depth == max_depth {
		return
	}

	prefix := fmt.Sprintf("[%v/%v] (%v) -> ", depth, max_depth, entry.value)
	lentry, rentry := c.blink(entry)

	if lentry != nil && rentry != nil {
		fmt.Printf("%v(%v, %v)\n", prefix, lentry.value, rentry.value)
	} else {
		fmt.Printf("%v(%v)\n", prefix, lentry.value)
	}

	walkLeft := false
	walkRight := false
	if lentry != nil {
		if entry.hasChild(lentry) {
			walkLeft = false
		} else {
			entry.addChild(lentry)
			walkLeft = true
		}
	}

	if rentry != nil {
		if entry.hasChild(rentry) {
			walkRight = false
		} else {
			entry.addChild(rentry)
			walkRight = true
		}
	}

	if walkLeft {
		c.walk(lentry, debug, depth+1, max_depth)
	}

	if walkRight {
		c.walk(rentry, debug, depth+1, max_depth)
	}

	if !walkLeft && !walkRight {
		fmt.Printf("%v complete.\n", prefix)
	}

}

func p1_blink(values []string) []string {
	result := make([]string, 0)
	for index := 0; index < len(values); index++ {
		value := values[index]
		if value == "0" {
			// rule 1
			value = strings.ReplaceAll(value, "0", "1")
			result = append(result, value)
		} else if len(value)%2 == 0 {
			// rule 2
			lpart := value[0 : len(value)/2]
			rpart := value[len(value)/2:]
			lpart_int, _ := strconv.Atoi(lpart)
			rpart_int, _ := strconv.Atoi(rpart)
			lpart = fmt.Sprintf("%v", lpart_int)
			rpart = fmt.Sprintf("%v", rpart_int)
			result = append(result, lpart)
			result = append(result, rpart)
		} else {
			// rule 3
			value_int, _ := strconv.Atoi(value)
			value_int *= 2024
			value = fmt.Sprintf("%v", value_int)
			result = append(result, value)
		}
	}
	return result
}

func blink(value string) (string, string) {
	if value == "0" {
		// rule 1
		value = strings.ReplaceAll(value, "0", "1")
		return value, ""
	} else if len(value)%2 == 0 {
		// rule 2
		lpart := value[0 : len(value)/2]
		rpart := value[len(value)/2:]
		lpart_int, _ := strconv.Atoi(lpart)
		rpart_int, _ := strconv.Atoi(rpart)
		lpart = fmt.Sprintf("%v", lpart_int)
		rpart = fmt.Sprintf("%v", rpart_int)
		return lpart, rpart
	} else {
		// rule 3
		value_int, _ := strconv.Atoi(value)
		value_int *= 2024
		value = fmt.Sprintf("%v", value_int)
		return value, ""
	}
}

func (puzzle *Puzzle) Part2() {
	puzzle.Load(REAL_DATA)
	// puzzle.p2(false, TEST_DATA_2, 6)
	puzzle.p2(false, TEST_DATA_2, 25)
	// puzzle.p2(TEST_DATA_2)
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}
