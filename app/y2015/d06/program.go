package d06

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
	"github.com/simonski/cli"
)

/*
Probably a Fire Hazard
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
	iday, _ := strconv.Atoi("06")
	p := Puzzle{year: iyear, day: iday, title: "Probably a Fire Hazard"}
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
	s.DateStarted = "d06"
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
	splits := strings.Split(REAL_DATA, "\n")
	grid := NewLightGrid()
	for _, instruction := range splits {
		grid.Execute(instruction)
	}
	countOn, countOff := grid.CountOnOff()
	fmt.Printf("On %v Off %v\n", countOn, countOff)
}

func (puzzle *Puzzle) Part2() {
	splits := strings.Split(REAL_DATA, "\n")
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
