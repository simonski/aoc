package main

import (
	"testing"
)

/*
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
func Test_AOC2015_07_Part1(t *testing.T) {

	circuit := NewCircuit(DAY_2015_07_TEST_DATA)
	if circuit.Size() != 8 {
		t.Errorf("Circuit.Size() should be 8, was %v\n", circuit.Size())
	}

	instruction := circuit.GetInstruction("x")
	if instruction.line != "123 -> x" {
		t.Errorf("instruction[0] line should be '123 -> x', was '%v'\n", instruction.line)
	}
	if instruction.logic != "123" {
		t.Errorf("instruction[0] logic should be '123', was '%v'\n", instruction.logic)
	}
	if instruction.wire != "x" {
		t.Errorf("instruction[0] wire should be 'x', was '%v'\n", instruction.wire)
	}

	circuit.Evaluate()
	if circuit.GetSignalValue("d") != 72 {
		t.Errorf("Circuit.GetSignalValue(d) != 72, is %v\n", circuit.GetSignalValue("d"))
	}

	if circuit.GetSignalValue("e") != 507 {
		t.Errorf("Circuit.GetSignalValue(e) != 507, is %v\n", circuit.GetSignalValue("e"))
	}

	if circuit.GetSignalValue("f") != 492 {
		t.Errorf("Circuit.GetSignalValue(f) != 492, is %v\n", circuit.GetSignalValue("f"))
	}

	if circuit.GetSignalValue("g") != 114 {
		t.Errorf("Circuit.GetSignalValue(g) != 114, is %v\n", circuit.GetSignalValue("g"))
	}

	if circuit.GetSignalValue("h") != 65412 {
		t.Errorf("Circuit.GetSignalValue(h) != 65412, is %v\n", circuit.GetSignalValue("h"))
	}

	if circuit.GetSignalValue("i") != 65079 {
		t.Errorf("Circuit.GetSignalValue(i) != 65079, is %v\n", circuit.GetSignalValue("i"))
	}

	if circuit.GetSignalValue("x") != 123 {
		t.Errorf("Circuit.GetSignalValue(x) != 123, is %v\n", circuit.GetSignalValue("x"))
	}

	if circuit.GetSignalValue("y") != 456 {
		t.Errorf("Circuit.GetSignalValue(d) != 456, is %v\n", circuit.GetSignalValue("y"))
	}

	instruction = circuit.GetInstruction("f")
	// x LSHIFT 2 -> f

	if !instruction.isLShift() {
		t.Errorf("instruction[f] should be LShift\n")
	}

	if instruction.line != "x LSHIFT 2 -> f" {
		t.Errorf("instruction[f] line should be 'x LSHIFT 2 -> f', was '%v'\n", instruction.line)
	}

	if instruction.logic != "x LSHIFT 2" {
		t.Errorf("instruction[f].logic should be 'x LSHIFT 2', was '%v'\n", instruction.logic)
	}

	if instruction.wire != "f" {
		t.Errorf("instruction[f].wire should be 'f', was '%v'\n", instruction.wire)
	}

	leftright := instruction.GetLeftRight()
	if leftright[0] != "x" {
		t.Errorf("instruction[f].leftright[0] should be 'x', was '%v'\n", leftright[0])
	}

	if leftright[1] != "2" {
		t.Errorf("instruction[f].leftright[1] should be '2', was '%v'\n", leftright[1])
	}

	instruction = circuit.GetInstruction("h")
	// NOT x -> h

	if !instruction.isNOT() {
		t.Errorf("instruction[h] should be NOT\n")
	}

	if instruction.line != "NOT x -> h" {
		t.Errorf("instruction[h] line should be 'NOT x -> h', was '%v'\n", instruction.line)
	}

	if instruction.logic != "NOT x" {
		t.Errorf("instruction[h].logic should be 'NOT x', was '%v'\n", instruction.logic)
	}

	if instruction.wire != "h" {
		t.Errorf("instruction[h].wire should be 'h', was '%v'\n", instruction.wire)
	}

	leftright = instruction.GetLeftRight()
	if leftright[0] != "x" {
		t.Errorf("instruction[f].leftright[0] should be 'x', was '%v'\n", leftright[0])
	}

	if len(leftright) != 1 {
		t.Errorf("instruction[f].leftright should only be length 1', was '%v'\n", len(leftright))
	}

}
