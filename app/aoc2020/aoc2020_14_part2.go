package aoc2020

import (
	"fmt"
	"strconv"

	goutils "github.com/simonski/goutils"
)

/*
--- Part Two ---
For some reason, the sea port's computer system still can't communicate with your ferry's docking program. It must be using version 2 of the decoder chip!

A version 2 decoder chip doesn't modify the values being written at all. Instead, it acts as a memory address decoder. Immediately before a value is written to memory, each bit in the bitmask modifies the corresponding bit of the destination memory address in the following way:

If the bitmask bit is 0, the corresponding memory address bit is unchanged.
If the bitmask bit is 1, the corresponding memory address bit is overwritten with 1.
If the bitmask bit is X, the corresponding memory address bit is floating.
A floating bit is not connected to anything and instead fluctuates unpredictably. In practice, this means the floating bits will take on all possible values, potentially causing many memory addresses to be written all at once!

For example, consider the following program:

mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1
When this program goes to write to memory address 42, it first applies the bitmask:

address: 000000000000000000000000000000101010  (decimal 42)
mask:    000000000000000000000000000000X1001X
result:  000000000000000000000000000000X1101X
After applying the mask, four bits are overwritten, three of which are different, and two of which are floating. Floating bits take on every possible combination of values; with two floating bits, four actual memory addresses are written:

000000000000000000000000000000011010  (decimal 26)
000000000000000000000000000000011011  (decimal 27)
000000000000000000000000000000111010  (decimal 58)
000000000000000000000000000000111011  (decimal 59)
Next, the program is about to write to memory address 26 with a different bitmask:

address: 000000000000000000000000000000011010  (decimal 26)
mask:    00000000000000000000000000000000X0XX
result:  00000000000000000000000000000001X0XX
This results in an address with three floating bits, causing writes to eight memory addresses:

000000000000000000000000000000010000  (decimal 16)
000000000000000000000000000000010001  (decimal 17)
000000000000000000000000000000010010  (decimal 18)
000000000000000000000000000000010011  (decimal 19)
000000000000000000000000000000011000  (decimal 24)
000000000000000000000000000000011001  (decimal 25)
000000000000000000000000000000011010  (decimal 26)
000000000000000000000000000000011011  (decimal 27)
The entire 36-bit address space still begins initialized to the value 0 at every address, and you still need the sum of all values left in memory at the end of the program. In this example, the sum is 208.

Execute the initialization program using an emulator for a version 2 decoder chip. What is the sum of all values left in memory after it completes?

*/

func (app *Application) Y2020D14P2() {
	filename := app.CLI.GetFileExistsOrDie("-input")
	p := NewDay14ProgramFromFilename(filename)
	p.RunV2()
	p.Debug()
	fmt.Printf("Total is %v\n", p.Sum())
}

func (m *Memory) SetV2(index int, value int64) {

	// convert the index (the memory position) to a bitset itself
	// memoryAddress := NewBitSet(int64(index))

	// now we mask the memoryAddress
	// the changes to memoryAddress give us the input to recurse on to find
	// the actual memory addresses to change
	// mask := m.GetMask()
	// memoryAddress.ApplyMask(mask)

	binaryMemoryAddress := goutils.Decimal_to_binary(int64(index))
	floatingMask := m.GetMask().DeriveNewMask(binaryMemoryAddress)
	addresses := floatingMask.GetVariations()
	for _, binaryAddress := range addresses {
		index := goutils.Binary_to_decimal(binaryAddress)
		b := m.Get(int(index))
		b.SetValue(value)
	}

	// b := m.Get(index)
	// b.SetValue(value)
	// b.ApplyMask(m.Mask)
}

func (p *Day14Program) ExecuteV2(instruction string) {
	fmt.Printf("ExecuteV2('%v')\n", instruction)
	if p.IsMask(instruction) {
		mask := p.ParseMask(instruction)
		p.Memory.SetMask(mask)
	} else if p.IsMem(instruction) {
		position, value := p.ParseMem(instruction)
		iposition, _ := strconv.Atoi(position)
		ivalue, _ := strconv.Atoi(value)
		p.Memory.SetV2(iposition, int64(ivalue))
	}
	fmt.Printf("\n")
}
