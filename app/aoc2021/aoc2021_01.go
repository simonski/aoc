package aoc2021

import (
	"fmt"

	utils "github.com/simonski/aoc/utils"
)

/*
--- Day 1: Sonar Sweep ---
You're minding your own business on a ship at sea when the overboard alarm goes off! You rush to see if you can help. Apparently, one of the Elves tripped and accidentally sent the sleigh keys flying into the ocean!

Before you know it, you're inside a submarine the Elves keep ready for situations like this. It's covered in Christmas lights (because of course it is), and it even has an experimental antenna that should be able to track the keys if you can boost its signal strength high enough; there's a little meter that indicates the antenna's signal strength by displaying 0-50 stars.

Your instincts tell you that in order to save Christmas, you'll need to get all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

As the submarine drops below the surface of the ocean, it automatically performs a sonar sweep of the nearby sea floor. On a small screen, the sonar sweep report (your puzzle input) appears: each line is a measurement of the sea floor depth as the sweep looks further and further away from the submarine.

For example, suppose you had the following report:

199
200
208
210
200
207
240
269
260
263
This report indicates that, scanning outward from the submarine, the sonar sweep found depths of 199, 200, 208, 210, and so on.

The first order of business is to figure out how quickly the depth increases, just so you know what you're dealing with - you never know if the keys will get carried into deeper water by an ocean current or a fish or something.

To do this, count the number of times a depth measurement increases from the previous measurement. (There is no measurement before the first measurement.) In the example above, the changes are as follows:

199 (N/A - no previous measurement)
200 (increased)
208 (increased)
210 (increased)
200 (decreased)
207 (increased)
240 (increased)
269 (increased)
260 (decreased)
263 (increased)
In this example, there are 7 measurements that are larger than the previous measurement.

How many measurements are larger than the previous measurement?
*/

func (app *Application) Y2021D01_Summary() *utils.Summary {
	s := utils.NewSummary(2021, 1)
	s.Name = "Sonar Sweep"
	s.ProgressP1 = utils.Completed
	s.ProgressP2 = utils.Completed
	return s
}

func (app *Application) Y2021D01P1() {
	ints := utils.SplitDataToListOfInts(DAY_2021_01_TEST_DATA, "\n")
	last_value := 0
	increase := 0
	for index, value := range ints {
		if index > 0 {
			if value > last_value {
				increase += 1
			}
		}
		last_value = value
	}
	fmt.Printf("2021 01/1 - Depth increased %v times in test data\n", increase)

	ints = utils.SplitDataToListOfInts(DAY_2021_01_DATA, "\n")
	last_value = 0
	increase = 0
	for index, value := range ints {
		if index > 0 {
			if value > last_value {
				increase += 1
			}
		}
		last_value = value
	}
	fmt.Printf("2021 01/1 - Depth increased %v times in real data\n", increase)

}

// rename this to the year and day in question
func (app *Application) Y2021D01P2() {
	ints := utils.SplitDataToListOfInts(DAY_2021_01_TEST_DATA, "\n")
	last_value := 0
	increase := 0
	for index := range ints {
		if index+2 < len(ints) {
			this_value := ints[index] + ints[index+1] + ints[index+2]
			if index > 0 && (this_value > last_value) {
				increase += 1
			}
			last_value = this_value

		} else {
			break
		}
	}
	fmt.Printf("2021 01/2 - Depth increased %v times in test data\n", increase)

	ints = utils.SplitDataToListOfInts(DAY_2021_01_DATA, "\n")
	last_value = 0
	increase = 0
	for index := range ints {
		if index+2 < len(ints) {
			this_value := ints[index] + ints[index+1] + ints[index+2]
			if index > 0 && (this_value > last_value) {
				increase += 1
			}
			last_value = this_value

		} else {
			break
		}
	}
	fmt.Printf("2021 01/2 - Depth increased %v times in real data\n", increase)

}

// rename and uncomment this to the year and day in question once complete for a gold star!
// func (app *Application) Y202021D01P1Render() {
// }

// rename and uncomment this to the year and day in question once complete for a gold star!
// func (app *Application) Y202021D01P2Render() {
// }

// this is what we will reflect and call - so both parts with run. It's up to you to make it print nicely etc.
// The app reference has a CLI for logging.
// func (app *Application) Y202021D01() {
// 	app.Y2021D01P1()
// 	app.Y2021D01P2()
// }
