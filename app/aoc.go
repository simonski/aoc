package app

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gookit/color"
	aoc2015 "github.com/simonski/aoc/app/aoc2015"
	"github.com/simonski/aoc/app/aoc2016"
	"github.com/simonski/aoc/app/aoc2017"
	"github.com/simonski/aoc/app/aoc2018"
	"github.com/simonski/aoc/app/aoc2019"
	"github.com/simonski/aoc/app/aoc2020"
	"github.com/simonski/aoc/app/aoc2021"
	"github.com/simonski/aoc/app/aoc2022"
	aoc2023d01 "github.com/simonski/aoc/app/aoc2023/d01"
	aoc2023d02 "github.com/simonski/aoc/app/aoc2023/d02"
	"github.com/simonski/aoc/app/constants"
	"github.com/simonski/aoc/utils"
	cli "github.com/simonski/cli"
	goutils "github.com/simonski/goutils"
)

type AOC struct {
	CLI     *cli.CLI
	Verbose bool
}

func NewAOC(cli *cli.CLI) AOC {
	app := AOC{CLI: cli}
	app.Verbose = cli.IndexOf("-v") > -1
	return app
}

func (app *AOC) Run(cli *cli.CLI) {
	USAGE := "Usage: aoc run (year) (day)"
	year := cli.GetStringOrDefault("run", "")
	if year == "" {
		fmt.Printf("%v\n", USAGE)
		os.Exit(1)
	}
	iyear, _ := strconv.Atoi(year)
	if iyear < 2022 {
		al := app.GetAppLogic(iyear)
		if al == nil {
			fmt.Printf("Sorry, we don't have year %v\n", year)
			return
		}
		al.Run(cli)
	} else {
		// iyear, _ := strconv.Atoi(year)
		day := cli.GetStringOrDie(year)
		// iday, _ := strconv.Atoi(day)
		puzzle := app.GetPuzzle(year, day)
		if puzzle == nil {
			fmt.Printf("Sorry, we don't have year %v\n", year)
			return
		} else {
			puzzle.Run()
		}

	}
}

func (app *AOC) Summary(cli *cli.CLI) {
	USAGE := "Usage: aoc summary (year) (day)"
	year := cli.GetStringOrDefault("summary", "")
	if year == "" {
		fmt.Printf("%v\n", USAGE)
		os.Exit(1)
	}
	iyear, _ := strconv.Atoi(year)

	day := cli.GetStringOrDefault(year, "")
	iday, _ := strconv.Atoi(day)

	al := app.GetAppLogic(iyear)
	if al == nil {
		fmt.Printf("Sorry, we don't have year %v\n", year)
		return
	}
	summary := al.Summary(iyear, iday)
	if summary != nil {
		j, _ := json.MarshalIndent(summary, "", "  ")
		fmt.Printf("%v\n", string(j))
	} else {
		fmt.Println("No summary.")
	}
}

func (app *AOC) Render(cli *cli.CLI) {
	USAGE := "Usage: aoc render (year) (day)"
	year := cli.GetStringOrDefault("render", "")
	if year == "" {
		fmt.Printf("%v\n", USAGE)
		os.Exit(1)
	}
	iyear, _ := strconv.Atoi(year)
	al := app.GetAppLogic(iyear)
	al.Render(cli)
}

func (app *AOC) Help(cli *cli.CLI) {
	output := strings.ReplaceAll(constants.USAGE_OVERALL, "LIST", app.List())
	fmt.Println(output)
	command := cli.GetStringOrDefault("help", "")
	if command != "" {
		fmt.Printf("aoc help %s: unknown help topic. Run `aoc help`.\n", command)
		os.Exit(1)
	}
}

