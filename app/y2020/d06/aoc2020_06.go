package aoc2020

/*

I decided not to use structs representing groups, just a single overall data with some functions; this has helped me in terms of the code is a mess and
super-procedural.  So it works but the maintainability I think is affected.



That said, people call this “idiomatic” in the go community.  Which I take it is meant to mean “the go-ish way of doing things” but I sort of feel it
may be dogma in a while.  So, my take is: use structs representing down to possibly the smallest data structure, then use those tested structures to verify.

As it is, I Went the other way this time and whilst it works, I require a rewrite for future simon!


*/
import (
	"fmt"
	"regexp"
	"strings"

	"github.com/simonski/aoc/utils"
	goutils "github.com/simonski/goutils"
)

func (app *Application) Y2020D06_Summary() *utils.Summary {
	s := utils.NewSummary(2020, 6)
	s.Name = "Custom Customs"
	s.ProgressP1 = utils.Completed
	s.ProgressP2 = utils.Completed
	return s
}

// AOC_2020_06 is the entrypoint
func (app *Application) Y2020D06P1() {
	AOC_2020_06_part1_attempt1(app)
}

func (app *Application) Y2020D06P2() {
	AOC_2020_06_part2_attempt1(app)
}

func AOC_2020_06_part1_attempt1(app *Application) {
	cli := app.CLI
	filename := cli.GetFileExistsOrDie("-input")
	q := NewQandAFromFile(filename)
	fmt.Printf("There are %v total answers.\n", q.TotalForAllGroups())
}

func AOC_2020_06_part2_attempt1(app *Application) {
	cli := app.CLI
	filename := cli.GetFileExistsOrDie("-input")
	q := NewQandAFromFile(filename)
	fmt.Printf("There are %v total answers in the second round.\n", q.SecondTotalForAllGroups())
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

func CheckGroupAnwsers(groupSize int, answers map[string]int) int {
	total := 0
	for key := range answers {
		value := answers[key]
		if value == groupSize {
			total++
		}
	}
	return total
}

func (q *QandA) SecondTotalForAllGroups() int {
	// for each group, it is the number of questions which yes count matches the group size
	answers := make(map[string]int)
	grandTotal := 0
	groupSize := 0
	for index := 0; index < len(q.lines); index++ {
		line := q.lines[index]
		line = strings.TrimSpace(line)
		if line == "" {
			// new line, take the last answers and get the totals from them
			grandTotal += CheckGroupAnwsers(groupSize, answers)
			answers = make(map[string]int)
			groupSize = 0
		} else {
			groupSize++
			for position := 0; position < len(line); position++ {
				key := line[position : position+1]
				if q.ValidCharacter(key) {
					count, exists := answers[key]
					if exists {
						count++
					} else {
						count = 1
					}
					answers[key] = count
				}
			}
		}

	}
	grandTotal += CheckGroupAnwsers(groupSize, answers)
	return grandTotal
}

func NewQandA(content string) *QandA {
	lines := strings.Split(content, "\n")
	q := &QandA{lines: lines}
	return q
}

func NewQandAFromFile(filename string) *QandA {
	lines := goutils.Load_file_to_strings(filename)
	q := &QandA{lines: lines}
	return q
}
