package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/gookit/color"
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

func main() {
	cli := goutils.CLI{Args: os.Args}
	c := &cli
	app := NewApplication(cli)
	if len(cli.Args) == 1 {
		app.Help(c)
		os.Exit(1)
	}
	command := cli.Args[1]
	if command == "run" {
		app.Run(c)
	} else if command == "render" {
		app.Render(c)
	} else if command == "info" {
		Info(c)
	} else if command == "version" {
		fmt.Printf("%v\n", VERSION)
	} else {
		fmt.Printf("I don't know how to '%v'.\n", command)
		os.Exit(1)
		// app.Help(c)
	}
}

type Application struct {
	CLI     goutils.CLI
	Verbose bool
}

func NewApplication(cli goutils.CLI) Application {
	app := Application{CLI: cli}
	app.Verbose = cli.IndexOf("-v") > -1
	return app
}

func (app *Application) Run(cli *goutils.CLI) {
	USAGE := "Usage: aoc run (year) (day)"
	// fmt.Printf("Application.Run(%v)", cli.Args)
	year := cli.GetStringOrDefault("run", "")
	if year == "" {
		fmt.Printf("%v\n", USAGE)
		os.Exit(1)
	}

	day := cli.GetStringOrDefault(year, "01")

	iyear, _ := strconv.Atoi(year)
	iday, _ := strconv.Atoi(day)

	methodNamePart1 := fmt.Sprintf("Y%vD%02dP1", iyear, iday)
	methodNamePart2 := fmt.Sprintf("Y%vD%02dP2", iyear, iday)

	_, m1, m1exists := app.GetMethod(methodNamePart1)
	_, m2, m2exists := app.GetMethod(methodNamePart2)

	if m1exists {
		// fmt.Printf("%v, %v %v Part 1 exists, calling.\n", methodNamePart1, year, day)
		m1.Call([]reflect.Value{})
	} else {
		fmt.Printf("%v(), %v %v Part 1 does not exist.\n", methodNamePart1, year, day)
	}

	if m2exists {
		// fmt.Printf("%v, %v %v Part 2 exists, calling.\n", methodNamePart2, year, day)
		m2.Call([]reflect.Value{})
	} else {
		fmt.Printf("%v(), %v %v Part 2 does not exist.\n", methodNamePart2, year, day)
	}

}

func (app *Application) List() string {

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
	bigline1 := "\u250F" + repeatstring(dash, 5) + repeatstring(biggydown, 24) + "\u2513"
	// bigline2 := "\u2523" + repeatstring(dash, repeat) + "\u252b"
	bigline2 := "\u2523" + repeatstring(dash, 5) + repeatstring(biggyup, 24) + "\u252b"
	bigline3 := "\u2517" + repeatstring(dash, 5) + repeatstring(biggyuponly, 24) + "\u251b"

	output += bigline1
	output += "\n"
	output += header
	output += "\n"
	output += bigline2
	output += "\n"
	for year := max_year; year >= min_year; year-- {
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

func (app *Application) Render(cli *goutils.CLI) {
	USAGE := "Usage: aoc render (year) (day)"
	// fmt.Printf("Application.Render(%v)", cli.Args)
	year := cli.GetStringOrDefault("run", "")
	if year == "" {
		fmt.Printf("%v\n", USAGE)
		os.Exit(1)
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

func Info(cli *goutils.CLI) {
	info := goutils.NewSysInfo()
	fmt.Printf("Platform %v CPU %v RAM %v\n", info.Platform, info.CPU, info.RAM)
}

func (app Application) GetMethod(methodName string) (reflect.Value, reflect.Value, bool) {

	// var a Application
	rvalue := reflect.ValueOf(&app)
	// fmt.Printf("GetMethod[%v], rvalue %v\n", methodName, rvalue)
	mvalue := rvalue.MethodByName(methodName)
	// fmt.Printf("GetMethod[%v], mvalue %v\n", methodName, mvalue)
	exists := false
	if reflect.Value.IsValid(mvalue) {
		exists = true
	}

	// if mvalue == nil {
	// exists = false
	// }
	// cvalue := mvalue.Call([]reflect.Value{})
	// fmt.Printf("cvalue %v\n", cvalue)
	// if false {
	// 	fmt.Printf("rvalue: %v, mvalue %v, cvalue %v\n", rvalue, mvalue, cvalue)
	// }
	return rvalue, mvalue, exists

}

func repeatstring(s string, times int) string {
	out := s
	for index := 0; index < times; index++ {
		out += s
	}
	return out
}
