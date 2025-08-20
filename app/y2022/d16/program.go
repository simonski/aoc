package d16

import (
	"fmt"
	"os"
	"strings"

	"github.com/simonski/aoc/utils"
	"github.com/simonski/cli"
)

/*
--- Day 05:  ---

*/

type Puzzle struct {
	title string
	year  int
	day   int
	input string
	lines []string
}

func NewPuzzleWithData(input string) *Puzzle {
	p := Puzzle{year: 2022, day: 16, title: "Proboscidea Volcanium"}
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

	if os.Args[5] == "test" {
		graph := NewGraph(TEST_DATA)
		aa := graph.Get("AA")
		path := NewPath()
		time := 30
		VERBOSE := true
		best_path := graph.dfs(aa, path, time, VERBOSE)
		fmt.Printf("\nBest=\n%v\n", best_path)
		fmt.Printf("Cache size %v, hits=%v, misses=%v\n", len(graph.Cache.data), graph.Cache.hits, graph.Cache.misses)
		fmt.Printf("best value in cache %v, path=%v\n", graph.Cache.max_value, graph.Cache.path)

	} else if os.Args[5] == "real" {
		graph := NewGraph(REAL_DATA)
		aa := graph.Get("AA")
		path := NewPath()
		time := 30
		VERBOSE := true
		best_path := graph.dfs(aa, path, time, VERBOSE)
		fmt.Printf("\nBest=\n%v\n", best_path)
		fmt.Printf("Cache size %v, hits=%v, misses=%v\n", len(graph.Cache.data), graph.Cache.hits, graph.Cache.misses)
		for key, value := range graph.Cache.data {
			if value == best_path {
				fmt.Println(key)
			}
		}
		fmt.Printf("best value in cache %v, path=%v\n", graph.Cache.max_value, graph.Cache.path)

	}
}

func (puzzle *Puzzle) Part2() {
	if os.Args[5] == "test" {
		graph := NewGraph(TEST_DATA)
		aa := graph.Get("AA")
		path := NewPath()
		time := 30
		VERBOSE := true
		best_path := graph.dfs2(aa, path, time, VERBOSE)
		fmt.Printf("\nBest=\n%v\n", best_path)
		fmt.Printf("Cache size %v, hits=%v, misses=%v\n", len(graph.Cache.data), graph.Cache.hits, graph.Cache.misses)
		fmt.Printf("best value in cache %v, path=%v\n", graph.Cache.max_value, graph.Cache.path)

	} else if os.Args[5] == "real" {
		graph := NewGraph(REAL_DATA)
		aa := graph.Get("AA")
		path := NewPath()
		time := 30
		VERBOSE := true
		best_path := graph.dfs2(aa, path, time, VERBOSE)
		fmt.Printf("\nBest=\n%v\n", best_path)
		fmt.Printf("Cache size %v, hits=%v, misses=%v\n", len(graph.Cache.data), graph.Cache.hits, graph.Cache.misses)
		for key, value := range graph.Cache.data {
			if value == best_path {
				fmt.Println(key)
			}
		}
		fmt.Printf("best value in cache %v, path=%v\n", graph.Cache.max_value, graph.Cache.path)

	}

}

func (puzzle *Puzzle) Run() {
	USAGE := "Usage: aoc run 2022 16 P1|P2 test|real (-v)"
	if len(os.Args) != 6 {
		fmt.Println(USAGE)
		os.Exit(1)
	} else if os.Args[4] == "P1" {
		puzzle.Part1()
	} else if os.Args[4] == "P2" {
		puzzle.Part2()
	} else {
		fmt.Println(USAGE)
		os.Exit(1)
	}
}

func (puzzle *Puzzle) GetSummary() utils.Summary {
	s := utils.Summary{Day: puzzle.day, Year: puzzle.year, Name: puzzle.title}
	s.ProgressP1 = utils.Started
	s.ProgressP2 = utils.Started
	s.DateStarted = "d01"
	return s
}

func (puzzle *Puzzle) RunWWW() string {
	return ""
}

func (puzzle *Puzzle) Render(cli *cli.CLI) {
}
