package app

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gookit/color"
	"github.com/simonski/aoc/app/aoc2015"
	"github.com/simonski/aoc/app/aoc2016"
	"github.com/simonski/aoc/app/aoc2017"
	"github.com/simonski/aoc/app/aoc2018"
	"github.com/simonski/aoc/app/aoc2019"
	"github.com/simonski/aoc/app/aoc2020"
	"github.com/simonski/aoc/app/aoc2021"
	"github.com/simonski/aoc/app/aoc2022"
	"github.com/simonski/aoc/app/aoc2023"
	"github.com/simonski/aoc/app/aoc2024"
	"github.com/simonski/aoc/app/constants"
	"github.com/simonski/aoc/utils"
	"github.com/simonski/cli"
	"github.com/simonski/goutils"
)

type AOC struct {
	CLI     *cli.CLI
	Verbose bool
}

func NewAOC(cli *cli.CLI) AOC {
	app := AOC{CLI: cli}
	app.Verbose = cli.IndexOf("-v") > -1
	return app
}

func (aoc *AOC) Run(cli *cli.CLI) {
	USAGE := "Usage: aoc run (year) (day)"
	year := cli.GetStringOrDefault("run", "")
	if year == "" {
		fmt.Printf("%v\n", USAGE)
		os.Exit(1)
	}
	iyear, _ := strconv.Atoi(year)
	day := cli.GetStringOrDie(year)
	iday, _ := strconv.Atoi(day)
	puzzle := aoc.GetPuzzleOrDie(iyear, iday)
	puzzle.Run()
}

func (aoc *AOC) Summary(cli *cli.CLI) {
	USAGE := "Usage: aoc summary (year) (day)"
	year := cli.GetStringOrDefault("summary", "")
	if year == "" {
		fmt.Printf("%v\n", USAGE)
		os.Exit(1)
	}
	iyear, _ := strconv.Atoi(year)
	day := cli.GetStringOrDefault(year, "")
	iday, _ := strconv.Atoi(day)

	puzzle := aoc.GetPuzzleOrDie(iyear, iday)
	summary := puzzle.GetSummary()
	j, _ := json.MarshalIndent(summary, "", "  ")
	fmt.Printf("%v\n", string(j))
}

func (aoc *AOC) Render(cli *cli.CLI) {
	USAGE := "Usage: aoc render (year) (day)"
	year := cli.GetStringOrDefault("render", "")
	if year == "" {
		fmt.Printf("%v\n", USAGE)
		os.Exit(1)
	}
	iyear, _ := strconv.Atoi(year)
	day := cli.GetStringOrDefault(year, "")
	iday, _ := strconv.Atoi(day)
	puzzle := aoc.GetPuzzleOrDie(iyear, iday)
	puzzle.Render(cli)
}

func (aoc *AOC) Help(cli *cli.CLI) {
	output := strings.ReplaceAll(constants.USAGE_OVERALL, "LIST", aoc.List())
	fmt.Println(output)
	command := cli.GetStringOrDefault("help", "")
	if command != "" {
		fmt.Printf("aoc help %s: unknown help topic. Run `aoc help`.\n", command)
		os.Exit(1)
	}
}

