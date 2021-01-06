package main

/*
 */

import (
	"fmt"
	"strconv"
)

type CrabCups2 struct {
	OriginalData string
	Data         []int
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

func (cc *CrabCups2) Play(Rounds int, DEBUG bool) {
	d := cc.Data
	currentCup := d[0]

	lowestCup := 1
	highestCup := 999999

	n := make([]int, 1000000)

	fmt.Printf("\n%v Round(s)\n", Rounds)
	for Round := 0; Round < Rounds; Round++ {
		if Round%100 == 0 {
			fmt.Printf("\n-- move %v --\n", Round+1)
		}

		// fmt.Printf("First 20 (Pre Shuffle)  : '%v'\n", d[0:20])
		d = Shuffle(currentCup, Round, d)
		// line := DebugLine(currentCup, d)
		// fmt.Printf("First 20 (Post Shuffle) : '%v'\n", d[0:20])
		index := IndexOf(currentCup, d)

		firstCupIndex := (index + 1) % 1000000
		secondCupIndex := (index + 2) % 1000000
		thirdCupIndex := (index + 3) % 1000000
		firstCup := d[firstCupIndex]
		secondCup := d[secondCupIndex]
		thirdCup := d[thirdCupIndex]
		remainder := make([]int, 0)
		for index2 := 0; index2 < 1000000; index2++ {
			if d[index2] == firstCup || d[index2] == secondCup || d[index2] == thirdCup {
				// fmt.Printf("d[%v]=%v\n", index2, d[index2])
			} else {
				remainder = append(remainder, d[index2])
			}
		}

		// fmt.Printf("Remainder length=%v\n", len(remainder))

		// fmt.Printf("(Round %v/%v d=%v, 3cups=%v,%v,%v remainder=%v\n", round, index, d, firstCup, secondCup, thirdCup, remainder)

		// now threeCups holds the current set
		// remainder holds everything else
		// remainder := d[4:]
		destinationCup := currentCup - 1
		if destinationCup == 0 {
			destinationCup = 999999
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
				destinationIndex = (x + 1) % 1000000
				break
			}
		}

		// if DEBUG {
		// 	fmt.Printf("cups: %v\n", line)
		// 	fmt.Printf("pick up: %v, %v, %v\n", firstCup, secondCup, thirdCup)
		// 	fmt.Printf("destination: %v\n", destinationCup)
		// }

		// fmt.Printf("Destination cup is %v, index is %v\n", destinationCup, destinationIndex)
		// fmt.Printf("(Round %v/[%v] - %v), currentCup: '%v', remainder: '%v'\n", round, index, d, currentCup, remainder)
		// fmt.Printf("taking left from 0:%v\n", destinationIndex)
		// fmt.Printf("taking right from %v:%v\n", destinationIndex, len(remainder))
		// fmt.Printf("remainder length=%v\n", len(remainder))

		leftPart := remainder[0:destinationIndex]
		rightPart := remainder[destinationIndex:len(remainder)]

		counter := 0
		for pos := 0; pos < len(leftPart); pos++ {
			n[pos] = leftPart[pos]
			counter++
		}
		n[len(leftPart)+1] = firstCup
		n[len(leftPart)+2] = secondCup
		n[len(leftPart)+3] = thirdCup
		counter += 3

		for pos := 0; pos < len(rightPart); pos++ {
			n[pos+counter] = rightPart[pos]
		}

		// fmt.Printf("len(n)=%v\n", len(n))
		// n = append(n, leftPart...)
		// // fmt.Printf("(after append leftPart (which itself is len '%v') len(n)=%v\n", len(leftPart), len(n))
		// n = append(n, firstCup)
		// // fmt.Printf("(after append firstCup '%v' ) len(n)=%v\n", firstCup, len(n))
		// n = append(n, secondCup)
		// // fmt.Printf("(after append second cup '%v' ) len(n)=%v\n", secondCup, len(n))
		// n = append(n, thirdCup)
		// fmt.Printf("(After append thirdCup '%v' ) len(n)=%v\n", thirdCup, len(n))
		// n = append(n, rightPart...)
		// fmt.Printf("(after append rightpart (which itself is len %v) total len(n)=%v\n", len(rightPart), len(n))

		// fmt.Printf("Remainder length=%v\n", len(remainder))
		// fmt.Printf("Remainder(Left) length=%v\n", len(leftPart))
		// fmt.Printf("Remainder(Right) length=%v\n", len(rightPart))

		// fmt.Printf("(Round %v/[%v] - %v), currentCup: '%v', threeCups '%v,%v,%v', remainder: '%v', leftPart: '%v' rightPart: '%v', new data will be: %v\n",
		// round, index, d, currentCup, firstCup, secondCup, thirdCup, remainder, leftPart, rightPart, n)
		// d = append(leftPart, threeCups, rightPart)

		// fmt.Printf("len(d)=%v, len(n)=%v\n", len(d), len(n))
		d = n

		index = IndexOf(currentCup, d) + 1
		if index >= len(d) {
			index = 0
		}
		currentCup = d[index]

	}
	cc.Data = d

	// s := cc.StringValueAfter("1")
	// fmt.Printf("y")
	// return s
}

func NewCrabCups2(input string) *CrabCups2 {
	cc := CrabCups2{OriginalData: input}
	cc.Reset()
	return &cc
}

func (cc *CrabCups2) Reset() {
	MAX_SIZE := 1000000
	arr := make([]int, MAX_SIZE)
	input := cc.OriginalData
	// fmt.Printf("Reset2()\n")
	for index := 0; index < len(input); index++ {
		sval := input[index : index+1]
		ival, _ := strconv.Atoi(sval)
		arr[index] = ival
		// fmt.Printf("arr[%v]=%v\n", index, ival)
	}
	// fmt.Printf("\n")

	inputSize := len(input)
	for index := inputSize; index < MAX_SIZE; index++ {
		arr[index] = index + 1
		// fmt.Printf("arr[%v]=%v\n", index, arr[index])
	}
	cc.Data = arr
}
