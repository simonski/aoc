package aoc2020

import (
	"fmt"
	"sort"

	"github.com/simonski/aoc/utils"
	cli "github.com/simonski/cli"
	goutils "github.com/simonski/goutils"
)

func (app *Application) Y2020D05_Summary() *utils.Summary {
	s := utils.NewSummary(2020, 5)
	s.Name = "Binary Boarding"
	s.ProgressP1 = utils.Completed
	s.ProgressP2 = utils.Completed
	return s
}

func (app *Application) Y2020D05() {
	app.Y2020D05P1()
	app.Y2020D05P2()
}

func (app *Application) Y2020D05P1() {
	app.AOC_2020_05_part1_attempt1(app.CLI)
}

func (app *Application) Y2020D05P2() {
	app.AOC_2020_05_part2_attempt1(app.CLI)
}

func (app *Application) AOC_2020_05_part1_attempt1(cli *cli.CLI) {
	filename := cli.GetFileExistsOrDie("-input")
	passes := LoadBoardingPassesFromFile(filename)
	maxPass := passes[0]
	maxSeatId := 0
	for index := 0; index < len(passes); index++ {
		pass := passes[index]
		seatId := pass.GetSeatId()
		if seatId > maxSeatId {
			maxSeatId = seatId
			maxPass = pass
		}
	}
	fmt.Printf("Highest seatId is %v, pass is %v", maxSeatId, maxPass.line)
}

func (app *Application) AOC_2020_05_part2_attempt1(cli *cli.CLI) {
	// so the seat ids with -1 and +1 will be present
	// only one will be missing
	// two ways of checking
	// brute force
	// attempt1: build list of ids, sort it, find first missing id
	// attempt2: something else

	filename := cli.GetFileExistsOrDie("-input")

	passes := LoadBoardingPassesFromFile(filename)

	// TODO look at sorting with functions
	// https://gobyexample.com/sorting-by-functions
	seatIds := make([]int, 10)
	for index := 0; index < len(passes); index++ {
		pass := passes[index]
		seatIds = append(seatIds, pass.GetSeatId())
	}
	sort.Ints(seatIds)

	// TODO improve on this performance
	seatId := -1
	for index := 0; index < len(seatIds); index++ {
		candidate := seatIds[index]
		// if the value of the next seat is +2 then the seatID is candidate+1
		nextValue := seatIds[index+1]
		fmt.Printf("%v,%v\n", candidate, nextValue)
		if nextValue == candidate+2 {
			seatId = candidate + 1
			break
		}
	}

	fmt.Printf("Seat Id is %v\n", seatId)

}

func LoadBoardingPassesFromFile(filename string) []*BoardingPass {
	results := goutils.Load_file_to_strings(filename)
	passes := make([]*BoardingPass, 0)
	for index := 0; index < len(results); index++ {
		line := results[index]
		bp := NewBoardingPass(line)
		passes = append(passes, bp)
	}
	return passes
}

type BoardingPass struct {
	line string
}

func (bp *BoardingPass) GetSeatId() int {
	return (bp.GetRow() * 8) + bp.GetCol()
}

func (bp *BoardingPass) GetRow() int {
	rows := bp.line[0:7]
	min_pos := 0
	max_pos := 127
	fmt.Printf("GetRow() line=%v, rows=%v\n", bp.line, rows)
	for index := 0; index < len(rows); index++ {
		instruction := rows[index : index+1]
		diff := (max_pos - min_pos) / 2
		if instruction == "F" {
			// take lower half
			max_pos = min_pos + diff
		} else if instruction == "B" {
			// take the upper half
			min_pos = max_pos - diff
		}
		fmt.Printf("[%v]: %v (min/max) (%v/%v)\n", index, instruction, min_pos, max_pos)
	}
	return min_pos
}

func (bp *BoardingPass) GetCol() int {
	rows := bp.line[7:10]
	min_pos := 0
	max_pos := 7
	fmt.Printf("GetCol() line=%v, rows=%v\n", bp.line, rows)
	for index := 0; index < len(rows); index++ {
		instruction := rows[index : index+1]
		diff := (max_pos - min_pos) / 2
		if instruction == "L" {
			// take lower half
			max_pos = min_pos + diff
		} else if instruction == "R" {
			// take the upper half
			min_pos = max_pos - diff
		}
		fmt.Printf("[%v]: %v (min/max) (%v/%v)\n", index, instruction, min_pos, max_pos)
	}
	return max_pos
}

func NewBoardingPass(line string) *BoardingPass {
	bp := BoardingPass{line: line}
	return &bp
}