// builds a table of progress for all problems for the terminal
func (a *AOC) List() string {

	output := ""

	max_year := constants.MAX_YEAR
	min_year := constants.MIN_YEAR

	tick := "\u2713"
	star := "\u2606"

	tick = star
	cross := " " // \u2717"

	dash := "\u2501"

	header := "\u2503           \u2503"
	for day := 1; day <= 25; day++ {
		header = fmt.Sprintf("%v%02d\u2503", header, day)
	}

	//https://en.wikipedia.org/wiki/Box-drawing_character
	biggydown := "\u2533" + dash + dash
	biggyup := "\u2543" + dash + dash
	biggyuponly := "\u253b" + dash + dash

	// repeat := 80
	// bigline1 := "\u250f" + repeatstring(dash, repeat) + "\u2513"
	bigline1 := "\u250F" + goutils.Repeatstring(dash, 10) + goutils.Repeatstring(biggydown, 24) + "\u2513"
	// bigline2 := "\u2523" + repeatstring(dash, repeat) + "\u252b"
	bigline2 := "\u2523" + goutils.Repeatstring(dash, 10) + goutils.Repeatstring(biggyup, 24) + "\u252b"
	bigline3 := "\u2517" + goutils.Repeatstring(dash, 10) + goutils.Repeatstring(biggyuponly, 24) + "\u251b"

	output += bigline1
	output += "\n"
	output += header
	output += "\n"
	output += bigline2
	output += "\n"
	totalStars := 0
	totalAvailableStars := 0
	for year := max_year; year >= min_year; year-- {
		app := a.GetAppLogic(year)
		if app == nil {
			fmt.Printf("AppLogic for year %v does not exist.\n", year)
			continue
		}
		line := ""
		stars := 0
		totalAvailableStars += 25
		for day := 1; day <= 25; day++ {
			methodNamePart1 := fmt.Sprintf("Y%vD%02dP1", year, day)
			methodNamePart2 := fmt.Sprintf("Y%vD%02dP2", year, day)
			methodNamePart1Render := fmt.Sprintf("Y%vD%02dP1Render", year, day)
			methodNamePart2Render := fmt.Sprintf("Y%vD%02dP2Render", year, day)

			_, _, m1exists := app.GetMethod(methodNamePart1)
			_, _, m2exists := app.GetMethod(methodNamePart2)
			_, _, m1existsRender := app.GetMethod(methodNamePart1Render)
			_, _, m2existsRender := app.GetMethod(methodNamePart2Render)

			part1 := cross
			part2 := cross

			gold := color.FgYellow.Render

			if m1exists {
				part1 = tick
				stars++
			} else {
				part1 = cross
			}

			if m2exists {
				part2 = tick
				stars++
			} else {
				part2 = cross
			}

			if m1existsRender {
				part1 = gold(part1)
			}
			if m2existsRender {
				part2 = gold(part2)
			}

			line = fmt.Sprintf("%v%v%v\u2503", line, part1, part2)

		}
		totalStars += stars
		line = fmt.Sprintf("\u2503 %04d (%02v) \u2503", year, stars) + line

		output += line
		output += "\n"
	}
	output += bigline3
	output += "\n"
	output += fmt.Sprintf("(Total Stars %v/%v)\n", totalStars, totalAvailableStars)

	return output

}

