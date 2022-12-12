package d11

import (
	"fmt"
	"testing"
)

func Test_MonkeyParseId(t *testing.T) {
	m := Monkey{}
	id := m.ParseId("Monkey 0:")
	if id != "0" {
		t.Fatalf("id should be 0, was %v\n", id)
	}
}

func Test_MonkeyParseItems(t *testing.T) {
	m := Monkey{}
	items := m.ParseItems("Starting items: 87, 57, 63, 86, 87, 53")
	if len(items) != 6 {
		t.Fatalf("items() should be len 6, was %v\n", len(items))
	}
	if items[0] != 87 {
		t.Fatalf("items[0] should be 87, was %v\n", items[0])
	}
	if items[3] != 86 {
		t.Fatalf("items[0] should be 86, was %v\n", items[3])
	}
}

func Test_MonkeyParseOperation(t *testing.T) {
	m := Monkey{}
	operation, value := m.ParseOperation("Operation: new = old * 19")
	if operation != "*" {
		t.Fatalf("operation shoudl be *, was %v\n", operation)
	}
	if value != 19 {
		t.Fatalf("operation should be 19, was %v\n", value)
	}
}

func Test_MonkeyParseOperation2(t *testing.T) {
	m := Monkey{}
	operation, value := m.ParseOperation("Operation: old = new / 119")
	if operation != "/" {
		t.Fatalf("operation shoudl be /, was %v\n", operation)
	}
	if value != 119 {
		t.Fatalf("operation should be 119, was %v\n", value)
	}
}

func Test_MonkeyParseOperation3(t *testing.T) {
	m := Monkey{}
	operation, value := m.ParseOperation("Operation: old = new - 5")
	if operation != "-" {
		t.Fatalf("operation shoudl be -, was %v\n", operation)
	}
	if value != 5 {
		t.Fatalf("operation should be 5, was %v\n", value)
	}
}

func Test_MonkeyParseTest(t *testing.T) {
	m := Monkey{}
	value := m.ParseTest("Test: divisible by 2")
	if value != 2 {
		t.Fatalf("value shoud lbe 2, was %v\n", value)
	}
}

func Test_MonkeyParseThrowTo(t *testing.T) {
	m := Monkey{}
	id := m.ParseThrow("If true: throw to monkey 1")
	if id != "1" {
		t.Fatalf("moneky id shoud be 1., was %v\n", id)
	}

	id = m.ParseThrow("If false: throw to monkey 9")
	if id != "9" {
		t.Fatalf("moneky id shoud be 9., was %v\n", id)
	}
}

func Test_NewTroupe_TestData(t *testing.T) {
	troupe := NewTroupe(TEST_DATA)
	if troupe.Size() != 4 {
		t.Fatalf("Troupe size should %v, was %v\n", 4, troupe.Size())
	}
}

func Test_NewTroupe_RealData(t *testing.T) {
	troupe := NewTroupe(REAL_DATA)
	if troupe.Size() != 8 {
		t.Fatalf("Troupe size should %v, was %v\n", 8, troupe.Size())
	}
}

func Test_NewTroupe_Test_Part1(t *testing.T) {
	troupe := NewTroupe(TEST_DATA)

	for index := 0; index < 20; index++ {
		troupe.Round(false, 3)
		fmt.Printf("\nRound %v\n", troupe.RoundNum)
		troupe.Debug()
	}

	for index, monkey := range troupe.Monkeys {
		fmt.Printf("Monkey[%v] %v inspections.\n", index, monkey.InspectCount)
	}

}

func Test_NewTroupe_Real_Part1(t *testing.T) {
	troupe := NewTroupe(REAL_DATA)

	for index := 0; index < 20; index++ {
		troupe.Round(false, 3)
		fmt.Printf("\nRound %v\n", troupe.RoundNum)
		troupe.Debug()
	}

	for index, monkey := range troupe.Monkeys {
		fmt.Printf("Monkey[%v] %v inspections.\n", index, monkey.InspectCount)
	}

}

func Test_NewTroupe_Test_Part2(t *testing.T) {
	troupe := NewTroupe(TEST_DATA)

	for index := 0; index < 1000; index++ {
		troupe.Round(false, 1)
		fmt.Printf("\nRound %v\n", troupe.RoundNum)

		for index, monkey := range troupe.Monkeys {
			fmt.Printf("Monkey[%v] %v inspections.\n", index, monkey.InspectCount)
		}
	}

}
