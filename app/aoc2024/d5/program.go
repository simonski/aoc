package d5

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
Day 5: Print Queue
*/

type Puzzle struct {
	title     string
	year      int
	day       int
	input     string
	lines     []string
	rulestore *RuleStore
}

type RuleStore struct {
	rules   map[int]*Rule
	updates []*Update
}

func (rs *RuleStore) Debug() {
	for _, r := range rs.rules {
		fmt.Printf("rule %v\n", r.rule_id)
		for k, _ := range r.children {
			fmt.Printf("    %v\n", k)
		}
	}
	fmt.Println()
}

func NewRuleStore() *RuleStore {
	rs := &RuleStore{}
	rs.rules = make(map[int]*Rule)
	rs.updates = make([]*Update, 0)
	return rs
}

func (rs *RuleStore) AddRule(r *Rule) {
	rs.rules[r.rule_id] = r
}

func (rs *RuleStore) AddUpdate(upd *Update) {
	rs.updates = append(rs.updates, upd)
}

func (rs *RuleStore) GetRule(key int) *Rule {
	return rs.rules[key]
}

func (rs *RuleStore) Load(data []string) {
	isRule := true
	for _, line := range data {
		if line == "" {
			isRule = false
			continue
		}
		if isRule {
			values := utils.SplitDataToListOfInts(line, "|")
			rule_id := values[0]
			child_id := values[1]
			rule := rs.GetRule(rule_id)
			if rule == nil {
				rule = NewRule(rule_id)
				rs.AddRule(rule)
			}

			child := rs.GetRule(child_id)
			if child == nil {
				child = NewRule(child_id)
				rs.AddRule(child)
			}
			rule.AddChild(child)

		} else {
			upd := NewUpdate(line)
			rs.AddUpdate(upd)
		}
	}
}

type Update struct {
	line   string
	values []int
}

func NewUpdate(line string) *Update {
	upd := Update{}
	upd.line = line
	upd.values = utils.SplitDataToListOfInts(line, ",")
	return &upd
}

func (upd *Update) MiddleValue() int {
	index := len(upd.values) / 2
	return upd.values[index]
}

type Rule struct {
	rule_id  int
	children map[int]*Rule
}

func (r *Rule) AddChild(child *Rule) {
	r.children[child.rule_id] = child
}

func (r *Rule) Contains(value int) bool {
	_, rule := r.children[value]
	return rule
}

func NewRule(rule_id int) *Rule {
	rule := Rule{}
	rule.rule_id = rule_id
	rule.children = make(map[int]*Rule, 0)
	return &rule
}

func (puzzle *Puzzle) GetSummary() utils.Summary {
	s := utils.Summary{Day: puzzle.day, Year: puzzle.year, Name: puzzle.title}
	s.ProgressP1 = utils.NotStarted
	s.ProgressP2 = utils.NotStarted
	s.DateStarted = "2024-12-05 14:00:42"
	return s
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("2024")
	iday, _ := strconv.Atoi("5")
	p := Puzzle{year: iyear, day: iday, title: "Day 5: Print Queue"}
	p.Load(input)
	return &p
}

func NewPuzzle() *Puzzle {
	return NewPuzzleWithData(REAL_DATA)
}

func (puzzle *Puzzle) Load(input string) {
	lines := strings.Split(input, "\n")
	puzzle.input = input
	puzzle.lines = lines
	rs := NewRuleStore()
	rs.Load(lines)
	puzzle.rulestore = rs
	rs.Debug()
}

func (rs *RuleStore) isValid(updates []int) bool {
	if len(updates) == 1 {
		return true
	}
	rule := rs.GetRule(updates[0])
	for index := 1; index < len(updates); index++ {
		if !rule.Contains(updates[index]) {
			return false
		}
	}
	next_updates := updates[1:]
	return rs.isValid(next_updates)
}

func (rs *RuleStore) reorder(updates []int) []int {
	// update is currently in the wrong order
	sort.Slice(updates, func(i int, j int) bool {
		r1 := rs.GetRule(updates[i])
		r2 := rs.GetRule(updates[j])
		if r1.Contains(r2.rule_id) {
			return true
		}
		return false
	})
	return updates
}

func (puzzle *Puzzle) Part1() {
	puzzle.Load(REAL_DATA)
	rs := puzzle.rulestore
	passes := make([]*Update, 0)

	for _, upd := range rs.updates {
		if rs.isValid(upd.values) {
			passes = append(passes, upd)
		}
	}

	total := 0
	for _, pass := range passes {
		fmt.Println(pass)
		total += pass.MiddleValue()
	}
	fmt.Println(total)
}

func (puzzle *Puzzle) Part2() {
	puzzle.Load(REAL_DATA)
	rs := puzzle.rulestore
	fails := make([][]int, 0)

	for _, upd := range rs.updates {
		if !rs.isValid(upd.values) {
			result := rs.reorder(upd.values)
			fails = append(fails, result)
		}
	}

	total := 0
	for _, result := range fails {
		fmt.Println(result)
		value := result[len(result)/2]
		total += value
	}
	fmt.Println(total)
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}
