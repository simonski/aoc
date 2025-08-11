package d12

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
Day 12: Hot Springs
*/

const TYPE_OPERATIONAL = "."
const TYPE_DAMAGED = "#"
const TYPE_UNKNOWN = "?"

type Grid struct {
	data []*Row
}

type Rule struct {
	rules   []int
	pattern string
	regex   *regexp.Regexp
}

func NewRule(rules []int) *Rule {
	r := Rule{}
	r.rules = rules
	pattern, regex := BuildRegex(rules)
	r.pattern = pattern
	r.regex = regex
	return &r
}

func BuildRegex(rules []int) (string, *regexp.Regexp) {

	pattern := "\\.*" // start any number of dots
	for index, rule := range rules {
		if index == 0 {
			pattern = fmt.Sprintf("\\.*#{%v}", rule) // any number of dots followed by a number of damageds
		} else {
			pattern = fmt.Sprintf("%v.\\.*#{%v}", pattern, rule) // previous follwoedb any number of dots followed by damageds
		}
	}

	// append optional any number of . on the end
	pattern = fmt.Sprintf("%v%v", pattern, "\\.*")

	// fmt.Printf("   IsValid: line='%v', Regex: '%v', Candidate: '%v'\n", r.line, pattern, candidate)
	regex, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}
	return pattern, regex

}

type Row struct {
	line                 string
	cols                 []string
	rules                []int
	totalDamagedRequired int
	Rule                 *Rule
	left                 string
	right                string
	bits                 int
}

func (r *Row) IsOperational(index int) bool {
	return r.cols[index] == TYPE_OPERATIONAL
}

func (r *Row) IsDamaged(index int) bool {
	return r.cols[index] == TYPE_DAMAGED
}

func (r *Row) IsUnknown(index int) bool {
	return r.cols[index] == TYPE_UNKNOWN
}

func NewRow(input string) *Row {
	row := Row{line: input}
	cols := make([]string, 0)
	splits := strings.Split(input, " ")
	left := splits[0]
	right := splits[1]
	row.left = left
	row.right = right

	for _, c := range left {
		cols = append(cols, string(c))
	}
	row.cols = cols
	row.rules = utils.SplitDataToListOfInts(right, ",")
	row.totalDamagedRequired = 0

	for _, r := range row.rules {
		row.totalDamagedRequired += r
	}

	l := row.left
	l = strings.ReplaceAll(l, "#", "")
	l = strings.ReplaceAll(l, ".", "")

	row.bits = len(l)

	row.Rule = NewRule(row.rules)
	return &row
}

func (r *Row) Debug() string {
	result := ""
	for _, c := range r.cols {
		result = fmt.Sprintf("%v%v", result, string(c))
	}
	return result
}

func NewGrid(input string) *Grid {
	rowData := strings.Split(input, "\n")
	rows := make([]*Row, 0)
	for _, rowStr := range rowData {
		row := NewRow(rowStr)
		rows = append(rows, row)
	}
	g := Grid{data: rows}
	return &g
}

func (g *Grid) Debug() string {
	result := ""
	for _, line := range g.data {
		result = fmt.Sprintf("%v%v\n", result, line.Debug())
	}
	return result
}

func (r *Row) GetUnknowns() []int {
	results := make([]int, 0)
	for index := range r.cols {
		if r.IsUnknown(index) {
			results = append(results, index)
		}
	}
	return results
}

func (r *Row) GetDamaged() []int {
	results := make([]int, 0)
	for index := range r.cols {
		if r.IsDamaged(index) {
			results = append(results, index)
		}
	}
	return results
}

func (r *Row) Grow() *Row {
	splits := strings.Split(r.line, " ")
	left := splits[0]
	right := splits[1]
	for index := 0; index < 4; index++ {
		left = fmt.Sprintf("%v?%v", left, splits[0])
		right = fmt.Sprintf("%v,%v", right, splits[1])
	}
	line := fmt.Sprintf("%v %v", left, right)
	return NewRow(line)
}

