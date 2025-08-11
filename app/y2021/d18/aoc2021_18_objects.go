package aoc2021

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	utils "github.com/simonski/goutils"
)

/*
Describe the problem
*/

const REDUCE_DEBUG = false
const REDUCE_TRACE = false

type Pair struct {
	Line     string
	Element1 *Pair
	Element2 *Pair
}

func NewPair(line string) *Pair {
	line = utils.StripWhitespace(line)
	p := Pair{Line: line}
	return &p
}

// func startAndEndsWith(line string, start string, end string) bool {
// 	lsame := line[0:1] == start
// 	rsame := line[len(line)-1:len(line)] == end
// 	return lsame && rsame
// }

// For example, [1,2] + [[3,4],5] becomes [[1,2],[[3,4],5]].
func (p *Pair) Add(line string) string {
	line = strings.ReplaceAll(line, " ", "")
	line = strings.ReplaceAll(line, "\t", "")
	// if startAndEndsWith(line, "[", "]") && startAndEndsWith(p.Line, "[", "]") {
	// 	newline := fmt.Sprintf("%v,%v", p.Line, line)
	// 	p.Line = newline
	// 	return p.Line
	// } else {
	newline := fmt.Sprintf("[%v,%v]", p.Line, line)
	p.Line = newline
	return p.Line

	// }

}

func (p *Pair) Magnitude() int {
	m := 0
	line := p.Line
	for {
		found, left, pair, right := p.FindFirstPair(line)
		if REDUCE_DEBUG {
			fmt.Printf("Magnitude(%v), pair=%v, left=%v, right=%v\n", line, pair, left, right)
		}
		if !found {
			value, _ := strconv.Atoi(line)
			return value
		} else {
			pair = strings.ReplaceAll(pair, "[", "")
			pair = strings.ReplaceAll(pair, "]", "")
			splits := strings.Split(pair, ",")
			lval, _ := strconv.Atoi(splits[0])
			rval, _ := strconv.Atoi(splits[1])
			val := 3*lval + 2*rval
			if left == "" && right == "" {
				return val
			}
			line = fmt.Sprintf("%v%v%v", left, val, right)
			if REDUCE_DEBUG {
				fmt.Printf(">> Magnitude(%v), pair=%v, left=%v, right=%v\n", line, pair, left, right)
				fmt.Println()
			}
		}
	}

	return m
}

func (p *Pair) FindFirstPair(line string) (bool, string, string, string) {
	pairCount := 0
	candidate := ""
	lindex := 0
	rindex := 0
	for index := 0; index < len(line); index++ {
		c := line[index : index+1]
		if c == "[" {
			lindex = index
			pairCount += 1
			candidate = c
		} else if c == "]" {
			candidate += c
			rindex = index
			if p.IsValuePair(candidate) {
				left := line[0:lindex]
				right := line[rindex+1:]
				return true, left, candidate, right
			}
		} else {
			candidate += c
		}
	}
	return false, "", "", ""
}

func (p *Pair) Reduce() int {
	return p.ReduceX()
}

func (p *Pair) DoReduce(style string) (int, string, string) {
	return p.DoReduceX(style)
}

func (p *Pair) ReduceX() int {
	total := 0
	step := 0
	style := "none"
	for {
		before := p.Line
		reductions, reduction, s := p.DoReduceX(style)
		style = s
		after := p.Line
		if REDUCE_DEBUG {
			fmt.Printf("Reduce(%v) style=%v, \n%v\n-> (%v)\n%v\n\n", step, style, before, reduction, after)
		}
		step += 1
		total += reductions
		if reductions == 0 && style == "none" {
			break
		} else if reductions == 0 {
			style = "none"
		}
	}
	return total
}

