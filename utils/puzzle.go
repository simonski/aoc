package utils

import "github.com/simonski/cli"

type Puzzle interface {
	Run()
	GetSummary() Summary
	Api() string
	Render(cli *cli.CLI)
}
