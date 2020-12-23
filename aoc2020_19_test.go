package main

import (
	"fmt"
	"strings"
	"testing"
)

const DAY_19_TEST_INPUT = `0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb
`

func Test_AOC2020_19_Test1(t *testing.T) {

	re := NewRegexRuleEngine(DAY_19_TEST_INPUT)
	if len(re.Rules) != 6 {
		t.Errorf("Day 19 Part 1: No. rules should be 6, was %v.\n", len(re.Rules))
	}

	if len(re.Messages) != 5 {
		t.Errorf("Day 19 Part 1: No. messages should be 5, was %v.\n", len(re.Rules))
	}

	if re.Rules["1"].Line != "1: 2 3 | 3 2" {
		t.Errorf("Day 19 Part 1: rule[1] should be `1: 2 3 | 3 2`, was '%v'\n", re.Rules["1"].Line)
	}

	if re.Messages[0] != "ababbb" {
		t.Errorf("Day 19 Part 1: messages[-1] should be `ababbb`, was '%v'\n", re.Messages[0])
	}

	re.ParseRules()
	total := re.Apply("0")
	if total != 2 {
		t.Errorf("Day 19 Part1: 2 messages should pass rules but %v did.", total)
	}

}

func Test_AOC2020_19_Test2(t *testing.T) {

	re := NewRegexRuleEngine(DAY_19_INPUT)
	re.ParseRules()
	rule := re.Rules["0"]
	regex := rule.Regex
	fmt.Printf("regex: '%v'\n", regex)
	regex = strings.ReplaceAll(regex, "(", "")
	regex = strings.ReplaceAll(regex, ")", "")
	regex = strings.ReplaceAll(regex, "a", "")
	regex = strings.ReplaceAll(regex, "b", "")
	regex = strings.ReplaceAll(regex, "|", "")
	fmt.Printf("regex: '%v'\n", regex)

	total := re.Apply("0")

	t.Errorf("Day 19 Part1: %v messages\n", total)

}
