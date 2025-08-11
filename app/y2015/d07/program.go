package d07

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
	"github.com/simonski/cli"
	"github.com/simonski/goutils"
)

/*
Some Assembly Required
*/

type Puzzle struct {
	title     string
	year      int
	day       int
	input     string
	lines     []string
	HasRender bool
	HasWWW    bool
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("2015")
	iday, _ := strconv.Atoi("07")
	p := Puzzle{year: iyear, day: iday, title: "Some Assembly Required"}
	p.Load(input)
	return &p
}

func NewPuzzle() *Puzzle {
	return NewPuzzleWithData(REAL_DATA)
}

func (puzzle *Puzzle) Load(input string) {
	lines := strings.Split(input, "\n")
	puzzle.input = input
	puzzle.lines = lines
}

func (puzzle *Puzzle) GetSummary() utils.Summary {
	s := utils.Summary{Day: puzzle.day, Year: puzzle.year, Name: puzzle.title}
	s.ProgressP1 = utils.Completed
	s.ProgressP2 = utils.Completed
	s.DateStarted = "d07"
	return s
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}

func (puzzle *Puzzle) RunWWW() string {
	return ""
}

func (puzzle *Puzzle) Render(cli *cli.CLI) {
}

func (puzzle *Puzzle) Part1() {
	circuit := NewCircuit(REAL_DATA)
	circuit.Evaluate()

	for key, instruction := range circuit.instructions {
		fmt.Printf("[%v] %v: %v   %v\n", instruction.evaluated, key, instruction.value, instruction.line)
	}

	fmt.Printf("a: %v\n", circuit.GetSignalValue("a"))
}

func (puzzle *Puzzle) Part2() {

	circuit1 := NewCircuit(REAL_DATA)
	circuit1.Evaluate()
	a := circuit1.GetInstruction("a")
	value := a.value

	circuit2 := NewCircuit(REAL_DATA)
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
					if !goutils.Isint(leftKey) {
						leftInstruction := c.GetInstruction(leftKey)
						if leftInstruction.evaluated {
							leftValue = leftInstruction.value
							leftGood = true
						}
					} else {
						leftValue = uint16(goutils.Intify(leftKey))
						leftGood = true
					}

					rightValue := uint16(0)
					rightGood := false
					rightKey := ""
					if len(keys) == 2 {
						rightKey = keys[1]
						if !goutils.Isint(rightKey) {
							rightInstruction := c.GetInstruction(rightKey)
							if rightInstruction.evaluated {
								rightValue = uint16(rightInstruction.value)
								rightGood = true
							}
						} else {
							rightValue = uint16(goutils.Intify(rightKey))
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
						i.value = uint16(goutils.Bitwisenot(int(leftValue)))
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
