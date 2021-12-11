package aoc2021

import (
	"fmt"
	"testing"
)

func Test_AOC2021_11_Part1(t *testing.T) {
	subject_number := 7
	card_public_key := DAY_25_CARD_PUBLIC_KEY
	door_public_key := DAY_25_DOOR_PUBLIC_KEY

	card_loop_size := FindLoopSize(subject_number, card_public_key)
	door_loop_size := FindLoopSize(subject_number, door_public_key)

	pk1 := FindPrivateKey(card_public_key, door_loop_size)
	pk2 := FindPrivateKey(door_public_key, card_loop_size)

	fmt.Printf("real PK is %v / %v\n", pk1, pk2)

}

func Test_AOC2021_11_Part2(t *testing.T) {
}
