package main

/*
--- Day 7: Some Assembly Required ---
This year, Santa brought little Bobby Tables a set of wires and bitwise logic gates! Unfortunately, little Bobby is a little under the recommended age range, and he needs help assembling the circuit.

Each wire has an identifier (some lowercase letters) and can carry a 16-bit signal (a number from 0 to 65535). A signal is provided to each wire by a gate, another wire, or some specific value. Each wire can only get a signal from one source, but can provide its signal to multiple destinations. A gate provides no signal until all of its inputs have a signal.

The included instructions booklet describes how to connect the parts together: x AND y -> z means to connect wires x and y to an AND gate, and then connect its output to wire z.

For example:

123 -> x means that the signal 123 is provided to wire x.
x AND y -> z means that the bitwise AND of wire x and wire y is provided to wire z.
p LSHIFT 2 -> q means that the value from wire p is left-shifted by 2 and then provided to wire q.
NOT e -> f means that the bitwise complement of the value from wire e is provided to wire f.
Other possible gates include OR (bitwise OR) and RSHIFT (right-shift). If, for some reason, you'd like to emulate the circuit instead, almost all programming languages (for example, C, JavaScript, or Python) provide operators for these gates.

For example, here is a simple circuit:

123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i
After it is run, these are the signals on the wires:

d: 72
e: 507
f: 492
g: 114
h: 65412
i: 65079
x: 123
y: 456
In little Bobby's kit's instructions booklet (provided as your puzzle input), what signal is ultimately provided to wire a?
*/

import (
	"fmt"
	"strconv"
	"strings"

	goutils "github.com/simonski/goutils"
)

// AOC_2015_07 is the entrypoint
func AOC_2015_07(cli *goutils.CLI) {
	AOC_2015_07_part1_attempt1(cli)
	AOC_2015_07_part2_attempt1(cli)
}

func AOC_2015_07_part1_attempt1(cli *goutils.CLI) {
	circuit := NewCircuit(DAY_2015_07_DATA)
	circuit.Evaluate()

	for key, instruction := range circuit.instructions {
		fmt.Printf("[%v] %v: %v   %v\n", instruction.evaluated, key, instruction.value, instruction.line)
	}

	fmt.Printf("a: %v\n", circuit.GetSignalValue("a"))
}

func AOC_2015_07_part2_attempt1(cli *goutils.CLI) {

	circuit1 := NewCircuit(DAY_2015_07_DATA)
	circuit1.Evaluate()
	a := circuit1.GetInstruction("a")
	value := a.value

	circuit2 := NewCircuit(DAY_2015_07_DATA)
	i := circuit2.GetInstruction("b")
	i.evaluated = true
	i.value = value
	circuit2.Evaluate()

	for key, instruction := range circuit2.instructions {
		fmt.Printf("[%v] %v: %v   %v\n", instruction.evaluated, key, instruction.value, instruction.line)
	}

	fmt.Printf("a: %v\n", circuit2.GetSignalValue("a"))
}

type CircuitInstruction struct {
	line      string // x LSHIFT 2 -> f     NOT x -> h
	logic     string // x LSHIFT 2          NOT x
	wire      string // f                   h
	evaluated bool
	value     uint16
}

// Operation returns the bitwise operation NOT, AND OR LSHIFT RSHIFT
func (ci *CircuitInstruction) isNOT() bool {
	return strings.Index(ci.logic, "NOT") > -1
}

func (ci *CircuitInstruction) isLShift() bool {
	return strings.Index(ci.logic, "LSHIFT") > -1
}

func (ci *CircuitInstruction) isRShift() bool {
	return strings.Index(ci.logic, "RSHIFT") > -1
}

func (ci *CircuitInstruction) isAND() bool {
	return strings.Index(ci.logic, "AND") > -1
}

func (ci *CircuitInstruction) isOR() bool {
	return strings.Index(ci.logic, "OR") > -1
}

func (ci *CircuitInstruction) isLiteral() bool {
	return !ci.isNOT() && !ci.isLShift() && !ci.isRShift() && !ci.isAND() && !ci.isOR()
}

