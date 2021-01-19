package main

/*
The elves are running low on wrapping paper, and so they need to submit an order for more. They have a list of the dimensions (length l, width w, and height h) of each present, and only want to order exactly as much as they need.

Fortunately, every present is a box (a perfect right rectangular prism), which makes calculating the required wrapping paper for each gift a little easier: find the surface area of the box, which is 2*l*w + 2*w*h + 2*h*l. The elves also need a little extra paper for each present: the area of the smallest side.

For example:

A present with dimensions 2x3x4 requires 2*6 + 2*12 + 2*8 = 52 square feet of wrapping paper plus 6 square feet of slack, for a total of 58 square feet.
A present with dimensions 1x1x10 requires 2*1 + 2*10 + 2*10 = 42 square feet of wrapping paper plus 1 square foot of slack, for a total of 43 square feet.
All numbers in the elves' list are in feet. How many total square feet of wrapping paper should they order?
*/

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	goutils "github.com/simonski/goutils"
)

// AOC_2015_02 is the entrypoint
func AOC_2015_02(cli *goutils.CLI) {
	AOC_2015_02_part1_attempt1(cli)
	AOC_2015_02_part2_attempt1(cli)
}

func AOC_2015_02_part1_attempt1(cli *goutils.CLI) {
	lines := strings.Split(DAY_2015_02_DATA, "\n")
	total := 0
	for _, line := range lines {
		p := NewPresent(line)
		total += p.Area()
	}
	fmt.Printf("Area: %v\n", total)
}

/*
--- Part Two ---
The elves are also running low on ribbon. Ribbon is all the same width, so they only have to worry about the length they need to order, which they would again like to be exact.

The ribbon required to wrap a present is the shortest distance around its sides, or the smallest perimeter of any one face. Each present also requires a bow made out of ribbon as well; the feet of ribbon required for the perfect bow is equal to the cubic feet of volume of the present. Don't ask how they tie the bow, though; they'll never tell.

For example:

A present with dimensions 2x3x4 requires 2+2+3+3 = 10 feet of ribbon to wrap the present plus 2*3*4 = 24 feet of ribbon for the bow, for a total of 34 feet.
A present with dimensions 1x1x10 requires 1+1+1+1 = 4 feet of ribbon to wrap the present plus 1*1*10 = 10 feet of ribbon for the bow, for a total of 14 feet.
How many total feet of ribbon should they order?
*/
func AOC_2015_02_part2_attempt1(cli *goutils.CLI) {
	lines := strings.Split(DAY_2015_02_DATA, "\n")
	volume := 0
	perimeter := 0
	for _, line := range lines {
		p := NewPresent(line)
		perimeter += p.Perimeter()
		volume += p.Volume()
	}
	total := perimeter + volume
	fmt.Printf("Volume %v Perimeter %v, total %v\n", volume, perimeter, total)
}

type Present struct {
	l int
	w int
	h int
}

func (p *Present) Area() int {
	l := p.l
	w := p.w
	h := p.h

	a1 := (2 * l * w)
	a2 := (2 * w * h)
	a3 := (2 * h * l)

	t := l * w
	t = goutils.Min(t, w*h)
	t = goutils.Min(t, h*l)

	return a1 + a2 + a3 + t
}

func (p *Present) Volume() int {
	return p.l * p.h * p.w
}

func (p *Present) Perimeter() int {
	arr := make([]int, 0)
	arr = append(arr, p.l)
	arr = append(arr, p.h)
	arr = append(arr, p.w)
	sort.Ints(arr)
	fmt.Printf("%v\n", arr)
	return arr[0] + arr[0] + arr[1] + arr[1]
}

func NewPresent(line string) *Present {
	splits := strings.Split(strings.TrimSpace(line), "x")
	l, _ := strconv.Atoi(splits[0])
	w, _ := strconv.Atoi(splits[1])
	h, _ := strconv.Atoi(splits[2])
	p := Present{l: l, w: w, h: h}
	return &p
}
