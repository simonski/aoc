package aoc2021

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/gookit/color"
	utils "github.com/simonski/aoc/utils"
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
	return "I am 2021 - the year I tried out running a webserver on top of the solutions. Phew!"
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

func (app *Application) List() string {

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
	if day == 5 {
		grid := NewGrid(DAY_2021_05_DATA)
		msgb, _ := json.Marshal(grid)
		msg := string(msgb)
		return msg
	} else if day == 9 {
		grid := NewDay9Grid(DAY_2021_09_DATA)
		grid.LoadBasins()
		msgb, _ := json.Marshal(grid.GetApiResponse())
		msg := string(msgb)
		return msg
	} else if day == 11 {
		response := app.Y2021D11P2Api()
		msgb, _ := json.Marshal(response)
		msg := string(msgb)
		return msg
	}
	return ""
}

func (app Application) GetPuzzle(year int, day int) utils.Puzzle {

	if year == 2021 {
		// if day == 1 {
		// 	return d01.NewPuzzle()
		// } else if day == 2 {
		// 	return d02.NewPuzzle()
		// } else if day == 3 {
		// 	return d03.NewPuzzle()
		// } else if day == 4 {
		// 	return d04.NewPuzzle()
		// } else if day == 5 {
		// 	return d05.NewPuzzle()
		// } else if day == 6 {
		// 	return d06.NewPuzzle()
		// } else if day == 7 {
		// 	return d07.NewPuzzle()
		// } else if day == 8 {
		// 	return d08.NewPuzzle()
		// } else if day == 9 {
		// 	return d09.NewPuzzle()
		// } else if day == 10 {
		// 	return d10.NewPuzzle()
		// } else if day == 11 {
		// 	return d11.NewPuzzle()
		// } else if day == 12 {
		// 	return d12.NewPuzzle()
		// } else if day == 13 {
		// 	return d13.NewPuzzle()
		// } else if day == 14 {
		// 	return d14.NewPuzzle()
		// } else if day == 15 {
		// 	return d15.NewPuzzle()
		// } else if day == 16 {
		// 	return d16.NewPuzzle()
		// } else if day == 17 {
		// 	return d17.NewPuzzle()
		// } else if day == 18 {
		// 	return d18.NewPuzzle()
		// } else if day == 19 {
		// 	return d19.NewPuzzle()
		// } else if day == 20 {
		// 	return d20.NewPuzzle()
		// } else if day == 21 {
		// 	return d21.NewPuzzle()
		// } else if day == 22 {
		// 	return d22.NewPuzzle()
		// } else if day == 23 {
		// 	return d23.NewPuzzle()
		// } else if day == 24 {
		// 	return d24.NewPuzzle()
		// } else if day == 25 {
		// 	return d25.NewPuzzle()
		// }
	}

	return nil
}
