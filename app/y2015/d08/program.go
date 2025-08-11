package d08

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
	"github.com/simonski/cli"
)

/*
Matchsticks
*/

type Puzzle struct {
	title     string
	year      int
	day       int
	input     string
	lines     []string
	HasRender bool
	HasWWW    bool
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("2015")
	iday, _ := strconv.Atoi("08")
	p := Puzzle{year: iyear, day: iday, title: "Matchsticks"}
	p.Load(input)
	return &p
}

func NewPuzzle() *Puzzle {
	return NewPuzzleWithData(REAL_DATA)
}

func (puzzle *Puzzle) Load(input string) {
	lines := strings.Split(input, "\n")
	puzzle.input = input
	puzzle.lines = lines
}

func (puzzle *Puzzle) GetSummary() utils.Summary {
	s := utils.Summary{Day: puzzle.day, Year: puzzle.year, Name: puzzle.title}
	s.ProgressP1 = utils.Completed
	s.ProgressP2 = utils.Completed
	s.DateStarted = "d08"
	return s
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}

func (puzzle *Puzzle) RunWWW() string {
	return ""
}

func (puzzle *Puzzle) Render(cli *cli.CLI) {
}

func (puzzle *Puzzle) Part1() {

	data := REAL_DATA
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

}

func (puzzle *Puzzle) Part2() {

	data := REAL_DATA
	splits := strings.Split(data, "\n")
	// length := 0
	// character_length := 0
	for _, line := range splits {
		total_quoted_character_length(line)
		break
	}
}

func total_quoted_character_length(line string) int {
	count := len(line)
	for index, _ := range line {
		c := line[index : index+1]
		if c == `\` || c == "\"" {
			count++
		}
		fmt.Printf("%v\n", c)
	}
	return count
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
