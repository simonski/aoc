package d7

import (
	"fmt"
	"strings"
	"testing"
)

// func Test_1(t *testing.T) {
// 	p := NewPuzzleWithData(TEST_DATA)
// 	fmt.Printf("There are %v lines.\n", len(p.lines))
// }

// func Test_P1_OnePair(t *testing.T) {
// 	h := NewHandP1("32T3K 765")
// 	if !h.IsOnePair {
// 		t.Fatalf("Expect one pair, got: %v", h.Debug())
// 	}

// 	h2 := NewHandP1("T55J5 684")
// 	if !h2.IsThreeOfAKind {
// 		t.Fatalf("Expect three of a kind, was: %v", h2.Debug())
// 	}

// 	h3 := NewHandP1("KK677 28")
// 	if !h3.IsTwoPair {
// 		t.Fatalf("Expect twopair, was: %v", h3.Debug())
// 	}
// }

// //
// // KK677 28
// // KTJJT 220
// // QQQJA 483)

// func Test_P1_Game(t *testing.T) {
// 	g := NewGameP1(strings.Split(TEST_DATA, "\n"))
// 	g.Sort()
// 	total := 0
// 	for _, h := range g.Hands {
// 		total += (h.Bid * h.Rank)
// 	}
// 	if total != 6441 {

// 		for _, h := range g.Hands {
// 			fmt.Printf("Rank: %v, Bid %v, Cards: %v, debug=%v\n", h.Rank, h.Bid, h.Line, h.Debug())
// 		}

// 		t.Fatalf("Expected 6440, was %v\n", total)
// 	}

// }

// func Test_P1_FH(t *testing.T) {

// 	data := "33332 100\n2AAAA 323"
// 	g := NewGameP1(strings.Split(data, "\n"))
// 	g.Sort()
// 	total := 0
// 	for _, h := range g.Hands {
// 		total += (h.Bid * h.Rank)
// 	}
// 	if total != 6440 {
// 		for _, h := range g.Hands {
// 			fmt.Printf("Rank: %v, Bid %v, Cards: %v, debug=%v\n", h.Rank, h.Bid, h.Line, h.Debug())
// 		}
// 		// 247422370 too low
// 		// 247422370
// 		t.Fatalf("Expected 6440, was %v\n", total)
// 	}
// }

// func Test_P1_FH2(t *testing.T) {

// 	data := "77888 100\n77788 323"
// 	g := NewGameP1(strings.Split(data, "\n"))
// 	g.Sort()
// 	total := 0
// 	for _, h := range g.Hands {
// 		total += (h.Bid * h.Rank)
// 	}
// 	if total != 6440 {
// 		for _, h := range g.Hands {
// 			fmt.Printf("Rank: %v, Bid %v, Cards: %v, debug=%v\n", h.Rank, h.Bid, h.Line, h.Debug())
// 		}
// 		t.Fatalf("Expected 6440, was %v\n", total)
// 	}
// }

// func Test_P1_GameReal(t *testing.T) {
// 	g := NewGameP1(strings.Split(REAL_DATA, "\n"))
// 	g.Sort()
// 	total := 0
// 	for _, h := range g.Hands {
// 		total += (h.Bid * h.Rank)
// 		// fmt.Printf("total %v\n", total)
// 	}
// 	if total != 6440 {
// 		// for _, h := range g.Hands {
// 		// if h.IsOnePair {
// 		// 	fmt.Printf("%v, %v\n", h.Rank, h.Line)
// 		// }
// 		// }
// 		// 247422370 too low
// 		// 247422370
// 		t.Fatalf("Expected 6440, was %v\n", total)
// 	}
// }

func Test_P2_GameTest(t *testing.T) {
	g := NewGameP2(strings.Split(TEST_DATA, "\n"))
	g.Sort()
	total := 0
	for _, h := range g.Hands {
		// fmt.Printf("%v %v\n", h.Rank, h.Debug())
		total += (h.Bid * h.Rank)
	}
	if total != 6441 {
		for _, h := range g.Hands {
			fmt.Printf("Rank: %v, Bid %v, Cards: %v, debug=%v\n", h.Rank, h.Bid, h.Line, h.Debug())
		}
		t.Fatalf("Expected 6440, was %v\n", total)
	}

}

func Test_P2_GameReal(t *testing.T) {
	g := NewGameP2(strings.Split(REAL_DATA, "\n"))
	g.Sort()
	total := 0
	// 248784657 too high
	// 248540703 // no
	// 248747492
	// 248454024 too low
	for _, h := range g.Hands {
		// fmt.Printf("%v %v\n", h.Rank, h.Debug())
		total += (h.Bid * h.Rank)
	}
	if total != 6441 {
		for _, h := range g.Hands {
			fmt.Printf("Rank: %v, Bid %v, Cards: %v, debug=%v\n", h.Rank, h.Bid, h.Line, h.Debug())
		}
		t.Fatalf("Expected 6440, was %v\n", total)
	}
}
