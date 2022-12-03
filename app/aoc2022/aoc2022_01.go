package aoc2022

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
--- Day 01: Description ---

*/

func (app *Application) Y2022D01_Summary() *utils.Summary {
	s := utils.NewSummary(2022, 1)
	s.Name = "Calorie Counting"
	s.ProgressP1 = utils.Started
	s.ProgressP2 = utils.NotStarted

	// entry := &utils.Entry{}
	// entry.Date = "2022-12-01"
	// entry.Title = "First entry."
	// entry.Notes = "This is the first entry in the blog."
	// entry.Summary = s
	// s.Entries = append(s.Entries, entry)
	return s
}

// rename this to the year and day in question
func (app *Application) Y2022D01P1() {
	splits := strings.Split(DAY_2022_01_TEST_DATA, "\n")
	total := 0
	max := 0
	max_index := 0
	index := 0

	for _, line := range splits {
		fmt.Println(line)
		line = strings.ReplaceAll(line, "\n", " ")
		if line == "" {
			if total > max {
				max = total
				max_index = index
			}
			index += 1
			total = 0
		} else {
			value, _ := strconv.Atoi(line)
			total += value
		}
	}
	fmt.Printf("Max Elf: [%v], Max Value [%v]\n", max_index, max)

}

// rename this to the year and day in question
func (app *Application) Y2022D01P2() {
	splits := strings.Split(DAY_2022_01_TEST_DATA, "\n")

	totals := make([]int, 0)

	total := 0

	for _, line := range splits {
		fmt.Println(line)
		line = strings.ReplaceAll(line, "\n", " ")
		if line == "" {
			totals = append(totals, total)
			total = 0
		} else {
			value, _ := strconv.Atoi(line)
			total += value
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(totals)))

	fmt.Printf("Elf 1: %v\n", totals[0])
	fmt.Printf("Elf 2: %v\n", totals[1])
	fmt.Printf("Elf 3: %v\n", totals[2])
	top_3 := (totals[0] + totals[1] + totals[2])
	fmt.Printf("Total: %v\n", top_3)
}

// rename and uncomment this to the year and day in question once complete for a gold star!
func (app *Application) Y20XXDXXP1Render() {
}

// rename and uncomment this to the year and day in question once complete for a gold star!
func (app *Application) Y20XXDXXP2Render() {
}

// this is what we will reflect and call - so both parts with run. It's up to you to make it print nicely etc.
// The app reference has a CLI for logging.
func (app *Application) Y2022D01() {
	app.Y2022D01P1()
	app.Y2022D01P2()
}
