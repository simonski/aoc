package aoc2020

/*

--- Day 19: Monster Messages ---
You land in an airport surrounded by dense forest. As you walk to your high-speed train, the Elves at the Mythical Information Bureau contact you again. They think their satellite has collected an image of a sea monster! Unfortunately, the connection to the satellite is having problems, and many of the messages sent back from the satellite have been corrupted.

They sent you a list of the rules valid messages should obey and a list of received messages they've collected so far (your puzzle input).

The rules for valid messages (the top part of your puzzle input) are numbered and build upon each other. For example:

0: 1 2
1: "a"
2: 1 3 | 3 1
3: "b"
Some rules, like 3: "b", simply match a single character (in this case, b).

The remaining rules list the sub-rules that must be followed; for example, the rule 0: 1 2 means that to match rule 0, the text being checked must match rule 1, and the text after the part that matched rule 1 must then match rule 2.

Some of the rules have multiple lists of sub-rules separated by a pipe (|). This means that at least one list of sub-rules must match. (The ones that match might be different each time the rule is encountered.) For example, the rule 2: 1 3 | 3 1 means that to match rule 2, the text being checked must match rule 1 followed by rule 3 or it must match rule 3 followed by rule 1.

Fortunately, there are no loops in the rules, so the list of possible matches will be finite. Since rule 1 matches a and rule 3 matches b, rule 2 matches either ab or ba. Therefore, rule 0 matches aab or aba.

Here's a more interesting example:

0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"
Here, because rule 4 matches a and rule 5 matches b, rule 2 matches two letters that are the same (aa or bb), and rule 3 matches two letters that are different (ab or ba).

Since rule 1 matches rules 2 and 3 once each in either order, it must match two pairs of letters, one pair with matching letters and one pair with different letters. This leaves eight possibilities: aaab, aaba, bbab, bbba, abaa, abbb, baaa, or babb.

Rule 0, therefore, matches a (rule 4), then any of the eight options from rule 1, then b (rule 5): aaaabb, aaabab, abbabb, abbbab, aabaab, aabbbb, abaaab, or ababbb.

The received messages (the bottom part of your puzzle input) need to be checked against the rules so you can determine which are valid and which are corrupted. Including the rules and the messages together, this might look like:

0: 4 1 5
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
Your goal is to determine the number of messages that completely match rule 0. In the above example, ababbb and abbbab match, but bababa, aaabbb, and aaaabbb do not, producing the answer 2. The whole message must match all of rule 0; there can't be extra unmatched characters in the message. (For example, aaaabbb might appear to match rule 0 above, but it has an extra unmatched b on the end.)

How many messages completely match rule 0?

*/

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	cli "github.com/simonski/cli"
	goutils "github.com/simonski/goutils"
)

func (app *Application) Y2020D19(cli *cli.CLI) {
	app.Y2020D19P1(cli)
	app.Y2020D19P2(cli)
}

func (app *Application) Y2020D19P1(cli *cli.CLI) {

	/*
		OK I’m guessing loop detection; a cheap way of doing that would be.. find what gives a loop then exclude any that include that in their rule to begin with
		10:17
		I non-nice way would be parse with a depth counter and just break after some depth charge goes off
		10:18
		that would mark that one as exploded and (probably) not good, then the remainder are the possibly valid messages, then evaluate only them

		that gets trickey as really what you want to do is assume some messages would pass a rule which had an OR where one was a loop, so you’d retain that message and evaluate only the non-loopy side auto-marking the loopy side as false
		10:28
		ok I have real life but some sort of depth charge as a loop detection followed by marking that rule as bad/no-eval and then any loop marked as false is my first attempt… later… real life time now
	*/

	rre := NewRegexRuleEngine(DAY_19_INPUT_PART_1)
	rr8 := NewRegexRule("8: 42 | 42 8")
	rr11 := NewRegexRule("11: 42 31 | 42 11 31")
	rre.Rules["8"] = rr8
	rre.Rules["11"] = rr11

	rre.Debug()
}

