package main

/*
The small crab challenges you to a game! The crab is going to mix up some cups, and you have to predict where they'll end up.

The cups will be arranged in a circle and labeled clockwise (your puzzle input). For example, if your labeling were 32415, there would be five cups in the circle; going clockwise around the circle from the first cup, the cups would be labeled 3, 2, 4, 1, 5, and then back to 3 again.

Before the crab starts, it will designate the first cup in your list as the current cup. The crab is then going to do 100 moves.

Each move, the crab does the following actions:

The crab picks up the three cups that are immediately clockwise of the current cup. They are removed from the circle; cup spacing is adjusted as necessary to maintain the circle.
The crab selects a destination cup: the cup with a label equal to the current cup's label minus one. If this would select one of the cups that was just picked up, the crab will keep subtracting one until it finds a cup that wasn't just picked up. If at any point in this process the value goes below the lowest value on any cup's label, it wraps around to the highest value on any cup's label instead.
The crab places the cups it just picked up so that they are immediately clockwise of the destination cup. They keep the same order as when they were picked up.
The crab selects a new current cup: the cup which is immediately clockwise of the current cup.
For example, suppose your cup labeling were 389125467. If the crab were to do merely 10 moves, the following changes would occur:

-- move 1 --
cups: (3) 8  9  1  2  5  4  6  7
pick up: 8, 9, 1
destination: 2

-- move 2 --
cups:  3 (2) 8  9  1  5  4  6  7
pick up: 8, 9, 1
destination: 7

-- move 3 --
cups:  3  2 (5) 4  6  7  8  9  1
pick up: 4, 6, 7
destination: 3

-- move 4 --
cups:  7  2  5 (8) 9  1  3  4  6
pick up: 9, 1, 3
destination: 7

-- move 5 --
cups:  3  2  5  8 (4) 6  7  9  1
pick up: 6, 7, 9
destination: 3

-- move 6 --
cups:  9  2  5  8  4 (1) 3  6  7
pick up: 3, 6, 7
destination: 9

-- move 7 --
cups:  7  2  5  8  4  1 (9) 3  6
pick up: 3, 6, 7
destination: 8

-- move 8 --
cups:  8  3  6  7  4  1  9 (2) 5
pick up: 5, 8, 3
destination: 1

-- move 9 --
cups:  7  4  1  5  8  3  9  2 (6)
pick up: 7, 4, 1
destination: 5

-- move 10 --
cups: (5) 7  4  1  8  3  9  2  6
pick up: 7, 4, 1
destination: 3

-- final --
cups:  5 (8) 3  7  4  1  9  2  6
In the above example, the cups' values are the labels as they appear moving clockwise around the circle; the current cup is marked with ( ).

After the crab is done, what order will the cups be in? Starting after the cup labeled 1, collect the other cups' labels clockwise into a single string with no extra characters; each number except 1 should appear exactly once. In the above example, after 10 moves, the cups clockwise from 1 are labeled 9, 2, 6, 5, and so on, producing 92658374. If the crab were to complete all 100 moves, the order after cup 1 would be 67384529.

Using your labeling, simulate 100 moves. What are the labels on the cups after cup 1?
*/

import (
	"fmt"
	"strconv"
	"time"

	goutils "github.com/simonski/goutils"
)

// AOC_2020_20 is the entrypoint
func AOC_2020_23(cli *goutils.CLI) {
	AOC_2020_23_part2_attempt1(cli)
}

func AOC_2020_23_part1_attempt1(cli *goutils.CLI) {
}

func AOC_2020_23_part2_attempt1(cli *goutils.CLI) {
	start := time.Now()
	input := "198753462"
	SIZE := 1000000
	ROUNDS := SIZE * 10
	data := make([]int, SIZE)
	for index := 0; index < len(input); index++ {
		sval := input[index : index+1]
		ival, _ := strconv.Atoi(sval)
		data[index] = ival
	}
	inputSize := len(input)
	for index := inputSize; index < SIZE; index++ {
		data[index] = index + 1
	}

	ring := NewRing(data)
	DEBUG := false
	ring.Play(ROUNDS, DEBUG)
	cup1 := ring.Find(1)
	r1 := cup1.Next
	r2 := cup1.Next.Next
	fmt.Printf("r1.Value=%v, r2.Value=%v, %v x %v = %v\n", r1.Value, r2.Value, r1.Value, r2.Value, r1.Value*r2.Value)
	end := time.Now()
	fmt.Printf("%v\n", end.Sub(start))

}

type CrabCups struct {
	OriginalData string
	Data         []int
}

func (cc *CrabCups) Play(Rounds int, DEBUG bool) string {
	return cc.Play1(Rounds, DEBUG)
}

// func (cc *CrabCups) Play2(Rounds int, DEBUG bool) string {
// 	d := cc.Data
// 	currentCupValue := d[currentCupIndex]
// 	// take off the 3 cups
// 	for index, value := range d {
// 		// I want currentCupIndex+1, +2, +3 as my 3cups
// 		if index == (currentCupIndex+1)%len(d) {
// 			threeCups = append(threeCups, d[index])
// 		} else {
// 			remainder = append(remainder, d[index])
// 		}
// 	}
// }

func IndexOf(value int, data []int) int {
	for index := 0; index < len(data); index++ {
		if data[index] == value {
			return index
		}
	}
	return -1
}

func DebugLine(currentCup int, data []int) string {
	line := ""
	for _, v := range data {
		if v == currentCup {
			line += fmt.Sprintf("(%v) ", v)
		} else {
			line += fmt.Sprintf("%v ", v)
		}
	}
	return line
}

