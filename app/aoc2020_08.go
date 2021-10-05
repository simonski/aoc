package app

/*
Day 08 - Handheld Halting
https://adventofcode.com/2020/day/8

Your flight to the major airline hub reaches cruising altitude without incident. While you consider checking the in-flight menu for one of those drinks that come with a little umbrella, you are interrupted by the kid sitting next to you.

Their handheld game console won't turn on! They ask if you can take a look.

You narrow the problem down to a strange infinite loop in the boot code (your puzzle input) of the device. You should be able to fix it, but first you need to be able to run the code in isolation.

The boot code is represented as a text file with one instruction per line of text. Each instruction consists of an operation (acc, jmp, or nop) and an argument (a signed number like +4 or -20).

acc increases or decreases a single global value called the accumulator by the value given in the argument. For example, acc +7 would increase the accumulator by 7. The accumulator starts at 0. After an acc instruction, the instruction immediately below it is executed next.
jmp jumps to a new instruction relative to itself. The next instruction to execute is found using the argument as an offset from the jmp instruction; for example, jmp +2 would skip the next instruction, jmp +1 would continue to the instruction immediately below it, and jmp -20 would cause the instruction 20 lines above to be executed next.
nop stands for No OPeration - it does nothing. The instruction immediately below it is executed next.
For example, consider the following program:

nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
These instructions are visited in this order:

nop +0  | 1
acc +1  | 2, 8(!)
jmp +4  | 3
acc +3  | 6
jmp -3  | 7
acc -99 |
acc +1  | 4
jmp -4  | 5
acc +6  |
First, the nop +0 does nothing. Then, the accumulator is increased from 0 to 1 (acc +1) and jmp +4 sets the next instruction to the other acc +1 near the bottom. After it increases the accumulator from 1 to 2, jmp -4 executes, setting the next instruction to the only acc +3. It sets the accumulator to 5, and jmp -3 causes the program to continue back at the first acc +1.

This is an infinite loop: with this sequence of jumps, the program will run forever. The moment the program tries to run any instruction a second time, you know it will never terminate.

Immediately before the program would run an instruction a second time, the value in the accumulator is 5.

Run your copy of the boot code. Immediately before any instruction is executed a second time, what value is in the accumulator?


*/
import (
	"fmt"
	"os"
	"strconv"
	"strings"

	goutils "github.com/simonski/goutils"
)

// AOC_2020_08 is the entrypoint
func AOC_2020_08(cli *goutils.CLI) {
	AOC_2020_08_part1_attempt1(cli)
	AOC_2020_08_part2_attempt1(cli)
}

func AOC_2020_08_part1_attempt1(cli *goutils.CLI) {
	filename := cli.GetFileExistsOrDie("-input")
	p := NewProgramFromFilename(filename)
	p.Debug()

	// Ok, step until Accumulator is 5
	for true {
		p.Step()
		instruction := p.GetCurrentInstruction()
		if instruction.ExecutionCount == 1 {
			// then it has run once.
			fmt.Printf("Accumulator is now %v, index is %v, instruction is %v\n", p.Accumulator, p.Index, instruction)
			break
		}
	}

}

func AOC_2020_08_part2_attempt1(cli *goutils.CLI) {
	filename := cli.GetFileExistsOrDie("-input")
	p := NewProgramFromFilename(filename)

	// one jmp is a nop
	// build list of jmps that exist and for each one run the test until complete or in loop
	jmps := p.FindInstructionIndexes("jmp")
	// okay so for each jmp, flip it to nop and run
	for index := range jmps {
		jmpIndex := jmps[index]
		// load the program, change the instruction
		testProgram := NewProgramFromFilename(filename)
		instruction := testProgram.GetInstructionAtIndex(jmpIndex)
		instruction.Operation = "nop"

		// Ok now run the program until we eithe rcomplete or move to a loop
		for true {
			if testProgram.Step() == false {
				// then we went to a loop
				fmt.Printf("jmp->nop at index %v causes a loop.\n", jmpIndex)
				break
			} else if testProgram.IsComplete() {
				// this is the one we want
				fmt.Printf("jmp->nop at index %v fixes our program, Accumulator is %v\n", jmpIndex, testProgram.Accumulator)
				os.Exit(0)
			}

		}
	}

	nops := p.FindInstructionIndexes("nop")
	// okay so for each jmp, flip it to nop and run
	for index := range nops {
		jmpIndex := nops[index]
		// load the program, change the instruction
		testProgram := NewProgramFromFilename(filename)
		instruction := testProgram.GetInstructionAtIndex(jmpIndex)
		instruction.Operation = "jmp"

		// Ok now run the program until we eithe rcomplete or move to a loop
		for true {
			if testProgram.Step() == false {
				// then we went to a loop
				fmt.Printf("nop->jmp at index %v causes a loop.\n", jmpIndex)
				break
			} else if testProgram.IsComplete() {
				// this is the one we want
				fmt.Printf("nop->jmp at index %v fixes our program, Accumulator is %v\n", jmpIndex, testProgram.Accumulator)
				os.Exit(0)
			}

		}
	}

}