func (a *AOC) GetPuzzle(year string, day string) utils.Puzzle {

	// if year == "2015" {
	// 	if day == "01" {
	// 		return aoc2015d01.NewPuzzle()
	// 	} else if day == "02" {
	// 		return aoc2015d02.NewPuzzle()
	// 	} else if day == "03" {
	// 		return aoc2015d03.NewPuzzle()
	// 	} else if day == "04" {
	// 		return aoc2015d04.NewPuzzle()
	// 	} else if day == "05" {
	// 		return aoc2015d05.NewPuzzle()
	// 	} else if day == "06" {
	// 		return aoc2015d06.NewPuzzle()
	// 	} else if day == "07" {
	// 		return aoc2015d07.NewPuzzle()
	// 	} else if day == "08" {
	// 		return aoc2015d08.NewPuzzle()
	// 	} else if day == "09" {
	// 		return aoc2015d09.NewPuzzle()
	// 	} else if day == "10" {
	// 		return aoc2015d10.NewPuzzle()
	// 	} else if day == "11" {
	// 		return aoc2015d11.NewPuzzle()
	// 	} else if day == "12" {
	// 		return aoc2015d12.NewPuzzle()
	// 	} else if day == "13" {
	// 		return aoc2015d13.NewPuzzle()
	// 	} else if day == "14" {
	// 		return aoc2015d14.NewPuzzle()
	// 	} else if day == "15" {
	// 		return aoc2015d15.NewPuzzle()
	// 	} else if day == "16" {
	// 		return aoc2015d16.NewPuzzle()
	// 	} else if day == "17" {
	// 		return aoc2015d17.NewPuzzle()
	// 	} else if day == "18" {
	// 		return aoc2015d18.NewPuzzle()
	// 	} else if day == "19" {
	// 		return aoc2015d19.NewPuzzle()
	// 	} else if day == "20" {
	// 		return aoc2015d20.NewPuzzle()
	// 	} else if day == "21" {
	// 		return aoc2015d21.NewPuzzle()
	// 	} else if day == "22" {
	// 		return aoc2015d22.NewPuzzle()
	// 	} else if day == "23" {
	// 		return aoc2015d23.NewPuzzle()
	// 	} else if day == "24" {
	// 		return aoc2015d24.NewPuzzle()
	// 	} else if day == "25" {
	// 		return aoc2015d25.NewPuzzle()
	// 	}
	// }

	// if year == "2016" {
	// 	if day == "01" {
	// 		return aoc2016d01.NewPuzzle()
	// 	} else if day == "02" {
	// 		return aoc2016d02.NewPuzzle()
	// 	} else if day == "03" {
	// 		return aoc2016d03.NewPuzzle()
	// 	} else if day == "04" {
	// 		return aoc2016d04.NewPuzzle()
	// 	} else if day == "05" {
	// 		return aoc2016d05.NewPuzzle()
	// 	} else if day == "06" {
	// 		return aoc2016d06.NewPuzzle()
	// 	} else if day == "07" {
	// 		return aoc2016d07.NewPuzzle()
	// 	} else if day == "08" {
	// 		return aoc2016d08.NewPuzzle()
	// 	} else if day == "09" {
	// 		return aoc2016d09.NewPuzzle()
	// 	} else if day == "10" {
	// 		return aoc2016d10.NewPuzzle()
	// 	} else if day == "11" {
	// 		return aoc2016d11.NewPuzzle()
	// 	} else if day == "12" {
	// 		return aoc2016d12.NewPuzzle()
	// 	} else if day == "13" {
	// 		return aoc2016d13.NewPuzzle()
	// 	} else if day == "14" {
	// 		return aoc2016d14.NewPuzzle()
	// 	} else if day == "15" {
	// 		return aoc2016d15.NewPuzzle()
	// 	} else if day == "16" {
	// 		return aoc2016d16.NewPuzzle()
	// 	} else if day == "17" {
	// 		return aoc2016d17.NewPuzzle()
	// 	} else if day == "18" {
	// 		return aoc2016d18.NewPuzzle()
	// 	} else if day == "19" {
	// 		return aoc2016d19.NewPuzzle()
	// 	} else if day == "20" {
	// 		return aoc2016d20.NewPuzzle()
	// 	} else if day == "21" {
	// 		return aoc2016d21.NewPuzzle()
	// 	} else if day == "22" {
	// 		return aoc2016d22.NewPuzzle()
	// 	} else if day == "23" {
	// 		return aoc2016d23.NewPuzzle()
	// 	} else if day == "24" {
	// 		return aoc2016d24.NewPuzzle()
	// 	} else if day == "25" {
	// 		return aoc2016d25.NewPuzzle()
	// 	}
	// }

	// if year == "2017" {
	// 	if day == "01" {
	// 		return aoc2017d01.NewPuzzle()
	// 	} else if day == "02" {
	// 		return aoc2017d02.NewPuzzle()
	// 	} else if day == "03" {
	// 		return aoc2017d03.NewPuzzle()
	// 	} else if day == "04" {
	// 		return aoc2017d04.NewPuzzle()
	// 	} else if day == "05" {
	// 		return aoc2017d05.NewPuzzle()
	// 	} else if day == "06" {
	// 		return aoc2017d06.NewPuzzle()
	// 	} else if day == "07" {
	// 		return aoc2017d07.NewPuzzle()
	// 	} else if day == "08" {
	// 		return aoc2017d08.NewPuzzle()
	// 	} else if day == "09" {
	// 		return aoc2017d09.NewPuzzle()
	// 	} else if day == "10" {
	// 		return aoc2017d10.NewPuzzle()
	// 	} else if day == "11" {
	// 		return aoc2017d11.NewPuzzle()
	// 	} else if day == "12" {
	// 		return aoc2017d12.NewPuzzle()
	// 	} else if day == "13" {
	// 		return aoc2017d13.NewPuzzle()
	// 	} else if day == "14" {
	// 		return aoc2017d14.NewPuzzle()
	// 	} else if day == "15" {
	// 		return aoc2017d15.NewPuzzle()
	// 	} else if day == "16" {
	// 		return aoc2017d16.NewPuzzle()
	// 	} else if day == "17" {
	// 		return aoc2017d17.NewPuzzle()
	// 	} else if day == "18" {
	// 		return aoc2017d18.NewPuzzle()
	// 	} else if day == "19" {
	// 		return aoc2017d19.NewPuzzle()
	// 	} else if day == "20" {
	// 		return aoc2017d20.NewPuzzle()
	// 	} else if day == "21" {
	// 		return aoc2017d21.NewPuzzle()
	// 	} else if day == "22" {
	// 		return aoc2017d22.NewPuzzle()
	// 	} else if day == "23" {
	// 		return aoc2017d23.NewPuzzle()
	// 	} else if day == "24" {
	// 		return aoc2017d24.NewPuzzle()
	// 	} else if day == "25" {
	// 		return aoc2017d25.NewPuzzle()
	// 	}
	// }

	// if year == "2018" {
	// 	if day == "01" {
	// 		return aoc2018d01.NewPuzzle()
	// 	} else if day == "02" {
	// 		return aoc2018d02.NewPuzzle()
	// 	} else if day == "03" {
	// 		return aoc2018d03.NewPuzzle()
	// 	} else if day == "04" {
	// 		return aoc2018d04.NewPuzzle()
	// 	} else if day == "05" {
	// 		return aoc2018d05.NewPuzzle()
	// 	} else if day == "06" {
	// 		return aoc2018d06.NewPuzzle()
	// 	} else if day == "07" {
	// 		return aoc2018d07.NewPuzzle()
	// 	} else if day == "08" {
	// 		return aoc2018d08.NewPuzzle()
	// 	} else if day == "09" {
	// 		return aoc2018d09.NewPuzzle()
	// 	} else if day == "10" {
	// 		return aoc2018d10.NewPuzzle()
	// 	} else if day == "11" {
	// 		return aoc2018d11.NewPuzzle()
	// 	} else if day == "12" {
	// 		return aoc2018d12.NewPuzzle()
	// 	} else if day == "13" {
	// 		return aoc2018d13.NewPuzzle()
	// 	} else if day == "14" {
	// 		return aoc2018d14.NewPuzzle()
	// 	} else if day == "15" {
	// 		return aoc2018d15.NewPuzzle()
	// 	} else if day == "16" {
	// 		return aoc2018d16.NewPuzzle()
	// 	} else if day == "17" {
	// 		return aoc2018d17.NewPuzzle()
	// 	} else if day == "18" {
	// 		return aoc2018d18.NewPuzzle()
	// 	} else if day == "19" {
	// 		return aoc2018d19.NewPuzzle()
	// 	} else if day == "20" {
	// 		return aoc2018d20.NewPuzzle()
	// 	} else if day == "21" {
	// 		return aoc2018d21.NewPuzzle()
	// 	} else if day == "22" {
	// 		return aoc2018d22.NewPuzzle()
	// 	} else if day == "23" {
	// 		return aoc2018d23.NewPuzzle()
	// 	} else if day == "24" {
	// 		return aoc2018d24.NewPuzzle()
	// 	} else if day == "25" {
	// 		return aoc2018d25.NewPuzzle()
	// 	}
	// }

	// if year == "2019" {
	// 	if day == "01" {
	// 		return aoc2019d01.NewPuzzle()
	// 	} else if day == "02" {
	// 		return aoc2019d02.NewPuzzle()
	// 	} else if day == "03" {
	// 		return aoc2019d03.NewPuzzle()
	// 	} else if day == "04" {
	// 		return aoc2019d04.NewPuzzle()
	// 	} else if day == "05" {
	// 		return aoc2019d05.NewPuzzle()
	// 	} else if day == "06" {
	// 		return aoc2019d06.NewPuzzle()
	// 	} else if day == "07" {
	// 		return aoc2019d07.NewPuzzle()
	// 	} else if day == "08" {
	// 		return aoc2019d08.NewPuzzle()
	// 	} else if day == "09" {
	// 		return aoc2019d09.NewPuzzle()
	// 	} else if day == "10" {
	// 		return aoc2019d10.NewPuzzle()
	// 	} else if day == "11" {
	// 		return aoc2019d11.NewPuzzle()
	// 	} else if day == "12" {
	// 		return aoc2019d12.NewPuzzle()
	// 	} else if day == "13" {
	// 		return aoc2019d13.NewPuzzle()
	// 	} else if day == "14" {
	// 		return aoc2019d14.NewPuzzle()
	// 	} else if day == "15" {
	// 		return aoc2019d15.NewPuzzle()
	// 	} else if day == "16" {
	// 		return aoc2019d16.NewPuzzle()
	// 	} else if day == "17" {
	// 		return aoc2019d17.NewPuzzle()
	// 	} else if day == "18" {
	// 		return aoc2019d18.NewPuzzle()
	// 	} else if day == "19" {
	// 		return aoc2019d19.NewPuzzle()
	// 	} else if day == "20" {
	// 		return aoc2019d20.NewPuzzle()
	// 	} else if day == "21" {
	// 		return aoc2019d21.NewPuzzle()
	// 	} else if day == "22" {
	// 		return aoc2019d22.NewPuzzle()
	// 	} else if day == "23" {
	// 		return aoc2019d23.NewPuzzle()
	// 	} else if day == "24" {
	// 		return aoc2019d24.NewPuzzle()
	// 	} else if day == "25" {
	// 		return aoc2019d25.NewPuzzle()
	// 	}
	// }

	// if year == "2020" {
	// 	if day == "01" {
	// 		return aoc2020d01.NewPuzzle()
	// 	} else if day == "02" {
	// 		return aoc2020d02.NewPuzzle()
	// 	} else if day == "03" {
	// 		return aoc2020d03.NewPuzzle()
	// 	} else if day == "04" {
	// 		return aoc2020d04.NewPuzzle()
	// 	} else if day == "05" {
	// 		return aoc2020d05.NewPuzzle()
	// 	} else if day == "06" {
	// 		return aoc2020d06.NewPuzzle()
	// 	} else if day == "07" {
	// 		return aoc2020d07.NewPuzzle()
	// 	} else if day == "08" {
	// 		return aoc2020d08.NewPuzzle()
	// 	} else if day == "09" {
	// 		return aoc2020d09.NewPuzzle()
	// 	} else if day == "10" {
	// 		return aoc2020d10.NewPuzzle()
	// 	} else if day == "11" {
	// 		return aoc2020d11.NewPuzzle()
	// 	} else if day == "12" {
	// 		return aoc2020d12.NewPuzzle()
	// 	} else if day == "13" {
	// 		return aoc2020d13.NewPuzzle()
	// 	} else if day == "14" {
	// 		return aoc2020d14.NewPuzzle()
	// 	} else if day == "15" {
	// 		return aoc2020d15.NewPuzzle()
	// 	} else if day == "16" {
	// 		return aoc2020d16.NewPuzzle()
	// 	} else if day == "17" {
	// 		return aoc2020d17.NewPuzzle()
	// 	} else if day == "18" {
	// 		return aoc2020d18.NewPuzzle()
	// 	} else if day == "19" {
	// 		return aoc2020d19.NewPuzzle()
	// 	} else if day == "20" {
	// 		return aoc2020d20.NewPuzzle()
	// 	} else if day == "21" {
	// 		return aoc2020d21.NewPuzzle()
	// 	} else if day == "22" {
	// 		return aoc2020d22.NewPuzzle()
	// 	} else if day == "23" {
	// 		return aoc2020d23.NewPuzzle()
	// 	} else if day == "24" {
	// 		return aoc2020d24.NewPuzzle()
	// 	} else if day == "25" {
	// 		return aoc2020d25.NewPuzzle()
	// 	}
	// }

	// if year == "2021" {
	// 	if day == "01" {
	// 		return aoc2021d01.NewPuzzle()
	// 	} else if day == "02" {
	// 		return aoc2021d02.NewPuzzle()
	// 	} else if day == "03" {
	// 		return aoc2021d03.NewPuzzle()
	// 	} else if day == "04" {
	// 		return aoc2021d04.NewPuzzle()
	// 	} else if day == "05" {
	// 		return aoc2021d05.NewPuzzle()
	// 	} else if day == "06" {
	// 		return aoc2021d06.NewPuzzle()
	// 	} else if day == "07" {
	// 		return aoc2021d07.NewPuzzle()
	// 	} else if day == "08" {
	// 		return aoc2021d08.NewPuzzle()
	// 	} else if day == "09" {
	// 		return aoc2021d09.NewPuzzle()
	// 	} else if day == "10" {
	// 		return aoc2021d10.NewPuzzle()
	// 	} else if day == "11" {
	// 		return aoc2021d11.NewPuzzle()
	// 	} else if day == "12" {
	// 		return aoc2021d12.NewPuzzle()
	// 	} else if day == "13" {
	// 		return aoc2021d13.NewPuzzle()
	// 	} else if day == "14" {
	// 		return aoc2021d14.NewPuzzle()
	// 	} else if day == "15" {
	// 		return aoc2021d15.NewPuzzle()
	// 	} else if day == "16" {
	// 		return aoc2021d16.NewPuzzle()
	// 	} else if day == "17" {
	// 		return aoc2021d17.NewPuzzle()
	// 	} else if day == "18" {
	// 		return aoc2021d18.NewPuzzle()
	// 	} else if day == "19" {
	// 		return aoc2021d19.NewPuzzle()
	// 	} else if day == "20" {
	// 		return aoc2021d20.NewPuzzle()
	// 	} else if day == "21" {
	// 		return aoc2021d21.NewPuzzle()
	// 	} else if day == "22" {
	// 		return aoc2021d22.NewPuzzle()
	// 	} else if day == "23" {
	// 		return aoc2021d23.NewPuzzle()
	// 	} else if day == "24" {
	// 		return aoc2021d24.NewPuzzle()
	// 	} else if day == "25" {
	// 		return aoc2021d25.NewPuzzle()
	// 	}
	// }

	// if year == "2022" {
	// 	if day == "01" {
	// 		return aoc2022d01.NewPuzzle()
	// 	} else if day == "02" {
	// 		return aoc2022d02.NewPuzzle()
	// 	} else if day == "03" {
	// 		return aoc2022d03.NewPuzzle()
	// 	} else if day == "04" {
	// 		return aoc2022d04.NewPuzzle()
	// 	} else if day == "05" {
	// 		return aoc2022d05.NewPuzzle()
	// 	} else if day == "06" {
	// 		return aoc2022d06.NewPuzzle()
	// 	} else if day == "07" {
	// 		return aoc2022d07.NewPuzzle()
	// 	} else if day == "08" {
	// 		return aoc2022d08.NewPuzzle()
	// 	} else if day == "09" {
	// 		return aoc2022d09.NewPuzzle()
	// 	} else if day == "10" {
	// 		return aoc2022d10.NewPuzzle()
	// 	} else if day == "11" {
	// 		return aoc2022d11.NewPuzzle()
	// 	} else if day == "12" {
	// 		return aoc2022d12.NewPuzzle()
	// 	} else if day == "13" {
	// 		return aoc2022d13.NewPuzzle()
	// 	} else if day == "14" {
	// 		return aoc2022d14.NewPuzzle()
	// 	} else if day == "15" {
	// 		return aoc2022d15.NewPuzzle()
	// 	} else if day == "16" {
	// 		return aoc2022d16.NewPuzzle()
	// 	} else if day == "17" {
	// 		return aoc2022d17.NewPuzzle()
	// 	} else if day == "18" {
	// 		return aoc2022d18.NewPuzzle()
	// 	} else if day == "19" {
	// 		return aoc2022d19.NewPuzzle()
	// 	} else if day == "20" {
	// 		return aoc2022d20.NewPuzzle()
	// 	} else if day == "21" {
	// 		return aoc2022d21.NewPuzzle()
	// 	} else if day == "22" {
	// 		return aoc2022d22.NewPuzzle()
	// 	} else if day == "23" {
	// 		return aoc2022d23.NewPuzzle()
	// 	} else if day == "24" {
	// 		return aoc2022d24.NewPuzzle()
	// 	} else if day == "25" {
	// 		return aoc2022d25.NewPuzzle()
	// 	}
	// }

	if year == "2023" {
		if day == "01" {
			return aoc2023d01.NewPuzzle()
		} else if day == "02" {
			return aoc2023d02.NewPuzzle()
			// } else if day == "03" {
			// 	return aoc2023d03.NewPuzzle()
			// } else if day == "04" {
			// 	return aoc2023d04.NewPuzzle()
			// } else if day == "05" {
			// 	return aoc2023d05.NewPuzzle()
			// } else if day == "06" {
			// 	return aoc2023d06.NewPuzzle()
			// } else if day == "07" {
			// 	return aoc2023d07.NewPuzzle()
			// } else if day == "08" {
			// 	return aoc2023d08.NewPuzzle()
			// } else if day == "09" {
			// 	return aoc2023d09.NewPuzzle()
			// } else if day == "10" {
			// 	return aoc2023d10.NewPuzzle()
			// } else if day == "11" {
			// 	return aoc2023d11.NewPuzzle()
			// } else if day == "12" {
			// 	return aoc2023d12.NewPuzzle()
			// } else if day == "13" {
			// 	return aoc2023d13.NewPuzzle()
			// } else if day == "14" {
			// 	return aoc2023d14.NewPuzzle()
			// } else if day == "15" {
			// 	return aoc2023d15.NewPuzzle()
			// } else if day == "16" {
			// 	return aoc2023d16.NewPuzzle()
			// } else if day == "17" {
			// 	return aoc2023d17.NewPuzzle()
			// } else if day == "18" {
			// 	return aoc2023d18.NewPuzzle()
			// } else if day == "19" {
			// 	return aoc2023d19.NewPuzzle()
			// } else if day == "20" {
			// 	return aoc2023d20.NewPuzzle()
			// } else if day == "21" {
			// 	return aoc2023d21.NewPuzzle()
			// } else if day == "22" {
			// 	return aoc2023d22.NewPuzzle()
			// } else if day == "23" {
			// 	return aoc2023d23.NewPuzzle()
			// } else if day == "24" {
			// 	return aoc2023d24.NewPuzzle()
			// } else if day == "25" {
			// 	return aoc2023d25.NewPuzzle()
		}
	}

	return nil
	// concreteTypeName := fmt.Sprintf("app.aoc%v.d%v.Puzzle", year, day)
	// puzzle := reflect.TypeOf(concreteTypeName)
	// return &puzzle
}

