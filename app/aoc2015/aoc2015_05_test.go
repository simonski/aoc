package aoc2015

import (
	"testing"
)

func Test_AOC2015_05_Part1(t *testing.T) {
	if !IsNice("ugknbfddgicrmopn") {
		t.Errorf("ugknbfddgicrmopn should be nice (vowel count is > 3, double count is 1)\n")
	}

	if !IsNice("aaa") {
		t.Errorf("aaa should be nice (vowel count is 3, double count is 1)\n")
	}

	if IsNice("jchzalrnumimnmhp") {
		t.Errorf("jchzalrnumimnmhp should be naughty\n")
	}

	if IsNice("haegwjzuvuyypxyu") {
		t.Errorf("haegwjzuvuyypxyu should be naughty\n")
	}

	if IsNice("dvszwmarrgswjxmb") {
		t.Errorf("dvszwmarrgswjxmb should be naughty\n")
	}

}

func Test_AOC2015_05_Part2(t *testing.T) {
	if !IsNice2("qjhvhtzxzqqjkmpb") {
		t.Errorf("qjhvhtzxzqqjkmpb should be nice.\n")
	}

	if !IsNice2("xxyxx") {
		t.Errorf("xxyxx should be nice.\n")
	}

	if IsNice2("uurcxstgmygtbstg") {
		t.Errorf("uurcxstgmygtbstg should be naughty\n")
	}

	if IsNice2("ieodomkazucvgmuy") {
		t.Errorf("ieodomkazucvgmuy should be naughty\n")
	}

}
