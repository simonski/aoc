package aoc2021

import (
	"fmt"
	"sort"

	"github.com/simonski/aoc/utils"
	goutils "github.com/simonski/goutils"
)

/*
-- Day 7: The Treachery of Whales ---
A giant whale has decided your submarine is its next meal, and it's much faster than you are. There's nowhere to run!

Suddenly, a swarm of crabs (each in its own tiny submarine - it's too deep for them otherwise) zooms in to rescue you! They seem to be preparing to blast a hole in the ocean floor; sensors indicate a massive underground cave system just beyond where they're aiming!

The crab submarines all need to be aligned before they'll have enough power to blast a large enough hole for your submarine to get through. However, it doesn't look like they'll be aligned before the whale catches you! Maybe you can help?

There's one major catch - crab submarines can only move horizontally.

You quickly make a list of the horizontal position of each crab (your puzzle input). Crab submarines have limited fuel, so you need to find a way to make all of their horizontal positions match while requiring them to spend as little fuel as possible.

For example, consider the following horizontal positions:

16,1,2,0,4,2,7,1,2,14
This means there's a crab with horizontal position 16, a crab with horizontal position 1, and so on.

Each change of 1 step in horizontal position of a single crab costs 1 fuel. You could choose any horizontal position to align them all on, but the one that costs the least fuel is horizontal position 2:

Move from 16 to 2: 14 fuel
Move from 1 to 2: 1 fuel
Move from 2 to 2: 0 fuel
Move from 0 to 2: 2 fuel
Move from 4 to 2: 2 fuel
Move from 2 to 2: 0 fuel
Move from 7 to 2: 5 fuel
Move from 1 to 2: 1 fuel
Move from 2 to 2: 0 fuel
Move from 14 to 2: 12 fuel
This costs a total of 37 fuel. This is the cheapest possible outcome; more expensive outcomes include aligning at position 1 (41 fuel), position 3 (39 fuel), or position 10 (71 fuel).

Determine the horizontal position that the crabs can align to using the least fuel possible. How much fuel must they spend to align to that position?
*/

type Crab struct {
	Initial  int
	Position int
}

func (crab *Crab) Fuel() int {
	return crab.Position - crab.Initial
}

type OceanOfCrabs struct {
	Crabs  []*Crab
	Total  int
	Min    int
	Max    int
	Mean   int
	Median int
}

func (o *OceanOfCrabs) Debug() string {
	return fmt.Sprintf("Ocean, %v crabs, total=%v, min=%v, max=%v, mean=%v, median=%v", len(o.Crabs), o.Total, o.Min, o.Max, o.Mean, o.Median)
}

// helper finds min, max, mean, median, mode
func (o *OceanOfCrabs) Scan() []int {

	total := 0
	positions := make([]int, 0)
	for _, crab := range o.Crabs {
		total += crab.Position
		positions = append(positions, crab.Position)
	}
	sort.Ints(positions)

	o.Total = total
	o.Min = positions[0]
	o.Max = positions[len(positions)-1]
	o.Mean = total / len(positions)
	o.Median = positions[len(positions)/2]

	return positions
}

func (o *OceanOfCrabs) CostOfMoveToPositionPart1(positionToMoveTo int) int {
	total := 0
	for _, crab := range o.Crabs {
		expense := goutils.Abs(crab.Position - positionToMoveTo)
		total += expense
	}
	return total
}

func (o *OceanOfCrabs) CostOfMoveToPositionPart2(positionToMoveTo int) int64 {
	total := int64(0)
	for _, crab := range o.Crabs {
		distance := goutils.Abs(crab.Position - positionToMoveTo)
		total += int64(Cost(distance))
	}
	return total
}

func NewCrab(position int) *Crab {
	c := Crab{Initial: position, Position: position}
	return &c
}

func NewOceanOfCrabs(data string) *OceanOfCrabs {
	ints := utils.SplitDataToListOfInts(data, ",")
	crabs := make([]*Crab, 0)

	for _, value := range ints {
		crab := NewCrab(value)
		crabs = append(crabs, crab)
	}

	ocean := OceanOfCrabs{Crabs: crabs}
	ocean.Scan()
	return &ocean

}