type Program struct {
	Instructions []*Instruction
	Accumulator  int
	Index        int
	CurrentStep  int
}

// Returns the posiiton of all Instructions with the specified Operation type
func (p *Program) FindInstructionIndexes(operation string) []int {
	results := make([]int, 0)
	for index := range p.Instructions {
		instruction := p.GetInstructionAtIndex(index)
		if instruction.Operation == operation {
			results = append(results, index)
		}
	}
	return results
}

func (p *Program) Size() int {
	return len(p.Instructions)
}

// GetInstructionAtIndex returns the instruct at any index
func (p *Program) GetInstructionAtIndex(index int) *Instruction {
	return p.Instructions[index]
}

// GetCurrentInstruct returns the Instruction at the current Index
func (p *Program) GetCurrentInstruction() *Instruction {
	return p.Instructions[p.Index]
}

// Reset set the Program to the original state, Index and Accumulator to 0
func (p *Program) Reset() {
	p.Accumulator = 0
	p.Index = 0
	p.CurrentStep = 0
	for index := range p.Instructions {
		instruction := p.GetInstructionAtIndex(index)
		instruction.ExecutionCount = 0
		instruction.ExecutedOnStep = 0
	}
}

// Performs the current instruction, moving the index to the next valid value
// returns true if the instruction executes and continues normally
// returns false if the insruction executes and puts us into an infinite loop
func (p *Program) Step() bool {
	p.CurrentStep++
	instruction := p.GetCurrentInstruction()
	if instruction.Operation == "acc" {
		// increase of decrease the accumulator by the value
		p.Accumulator += instruction.Argument
		p.Index++
	} else if instruction.Operation == "jmp" {
		p.Index += instruction.Argument
	} else if instruction.Operation == "nop" {
		p.Index++
	}
	instruction.ExecutionCount++
	instruction.ExecutedOnStep = p.CurrentStep
	if instruction.ExecutionCount == 1 {
		return true
	}
	return false
}

// IsComplete indicates if the program has completed
func (p *Program) IsComplete() bool {
	return p.Index == len(p.Instructions)
}

func (p *Program) Debug() {
	for index := range p.Instructions {
		instruction := p.GetInstructionAtIndex(index)
		instruction.Debug()
	}
}

func NewProgram(lines []string) *Program {
	instructions := make([]*Instruction, 0)
	for index := range lines {
		line := lines[index]
		if strings.TrimSpace(line) != "" {
			i := NewInstruction(line)
			instructions = append(instructions, i)
		}
	}
	return &Program{Instructions: instructions, Accumulator: 0, Index: 0}
}

func NewProgramFromFilename(filename string) *Program {
	lines := goutils.Load_file_to_strings(filename)
	return NewProgram(lines)
}

type Instruction struct {
	Operation      string
	Argument       int
	ExecutionCount int
	ExecutedOnStep int
}

func (i *Instruction) Debug() {
	fmt.Printf("[%v] [%v] %v %v\n", i.ExecutionCount, i.ExecutedOnStep, i.Operation, i.Argument)
}

func NewInstruction(line string) *Instruction {
	line = strings.TrimSpace(line)
	splits := strings.Split(line, " ")
	operation := splits[0]
	argument, _ := strconv.Atoi(splits[1])
	i := Instruction{Operation: operation, Argument: argument, ExecutionCount: 0}
	return &i
}