func (r *Row) CountArrangementsP1() int {

	// return the number of arrangements of unknown that satisfy the rules
	// e.g. a rules of 1,1,3
	// #.#.### 1,1,3					no unknowns, as-is
	// ? ? ### 1,1,3                    ? can be . or #, generate all combinations and assert if true against rule
	// if rule passes then combo good.

	// search space is the number of unknowns as a bitset

	// the combinations are NOT for every combo on this line
	// a 20-bit line is only 1m variations so it's not too bad if I just do it anyway but I want to reduce it further
	// so I make an array of the unknowns and have that generate
	// then if it matches based on unknown count, I can use the index of each to replace in the candidate and
	// see if that passes; so teh nuber of regex searches I do in the end is less.

	size := len(r.GetUnknowns())
	totalCombinations := 1 << uint(size)
	// combinations := make([][]bool, 0)
	count := 0

	// unknowns := r.GetUnknowns()

	totalDamagedAlready := len(r.GetDamaged())
	candidateDamageRequired := r.totalDamagedRequired - totalDamagedAlready

	// valids := make([]string, 0)
	// fmt.Printf("combinations[%v] = %v\n", size, totalCombinations)
	// if true {
	// 	return 0
	// }
	for index := 0; index < totalCombinations; index++ {
		combination := make([]bool, size)
		candidateDamageCount := 0
		// if index%1000000 == 0 {
		// 	fmt.Printf("%v/%v (%v)\n", index, totalCombinations, count)
		// }
		for j := 0; j < size; j++ {
			// Check if the j-th bit is set in the current combination
			result := (index>>uint(j))&1 == 1
			combination[j] = result

			if result {
				candidateDamageCount++
			}
			// if candidateDamageCount == candidateDamageRequired {
			// 	break
			// }
		}

		// if true {
		// 	return 0
		// }
		// this combination can now be tested
		if candidateDamageRequired == candidateDamageCount {
			// must have the correct number of damageds to be a candidate
			candidate := r.CreateCandidate(combination)

			// now I inflate an attempt
			valid := r.IsValid(candidate)
			if valid {
				// fmt.Printf(" valid   CANDIDATE[%v] : '%v' (pattern='%v') '%v' total required is %v, missing is %v, candidate provides %v \n", index, r.line, pattern, candidate, r.totalDamagedRequired, candidateDamageRequired, candidateDamageCount)
				count++
			} else {
				// fmt.Printf(" invalid CANDIDATE[%v] : '%v' (pattern='%v') '%v' total required is %v, missing is %v, candidate provides %v \n", index, r.line, pattern, candidate, r.totalDamagedRequired, candidateDamageRequired, candidateDamageCount)

			}

		} else {
			// fmt.Printf("! CANDIDATE[%v] : '%v' '%v' total required is %v, missing is %v, candidate provides %v \n", index, r.line, candidate, r.totalDamagedRequired, candidateDamageRequired, candidateDamageCount)

		}
		// combinations = append(combinations, combination)
	}

	// for _, v := range valids {
	// 	fmt.Printf("YES: '%v'\n", v)
	// }
	return count

}

func (r *Row) IsValid(candidate string) bool {

	// pattern := "\\.*" // start any number of dots
	// for index, rule := range r.rules {
	// 	if index == 0 {
	// 		pattern = fmt.Sprintf("\\.*#{%v}", rule) // any number of dots followed by a number of damageds
	// 	} else {
	// 		pattern = fmt.Sprintf("%v.\\.*#{%v}", pattern, rule) // previous follwoedb any number of dots followed by damageds
	// 	}
	// }

	// // append optional any number of . on the end
	// pattern = fmt.Sprintf("%v%v", pattern, "\\.*")

	// // fmt.Printf("   IsValid: line='%v', Regex: '%v', Candidate: '%v'\n", r.line, pattern, candidate)

	// regex, err := regexp.Compile(pattern)
	// if err != nil {
	// 	// fmt.Printf("Bad regex: %v\n", err)
	// 	panic(err)
	// }
	matches := r.Rule.regex.FindAllString(candidate, -1)
	return len(matches) > 0
	// 1,1,4,1
	// make a regex that is
	// anythingONEanythingONEanythingFOURanythingONEanything

	// pattern := `\.[a-zA-Z0-9]{4}\.`
	// pattern := `\.\#{3}\.?`

}

func (r *Row) CreateCandidate(combination []bool) string {
	result := ""
	uindex := 0
	for index := 0; index < len(r.cols); index++ {
		if r.cols[index] == TYPE_UNKNOWN {
			if combination[uindex] {
				result = fmt.Sprintf("%v%v", result, TYPE_DAMAGED)
			} else {
				result = fmt.Sprintf("%v%v", result, TYPE_OPERATIONAL)
			}
			uindex += 1
		} else {
			result = fmt.Sprintf("%v%v", result, r.cols[index])
		}
	}
	return result
}

func (g *Grid) Part1(verbose bool) {
	total := 0
	for _, row := range g.data {
		count := row.CountArrangementsP1()
		total += count
		if verbose {
			fmt.Printf("Row '%v' total arrangements is %v\n", row.line, count)
		}
	}
	fmt.Printf("P1 Total is %v\n", total)
}
