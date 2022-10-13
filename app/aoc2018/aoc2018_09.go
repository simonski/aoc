package aoc2018

import (
	"fmt"

	cli "github.com/simonski/cli"
)

/*
--- Day 9: All in a Single Night ---
Every year, Santa manages to deliver all of his presents in a single night.

This year, however, he has some new locations to visit; his elves have provided him the distances between every pair of locations. He can start and end at any two (different) locations he wants, but he must visit each location exactly once. What is the shortest distance he can travel to achieve this?

For example, given the following distances:

London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141
The possible routes are therefore:

Dublin -> London -> Belfast = 982
London -> Dublin -> Belfast = 605
London -> Belfast -> Dublin = 659
Dublin -> Belfast -> London = 659
Belfast -> Dublin -> London = 605
Belfast -> London -> Dublin = 982
The shortest of these is London -> Dublin -> Belfast = 605, and so the answer is 605 in this example.

What is the distance of the shortest route?
*/

/*
thinking
This is the travelling salesman problem, which is NP-hard
Instant thoughts
node, path, route
the path can only cross once

1. write a parser with a test, parse to a struct
route.source
route.target
route.distance

2. calculate permutations (search space)

3. walk the tree of permutations to find least distance;
	for any tree fail any that exceed
	so walk the first fully to get distance


2. start and end any two

bool can(route)
bool canGoBefore(route)
a to b = distance



*/

// rename this to the year and day in question
func (app *Application) Y2018D09P1_inprogress(cli *cli.CLI) {
	rl := NewRouteLogic(DAY_2018_09_TEST_DATA)
	fmt.Printf("There are %v locations, %v test routes.\n", len(rl.Locations), len(rl.Routes))

	firstLocation, lastLocation := rl.FindFirstAndLastLocations()
	fmt.Printf("firstLocationis %v, lastLocation is %v\n", firstLocation.Name, lastLocation.Name)
	// var keys []string
	for _, route := range rl.Routes {
		fmt.Printf("\t%v\n", route.Debug())
	}

	// at each leaf we ask: is this path calculated and if so what was the depth and distance
	rl = NewRouteLogic(DAY_2018_09_DATA)
	fmt.Printf("\nThere are %v locations, there are %v routes.\n", len(rl.Locations), len(rl.Routes))
	for _, route := range rl.Routes {
		fmt.Printf("\t%v\n", route.Debug())
	}
	fmt.Printf("\n")

	// for index := uint64(1); index <= uint64(len(rl.Routes)); index++ {
	// 	fmt.Printf("!%v = %v\n", index, Factorial(index))
	// }
}

// rename this to the year and day in question
func (app *Application) Y2018D09P2_inprogress(cli *cli.CLI) {
}

// rename and uncomment this to the year and day in question once complete for a gold star!
// func (app *Application) Y202018D09P1Render() {
// }

// rename and uncomment this to the year and day in question once complete for a gold star!
// func (app *Application) Y202018D09P2Render() {
// }

// this is what we will reflect and call - so both parts with run. It's up to you to make it print nicely etc.
// The app reference has a CLI for logging.
func (app *Application) Y202018D09(cli *cli.CLI) {
	app.Y2018D09P1_inprogress(cli)
	app.Y2018D09P2_inprogress(cli)
}
