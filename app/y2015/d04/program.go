package d04

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
	"github.com/simonski/cli"
)

/*
The Ideal Stocking Stuffer
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
	iday, _ := strconv.Atoi("04")
	p := Puzzle{year: iyear, day: iday, title: "The Ideal Stocking Stuffer"}
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
	s.DateStarted = "d04"
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

func (puzzle *Puzzle) Part2() {
	puzzle.Load(REAL_DATA)
}

func (puzzle *Puzzle) Part1() {
	prefix := "bgvyzdsv"
	// prefix = "abcdef"
	counter := 0

	// word := fmt.Sprintf("%v%v", prefix, 609043)
	// bytes := []byte(word)
	// hash := fmt.Sprintf("%x", md5.Sum(bytes))
	// fmt.Printf("%v %v : %v\n", prefix, 609043, hash)
	// fmt.Printf("\n\n")

	for {
		word := fmt.Sprintf("%v%v", prefix, counter)
		bytes := []byte(word)
		hash := fmt.Sprintf("%x", md5.Sum(bytes))
		if strings.Index(hash, "000000") == 0 {
			fmt.Printf("%v %v\n", hash, counter)
			break
		}
		counter++
	}
}
