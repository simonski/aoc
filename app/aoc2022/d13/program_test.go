package d13

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func Test_1(t *testing.T) {
	p := NewPuzzleWithData(TEST_DATA)
	fmt.Printf("There are %v lines.\n", len(p.lines))
	fmt.Printf("There are %v pairs.\n", len(p.Pairs))
	if len(p.Pairs) != 8 {
		t.Fatalf("There should be %v paris, there are %v\n", 8, len(p.Pairs))
	}
	for index, p := range p.Pairs {
		if p.Left.String() == "" {
			t.Fatalf("Pair[%v] line 1 was empty.\n", index)
		}
		if p.Right.String() == "" {
			t.Fatalf("Pair[%v] line 2 was empty.\n", index)
		}
	}

	pair := p.Pairs[0]
	if pair.Left.String() != "[1,1,3,1,1]" {
		t.Fatalf("Pair[0] Line1 is invalid.")
	}
}

func Test_2(t *testing.T) {
	p := NewPuzzleWithData(REAL_DATA)
	fmt.Printf("There are %v lines.\n", len(p.lines))
	fmt.Printf("There are %v pairs.\n", len(p.Pairs))
	if len(p.Pairs) != 150 {
		t.Fatalf("There should be %v paris, there are %v\n", 150, len(p.Pairs))
	}
	for index, p := range p.Pairs {
		if p.Left.String() == "" {
			t.Fatalf("Pair[%v] line 1 was empty.\n", index)
		}
		if p.Right.String() == "" {
			t.Fatalf("Pair[%v] line 2 was empty.\n", index)
		}
	}

}

func Test_4(t *testing.T) {
	p := BuildPacket(true, "[1,221,3,4,5]")
	if p.PacketType != LIST {
		t.Fatalf("Packet should be LIST, was %v\n", p.PacketType)
	}
	if p.Size() != 5 {
		t.Fatalf("Packet should be sizse 5, was %v\n", p.Size())
	}
	if p.Entries[0].PacketType != INTEGER {
		t.Fatalf("Packet[%v] should be type INTEGER, was %v\n", 0, p.Entries[0].PacketType)
	}
	if p.Entries[1].Value != 221 {
		t.Fatalf("Packet[%v] Value should be %v, was %v\n", 1, 221, p.Entries[1].Value)
	}

}

func Test_5(t *testing.T) {
	p := BuildPacket(true, "[1,221,[],3,4,5]")
	if p.PacketType != LIST {
		t.Fatalf("Packet should be LIST, was %v\n", p.PacketType)
	}
	if p.Size() != 6 {
		t.Fatalf("Packet should be size 6, was %v\n", p.Size())
	}
	if p.Entries[0].PacketType != INTEGER {
		t.Fatalf("Packet[%v] should be type INTEGER, was %v\n", 0, p.Entries[0].PacketType)
	}
	if p.Entries[1].Value != 221 {
		t.Fatalf("Packet[%v] Value should be %v, was %v\n", 1, 221, p.Entries[1].Value)
	}
	if p.Entries[2].PacketType != LIST {
		t.Fatalf("Packet[%v] Value should be LIST, was %v\n", 2, p.Entries[2].PacketType)
	}

	if p.Entries[2].Size() != 0 {
		t.Fatalf("Packet[%v] Size() should be 0, was %v\n", 2, p.Entries[2].Size())
	}

}

func Test_6(t *testing.T) {
	p := BuildPacket(true, "[1,221,[1,2,3],3,4,5]")
	if p.PacketType != LIST {
		t.Fatalf("Packet should be LIST, was %v\n", p.PacketType)
	}
	if p.Size() != 6 {
		t.Fatalf("Packet should be size 6, was %v\n", p.Size())
	}
	if p.Entries[0].PacketType != INTEGER {
		t.Fatalf("Packet[%v] should be type INTEGER, was %v\n", 0, p.Entries[0].PacketType)
	}
	if p.Entries[1].Value != 221 {
		t.Fatalf("Packet[%v] Value should be %v, was %v\n", 1, 221, p.Entries[1].Value)
	}
	if p.Entries[2].PacketType != LIST {
		t.Fatalf("Packet[%v] Value should be LIST, was %v\n", 2, p.Entries[2].PacketType)
	}

	if p.Entries[2].Size() != 3 {
		t.Fatalf("Packet[%v] Size() should be 3, was %v\n", 2, p.Entries[2].Size())
	}

}

