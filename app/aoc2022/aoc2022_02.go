package aoc2022

import (
	"fmt"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
--- Day 01: Description ---

*/

func (app *Application) Y2022D02_Summary() *utils.Summary {
	s := utils.NewSummary(2022, 2)
	s.Name = "Rock Paper Scissors"
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

type D2Game struct {
	data string
}

func (g *D2Game) PlayGame1(game string) int {
	splits := strings.Split(game, "\n")
	overall := 0
	for _, round := range splits {
		score := g.PlayRound1(round)
		overall += score
	}
	return overall
}

func (g *D2Game) PlayGame2(game string) int {
	splits := strings.Split(game, "\n")
	overall := 0
	for _, round := range splits {
		score := g.PlayRound2(round)
		overall += score
	}
	return overall
}

func (g *D2Game) PlayRound1(round string) int {
	// fmt.Println(round)
	round = strings.TrimRight(round, "\n")
	actions := strings.Split(round, " ")
	p1_action := actions[0]
	p2_action := actions[1]

	p1Rock := p1_action == "A"
	p1Paper := p1_action == "B"
	p1Scissors := p1_action == "C"

	p2Rock := p2_action == "X"
	p2Paper := p2_action == "Y"
	p2Scissors := p2_action == "Z"

	p2Win := (p2Rock && p1Scissors) || (p2Scissors && p1Paper) || (p2Paper && p1Rock)
	p2Draw := (p2Rock && p1Rock) || (p2Scissors && p1Scissors) || (p2Paper && p1Paper)
	p2Loss := !p2Win && !p2Draw

	roundScore := 0
	if p2Win {
		roundScore += 6
	} else if p2Draw {
		roundScore += 3
	} else if p2Loss {
		roundScore += 0
	}

	if p2Rock {
		roundScore += 1
	} else if p2Paper {
		roundScore += 2
	} else if p2Scissors {
		roundScore += 3
	}

	return roundScore
}

func (g *D2Game) PlayRound2(round string) int {
	// fmt.Println(round)
	round = strings.TrimRight(round, "\n")
	actions := strings.Split(round, " ")
	p1_action := actions[0]
	p2_action := actions[1]

	p2_mustLose := p2_action == "X"
	p2_mustWin := p2_action == "Z"

	p1Rock := p1_action == "A"
	p1Paper := p1_action == "B"
	p1Scissors := p1_action == "C"

	p2Rock := false
	p2Paper := false
	p2Scissors := false

	if p2_mustLose {
		if p1Rock {
			p2Scissors = true
		} else if p1Paper {
			p2Rock = true
		} else if p1Scissors {
			p2Paper = true
		}
	} else if p2_mustWin {
		if p1Rock {
			p2Paper = true
		} else if p1Paper {
			p2Scissors = true
		} else if p1Scissors {
			p2Rock = true
		}

	} else {
		p2Rock = p1Rock
		p2Paper = p1Paper
		p2Scissors = p1Scissors
	}

	p2Win := (p2Rock && p1Scissors) || (p2Scissors && p1Paper) || (p2Paper && p1Rock)
	p2Draw := (p2Rock && p1Rock) || (p2Scissors && p1Scissors) || (p2Paper && p1Paper)
	p2Loss := !p2Win && !p2Draw

	roundScore := 0
	if p2Win {
		roundScore += 6
	} else if p2Draw {
		roundScore += 3
	} else if p2Loss {
		roundScore += 0
	}

	if p2Rock {
		roundScore += 1
	} else if p2Paper {
		roundScore += 2
	} else if p2Scissors {
		roundScore += 3
	}

	return roundScore
}

// rename this to the year and day in question
func (app *Application) Y2022D02P1() {
	game := D2Game{}

	score1 := game.PlayGame1(DAY_2022_02_TEST_DATA)
	fmt.Printf("Game1 score is %v\n", score1)

	score2 := game.PlayGame1(DAY_2022_02_DATA)
	fmt.Printf("Game1 score is %v\n", score2)
}

// rename this to the year and day in question
func (app *Application) Y2022D02P2() {
	game := D2Game{}

	score1 := game.PlayGame2(DAY_2022_02_TEST_DATA)
	fmt.Printf("Game2 score is %v\n", score1)

	score2 := game.PlayGame2(DAY_2022_02_DATA)
	fmt.Printf("Game2 score is %v\n", score2)
}

// this is what we will reflect and call - so both parts with run. It's up to you to make it print nicely etc.
// The app reference has a CLI for logging.
func (app *Application) Y2022D02() {
	app.Y2022D02P1()
	app.Y2022D02P2()
}