func (app *Application) Y2020D19P2(cli *cli.CLI) {

	/*
		I think that this
		rr8 := NewRegexRule("8: 42 | 42 8")			    >> 42 | (42+)   == 42+
		rr11 := NewRegexRule("11: 42 31 | 42 11 31")    >> 42 11 31  >>   42 (42 31) 31  >>

		step0: 42 31 | 42 11 31
		step1: 42 31 | 42 (42 31) 31
		step2: 42 31 | 42 (42 (42 31) 31) 31
		step3: 42 31 | 42 (42 (42 (42 31) 31) 31) 31
		step4: 42 31 | 42 (42 (42 (42 (42 31) 31) 31) 31) 31
		step5: 42 31 | 42 (42 (42 (42 (42 (42 31) 31) 31) 31) 31) 31
		step6: 42 31 | 42 (42 (42 (42 (42 (42 (42 31) 31) 31) 31) 31) 31) 31
		step7: 42 31 | 42 (42 (42 (42 (42 (42 (42 (42 31) 31) 31) 31) 31) 31) 31) 31
		step8: 42 31 | 42 (42 (42 (42 (42 (42 (42 (42 (42 31) 31) 31) 31) 31) 31) 31) 31) 31
		step9: 42 31 | 42 (42 (42 (42 (42 (42 (42 (42 (42 (42 31) 31) 31) 31) 31) 31) 31) 31) 31) 31
		step10: 42 31 | 42 (42 (42 (42 (42 (42 (42 (42 (42 (42 (42 31) 31) 31) 31) 31) 31) 31) 31) 31) 31) 31
		step11: 42 31 | 42 (42 (42 (42 (42 (42 (42 (42 (42 (42 (42 (42 31) 31) 31) 31) 31) 31) 31) 31) 31) 31) 31) 31

		step2: 42 31 | (42 (42 31) 31) | (42 (42 (42 31) 31) 31) | (42 (42 (42 (42 31) 31) 31) 31) | (42 (42 (42 (42 (42 31) 31) 31) 31) 31) (42 (42 (42 (42 (42 (42 31) 31) 31) 31) 31) 31) | (42 (42 (42 (42 (42 (42 (42 31) 31) 31) 31) 31) 31) 31) | (42 (42 (42 (42 (42 (42 (42 (42 31) 31) 31) 31) 31) 31) 31) 31) | (42 (42 (42 (42 (42 (42 (42 (42 (42 31) 31) 31) 31) 31) 31) 31) 31) 31)


		rule = (step1) | (step2) | (step3) ....

		is this
		rr8 := NewRegexRule("8: 42 | (42)+")
		rr11 := NewRegexRule("11: 42 31 | 42 (42 31)+ 31")

		so I can just manually do this
		x1
		rr8 := NewRegexRule("8: 42 | 42 (42)")
		rr11 := NewRegexRule("11: 42 31 | 42 (42 31) 31")

		x2
		rr8 := NewRegexRule("8: 42 | 42 42 42")
		WRONG rr11 := NewRegexRule("11: 42 31 | 42 (42 31) (42 31) 31")

		RIGHT rr11 := NewRegexRule("11: 42 31 | (42 (42 31) 31)")


		xN
		which means I have a 'source' regex that I keep adding to
		rr8 := NewRegexRule("8: 42 | 42 RR8")
		rr11 := NewRegexRule("11: 42 31 | 42 RR11 31")

		Where I replace
		RR8 with N x 42
		and
		RR11 with N x 42 31
		add in the regex each time and see which messages pass
		keep going until no messages pass
		keep adding the passing messages to a set
		at the end, the set size is the number of messages that pass in total
	*/
	index := 1
	rr8_source := "8: RR8"
	rr8_template := " 42"
	rr8_replacewith := ""

	// rr11_source := "11: 42 31 | RR42 RR11 RR31"
	// rr11_template := " 42 31 "
	// rr11_template42 := " 42 "
	// rr11_template31 := " 31 "

	// rr11_replacewith := ""
	// r11 := "11: 42 31 | ((42 (42 31) 31) | (42 (42 (42 31) 31) 31) | (42 (42 (42 (42 31) 31) 31) 31) | (42 (42 (42 (42 (42 31) 31) 31) 31) 31) (42 (42 (42 (42 (42 (42 31) 31) 31) 31) 31) 31) | (42 (42 (42 (42 (42 (42 (42 31) 31) 31) 31) 31) 31) 31) | (42 (42 (42 (42 (42 (42 (42 (42 31) 31) 31) 31) 31) 31) 31) 31) | (42 (42 (42 (42 (42 (42 (42 (42 (42 31) 31) 31) 31) 31) 31) 31) 31) 31))"
	r11 := "11: 42 31 | ( (42 (42 31) 31) | (42 (42 (42 31) 31) 31) | (42 (42 (42 (42 31) 31) 31) 31) | (42 (42 (42 (42 (42 31) 31) 31) 31) 31) | (42 (42 (42 (42 (42 (42 31) 31) 31) 31) 31) 31) | (42 (42 (42 (42 (42 (42 (42 31) 31) 31) 31) 31) 31) 31) | (42 (42 (42 (42 (42 (42 (42 (42 31) 31) 31) 31) 31) 31) 31) 31) | (42 (42 (42 (42 (42 (42 (42 (42 (42 31) 31) 31) 31) 31) 31) 31) 31) 31) | (42 (42 (42 (42 (42 (42 (42 (42 (42 (42 31) 31) 31) 31) 31) 31) 31) 31) 31) 31) | (42 (42 (42 (42 (42 (42 (42 (42 (42 (42 (42 31) 31) 31) 31) 31) 31) 31) 31) 31) 31) 31) | (42 (42 (42 (42 (42 (42 (42 (42 (42 (42 (42 (42 31) 31) 31) 31) 31) 31) 31) 31) 31) 31) 31) 31) )"
	// passedMessages := make(map[string]int)
	// overallTotal := 0

	// here I will keep all "previouly passed" messages and then
	// on each iteration I will remove them so I don't query them
	allPassingMessages := make([]string, 0)

	// not 307, 331? NO
	// 374

	for {
		rr8_replacewith += rr8_template
		// rr11_replacewith += rr11_template
		rr8_regex := strings.ReplaceAll(rr8_source, "RR8", rr8_replacewith)
		// rr11_regex := strings.ReplaceAll(rr11_source, "RR11", rr11_replacewith)
		rr8_regex = strings.ReplaceAll(rr8_regex, "  ", " ")
		// rr11_regex = strings.ReplaceAll(rr11_regex, "  ", " ")

		// fmt.Printf("%v\n", rr8_regex)
		// fmt.Printf("%v\n", r11)

		rre := NewRegexRuleEngine(DAY_19_INPUT_PART_1)
		rre.RemoveMessages(allPassingMessages)
		fmt.Printf("[%v] Removed %v messages that pass, now have %v messages to check.\n", index, len(allPassingMessages), len(rre.Messages))
		rr8 := NewRegexRule(rr8_regex)
		// rr11 := NewRegexRule(rr11_regex)
		rr11 := NewRegexRule(r11)
		rre.Rules["8"] = rr8
		rre.Rules["11"] = rr11

		// fmt.Printf("R8: %v\n", rr8_regex)
		// fmt.Printf("R11: %v\n", rr11_regex)
		rre.ParseRules(false)

		r42 := rre.Rules["42"]
		r31 := rre.Rules["31"]

		r0 := rre.Rules["0"]
		r0.Regex = strings.ReplaceAll(r0.Regex, "4231", "4231")
		r0.Regex = strings.ReplaceAll(r0.Regex, "42", r42.Regex)
		r0.Regex = strings.ReplaceAll(r0.Regex, "31", r31.Regex)

		r11 := rre.Rules["11"]
		r11.Regex = strings.ReplaceAll(r11.Regex, "4231", "4231")
		r11.Regex = strings.ReplaceAll(r11.Regex, "42", r42.Regex)
		r11.Regex = strings.ReplaceAll(r11.Regex, "31", r31.Regex)

		// r0.Regex = strings.ReplaceAll(r0.Regex, "4231", "42 31")
		// for _, rule := range rre.Rules {
		// 	fmt.Printf("%v %v\n", rule.Key, rule.Regex)
		// }

		// os.Exit(1)
		total, passingMessages := rre.Apply("0", false)
		fmt.Printf("%v messages pass.\n", total)
		for _, message := range passingMessages {
			allPassingMessages = append(allPassingMessages, message)
		}
		fmt.Printf("[%v] %v rules pass this time, so far %v has passed in total.\n", index, len(passingMessages), len(allPassingMessages))
		// thisLength := len(passedRules)
		// if overallTotal == thisLength {
		// 	os.Exit(1)
		// } else {
		// overallTotal = thisLength
		// }

		// fmt.Printf("[%v] RR8  %v\n", index, rr8_regex)
		// fmt.Printf("[%v] RR11 %v\n\n", index, rr11_regex)
		// if index == 10 {
		// 	os.Exit(1)
		// }
		index++

	}

	// rr8 := NewRegexRule("8: 42 | 42 (42)+")
	// rr11 := NewRegexRule("11: 42 31 | 42 (42 31)+ 31")

	// now this is also equivalent to just throwing in multiple entries to see if any of these rules pass
	// 337 pass  +1 times

	// rr8 := NewRegexRule("8: 42 | 42 42")
	// rr11 := NewRegexRule("11: 42 31 | 42 42 31 31")

	// 285 +2 times
	// rr8 := NewRegexRule("8: 42 | 42 42 42")
	// rr11 := NewRegexRule("11: 42 31 | 42 42 31 42 31 31")

	// +3 times
	// 279
	// rr8 := NewRegexRule("8: 42 | 42 42 42 42")
	// rr11 := NewRegexRule("11: 42 31 | 42 42 31 42 31 42 31 31")

	// +4 times
	// 276
	// rr8 := NewRegexRule("8: 42 | 42 42 42 42 42")
	// rr11 := NewRegexRule("11: 42 31 | 42 42 31 42 31 42 31 42 31 31")

	// +5 times
	// 272
	// rre := NewRegexRuleEngine(DAY_19_INPUT_PART_1)

	// rr8 := NewRegexRule("8: 42 | 42 42 42 42 42 42")
	// rr11 := NewRegexRule("11: 42 31 | 42 42 31 42 31 42 31 42 31 42 31 31")

	// rre.Rules["8"] = rr8
	// rre.Rules["11"] = rr11

	// rre.ParseRules()
	// total, _ := rre.Apply("0")
	// fmt.Printf("Day19.1: %v messages pass.\n", total)

	// rre.Debug()
}

