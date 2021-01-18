package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/gookit/color"

	goutils "github.com/simonski/goutils"
)

const USAGE_OVERALL = `aoc is my Advent Of Code set of attempts.

Usage:
    aoc <command> [arguments]
	
The commands are:

    run    (year) (day)      run a puzzle 
	
    render (year) (day)      render a puzzle to an animated gif

    list   (year)            list all the puzzles done so far
	
    version                  prints aoc version
	
Usage "aoc help <topic>" for more information.

`

type Application struct {
	CLI     goutils.CLI
	Verbose bool
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

func (app *Application) List(cli *goutils.CLI) {
	// USAGE := "Usage: aoc list (year) (day)"
	// fmt.Printf("Application.List(%v)", cli.Args)
	// year := cli.GetStringOrDefault("run", "")
	// if year == "" {
	// 	fmt.Printf("%v\n", USAGE)
	// 	os.Exit(1)
	// }

	max_year := 2020
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
	fmt.Println(bigline1)
	fmt.Println(header)
	fmt.Println(bigline2)
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
		fmt.Println(line)
	}
	fmt.Println(bigline3)

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
	fmt.Printf("Application.Help(%v)", cli.Args)
}

func NewApplication(cli goutils.CLI) Application {
	app := Application{CLI: cli}
	app.Verbose = cli.IndexOf("-v") > -1
	return app
}

// func Test(cli *goutils.CLI) {
// 	app := NewApplication(cli)
// 	reflectycall(app)
// }

func reflectycall(app Application) {
	// gimme an application
	// var app Application
	// var a Application
	rvalue := reflect.ValueOf(&app)
	fmt.Printf("rvalue %v\n", rvalue)
	mvalue := rvalue.MethodByName("Foo")
	fmt.Printf("mvalue %v\n", mvalue)
	cvalue := mvalue.Call([]reflect.Value{})
	fmt.Printf("cvalue %v\n", cvalue)
	if false {
		fmt.Printf("rvalue: %v, mvalue %v, cvalue %v\n", rvalue, mvalue, cvalue)
	}
	// reflect.ValueOf(&t).MethodByName("Foo").Call([]reflect.Value{})

}

func (app *Application) ReflectUponMySelf(name string, cli *goutils.CLI) {
	// gimme an application
	inputs := make([]reflect.Value, 1) // len(args))
	inputs[0] = reflect.ValueOf(cli)
	// for i, _ := range args {
	// 	inputs[i] = reflect.ValueOf(args[i])
	// }
	rvalue := reflect.ValueOf(&app)
	mvalue := rvalue.MethodByName(name)
	// inputs := cli
	cvalue := mvalue.Call(inputs) // []reflect.Value{})
	if false {
		fmt.Printf("rvalue: %v, mvalue %v, cvalue %v\n", rvalue, mvalue, cvalue)
	}
	// reflect.ValueOf(&t).MethodByName("Foo").Call([]reflect.Value{})

}

func (app *Application) AThing(cli *goutils.CLI) {
	fmt.Printf("AThing, cli length is %v\n", len(cli.Args))
}

func main() {
	cli := goutils.CLI{os.Args}
	if len(cli.Args) == 1 {
		Help(&cli)
		os.Exit(1)
	}
	app := NewApplication(cli)
	command := cli.Args[1]
	c := &cli
	if command == "run" {
		app.Run(c)
	} else if command == "list" {
		app.List(c)
	} else if command == "render" {
		app.Render(c)
	} else if command == "render" {
		app.Help(c)
	} else {
		Usage()
	}
}

