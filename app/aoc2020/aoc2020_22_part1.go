package aoc2020

/*
--- Day 22: Crab Combat ---
It only takes a few hours of sailing the ocean on a raft for boredom to sink in. Fortunately, you brought a small deck of space cards! You'd like to play a game of Combat, and there's even an opponent available: a small crab that climbed aboard your raft before you left.

Fortunately, it doesn't take long to teach the crab the rules.

Before the game starts, split the cards so each player has their own deck (your puzzle input). Then, the game consists of a series of rounds: both players draw their top card, and the player with the higher-valued card wins the round. The winner keeps both cards, placing them on the bottom of their own deck so that the winner's card is above the other card. If this causes a player to have all of the cards, they win, and the game ends.

For example, consider the following starting decks:

Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10
This arrangement means that player 1's deck contains 5 cards, with 9 on top and 1 on the bottom; player 2's deck also contains 5 cards, with 5 on top and 10 on the bottom.

The first round begins with both players drawing the top card of their decks: 9 and 5. Player 1 has the higher card, so both cards move to the bottom of player 1's deck such that 9 is above 5. In total, it takes 29 rounds before a player has all of the cards:

-- Round 1 --
Player 1's deck: 9, 2, 6, 3, 1
Player 2's deck: 5, 8, 4, 7, 10
Player 1 plays: 9
Player 2 plays: 5
Player 1 wins the round!

-- Round 2 --
Player 1's deck: 2, 6, 3, 1, 9, 5
Player 2's deck: 8, 4, 7, 10
Player 1 plays: 2
Player 2 plays: 8
Player 2 wins the round!

-- Round 3 --
Player 1's deck: 6, 3, 1, 9, 5
Player 2's deck: 4, 7, 10, 8, 2
Player 1 plays: 6
Player 2 plays: 4
Player 1 wins the round!

-- Round 4 --
Player 1's deck: 3, 1, 9, 5, 6, 4
Player 2's deck: 7, 10, 8, 2
Player 1 plays: 3
Player 2 plays: 7
Player 2 wins the round!

-- Round 5 --
Player 1's deck: 1, 9, 5, 6, 4
Player 2's deck: 10, 8, 2, 7, 3
Player 1 plays: 1
Player 2 plays: 10
Player 2 wins the round!

...several more rounds pass...

-- Round 27 --
Player 1's deck: 5, 4, 1
Player 2's deck: 8, 9, 7, 3, 2, 10, 6
Player 1 plays: 5
Player 2 plays: 8
Player 2 wins the round!

-- Round 28 --
Player 1's deck: 4, 1
Player 2's deck: 9, 7, 3, 2, 10, 6, 8, 5
Player 1 plays: 4
Player 2 plays: 9
Player 2 wins the round!

-- Round 29 --
Player 1's deck: 1
Player 2's deck: 7, 3, 2, 10, 6, 8, 5, 9, 4
Player 1 plays: 1
Player 2 plays: 7
Player 2 wins the round!


== Post-game results ==
Player 1's deck:
Player 2's deck: 3, 2, 10, 6, 8, 5, 9, 4, 7, 1
Once the game ends, you can calculate the winning player's score. The bottom card in their deck is worth the value of the card multiplied by 1, the second-from-the-bottom card is worth the value of the card multiplied by 2, and so on. With 10 cards, the top card is worth the value on the card multiplied by 10. In this example, the winning player's score is:

   3 * 10
+  2 *  9
+ 10 *  8
+  6 *  7
+  8 *  6
+  5 *  5
+  9 *  4
+  4 *  3
+  7 *  2
+  1 *  1
= 306
So, once the game ends, the winning player's score is 306.

Play the small crab in a game of Combat using the two decks you just dealt. What is the winning player's score?*/

import (
	"fmt"
	"strconv"
	"strings"
)

func (app *Application) Y2020D22() {
	app.Y2020D22P1()
	app.Y2020D22P2()
}

func (app *Application) Y2020D22P1() {
}
func (app *Application) Y2020D22P2() {
}

type Combat struct {
	Player1 *Player
	Player2 *Player
	Round   int
}

func (c *Combat) Play() {
	for {
		p1 := c.Player1.Draw()
		p2 := c.Player2.Draw()
		if p1 > p2 {
			c.Player1.AddCard(p1)
			c.Player1.AddCard(p2)
		} else {
			c.Player2.AddCard(p2)
			c.Player2.AddCard(p1)
		}
		c.Round++
		if c.Player1.Size() == 0 || c.Player2.Size() == 0 {
			break
		}
	}
}

func (c *Combat) GetRound() int {
	return c.Round
}

func (c *Combat) GetScore() int {
	return c.GetWinner().GetScore()
}

func (c *Combat) GetWinner() *Player {
	if c.Player1.Size() == 0 {
		return c.Player2
	} else {
		return c.Player1
	}
}

type Player struct {
	Cards []int
}

func (player *Player) AddCard(card int) {
	player.Cards = append(player.Cards, card)
}

func (player *Player) Draw() int {
	value := player.Cards[0]
	player.Cards = player.Cards[1:]
	return value
}

func (player *Player) Size() int {
	return len(player.Cards)
}

func (player *Player) GetScore() int {
	score := 0
	for index := 0; index < len(player.Cards); index++ {
		multiplier := index + 1
		cardIndex := len(player.Cards) - index - 1
		value := player.Cards[cardIndex] * multiplier
		score += value
	}
	return score
}

func NewCombat(cards string) *Combat {
	splits := strings.Split(cards, "\n")
	player1 := &Player{Cards: make([]int, 0)}
	player2 := &Player{Cards: make([]int, 0)}
	player := player1
	play := "Player 1"
	for index, line := range splits {
		line = strings.TrimSpace(line)
		if line == "Player 1:" {
			player = player1
			continue
		} else if line == "Player 2:" {
			fmt.Printf("[%v %v] '%v' is Player2\n", index, play, line)
			player = player2
			play = "Player 2"
			continue
		} else if strings.TrimSpace(line) == "" {
			fmt.Printf("[%v %v] '%v' is empty\n", index, play, line)
			continue
		}
		fmt.Printf("[%v %v] '%v' is value.\n", index, play, line)
		ival, _ := strconv.Atoi(line)
		player.AddCard(ival)
	}
	c := Combat{Player1: player1, Player2: player2}
	return &c
}
