package d11

import (
	"strconv"
	"strings"
)

type Monkey struct {
	Id                 string
	Items              []int
	Operation          string
	OperationNew       string
	OperationOld       string
	OperationOp        string
	OperationValue     int
	Test               int
	OutcomeMonkeyTrue  string
	OutcomeMonkeyFalse string
}

func (d *Monkey) Throw() {

}

func NewMonkey(id string, items_line string, operation_line string, test_line string, true_id_line string, false_id_line string) *Monkey {
	m := Monkey{}
	m.Id = m.ParseId(id)
	m.Items = m.ParseItems(items_line)
	m.Operation = operation_line
	m.OperationNew, m.OperationOld, m.OperationOp, m.OperationValue = m.ParseOperation(operation_line)
	m.Test = m.ParseTest(test_line)
	m.OutcomeMonkeyTrue = m.ParseThrow(true_id_line)
	m.OutcomeMonkeyFalse = m.ParseThrow(false_id_line)
	return &m
}

type Troupe struct {
	Monkeys map[string]*Monkey
}

func (t *Troupe) Add(monkey *Monkey) {
	t.Monkeys[monkey.Id] = monkey
}

func (t *Troupe) Size() int {
	return len(t.Monkeys)
}

func (t *Troupe) Get(id string) *Monkey {
	return t.Monkeys[id]
}

func NewTroupe(input string) *Troupe {
	lines := strings.Split(input, "\n")
	monkeys := make(map[string]*Monkey)
	troupe := &Troupe{Monkeys: monkeys}
	for index := 0; index < len(lines)-8; index += 7 {
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

func (m *Monkey) ParseId(input string) string {
	// "Monkey 0:"
	s := strings.Trim(input, " ")
	s = strings.ReplaceAll(s, "Monkey ", "")
	s = strings.ReplaceAll(s, ":", "")
	return s
}

func (m *Monkey) ParseItems(input string) []int {
	// "Starting items: 87, 57, 63, 86, 87, 53"
	s := strings.Trim(input, " ")
	s = strings.ReplaceAll(s, "Starting items: ", "")
	s = strings.ReplaceAll(s, " ", "")
	splits := strings.Split(s, ",")
	results := make([]int, 0)
	for _, value := range splits {
		v, _ := strconv.Atoi(value)
		results = append(results, v)
	}
	return results
}

func (m *Monkey) ParseOperation(input string) (string, string, string, int) {
	// "Operation: new = old * 19"
	s := strings.Trim(input, " ")
	s = strings.ReplaceAll(s, "Operation: ", "")
	splits := strings.Split(s, " ")
	first := splits[0]
	second := splits[2]
	operation := splits[3]
	value, _ := strconv.Atoi(splits[4])
	return first, second, operation, value
}

func (m *Monkey) ParseTest(input string) int {
	// Test: divisible by 2
	s := strings.Trim(input, " ")
	s = strings.ReplaceAll(s, "Test: divisible by ", "")
	value, _ := strconv.Atoi(s)
	return value

}

func (m *Monkey) ParseThrow(input string) string {
	// If true: throw to monkey 1
	s := strings.Trim(input, " ")
	s = strings.ReplaceAll(s, "If true: throw to monkey ", "")
	s = strings.ReplaceAll(s, "If false: throw to monkey ", "")
	return s
}
