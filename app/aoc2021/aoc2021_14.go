package aoc2021

import (
	"fmt"
	"strings"
)

/*
--- Day 14: Extended Polymerization ---
The incredible pressures at this depth are starting to put a strain on your submarine. The submarine has polymerization equipment that would produce suitable materials to reinforce the submarine, and the nearby volcanically-active caves should even have the necessary input elements in sufficient quantities.

The submarine manual contains instructions for finding the optimal polymer formula; specifically, it offers a polymer template and a list of pair insertion rules (your puzzle input). You just need to work out what polymer would result after repeating the pair insertion process a few times.

For example:

NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C
The first line is the polymer template - this is the starting point of the process.

The following section defines the pair insertion rules. A rule like AB -> C means that when elements A and B are immediately adjacent, element C should be inserted between them. These insertions all happen simultaneously.

So, starting with the polymer template NNCB, the first step simultaneously considers all three pairs:

The first pair (NN) matches the rule NN -> C, so element C is inserted between the first N and the second N.
The second pair (NC) matches the rule NC -> B, so element B is inserted between the N and the C.
The third pair (CB) matches the rule CB -> H, so element H is inserted between the C and the B.
Note that these pairs overlap: the second element of one pair is the first element of the next pair. Also, because all pairs are considered simultaneously, inserted elements are not considered to be part of a pair until the next step.

After the first step of this process, the polymer becomes NCNBCHB.

Here are the results of a few steps using the above rules:

Template:     NNCB
After step 1: NCNBCHB
After step 2: NBCCNBBBCBHCB
After step 3: NBBBCNCCNBBNBNBBCHBHHBCHB
After step 4: NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB

This polymer grows quickly. After step 5, it has length 97; After step 10, it has length 3073. After step 10, B occurs 1749 times, C occurs 298 times, H occurs 161 times, and N occurs 865 times; taking the quantity of the most common element (B, 1749) and subtracting the quantity of the least common element (H, 161) produces 1749 - 161 = 1588.

Apply 10 steps of pair insertion to the polymer template and find the most and least common elements in the result. What do you get if you take the quantity of the most common element and subtract the quantity of the least common element?


*/

var DEBUG_TIMING = false

// type InstructionLoopEntry struct {
// 	Instruction *Instruction
// 	Counter     *Counter
// 	Path        string
// }
// type InstructionLoop struct {
// 	Key         string
// 	Size        int
// 	Instruction *Instruction
// 	// Instructions     []*Instruction
// 	// FinalInstruction *Instruction
// 	RepeatIndex int
// 	Index       int
// 	// Counters         []*Counter
// 	Entries []*InstructionLoopEntry
// }

// func (l *InstructionLoop) GetCounterAtDepth(depth int) *Counter {
// 	return l.Entries[depth].Counter
// }

// func (l *InstructionLoop) Debug() string {
// 	line := fmt.Sprintf("Loop[%v->%v] ", l.Key, l.Instruction.InjectWith)
// 	for index := 0; index < len(l.Entries); index++ {
// 		line += l.Entries[index].Instruction.Pair
// 		if index+1 < len(l.Entries) {
// 			line += "->"
// 		}
// 	}
// 	line = fmt.Sprintf("%v", line)
// 	return line
// }

type Instruction struct {
	Pair       string
	InjectWith string
}

func NewInstruction(line string) *Instruction {
	line2 := strings.ReplaceAll(line, " ", "")
	splits := strings.Split(line2, "->")
	pair := splits[0]
	injectWith := splits[1]
	return &Instruction{Pair: pair, InjectWith: injectWith}
}

type Polymer struct {
	Data         string
	Value        string
	Instructions map[string]*Instruction
}

func NewPolymer(data string) *Polymer {
	fmt.Printf("%v\n", data)
	lines := strings.Split(data, "\n")
	value := lines[0]
	instructions := make(map[string]*Instruction)
	for index := 2; index < len(lines); index++ {
		i := NewInstruction(lines[index])
		instructions[i.Pair] = i
	}
	return &Polymer{Data: data, Value: value, Instructions: instructions}
}

// rename this to the year and day in question
func (app *Application) Y2021D14P1() {
	// d14P1(DAY_2021_14_TEST_DATA, 10)
	// d14P1(DAY_2021_14_DATA, 10)
}

func (app *Application) Y2021D14P2() {
	// go_on_then_1_working(DAY_2021_14_TEST_DATA, 10)
	// go_on_then_2(DAY_2021_14_TEST_DATA, 40)
	go_on_then_2(DAY_2021_14_DATA, 40)
}

func countLetters(data string) map[string]int64 {
	m := make(map[string]int64)
	for index := 0; index < len(data); index++ {
		letter := data[index : index+1]
		value := m[letter]
		value += 1
		m[letter] = value
	}
	return m
}

func go_on_then_2(data string, iterations int) {
	polymer := NewPolymer(data)
	pairs := Pairs(polymer.Value)

	pairMap := make(map[string]int64)
	for _, pair := range pairs {
		pairMap[pair] += 1
	}
	letterCounter := countLetters(polymer.Value)
	for index := 0; index < iterations; index++ {
		new_pairMap := make(map[string]int64)
		// for pair, pairCount := range pairMap {
		// 	new_pairMap[pair] = pairCount
		// }

		for pair, pairCount := range pairMap {
			// take a pair, e.g. NC, inject H, this means we replace NC with NH and HC
			// so our paircounter will go down by 1 NC and up by 1 NH and one HC
			instruction := polymer.Instructions[pair]
			leftPair := pair[0:1] + instruction.InjectWith
			rightPair := instruction.InjectWith + pair[1:2]

			new_pairMap[leftPair] += pairCount
			new_pairMap[rightPair] += pairCount

			letterCounter[instruction.InjectWith] += pairCount

		}
		pairMap = new_pairMap
	}

	for key, value := range letterCounter {
		fmt.Printf("%v = %v\n", key, value)
	}
}
