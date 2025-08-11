package d7

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Game struct {
	Hands []*Hand
}

func NewGameP1(input []string) *Game {
	hands := make([]*Hand, 0)
	for _, line := range input {
		h := NewHandP1(line)
		hands = append(hands, h)
	}
	g := Game{}
	g.Hands = hands
	return &g
}

func NewGameP2(input []string) *Game {
	hands := make([]*Hand, 0)
	for _, line := range input {
		h := NewHandP2(line)
		hands = append(hands, h)
	}
	g := Game{}
	g.Hands = hands
	return &g
}

func (g *Game) Sort() {
	sort.Slice(g.Hands, func(i, j int) bool {
		h1 := g.Hands[i]
		h2 := g.Hands[j]
		return Winner(h1, h2) == h1
	})
	for index, h := range g.Hands {
		h.Rank = index + 1
	}
}

type Hand struct {
	Line           string
	Cards          []*Card
	Score          int
	Rank           int
	Bid            int
	IsFiveOfAKind  bool
	IsFourOfAKind  bool
	IsThreeOfAKind bool
	IsTwoPair      bool
	IsOnePair      bool
	IsFullHouse    bool
	IsHighCard     bool
}

func (h *Hand) Debug() string {
	cardsStr := ""
	for _, c := range h.Cards {
		cardsStr = fmt.Sprintf("%v%v", cardsStr, c.Name)
	}

	return fmt.Sprintf("Line='%v', Cards='%v', Bid='%v', Is5=%v, is4=%v, Is3=%v, is2P=%v, Is1P=%v, isFH=%v, isHigh=%v",
		h.Line, cardsStr, h.Bid, h.IsFiveOfAKind, h.IsFourOfAKind, h.IsThreeOfAKind, h.IsTwoPair, h.IsOnePair, h.IsFullHouse, h.IsHighCard)
}

func NewHandP1(line string) *Hand {

	splits := strings.Split(line, " ")
	cardsS := splits[0]
	bidS := splits[1]
	bid, _ := strconv.Atoi(bidS)

	cards := make([]*Card, 0)
	cardMap := BuildCardMapP1()

	cardCounter := make(map[string]int)
	for letter, _ := range cardMap {
		cardCounter[letter] = 0
	}

	for index := 0; index < len(cardsS); index++ {
		letter := cardsS[index : index+1]
		card := cardMap[letter]
		cards = append(cards, card)
		count := cardCounter[card.Name]
		count += 1
		cardCounter[card.Name] = count
	}

	hand := Hand{Line: line, Cards: cards, Bid: bid}

	walk := func(cardCounter map[string]int) (int, int, int, int) {
		fives := 0
		fours := 0
		threes := 0
		pairs := 0
		for _, value := range cardCounter {
			if value == 5 {
				fives += 1
			} else if value == 4 {
				fours += 1
			} else if value == 3 {
				threes += 1
			} else if value == 2 {
				pairs += 1
			}
		}
		return fives, fours, threes, pairs

	}

	five, four, three, pairs := walk(cardCounter)
	if five == 1 {
		hand.IsFiveOfAKind = true
		hand.Score = 6
	} else if four == 1 {
		hand.IsFourOfAKind = true
		hand.Score = 5
	} else if three == 1 && pairs == 1 {
		hand.IsFullHouse = true
		hand.Score = 4
	} else if three == 1 && pairs == 0 {
		hand.IsThreeOfAKind = true
		hand.Score = 3
	} else if pairs == 2 {
		hand.IsTwoPair = true
		hand.Score = 2
	} else if pairs == 1 && three == 0 {
		hand.IsOnePair = true
		hand.Score = 1
	} else {
		hand.Score = 0
		hand.IsHighCard = true
	}

	return &hand

}

