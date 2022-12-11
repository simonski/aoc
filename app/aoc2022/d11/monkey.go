package d11

import (
	"fmt"
	"strconv"
	"strings"
)

type Monkey struct {
	Id                 string
	Items              []uint64
	Operation          string
	OperationOp        string
	OperationValue     uint64
	UseOperationValue  bool
	Test               uint64
	OutcomeMonkeyTrue  string
	OutcomeMonkeyFalse string
	InspectCount       uint64
}

func (m *Monkey) Turn(monkeyIndex int, DEBUG bool, t *Troupe, divideBy uint64) {
	if len(m.Items) == 0 {
		return
	}

	if DEBUG {
		fmt.Printf("\nRound [%v] Monkey[%v]\n", t.RoundNum, monkeyIndex)
		fmt.Printf("  Starting items %v\n", m.Items)
	}
	for _, item := range m.Items {
		m.InspectCount += 1
		if DEBUG {
			fmt.Printf("  Monkey inspects %v\n", item)
		}
		newItem := m.Inspect(item)
		if DEBUG {
			fmt.Printf("    Worry level %v becomes %v\n", item, newItem)
		}
		newItem = newItem / divideBy
		if DEBUG {
			fmt.Printf("    Monkey bored, new value is %v\n", newItem)
		}
		if newItem%m.Test == 0 {
			monkeyId := m.OutcomeMonkeyTrue
			if DEBUG {
				fmt.Printf("    Item %v is thrown to monkey %v\n", newItem, monkeyId)
			}
			t.Get(monkeyId).Add(newItem)
		} else {
			monkeyId := m.OutcomeMonkeyFalse
			if DEBUG {
				fmt.Printf("    Item %v is thrown to monkey %v\n", newItem, monkeyId)
			}
			t.Get(monkeyId).Add(newItem)
		}
	}
	m.Items = make([]uint64, 0)
}

func (m *Monkey) Inspect(item uint64) uint64 {
	var value uint64
	if m.UseOperationValue {
		value = item
	} else {
		value = m.OperationValue
	}
	if m.OperationOp == "/" {
		return item / value
	} else if m.OperationOp == "*" {
		return item * value
	} else if m.OperationOp == "+" {
		return item + value
	} else if m.OperationOp == "-" {
		return item - value
	} else {
		return item
	}
}

func NewMonkey(id string, items_line string, operation_line string, test_line string, true_id_line string, false_id_line string) *Monkey {
	m := Monkey{}
	m.Id = m.ParseId(id)
	m.Items = m.ParseItems(items_line)
	m.Operation = operation_line
	m.OperationOp, m.OperationValue = m.ParseOperation(operation_line)
	m.Test = m.ParseTest(test_line)
	m.OutcomeMonkeyTrue = m.ParseThrow(true_id_line)
	m.OutcomeMonkeyFalse = m.ParseThrow(false_id_line)
	return &m
}

func (m *Monkey) Add(item uint64) {
	m.Items = append(m.Items, item)
}

func (m *Monkey) ParseId(input string) string {
	// "Monkey 0:"
	s := strings.Trim(input, " ")
	s = strings.ReplaceAll(s, "Monkey ", "")
	s = strings.ReplaceAll(s, ":", "")
	return s
}

func (m *Monkey) ParseItems(input string) []uint64 {
	// "Starting items: 87, 57, 63, 86, 87, 53"
	s := strings.Trim(input, " ")
	s = strings.ReplaceAll(s, "Starting items: ", "")
	s = strings.ReplaceAll(s, " ", "")
	splits := strings.Split(s, ",")
	results := make([]uint64, 0)
	for _, value := range splits {
		v, _ := strconv.Atoi(value)
		results = append(results, uint64(v))
	}
	return results
}

func (m *Monkey) ParseOperation(input string) (string, uint64) {
	// "Operation: new = old * 19"
	s := strings.Trim(input, " ")
	s = strings.ReplaceAll(s, "Operation: ", "")
	splits := strings.Split(s, " ")
	operation := splits[3]
	if splits[4] == "old" {
		m.UseOperationValue = true
		return operation, 0
	} else {
		m.UseOperationValue = false
		value, _ := strconv.Atoi(splits[4])
		return operation, uint64(value)
	}
}

func (m *Monkey) ParseTest(input string) uint64 {
	// Test: divisible by 2
	s := strings.Trim(input, " ")
	s = strings.ReplaceAll(s, "Test: divisible by ", "")
	value, _ := strconv.Atoi(s)
	return uint64(value)

}

func (m *Monkey) ParseThrow(input string) string {
	// If true: throw to monkey 1
	s := strings.Trim(input, " ")
	s = strings.ReplaceAll(s, "If true: throw to monkey ", "")
	s = strings.ReplaceAll(s, "If false: throw to monkey ", "")
	return s
}
