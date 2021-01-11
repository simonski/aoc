package main

import (
	"fmt"
	"testing"
)

func Test_AOC2020_25_TestData(t *testing.T) {
	subject_number := 7
	card_public_key_expected := 5764801
	card_loop_size := FindLoopSize(subject_number, card_public_key_expected)
	// fmt.Printf("Loop size for %v is %v\n", target, loop_size)
	if card_loop_size != 8 {
		t.Errorf("FindLoopSize: card expected 8, got %v\n", card_loop_size)
	}

	door_public_key_expected := 17807724
	door_loop_size := FindLoopSize(subject_number, door_public_key_expected)
	if door_loop_size != 11 {
		t.Errorf("FindLoopSize: door expected 11 got %v\n", door_loop_size)
	}

	pk1 := FindPrivateKey(card_public_key_expected, door_loop_size)
	pk2 := FindPrivateKey(door_public_key_expected, card_loop_size)

	if pk1 != pk2 {
		t.Errorf("PK does not match, card: %v, door %v\n", pk1, pk2)
	}

}

func Test_AOC2020_25_Part1(t *testing.T) {
	subject_number := 7
	card_public_key := DAY_25_CARD_PUBLIC_KEY
	door_public_key := DAY_25_DOOR_PUBLIC_KEY

	card_loop_size := FindLoopSize(subject_number, card_public_key)
	door_loop_size := FindLoopSize(subject_number, door_public_key)

	pk1 := FindPrivateKey(card_public_key, door_loop_size)
	pk2 := FindPrivateKey(door_public_key, card_loop_size)

	fmt.Printf("real PK is %v / %v\n", pk1, pk2)

}