func AOC_2020_19_part1_attempt1(cli *cli.CLI) {
	rre := NewRegexRuleEngine(DAY_19_INPUT_PART_1)
	rre.ParseRules(true)
	total, _ := rre.Apply("0", false)

	fmt.Printf("Day19.1: 133 rules, 471 messages? : %v, %v\n", len(rre.Rules), len(rre.Messages))
	fmt.Printf("Day19.1: %v messages pass.\n", total)

}

type RegexRuleEngine struct {
	Rules    map[string]*RegexRule
	Messages []string
}

// Init parses all Rules to create RegexRules
func (r *RegexRuleEngine) ParseRules(DEBUG bool) {
	// first find our literal rules
	for _, rule := range r.Rules {
		rule.Value = strings.ReplaceAll(rule.Value, "\"", "")
		if rule.Value == "\"a\"" || rule.Value == "\"b\"" {
			regex := strings.ReplaceAll(rule.Value, "\"", "")
			rule.Regex = regex
			rule.Evaluated = true
		}
	}

	if DEBUG {
		fmt.Printf(" >>>>>>> RegexRule: before >>>>> \n\n")
		r.Debug()
		fmt.Printf("\n <<<<<<< RegexRule: before <<<<< \n\n")
	}

	// 	0: 4 1 5
	// 1: 2 3 | 3 2
	// 2: 4 4 | 5 5
	// 3: 4 5 | 5 4
	// 4: "a"
	// 5: "b"

	// now go over each rule until we have a fully evaluated sequence
	totalToEvaluate := len(r.Rules)
	for {
		// loop over until we evaluate everything
		totalEvaluated := 0
		for _, rule := range r.Rules {
			if rule.Evaluated {
				totalEvaluated++
			} else {
				// this requires evaluation
				// split this rule to its sub-rules
				// for any sub-rule that is evaluated, replace the literal with the sub rule regex
				// newValue is the value that contains e.g. 4 5 | 45 45
				// for each evaluated subrule, replace it with the evaluated content
				// newValue := rule.Value

				// the subrules are an array of [ 4, 5, |, 45 ,45 ]
				subRules := strings.Split(rule.Value, " ")
				if DEBUG {
					fmt.Printf("starting value is : '%v'\n", rule.Value)
					fmt.Printf("subrules are      : '%v'\n", subRules)
				}

				// we will go over all subrules and replace any that are evaluated
				// then we will rebuild the rule based on the subrules
				allSubRulesEvaluated := true
				changes := 0
				for subRuleIndex, subRuleKey := range subRules {
					if subRuleKey == "|" {
						continue
					} else if subRuleKey == " " || subRuleKey == "" {
						continue
					} else if subRuleKey == "a" || subRuleKey == "b" {
						continue
					} else {
						// fmt.Printf("rule[%v], SubRules %v SubRule %v\n", rule.Key, subRules, subRuleKey)
						subRule, exists := r.Rules[subRuleKey]
						if !exists {
							if DEBUG {
								fmt.Printf("subRule '%v', exists=%v\n", subRuleKey, exists)
							}
							continue
						}
						if DEBUG {
							fmt.Printf("subRule '%v', exists=%v, subruleValue=%v\n", subRuleKey, exists, subRule.Value)
						}
						if subRule.Evaluated {
							// then this sub-rule has been fully evalulated; we can replace
							// the reference to it with the evaluated content
							// if the ruleKey was 4 and I had others like 4 44 444
							// then I need to e "careful" about how I replace using this key
							// splits := strings.Split(newValue, " ")
							// for _, splitValue := range splits {
							// 	if splitValue == subRuleKey {
							// 		splitValue = subRule.Value
							// 	}
							// 	newValue += " "
							// 	newValue += splitValue
							// }
							// subRuleKey = subRule.Value
							subRules[subRuleIndex] = subRule.Value
							changes++

							// newValue = strings.ReplaceAll(newValue, subRuleKey, subRule.Value)
							// rule.Value = newValue
						} else {
							allSubRulesEvaluated = false
						}
					}
				}

				if changes > 0 {
					// make a string from the subRules as they were changed
					newValue := ""
					for _, value := range subRules {
						newValue += " "
						newValue += value
					}
					rule.Value = newValue
				}

				// now rebuild the subrules to a single value for the rule

				if allSubRulesEvaluated {

					// there seems to be a bug where I am including a number... I don't know why
					testValue := rule.Value
					testValue = strings.ReplaceAll(testValue, "(", "")
					testValue = strings.ReplaceAll(testValue, ")", "")
					testValue = strings.ReplaceAll(testValue, "a", "")
					testValue = strings.ReplaceAll(testValue, "b", "")
					testValue = strings.ReplaceAll(testValue, "|", "")
					testValue = strings.ReplaceAll(testValue, " ", "")

					if DEBUG && testValue != "" {
						fmt.Printf(">>>>>>>>>> INVALID RULE %v >>>>>>>\n", testValue)
						fmt.Printf("%v\n", rule.Line)
						fmt.Printf("%v\n", rule.Value)
						fmt.Printf("<<<<<<<<<< INVALID RULE %v <<<<<<<\n", testValue)

						// fmt.Printf("%vINVALID RULE \n", rule.Debug())
						// fmt.Printf("<<<<<<<<<< INVALID RULE <<<<<<<<\n")

					}

					rule.Evaluated = true
					rule.Regex = strings.ReplaceAll(rule.Value, " ", "")
					if len(rule.Regex) > 1 {
						rule.Regex = "(" + rule.Regex + ")"
					}
					rule.Value = rule.Regex
				}

			}
		}

		if totalEvaluated == totalToEvaluate {
			break
		}

		if DEBUG {
			fmt.Printf(" >>>>>>> RegexRule: during >>>>> \n\n")
			r.Debug()
			fmt.Printf("\n\n <<<<<<< RegexRule: during <<<<< \n\n")
		}

	}

	if DEBUG {
		fmt.Printf(" >>>>>>> RegexRule: after >>>>> \n\n")
		r.Debug()
		fmt.Printf("\n\n <<<<<<< RegexRule: after <<<<< \n\n")
	}

}

