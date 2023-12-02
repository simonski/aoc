package aoc2023

import (
	"github.com/simonski/aoc/app/aoc2023/d01"
	"github.com/simonski/aoc/app/aoc2023/d02"
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

// func (app Application) Run(cli *cli.CLI) {
// }

// func (app Application) Render(cli *cli.CLI) {
// }

// func (app Application) Api(day int) string {
// 	return ""
// }

func (app Application) Summary(year int, day int) *utils.Summary {
	p := app.GetPuzzle(year, day)
	if p == nil {
		return utils.NewSummary(year, day)
	}
	s := p.GetSummary()
	return &s
}

// func (app Application) GetMethod(methodName string) (reflect.Value, reflect.Value, bool) {
// 	return reflect.ValueOf(""), reflect.ValueOf(""), false
// }

func (app Application) GetPuzzle(year int, day int) utils.Puzzle {

	if year == 2023 {
		if day == 1 {
			return d01.NewPuzzle()
		} else if day == 2 {
			return d02.NewPuzzle()
			// } else if day == 3 {
			// 	return d03.NewPuzzle()
			// } else if day == 4 {
			// 	return d04.NewPuzzle()
			// } else if day == 5 {
			// 	return d05.NewPuzzle()
			// } else if day == 6 {
			// 	return d06.NewPuzzle()
			// } else if day == 7 {
			// 	return d07.NewPuzzle()
			// } else if day == 8 {
			// 	return d08.NewPuzzle()
			// } else if day == 9 {
			// 	return d09.NewPuzzle()
			// } else if day == 10 {
			// 	return d10.NewPuzzle()
			// } else if day == 11 {
			// 	return d11.NewPuzzle()
			// } else if day == 12 {
			// 	return d12.NewPuzzle()
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
