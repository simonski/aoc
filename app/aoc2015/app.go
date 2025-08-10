package aoc2015

import (
	"github.com/simonski/aoc/app/aoc2015/d1"
	"github.com/simonski/aoc/app/aoc2015/d2"
	"github.com/simonski/aoc/app/aoc2015/d3"
	"github.com/simonski/aoc/app/aoc2015/d4"
	"github.com/simonski/aoc/app/aoc2015/d5"
	"github.com/simonski/aoc/app/aoc2015/d6"
	"github.com/simonski/aoc/app/aoc2015/d7"
	"github.com/simonski/aoc/app/aoc2015/d8"
	"github.com/simonski/aoc/utils"
	cli "github.com/simonski/cli"
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

func (app Application) GetPuzzle(year int, day int) utils.Puzzle {

	if day == 1 {
		return d1.NewPuzzle()
	} else if day == 2 {
		return d2.NewPuzzle()
	} else if day == 3 {
		return d3.NewPuzzle()
	} else if day == 4 {
		return d4.NewPuzzle()
	} else if day == 5 {
		return d5.NewPuzzle()
	} else if day == 6 {
		return d6.NewPuzzle()
	} else if day == 7 {
		return d7.NewPuzzle()
	} else if day == 8 {
		return d8.NewPuzzle()
	}

	return nil

}