func NewCircuitInstruction(line string) *CircuitInstruction {
	splits := strings.Split(line, "->")
	logic := strings.TrimSpace(splits[0])
	wire := strings.TrimSpace(splits[1])
	ci := CircuitInstruction{line: line, logic: logic, wire: wire, evaluated: false}
	return &ci
}

func (i *CircuitInstruction) GetLeftRight() []string {
	// "x LSHIFT 2" >> [x, 2]
	// "4 LSHIFT x" >> [4, x]
	// "x AND y" >> [x, y]
	// "NOT xx" >> [xx]
	// "a or b" >> [a, b]
	l := strings.ReplaceAll(i.logic, "AND", "")
	l = strings.ReplaceAll(l, "OR", "")
	l = strings.ReplaceAll(l, "LSHIFT", "")
	l = strings.ReplaceAll(l, "RSHIFT", "")
	l = strings.ReplaceAll(l, "NOT", "")
	l = strings.ReplaceAll(l, "  ", " ")
	l = strings.TrimSpace(l)
	splits := strings.Split(l, " ")
	return splits
}

type Circuit struct {
	instructions map[string]*CircuitInstruction
}

func (c *Circuit) GetInstruction(key string) *CircuitInstruction {
	i, _ := c.instructions[key]
	return i
}

func (c *Circuit) Evaluate() {
	for {
		evaluated := 0
		for _, i := range c.instructions {
			if !i.evaluated {
				if i.isLiteral() {
					value, err := strconv.ParseInt(i.logic, 10, 16)
					if err == nil {
						i.value = uint16(value)
						i.evaluated = true
						evaluated++
						fmt.Printf("Literal: %v, value=%v, err=%v\n", i.line, value, err)
					} else {
						// it is an assigment of a gate
						instruction := c.GetInstruction(i.logic)
						if instruction != nil && instruction.evaluated {
							i.value = instruction.value
							i.evaluated = true
							evaluated++

						}
					}
				} else {

					keys := i.GetLeftRight()
					leftKey := keys[0]
					leftValue := uint16(0)
					leftGood := false
					if !isint(leftKey) {
						leftInstruction := c.GetInstruction(leftKey)
						if leftInstruction.evaluated {
							leftValue = leftInstruction.value
							leftGood = true
						}
					} else {
						leftValue = uint16(intify(leftKey))
						leftGood = true
					}

					rightValue := uint16(0)
					rightGood := false
					rightKey := ""
					if len(keys) == 2 {
						rightKey = keys[1]
						if !isint(rightKey) {
							rightInstruction := c.GetInstruction(rightKey)
							if rightInstruction.evaluated {
								rightValue = uint16(rightInstruction.value)
								rightGood = true
							}
						} else {
							rightValue = uint16(intify(rightKey))
							rightGood = true
						}
					}

					if i.isAND() && leftGood && rightGood {
						i.value = leftValue & rightValue
						i.evaluated = true
						evaluated++
					} else if i.isOR() && leftGood && rightGood {
						i.value = leftValue | rightValue
						i.evaluated = true
						evaluated++
						// } else if i.isNOT() && leftGood {
						// 	i.value = leftValue | rightValue
						// 	i.evaluated = true
						// 	evaluated++
					} else if i.isRShift() && leftGood && rightGood {
						i.value = leftValue >> rightValue
						i.evaluated = true
						evaluated++
					} else if i.isLShift() && leftGood && rightGood {
						i.value = leftValue << rightValue
						i.evaluated = true
						evaluated++
					} else if i.isNOT() && leftGood {
						i.value = uint16(bitwisenot(int(leftValue)))
						i.evaluated = true
						evaluated++
					}

				}

			}
		}
		if evaluated == 0 {
			break
		}
	}

}

func (c *Circuit) GetSignalValue(key string) uint16 {
	return c.instructions[key].value
}

func (c *Circuit) Size() int {
	return len(c.instructions)
}

func NewCircuit(instructions string) *Circuit {
	splits := strings.Split(instructions, "\n")
	m := make(map[string]*CircuitInstruction)
	for _, instruction := range splits {
		i := NewCircuitInstruction(instruction)
		m[i.wire] = i
	}
	circuit := Circuit{instructions: m}
	return &circuit
}