// rename this to the year and day in question
func (app *Application) Y2021D07P1() {
	ocean := NewOceanOfCrabs(DAY_2021_07_TEST_DATA)
	fmt.Println(ocean.Debug())
	median := ocean.Median
	cost := ocean.CostOfMoveToPositionPart1(median)
	fmt.Printf("The cost of moving to position %v is %v.\n", median, cost)

	oceanX := NewOceanOfCrabs(DAY_2021_07_DATA)
	fmt.Println(oceanX.Debug())
	medianX := oceanX.Median
	costX := oceanX.CostOfMoveToPositionPart1(medianX)
	fmt.Printf("The cost of moving to position %v is %v.\n", medianX, costX)

}

// // rename this to the year and day in question

/*
--- Part Two ---
The crabs don't seem interested in your proposed solution. Perhaps you misunderstand crab engineering?

As it turns out, crab submarine engines don't burn fuel at a constant rate. Instead, each change of 1 step in horizontal position costs 1 more unit of fuel than the last: the first step costs 1, the second step costs 2, the third step costs 3, and so on.

As each crab moves, moving further becomes more expensive. This changes the best horizontal position to align them all on; in the example above, this becomes 5:

Move from 16 to 5: 66 fuel
Move from 1 to 5: 10 fuel
Move from 2 to 5: 6 fuel
Move from 0 to 5: 15 fuel
Move from 4 to 5: 1 fuel
Move from 2 to 5: 6 fuel
Move from 7 to 5: 3 fuel
Move from 1 to 5: 10 fuel
Move from 2 to 5: 6 fuel
Move from 14 to 5: 45 fuel
This costs a total of 168 fuel. This is the new cheapest possible outcome; the old alignment position (2) now costs 206 fuel instead.

Determine the horizontal position that the crabs can align to using the least fuel possible so they can make you an escape route! How much fuel must they spend to align to that position?
*/

func Cost(distance int) int {
	total := 0
	for index := 1; index <= distance; index++ {
		total += index
	}
	return total
}

func (app *Application) Y2021D07P2() {
	ocean := NewOceanOfCrabs(DAY_2021_07_DATA)
	fmt.Println(ocean.Debug())
	median := ocean.Median
	cost := ocean.CostOfMoveToPositionPart2(median)
	fmt.Printf("The cost of moving to position %v is %v.\n", median, cost)

	positions := ocean.Scan()
	// size := len(ocean.Crabs) / 2
	// mean := ocean.Mean
	minPosition := -1
	minCost := int64(100000000000)
	for index := 0; index < len(positions)/2; index++ {
		rpos := len(positions)/2 + index
		lpos := len(positions)/2 - index

		cost = int64(ocean.CostOfMoveToPositionPart2(rpos))
		fmt.Printf("The cost of moving to position %v is %v.\n", rpos, cost)
		if cost < minCost {
			minPosition = rpos
			minCost = cost
		}

		cost = ocean.CostOfMoveToPositionPart2(lpos)
		fmt.Printf("The cost of moving to position %v is %v.\n", lpos, cost)
		if cost < minCost {
			minPosition = lpos
			minCost = cost
		}
	}

	fmt.Printf("The minimumCost is at position %v is %v.\n", minPosition, minCost)

	// oceanX := NewOceanOfCrabs(DAY_2021_07_DATA)
	// fmt.Println(oceanX.Debug())
	// medianX := oceanX.Median
	// costX := oceanX.CostOfMoveToPosition(medianX)
	// fmt.Printf("The cost of moving to position %v is %v.\n", medianX, costX)

}

// rename and uncomment this to the year and day in question once complete for a gold star!
// func (app *Application) Y20XXDXXP1Render() {
// }

// rename and uncomment this to the year and day in question once complete for a gold star!
// func (app *Application) Y20XXDXXP2Render() {
// }

// this is what we will reflect and call - so both parts with run. It's up to you to make it print nicely etc.
// The app reference has a CLI for logging.
// func (app *Application) Y2021D07() {
// 	app.Y2021D07P1()
// 	app.Y2021D07P2()
// }
