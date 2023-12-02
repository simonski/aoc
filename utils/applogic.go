package utils

import (
	"reflect"

	"github.com/simonski/cli"
)

type AppLogic interface {
	Run(cli *cli.CLI)
	Render(cli *cli.CLI)
	Help(cli *cli.CLI)
	GetMethod(methodName string) (reflect.Value, reflect.Value, bool)
	GetName() string
	Api(day int) string
	Summary(year int, day int) *Summary
	GetPuzzle(year int, day int) Puzzle
}
