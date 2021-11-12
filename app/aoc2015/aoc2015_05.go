package aoc2015

/*
--- Day 5: Doesn't He Have Intern-Elves For This? ---
Santa needs help figuring out which strings in his text file are naughty or nice.

A nice string is one with all of the following properties:

It contains at least three vowels (aeiou only), like aei, xazegov, or aeiouaeiouaeiou.
It contains at least one letter that appears twice in a row, like xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).
It does not contain the strings ab, cd, pq, or xy, even if they are part of one of the other requirements.
For example:

ugknbfddgicrmopn is nice because it has at least three vowels (u...i...o...), a double letter (...dd...), and none of the disallowed substrings.
aaa is nice because it has at least three vowels and a double letter, even though the letters used by different rules overlap.
jchzalrnumimnmhp is naughty because it has no double letter.
haegwjzuvuyypxyu is naughty because it contains the string xy.
dvszwmarrgswjxmb is naughty because it contains only one vowel.
How many strings are nice?

*/

import (
	"fmt"
	"regexp"
	"strings"
)

// AOC_2015_05 is the entrypoint
func (app *Application) Y2015D05() {
	app.Y2015D05P1()
	app.Y2015D05P2()
}

func (app *Application) Y2015D05P1() {
	splits := strings.Split(DAY_2015_05_DATA, "\n")
	niceCount := 0
	naughtyCount := 0
	for _, line := range splits {
		if IsNice(line) {
			niceCount++
		} else {
			naughtyCount++
		}
	}
	fmt.Printf("Part 1 Nice count %v, naughty count %v\n", niceCount, naughtyCount)
}

func (app *Application) Y2015D05P2() {
	splits := strings.Split(DAY_2015_05_DATA, "\n")
	niceCount := 0
	naughtyCount := 0
	for _, line := range splits {
		if IsNice2(line) {
			niceCount++
		} else {
			naughtyCount++
		}
	}
	fmt.Printf("Part 2 Nice count %v, naughty count %v\n", niceCount, naughtyCount)
}

func IsNice(line string) bool {
	if strings.Index(line, "ab") > -1 || strings.Index(line, "cd") > -1 || strings.Index(line, "pq") > -1 || strings.Index(line, "xy") > -1 {
		return false
	}

	vowelless := strings.ReplaceAll(line, "a", "")
	vowelless = strings.ReplaceAll(vowelless, "e", "")
	vowelless = strings.ReplaceAll(vowelless, "i", "")
	vowelless = strings.ReplaceAll(vowelless, "o", "")
	vowelless = strings.ReplaceAll(vowelless, "u", "")

	diff := len(line) - len(vowelless)
	if diff < 3 {
		return false
	}

	doubleCount := 0
	for iletter := 97; iletter < 97+26; iletter++ {
		c := string(rune(iletter))
		c += c
		if strings.Index(line, c) > -1 {
			// fmt.Printf("%v in %v? : yes!\n", c, line)
			doubleCount++
			// contains a double
		} else {
			// fmt.Printf("%v in %v? : no!\n", c, line)

		}
	}

	if doubleCount > 0 {
		return true
	} else {
		return false
	}

}

func IsNice2(line string) bool {

	/*
			It contains a pair of any two letters that appears at least twice in the string without overlapping, like xyxy (xy) or aabcdefgaa (aa), but not like aaa (aa, but it overlaps).
		It contains at least one letter which repeats with exactly one letter between them, like xyx, abcdefeghi (efe), or even aaa.
	*/

	foundDouble := false
	for iletter1 := 97; iletter1 < 97+26; iletter1++ {
		l1 := string(rune(iletter1))
		for iletter2 := 97; iletter2 < 97+26; iletter2++ {
			l2 := string(rune(iletter2))
			l := l1 + l2
			s := strings.ReplaceAll(line, l, "")
			if len(line)-len(s) >= 4 {
				foundDouble = true
				break
			}
		}
		if foundDouble {
			break
		}
	}

	foundRegex := false
	for iletter1 := 97; iletter1 < 97+26; iletter1++ {
		l1 := string(rune(iletter1))
		pattern := l1 + "." + l1
		match, _ := regexp.MatchString(pattern, line)
		if match {
			foundRegex = true
			break
		}
	}

	if foundDouble && foundRegex {
		return true
	}
	return false

}