func mainx() {
	cli := goutils.CLI{os.Args}
	if len(cli.Args) == 1 {
		Help(&cli)
		os.Exit(1)
	}

	command := cli.Args[1]
	if len(os.Args) < 2 {
		Usage()
	} else {
		// if command == "test" {
		// Test(&cli)
		if command == "help" {
			Help(&cli)
		} else if command == "render" {
			Render(&cli)
		} else if command == "info" {
			Info(&cli)
		} else if command == "2015" {
			AOC_2015(&cli)
		} else if command == "2015-01" {
			AOC_2015_01(&cli)
		} else if command == "2015-02" {
			AOC_2015_02(&cli)
		} else if command == "2015-03" {
			AOC_2015_03(&cli)
		} else if command == "2015-04" {
			AOC_2015_04(&cli)
		} else if command == "2015-05" {
			AOC_2015_05(&cli)
		} else if command == "2015-06" {
			AOC_2015_06(&cli)
		} else if command == "2015-07" {
			AOC_2015_07(&cli)

		} else if command == "2020" {
			AOC_2020(&cli)
		} else if command == "2020-01" {
			AOC_2020_01(&cli)
		} else if command == "2020-02" {
			AOC_2020_02(&cli)
		} else if command == "2020-03" {
			AOC_2020_03(&cli)
		} else if command == "2020-04" {
			AOC_2020_04(&cli)
		} else if command == "2020-05" {
			AOC_2020_05(&cli)
		} else if command == "2020-06" {
			AOC_2020_06(&cli)
		} else if command == "2020-07" {
			AOC_2020_07(&cli)
		} else if command == "2020-08" {
			AOC_2020_08(&cli)
		} else if command == "2020-09" {
			AOC_2020_09(&cli)
		} else if command == "2020-10" {
			AOC_2020_10(&cli)
		} else if command == "2020-11" {
			AOC_2020_11(&cli)
		} else if command == "2020-12" {
			AOC_2020_12(&cli)
		} else if command == "2020-13" {
			AOC_2020_13(&cli)
		} else if command == "2020-14" {
			AOC_2020_14(&cli)
		} else if command == "2020-15" {
			AOC_2020_15(&cli)
		} else if command == "2020-16" {
			AOC_2020_16(&cli)
		} else if command == "2020-17" {
			AOC_2020_17(&cli)
		} else if command == "2020-18" {
			AOC_2020_18(&cli)
		} else if command == "2020-19" {
			AOC_2020_19(&cli)
		} else if command == "2020-20" {
			AOC_2020_20(&cli)
		} else if command == "2020-21" {
			AOC_2020_21(&cli)
		} else if command == "2020-22" {
			AOC_2020_22(&cli)
		} else if command == "2020-23" {
			AOC_2020_23(&cli)
		} else if command == "2020-24" {
			AOC_2020_24(&cli)
		} else if command == "2020-25" {
			AOC_2020_25(&cli)
		} else if command == "version" {
			fmt.Printf("aoc version %v\n", VERSION)
		} else {
			// usage()
			fmt.Printf("aoc %s: unknown command.  Run `aoc help`.\n", command)
		}
	}

}

// Usage displays in terminal how to use the application
func Usage() {
	fmt.Printf(USAGE_OVERALL)
}

// Help shows detailed help on each command
// Usage: aoc help <command>
func Help(cli *goutils.CLI) {
	command := cli.GetStringOrDefault("help", "")
	if command == "" {
		Usage()
	} else {
		fmt.Printf("aoc help %s: unknown help topic. Run `aoc help`.\n", command)
		os.Exit(1)
	}
}

func Info(cli *goutils.CLI) {
	info := NewSysInfo()
	fmt.Printf("Platform %v CPU %v RAM %v\n", info.Platform, info.CPU, info.RAM)
}

func Render(cli *goutils.CLI) {

	USAGE := `aoc render <day> will call the render logic on that day

renderable days:
	2015-06
	2015-08

	2020-24
`
	fmt.Printf(USAGE)
	// if len(cli.Args) == 2 {
	// 	// they typed render only
	// } else {

	// }

	// }
	// command := cli.Args[2]
	// if len(os.Args) < 2 {
	// 	Usage()
	// }

}

func AOC_2015(cli *goutils.CLI) {
	AOC_2015_01(cli)
	AOC_2015_02(cli)
	AOC_2015_03(cli)
	AOC_2015_04(cli)
	AOC_2015_05(cli)
	AOC_2015_06(cli)
	AOC_2015_07(cli)
	// AOC_2015_08(cli)
	// AOC_2015_09(cli)
	// AOC_2015_10(cli)
	// AOC_2015_11(cli)
	// AOC_2015_13(cli)
	// AOC_2015_14(cli)
	// AOC_2015_15(cli)
	// AOC_2015_16(cli)
	// AOC_2015_17(cli)
	// AOC_2015_18(cli)
	// AOC_2015_19(cli)
	// AOC_2015_20(cli)
	// AOC_2015_21(cli)
	// AOC_2015_22(cli)
	// AOC_2015_23(cli)
	// AOC_2015_24(cli)
	// AOC_2015_25(cli)
}

// Usage displays in terminal how to use the application
func AOC_2020(cli *goutils.CLI) {
	AOC_2020_01(cli)
	AOC_2020_02(cli)
	AOC_2020_03(cli)
	AOC_2020_04(cli)
	AOC_2020_05(cli)
	AOC_2020_06(cli)
	AOC_2020_07(cli)
	AOC_2020_08(cli)
	AOC_2020_09(cli)
	AOC_2020_10(cli)
	AOC_2020_11(cli)
	AOC_2020_13(cli)
	AOC_2020_14(cli)
	AOC_2020_15(cli)
	AOC_2020_16(cli)
	AOC_2020_17(cli)
	AOC_2020_18(cli)
	AOC_2020_19(cli)
	AOC_2020_20(cli)
	AOC_2020_21(cli)
	AOC_2020_22(cli)
	AOC_2020_23(cli)
	AOC_2020_24(cli)
	AOC_2020_25(cli)
}