func NewHandP2(line string) *Hand {

	splits := strings.Split(line, " ")
	cardsS := splits[0]
	bidS := splits[1]
	bid, _ := strconv.Atoi(bidS)

	cards := make([]*Card, 0)
	cardMap := BuildCardMapP2()

	cardCounter := make(map[string]int)
	for letter, _ := range cardMap {
		cardCounter[letter] = 0
	}

	for index := 0; index < len(cardsS); index++ {
		letter := cardsS[index : index+1]
		card := cardMap[letter]
		cards = append(cards, card)
		count := cardCounter[card.Name]
		count += 1
		cardCounter[card.Name] = count
	}

	hand := Hand{Line: line, Cards: cards, Bid: bid}

	walk := func(cardCounter map[string]int) (int, int, int, int) {
		fives := 0
		fours := 0
		threes := 0
		pairs := 0
		for key, value := range cardCounter {
			if key == "J" {
				continue
			}
			if value == 5 {
				fives += 1
			} else if value == 4 {
				fours += 1
			} else if value == 3 {
				threes += 1
			} else if value == 2 {
				pairs += 1
			}
		}
		return fives, fours, threes, pairs

	}

	jCount := cardCounter["J"]
	five, four, three, pairs := walk(cardCounter)
	if five == 1 || (four == 1 && jCount == 1) || (three == 1 && jCount == 2) || (pairs == 1 && jCount == 3) || jCount == 5 {
		hand.IsFiveOfAKind = true
		hand.Score = 6
	} else if four == 1 || (three == 1 && jCount == 1) || (pairs == 1 && jCount == 2) {
		hand.IsFourOfAKind = true
		hand.Score = 5
	} else if (three == 1 && pairs == 1) || (pairs == 1 && jCount == 3) || (pairs == 2 && jCount == 1) {
		hand.IsFullHouse = true
		hand.Score = 4
	} else if (three == 1 && pairs == 0) || (pairs == 1 && jCount == 1) {
		hand.IsThreeOfAKind = true
		hand.Score = 3
	} else if pairs == 2 || (pairs == 1 && jCount == 1) {
		hand.IsTwoPair = true
		hand.Score = 2
	} else if pairs == 1 && three == 0 || (pairs == 0 && jCount == 1) {
		hand.IsOnePair = true
		hand.Score = 1
	} else if jCount == 5 || jCount == 4 {
		hand.IsFiveOfAKind = true
		hand.Score = 6
	} else if jCount == 3 {
		hand.IsFourOfAKind = true
		hand.Score = 5
	} else if jCount == 2 {
		hand.IsThreeOfAKind = true
		hand.Score = 3
	} else if jCount == 1 {
		hand.IsOnePair = true
		hand.Score = 1
	} else {
		hand.Score = 0
		hand.IsHighCard = true
	}

	return &hand

}

func Winner(left *Hand, right *Hand) *Hand {
	if left.Score > right.Score {
		return right
	} else if right.Score > left.Score {
		return left
	} else {
		return CompareHighest(left, right)
	}
}

func CompareHighest(left *Hand, right *Hand) *Hand {
	for i := 0; i < 5; i++ {
		c1 := left.Cards[i]
		c2 := right.Cards[i]
		if c1.Value < c2.Value {
			// fmt.Printf("CompareHighest((%v/%v) [%v] %v <> %v ? return left - %v\n\n", left.Line, right.Line, i, c1.Name, c2.Name, c1.Name)
			return left
		} else if c1.Value > c2.Value {
			// fmt.Printf("CompareHighest((%v/%v) [%v] %v <> %v ? return right - %v\n\n", left.Line, right.Line, i, c1.Name, c2.Name, c2.Name)
			return right
		} else {
			// fmt.Printf("CompareHighest((%v/%v) [%v] %v <> %v ? SAME, continue\n", left.Line, right.Line, i, c1.Name, c2.Name)
		}
	}
	fmt.Println("<<< NOTHING?? RETURNING RIGHT>>>>")
	return nil
}

type Card struct {
	Name  string
	Value int
}

func NewCard(name string, value int) *Card {
	return &Card{Name: name, Value: value}
}

func BuildCardMapP1() map[string]*Card {
	cm := make(map[string]*Card)
	cm["A"] = NewCard("A", 14)
	cm["K"] = NewCard("K", 13)
	cm["Q"] = NewCard("Q", 12)
	cm["J"] = NewCard("J", 11)
	cm["T"] = NewCard("T", 10)
	cm["9"] = NewCard("9", 9)
	cm["8"] = NewCard("8", 8)
	cm["7"] = NewCard("7", 7)
	cm["6"] = NewCard("6", 6)
	cm["5"] = NewCard("5", 5)
	cm["4"] = NewCard("4", 4)
	cm["3"] = NewCard("3", 3)
	cm["2"] = NewCard("2", 2)
	return cm
}

func BuildCardMapP2() map[string]*Card {
	cm := make(map[string]*Card)
	cm["A"] = NewCard("A", 14)
	cm["K"] = NewCard("K", 13)
	cm["Q"] = NewCard("Q", 12)
	cm["T"] = NewCard("T", 10)
	cm["9"] = NewCard("9", 9)
	cm["8"] = NewCard("8", 8)
	cm["7"] = NewCard("7", 7)
	cm["6"] = NewCard("6", 6)
	cm["5"] = NewCard("5", 5)
	cm["4"] = NewCard("4", 4)
	cm["3"] = NewCard("3", 3)
	cm["2"] = NewCard("2", 2)
	cm["J"] = NewCard("J", 1)
	return cm
}