func (r *RegexRuleEngine) Debug() {
	for index := 0; index < 10000; index++ {
		sindex := fmt.Sprintf("%v", index)
		rule, exists := r.Rules[sindex]
		if !exists {
			break
		}
		line := rule.Debug()
		fmt.Printf("%v\n", line)
	}
}

func (r *RegexRuleEngine) CheckMessage(message string) (*RegexRule, bool) {
	for _, rule := range r.Rules {
		if rule.IsMessageValid(message) {
			return rule, true
		}
	}
	return nil, false
}

func (r *RegexRuleEngine) Apply(ruleId string, debug bool) (int, []string) {
	rule := r.Rules[ruleId]
	passingMessages := make([]string, 0)
	total := 0
	for _, message := range r.Messages {
		if rule.IsMessageValid(message) {
			total++
			if debug {
				fmt.Printf("REGEX PASSED [id=%v] [regex=%v] '%v'\n", rule.Key, rule.Regex, message)
			}
			passingMessages = append(passingMessages, message)
		} else {
			if debug {
				fmt.Printf("REGEX FAILED [id=%v] [regex=%v] '%v'\n", rule.Key, rule.Regex, message)
			}
		}
	}
	// fmt.Printf("REGEX VALUE %v\n", rule.Regex)
	return total, passingMessages
}

