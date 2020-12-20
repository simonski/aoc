package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func Test_AOC2020_14_BitSet(t *testing.T) {

	b := NewBitSet(1)
	s := b.ToBinaryString(8)
	if s != "00000001" {
		t.Error("Test_AOC2020_14_BitSet")
		t.Errorf("8-bit 1 no good, was %v, expected %v\n", s, "00000001")
	}

}
func Test_AOC2020_14_Program(t *testing.T) {

	instructions := `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`

	p := NewDay14ProgramFromStrings(strings.Split(instructions, "\n"))
	p.RunV1()
	sum := p.Sum()
	if sum != 165 {
		t.Error("Sum Error")
		t.Errorf("Test Sum() does not match, expected %v, got %v.", 165, sum)
	}
}

func Test_AOC2020_14_Mask(t *testing.T) {

	maskInput := "mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"
	p := NewDay14ProgramFromStrings(strings.Split(maskInput, "\n"))
	isMask := p.IsMask(maskInput)
	isMem := p.IsMem(maskInput)
	maskValue := p.ParseMask(maskInput)
	if maskValue.data != "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X" {
		t.Errorf("Mask value was incorrect.")
	}

	if !isMask {
		t.Errorf("Should be IsMask=True")
	}

	if isMem {
		t.Errorf("Should be IsMem=False")
	}

}

func Test_AOC2020_14_Mem(t *testing.T) {

	memInput := "mem[11] = 43"
	p := NewDay14ProgramFromStrings(strings.Split(memInput, "\n"))
	isMask := p.IsMask(memInput)
	isMem := p.IsMem(memInput)
	position, value := p.ParseMem(memInput)
	if position != "11" {
		t.Errorf("mem position was incorrect.")
	}
	if value != "43" {
		t.Errorf("mem value was incorrect.")
	}

	if isMask {
		t.Errorf("Should be IsMask=false")
	}

	if !isMem {
		t.Errorf("Should be IsMem=true")
	}

}

func Test_AOC2020_14_ParseMask(t *testing.T) {

	pattern1 := "mask = (?P<mask>.*)"
	input1 := "mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"
	expr := regexp.MustCompile(pattern1)
	names := expr.SubexpNames()
	result := expr.FindAllStringSubmatch(input1, -1)
	m := map[string]string{}
	for i, n := range result[0] {
		m[names[i]] = n
	}

	fmt.Printf("input: %v\n", input1)
	fmt.Printf("regex: %v\n", pattern1)
	fmt.Printf("names: %v\n", names)
	fmt.Printf("map: %v\n", m)

}

func Test_AOC2020_14_ParseMem(t *testing.T) {

	pattern1 := "mem\\[(?P<position>\\d+)\\] = (?P<mem>\\d+)"
	input1 := "mem[11] = 43"
	expr := regexp.MustCompile(pattern1)
	names := expr.SubexpNames()
	result := expr.FindAllStringSubmatch(input1, -1)
	m := map[string]string{}
	for i, n := range result[0] {
		m[names[i]] = n
	}
	fmt.Printf("values: %v\n", m)

}

func Test_AOC2020_14_Part2_Overall(t *testing.T) {

	program := `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`
	p := NewDay14ProgramFromStrings(strings.Split(program, "\n"))
	p.RunV2()
	sum := p.Sum()
	if sum != 208 {
		t.Error("Sum Error")
		t.Errorf("Test Sum() does not match, expected %v, got %v.", 165, sum)
	}

}

func Test_AOC2020_14_Part2_Section1(t *testing.T) {

	// 	program := `mask = 000000000000000000000000000000X1001X
	// mem[42] = 100
	// mask = 00000000000000000000000000000000X0XX
	// mem[26] = 1`
	p := NewDay14ProgramFromStrings(strings.Split("", "\n"))
	p.ExecuteV2("mask = 000000000000000000000000000000X1001X")
	expected := "000000000000000000000000000000X1001X"
	actual := p.Memory.GetMask()
	if actual.data != expected {
		t.Error("Mask is incorrect.")
	}

	p.ExecuteV2("mem[42] = 100")
	memPosition, _ := p.ParseMem("mem[42] = 100")
	ipos, _ := strconv.Atoi(memPosition)
	// ivalue, _ := strconv.Atoi(memValue)

	// convert the memory location to a bitset
	binaryMemoryAddress := decimal_to_binary(int64(ipos))
	if binaryMemoryAddress != "000000000000000000000000000000101010" {
		t.Errorf("binaryMemoryAddress should be 000000000000000000000000000000101010")
	}

	actualMask := p.Memory.GetMask()
	if actualMask.data != "000000000000000000000000000000X1001X" {
		t.Errorf("actualMask should be 000000000000000000000000000000X1001X")
	}

	// create a new mask from this as the current mask
	floatingMask := p.Memory.GetMask().DeriveNewMask(binaryMemoryAddress)
	if floatingMask.data != "000000000000000000000000000000X1101X" {
		t.Errorf("")
		t.Errorf("derivedMask is incorect")
		t.Errorf("original: %v\n", actualMask.data)
		t.Errorf("address:  %v\n", binaryMemoryAddress)
		t.Errorf("Expected: 000000000000000000000000000000X1101X")
		t.Errorf("Actual  : %v", floatingMask.data)
		t.Errorf("")
	}

	// Ok so we know that when we set the memoryValue we will calculate the floating mask correctly
	// now we need to verify that we will receive the correct list of masks to change; in this case
	// I want to check the recursive function is correct by having it generate the various addresses we *will* change
	// once that test works, then I can retain and adapt it to actually apply against all addresses and verify the sum
	addresses := floatingMask.GetVariations()
	if len(addresses) != 4 {
		t.Errorf("Incorrect number of addresses calculated, expected 4, got %v\n.", len(addresses))
		for address := range addresses {
			t.Errorf("%v\n", address)
		}
	}
	for index, address := range addresses {
		fmt.Printf("[%v] %v\n", index, address)
	}

	// realMasks := NewMaskV2(p.Memory.GetMask(), bitset_memPosition)

	// use this mask to calculate all memory positions to actually set.. and set them

	// bin_ipos := p.ConvertToBinary(memPosition)
	// bin_ivalue := p.ConvertToBinary(memValue)

	// fmt.Printf("memory position %v as binary  %v\n", ipos, bin_ipos)
	// fmt.Printf("memory value    %v as binary %v\n", ivalue, bin_ivalue)

	// sum := p.Sum()
	// if sum != 208 {
	// 	t.Error("Sum Error")
	// 	t.Errorf("Test Sum() does not match, expected %v, got %v.", 165, sum)
	// }

	// mask := m.GetMask()
	// memoryAddress.ApplyMask(mask)

}
