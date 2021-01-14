package main

/*
Space on the sleigh is limited this year, and so Santa will be bringing his list as a digital copy. He needs to know how much space it will take up when stored.

It is common in many programming languages to provide a way to escape special characters in strings. For example, C, JavaScript, Perl, Python, and even PHP handle special characters in very similar ways.

However, it is important to realize the difference between the number of characters in the code representation of the string literal and the number of characters in the in-memory string itself.

For example:

"" is 2 characters of code (the two double quotes), but the string contains zero characters.
"abc" is 5 characters of code, but 3 characters in the string data.
"aaa\"aaa" is 10 characters of code, but the string itself contains six "a" characters and a single, escaped quote character, for a total of 7 characters in the string data.
"\x27" is 6 characters of code, but the string itself contains just one - an apostrophe ('), escaped using hexadecimal notation.
Santa's list is a file that contains many double-quoted string literals, one on each line. The only escape sequences used are \\ (which represents a single backslash), \" (which represents a lone double-quote character), and \x plus two hexadecimal characters (which represents a single character with that ASCII code).

Disregarding the whitespace in the file, what is the number of characters of code for string literals minus the number of characters in memory for the values of the strings in total for the entire file?

For example, given the four strings above, the total number of characters of string code (2 + 5 + 10 + 6 = 23) minus the total number of characters in memory for string values (0 + 3 + 7 + 1 = 11) is 23 - 11 = 12.

*/

import (
	"fmt"
	"strings"

	goutils "github.com/simonski/goutils"
)

// AOC_2015_07 is the entrypoint
func AOC_2015_08(cli *goutils.CLI) {
	AOC_2015_08_part1_attempt1(cli)
	// AOC_2015_08_part2_attempt1(cli)
}

func AOC_2015_08_part1_attempt1(cli *goutils.CLI) {

	data := DAY_2015_08_DATA
	splits := strings.Split(data, "\n")
	length := 0
	character_length := 0
	for _, line := range splits {
		line = strings.TrimSpace(line)
		l := len(line)
		cl := total_parsed_character_length(line)

		length += l
		character_length += cl
	}
	fmt.Printf("%v - %v = %v\n", length, character_length, length-character_length)

	// 1206 too low
	// 1246 wrong
	// 1257 wrong

}

func total_parsed_character_length(line string) int {
	// delim1 := `\\`
	// delim2 := `\"`
	// delim3 := `\xdd`

	// replace all \\

	original := line
	line = line[1:]
	line = line[:len(line)-1]

	line4 := ""
	for {

		// o := line
		// fmt.Printf("Received '%v', working with '%v'\n", original, line)

		line2 := strings.ReplaceAll(line, `\\`, `\`)
		// fmt.Printf("\\\\ '%v'\n", line2)
		// quotedBackslashes := len(line) - len(line2)

		line3 := strings.ReplaceAll(line2, `\"`, `"`)
		// fmt.Printf("\\\" '%v'\n", line3)
		// quotedQuotes := (len(line2) - len(line3))

		line4 = line3
		for index := 0; index < 255; index++ {
			value := fmt.Sprintf(`\x%x`, index) // as hex
			for {
				if strings.Index(line4, value) > -1 {
					i := strings.Index(line4, value)
					line4 = line4[0:i] + "X" + line4[i+4:]
					fmt.Printf("''%v'  -> %v'\n", line4, value)
					// line4 = strings.ReplaceAll(line4, value, "X")
				} else {
					break
				}
			}
		}
		// fmt.Printf("\\xNN '%v'\n", line4)

		// hexchars := (len(line3) - len(line4)) / 4
		// regularChars := len(line4)

		fmt.Printf("%v -> %v\n", original, line4)

		// if line4 == o {
		break
		// }
		// line = line4

	}
	// total := quotedBackslashes + quotedQuotes + hexchars + regularChars
	// fmt.Printf("%v + %v + %v + %v = %v\n", quotedBackslashes, quotedQuotes, hexchars, regularChars, total)
	// return total
	return len(line4)

}
