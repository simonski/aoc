package main

import (
	"testing"
)

func Test_AOC2020_15_Test1(t *testing.T) {

	verifySequence("0,3,6", 436, t)
	verifySequence("2,1,3", 10, t)
	verifySequence("1,2,3", 27, t)
	verifySequence("2,3,1", 78, t)
	verifySequence("3,2,1", 438, t)
	verifySequence("3,1,2", 1836, t)

}

func verifySequence(sequence string, expected int, t *testing.T) {
	actual := NextInDay15Part1Sequence(sequence, 2020)
	if actual != expected {
		t.Errorf("In sequence %v, 2020th should be %v, was %v\n", sequence, expected, actual)
	}

}
