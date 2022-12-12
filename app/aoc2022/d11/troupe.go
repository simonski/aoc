package d11

import (
	"fmt"
	"strings"
)

type Troupe struct {
	Monkeys   []*Monkey
	MonkeyMap map[string]*Monkey
	RoundNum  int
	Cache     map[string]int
}

func (t *Troupe) Add(monkey *Monkey) {
	t.Monkeys = append(t.Monkeys, monkey)
	t.MonkeyMap[monkey.Id] = monkey
}

func (t *Troupe) Size() int {
	return len(t.Monkeys)
}

func (t *Troupe) Get(id string) *Monkey {
	return t.MonkeyMap[id]
}

func NewTroupe(input string) *Troupe {
	lines := strings.Split(input, "\n")
	monkeyMap := make(map[string]*Monkey)
	monkeys := make([]*Monkey, 0)
	cache := make(map[string]int)
	troupe := &Troupe{MonkeyMap: monkeyMap, Monkeys: monkeys, Cache: cache}
	for index := 0; index < len(lines); index += 7 {
		idline := lines[index]
		items := lines[index+1]
		operation := lines[index+2]
		test := lines[index+3]
		outcome_true := lines[index+4]
		outcome_false := lines[index+5]
		monkey := NewMonkey(idline, items, operation, test, outcome_true, outcome_false)
		troupe.Add(monkey)
	}
	return troupe
}

func (t *Troupe) Round(DEBUG bool, divideBy uint) {
	t.RoundNum += 1
	for index, monkey := range t.Monkeys {
		monkey.Turn(index, DEBUG, t, divideBy)
	}
}

func (t *Troupe) Debug() {
	for index, monkey := range t.Monkeys {
		fmt.Printf("Monkey[%v] %v\n", index, monkey.Items)
	}
}