func Test_Compare1(t *testing.T) {
	left := BuildPacket(true, "[1,1,3,1,1]")
	right := BuildPacket(true, "[1,1,5,1,1]")
	pair1 := &Pair{Left: left, Right: right}
	if !pair1.IsCorrect() {
		t.Fatalf("Pair1 should be correct.")
	}

	left = BuildPacket(true, "[[1],[2,3,4]]")
	right = BuildPacket(true, "[[1],4]")
	pair2 := &Pair{Left: left, Right: right}
	if !pair2.IsCorrect() {
		t.Fatalf("Pair2 should be correct.")
	}

	left = BuildPacket(true, "[9]")
	right = BuildPacket(true, "[[8,7,6]]")
	pair3 := &Pair{Left: left, Right: right}
	if pair3.IsCorrect() {
		t.Fatalf("Pair3 should not be correct.")
	}

	left = BuildPacket(true, "[[4,4],4,4]")
	right = BuildPacket(true, "[[4,4],4,4,4]")
	pair4 := &Pair{Left: left, Right: right}
	if !pair4.IsCorrect() {
		t.Fatalf("Pair4 should be correct.")
	}

	left = BuildPacket(true, "[7,7,7,7]")
	right = BuildPacket(true, "[7,7,7]")
	pair5 := &Pair{Left: left, Right: right}
	if pair5.IsCorrect() {
		t.Fatalf("Pair5 should not be correct.")
	}

	left = BuildPacket(true, "[]")
	right = BuildPacket(true, "[3]")
	pair6 := &Pair{Left: left, Right: right}
	if !pair6.IsCorrect() {
		t.Fatalf("Pair6 should be correct.")
	}

	left = BuildPacket(true, "[[[]]]")
	right = BuildPacket(true, "[[]]")
	pair7 := &Pair{Left: left, Right: right}
	if pair7.IsCorrect() {
		t.Fatalf("Pair7 should not be correct.")
	}

	left = BuildPacket(true, "[1,[2,[3,[4,[5,6,7]]]],8,9]")
	right = BuildPacket(true, "[1,[2,[3,[4,[5,6,0]]]],8,9]")
	pair8 := &Pair{Left: left, Right: right}
	if pair8.IsCorrect() {
		t.Fatalf("Pair8 should not be correct.")
	}

}

func Test_Compare2(t *testing.T) {

	lines := strings.Split(TEST_DATA, "\n")
	pairs := make([]*Pair, 0)
	VERBOSE := true
	for index := 0; index < len(lines); index += 3 {
		line1 := lines[index]
		line2 := lines[index+1]
		pair := &Pair{Left: BuildPacket(VERBOSE, line1), Right: BuildPacket(VERBOSE, line2)}
		pairs = append(pairs, pair)
	}

	total := 0
	fmt.Println("")
	for index, pair := range pairs {
		fmt.Println("")
		if pair.IsCorrect() {
			fmt.Printf("pair [%v] is correct\n", index+1)
			total += (index + 1)
		}
	}
	fmt.Printf("Sum of the indices is %v\n", total)
	t.Fatalf("mm")

}

func Test_Compare3(t *testing.T) {
	//5905 teh correct answer
	lines := strings.Split(REAL_DATA, "\n")
	pairs := make([]*Pair, 0)
	VERBOSE := true
	for index := 0; index < len(lines); index += 3 {
		line1 := lines[index]
		line2 := lines[index+1]
		pair := &Pair{Left: BuildPacket(VERBOSE, line1), Right: BuildPacket(VERBOSE, line2)}
		pairs = append(pairs, pair)
	}

	total := 0
	for index, pair := range pairs {
		if pair.IsCorrect() {
			fmt.Printf("pair [%v] is correct\n", index+1)
			total += (index + 1)
		}
	}
	fmt.Printf("Sum of the indices is %v\n", total)
	t.Fatalf("mm")
	// NOT 7181

}

func Test_Compare_TestData_Part2(t *testing.T) {

	lines := strings.Split(TEST_DATA, "\n")
	pairs := make([]*Pair, 0)
	VERBOSE := true
	for index := 0; index < len(lines); index += 3 {
		line1 := lines[index]
		line2 := lines[index+1]
		pair := &Pair{Left: BuildPacket(VERBOSE, line1), Right: BuildPacket(VERBOSE, line2)}
		pairs = append(pairs, pair)
	}

	fmt.Println("")

	results := make([]*Packet, 0)
	for _, pair := range pairs {
		results = append(results, pair.Left)
		results = append(results, pair.Right)
	}

	results = append(results, BuildPacket(false, "[[6]]"))
	results = append(results, BuildPacket(false, "[[2]]"))

	sort.Slice(results, func(i int, j int) bool {
		left := results[i]
		right := results[j]
		p := Pair{Left: left, Right: right}
		return p.IsCorrect()
	})

	value1 := 0
	value2 := 0
	for index, packet := range results {
		fmt.Printf("%v\n", packet.String())
		if packet.String() == "[[2]]" {
			value1 = index + 1
		} else if packet.String() == "[[6]]" {
			value2 = index + 1
		}
	}
	fmt.Printf("index=%v * %v = %v\n", value1, value2, value1*value2)

	t.Fatal("ff")
}

func Test_Compare_RealData_Part2(t *testing.T) {

	lines := strings.Split(REAL_DATA, "\n")
	pairs := make([]*Pair, 0)
	VERBOSE := true
	for index := 0; index < len(lines); index += 3 {
		line1 := lines[index]
		line2 := lines[index+1]
		pair := &Pair{Left: BuildPacket(VERBOSE, line1), Right: BuildPacket(VERBOSE, line2)}
		pairs = append(pairs, pair)
	}

	fmt.Println("")

	results := make([]*Packet, 0)
	for _, pair := range pairs {
		results = append(results, pair.Left)
		results = append(results, pair.Right)
	}

	results = append(results, BuildPacket(false, "[[6]]"))
	results = append(results, BuildPacket(false, "[[2]]"))

	sort.Slice(results, func(i int, j int) bool {
		left := results[i]
		right := results[j]
		p := Pair{Left: left, Right: right}
		return p.IsCorrect()
	})

	value1 := 0
	value2 := 0
	for index, packet := range results {
		fmt.Printf("%v\n", packet.String())
		if packet.String() == "[[2]]" {
			value1 = index + 1
		} else if packet.String() == "[[6]]" {
			value2 = index + 1
		}
	}
	fmt.Printf("index=%v * %v = %v\n", value1, value2, value1*value2)

	t.Fatal("ff")
}
