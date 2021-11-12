package aoc2020

import (
	"fmt"
	"strconv"
	"strings"

	goutils "github.com/simonski/goutils"
)

func (app *Application) Y2020D02P1() {
	verbose := app.CLI.IndexOf("-v") > -1
	AOC_2020_02_part1_attempt1(app.CLI, verbose)
}

func (app *Application) Y2020D02P2() {
	verbose := app.CLI.IndexOf("-v") > -1
	AOC_2020_02_part2_attempt1(app.CLI, verbose)
}

// AOC_2020_02 is the entrypoint to the various attempts for day two
func AOC_2020_02(cli *goutils.CLI) {
	verbose := cli.IndexOf("-v") > -1
	AOC_2020_02_part1_attempt1(cli, verbose)
	AOC_2020_02_part2_attempt1(cli, verbose)
}

// AOC_2020_01_part1_attempt1 this is part 1 of day 1, attempt 1
// a brute-force attempt which as the volume is small works fine
// we have an inner loop giving o(n^2) which again works but is not fast
func AOC_2020_02_part1_attempt1(cli *goutils.CLI, verbose bool) {

	// make it work
	// make it right
	// make it fast

	// find two entries in input that sum to 2020
	// find combination that yields highest product

	// left to right
	// attempt 1: brute force attempt first entry and walk up to find entry that totals
	// retain maximum

	// then come back and do it properly

	filename := cli.GetStringOrDie("-input")
	data := goutils.Load_file_to_strings(filename)
	validCount := 0
	for index := 0; index < len(data); index++ {
		line := data[index]
		if isValidPart1(line, verbose) {
			validCount++
		}
	}

	fmt.Printf("AOC_2020_02_part1_attempt1: %v valid passwords.\n", validCount)

}

func AOC_2020_02_part2_attempt1(cli *goutils.CLI, verbose bool) {

	// make it work
	// make it right
	// make it fast

	// find two entries in input that sum to 2020
	// find combination that yields highest product

	// left to right
	// attempt 1: brute force attempt first entry and walk up to find entry that totals
	// retain maximum

	// then come back and do it properly

	filename := cli.GetStringOrDie("-input")
	data := goutils.Load_file_to_strings(filename)
	validCount := 0
	for index := 0; index < len(data); index++ {
		line := data[index]
		if isValidPart2(line, verbose) {
			validCount++
		}
	}

	fmt.Printf("AOC_2020_02_part1_attempt2: %v valid passwords.\n", validCount)

}

// parseLine returns the min, max letter and word to search
func parseLine(line string) (int, int, string, string) {
	results := strings.Split(line, ": ") // [ "7-11 w", "hwlwwwqwcwrwwwww" ]
	left := results[0]                   // "7-11 w"
	word := results[1]                   // "hwlwwwqwcwrwwwww"

	results = strings.Split(left, " ")               // [ "7-11", "w" ]
	minMaxValue := results[0]                        // "7-11"
	letter := results[1]                             // "w"
	minMaxResults := strings.Split(minMaxValue, "-") // [ "7", "11" ]
	minValue, _ := strconv.Atoi(minMaxResults[0])
	maxValue, _ := strconv.Atoi(minMaxResults[1])

	return minValue, maxValue, letter, word

}
func isValidPart1(line string, VERBOSE bool) bool {
	// 7-11 w: hwlwwwqwcwrwwwww
	minValue, maxValue, letter, word := parseLine(line)
	// results := strings.Split(line, ": ") // [ "7-11 w", "hwlwwwqwcwrwwwww" ]
	// left := results[0]                   // "7-11 w"
	// value := results[1]                  // "hwlwwwqwcwrwwwww"

	// results = strings.Split(left, " ")               // [ "7-11", "w" ]
	// minMaxValue := results[0]                        // "7-11"
	// letter := results[1]                             // "w"
	// minMaxResults := strings.Split(minMaxValue, "-") // [ "7", "11" ]
	// minValue, _ := strconv.Atoi(minMaxResults[0])
	// maxValue, _ := strconv.Atoi(minMaxResults[1])

	newValue := strings.ReplaceAll(word, letter, "")
	diff := len(word) - len(newValue)
	if diff >= minValue && diff <= maxValue {
		if VERBOSE {
			fmt.Printf("TRUE actual=%v, line='%v', min=%v, max=%v, letter=%v, value=%v\n", diff, line, minValue, maxValue, letter, word)
		}
		return true
	}

	if VERBOSE {
		fmt.Printf("FALSE actual=%v, line='%v', min=%v, max=%v, letter=%v, value=%v\n", diff, line, minValue, maxValue, letter, word)
	}
	return false
}

func isValidPart2(line string, VERBOSE bool) bool {
	// 7-11 w: hwlwwwqwcwrwwwww
	// Actually means at least 7 or 11 must contain w but not both
	index1, index2, letter, word := parseLine(line)

	if len(word) < index1 || len(word) < index2 {
		return false
	}
	letter1 := word[index1-1 : index1]
	letter2 := word[index2-1 : index2]

	isLetter1 := letter1 == letter
	isLetter2 := letter2 == letter

	if (isLetter1 && !isLetter2) || (!isLetter1 && isLetter2) {
		if VERBOSE {
			fmt.Printf("TRUE letter1=%v, letter2=%v, letter=%v, value=%v\n", letter1, letter2, letter, word)
		}
		return true
	} else {
		if VERBOSE {
			fmt.Printf("FALSE letter1=%v, letter2=%v, letter=%v, value=%v\n", letter1, letter2, letter, word)
		}
		return false
	}

}
