package app

import (
	"strings"
	"testing"
)

func Test_AOC2020_08_NewInstruction(t *testing.T) {
	line := "nop +0"
	instruction := NewInstruction(line)
	expected := "nop"
	if instruction.Operation != expected {
		t.Errorf("Expected '%v', got '%v', line='%v'.\n", expected, instruction.Operation, line)
	}
	iexpected := 0
	if instruction.Argument != iexpected {
		t.Errorf("Expected '%v', got '%v', line='%v'.\n", expected, instruction.Argument, line)
	}
}

func Test_AOC2020_08_NewProgram(t *testing.T) {
	commands := `nop +0
	acc +4

	jmp -1
	acc +5
	nop +1
	`

	p := NewProgram(strings.Split(commands, "\n"))
	if p.Size() != 5 {
		t.Errorf("Expected '5', got '%v'.\n", p.Size())
	}

	i1 := p.GetInstructionAtIndex(0)
	checkInstruction("nop", 0, i1, t)
	i2 := p.GetInstructionAtIndex(1)
	checkInstruction("acc", 4, i2, t)
	i3 := p.GetInstructionAtIndex(2)
	checkInstruction("jmp", -1, i3, t)
	i4 := p.GetInstructionAtIndex(3)
	checkInstruction("acc", 5, i4, t)
	i5 := p.GetInstructionAtIndex(4)
	checkInstruction("nop", 1, i5, t)
}

func checkInstruction(operation string, argument int, instruction *Instruction, t *testing.T) {
	if instruction.Operation != operation {
		t.Errorf("Expected '%v', got '%v'.\n", operation, instruction.Operation)
	}
	if instruction.Argument != argument {
		t.Errorf("Expected '%v', got '%v'.\n", argument, instruction.Argument)
	}
}
