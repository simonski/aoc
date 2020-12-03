package main

import "testing"

func Test_AOC2020_03_TobogganMap_Get(t *testing.T) {

	line := ".....#....."
	tmap := NewTobogganMap()
	tmap.Add(line)

	verify(tmap, 0, 0, ".", t)
	verify(tmap, 5, 0, "#", t)
	verify(tmap, 6, 0, ".", t)
	verify(tmap, 11, 0, ".", t)
	verify(tmap, 16, 0, "#", t)

}

func Test_AOC2020_03_TobogganMap_IsTree(t *testing.T) {

	line := ".....#....."
	tmap := NewTobogganMap()
	tmap.Add(line)

	if tmap.IsTree(0, 0) {
		t.Errorf("(0,0) expected not IsTree")
	}

	if !tmap.IsTree(5, 0) {
		t.Errorf("(5,0) expected IsTree")
	}

	if tmap.IsTree(6, 0) {
		t.Errorf("(6,0) expected not IsTree")
	}

	if tmap.IsTree(11, 0) {
		t.Errorf("(11,0) expected not IsTree")
	}

	if !tmap.IsTree(16, 0) {
		t.Errorf("(16,0) expected IsTree")
	}

}

func verify(tmap *TobogganMap, col int, row int, expected string, t *testing.T) bool {
	actual := tmap.Get(col, row)
	if actual != expected {
		t.Errorf("(%v,%v) expected %v actual %v", row, col, expected, actual)
		return false
	}
	return true
}
