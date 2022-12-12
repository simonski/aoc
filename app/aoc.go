package app

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gookit/color"
	aoc2015 "github.com/simonski/aoc/app/aoc2015"
	"github.com/simonski/aoc/app/aoc2016"
	"github.com/simonski/aoc/app/aoc2017"
	"github.com/simonski/aoc/app/aoc2018"
	"github.com/simonski/aoc/app/aoc2019"
	"github.com/simonski/aoc/app/aoc2020"
	"github.com/simonski/aoc/app/aoc2021"
	"github.com/simonski/aoc/app/aoc2022"
	aoc2022d06 "github.com/simonski/aoc/app/aoc2022/d06"
	aoc2022d11 "github.com/simonski/aoc/app/aoc2022/d11"
	aoc2022d12 "github.com/simonski/aoc/app/aoc2022/d12"
	"github.com/simonski/aoc/app/constants"
	"github.com/simonski/aoc/utils"
	cli "github.com/simonski/cli"
	goutils "github.com/simonski/goutils"
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

func (app *AOC) Run(cli *cli.CLI) {
	USAGE := "Usage: aoc run (year) (day)"
	year := cli.GetStringOrDefault("run", "")
	if year == "" {
		fmt.Printf("%v\n", USAGE)
		os.Exit(1)
	}
	iyear, _ := strconv.Atoi(year)
	if iyear < 2022 {
		al := app.GetAppLogic(iyear)
		if al == nil {
			fmt.Printf("Sorry, we don't have year %v\n", year)
			return
		}
		al.Run(cli)
	} else {
		// iyear, _ := strconv.Atoi(year)
		day := cli.GetStringOrDie(year)
		// iday, _ := strconv.Atoi(day)
		puzzle := app.GetPuzzle(year, day)
		if puzzle == nil {
			fmt.Printf("Sorry, we don't have year %v\n", year)
			return
		} else {
			puzzle.Run()
		}

	}
}

func (app *AOC) Summary(cli *cli.CLI) {
	USAGE := "Usage: aoc summary (year) (day)"
	year := cli.GetStringOrDefault("summary", "")
	if year == "" {
		fmt.Printf("%v\n", USAGE)
		os.Exit(1)
	}
	iyear, _ := strconv.Atoi(year)

	day := cli.GetStringOrDefault(year, "")
	iday, _ := strconv.Atoi(day)

	al := app.GetAppLogic(iyear)
	if al == nil {
		fmt.Printf("Sorry, we don't have year %v\n", year)
		return
	}
	summary := al.Summary(iyear, iday)
	if summary != nil {
		j, _ := json.MarshalIndent(summary, "", "  ")
		fmt.Printf("%v\n", string(j))
	} else {
		fmt.Println("No summary.")
	}

}

func (app *AOC) Render(cli *cli.CLI) {
	USAGE := "Usage: aoc render (year) (day)"
	year := cli.GetStringOrDefault("render", "")
	if year == "" {
		fmt.Printf("%v\n", USAGE)
		os.Exit(1)
	}
	iyear, _ := strconv.Atoi(year)
	al := app.GetAppLogic(iyear)
	al.Render(cli)
}

func (app *AOC) Help(cli *cli.CLI) {
	output := strings.ReplaceAll(constants.USAGE_OVERALL, "LIST", app.List())
	fmt.Println(output)
	command := cli.GetStringOrDefault("help", "")
	if command != "" {
		fmt.Printf("aoc help %s: unknown help topic. Run `aoc help`.\n", command)
		os.Exit(1)
	}
}

// builds a table of progress for all problems for the terminal
func (a *AOC) List() string {

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
		app := a.GetAppLogic(year)
		if app == nil {
			fmt.Printf("AppLogic for year %v does not exist.\n", year)
			continue
		}
		line := ""
		stars := 0
		totalAvailableStars += 25
		for day := 1; day <= 25; day++ {
			methodNamePart1 := fmt.Sprintf("Y%vD%02dP1", year, day)
			methodNamePart2 := fmt.Sprintf("Y%vD%02dP2", year, day)
			methodNamePart1Render := fmt.Sprintf("Y%vD%02dP1Render", year, day)
			methodNamePart2Render := fmt.Sprintf("Y%vD%02dP2Render", year, day)

			_, _, m1exists := app.GetMethod(methodNamePart1)
			_, _, m2exists := app.GetMethod(methodNamePart2)
			_, _, m1existsRender := app.GetMethod(methodNamePart1Render)
			_, _, m2existsRender := app.GetMethod(methodNamePart2Render)

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

			if m1existsRender {
				part1 = gold(part1)
			}
			if m2existsRender {
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

func (a *AOC) GetPuzzle(year string, day string) utils.Puzzle {
	if year == "2022" {
		if day == "06" {
			return aoc2022d06.NewPuzzle()
		} else if day == "11" {
			return aoc2022d11.NewPuzzle()
		} else if day == "12" {
			return aoc2022d12.NewPuzzle()
			// } else if day == "05" {
			// 	return aoc2022d05.NewPuzzle()
			// } else if day == "04" {
			// 	return aoc2022d04.NewPuzzle()
			// } else if day == "03" {
			// 	return aoc2022d03.NewPuzzle()
			// } else if day == "02" {
			// 	return aoc2022d02.NewPuzzle()
			// } else if day == "01" {
			// 	return aoc2022d01.NewPuzzle()
		} else {
			return nil
		}

	} else {
		return nil
	}
	// concreteTypeName := fmt.Sprintf("app.aoc%v.d%v.Puzzle", year, day)
	// puzzle := reflect.TypeOf(concreteTypeName)
	// return &puzzle
}

func (a *AOC) GetAppLogic(year int) utils.AppLogic {
	if year == 2015 {
		return aoc2015.NewApplication(a.CLI)
	} else if year == 2016 {
		return aoc2016.NewApplication(a.CLI)
	} else if year == 2017 {
		return aoc2017.NewApplication(a.CLI)
	} else if year == 2018 {
		return aoc2018.NewApplication(a.CLI)
	} else if year == 2019 {
		return aoc2019.NewApplication(a.CLI)
	} else if year == 2020 {
		return aoc2020.NewApplication(a.CLI)
	} else if year == 2021 {
		return aoc2021.NewApplication(a.CLI)
	} else if year == 2022 {
		return aoc2022.NewApplication(a.CLI)
	} else {
		return nil
	}
}

func (a *AOC) GetSummary(year int, day int) *utils.Summary {
	appLogic := a.GetAppLogic(year)
	if appLogic != nil {
		return appLogic.Summary(year, day)
	} else {
		return nil
	}
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