func NewRegexRuleEngine(input string) *RegexRuleEngine {
	splits := strings.Split(input, "\n")
	messages := make([]string, 0)
	ruleMap := make(map[string]*RegexRule)
	useMessages := false
	for _, line := range splits {
		if strings.TrimSpace(line) == "" {
			useMessages = true
			continue
		}
		if useMessages {
			messages = append(messages, line)
		} else {
			rule := NewRegexRule(line)
			ruleMap[rule.Key] = rule
		}
	}
	return &RegexRuleEngine{Rules: ruleMap, Messages: messages}
}

func (rre *RegexRuleEngine) RemoveMessages(messages []string) {
	if len(messages) == 0 {
		return
	}
	retain := make([]string, 0)
	for _, message := range rre.Messages {
		shouldRetain := true
		for _, messageToRemove := range messages {
			if message == messageToRemove {
				// we already have it
				shouldRetain = false
				break
			}
		}
		if shouldRetain {
			retain = append(retain, message)
		}
	}
	rre.Messages = retain
}

func (r *RegexRuleEngine) ParseRulesV3() {
	// first find our literal rules
	min_key := 10
	max_key := 0
	for key, rule := range r.Rules {
		ikey, _ := strconv.Atoi(key)
		min_key = goutils.Min(min_key, ikey)
		max_key = goutils.Max(max_key, ikey)
		rule.Value = strings.ReplaceAll(rule.Value, "\"", "")
		if rule.Value == "\"a\"" || rule.Value == "\"b\"" {
			regex := strings.ReplaceAll(rule.Value, "\"", "")
			rule.Regex = regex
			rule.Evaluated = true
		}
	}

	fmt.Printf("max_key %v\n", max_key)
	fmt.Printf("min_key %v\n", min_key)

	// ascending
	for key := min_key; key <= max_key; key++ {
		skey := fmt.Sprintf("%v", key)
		sourceRule, exists := r.Rules[skey]
		if !exists {
			continue
		}
		if sourceRule.Evaluated {
			for targetKey, targetRule := range r.Rules {
				if targetKey == skey {
					continue
				} else if targetRule.Evaluated {
					continue
				} else {
					// not evaluated; replace any occurances of this sourceRule in the targetRule
					targetRule.Replace(sourceRule)
				}
			}
		}
	}

	// sort the list of rules by key order
}

