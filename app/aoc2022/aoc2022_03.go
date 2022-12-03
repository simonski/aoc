package aoc2022

import (
	"fmt"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
--- Day 03: Description ---

*/

func (app *Application) Y2022D03_Summary() *utils.Summary {
	s := utils.Summary{}
	s.Name = "Rucksack Reorganization"
	s.Year = 2023
	s.Day = 03
	return &s
}

type Y2022D3 struct {
	Rucksacks []*Rucksack
}

type Rucksack struct {
	Left  *Compartment
	Right *Compartment
	Line  string
}

type Compartment struct {
	az []int
	AZ []int
}

func NewRucksack(line string) *Rucksack {
	rs := Rucksack{}
	line = strings.ReplaceAll(line, "\n", "")
	length := len(line)
	left := line[0 : length/2]
	right := line[length/2:]
	fmt.Printf("RS: line = %v, left=%v, right=%v\n", line, left, right)
	rs.Left = NewCompartment(left)
	rs.Right = NewCompartment(right)
	rs.Line = line
	return &rs
}

func NewCompartment(line string) *Compartment {
	rs := Compartment{}
	rs.az = make([]int, 26) // a..z 97 .. 122
	rs.AZ = make([]int, 26) // A..Z 65 .. 90
	runes := []rune(line)
	for index := 0; index < len(runes); index++ {
		// character := line[index : index+1]
		// fmt.Printf("line[%v] = %v, rune=%v\n", index, character, runevalue)
		runevalue := runes[index]

		if runevalue >= 97 {
			rune_index := runevalue - 97
			rs.az[rune_index] += 1
		} else {
			rune_index := runevalue - 65
			rs.AZ[rune_index] += 1
		}
	}

	return &rs
}

func (rs *Rucksack) Frequency() []int {
	// rs.az = make([]int, 26) // a..z 97 .. 122
	// rs.AZ = make([]int, 26) // A..Z 65 .. 90

	result := make([]int, 52)
	runes := []rune(rs.Line)
	for index := 0; index < len(runes); index++ {
		// character := line[index : index+1]
		// fmt.Printf("line[%v] = %v, rune=%v\n", index, character, runevalue)
		runevalue := runes[index]

		if runevalue >= 97 {
			rune_index := runevalue - 97
			result[rune_index] += 1
		} else {
			rune_index := runevalue - 65
			result[rune_index+26] += 1
		}
	}

	return result
}
func (rs *Rucksack) Common() []int {
	results := make([]int, 52)
	for index := 0; index < 26; index++ {
		if rs.Left.az[index] > 0 && rs.Right.az[index] > 0 {
			// fmt.Printf("common: az[%v]\n", index)
			results[index] = 1
		}
	}
	for index := 0; index < 26; index++ {
		if rs.Left.AZ[index] > 0 && rs.Right.AZ[index] > 0 {
			// fmt.Printf("common: AZ[%v]\n", index)
			results[index+26] = 1
		}
	}
	return results
}

const AZ = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func (rs *Rucksack) DebugCommon() {
	common := rs.Common()
	for index := 0; index < len(common); index++ {
		if common[index] > 0 {
			fmt.Printf("line %v shares %v\n", rs.Line, AZ[index:index+1])
		}
	}
}

func (rs *Rucksack) SumCommon() int {
	common := rs.Common()
	total := 0
	for index := 0; index < len(common); index++ {
		if common[index] > 0 {
			total += (index + 1)
		}
	}
	return total
}

func IntValue(c string) int {
	runes := []rune(c)
	rvar := int(runes[0])
	if rvar >= 97 {
		return (rvar - 65) + 1
	} else {
		return (rvar - 65) + 26
	}

}

func (logic Y2022D3) Load(data string) {
	splits := strings.Split(data, "\n")
	rucksacks := make([]*Rucksack, 0)
	for _, line := range splits {
		rs := NewRucksack(line)
		rucksacks = append(rucksacks, rs)
	}
	logic.Rucksacks = rucksacks
}

func (app *Application) Y2022D03P1() {
}

// rename this to the year and day in question
func (app *Application) Y2022D03P2() {
}

// this is what we will reflect and call - so both parts with run. It's up to you to make it print nicely etc.
func (app *Application) Y2022D03() {
	app.Y2022D03P1()
	app.Y2022D03P2()
}
