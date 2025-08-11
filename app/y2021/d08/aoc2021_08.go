package aoc2021

import (
	"fmt"
	"strconv"
	"strings"

	utils "github.com/simonski/aoc/utils"
)

const SMALL = `acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf`

const BIG = `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`

func (app *Application) Y2021D08_Summary() *utils.Summary {
	s := utils.NewSummary(2021, 8)
	s.Name = "Seven Segment Search"
	s.ProgressP1 = utils.Completed
	s.ProgressP2 = utils.Completed
	return s
}

func (app *Application) Y2021D08P1() {
	runPart1(BIG)
	runPart1(DATA)
}

func (app *Application) Y2021D08P2() {
	// runPart2(SMALL)
	// runPart2(BIG)
	runPart2(DATA)
}

func measure(s6n1 string, s6n2 string, s6n3 string) (string, string, string, int, int, int) {
	if len(s6n1) < len(s6n2) {
		return s6n1, s6n2, s6n3, 1, 2, 3
	} else if len(s6n2) < len(s6n1) {
		return s6n2, s6n1, s6n3, 2, 1, 3
	} else {
		return s6n3, s6n1, s6n2, 3, 1, 2
	}
}

func runPart2(data string) {
	lines := strings.Split(data, "\n")
	total := 0
	for _, line := range lines {
		value := part2(line)
		total += value
	}
	fmt.Printf("total is %v\n", total)

}

func part2(data string) int {
	// 1. find 1, 7, 4, 8
	p := NewPattern(data)
	// one, four, seven, eight := findOneFourSevenEight(data)

	// 2. calculate TOP (a)
	top := p.Subtract(p.Seven, p.One)
	a := top
	fmt.Printf("seven=%v, one=%v, top=%v\n", p.Seven, p.One, top)

	// 2. find the size 6 numbers
	sixes := p.Sixes

	// // 3. find the size 5 numbers
	fives := p.Fives

	fmt.Printf("fives: %v\n", fives)
	fmt.Printf("sixes: %v\n", sixes)

	// // there are three sixes (0, 6, 9)
	// // there are three fives (2, 3, 5)

	// 4. finding e and 9
	s6n1 := p.Subtract(p.Subtract(p.Subtract(sixes[0], p.Four), p.One), top)
	s6n2 := p.Subtract(p.Subtract(p.Subtract(sixes[1], p.Four), p.One), top)
	s6n3 := p.Subtract(p.Subtract(p.Subtract(sixes[2], p.Four), p.One), top)

	smallest, larger1, larger2, _, index2, index3 := measure(s6n1, s6n2, s6n3)
	// _ := sixes[index1-1]
	originalLarger1 := sixes[index2-1]
	originalLarger2 := sixes[index3-1]

	// // subtract 4 from one of the unknown sets to find the difference
	// // 9 - 4 = (abcdfg) - (a) - (bcdf) = g
	// // 6 - 4 = (abdefg) - (a) - (bcdf) = eg
	// // 0 - 4 = (abcefg) - (a) - (bcdf) = eg

	fmt.Printf("smallest=%v, larger1=%v, larger2=%v\n", smallest, larger1, larger2)
	p.Nine = smallest
	// p.Nine = smallest
	// smallest, larger1, larger2 = smallestWord(s6n1, s6n2, s6n3)
	g := smallest
	// number9 := smallest
	e := p.Subtract(larger1, smallest)

	fmt.Printf("a=%v, g=%v, e=%v\n", a, g, e)
	// // the smaller is g
	// // the diffference (e) is now found, too
	// // also we now know 9 (the smaller)

	// // 5. if we know 9 is one of the size six entries, we know
	// // 	one of the others is a six, so the other must be the 0

	fmt.Printf("originalLarger1=%v, originalLarger2=%v, p.Seven=%v\n", larger1, larger2, p.Seven)
	result1 := p.Subtract(originalLarger1, p.Seven)
	result2 := p.Subtract(originalLarger2, p.Seven)
	six := ""
	zero := ""
	originalSix := ""
	originalZero := ""
	if len(result1) == 4 {
		six = result1
		zero = result2
		originalSix = originalLarger1
		originalZero = originalLarger2
	} else {
		six = result2
		zero = result1
		originalSix = originalLarger2
		originalZero = originalLarger1
	}
	d := p.Subtract(six, zero)

	fmt.Printf("a=%v, g=%v, e=%v, d=%v\n", a, g, e, d)

	// // 	6 - 7 = (abdefg) - (acf) = bdeg
	// // 	0 - 7 = (abcefg) - (acf) = beg

	// // So the 6 contains is size 4
	// // So teh 0 is size 0
	// // The difference bwteeen them is the d in 6 now we have middle

	// // now we know
	// // d
	// // 6
	// // 9
	// // 0

	// // 6. now we can calculate 4 - 6  leaving top-right, c
	// // 	now we know c
	fmt.Printf("p.Seven=%v\n", p.Seven)
	fmt.Printf("p.Four=%v\n", p.Four)
	fmt.Printf("six=%v\n", six)
	fmt.Printf("zero=%v\n", zero)

	fmt.Printf("originalSix=%v\n", originalSix)
	fmt.Printf("originalZero=%v\n", originalZero)

	c := p.Subtract(p.Four, originalSix)
	fmt.Printf("a=%v, g=%v, e=%v, d=%v, c=%v\n", a, g, e, d, c)

	// // 7. look at 1 again, now we know c, we know f
	f := p.Subtract(p.One, c)
	fmt.Printf("a=%v, g=%v, e=%v, d=%v, c=%v, f=%v\n", a, g, e, d, c, f)

	// // 8. now we know f, the last remaining is b
	// b := remainder
	b := p.Subtract("abcdefg", a+g+e+d+c+f)
	fmt.Printf("a=%v, g=%v, e=%v, d=%v, c=%v, f=%v, b=%v\n", a, g, e, d, c, f, b)

	// so now I have a,b,c,d,e,f
	// given I know the mappings, I can construct "reprogram" a display to accept
	// one for the other, then ask the display its value
	p.Reprogram(a, b, c, d, e, f, g)
	for _, sp := range p.SignalPatterns {
		value := p.GetValue(sp)
		fmt.Printf("%v=%v\n", sp, value)
	}

	total := ""
	for _, ov := range p.OutputValue {
		value := p.GetValue(ov)
		fmt.Printf("%v=%v\n", ov, value)
		total += fmt.Sprintf("%v", value)
	}
	i, _ := strconv.Atoi(total)
	return i

	// 	acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab |
	// cdfeb fcadb cdfeb cdbaf

}

func runPart1(data string) {
	lines := strings.Split(data, "\n")
	ones := 0
	fours := 0
	sevens := 0
	eights := 0

	for _, line := range lines {
		p := NewPattern(line)
		for _, ov := range p.OutputValue {
			value := p.DigitValue(ov)
			if value == 1 {
				ones++
			} else if value == 4 {
				fours++
			} else if value == 7 {
				sevens++
			} else if value == 8 {
				eights++
			}
		}
	}

	total := ones + fours + sevens + eights
	fmt.Printf("2021-08/1: ones=%v, fours=%v, sevens=%v, eights=%v, total=%v\n", ones, fours, sevens, eights, total)
}
