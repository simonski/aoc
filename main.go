package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/simonski/aoc/api"
	"github.com/simonski/aoc/app"
	"github.com/simonski/aoc/app/constants"
	cli "github.com/simonski/cli"
	goutils "github.com/simonski/goutils"
)

func main() {
	c := &cli.CLI{Args: os.Args}
	c.Shift()
	command := c.GetCommand()
	app := app.NewAOC(c)

	if command == "" {
		app.Help(c)
		os.Exit(1)
	} else if command == "run" {
		app.Run(c)
	} else if command == "render" {
		app.Render(c)
	} else if command == "server" {
		server := api.NewServer(c, &app)
		server.Run()
	} else if command == "summary" {
		app.Summary(c)
	} else if command == "info" {
		Info(c)
	} else if command == "version" {
		fmt.Printf("%v\n", constants.VERSION)
	} else {
		fmt.Printf("I don't know how to '%v'.\n", command)
		os.Exit(1)
	}
}

func Info(cli *cli.CLI) {
	info := goutils.NewSysInfo()
	fmt.Printf("Platform %v CPU %v RAM %v\n", info.Platform, info.CPU, info.RAM)
	buildInfo, ok := debug.ReadBuildInfo()
	if ok {
		fmt.Printf("%v\n", buildInfo)
	}
}
