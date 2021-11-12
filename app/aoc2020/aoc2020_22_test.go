package aoc2020

import (
	"testing"
)

func Test_AOC2020_22_1_Test(t *testing.T) {
	DATA := DAY_22_DATA
	// DATA := DAY_22_TEST_DATA
	combat := NewCombat(DATA)

	if combat.Player1.Size() != 5 {
		t.Errorf("Expected 4 cards in player 1 at the start, got %v.\n", combat.Player1.Size())
	}

	if combat.Player2.Size() != 5 {
		t.Errorf("Expected 4 cards in player 2 at the start, got %v.\n", combat.Player2.Size())
	}

	if combat.Player1.GetScore() != 78 {
		t.Errorf("Expected Score 78 at start for player1, got %v.\n", combat.Player1.GetScore())
	}

	combat.Play()
	expectedRound := 29
	actualRound := combat.GetRound()
	if expectedRound != actualRound {
		t.Errorf("Expected round to be %v, was %v.\n", expectedRound, actualRound)
	}

	expectedWinner := combat.Player2
	actualWinner := combat.GetWinner()
	if expectedWinner != actualWinner {
		t.Errorf("Expected winner to be Player 2.\n")
	}

	expectedScore := 306
	actualScore := combat.GetScore()
	if expectedScore != actualScore {
		t.Errorf("Expected score to be %v, was %v.\n", expectedScore, actualScore)
	}

}

func Test_AOC2020_22_2_Test(t *testing.T) {
	DATA := DAY_22_DATA
	// DATA := DAY_22_TEST_DATA
	combat := NewCombat(DATA)

	if combat.Player1.Size() != 5 {
		t.Errorf("Expected 4 cards in player 1 at the start, got %v.\n", combat.Player1.Size())
	}

	if combat.Player2.Size() != 5 {
		t.Errorf("Expected 4 cards in player 2 at the start, got %v.\n", combat.Player2.Size())
	}

	if combat.Player1.GetScore() != 78 {
		t.Errorf("Expected Score 78 at start for player1, got %v.\n", combat.Player1.GetScore())
	}

	combat.PlayDay2()
	// expectedRound := 29
	// actualRound := combat.GetRound()
	// if expectedRound != actualRound {
	// 	t.Errorf("Expected round to be %v, was %v.\n", expectedRound, actualRound)
	// }

	// expectedWinner := combat.Player2
	// actualWinner := combat.GetWinner()
	// if expectedWinner != actualWinner {
	// 	t.Errorf("Expected winner to be Player 2.\n")
	// }

	// expectedScore := 306
	// actualScore := combat.GetScore()
	// if expectedScore != actualScore {
	// 	t.Errorf("Expected score to be %v, was %v.\n", expectedScore, actualScore)
	// }

}
