package d4

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
Day 4: Scratchcards
*/

type Puzzle struct {
	title string
	year  int
	day   int
	input string
	lines []string
	Cards map[int]*Card
}

func (puzzle *Puzzle) GetSummary() utils.Summary {
	s := utils.Summary{Day: puzzle.day, Year: puzzle.year, Name: puzzle.title}
	s.ProgressP1 = utils.Completed
	s.ProgressP2 = utils.Completed
	s.DateStarted = "2023-12-04 07:30:16"
	s.DateCompleted = "2023-12-04 10:41:00" /// I did actually work this morning :)
	return s
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("2023")
	iday, _ := strconv.Atoi("4")
	p := Puzzle{year: iyear, day: iday, title: "Day 4: Scratchcards"}
	p.Load(input)
	return &p
}

func NewPuzzle() *Puzzle {
	return NewPuzzleWithData(REAL_DATA)
}

func (puzzle *Puzzle) Load(input string) {
	lines := strings.Split(input, "\n")
	puzzle.input = input
	puzzle.lines = lines

	cards := make(map[int]*Card)
	for _, line := range lines {
		card := NewCard(line)
		cards[card.CardNumber] = card
	}
	puzzle.Cards = cards
}

func (puzzle *Puzzle) Part1() {
	puzzle.Load(REAL_DATA)

	p := NewPuzzleWithData(TEST_DATA)
	score := 0
	for _, line := range p.lines {
		card := NewCard(line)
		score += card.Score()
	}
	fmt.Printf("Part1: test score %v\n", score)

	p = NewPuzzleWithData(REAL_DATA)
	score = 0
	for _, line := range p.lines {
		card := NewCard(line)
		score += card.Score()
	}
	fmt.Printf("Part1: test score %v\n", score)

}

func (puzzle *Puzzle) Part2() {

	p := NewPuzzleWithData(TEST_DATA)

	// build a list of winners
	results := make([]*Card, 0)
	for _, card := range p.Cards {
		results = descend(card, p.Cards, results)
	}

	fmt.Printf("Part2: test data: there are %v won cards.\n", len(results))

	p = NewPuzzleWithData(REAL_DATA)

	// build a list of winners
	results = make([]*Card, 0)
	for _, card := range p.Cards {
		results = descend(card, p.Cards, results)
	}

	fmt.Printf("Part2: real data: there are %v won cards.\n", len(results))

}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}

func descend(card *Card, cards map[int]*Card, results []*Card) []*Card {
	// card is a winning card, so we will descend on it
	results = append(results, card)
	if len(card.WinningNumbers) == 0 {
		// fmt.Printf("descend: [%v] no winners, not descending.\n", card.CardNumber)
		return results
	}
	numberOfWinners := len(card.WinningNumbers)
	// fmt.Printf("descend: [%v] %v winners, will descend\n", card.CardNumber, len(card.WinningNumbers))
	for index := 0; index < numberOfWinners; index++ {
		realIndex := index + card.CardNumber + 1
		winningCard := cards[realIndex]
		results = descend(winningCard, cards, results)
	}
	return results
}
