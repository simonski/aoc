package d11

import (
	"fmt"
	"strconv"
	"strings"
)

type Monkey struct {
	Id                 string
	Items              []uint
	Operation          string
	OperationOp        string
	OperationValue     uint
	UseOperationValue  bool
	Test               uint
	OutcomeMonkeyTrue  string
	OutcomeMonkeyFalse string
	InspectCount       int
}

func (m *Monkey) Turn(monkeyIndex int, DEBUG bool, troupe *Troupe, divideBy uint) {
	if len(m.Items) == 0 {
		return
	}

	if DEBUG {
		fmt.Printf("\nRound [%v] Monkey[%v]\n", troupe.RoundNum, monkeyIndex)
		fmt.Printf("  Starting items %v\n", m.Items)
	}
	for _, item := range m.Items {
		// m.Add(1)
		m.InspectCount += 1
		if DEBUG {
			fmt.Printf("  Monkey inspects %v\n", item)
		}
		newItem := m.Inspect(item)
		if DEBUG {
			fmt.Printf("    Worry level %v becomes %v\n", item, newItem)
		}
		if divideBy > 1 {
			newItem = newItem / divideBy //newItem.Div(newItem, big.NewInt(int64(divideBy)))
		}
		if DEBUG {
			fmt.Printf("    Monkey bored, new value is %v\n", newItem)
		}

		if newItem%m.Test == 0 {
			newItem = m.Test
			monkeyId := m.OutcomeMonkeyTrue
			if DEBUG {
				fmt.Printf("    Item %v is thrown to monkey %v\n", newItem, monkeyId)
			}
			troupe.Get(monkeyId).Add(newItem)
		} else {
			monkeyId := m.OutcomeMonkeyFalse
			if DEBUG {
				fmt.Printf("    Item %v is thrown to monkey %v\n", newItem, monkeyId)
			}
			troupe.Get(monkeyId).Add(newItem)
		}
	}
	m.Items = make([]uint, 0)
}

func (m *Monkey) Inspect(item uint) uint {
	var value uint
	if m.UseOperationValue {
		value = item
	} else {
		value = m.OperationValue
	}
	if m.OperationOp == "*" {
		return item * value
		// return item.Mul(item, value)
	} else if m.OperationOp == "+" {
		return item + value
		// return item.Add(item, value)
	} else if m.OperationOp == "-" {
		return item - value
		// return item.Sub(item, value)
	} else if m.OperationOp == "/" {
		return item / value
		// return item.Div(item, value)
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

func (m *Monkey) Add(item uint) {
	m.Items = append(m.Items, item)
}

func (m *Monkey) ParseId(input string) string {
	// "Monkey 0:"
	s := strings.Trim(input, " ")
	s = strings.ReplaceAll(s, "Monkey ", "")
	s = strings.ReplaceAll(s, ":", "")
	return s
}

func (m *Monkey) ParseItems(input string) []uint {
	// "Starting items: 87, 57, 63, 86, 87, 53"
	s := strings.Trim(input, " ")
	s = strings.ReplaceAll(s, "Starting items: ", "")
	s = strings.ReplaceAll(s, " ", "")
	splits := strings.Split(s, ",")
	results := make([]uint, 0)
	for _, value := range splits {
		v, _ := strconv.Atoi(value)
		// bi := big.NewInt(int64(v))
		results = append(results, uint(v))
	}
	return results
}

func (m *Monkey) ParseOperation(input string) (string, uint) {
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
		bi := uint(value)
		return operation, bi
	}
}

func (m *Monkey) ParseTest(input string) uint {
	// Test: divisible by 2
	s := strings.Trim(input, " ")
	s = strings.ReplaceAll(s, "Test: divisible by ", "")
	value, _ := strconv.Atoi(s)
	return uint(value)

}

func (m *Monkey) ParseThrow(input string) string {
	// If true: throw to monkey 1
	s := strings.Trim(input, " ")
	s = strings.ReplaceAll(s, "If true: throw to monkey ", "")
	s = strings.ReplaceAll(s, "If false: throw to monkey ", "")
	return s
}
