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
	"strings"
)

// AOC_2015_07 is the entrypoint
func (app *Application) Y2015D07() {
	app.Y2015D07P1()
	// AOC_2015_06_part2_attempt1(cli)
}

func (app *Application) Y2015D07P1() {
	splits := strings.Split(DAY_2015_06_DATA, "\n")
	grid := NewLightGrid()
	for _, instruction := range splits {
		grid.Execute(instruction)
	}
	countOn, countOff := grid.CountOnOff()
	fmt.Printf("On %v Off %v\n", countOn, countOff)
}

type CircuitInstruction struct {
	line  string // x LSHIFT 2 -> f     NOT x -> h
	logic string // x LSHIFT 2          NOT x
	wire  string // f                   h
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
	logic := splits[0]
	wire := splits[1]
	ci := CircuitInstruction{line: line, logic: logic, wire: wire}
	return &ci
}

type Circuit struct {
	instructions []*CircuitInstruction
}

func NewCircuit(instructions string) *Circuit {
	splits := strings.Split(instructions, "\n")
	arr := make([]*CircuitInstruction, 0)
	for _, instruction := range splits {
		i := NewCircuitInstruction(instruction)
		arr = append(arr, i)
	}
	circuit := Circuit{instructions: arr}
	return &circuit
}