// Init parses all Rules to create RegexRules
func (r *RegexRuleEngine) ParseRulesV2() {
	// first find our literal rules
	for _, rule := range r.Rules {
		rule.Value = strings.ReplaceAll(rule.Value, "\"", "")
		if rule.Value == "\"a\"" || rule.Value == "\"b\"" {
			regex := strings.ReplaceAll(rule.Value, "\"", "")
			rule.Regex = regex
			rule.Evaluated = true
		}
	}

	fmt.Printf(" >>>>>>> RegexRule: before >>>>> \n\n")
	r.Debug()
	fmt.Printf("\n <<<<<<< RegexRule: before <<<<< \n\n")

	// 	0: 4 1 5
	// 1: 2 3 | 3 2
	// 2: 4 4 | 5 5
	// 3: 4 5 | 5 4
	// 4: "a"
	// 5: "b"

	// now go over each rule until we have a fully evaluated sequence
	DEPTH_CHARGE := 10000

	totalToEvaluate := len(r.Rules)
	for {
		// loop over until we evaluate everything
		totalToEvaluate = len(r.Rules)

		totalEvaluated := 0
		for _, rule := range r.Rules {
			if rule.Evaluated {
				totalEvaluated++
			} else {
				if rule.Attempt >= DEPTH_CHARGE {
					rule.IsLooping = true
					rule.Evaluated = true
				}
				rule.Attempt++
				// this requires evaluation
				// split this rule to its sub-rules
				// for any sub-rule that is evaluated, replace the literal with the sub rule regex
				// newValue is the value that contains e.g. 4 5 | 45 45
				// for each evaluated subrule, replace it with the evaluated content
				// newValue := rule.Value

				// the subrules are an array of [ 4, 5, |, 45 ,45 ]
				subRules := strings.Split(rule.Value, " ")
				fmt.Printf("starting value is : '%v'\n", rule.Value)
				fmt.Printf("subrules are      : '%v'\n", subRules)

				// we will go over all subrules and replace any that are evaluated
				// then we will rebuild the rule based on the subrules
				allSubRulesEvaluated := true
				changes := 0
				for subRuleIndex, subRuleKey := range subRules {
					if subRuleKey == "|" {
						continue
					} else if subRuleKey == " " || subRuleKey == "" {
						continue
					} else if subRuleKey == "a" || subRuleKey == "b" {
						continue
					} else {
						// fmt.Printf("rule[%v], SubRules %v SubRule %v\n", rule.Key, subRules, subRuleKey)
						subRule, exists := r.Rules[subRuleKey]
						if !exists {
							fmt.Printf("subRule '%v', exists=%v\n", subRuleKey, exists)
							continue
						}
						fmt.Printf("subRule '%v', exists=%v, subruleValue=%v\n", subRuleKey, exists, subRule.Value)
						if subRule.Evaluated {
							// then this sub-rule has been fully evalulated; we can replace
							// the reference to it with the evaluated content
							// if the ruleKey was 4 and I had others like 4 44 444
							// then I need to e "careful" about how I replace using this key
							// splits := strings.Split(newValue, " ")
							// for _, splitValue := range splits {
							// 	if splitValue == subRuleKey {
							// 		splitValue = subRule.Value
							// 	}
							// 	newValue += " "
							// 	newValue += splitValue
							// }
							// subRuleKey = subRule.Value
							subRules[subRuleIndex] = subRule.Value
							changes++

							// newValue = strings.ReplaceAll(newValue, subRuleKey, subRule.Value)
							// rule.Value = newValue
						} else {
							allSubRulesEvaluated = false
						}
					}
				}

				if changes > 0 {
					// make a string from the subRules as they were changed
					newValue := ""
					for _, value := range subRules {
						newValue += " "
						newValue += value
					}
					rule.Value = newValue
				}

				// now rebuild the subrules to a single value for the rule

				if allSubRulesEvaluated {

					// there seems to be a bug where I am including a number... I don't know why
					testValue := rule.Value
					testValue = strings.ReplaceAll(testValue, "(", "")
					testValue = strings.ReplaceAll(testValue, ")", "")
					testValue = strings.ReplaceAll(testValue, "a", "")
					testValue = strings.ReplaceAll(testValue, "b", "")
					testValue = strings.ReplaceAll(testValue, "|", "")
					testValue = strings.ReplaceAll(testValue, " ", "")
					if testValue != "" {
						fmt.Printf(">>>>>>>>>> INVALID RULE %v >>>>>>>\n", testValue)
						fmt.Printf("%v\n", rule.Line)
						fmt.Printf("%v\n", rule.Value)
						fmt.Printf("<<<<<<<<<< INVALID RULE %v <<<<<<<\n", testValue)

						// fmt.Printf("%vINVALID RULE \n", rule.Debug())
						// fmt.Printf("<<<<<<<<<< INVALID RULE <<<<<<<<\n")

					}

					rule.Evaluated = true
					rule.Regex = strings.ReplaceAll(rule.Value, " ", "")
					if len(rule.Regex) > 1 {
						rule.Regex = "(" + rule.Regex + ")"
					}
					rule.Value = rule.Regex
				}

			}
		}

		if totalEvaluated == totalToEvaluate {
			break
		}

		fmt.Printf(" >>>>>>> RegexRule: during >>>>> \n\n")
		r.Debug()
		fmt.Printf("\n\n <<<<<<< RegexRule: during <<<<< \n\n")

	}

	fmt.Printf(" >>>>>>> RegexRule: after >>>>> \n\n")
	r.Debug()
	fmt.Printf("\n\n <<<<<<< RegexRule: after <<<<< \n\n")

}

