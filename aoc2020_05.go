package main

import (
	"fmt"

	goutils "github.com/simonski/goutils"
)

// AOC_2020_05 is the entrypoint to the various attempts for day two
func AOC_2020_05(cli *goutils.CLI) {
	AOC_2020_05_part1_attempt1(cli)
}

func AOC_2020_05_part1_attempt1(cli *goutils.CLI) {
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

func LoadBoardingPassesFromFile(filename string) []*BoardingPass {
	results := load_file_to_strings(filename)
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