func (p *Pair) DoReduceX(style string) (int, string, string) {

	if REDUCE_DEBUG {
		fmt.Printf("->Reduce(%v):\n", style)
	}

	line := p.Line
	indexOfFifth := p.FindIndexOfFifthPair(line)
	indexOfBiggieSmalls := p.FindIndexOfRegularNumberGreaterOrEqualToTen(line)
	if REDUCE_DEBUG {
		fmt.Printf("->Reduce(%v): indexOfFifth=%v, indexOfBig=%v\n", style, indexOfFifth, indexOfBiggieSmalls)
	}
	if (style == "pair" || style == "none") && indexOfFifth > -1 {
		style = "pair"
		if REDUCE_TRACE {
			fmt.Printf("DoReduce(): 4-deep pair found at %v (will explode)\n.", indexOfFifth)
		}
		newLine, pairExploded := p.ExplodeLeftMostPair(indexOfFifth, line)
		p.Line = newLine
		return 1, fmt.Sprintf("exploded %v", pairExploded), style
	} else if "style" == "pair" && indexOfFifth == -1 {
		if REDUCE_TRACE {
			fmt.Printf("DoReduce() no 4-deep pair found %v.\n", p.Line)
		}
		return 0, "pair", style
	}

	if (style == "big" || style == "none") && indexOfBiggieSmalls > -1 {
		if REDUCE_TRACE {
			fmt.Printf("Reduce: number > 10 found at %v (will split).\n", indexOfBiggieSmalls)
		}
		newLine, pairSplit := p.SplitAt(indexOfBiggieSmalls, line)
		p.Line = newLine
		return 1, fmt.Sprintf("split %v", pairSplit), style
	} else {
		if REDUCE_TRACE {
			fmt.Printf("DoReduce(%v): no number > 10 found.'n", p.Line)
		}
		return 0, "big", style
	}
}

func (p *Pair) ReduceY() int {
	total := 0
	step := 0
	for {
		before := p.Line
		reductions, reduction := p.DoReduceY()
		after := p.Line
		if REDUCE_DEBUG {
			fmt.Printf("Reduce(%v) \n%v\n-> (%v)\n%v\n\n", step, before, reduction, after)
		}
		step += 1
		total += reductions
		if reductions == 0 {
			break
		}
	}
	return total
}

func (p *Pair) DoReduceY() (int, string) {

	line := p.Line
	indexOfFifth := p.FindIndexOfFifthPair(line)
	indexOfBiggieSmalls := p.FindIndexOfRegularNumberGreaterOrEqualToTen(line)

	if (indexOfFifth > -1 && indexOfBiggieSmalls == -1) || ((indexOfFifth > -1 && indexOfBiggieSmalls > -1) && indexOfFifth < indexOfBiggieSmalls) {
		if REDUCE_TRACE {
			fmt.Printf("DoReduce(): 4-deep pair found at %v (will explode)\n.", indexOfFifth)
		}
		newLine, pairExploded := p.ExplodeLeftMostPair(indexOfFifth, line)
		p.Line = newLine
		return 1, fmt.Sprintf("exploded %v", pairExploded)
	} else {
		if REDUCE_TRACE {
			fmt.Printf("DoReduce() no 4-deep pair found %v.\n", p.Line)
		}
	}

	if indexOfBiggieSmalls > -1 {
		if REDUCE_TRACE {
			fmt.Printf("Reduce: number > 10 found at %v (will split).\n", indexOfBiggieSmalls)
		}
		newLine, pairSplit := p.SplitAt(indexOfBiggieSmalls, line)
		p.Line = newLine
		return 1, fmt.Sprintf("split %v", pairSplit)
	} else {
		if REDUCE_TRACE {
			fmt.Printf("DoReduce(%v): no number > 10 found.'n", p.Line)
		}
	}
	return 0, "none"
}

func (p *Pair) IsValuePair(candidate string) bool {

	regex := `^\[\d+,\d+\]$`
	match, err := regexp.MatchString(regex, candidate)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	return match
}

func (p *Pair) FindIndexOfFifthPair(line string) int {
	pairCount := 0
	candidate := ""
	openingIndex := -1
	for index := 0; index < len(line); index++ {
		c := line[index : index+1]
		if c == "[" {
			openingIndex = index
			pairCount += 1
			candidate = c
		} else if c == "]" {
			candidate += c
			// check if this candidate is a value pair
			isPair := p.IsValuePair(candidate)
			if pairCount >= 5 && isPair {
				if REDUCE_TRACE {
					fmt.Printf("RETURN THIS index=%v, openingIndex=%v, pairCount=`%v`, candidate=`%v`, isPair=`%v`\n", index, openingIndex, pairCount, candidate, isPair)
				}
				return openingIndex
			} else {
				if REDUCE_TRACE {
					fmt.Printf("CONTINUE SEARCHING index=%v, openingIndex=%v, pairCount=`%v`, candidate=`%v`, isPair=`%v`\n", index, openingIndex, pairCount, candidate, isPair)
				}

			}
			pairCount -= 1
		} else {
			candidate += c
			// isPair := p.IsValuePair(candidate)
			// fmt.Printf("index=%v, openingIndex=%v, pairCount=%v, candidate=`%v`, isPair=%v\n", index, openingIndex, pairCount, candidate, isPair)
		}

	}
	return -1
}

