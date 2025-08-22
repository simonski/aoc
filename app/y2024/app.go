package y2024

import (
	"github.com/simonski/aoc/app/y2024/d01"
	"github.com/simonski/aoc/app/y2024/d02"
	"github.com/simonski/aoc/app/y2024/d03"
	"github.com/simonski/aoc/app/y2024/d04"
	"github.com/simonski/aoc/app/y2024/d05"
	"github.com/simonski/aoc/app/y2024/d06"
	"github.com/simonski/aoc/app/y2024/d07"
	"github.com/simonski/aoc/app/y2024/d08"
	"github.com/simonski/aoc/app/y2024/d09"
	"github.com/simonski/aoc/app/y2024/d10"
	"github.com/simonski/aoc/app/y2024/d11"
	"github.com/simonski/aoc/app/y2024/d12"
	"github.com/simonski/aoc/app/y2024/d13"
	"github.com/simonski/aoc/app/y2024/d14"
	"github.com/simonski/aoc/app/y2024/d15"
	"github.com/simonski/aoc/app/y2024/d16"
	"github.com/simonski/aoc/app/y2024/d17"
	"github.com/simonski/aoc/app/y2024/d18"
	"github.com/simonski/aoc/app/y2024/d19"
	"github.com/simonski/aoc/app/y2024/d20"
	"github.com/simonski/aoc/app/y2024/d21"
	"github.com/simonski/aoc/app/y2024/d22"
	"github.com/simonski/aoc/app/y2024/d23"
	"github.com/simonski/aoc/app/y2024/d24"
	"github.com/simonski/aoc/app/y2024/d25"
	"github.com/simonski/aoc/utils"
	cli "github.com/simonski/cli"
)

type App struct {
	CLI     *cli.CLI
	Verbose bool
}

func NewApp(cli *cli.CLI) App {
	app := App{CLI: cli}
	app.Verbose = cli.IndexOf("-v") > -1
	return app
}

func (app App) Summary(year int, day int) *utils.Summary {
	p := app.GetPuzzle(year, day)
	if p == nil {
		return utils.NewSummary(year, day)
	}
	s := p.GetSummary()
	return &s
}

func (app App) GetPuzzle(year int, day int) utils.Puzzle {

	if day == 1 {
		return d01.NewPuzzle()
	} else if day == 2 {
		return d02.NewPuzzle()
	} else if day == 3 {
		return d03.NewPuzzle()
	} else if day == 4 {
		return d04.NewPuzzle()
	} else if day == 5 {
		return d05.NewPuzzle()
	} else if day == 6 {
		return d06.NewPuzzle()
	} else if day == 7 {
		return d07.NewPuzzle()
	} else if day == 8 {
		return d08.NewPuzzle()
	} else if day == 9 {
		return d09.NewPuzzle()
	} else if day == 10 {
		return d10.NewPuzzle()
	} else if day == 11 {
		return d11.NewPuzzle()
	} else if day == 12 {
		return d12.NewPuzzle()
	} else if day == 13 {
		return d13.NewPuzzle()
	} else if day == 14 {
		return d14.NewPuzzle()
	} else if day == 15 {
		return d15.NewPuzzle()
	} else if day == 16 {
		return d16.NewPuzzle()
	} else if day == 17 {
		return d17.NewPuzzle()
	} else if day == 18 {
		return d18.NewPuzzle()
	} else if day == 19 {
		return d19.NewPuzzle()
	} else if day == 20 {
		return d20.NewPuzzle()
	} else if day == 21 {
		return d21.NewPuzzle()
	} else if day == 22 {
		return d22.NewPuzzle()
	} else if day == 23 {
		return d23.NewPuzzle()
	} else if day == 24 {
		return d24.NewPuzzle()
	} else if day == 25 {
		return d25.NewPuzzle()
	}

	return nil
}
