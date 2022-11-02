package aoc2018

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/gookit/color"
	"github.com/simonski/aoc/utils"
	cli "github.com/simonski/cli"
	goutils "github.com/simonski/goutils"
)

type Application struct {
	CLI     *cli.CLI
	Verbose bool
}

func NewApplication(cli *cli.CLI) Application {
	app := Application{CLI: cli}
	app.Verbose = cli.IndexOf("-v") > -1
	return app
}

func (app Application) Summary(year int, day int) *utils.Summary {
	methodName := fmt.Sprintf("Y%vD%02d_Summary", year, day)
	_, rvar, exists := app.GetMethod(methodName)
	if exists {
		results := rvar.Call([]reflect.Value{})
		v := results[0]
		concrete, _ := v.Interface().(*utils.Summary)
		return concrete
	} else {
		return utils.NewSummary(year, day)
	}
}

func (app Application) GetName() string {
	return "I am 2018"
}

func (app Application) Run(cli *cli.CLI) {
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

func (app Application) List() string {

	output := ""

	max_year := 2021
	min_year := 2015

	tick := "\u2713"
	star := "\u2606"

	tick = star
	cross := " " // \u2717"

	dash := "\u2501"

	header := "\u2503      \u2503"
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

func (app Application) Render(cli *cli.CLI) {
	USAGE := "Usage: aoc render (year) (day)"
	// fmt.Printf("Application.Render(%v)", cli.Args)
	year := cli.GetStringOrDefault("run", "")
	if year == "" {
		fmt.Printf("%v\n", USAGE)
		os.Exit(1)
	}
}

func (app Application) Help(cli *cli.CLI) {
}

func (app Application) GetMethod(methodName string) (reflect.Value, reflect.Value, bool) {
	rvalue := reflect.ValueOf(&app)
	mvalue := rvalue.MethodByName(methodName)
	exists := false
	if reflect.Value.IsValid(mvalue) {
		exists = true
	}
	return rvalue, mvalue, exists
}

func (app Application) Api(day int) string {
	return ""
}
