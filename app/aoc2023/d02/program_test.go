package d02

import (
	"fmt"
	"testing"
)

func Test_1(t *testing.T) {
	p := NewPuzzleWithData(TEST_DATA)
	fmt.Printf("There are %v lines.\n", len(p.lines))
}

func Test_ParseGame(t *testing.T) {
	// g := NewGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	// fmt.Println(g.Debug())
	// if g.Red != 5 {
	// 	t.Fatalf("Expected 5 red, got %v\n", g.Red)
	// }
	// if g.Green != 4 {
	// 	t.Fatalf("Expected 4 green, got %v\n", g.Green)
	// }
	// if g.Blue != 9 {
	// 	t.Fatalf("Expected 9 red, got %v\n", g.Blue)
	// }
	// 	Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
	// 	Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
	// 	Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
	// 	Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green)
	// fmt.Printf("There are %v lines.\n", len(p.lines))
}

func Test_ParseGames(t *testing.T) {
	p := NewPuzzle()
	p.Load(TEST_DATA)
	if len(p.games) != 5 {
		t.Fatalf("Expected %v games, got %v\n", 5, len(p.games))
	}
}

func Test_FindGames(t *testing.T) {
	p := NewPuzzle()
	p.Load(TEST_DATA)
	games := p.FindGames(12, 13, 14)
	if len(games) != 3 {
		t.Fatalf("Expected %v games, got %v\n", 4, len(games))
	}
	for _, g := range games {
		fmt.Println(g.ID)
	}
}

func Test_Fewest(t *testing.T) {
	game := NewGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	r, g, b := game.Fewest()
	if r != 4 {
		t.Fatalf("Fewest shoudl be 4 red, was %v\n", r)
	}
	if g != 2 {
		t.Fatalf("Fewest shoudl be 2 green, was %v\n", g)
	}
	if b != 6 {
		t.Fatalf("Fewest shoudl be 6 blue, was %v\n", b)
	}
}

func Test_Power(t *testing.T) {
	game1 := NewGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	pgame1 := game1.Power()
	if pgame1 != 48 {
		t.Fatalf("Power shoudl be 48 , was %v\n", pgame1)
	}
}

func Test_PowerZ(t *testing.T) {
	puzzle := NewPuzzle()
	puzzle.Load(TEST_DATA)
	totalPower := 0
	for _, g := range puzzle.games {
		totalPower += g.Power()
	}
	if totalPower != 2286 {
		t.Fatalf("TotalPOwer shoudl be %v, was %v\n", 2286, totalPower)
	}
}
