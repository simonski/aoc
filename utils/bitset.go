package utils

import (
	"fmt"
	"math"
	"strconv"
)

type BitSet struct {
	Bits   map[int]int
	Length int
}

func (b *BitSet) Get(index int) int {
	value, exists := b.Bits[index]
	if exists {
		return value
	} else {
		return 0
	}
}

func (b *BitSet) SetValue(value int64) {
	b.Clear()
	binary := strconv.FormatInt(value, 2)
	for index := 0; index < len(binary); index++ {
		bitindex := len(binary) - index
		bitvalue := binary[bitindex-1 : bitindex]
		ivalue, _ := strconv.Atoi(bitvalue)
		b.Bits[index] = ivalue
		if index > b.Length {
			b.Length = index
		}
	}
}

func (b *BitSet) ApplyMask(mask *Mask) {
	fmt.Printf("BitSet.ApplyMask()\n")
	fmt.Printf("Original  %v\n", b.ToBinaryString(36))
	fmt.Printf("ApplyMask %v\n", mask.Data)
	for index := 0; index < len(mask.Data); index++ {
		maskvalue := mask.Get(index)
		if maskvalue == "X" {
			// ingnore
		} else if maskvalue == "1" {
			b.Bits[index] = 1
		} else if maskvalue == "0" {
			b.Bits[index] = 0
		}
	}
	fmt.Printf("Modified  %v\n", b.ToBinaryString(36))
	fmt.Printf("Value     %v\n", b.GetValue())
}

func (b *BitSet) ToBinaryString(bits int) string {
	if b.Length > bits {
		bits = b.Length + 1
	}
	result := ""
	for index := bits - 1; index >= 0; index-- {
		if b.Get(index) == 0 {
			result += "0"
		} else {
			result += "1"
		}
	}
	return result
}

func (b *BitSet) GetValue() int64 {
	total := int64(0)
	for key, value := range b.Bits {
		if value == 1 {
			total += int64(math.Pow(2, float64(key)))
		}
	}
	return total
}

func (b *BitSet) Clear() {
	b.Bits = make(map[int]int)
}

func NewBitSet(value int64) *BitSet {
	bits := make(map[int]int)
	b := BitSet{Bits: bits}
	b.SetValue(value)
	return &b
}
