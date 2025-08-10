package utils

import "github.com/simonski/cli"

type Puzzle interface {

	// run all parts for STDOUT
	Run()

	// when called from the server
	RunWWW() string

	// stdout calls
	Part1()
	Part2()

	GetSummary() Summary

	Render(cli *cli.CLI)
}
