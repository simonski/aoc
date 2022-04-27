package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/simonski/aoc/api"
	"github.com/simonski/aoc/app"
	cli "github.com/simonski/cli"
	goutils "github.com/simonski/goutils"
)

func main() {
	c := &cli.CLI{Args: os.Args}
	app := app.NewAOCApplication(c)
	if len(c.Args) == 1 {
		app.Help(c)
		os.Exit(1)
	}
	command := c.Args[1]
	if command == "run" {
		app.Run(c)
	} else if command == "render" {
		app.Render(c)
	} else if command == "server" {
		server := api.NewServer(c, &app)
		server.Run()
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

func Info(cli *cli.CLI) {
	info := goutils.NewSysInfo()
	fmt.Printf("Platform %v CPU %v RAM %v\n", info.Platform, info.CPU, info.RAM)

	buildInfo, ok := debug.ReadBuildInfo()
	if ok {
		fmt.Printf("%v\n", buildInfo)
	}
}
