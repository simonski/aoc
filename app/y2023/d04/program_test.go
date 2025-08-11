package d4

import (
	"fmt"
	"testing"
)

func Test_1(t *testing.T) {
	p := NewPuzzleWithData(REAL_DATA)
	fmt.Printf("There are %v lines.\n", len(p.lines))

	score := 0
	for _, line := range p.lines {
		card := NewCard(line)
		fmt.Println(card.Debug())
		score += card.Score()
	}

	fmt.Printf("Score: %v\n", score)
	t.Fatal("Nope")

}

func Test_2(t *testing.T) {
	p := NewPuzzleWithData(REAL_DATA)
	fmt.Printf("There are %v lines.\n", len(p.lines))

	score := 0
	for _, line := range p.lines {
		card := NewCard(line)
		fmt.Println(card.Debug())
		score += card.Score()
	}

	fmt.Printf("Score: %v\n", score)
	t.Fatal("Nope")

}
