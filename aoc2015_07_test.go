package main

import (
	"testing"
)

func Test_AOC2015_06_Part1(t *testing.T) {
	grid := NewLightGrid()
	actualOn, actualOff := grid.CountOnOff()
	if actualOn != 0 {
		t.Errorf("LightGrid should start with 0 on, was %v\n", actualOn)
	}
	if actualOff != 1000000 {
		t.Errorf("LightGrid should start with 1,000,000 off, was %v\n", actualOff)
	}

	grid.Execute("turn on 0,0 through 999,999")
	actualOn, actualOff = grid.CountOnOff()
	if actualOn != 1000000 {
		t.Errorf("LightGrid 'turn on 0,0 through 999,999' should turn on 1,000,000, was %v\n", actualOn)
	}
	if actualOff != 0 {
		t.Errorf("LightGrid 'turn off 0,0 through 999,999' should turn all off to 0 was %v\n", actualOff)
	}

	grid.Execute("turn off 0,0 through 999,999")
	actualOn, actualOff = grid.CountOnOff()
	if actualOff != 1000000 {
		t.Errorf("XXXX LightGrid 'turn off 0,0 through 999,999' should turn off 1,000,000, was %v\n", actualOff)
	}

	grid.Execute("toggle 499,499 through 500,500")
	actualOn, actualOff = grid.CountOnOff()
	if actualOn != 4 {
		t.Errorf("XXXX LightGrid 'toggle 499,499 through 500,500' should turn on 4, was %v\n", actualOn)
	}

}

