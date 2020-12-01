package main

import (
	"fmt"

	goutils "github.com/simonski/goutils"
)

// AOC_2020_01 is the entrypoint to the various attempts for day one
func AOC_2020_01(cli *goutils.CLI) {

	AOC_2020_01_part1_attempt1(cli)
	AOC_2020_01_part1_attempt2(cli)

	AOC_2020_01_part2_attempt1(cli)

}

// AOC_2020_01_part2_attempt1 the second part of day 1, attempt1
// this is a brute-force attempt that gets over the line
// in the spirit of make it work, make it fast, this is make it work
// so this is 3 inner loops, giving o(n^3) performance I believe - it works, it is not fast.
func AOC_2020_01_part2_attempt1(cli *goutils.CLI) {

	// now we need to find 3 numbers that meet our total

	// so that is a case of

	// 1... 2020
	// a * b * c ! > maximum

	// make it work
	// make it right
	// make it fast

	// find two entries in input that sum to 2020
	// find combination that yields highest product

	// left to right
	// attempt 1: brute force attempt first entry and walk up to find entry that totals
	// retain maximum

	// then come back and do it properly
	fmt.Printf("\nPart2:\n")
	filename := cli.GetStringOrDie("-input")
	data := load_file_to_ints(filename)
	totalRequired := 2020
	maxSoFar := 0
	maxValue1 := 0
	maxValue2 := 0
	maxValue3 := 0
	oCount := 0
	for index1 := 0; index1 < len(data); index1++ {
		value1 := data[index1]
		for index2 := 0; index2 < len(data); index2++ {
			if index2 == index1 {
				continue
			}
			value2 := data[index2]
			if value1+value2 >= totalRequired {
				continue
			}
			for index3 := 0; index3 < len(data); index3++ {
				oCount++
				value3 := data[index3]
				if value1+value2+value3 == totalRequired {
					product := value1 * value2 * value3
					if product > maxSoFar {
						maxSoFar = product
						maxValue1 = value1
						maxValue2 = value2
						maxValue3 = value3
						fmt.Printf("Part2: New maximum: %v+%v+%v=%v, %v*%v*%v=%v\n", value1, value2, value3, value1+value2+value3, value1, value2, value3, value1*value2*value3)
					}
				}
			}

		}
	}

	fmt.Printf("Part2: Maximum: %v, (%v * %v * %v )\n", maxSoFar, maxValue1, maxValue2, maxValue3)
	fmt.Printf("Part2: o(n) is o(%v)=%v\n", len(data), oCount)

}

// AOC_2020_01_part1_attempt1 this is part 1 of day 1, attempt 1
// a brute-force attempt which as the volume is small works fine
// we have an inner loop giving o(n^2) which again works but is not fast
func AOC_2020_01_part1_attempt1(cli *goutils.CLI) {

	// make it work
	// make it right
	// make it fast

	// find two entries in input that sum to 2020
	// find combination that yields highest product

	// left to right
	// attempt 1: brute force attempt first entry and walk up to find entry that totals
	// retain maximum

	// then come back and do it properly

	fmt.Printf("\nPart1:\n")
	filename := cli.GetStringOrDie("-input")
	data := load_file_to_ints(filename)
	totalRequired := 2020
	maxSoFar := 0
	maxValue1 := 0
	maxValue2 := 0
	oCount := 0
	for index1 := 0; index1 < len(data); index1++ {
		value1 := data[index1]
		for index2 := 0; index2 < len(data); index2++ {
			oCount++
			if index2 == index1 {
				continue
			}

			value2 := data[index2]

			if value1+value2 == totalRequired {
				if value1*value2 > maxSoFar {
					maxSoFar = value1 * value2
					maxValue1 = value1
					maxValue2 = value2
					fmt.Printf("Part1: New maximum: %v+%v=%v, %v * %v=%v\n", value1, value2, value1+value2, value1, value2, value1*value2)
				}
			}

		}
	}

	fmt.Printf("Part1: Maximum: %v, (%v x %v)\n", maxSoFar, maxValue1, maxValue2)
	fmt.Printf("Part1: o(n^2)=%v\n", oCount)

}

// AOC_2020_01_part1_attempt2
// in this atempt I preload an "inty" map affording go's own probably binsearch by keying on the int value
// itself
// this uses more memory (the inty map in addition to the list) but avoids an initial sort and binsearch
// I intend to do my own sort and binsearch as an attempt3
func AOC_2020_01_part1_attempt2(cli *goutils.CLI) {

	// make it fast
	// so now I think sorting the numbers and doing a binary chop will give me o(log n) performance
	// as the first impl gave me my inner loop with is o(n^2) as I have to search everything; this way
	// I'll reduce my search space down somewhat

	// I think I'll start again in a loop but for each I'll workout my maximum value I can multiply with by
	// list = sorted(list)
	// for index in list:
	// 	entry = list[index]
	//  2020 / entry = ?   if value is integer, binsearch, else discard
	//	if found, retain if > max

	fmt.Printf("\nPart1:\n")
	filename := cli.GetStringOrDie("-input")
	data := load_file_to_ints(filename)

	mapints := make_map_of_inty_list(data)

	// don't need to binsearch if use an inbuild map

	totalRequired := 2020
	maxSoFar := 0
	maxValue1 := 0
	maxValue2 := 0
	oCount := 0
	for index := 0; index < len(data); index++ {
		// check - will there be an int availble?
		oCount++
		value := data[index]

		// we want searchFor exactly
		searchFor := totalRequired - value

		// otherwise it's an int. Do we have it?
		_, exists := mapints[searchFor]
		if exists {
			// yes, it exists.  These sum to our max
			product := value * searchFor
			if product > maxSoFar {
				maxSoFar = product
				maxValue1 = value
				maxValue2 = searchFor
				fmt.Printf("Part1: New maximum: %v+%v=%v, %v * %v=%v\n", maxValue1, maxValue2, maxValue1+maxValue2, maxValue1, maxValue2, maxValue1*maxValue2)
			}
		}
	}

	fmt.Printf("Part1: Maximum: %v, (%v x %v)\n", maxSoFar, maxValue1, maxValue2)
	fmt.Printf("Part1: o(n) is o(n log n)=%v\n", oCount)

}

// make_map_of_inty_list helper makes a map[int]int of a []int to give me
// whatever go's maps key algorithm performance is, at the cost of the memory
func make_map_of_inty_list(data []int) map[int]int {
	m := make(map[int]int)
	for index := 0; index < len(data); index++ {
		value := data[index]
		m[value] = value
	}
	return m
}
