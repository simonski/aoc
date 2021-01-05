package main

import (
	"fmt"
	"testing"
)

func Test_AOC2020_23_1_Test(t *testing.T) {
	cc := NewCrabCups("389125467")
	if len(cc.Data) != 9 {
		t.Errorf("Test data should produce lenght 9 array.\n")
	}

	if cc.Data[0] != 3 {
		t.Errorf("data[0] should be 3, was %v.\n", cc.Data[0])
	}

	if cc.Data[1] != 8 {
		t.Errorf("data[1] should be 8, was %v.\n", cc.Data[1])
	}

	if cc.Data[2] != 9 {
		t.Errorf("data[2] should be 9, was %v.\n", cc.Data[2])
	}

	if cc.Data[3] != 1 {
		t.Errorf("data[3] should be 1, was %v.\n", cc.Data[3])
	}

	if cc.Data[4] != 2 {
		t.Errorf("data[4] should be 2, was %v.\n", cc.Data[4])
	}

	if cc.Data[5] != 5 {
		t.Errorf("data[5] should be 5, was %v.\n", cc.Data[5])
	}

	if cc.Data[6] != 4 {
		t.Errorf("data[6] should be 4, was %v.\n", cc.Data[6])
	}

	if cc.Data[7] != 6 {
		t.Errorf("data[7] should be 6, was %v.\n", cc.Data[7])
	}

	if cc.Data[8] != 7 {
		t.Errorf("data[8] should be 7, was %v.\n", cc.Data[8])
	}

	// cc.Reset()
	// result1 := cc.Play(1, true)
	// if result1 != "328915467" {
	// 	t.Errorf("CrabCups(1, 389125467) should give 328915467 but gave %v.\n", result1)
	// }

	// cc.Reset()
	// result2 := cc.Play(2, true)
	// if result2 != "325467891" {
	// 	t.Errorf("CrabCups(2) should give 325467891 but gave %v.\n", result2)
	// }

	// cc.Reset()
	// result3 := cc.Play(3, true)
	// if result3 != "725891346" {
	// 	t.Errorf("CrabCups(3) should give 725891346 but gave %v.\n", result3)
	// }

	cc.Reset()
	result10 := cc.Play(10, true)
	if result10 != "92658374" {
		t.Errorf("CrabCups(10) should give 92658374 but gave %v.\n", result10)
	} else {
		fmt.Printf("CrabCups(10) correctly gives %v\n", result10)
	}

	cc.Reset()
	result100 := cc.Play(100, true)
	if result100 != "67384529" {
		t.Errorf("CrabCups(100) should give 67384529 but gave %v.\n", result100)
	} else {
		fmt.Printf("CrabCups(100) correctly gives %v\n", result100)
	}

	ccReal := NewCrabCups("198753462")
	ccReal.Reset()
	result100Real := ccReal.Play(100, true)
	fmt.Printf("CrabCups(100) REAL gives %v\n", result100Real)

	// cc.Reset()
	// result100 := cc.Play(100, true)
	// if result100 != "92658374" {
	// 	t.Errorf("CrabCups(100) should give 67384529 but gave %v.\n", result100)
	// } else {
	// 	t.Errorf("CrabCups(100) correctly gives %v\n", result100)
	// }

	// not 67384529

	// cc.Reset()
	// result100 := cc.Play(100, true)
	// if result100 != "92658374" {
	// 	t.Errorf("CrabCups(10) should give 92658374 but gave %v.\n", result100)
	// }

	// cc.Reset()
	// result100 := cc.Play(100, true)
	// if result100 != "67384529" {
	// 	t.Errorf("CrabCups(100) should give 67384529 but gave %v.\n", result100)
	// }

}
