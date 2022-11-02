package aoc2021

import (
	"fmt"
	"sort"
	"strings"

	utils "github.com/simonski/aoc/utils"
)

/*
--- Day 10: Syntax Scoring ---
You ask the submarine to determine the best route out of the deep-sea cave, but it only replies:

Syntax error in navigation subsystem on line: all of them
All of them?! The damage is worse than you thought. You bring up a copy of the navigation subsystem (your puzzle input).

The navigation subsystem syntax is made of several lines containing chunks. There are one or more chunks on each line, and chunks contain zero or more other chunks. Adjacent chunks are not separated by any delimiter; if one chunk stops, the next chunk (if any) can immediately start. Every chunk must open and close with one of four legal pairs of matching characters:

If a chunk opens with (, it must close with ).
If a chunk opens with [, it must close with ].
If a chunk opens with {, it must close with }.
If a chunk opens with <, it must close with >.
So, () is a legal chunk that contains no other chunks, as is []. More complex but valid chunks include ([]), {()()()}, <([{}])>, [<>({}){}[([])<>]], and even (((((((((()))))))))).

Some lines are incomplete, but others are corrupted. Find and discard the corrupted lines first.

A corrupted line is one where a chunk closes with the wrong character - that is, where the characters it opens and closes with do not form one of the four legal pairs listed above.

Examples of corrupted chunks include (], {()()()>, (((()))}, and <([]){()}[{}]). Such a chunk can appear anywhere within a line, and its presence causes the whole line to be considered corrupted.

For example, consider the following navigation subsystem:

[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]
Some of the lines aren't corrupted, just incomplete; you can ignore these lines for now. The remaining five lines are corrupted:

{([(<{}[<>[]}>{[]{[(<()> - Expected ], but found } instead.
[[<[([]))<([[{}[[()]]] - Expected ], but found ) instead.
[{[{({}]{}}([{[{{{}}([] - Expected ), but found ] instead.
[<(<(<(<{}))><([]([]() - Expected >, but found ) instead.
<{([([[(<>()){}]>(<<{{ - Expected ], but found > instead.
Stop at the first incorrect closing character on each corrupted line.

Did you know that syntax checkers actually have contests to see who can get the high score for syntax errors in a file? It's true! To calculate the syntax error score for a line, take the first illegal character on the line and look it up in the following table:

): 3 points.
]: 57 points.
}: 1197 points.
>: 25137 points.
In the above example, an illegal ) was found twice (2*3 = 6 points), an illegal ] was found once (57 points), an illegal } was found once (1197 points), and an illegal > was found once (25137 points). So, the total syntax error score for this file is 6+57+1197+25137 = 26397 points!

Find the first illegal character in each corrupted line of the navigation subsystem. What is the total syntax error score for those errors?

--- Part 2 ---

Now, discard the corrupted lines. The remaining lines are incomplete.

Incomplete lines don't have any incorrect characters - instead, they're missing some closing characters at the end of the line. To repair the navigation subsystem, you just need to figure out the sequence of closing characters that complete all open chunks in the line.

You can only use closing characters (), ], }, or >), and you must add them in the correct order so that only legal pairs are formed and all chunks end up closed.

In the example above, there are five incomplete lines:

[({(<(())[]>[[{[]{<()<>> - Complete by adding }}]])})].
[(()[<>])]({[<{<<[]>>( - Complete by adding )}>]}).
(((({<>}<{<{<>}{[]{[]{} - Complete by adding }}>}>)))).
{<[[]]>}<{[{[{[]{()[[[] - Complete by adding ]]}}]}]}>.
<{([{{}}[<[[[<>{}]]]>[]] - Complete by adding ])}>.
Did you know that autocomplete tools also have contests? It's true! The score is determined by considering the completion string character-by-character. Start with a total score of 0. Then, for each character, multiply the total score by 5 and then increase the total score by the point value given for the character in the following table:

): 1 point.
]: 2 points.
}: 3 points.
>: 4 points.
So, the last completion string above - ])}> - would be scored as follows:

Start with a total score of 0.
Multiply the total score by 5 to get 0, then add the value of ] (2) to get a new total score of 2.
Multiply the total score by 5 to get 10, then add the value of ) (1) to get a new total score of 11.
Multiply the total score by 5 to get 55, then add the value of } (3) to get a new total score of 58.
Multiply the total score by 5 to get 290, then add the value of > (4) to get a new total score of 294.
The five lines' completion strings have total scores as follows:

}}]])})] - 288957 total points.
)}>]}) - 5566 total points.
}}>}>)))) - 1480781 total points.
]]}}]}]}> - 995444 total points.
])}> - 294 total points.
Autocomplete tools are an odd bunch: the winner is found by sorting all of the scores and then taking the middle score. (There will always be an odd number of scores to consider.) In this example, the middle score is 288957 because there are the same number of scores smaller and larger than it.

Find the completion string for each incomplete line, score the completion strings, and sort the scores. What is the middle score?



*/

