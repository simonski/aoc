package main

/*
--- Day 6: Probably a Fire Hazard ---
Because your neighbors keep defeating you in the holiday house decorating contest year after year, you've decided to deploy one million lights in a 1000x1000 grid.

Furthermore, because you've been especially nice this year, Santa has mailed you instructions on how to display the ideal lighting configuration.

Lights in your grid are numbered from 0 to 999 in each direction; the lights at each corner are at 0,0, 0,999, 999,999, and 999,0. The instructions include whether to turn on, turn off, or toggle various inclusive ranges given as coordinate pairs. Each coordinate pair represents opposite corners of a rectangle, inclusive; a coordinate pair like 0,0 through 2,2 therefore refers to 9 lights in a 3x3 square. The lights all start turned off.

To defeat your neighbors this year, all you have to do is set up your lights by doing the instructions Santa sent you in order.

For example:

turn on 0,0 through 999,999 would turn on (or leave on) every light.
toggle 0,0 through 999,0 would toggle the first line of 1000 lights, turning off the ones that were on, and turning on the ones that were off.
turn off 499,499 through 500,500 would turn off (or leave off) the middle four lights.
After following the instructions, how many lights are lit?
*/

import (
	"fmt"
	"strconv"
	"strings"

	goutils "github.com/simonski/goutils"
)

// AOC_2015_05 is the entrypoint
func AOC_2015_06(cli *goutils.CLI) {
	AOC_2015_06_part1_attempt1(cli)
	AOC_2015_06_part2_attempt1(cli)
}

func AOC_2015_06_part1_attempt1(cli *goutils.CLI) {
	splits := strings.Split(DAY_2015_06_DATA, "\n")
	grid := NewLightGrid()
	for _, instruction := range splits {
		grid.Execute(instruction)
	}
	countOn, countOff := grid.CountOnOff()
	fmt.Printf("On %v Off %v\n", countOn, countOff)
}

/*
You just finish implementing your winning light pattern when you realize you mistranslated Santa's message from Ancient Nordic Elvish.

The light grid you bought actually has individual brightness controls; each light can have a brightness of zero or more. The lights all start at zero.

The phrase turn on actually means that you should increase the brightness of those lights by 1.

The phrase turn off actually means that you should decrease the brightness of those lights by 1, to a minimum of zero.

The phrase toggle actually means that you should increase the brightness of those lights by 2.

What is the total brightness of all lights combined after following Santa's instructions?

For example:

turn on 0,0 through 0,0 would increase the total brightness by 1.
toggle 0,0 through 999,999 would increase the total brightness by 2000000.
*/
func AOC_2015_06_part2_attempt1(cli *goutils.CLI) {
	splits := strings.Split(DAY_2015_06_DATA, "\n")
	grid := NewLightGrid()
	for _, instruction := range splits {
		grid.Execute(instruction)
	}
	countOn, countOff := grid.CountOnOff()
	b := grid.TotalBrightness()
	fmt.Printf("On %v Off %v\n", countOn, countOff)
	fmt.Printf("Brightness %v\n", b)

}

type LightGrid struct {
	coordinates map[string]bool
	brightness  map[string]int
}

func NewLightGrid() *LightGrid {
	lg := LightGrid{coordinates: make(map[string]bool), brightness: make(map[string]int)}
	return &lg
}

func (grid *LightGrid) ExecuteAll(instructions string) {
	splits := strings.Split(instructions, "\n")
	for _, instruction := range splits {
		grid.Execute(instruction)
	}
}

func (grid *LightGrid) Execute(instruction string) {
	x1, y1, x2, y2 := grid.ParseCoordinates(instruction)
	if strings.Index(instruction, "toggle") > -1 {
		// toggle
		fmt.Printf("instruction = '%v', toggle (%v,%v)->(%v,%v)\n", instruction, x1, y1, x2, y2)
		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				grid.Toggle(x, y)
			}
		}

	} else if strings.Index(instruction, "turn on") > -1 {
		fmt.Printf("instruction = '%v', turn on (%v,%v)->(%v,%v)\n", instruction, x1, y1, x2, y2)
		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				grid.TurnOn(x, y)
			}
		}
	} else if strings.Index(instruction, "turn off") > -1 {
		fmt.Printf("instruction = '%v', turn off (%v,%v)->(%v,%v)\n", instruction, x1, y1, x2, y2)
		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				grid.TurnOff(x, y)
			}
		}
	}
}

func (grid *LightGrid) ParseCoordinates(instruction string) (int, int, int, int) {
	instruction = strings.ReplaceAll(instruction, "toggle ", "")
	instruction = strings.ReplaceAll(instruction, "turn on ", "")
	instruction = strings.ReplaceAll(instruction, "turn off ", "")
	coordinates := strings.ReplaceAll(instruction, " through ", ",")
	splits := strings.Split(coordinates, ",")
	x1, _ := strconv.Atoi(splits[0])
	y1, _ := strconv.Atoi(splits[1])
	x2, _ := strconv.Atoi(splits[2])
	y2, _ := strconv.Atoi(splits[3])
	return x1, y1, x2, y2
}

func (grid *LightGrid) Toggle(x int, y int) {
	key := fmt.Sprintf("%v,%v", x, y)
	value, exists := grid.coordinates[key]
	brightness, _ := grid.brightness[key]
	if exists {
		value = !value
		brightness += 2
	} else {
		value = true
		brightness = 2
	}
	grid.coordinates[key] = value
	grid.brightness[key] = brightness
}

func (grid *LightGrid) TurnOn(x int, y int) {
	key := fmt.Sprintf("%v,%v", x, y)
	grid.coordinates[key] = true

	brightness, exists := grid.brightness[key]
	if exists {
		brightness += 1
	} else {
		brightness = 1
	}
	grid.brightness[key] = brightness
}

func (grid *LightGrid) TotalBrightness() int {
	total := 0
	for _, value := range grid.brightness {
		total += value
	}
	return total
}

func (grid *LightGrid) TurnOff(x int, y int) {
	key := fmt.Sprintf("%v,%v", x, y)
	grid.coordinates[key] = false

	brightness, exists := grid.brightness[key]
	if exists {
		brightness -= 1
		if brightness < 0 {
			brightness = 0
		}
	} else {
		brightness = 0
	}
	grid.brightness[key] = brightness

}

func (grid *LightGrid) Get(x int, y int) bool {
	key := fmt.Sprintf("%v,%v", x, y)
	value, exists := grid.coordinates[key]
	if !exists {
		return false
	} else {
		return value
	}

}

func (grid *LightGrid) CountOnOff() (int, int) {
	on := 0
	off := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if grid.Get(x, y) {
				on++
			} else {
				off++
			}
		}
	}
	return on, off
}
