package main

import (
	"testing"
)

func Test_AOC2020_13_Test1(t *testing.T) {

	time := 939
	buses := []int{7, 13, 59, 31, 19}
	earliest, bus := Day13Logic(time, buses)
	if bus != 59 {
		t.Errorf("Earliest bus should be %v, was %v\n", 59, bus)
	}
	if earliest != 944 {
		t.Errorf("Earliest time should be %v, was %v\n", 944, earliest)
	}

	waitTime := earliest - time
	if waitTime != 5 {
		t.Errorf("Time should be 5, was %v\n", waitTime)
	}

}
