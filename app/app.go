package app

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/gookit/color"
	aoc2015 "github.com/simonski/aoc/app/aoc2015"
	aoc2016 "github.com/simonski/aoc/app/aoc2016"
	aoc2017 "github.com/simonski/aoc/app/aoc2017"
	aoc2018 "github.com/simonski/aoc/app/aoc2018"
	aoc2019 "github.com/simonski/aoc/app/aoc2019"
	aoc2020 "github.com/simonski/aoc/app/aoc2020"
	aoc2021 "github.com/simonski/aoc/app/aoc2021"
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

type Application struct {
	CLI     *goutils.CLI
	Verbose bool
}

type AppLogic interface {
	Run(cli *goutils.CLI)
	Render(cli *goutils.CLI)
	Help(cli *goutils.CLI)
	GetMethod(methodName string) (reflect.Value, reflect.Value, bool)
}

func NewApplication(cli *goutils.CLI) Application {
	app := Application{CLI: cli}
	app.Verbose = cli.IndexOf("-v") > -1
	return app
}

func (app *Application) Run(cli *goutils.CLI) {
	USAGE := "Usage: aoc run (year) (day)"
	year := cli.GetStringOrDefault("run", "")
	if year == "" {
		fmt.Printf("%v\n", USAGE)
		os.Exit(1)
	}
	if year == "2015" {
		a := aoc2015.NewApplication(cli)
		a.Run(cli)
	} else if year == "2016" {
		a := aoc2016.NewApplication(cli)
		a.Run(cli)
	} else if year == "2017" {
		a := aoc2017.NewApplication(cli)
		a.Run(cli)
	} else if year == "2018" {
		a := aoc2018.NewApplication(cli)
		a.Run(cli)
	} else if year == "2019" {
		a := aoc2019.NewApplication(cli)
		a.Run(cli)
	} else if year == "2020" {
		a := aoc2020.NewApplication(cli)
		a.Run(cli)
	} else if year == "2021" {
		a := aoc2021.NewApplication(cli)
		a.Run(cli)
	}

}

func (app *Application) Render(cli *goutils.CLI) {
	USAGE := "Usage: aoc render (year) (day)"
	year := cli.GetStringOrDefault("run", "")
	if year == "" {
		fmt.Printf("%v\n", USAGE)
		os.Exit(1)
	}
	if year == "2015" {
		a := aoc2015.NewApplication(cli)
		a.Render(cli)
	} else if year == "2016" {
		a := aoc2016.NewApplication(cli)
		a.Render(cli)
	} else if year == "2017" {
		a := aoc2017.NewApplication(cli)
		a.Render(cli)
	} else if year == "2018" {
		a := aoc2018.NewApplication(cli)
		a.Render(cli)
	} else if year == "2019" {
		a := aoc2019.NewApplication(cli)
		a.Render(cli)
	} else if year == "2020" {
		a := aoc2020.NewApplication(cli)
		a.Render(cli)
	} else if year == "2021" {
		a := aoc2021.NewApplication(cli)
		a.Render(cli)
	}
}

func (app *Application) Help(cli *goutils.CLI) {
	output := strings.ReplaceAll(USAGE_OVERALL, "LIST", app.List())
	fmt.Println(output)
	command := cli.GetStringOrDefault("help", "")
	if command != "" {
		fmt.Printf("aoc help %s: unknown help topic. Run `aoc help`.\n", command)
		os.Exit(1)
	}
}

func (a *Application) List() string {

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

			_, _, m1exists := GetMethod(app, methodNamePart1)
			_, _, m2exists := GetMethod(app, methodNamePart2)
			_, _, m1existsRender := GetMethod(app, methodNamePart1Render)
			_, _, m2existsRender := GetMethod(app, methodNamePart2Render)

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

func (a *Application) GetAppLogic(year int) AppLogic {
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

func GetMethod(app AppLogic, methodName string) (reflect.Value, reflect.Value, bool) {
	rvalue := reflect.ValueOf(app)
	mvalue := rvalue.MethodByName(methodName)
	exists := false
	if reflect.Value.IsValid(mvalue) {
		exists = true
	}
	return rvalue, mvalue, exists
}