const DEBUG = false

type StackDay10 struct {
	Data []string
}

func NewStackDay10() *StackDay10 {
	return &StackDay10{Data: make([]string, 0)}
}
func (stack *StackDay10) Push(value string) {
	stack.Data = append(stack.Data, value)
	if DEBUG {
		fmt.Printf("push: %v     %v\n", value, stack.Flatten())
	}
}
func (stack *StackDay10) Flatten() string {
	result := ""
	for index := 0; index < len(stack.Data); index++ {
		result += stack.Data[index]
	}
	return result
}

func (stack *StackDay10) Pop() string {
	value := stack.Data[len(stack.Data)-1]
	stack.Data = stack.Data[0 : len(stack.Data)-1]
	if DEBUG {
		fmt.Printf("pop : %v     %v\n", value, stack.Flatten())
	}
	return value
}

func (stack *StackDay10) Peek(string) string {
	return stack.Data[len(stack.Data)-1]
}

func (stack *StackDay10) Size() int {
	return len(stack.Data)
}

func isStartCharacter(c string) bool {
	return c == "(" || c == "[" || c == "{" || c == "<"
}

func isEndCharacter(c string) bool {
	return c == ")" || c == "]" || c == "}" || c == ">"
}

func compliment(a string) string {
	if a == "{" {
		return "}"
	} else if a == "(" {
		return ")"
	} else if a == "[" {
		return "]"
	} else if a == "<" {
		return ">"
	} else {
		return ""
	}
}

func doesStartMatch(start string, end string) bool {
	if start == "{" {
		return end == "}"
	}
	if start == "(" {
		return end == ")"
	}
	if start == "[" {
		return end == "]"
	}
	if start == "<" {
		return end == ">"
	}
	return false
}

func day10TestPart1(data string) {
	lines := strings.Split(data, "\n")
	score := 0
	corrupt_lines := 0
	for index, line := range lines {
		line = strings.Trim(line, " ")
		corrupt, c, cindex := isCorrupt(line)
		if corrupt {
			// if DEBUG {
			fmt.Printf("line[%v], corrupt at %v, character=%v, line=%v\n", index, cindex, c, line)
			// }

			if c == ")" {
				score += 3
			} else if c == "]" {
				score += 57
			} else if c == "}" {
				score += 1197
			} else if c == ">" {
				score += 25137
			}
			corrupt_lines += 1
		} else {
			// fmt.Printf("line[%v], not corrupt, line=%v\n", index, line)
		}
		// break
	}
	fmt.Printf("%v corrupt lines scoring %v\n", corrupt_lines, score)

	// ): 3 points.
	// ]: 57 points.
	// }: 1197 points.
	// >: 25137 points.

}

