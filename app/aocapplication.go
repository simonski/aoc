package app

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gookit/color"
	aoc2015 "github.com/simonski/aoc/app/aoc2015"
	aoc2016 "github.com/simonski/aoc/app/aoc2016"
	aoc2017 "github.com/simonski/aoc/app/aoc2017"
	aoc2018 "github.com/simonski/aoc/app/aoc2018"
	aoc2019 "github.com/simonski/aoc/app/aoc2019"
	aoc2020 "github.com/simonski/aoc/app/aoc2020"
	aoc2021 "github.com/simonski/aoc/app/aoc2021"
	"github.com/simonski/aoc/utils"
	cli "github.com/simonski/cli"
	goutils "github.com/simonski/goutils"
)

const USAGE_OVERALL = `aoc is my Advent Of Code set of attempts.

Usage:
    aoc <command> [arguments]
	
Status:

LIST
Commands:

The commands are:

    run    (year) (day)      run a puzzle 
	
    render (year) (day)      render a puzzle to an animated gif

    info                     prints computer information

    version                  prints aoc version
	
`

type AOCApplication struct {
	CLI     *cli.CLI
	Verbose bool
}

func NewAOCApplication(cli *cli.CLI) AOCApplication {
	app := AOCApplication{CLI: cli}
	app.Verbose = cli.IndexOf("-v") > -1
	return app
}

func (app *AOCApplication) Run(cli *cli.CLI) {
	USAGE := "Usage: aoc run (year) (day)"
	year := cli.GetStringOrDefault("run", "")
	if year == "" {
		fmt.Printf("%v\n", USAGE)
		os.Exit(1)
	}
	iyear, _ := strconv.Atoi(year)
	al := app.GetAppLogic(iyear)
	if al == nil {
		fmt.Printf("Sorry, we don't have year %v\n", year)
		return
	}
	al.Run(cli)

}

func (app *AOCApplication) Render(cli *cli.CLI) {
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

func (app *AOCApplication) Help(cli *cli.CLI) {
	output := strings.ReplaceAll(USAGE_OVERALL, "LIST", app.List())
	fmt.Println(output)
	command := cli.GetStringOrDefault("help", "")
	if command != "" {
		fmt.Printf("aoc help %s: unknown help topic. Run `aoc help`.\n", command)
		os.Exit(1)
	}
}

func (a *AOCApplication) List() string {

	output := ""

	max_year := 2021
	min_year := 2015

	tick := "\u2713"
	star := "\u2606"

	tick = star
	cross := " " // \u2717"

	dash := "\u2501"

	header := fmt.Sprintf("\u2503      \u2503")
	for day := 1; day <= 25; day++ {
		header = fmt.Sprintf("%v%02d\u2503", header, day)
	}

	//https://en.wikipedia.org/wiki/Box-drawing_character
	biggydown := "\u2533" + dash + dash
	biggyup := "\u2543" + dash + dash
	biggyuponly := "\u253b" + dash + dash

	// repeat := 80
	// bigline1 := "\u250f" + repeatstring(dash, repeat) + "\u2513"
	bigline1 := "\u250F" + goutils.Repeatstring(dash, 5) + goutils.Repeatstring(biggydown, 24) + "\u2513"
	// bigline2 := "\u2523" + repeatstring(dash, repeat) + "\u252b"
	bigline2 := "\u2523" + goutils.Repeatstring(dash, 5) + goutils.Repeatstring(biggyup, 24) + "\u252b"
	bigline3 := "\u2517" + goutils.Repeatstring(dash, 5) + goutils.Repeatstring(biggyuponly, 24) + "\u251b"

	output += bigline1
	output += "\n"
	output += header
	output += "\n"
	output += bigline2
	output += "\n"
	for year := max_year; year >= min_year; year-- {
		app := a.GetAppLogic(year)
		line := fmt.Sprintf("\u2503 %04d \u2503", year)
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
				// if m1existsRender {
				// 	part1 = tick
				// }
				// } else {
				// 	part1 = tick + tick
			} else {
				part1 = cross
			}

			if m2exists {
				part2 = tick
			} else {
				part2 = cross
			}
			// 	if m2existsRender {
			// 		part2 = tick
			// 	}
			// } else {
			// 	part2 = tick + tick

			// }

			if m1existsRender {
				part1 = gold(part1)
			}
			if m2existsRender {
				part2 = gold(part2)
			}

			line = fmt.Sprintf("%v%v%v\u2503", line, part1, part2)

		}
		output += line
		output += "\n"
	}
	output += bigline3
	output += "\n"

	return output

}

func (a *AOCApplication) GetAppLogic(year int) utils.AppLogic {
	if year == 2015 {
		return aoc2015.NewApplication(a.CLI)
	} else if year == 2016 {
		return aoc2016.NewApplication(a.CLI)
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
	} else {
		return nil
	}
}
