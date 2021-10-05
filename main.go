package main

import (
	"fmt"
	"github.com/simonski/aoc/app"
	goutils "github.com/simonski/goutils"
	"os"
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
	app := app.NewApplication(cli)
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

func Info(cli *goutils.CLI) {
	info := goutils.NewSysInfo()
	fmt.Printf("Platform %v CPU %v RAM %v\n", info.Platform, info.CPU, info.RAM)
}

// func repeatstring(s string, times int) string {
// 	out := s
// 	for index := 0; index < times; index++ {
// 		out += s
// 	}
// 	return out
// }