func day10TestPart2(data string) {
	/*

	   // (), ], }, or >)
	   // [({(<(())[]>[[{[]{<()<>> - Complete by adding }}]])})].
	   // [(()[<>])]({[<{<<[]>>( - Complete by adding )}>]}).
	   // (((({<>}<{<{<>}{[]{[]{} - Complete by adding }}>}>)))).
	   // {<[[]]>}<{[{[{[]{()[[[] - Complete by adding ]]}}]}]}>.
	   // <{([{{}}[<[[[<>{}]]]>[]] - Complete by adding ])}>.


	   // (), ], }, or >)
	   // [({(<(())[]>[[{[]{<()<>> - Complete by adding }}]])})].

	   // [({(<(())[]>[[{[]{<()<>> - Complete by adding }}]])})].
	   // for index=0.. length
	   	// 0: if string from current index !isClosed   index0: "[" is not close
	   	//    add compliment to "start" of the suffix, so.. ]
	   	//
	   	// 1: if string from current index !isClosed   index1: "("
	   	      // add compliment to start of the suffix, so )]

	   	// 2: if string from current index !isClosed   index2: "{"
	   	      // add compliment to start of the suffix, so )]
	   //



	   // first must be last
	*/
	lines := strings.Split(data, "\n")

	// discard corrupt, retain only incomplete lines
	incomplete := make([]string, 0)
	for _, line := range lines {
		line = strings.Trim(line, " ")
		corrupt, _, _ := isCorrupt(line)
		if !corrupt {
			incomplete = append(incomplete, line)
		}
	}

	fmt.Printf("We have %v incomplete lines.\n", len(incomplete))

	scores := make([]int, 0)
	for _, line := range incomplete {
		line = strings.Trim(line, " ")
		stack := NewStackDay10()
		for index := 0; index < len(line); index++ {
			c := line[index : index+1]
			if isStartCharacter(c) {
				stack.Push(c)
			} else if isEndCharacter(c) {
				check := stack.Pop()
				if doesStartMatch(check, c) {
					// then we have 'cleared' this good bit from the stack
				}
			}
		}
		// now we are at the end, the stack is whatever is left
		// fmt.Printf("Part2: Line %v becomes %v\n", line, stack.Flatten())

		left := stack.Flatten()
		right := ""
		for index := 0; index < len(left); index++ {
			c := left[index : index+1]
			right = compliment(c) + right
		}

		score := 0
		for index := 0; index < len(right); index++ {
			c := right[index : index+1]
			var value int
			if c == ")" {
				value = 1
			} else if c == "]" {
				value = 2
			} else if c == "}" {
				value = 3
			} else if c == ">" {
				value = 4
			}
			score *= 5
			score += value
		}

		scores = append(scores, score)
		fmt.Printf("%v + %v, =%v\n", left, right, score)
	}
	sort.Ints(scores)
	index := (len(scores)) / 2
	fmt.Printf("Middle score of %v (index=%v) is %v\n", scores, index, scores[index])
}

func isCorrupt(line string) (bool, string, int) {
	stack := NewStackDay10()
	for index := 0; index < len(line); index++ {
		c := line[index : index+1]
		if isStartCharacter(c) {
			stack.Push(c)
		} else if isEndCharacter(c) {
			check := stack.Pop()
			if doesStartMatch(check, c) {
				// valid.. so far
				if DEBUG {
					fmt.Printf("%v is an end character, the last start was %v - matches!.\n", c, check)
				}
			} else {
				// corrupt
				if DEBUG {
					fmt.Printf("%v is an end character, the last start was %v - matches - NO!.\n", c, check)
				}
				return true, c, index
			}
		}
	}

	// all good
	return false, "", -1

}

// // rename this to the year and day in question
// func (app *Application) Y2021D10P2() {
// }

// rename and uncomment this to the year and day in question once complete for a gold star!
// func (app *Application) Y20XXDXXP1Render() {
// }

// rename and uncomment this to the year and day in question once complete for a gold star!
// func (app *Application) Y20XXDXXP2Render() {
// }

// this is what we will reflect and call - so both parts with run. It's up to you to make it print nicely etc.
// The app reference has a CLI for logging.
// func (app *Application) Y2021D10() {
// 	app.Y2021D10P1()
// 	app.Y2021D10P2()
// }

func (app *Application) Y2021D10_Summary() *utils.Summary {
	s := utils.NewSummary(2021, 10)
	s.Name = "Syntax Scoring"
	s.ProgressP1 = utils.Completed
	s.ProgressP2 = utils.Completed
	return s
}

// rename this to the year and day in question
func (app *Application) Y2021D10P1() {
	day10TestPart1(DAY_2021_10_TEST_DATA)
	day10TestPart1(DAY_2021_10_DATA)
}

func (app *Application) Y2021D10P2() {
	day10TestPart2(DAY_2021_10_TEST_DATA)
	day10TestPart2(DAY_2021_10_DATA)
}

// (), ], }, or >)
// [({(<(())[]>[[{[]{<()<>> - Complete by adding }}]])})].

// 1. "clean" the original (WRONG)
// [({(<(())[]>[[{[]{<()<>> - Complete by adding }}]])})].
// [({([[{{ - Complete by adding }}]])})].
