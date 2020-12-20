package main

import (
	"fmt"
	"regexp"
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
	p.Run()
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