func (a *AOC) GetAppLogic(year int) utils.AppLogic {
	if year == 2015 {
		return aoc2015.NewApplication(a.CLI)
	} else if year == 2016 {
		return aoc2016.NewApplication(a.CLI)
	} else if year == 2017 {
		return aoc2017.NewApplication(a.CLI)
	} else if year == 2018 {
		return aoc2018.NewApplication(a.CLI)
	} else if year == 2019 {
		return aoc2019.NewApplication(a.CLI)
	} else if year == 2020 {
		return aoc2020.NewApplication(a.CLI)
	} else if year == 2021 {
		return aoc2021.NewApplication(a.CLI)
	} else if year == 2022 {
		return aoc2022.NewApplication(a.CLI)
	} else {
		return nil
	}
}

func (a *AOC) GetSummary(year int, day int) *utils.Summary {
	appLogic := a.GetAppLogic(year)
	if appLogic != nil {
		return appLogic.Summary(year, day)
	} else {
		return nil
	}
}

func (a *AOC) GetSummaries() []*utils.Summary {
	results := make([]*utils.Summary, 0)
	for year := constants.MIN_YEAR; year <= constants.MAX_YEAR; year++ {
		for day := 1; day <= 25; day++ {
			summary := a.GetSummary(year, day)
			results = append(results, summary)
		}
	}
	return results
}
