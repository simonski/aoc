package aoc2019

import (
	"fmt"
	"strconv"
	"strings"
)

/*
--- Day XX: Description ---

*/

/*
For a mass of 12, divide by 3 and round down to get 4, then subtract 2 to get 2.
For a mass of 14, dividing by 3 and rounding down still yields 4, so the fuel required is also 2.
For a mass of 1969, the fuel required is 654.
For a mass of 100756, the fuel required is 33583.

Fuel required to launch a given module is based on its mass. Specifically, to
find the fuel required for a module,
take its mass, divide by three, round down, and subtract 2.

*/
func (app *Application) FuelRequired(mass int) int {
	divided := (mass / 3) - 2
	return divided
}

func (app *Application) FuelRequiredPart2(mass int) int {
	// recursive
	result := (mass / 3) - 2
	if result <= 0 {
		// fmt.Printf("discounted %v\n", result)
		return mass
	} else {
		// fmt.Printf("remaining  %v\n", result)
		return mass + app.FuelRequiredPart2(result)
	}
}

// rename this to the year and day in question
func (app *Application) Y2019D01P1() {
	fmt.Printf("Y2019 D01 ok!\n")
	lines := strings.Split(DAY_2019_01_DATA, "\n")
	total := 0
	for _, line := range lines {
		mass, _ := strconv.Atoi(line)
		fuelRequired := app.FuelRequired(mass)
		total += fuelRequired
	}
	fmt.Printf("Part 1 total: %v\n", total)

}

// rename this to the year and day in question
func (app *Application) Y2019D01P2() {
	fmt.Printf("Y2019 D02 ok!\n")
	lines := strings.Split(DAY_2019_01_DATA, "\n")
	total := 0
	for _, line := range lines {
		mass, _ := strconv.Atoi(line)
		fuelRequired := app.FuelRequiredPart2(mass) - mass
		total += fuelRequired
	}
	fmt.Printf("Part 2 total: %v\n", total)

	// answer := app.FuelRequiredPart2(100756) - 100756
	// fmt.Printf("Y2019 D02 Test: 100756 (requires 50346), our answer: %v\n", answer)

	// fmt.Printf("Y2019 D02 ok!\n")
	// lines := strings.Split(DAY_2019_01_DATA, "\n")
	// total := 0
	// for _, line := range lines {
	// 	mass, _ := strconv.Atoi(line)
	// 	fuelRequired := app.FuelRequiredPart2(mass)
	// 	total += fuelRequired
	// }
	// fmt.Printf("Part 2 total: %v\n", total)

}

// rename and uncomment this to the year and day in question once complete for a gold star!
// func (app *Application) Y20XXDXXP1Render() {
// }

// rename and uncomment this to the year and day in question once complete for a gold star!
// func (app *Application) Y20XXDXXP2Render() {
// }

// this is what we will reflect and call - so both parts with run. It's up to you to make it print nicely etc.
// The app reference has a CLI for logging.
func (app *Application) Y2019D01() {
	app.Y2019D01P1()
	app.Y2019D01P2()
}
