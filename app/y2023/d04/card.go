package d4

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
Day 4: Scratchcards
*/

// "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1"
type Card struct {
	Line           string
	CardNumber     int
	WinningNumbers []int
}

func makeNumbers(ns []string) []int {
	results := make([]int, 0)
	for _, v := range ns {
		if v == " " || v == "" {
			continue
		}
		i, err := strconv.Atoi(v)
		if err == nil {
			results = append(results, i)
		}
	}
	sort.Ints(results)
	return results
}

func union(left []int, right []int) []int {
	retain := make([]int, 0)
	for il := 0; il < len(left); il++ {
		leftValue := left[il]
		for ir := 0; ir < len(right); ir++ {
			rightValue := right[ir]
			if leftValue == rightValue {
				retain = append(retain, leftValue)
				break
			}
		}
	}
	return retain
}

func NewCard(line string) *Card {

	line = strings.ReplaceAll(line, "  ", " ")
	line = strings.ReplaceAll(line, "\t", " ")

	card := Card{Line: line, WinningNumbers: make([]int, 0)}

	splits := strings.Split(line, ":")
	l := splits[0]
	r := splits[1]
	lpart := strings.ReplaceAll(l, "Card ", "")
	lpart = strings.ReplaceAll(lpart, " ", "")
	cardNumber, err := strconv.Atoi(lpart)
	card.CardNumber = cardNumber
	if err != nil {
		panic(err)
	}

	splits = strings.Split(r, "|")
	lnumbersStr := strings.Split(splits[0], " ")
	rnumbersStr := strings.Split(splits[1], " ")

	lnumbers := makeNumbers(lnumbersStr)
	rnumbers := makeNumbers(rnumbersStr)
	union := union(lnumbers, rnumbers)
	card.WinningNumbers = union
	return &card
}

func (c *Card) Score() int {
	// The first match makes the card worth one point and each match after the first doubles the point value of that card.
	if len(c.WinningNumbers) == 0 {
		return 0
	}
	score := 1
	for index := 0; index < len(c.WinningNumbers)-1; index++ {
		score *= 2
	}
	return score
}

func (c *Card) Debug() string {
	return fmt.Sprintf("Line='%v', Number='%v', winning='%v', score='%v'", c.Line, c.CardNumber, c.WinningNumbers, c.Score())
}
