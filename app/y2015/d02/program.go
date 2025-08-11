package d02

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
	"github.com/simonski/cli"
	"github.com/simonski/goutils"
)

/*
I Was Told There Would Be No Math
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
	iyear, _ := strconv.Atoi("2015")
	iday, _ := strconv.Atoi("02")
	p := Puzzle{year: iyear, day: iday, title: "I Was Told There Would Be No Math"}
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
	s.DateStarted = "d02"
	return s
}

func (puzzle *Puzzle) Part1() {
	puzzle.Load(REAL_DATA)
	lines := strings.Split(REAL_DATA, "\n")
	total := 0
	for _, line := range lines {
		p := NewPresent(line)
		total += p.Area()
	}
	fmt.Printf("Area: %v\n", total)

}

func (puzzle *Puzzle) Part2() {
	puzzle.Load(REAL_DATA)
	lines := strings.Split(REAL_DATA, "\n")
	volume := 0
	perimeter := 0
	for _, line := range lines {
		p := NewPresent(line)
		perimeter += p.Perimeter()
		volume += p.Volume()
	}
	total := perimeter + volume
	fmt.Printf("Volume %v Perimeter %v, total %v\n", volume, perimeter, total)

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

type Present struct {
	l int
	w int
	h int
}

func (p *Present) Area() int {
	l := p.l
	w := p.w
	h := p.h

	a1 := (2 * l * w)
	a2 := (2 * w * h)
	a3 := (2 * h * l)

	t := l * w
	t = goutils.Min(t, w*h)
	t = goutils.Min(t, h*l)

	return a1 + a2 + a3 + t
}

func (p *Present) Volume() int {
	return p.l * p.h * p.w
}

func (p *Present) Perimeter() int {
	arr := make([]int, 0)
	arr = append(arr, p.l)
	arr = append(arr, p.h)
	arr = append(arr, p.w)
	sort.Ints(arr)
	fmt.Printf("%v\n", arr)
	return arr[0] + arr[0] + arr[1] + arr[1]
}

func NewPresent(line string) *Present {
	splits := strings.Split(strings.TrimSpace(line), "x")
	l, _ := strconv.Atoi(splits[0])
	w, _ := strconv.Atoi(splits[1])
	h, _ := strconv.Atoi(splits[2])
	p := Present{l: l, w: w, h: h}
	return &p
}
