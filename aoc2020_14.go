package main

/*

https://adventofcode.com/2020/day/14

--- Day 14: Docking Data ---
As your ferry approaches the sea port, the captain asks for your help again. The computer system that runs this port isn't compatible with the docking program on the ferry, so the docking parameters aren't being correctly initialized in the docking program's memory.

After a brief inspection, you discover that the sea port's computer system uses a strange bitmask system in its initialization program. Although you don't have the correct decoder chip handy, you can emulate it in software!

The initialization program (your puzzle input) can either update the bitmask or write a value to memory. Values and memory addresses are both 36-bit unsigned integers. For example, ignoring bitmasks for a moment, a line like mem[8] = 11 would write the value 11 to memory address 8.

The bitmask is always given as a string of 36 bits, written with the most significant bit (representing 2^35) on the left and the least significant bit (2^0, that is, the 1s bit) on the right. The current bitmask is applied to values immediately before they are written to memory: a 0 or 1 overwrites the corresponding bit in the value, while an X leaves the bit in the value unchanged.

For example, consider the following program:

mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0
This program starts by specifying a bitmask (mask = ....). The mask it specifies will overwrite two bits in every written value: the 2s bit is overwritten with 0, and the 64s bit is overwritten with 1.

The program then attempts to write the value 11 to memory address 8. By expanding everything out to individual bits, the mask is applied as follows:

value:  000000000000000000000000000000001011  (decimal 11)
mask:   XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
result: 000000000000000000000000000001001001  (decimal 73)
So, because of the mask, the value 73 is written to memory address 8 instead. Then, the program tries to write 101 to address 7:

value:  000000000000000000000000000001100101  (decimal 101)
mask:   XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
result: 000000000000000000000000000001100101  (decimal 101)
This time, the mask has no effect, as the bits it overwrote were already the values the mask tried to set. Finally, the program tries to write 0 to address 8:

value:  000000000000000000000000000000000000  (decimal 0)
mask:   XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
result: 000000000000000000000000000001000000  (decimal 64)
64 is written to address 8 instead, overwriting the value that was there previously.

To initialize your ferry's docking program, you need the sum of all values left in memory after the initialization program completes. (The entire 36-bit address space begins initialized to the value 0 at every address.) In the above example, only two values in memory are not zero - 101 (at address 7) and 64 (at address 8) - producing a sum of 165.

Execute the initialization program. What is the sum of all values left in memory after it completes? (Do not truncate the sum to 36 bits.)

*/
import (
	"fmt"
	"regexp"
	"strconv"

	goutils "github.com/simonski/goutils"
)

// AOC_2020_14 is the entrypoint
func AOC_2020_14(cli *goutils.CLI) {
	AOC_2020_14_part1_attempt1(cli)
	AOC_2020_14_part2_attempt1(cli)
}

func AOC_2020_14_part1_attempt1(cli *goutils.CLI) {
	filename := cli.GetFileExistsOrDie("-input")
	p := NewDay14ProgramFromFilename(filename)
	p.RunV1()
	p.Debug()
	fmt.Printf("Total is %v\n", p.Sum())
}

type Memory struct {
	data map[int]*BitSet
	Mask *Mask
}

func (m *Memory) Get(index int) *BitSet {
	b, exists := m.data[index]
	if exists {
		return b
	} else {
		b := NewBitSet(0)
		m.data[index] = b
		return b
	}
}

func (m *Memory) SetMask(mask *Mask) {
	m.Mask = mask
}

func (m *Memory) GetMask() *Mask {
	return m.Mask
}

func (m *Memory) Set(index int, value int64) {
	b := m.Get(index)
	b.SetValue(value)
	b.ApplyMask(m.Mask)
}

func (m *Memory) Debug() {
	for key, value := range m.data {
		fmt.Printf("[%v] %v\n", key, value.ToBinaryString(36))
	}
}

func (m *Memory) Sum() int64 {
	total := int64(0)
	for _, value := range m.data {
		total += value.GetValue()
	}
	return total
}

func NewMemory() *Memory {
	data := make(map[int]*BitSet)
	return &Memory{data: data}
}

type Day14Program struct {
	Instructions []string
	Position     int
	Memory       *Memory
}

func NewDay14ProgramFromFilename(filename string) *Day14Program {
	instructions := load_file_to_strings(filename)
	return NewDay14ProgramFromStrings(instructions)
}

func NewDay14ProgramFromStrings(instructions []string) *Day14Program {
	m := NewMemory()
	return &Day14Program{Instructions: instructions, Position: 0, Memory: m}
}

func (p *Day14Program) RunV1() {
	for _, line := range p.Instructions {
		p.ExecuteV1(line)
	}
}

func (p *Day14Program) RunV2() {
	for _, line := range p.Instructions {
		p.ExecuteV2(line)
	}
}

func (p *Day14Program) Debug() {
	for index, line := range p.Instructions {
		fmt.Printf("%v  [%v] \n", line, index)
	}
}

func (p *Day14Program) ExecuteV1(instruction string) {
	fmt.Printf("ExecuteV1('%v')\n", instruction)
	if p.IsMask(instruction) {
		mask := p.ParseMask(instruction)
		p.Memory.SetMask(mask)
	} else if p.IsMem(instruction) {
		position, value := p.ParseMem(instruction)
		iposition, _ := strconv.Atoi(position)
		ivalue, _ := strconv.Atoi(value)
		p.Memory.Set(iposition, int64(ivalue))
	}
	fmt.Printf("\n")
}

func (p *Day14Program) Sum() int64 {
	return p.Memory.Sum()
}

func (p *Day14Program) IsMask(instruction string) bool {
	pattern := "mask = (?P<mask>.*)"
	match, _ := regexp.MatchString(pattern, instruction)
	return match
}

func (p *Day14Program) IsMem(instruction string) bool {
	pattern := "mem\\[(?P<position>\\d+)\\] = (?P<mem>\\d+)"
	match, _ := regexp.MatchString(pattern, instruction)
	return match
}

func (p *Day14Program) ParseMem(instruction string) (string, string) {
	pattern1 := "mem\\[(?P<position>\\d+)\\] = (?P<mem>\\d+)"
	expr := regexp.MustCompile(pattern1)
	names := expr.SubexpNames()
	result := expr.FindAllStringSubmatch(instruction, -1)
	m := map[string]string{}
	for i, n := range result[0] {
		m[names[i]] = n
	}
	return m["position"], m["mem"]
}

func (p *Day14Program) ParseMask(instruction string) *Mask {
	pattern1 := "mask = (?P<mask>.*)"
	expr := regexp.MustCompile(pattern1)
	names := expr.SubexpNames()
	result := expr.FindAllStringSubmatch(instruction, -1)
	m := map[string]string{}
	for i, n := range result[0] {
		m[names[i]] = n
	}
	return NewMask(m["mask"])
}

func (p *Day14Program) ConvertToBinary(value string) string {
	ivalue, _ := strconv.Atoi(value)
	b := NewBitSet(int64(ivalue))
	return b.ToBinaryString(36)
}

func (p *Day14Program) ConvertToBitSet(value string) *BitSet {
	ivalue, _ := strconv.Atoi(value)
	b := NewBitSet(int64(ivalue))
	return b
}

func Regex(value string, pattern string) map[string]string {
	myExp := regexp.MustCompile(pattern)
	match := myExp.FindStringSubmatch(value)
	result := make(map[string]string)
	for i, name := range myExp.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	return result
}
