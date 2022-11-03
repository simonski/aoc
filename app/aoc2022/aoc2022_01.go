package aoc2022

import "github.com/simonski/aoc/utils"

/*
--- Day 01: Description ---

*/

func (app *Application) Y2022D01_Summary() *utils.Summary {
	s := utils.NewSummary(2022, 1)
	s.Name = "Unknown"
	s.ProgressP1 = utils.Started
	s.ProgressP2 = utils.NotStarted

	entry := &utils.Entry{}
	entry.Date = "2022-11-03"
	entry.Title = "First entry."
	entry.Notes = "This is the first entry in the blog."
	entry.Summary = s
	s.Entries = append(s.Entries, entry)
	return s
}

// rename this to the year and day in question
func (app *Application) Y2022D01P1() {
}

// rename this to the year and day in question
func (app *Application) Y2022D01P2() {
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
