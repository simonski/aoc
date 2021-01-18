package main

import (
	"fmt"
	"os"
	"reflect"

	goutils "github.com/simonski/goutils"
)

const USAGE_OVERALL = `aoc is my Advent Of Code set of attempts.

Usage:
    aoc <command> [arguments]
	
The commands are:

    run (year) (day)         run a puzzle 
	
    render (year) (day)      render a puzzle to an animated gif

    list (year)              list all the puzzles done so far
	
    version                  prints aoc version
	
Usage "aoc help <topic>" for more information.

`

type AOC2015_06 struct {
}

func (logic *AOC2015_06) Part1(app *Application) {
}

type Application struct {
}

func (app *Application) Foo() {
	fmt.Printf("Application.Foo()")
}

func reflectycall() {
	// gimme an application
	var app Application
	rvalue := reflect.ValueOf(&app)
	mvalue := rvalue.MethodByName("foo")
	cvalue := mvalue.Call([]reflect.Value{})
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

	command := cli.Args[1]
	if len(os.Args) < 2 {
		Usage()
	} else {
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