func Shuffle(currentCup int, newIndexRequired int, data []int) []int {
	index := IndexOf(currentCup, data)
	offset := newIndexRequired - index
	d := make([]int, len(data))
	for index := 0; index < len(data); index++ {
		new_index := (index + offset) % len(data)
		if index+offset < 0 {
			new_index = len(data) + (index + offset)
		}
		// fmt.Printf("Shuffle: index=%v, offset=%v, newIndex=%v len(data)=%v\n", index, offset, new_index, len(data))
		d[new_index] = data[index]
	}
	// fmt.Printf("Shuffle: currentCup=%v, indexOfCurrentCup=%v, newIndex=%v, offset=%v, original=%v, new=%v\n", currentCup, index, index+offset, offset, data, d)
	return d
}

func (cc *CrabCups) Play1(Rounds int, DEBUG bool) string {
	d := cc.Data
	currentCup := d[0]

	lowestCup := 9
	highestCup := 0
	for _, value := range d {
		lowestCup = Min(lowestCup, value)
		highestCup = Max(highestCup, value)
	}

	fmt.Printf("\n%v Round(s)\n", Rounds)
	for Round := 0; Round < Rounds; Round++ {
		fmt.Printf("\n-- move %v --\n", Round+1)
		d = Shuffle(currentCup, Round, d)
		line := DebugLine(currentCup, d)
		index := IndexOf(currentCup, d)

		firstCupIndex := (index + 1) % len(d)
		secondCupIndex := (index + 2) % len(d)
		thirdCupIndex := (index + 3) % len(d)
		firstCup := d[firstCupIndex]
		secondCup := d[secondCupIndex]
		thirdCup := d[thirdCupIndex]
		remainder := make([]int, 0)
		for index2 := 0; index2 < len(d); index2++ {
			if d[index2] == firstCup || d[index2] == secondCup || d[index2] == thirdCup {
			} else {
				remainder = append(remainder, d[index2])
			}
		}
		// fmt.Printf("(Round %v/%v d=%v, 3cups=%v,%v,%v remainder=%v\n", round, index, d, firstCup, secondCup, thirdCup, remainder)

		// now threeCups holds the current set
		// remainder holds everything else
		// remainder := d[4:]
		destinationCup := currentCup - 1
		if destinationCup == 0 {
			destinationCup = 9
		}
		for {
			quit := true
			// fmt.Printf("Desintation cup %v ? \n", destinationCup)
			if destinationCup == firstCup || destinationCup == secondCup || destinationCup == thirdCup {
				// fmt.Printf("Desintation cup equals one of our 3 cups %v yes \n", destinationCup)
				destinationCup -= 1
				if destinationCup < lowestCup {
					// fmt.Printf("Desintation cup < lowestCup %v < %v yes, setting to %v\n ", destinationCup, lowestCup, highestCup)
					destinationCup = highestCup
				}
				quit = false
			}
			if quit {
				break
			}
		}
		destinationIndex := 0
		for x, value := range remainder {
			if value == destinationCup {
				destinationIndex = (x + 1) % len(d)
				break
			}
		}

		fmt.Printf("cups: %v\n", line)
		fmt.Printf("pick up: %v, %v, %v\n", firstCup, secondCup, thirdCup)
		fmt.Printf("destination: %v\n", destinationCup)

		// fmt.Printf("Dstination cup is %v, index is %v\n", destinationCup, destinationIndex)
		// fmt.Printf("(Round %v/[%v] - %v), currentCup: '%v', remainder: '%v'\n", round, index, d, currentCup, remainder)
		leftPart := remainder[0:destinationIndex]
		rightPart := remainder[destinationIndex:len(remainder)]
		n := make([]int, 0)
		n = append(n, leftPart...)
		n = append(n, firstCup)
		n = append(n, secondCup)
		n = append(n, thirdCup)
		n = append(n, rightPart...)
		// fmt.Printf("(Round %v/[%v] - %v), currentCup: '%v', threeCups '%v,%v,%v', remainder: '%v', leftPart: '%v' rightPart: '%v', new data will be: %v\n",
		// round, index, d, currentCup, firstCup, secondCup, thirdCup, remainder, leftPart, rightPart, n)
		// d = append(leftPart, threeCups, rightPart)

		d = n

		index = IndexOf(currentCup, d) + 1
		if index >= len(d) {
			index = 0
		}
		currentCup = d[index]

	}
	cc.Data = d
	return cc.StringValueAfter("1")
}

func (cc *CrabCups) StringValueAfter(after string) string {
	d := cc.Data
	offset := IndexOf(1, cc.Data)
	line := ""
	for index := 0; index < len(cc.Data); index++ {
		new_index := (index + offset) % len(d)
		if cc.Data[new_index] != 1 {
			line += fmt.Sprintf("%v", d[new_index])
		}
	}
	return line
}

func (cc *CrabCups) StringValue() string {
	d := cc.Data
	line := ""
	for index := 0; index < len(cc.Data); index++ {
		line += fmt.Sprintf("%v", d[index])
	}
	return line
}

func NewCrabCups(input string) *CrabCups {
	cc := CrabCups{OriginalData: input}
	cc.Reset()
	return &cc
}

func (cc *CrabCups) Reset() {
	arr := make([]int, 0)
	input := cc.OriginalData
	for index := 0; index < len(input); index++ {
		sval := input[index : index+1]
		ival, _ := strconv.Atoi(sval)
		arr = append(arr, ival)
	}
	cc.Data = arr
}
