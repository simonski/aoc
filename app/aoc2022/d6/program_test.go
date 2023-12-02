package d6

import (
	"testing"
)

func Test_1(t *testing.T) {
	requireMarker("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7, t)
	requireMarker("bvwbjplbgvbhsrlpgdmjqwftvncz", 5, t)
	requireMarker("nppdvjthqldpwncqszvftbrmjlhg", 6, t)
	requireMarker("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10, t)
	requireMarker("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11, t)
}

func Test_14(t *testing.T) {
	requireMarker14("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19, t)
	requireMarker14("bvwbjplbgvbhsrlpgdmjqwftvncz", 23, t)
	requireMarker14("nppdvjthqldpwncqszvftbrmjlhg", 23, t)
	requireMarker14("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29, t)
	requireMarker14("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26, t)
}

func Test_Day1(t *testing.T) {
	requireMarker(REAL_DATA, 0, t)
}

func Test_Day2(t *testing.T) {
	requireMarker14(REAL_DATA, 0, t)
}

func requireMarker(input string, expected int, t *testing.T) {
	p := NewPuzzleWithData(input)
	actual := p.NextPacketMarker(4)
	if actual != expected {
		t.Fatalf("Input %v expected %v but was %v.\n", input, expected, actual)
	}
}

func requireMarker14(input string, expected int, t *testing.T) {
	p := NewPuzzleWithData(input)
	actual := p.NextPacketMarker(14)
	if actual != expected {
		t.Fatalf("Input %v expected %v but was %v.\n", input, expected, actual)
	}
}
