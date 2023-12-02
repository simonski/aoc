package d02

import (
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
--- Day 05:  ---

*/

type Game struct {
	ID     int
	Line   string
	Rounds []*Round
}

type Round struct {
	Red   int
	Green int
	Blue  int
}

func NewGame(original string) *Game {
	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	splits := strings.Split(original, ":")
	g := Game{}
	gameID := splits[0]
	idStr := strings.Split(gameID, " ")[1]
	id, _ := strconv.Atoi(idStr)
	g.ID = id

	g.Line = original
	rounds := strings.Split(splits[1], ";")
	// fmt.Printf("rounds: '%v'\n", rounds)
	for _, round := range rounds {
		r := Round{}
		colorValues := strings.Split(round, ",")
		// fmt.Printf("colorValues: '%v'\n", colorValues)
		for _, colorValue := range colorValues {
			tokens := strings.Split(colorValue, " ")
			// fmt.Printf("tokens: '%v'\n", tokens)
			number, _ := strconv.Atoi(tokens[1])
			color := tokens[2]
			if color == "red" {
				r.Red = number
			} else if color == "blue" {
				r.Blue = number
			} else if color == "green" {
				r.Green = number
			}
		}
		g.Rounds = append(g.Rounds, &r)
	}

	return &g

}

func (g *Game) Passes(red int, green int, blue int) bool {
	for _, r := range g.Rounds {
		if r.Red <= red && r.Green <= green && r.Blue <= blue {
			// ok
		} else {
			return false
		}
	}
	return true

}

func (game *Game) Fewest() (int, int, int) {
	r := 0
	g := 0
	b := 0
	for _, round := range game.Rounds {
		r = utils.MaxInt(r, round.Red)
		g = utils.MaxInt(g, round.Green)
		b = utils.MaxInt(b, round.Blue)
	}
	return r, g, b

}

func (game *Game) Power() int {
	r, g, b := game.Fewest()
	return r * g * b
}
