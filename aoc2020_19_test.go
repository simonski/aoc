package main

import (
	"fmt"
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

// 8: 42 | 42 8
// 11: 42 31 | 42 11 31

// const DAY_19_INPUT_PART_2_EXAMPLE_ALTERED = `42: 9 14 | 10 1
// 9: 14 27 | 1 26
// 10: 23 14 | 28 1
// 1: "a"
// 11: 42 31
// 5: 1 14 | 15 1
// 19: 14 1 | 14 14
// 12: 24 14 | 19 1
// 16: 15 1 | 14 14
// 31: 14 17 | 1 13
// 6: 14 14 | 1 14
// 2: 1 24 | 14 4
// 0: 8 11
// 13: 14 3 | 1 12
// 15: 1 | 14
// 17: 14 2 | 1 7
// 23: 25 1 | 22 14
// 28: 16 1
// 4: 1 1
// 20: 14 14 | 1 15
// 3: 5 14 | 16 1
// 27: 1 6 | 14 18
// 14: "b"
// 21: 14 1 | 1 14
// 25: 1 1 | 1 14
// 22: 14 14
// 8: 42
// 26: 14 22 | 1 20
// 18: 15 15
// 7: 14 5 | 1 21
// 24: 14 1

// abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa
// bbabbbbaabaabba
// babbbbaabbbbbabbbbbbaabaaabaaa
// aaabbbbbbaaaabaababaabababbabaaabbababababaaa
// bbbbbbbaaaabbbbaaabbabaaa
// bbbababbbbaaaaaaaabbababaaababaabab
// ababaaaaaabaaab
// ababaaaaabbbaba
// baabbaaaabbaaaababbaababb
// abbbbabbbbaaaababbbbbbaaaababb
// aaaaabbaabaaaaababaa
// aaaabbaaaabbaaa
// aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
// babaaabbbaaabaababbaabababaaab
// aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba`

/*
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

	t.Errorf("Day 19 Part1 Test 2 (Not a test, just running): %v messages\n", total)

}

const DAY_19_INPUT_PART_2_EXAMPLE = `42: 9 14 | 10 1
9: 14 27 | 1 26
10: 23 14 | 28 1
1: "a"
11: 42 31
5: 1 14 | 15 1
19: 14 1 | 14 14
12: 24 14 | 19 1
16: 15 1 | 14 14
31: 14 17 | 1 13
6: 14 14 | 1 14
2: 1 24 | 14 4
0: 8 11
13: 14 3 | 1 12
15: 1 | 14
17: 14 2 | 1 7
23: 25 1 | 22 14
28: 16 1
4: 1 1
20: 14 14 | 1 15
3: 5 14 | 16 1
27: 1 6 | 14 18
14: "b"
21: 14 1 | 1 14
25: 1 1 | 1 14
22: 14 14
8: 42
26: 14 22 | 1 20
18: 15 15
7: 14 5 | 1 21
24: 14 1

abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa
bbabbbbaabaabba
babbbbaabbbbbabbbbbbaabaaabaaa
aaabbbbbbaaaabaababaabababbabaaabbababababaaa
bbbbbbbaaaabbbbaaabbabaaa
bbbababbbbaaaaaaaabbababaaababaabab
ababaaaaaabaaab
ababaaaaabbbaba
baabbaaaabbaaaababbaababb
abbbbabbbbaaaababbbbbbaaaababb
aaaaabbaabaaaaababaa
aaaabbaaaabbaaa
aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
babaaabbbaaabaababbaabababaaab
aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba`

func Test_AOC2020_19_Part2_Test3(t *testing.T) {

	re := NewRegexRuleEngine(DAY_19_INPUT_PART_2_EXAMPLE)
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

	if total != 3 {
		t.Errorf("Day 19 Part2: %v messages\n", total)
	}

}
func Test_AOC2020_19_Test4(t *testing.T) {

	re := NewRegexRuleEngine(DAY_19_INPUT_PART_2_EXAMPLE_ALTERED)
	re.ParseRulesV2()
	message := "bbabbbbaabaabba"
	fmt.Printf("%v\n", message)
	for _, rule := range re.Rules {
		if rule.IsMessageValid(message) {
			fmt.Printf("PASSED (%v) '%v' : %v\n", message, rule.Line, rule.Value)
		} else {
			fmt.Printf("FAILED (%v) '%v' : %v\n", message, rule.Line, rule.Value)
		}
	}

}
*/

// func Test_AOC2020_19_Test4(t *testing.T) {

// 	re := NewRegexRuleEngine(DAY_19_INPUT_PART_2_EXAMPLE_ALTERED)
// 	re.ParseRulesV2()
// 	message := "bbabbbbaabaabba"
// 	fmt.Printf("%v\n", message)
// 	for _, rule := range re.Rules {
// 		if rule.IsMessageValid(message) {
// 			fmt.Printf("PASSED (%v) '%v' : %v\n", message, rule.Line, rule.Value)
// 		} else {
// 			// fmt.Printf("FAILED (%v) '%v' : %v\n", message, rule.Line, rule.Value)
// 		}
// 	}

// }

func Test_Day19_Part2_ParseRules(t *testing.T) {
	re := NewRegexRuleEngine(DAY_19_INPUT_PART_2)
	re.ParseRulesV2()
	rule := re.Rules["0"]
	total := 0

	// 	test_messages := `abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa
	// bbabbbbaabaabba
	// babbbbaabbbbbabbbbbbaabaaabaaa
	// aaabbbbbbaaaabaababaabababbabaaabbababababaaa
	// bbbbbbbaaaabbbbaaabbabaaa
	// bbbababbbbaaaaaaaabbababaaababaabab
	// ababaaaaaabaaab
	// ababaaaaabbbaba
	// baabbaaaabbaaaababbaababb
	// abbbbabbbbaaaababbbbbbaaaababb
	// aaaaabbaabaaaaababaa
	// aaaabbaaaabbaaa
	// aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
	// babaaabbbaaabaababbaabababaaab
	// aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba`

	for _, message := range re.Messages {
		if rule.IsMessageValid(message) {
			fmt.Printf("PASSED (%v) '%v' : %v\n", message, rule.Line, rule.Value)
			total++
		} else {
			fmt.Printf("FAILED (%v) '%v' : %v\n", message, rule.Line, rule.Value)
		}
	}
	// total := re.Apply("0")
	fmt.Printf("Part2 rules gives %v rules passing.\n", total)
}