/*
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
*/

type RegexRule struct {
	Line      string // the original line    1: 2 3 | 3 2
	Key       string // the rule id 1
	Value     string // the 2 3 | 3 2
	Evaluated bool   // indicates if the regex has been evaluated
	Regex     string // the regex for this rule
	FullRegex string // the "full" regex composing all regexes (after evaluation)
	Attempt   int
	IsLooping bool
}

func (rr *RegexRule) Debug() string {
	line := fmt.Sprintf("%v", rr.Line)
	return line
}

func NewRegexRule(line string) *RegexRule {
	splits := strings.Split(line, ":")
	key := strings.TrimSpace(splits[0])
	value := strings.TrimSpace(splits[1])
	value = strings.ReplaceAll(value, "\"", "")
	rr := RegexRule{Key: key, Value: value, Line: line, Evaluated: false, IsLooping: false, Attempt: 0}
	return &rr
}

func (rr *RegexRule) IsMessageValid(message string) bool {
	if rr.IsLooping {
		return false
	}
	expr, _ := regexp.Compile("^" + rr.Regex + "$")
	return expr.MatchString(message)
}

// Replace replaces references to the passed rule with the actual regex of the passed rule
func (rr *RegexRule) Replace(rule *RegexRule) {
	splits := strings.Split(rr.Value, " ")
	replacedValue := ""
	changed := false
	for index, value := range splits {
		if value == rule.Key {
			changed = true
			splits[index] = rule.Regex
		}
		replacedValue += " "
		replacedValue += splits[index]
	}
	if changed {
		rr.Value = replacedValue
	}
}

/*
ParseRules
	for each rule, create a regex
	scan list and find literals and create a rule for them
		"a" becomes a



	for each ruleId
		if isLiteral? ("a", "b")
			rule_regex = "a"
		else:
			0: 4 1 5
			1: 2 3 | 3 2
			2: 4 4 | 5 5
			3: 4 5 | 5 4
			4: "a"
			5: "b"

			total_regex = '^regex$

	find literal rules first


	expect 1:1 Rule
	a
	Rule.IsValid(message)
		rule has a regex
		regex.Match(message)

*/