// builds a table of progress for all problems for the terminal
func (aoc *AOC) List() string {

	output := ""

	max_year := constants.MAX_YEAR
	min_year := constants.MIN_YEAR

	tick := "\u2713"
	star := "\u2606"

	tick = star
	cross := " " // \u2717"

	dash := "\u2501"

	header := "\u2503           \u2503"
	for day := 1; day <= 25; day++ {
		header = fmt.Sprintf("%v%02d\u2503", header, day)
	}

	//https://en.wikipedia.org/wiki/Box-drawing_character
	biggydown := "\u2533" + dash + dash
	biggyup := "\u2543" + dash + dash
	biggyuponly := "\u253b" + dash + dash

	// repeat := 80
	// bigline1 := "\u250f" + repeatstring(dash, repeat) + "\u2513"
	bigline1 := "\u250F" + goutils.Repeatstring(dash, 10) + goutils.Repeatstring(biggydown, 24) + "\u2513"
	// bigline2 := "\u2523" + repeatstring(dash, repeat) + "\u252b"
	bigline2 := "\u2523" + goutils.Repeatstring(dash, 10) + goutils.Repeatstring(biggyup, 24) + "\u252b"
	bigline3 := "\u2517" + goutils.Repeatstring(dash, 10) + goutils.Repeatstring(biggyuponly, 24) + "\u251b"

	output += bigline1
	output += "\n"
	output += header
	output += "\n"
	output += bigline2
	output += "\n"
	totalStars := 0
	totalAvailableStars := 0
	for year := max_year; year >= min_year; year-- {
		line := ""
		stars := 0
		totalAvailableStars += 25
		for day := 1; day <= 25; day++ {

			summary := aoc.GetSummary(year, day)
			m1exists := summary.ProgressP1 == utils.Completed
			m2exists := summary.ProgressP2 == utils.Completed

			part1 := cross
			part2 := cross

			gold := color.FgYellow.Render

			if m1exists {
				part1 = tick
				stars++
			} else {
				part1 = cross
			}

			if m2exists {
				part2 = tick
				stars++
			} else {
				part2 = cross
			}

			if summary.VisualisationP1 {
				part1 = gold(part1)
			}
			if summary.VisualisationP1 {
				part2 = gold(part2)
			}

			line = fmt.Sprintf("%v%v%v\u2503", line, part1, part2)

		}
		totalStars += stars
		line = fmt.Sprintf("\u2503 %04d (%02v) \u2503", year, stars) + line

		output += line
		output += "\n"
	}
	output += bigline3
	output += "\n"
	output += fmt.Sprintf("(Total Stars %v/%v)\n", totalStars, totalAvailableStars)

	return output
}

func (a *AOC) GetPuzzleOrDie(year int, day int) utils.Puzzle {
	p := a.GetPuzzle(year, day)
	if p == nil {
		fmt.Printf("Sorry, we don't have year-day %v-%v\n", year, day)
		os.Exit(1)
		return nil
	}
	return p
}

func (a *AOC) GetPuzzle(year int, day int) utils.Puzzle {
	if year == 2015 {
		a := aoc2015.NewApplication(a.CLI)
		return a.GetPuzzle(year, day)
	} else if year == 2016 {
		a := aoc2016.NewApplication(a.CLI)
		return a.GetPuzzle(year, day)
	} else if year == 2017 {
		a := aoc2017.NewApplication(a.CLI)
		return a.GetPuzzle(year, day)
	} else if year == 2018 {
		a := aoc2018.NewApplication(a.CLI)
		return a.GetPuzzle(year, day)
	} else if year == 2019 {
		a := aoc2019.NewApplication(a.CLI)
		return a.GetPuzzle(year, day)
	} else if year == 2020 {
		a := aoc2020.NewApplication(a.CLI)
		return a.GetPuzzle(year, day)
	} else if year == 2021 {
		a := aoc2021.NewApplication(a.CLI)
		return a.GetPuzzle(year, day)
	} else if year == 2022 {
		a := aoc2022.NewApplication(a.CLI)
		return a.GetPuzzle(year, day)
	} else if year == 2023 {
		a := aoc2023.NewApplication(a.CLI)
		return a.GetPuzzle(year, day)
	} else if year == 2024 {
		a := aoc2024.NewApplication(a.CLI)
		return a.GetPuzzle(year, day)
	} else {
		return nil
	}
}

func (a *AOC) GetSummary(year int, day int) *utils.Summary {
	puzzle := a.GetPuzzleOrDie(year, day)
	s := puzzle.GetSummary()
	return &s
}

func (a *AOC) GetSummaries() []*utils.Summary {
	results := make([]*utils.Summary, 0)
	for year := constants.MIN_YEAR; year <= constants.MAX_YEAR; year++ {
		for day := 1; day <= 25; day++ {
			summary := a.GetSummary(year, day)
			results = append(results, summary)
		}
	}
	return results
}
