package main

import (
	"fmt"
	"regexp"
	"strings"

	goutils "github.com/simonski/goutils"
)

// AOC_2020_06 is the entrypoint to the various attempts for day six
func AOC_2020_06(cli *goutils.CLI) {
	AOC_2020_06_part1_attempt1(cli)
}

func AOC_2020_06_part1_attempt1(cli *goutils.CLI) {
	filename := cli.GetFileExistsOrDie("-input")
	q := NewQandAFromFile(filename)
	fmt.Printf("There are %v total answers.\n", q.TotalForAllGroups())
}

type QandA struct {
	lines []string
}

func (q *QandA) ValidCharacter(answer string) bool {
	expression, _ := regexp.Compile(`[a-z]`)
	return expression.MatchString(answer)
}

// parse queries the content in QandA to store total answer
func (q *QandA) TotalForAllGroups() int {
	answers := make(map[string]int)
	grandTotal := 0
	for index := 0; index < len(q.lines); index++ {
		line := q.lines[index]
		if line == "" {
			grandTotal += len(answers)
			answers = make(map[string]int)
		}
		for position := 0; position < len(line); position++ {
			key := line[position : position+1]
			if q.ValidCharacter(key) {
				answers[key] = 1
			}
		}
	}
	grandTotal += len(answers)
	return grandTotal
}

func NewQandA(content string) *QandA {
	lines := strings.Split(content, "\n")
	q := &QandA{lines: lines}
	return q
}

func NewQandAFromFile(filename string) *QandA {
	lines := load_file_to_strings(filename)
	q := &QandA{lines: lines}
	return q
}
