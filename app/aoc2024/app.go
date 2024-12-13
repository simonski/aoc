package aoc2024

import (
	"github.com/simonski/aoc/app/aoc2024/d1"
	"github.com/simonski/aoc/app/aoc2024/d10"
	"github.com/simonski/aoc/app/aoc2024/d11"
	"github.com/simonski/aoc/app/aoc2024/d12"
	"github.com/simonski/aoc/app/aoc2024/d2"
	"github.com/simonski/aoc/app/aoc2024/d3"
	"github.com/simonski/aoc/app/aoc2024/d4"
	"github.com/simonski/aoc/app/aoc2024/d5"
	"github.com/simonski/aoc/app/aoc2024/d6"
	"github.com/simonski/aoc/app/aoc2024/d7"
	"github.com/simonski/aoc/app/aoc2024/d8"
	"github.com/simonski/aoc/app/aoc2024/d9"
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

func (app Application) Summary(year int, day int) *utils.Summary {
	p := app.GetPuzzle(year, day)
	if p == nil {
		return utils.NewSummary(year, day)
	}
	s := p.GetSummary()
	return &s
}

func (app Application) GetPuzzle(year int, day int) utils.Puzzle {

	if year == 2024 {
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
		} else if day == 9 {
			return d9.NewPuzzle()
		} else if day == 10 {
			return d10.NewPuzzle()
		} else if day == 11 {
			return d11.NewPuzzle()
		} else if day == 12 {
			return d12.NewPuzzle()
			// } else if day == 13 {
			// 	return d13.NewPuzzle()
			// } else if day == 14 {
			// 	return d14.NewPuzzle()
			// } else if day == 15 {
			// 	return d15.NewPuzzle()
			// } else if day == 16 {
			// 	return d16.NewPuzzle()
			// } else if day == 17 {
			// 	return d17.NewPuzzle()
			// } else if day == 18 {
			// 	return d18.NewPuzzle()
			// } else if day == 19 {
			// 	return d19.NewPuzzle()
			// } else if day == 20 {
			// 	return d20.NewPuzzle()
			// } else if day == 21 {
			// 	return d21.NewPuzzle()
			// } else if day == 22 {
			// 	return d22.NewPuzzle()
			// } else if day == 23 {
			// 	return d23.NewPuzzle()
			// } else if day == 24 {
			// 	return d24.NewPuzzle()
			// } else if day == 25 {
			// 	return d25.NewPuzzle()
		}
	}

	return nil
}