func (p *Pair) FindIndexOfRegularNumberGreaterOrEqualToTen(line string) int {
	value := "0"
	for index := 0; index < len(line); index++ {
		letter := line[index : index+1]
		if letter == "[" || letter == "]" || letter == "," {
			// start again
			ival, _ := strconv.Atoi(value)
			if ival > 9 {
				return index - 2
			}
			value = "0"
		} else {
			value += letter
		}
	}
	return -1
}

func (p *Pair) ExplodeLeftMostPair(index int, line string) (string, string) {
	/*
		To explode a pair, the pair's left value is added to the first regular number to the left of the exploding pair (if any), and the pair's right value is added to the first regular number to the right of the exploding pair (if any). Exploding pairs will always consist of two regular numbers. Then, the entire exploding pair is replaced with the regular number 0.
	*/

	// left is everything up to this pair

	subpair := p.FindPairAtIndex(index, line)
	right := line[index+len(subpair):]
	left := line[0:index]
	if REDUCE_TRACE {
		fmt.Printf("original: %v\nleft    : %v\npair    : %v\nright   : %v\n", line, left, subpair, right)
	}

	isubpair := strings.ReplaceAll(subpair, "[", "")
	isubpair = strings.ReplaceAll(isubpair, "]", "")

	left_value := strings.Split(isubpair, ",")[0]
	right_value := strings.Split(isubpair, ",")[1]

	/*
		To explode a pair, the pair's left value is added to the first regular number to the left of the exploding pair (if any), and the pair's right value is added to the first regular number to the right of the exploding pair (if any). Exploding pairs will always consist of two regular numbers. Then, the entire exploding pair is replaced with the regular number 0.
	*/
	// look left for the first number

	// Step 1. Figure out the left side
	circuitBreaker := false
	leftNumber := ""
	leftNumberIndexStart := -1
	leftNumberIndexEnd := -1
	for lindex := index - 1; lindex >= 0; lindex-- {
		c := left[lindex : lindex+1]
		if c == "]" || c == "," || c == "[" {
			// if we have stsrted reading a number then we want to stop at this pooint
			if circuitBreaker {
				leftNumberIndexStart = lindex
				break
			}
		} else {
			// now reading a number
			if !circuitBreaker {
				leftNumberIndexEnd = lindex
			}
			circuitBreaker = true
			leftNumber = c + leftNumber
		}
	}

	if circuitBreaker {
		if REDUCE_TRACE {
			fmt.Println("We find a left number.")
			fmt.Printf("leftNumber           : %v\n", leftNumber)
			fmt.Printf("leftNumberIndexStart : %v\n", leftNumberIndexStart)
			fmt.Printf("leftNumberIndexEnd   : %v\n", leftNumberIndexEnd)
		}
		left_ival, _ := strconv.Atoi(left_value)
		left_inumber, _ := strconv.Atoi(leftNumber)
		if REDUCE_TRACE {
			fmt.Printf("left_ival: %v, left_inumber: %v\n", left_ival, left_inumber)
		}
		new_left_number := left_ival + left_inumber

		left_without_left_number := left[0 : leftNumberIndexStart+1]

		new_left := fmt.Sprintf("%v%v%v", left_without_left_number, new_left_number, left[leftNumberIndexEnd+1:])
		left = new_left

		if REDUCE_TRACE {
			fmt.Printf("We will add our leftmost number to the left number (%v + %v)\n", leftNumber, left_value)
			fmt.Printf("left without the left nubmer is %v\n", left_without_left_number)
			fmt.Printf("New left will be %v\n", left)
		}

	} else {
		if REDUCE_TRACE {
			fmt.Println("We did not find a left number.")
			fmt.Println("So we won't make a change to left at all.")
		}
	}

	// BREAK FOR A MOMENT
	// so "left" is now correct.  We want to now figure out "right"
	// once we figure out right, we can then add
	// left + "0" + right

	// Step 2. Figure out the right hand side
	/*
		To explode a pair, the pair's left value is added to the first regular number to the left of the exploding pair (if any), and the pair's right value is added to the first regular number to the right of the exploding pair (if any). Exploding pairs will always consist of two regular numbers. Then, the entire exploding pair is replaced with the regular number 0.
	*/
	// we have right_value to add to the first available number on the right string

	// read forward until we get a whole number
	circuitBreaker = false
	rightNumberIndexStart := -1
	rightNumberIndexEnd := -1
	rnumber := ""
	for rindex := 0; rindex < len(right); rindex++ {
		c := right[rindex : rindex+1]
		if c == "[" || c == "]" || c == "," {
			if circuitBreaker {
				// then we have already finished our number time
				rightNumberIndexEnd = rindex
				break
			}
		} else {
			rnumber = rnumber + c
			if !circuitBreaker {
				circuitBreaker = true // we want only numbers now
				rightNumberIndexStart = rindex
			}
		}
	}

	if circuitBreaker {
		if REDUCE_TRACE {
			fmt.Print("We found a right number to add to\n")
		}
		// now we have a nubmer (rnumber) and the position it occupied
		// so we can convert it to an integer, add the right nubmer to it, replace it, then join the whole thign back

		if REDUCE_TRACE {
			fmt.Printf("we are going to add %v to the right number is %v\n", right_value, rnumber)
		}
		rval, _ := strconv.Atoi(rnumber)
		right_ival, _ := strconv.Atoi(right_value)
		new_rval := right_ival + rval
		rstring := fmt.Sprintf("%v%v%v", right[0:rightNumberIndexStart], new_rval, right[rightNumberIndexEnd:])

		if REDUCE_TRACE {
			fmt.Printf("new right will be %v\n", rstring)
		}

		newline := left + "0" + rstring
		if REDUCE_TRACE {
			fmt.Printf("newline is %v\n", newline)
		}

		return newline, subpair
	} else {
		if REDUCE_TRACE {
			fmt.Print("We didn not find a number to the right, so we won't add it to anything.\n")
		}
		newline := left + "0" + right
		if REDUCE_TRACE {
			fmt.Printf("newline is %v\n", newline)
		}
		return newline, subpair
	}
}

