package d12

import (
	"fmt"
	"strings"

	"github.com/simonski/aoc/utils"
)

func countup(conditions string, rules []int) int {
	fmt.Printf("??  %v %V\n", conditions, rules)
	if len(rules) == 0 {
		if strings.Contains("#", conditions) {
			fmt.Printf("0  << %v %v\n", conditions, rules)
			return 0
		} else {
			fmt.Printf("1  << %v %v\n", conditions, rules)
			return 1
		}
	}
	if conditions == "" {
		if len(rules) == 0 {
			fmt.Printf("1  << %v %v\n", conditions, rules)
			return 1
		} else {
			fmt.Printf("0  << %v %v\n", conditions, rules)
			return 0
		}
	}

	result := 0

	letter := string(conditions[0:1])
	if letter == "." || letter == "?" {
		result += countup(conditions[1:], rules)
	}
	if letter == "#" || letter == "?" {
		var subcond string
		if len(conditions) <= rules[0] {
			subcond = ""
		} else {
			subcond = conditions[:rules[0]]
		}
		cond1 := rules[0] <= len(conditions)
		cond2 := utils.ContainsString(".", subcond)
		cond3 := false

		cond3_1 := len(rules) > 0
		cond3_2 := false
		if cond3_1 {
			cond3_2 = len(conditions) > rules[0]
		}
		if cond3_1 && cond3_2 {
			cond3 = (rules[0] == len(conditions) || string(conditions[rules[0]]) != "#")
		}

		if cond1 && cond2 && cond3 {
			subconditions := conditions[rules[0]+1:]
			result += countup(subconditions, rules[1:])

		}
	}

	return result

}
