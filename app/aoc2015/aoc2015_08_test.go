package app

import (
	"testing"
)

func Test_AOC2015_08_Part1(t *testing.T) {

	c1 := `""`
	c2 := `"abc"`
	c3 := `"aaa\"aaa"`
	c4 := `"\x27"`

	c1len := total_parsed_character_length(c1)
	if c1len != 0 {
		t.Errorf("character len on '%v' should be 0, was %v.", c1, c1len)
	}
	if len(c1) != 2 {
		t.Errorf("len on '%v' should be 2, was %v.", c1, len(c1))
	}

	c2len := total_parsed_character_length(c2)
	if c2len != 3 {
		t.Errorf("character len on '%v' should be 5, was %v.", c2, c2len)
	}
	if len(c2) != 5 {
		t.Errorf("len on '%v' should be 5, was %v.", c2, len(c2))
	}

	c3len := total_parsed_character_length(c3)
	if c3len != 7 {
		t.Errorf("character len on '%v' should be 7, was %v.", c3, c3len)
	}
	if len(c3) != 10 {
		t.Errorf("len on '%v' should be 10, was %v.", c3, len(c3))
	}

	c4len := total_parsed_character_length(c4)
	if c4len != 1 {
		t.Errorf("character len on '%v' should be 1, was %v.", c4, c4len)
	}
	if len(c4) != 6 {
		t.Errorf("len on '%v' should be 6, was %v.", c4, len(c4))
	}

}
