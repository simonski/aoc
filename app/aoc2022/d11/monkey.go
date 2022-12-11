package d11

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

type Monkey struct {
	Id                 string
	Items              []*big.Int
	Operation          string
	OperationOp        string
	OperationValue     *big.Int
	UseOperationValue  bool
	Test               *big.Int
	OutcomeMonkeyTrue  string
	OutcomeMonkeyFalse string
	InspectCount       int
}

func (m *Monkey) Turn(monkeyIndex int, DEBUG bool, troupe *Troupe, divideBy uint64) {
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
			newItem = newItem.Div(newItem, big.NewInt(int64(divideBy)))
		}
		if DEBUG {
			fmt.Printf("    Monkey bored, new value is %v\n", newItem)
		}

		// key := fmt.Sprintf("%v/%v", newItem, m.Test)
		// hitCount := troupe.Cache[key]
		// hitCount += 1
		// troupe.Cache[key] = hitCount
		// if hitCount > 1 {
		// 	fmt.Printf("%v = %v\n", hitCount, key)
		// }

		var x big.Int
		x.Mod(newItem, m.Test)
		zero := big.NewInt(0)
		// fmt.Printf("item is %v digits.\n", len(newItem.String()))
		if x.Cmp(zero) == 0 { //x.Int64() == 0 {
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
	m.Items = make([]*big.Int, 0)
}

func (m *Monkey) Inspect(item *big.Int) *big.Int {
	var value *big.Int
	if m.UseOperationValue {
		value = item
	} else {
		value = m.OperationValue
	}
	if m.OperationOp == "*" {
		return item.Mul(item, value)
	} else if m.OperationOp == "+" {
		return item.Add(item, value)
	} else if m.OperationOp == "-" {
		return item.Sub(item, value)
	} else if m.OperationOp == "/" {
		return item.Div(item, value)
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

func (m *Monkey) Add(item *big.Int) {
	m.Items = append(m.Items, item)
}

func (m *Monkey) ParseId(input string) string {
	// "Monkey 0:"
	s := strings.Trim(input, " ")
	s = strings.ReplaceAll(s, "Monkey ", "")
	s = strings.ReplaceAll(s, ":", "")
	return s
}

func (m *Monkey) ParseItems(input string) []*big.Int {
	// "Starting items: 87, 57, 63, 86, 87, 53"
	s := strings.Trim(input, " ")
	s = strings.ReplaceAll(s, "Starting items: ", "")
	s = strings.ReplaceAll(s, " ", "")
	splits := strings.Split(s, ",")
	results := make([]*big.Int, 0)
	for _, value := range splits {
		v, _ := strconv.Atoi(value)
		bi := big.NewInt(int64(v))
		results = append(results, bi)
	}
	return results
}

func (m *Monkey) ParseOperation(input string) (string, *big.Int) {
	// "Operation: new = old * 19"
	s := strings.Trim(input, " ")
	s = strings.ReplaceAll(s, "Operation: ", "")
	splits := strings.Split(s, " ")
	operation := splits[3]
	if splits[4] == "old" {
		m.UseOperationValue = true
		return operation, big.NewInt(0)
	} else {
		m.UseOperationValue = false
		value, _ := strconv.Atoi(splits[4])
		bi := big.NewInt(int64(value))
		return operation, bi
	}
}

func (m *Monkey) ParseTest(input string) *big.Int {
	// Test: divisible by 2
	s := strings.Trim(input, " ")
	s = strings.ReplaceAll(s, "Test: divisible by ", "")
	value, _ := strconv.Atoi(s)
	return big.NewInt(int64(value))

}

func (m *Monkey) ParseThrow(input string) string {
	// If true: throw to monkey 1
	s := strings.Trim(input, " ")
	s = strings.ReplaceAll(s, "If true: throw to monkey ", "")
	s = strings.ReplaceAll(s, "If false: throw to monkey ", "")
	return s
}
