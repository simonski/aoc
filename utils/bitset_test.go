package utils

import (
	"testing"
)

func Test_BitSet(t *testing.T) {
	verifyBitSet(0, "00000000", 8, t)
	verifyBitSet(1, "00000001", 8, t)
	verifyBitSet(2, "00000010", 8, t)
	verifyBitSet(4, "00000100", 8, t)
	verifyBitSet(8, "00001000", 8, t)
	verifyBitSet(16, "00010000", 8, t)
	verifyBitSet(32, "00100000", 8, t)
	verifyBitSet(64, "01000000", 8, t)
	verifyBitSet(128, "10000000", 8, t)
	verifyBitSet(255, "11111111", 8, t)
	verifyBitSet(256, "100000000", 9, t)
	verifyBitSet(512, "1000000000", 10, t)
	verifyBitSet(1024, "10000000000", 11, t)
	verifyBitSet(2048, "100000000000", 12, t)
}

func Test_BitSet2(t *testing.T) {
	verifyBits(0, 0, 0, 0, 0, 0, 0, 0, 0, "00000000", 8, t)
	verifyBits(1, 0, 0, 0, 0, 0, 0, 0, 1, "00000001", 8, t)
	verifyBits(2, 0, 0, 0, 0, 0, 0, 1, 0, "00000010", 8, t)
	verifyBits(4, 0, 0, 0, 0, 0, 1, 0, 0, "00000100", 8, t)
	verifyBits(8, 0, 0, 0, 0, 1, 0, 0, 0, "00001000", 8, t)
	verifyBits(16, 0, 0, 0, 1, 0, 0, 0, 0, "00010000", 8, t)
	verifyBits(32, 0, 0, 1, 0, 0, 0, 0, 0, "00100000", 8, t)
	verifyBits(64, 0, 1, 0, 0, 0, 0, 0, 0, "01000000", 8, t)
	verifyBits(128, 1, 0, 0, 0, 0, 0, 0, 0, "10000000", 8, t)
	verifyBits(255, 1, 1, 1, 1, 1, 1, 1, 1, "11111111", 8, t)

}

func verifyBits(value int64, bit7 int, bit6 int, bit5 int, bit4 int, bit3 int, bit2 int, bit1 int, bit0 int, expected string, length int, t *testing.T) {
	b := NewBitSet(value)
	actual := b.ToBinaryString(length)
	if expected != actual {
		t.Errorf("BitSet(%v) expected %v actual %v\n", value, expected, actual)
	}
	if b.Get(0) != bit0 {
		t.Errorf("BitSet(%v)[0] expected %v actual %v\n", actual, bit0, b.Get(0))
	}
	if b.Get(1) != bit1 {
		t.Errorf("BitSet(%v)[1] expected %v actual %v\n", actual, bit1, b.Get(1))
	}
	if b.Get(2) != bit2 {
		t.Errorf("BitSet(%v)[2] expected %v actual %v\n", actual, bit2, b.Get(2))
	}
	if b.Get(3) != bit3 {
		t.Errorf("BitSet(%v)[3] expected %v actual %v\n", actual, bit3, b.Get(3))
	}
	if b.Get(4) != bit4 {
		t.Errorf("BitSet(%v)[4] expected %v actual %v\n", actual, bit4, b.Get(4))
	}
	if b.Get(5) != bit5 {
		t.Errorf("BitSet(%v)[5] expected %v actual %v\n", actual, bit5, b.Get(5))
	}
	if b.Get(6) != bit6 {
		t.Errorf("BitSet(%v)[6] expected %v actual %v\n", actual, bit6, b.Get(6))
	}
	if b.Get(7) != bit7 {
		t.Errorf("BitSet(%v)[7] expected %v actual %v\n", actual, bit7, b.Get(7))
	}

}

func verifyBitSet(value int64, expected string, length int, t *testing.T) {
	b := NewBitSet(value)
	actual := b.ToBinaryString(length)
	if expected != actual {
		t.Errorf("BitSet(%v) expected %v actual %v\n", value, expected, actual)
	}

}

func Test_BitSetMask(t *testing.T) {
	// b := NewBitSet(11)
	verifyBitSet(11, "000000000000000000000000000000001011", 36, t)
	mask := NewMask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")
	b := NewBitSet(11)
	original := b.ToBinaryString(36)
	b.ApplyMask(mask)
	actual := b.ToBinaryString(36)
	expected := "000000000000000000000000000001001001"
	if actual != expected {
		t.Error("Mask not applied properly.")
		t.Errorf("Value    : %v\n", "000000000000000000000000000000001011")
		t.Errorf("Original : %v\n", original)
		t.Errorf("Mask     : %v\n", mask.data)
		t.Errorf("Expceted : %v\n", expected)
		t.Errorf("Actual   : %v\n", actual)
	}

}
