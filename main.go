package main

import (
	"fmt"
	"os"

	goutils "github.com/simonski/goutils"
)

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
	Console("aoc is my Advent Of Code set of attempts.")
	Console("")
	Console("Usage:")
	Console("\taoc <command> [arguments]")
	Console("")
	Console("The commands are:")
	Console("")
	Console("\t2020", "run all 2020 examples")
	Console("\t  2020-01", "run only 2020-01")
	Console("\t  2020-02", "run the 2020-02")
	Console("\t  ..")
	Console("")
	Console("\t  2020-07", "run the 2020-07")
	Console("")
	Console("\tversion", "prints aoc version ")
	Console("")
	Console("Usage \"aoc help <topic>\" for more information.")
	Console("")
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

// Usage displays in terminal how to use the application
func AOC_2020(cli *goutils.CLI) {
	AOC_2020_01(cli)
	AOC_2020_02(cli)
	AOC_2020_03(cli)
	AOC_2020_04(cli)
	AOC_2020_05(cli)
}