func (p *Pair) FindPairAtIndex(index int, line string) string {
	size := 0
	pair := ""
	for position := index; position < len(line); position++ {
		v := line[position : position+1]
		pair += v
		if v == "[" {
			size += 1
		} else if v == "]" {
			size -= 1
			if size == 0 {
				return pair
			}
		}
	}
	return pair
}

/**
To split a regular number, replace it with a pair; the left element of the pair should be the regular number divided by two and rounded down, while the right element of the pair should be the regular number divided by two and rounded up. For example, 10 becomes [5,5], 11 becomes [5,6], 12 becomes [6,6], and so on.

*/

func (p *Pair) SplitAt(index int, value string) (string, string) {
	/*
		To split a regular number, replace it with a pair; the left element of the pair should be the regular number divided by two and rounded down, while the right element of the pair should be the regular number divided by two and rounded up. For example, 10 becomes [5,5], 11 becomes [5,6], 12 becomes [6,6], and so on.
	*/
	left := value[0:index]
	right := value[index+2:]
	v := value[index : index+2]
	replacement := p.Split(v)
	if REDUCE_TRACE {
		fmt.Printf("SplitAt(index=%v, value=%v), left=%v, right=%v, v=%v, replacement=%v\n", index, value, left, right, v, replacement)
	}
	result := fmt.Sprintf("%v%v%v", left, replacement, right)
	if REDUCE_TRACE {
		fmt.Printf("SplitAt(index=%v)\ninput= %v\noutput=%v\n", index, value, result)
	}
	return result, v

}

func (p *Pair) Split(value string) string {
	intValue, _ := strconv.Atoi(value)
	line := ""
	if intValue%2 == 0 {
		left := intValue / 2
		right := left
		line = fmt.Sprintf("[%v,%v]", left, right)
	} else {
		left := intValue / 2
		right := left + 1
		line = fmt.Sprintf("[%v,%v]", left, right)
	}
	return line
}
