package aoc2021

import (
	"fmt"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
--- Day 6: Lanternfish ---
The sea floor is getting steeper. Maybe the sleigh keys got carried this way?

A massive school of glowing lanternfish swims past. They must spawn quickly to reach such large numbers - maybe exponentially quickly? You should model their growth rate to be sure.

Although you know nothing about this specific species of lanternfish, you make some guesses about their attributes. Surely, each lanternfish creates a new lanternfish once every 7 days.

However, this process isn't necessarily synchronized between every lanternfish - one lanternfish might have 2 days left until it creates another lanternfish, while another might have 4. So, you can model each fish as a single number that represents the number of days until it creates a new lanternfish.

Furthermore, you reason, a new lanternfish would surely need slightly longer before it's capable of producing more lanternfish: two more days for its first cycle.

So, suppose you have a lanternfish with an internal timer value of 3:

After one day, its internal timer would become 2.
After another day, its internal timer would become 1.
After another day, its internal timer would become 0.
After another day, its internal timer would reset to 6, and it would create a new lanternfish with an internal timer of 8.
After another day, the first lanternfish would have an internal timer of 5, and the second lanternfish would have an internal timer of 7.
A lanternfish that creates a new fish resets its timer to 6, not 7 (because 0 is included as a valid timer value). The new lanternfish starts with an internal timer of 8 and does not start counting down until the next day.

Realizing what you're trying to do, the submarine automatically produces a list of the ages of several hundred nearby lanternfish (your puzzle input). For example, suppose you were given the following list:

3,4,3,1,2
This list means that the first fish has an internal timer of 3, the second fish has an internal timer of 4, and so on until the fifth fish, which has an internal timer of 2. Simulating these fish over several days would proceed as follows:

Initial state: 3,4,3,1,2
After  1 day:  2,3,2,0,1
After  2 days: 1,2,1,6,0,8
After  3 days: 0,1,0,5,6,7,8
After  4 days: 6,0,6,4,5,6,7,8,8
After  5 days: 5,6,5,3,4,5,6,7,7,8
After  6 days: 4,5,4,2,3,4,5,6,6,7
After  7 days: 3,4,3,1,2,3,4,5,5,6
After  8 days: 2,3,2,0,1,2,3,4,4,5
After  9 days: 1,2,1,6,0,1,2,3,3,4,8
After 10 days: 0,1,0,5,6,0,1,2,2,3,7,8
After 11 days: 6,0,6,4,5,6,0,1,1,2,6,7,8,8,8
After 12 days: 5,6,5,3,4,5,6,0,0,1,5,6,7,7,7,8,8
After 13 days: 4,5,4,2,3,4,5,6,6,0,4,5,6,6,6,7,7,8,8
After 14 days: 3,4,3,1,2,3,4,5,5,6,3,4,5,5,5,6,6,7,7,8
After 15 days: 2,3,2,0,1,2,3,4,4,5,2,3,4,4,4,5,5,6,6,7
After 16 days: 1,2,1,6,0,1,2,3,3,4,1,2,3,3,3,4,4,5,5,6,8
After 17 days: 0,1,0,5,6,0,1,2,2,3,0,1,2,2,2,3,3,4,4,5,7,8
After 18 days: 6,0,6,4,5,6,0,1,1,2,6,0,1,1,1,2,2,3,3,4,6,7,8,8,8,8
Each day, a 0 becomes a 6 and adds a new 8 to the end of the list, while each other number decreases by 1 if it was present at the start of the day.

In this example, after 18 days, there are a total of 26 fish. After 80 days, there would be a total of 5934.

Find a way to simulate lanternfish. How many lanternfish would there be after 80 days?


*/

func (app *Application) Y2021D06_Summary() *utils.Summary {
	s := utils.NewSummary(2021, 6)
	s.Name = "Lanternfish"
	s.ProgressP1 = utils.Completed
	s.ProgressP2 = utils.Completed
	return s
}

// rename this to the year and day in question
func (app *Application) Y2021D06P1() {
	bruteForceAttempt(80, DAY_2021_06_TEST_DATA)
	// algoAttempt(80, DAY_2021_06_TEST_DATA)
	// attempt1(256, DAY_2021_06_TEST_DATA)
}

func bruteForceAttempt(days int, dataStr string) int {
	s := strings.ReplaceAll(dataStr, " ", "")
	data := utils.SplitDataToListOfInts(s, ",")
	fish := data
	fmt.Printf("hi, fish=%v\n", fish)
	for day := 0; day < days; day++ {
		new_fish := 0
		for index := 0; index < len(fish); index++ {
			f := fish[index]
			f -= 1
			fish[index] = f
			if f < 0 {
				new_fish += 1
				fish[index] = 6
			}
		}
		for index := 0; index < new_fish; index++ {
			fish = append(fish, 8)
		}
		fmt.Printf("Day[%v/%v] = %v\n", day, days, len(fish))
	}
	return len(fish)
}

func algo(days int, data []int) {
	total := 0
	depth := 0

	cache := make(map[int]int)

	for _, value := range data {
		count := count_children(cache, depth, days-(value+1)) // 1 is this fish itself
		total += count
	}
	total += len(data)
	fmt.Printf("algo  total=%v\n", total)
}

// this day days is the creation day
func count_children(cache map[int]int, depth int, days int) int {
	if cache[days] != 0 {
		return cache[days]
	}
	if days < 0 {
		return 0
	}
	if days < 7 {
		// we KNOW we can't create a child, so we can safely return only this fish
		return 1
	}

	// so today is a count of 1 because it has created 1 fish
	// now we look at where we are and work out how many can we create from here

	// direct_spawns := (days / 7) // the +1 is "this" fish
	// c := direct_spawns

	total := 0
	for test_day := days; test_day >= 0; test_day -= 7 {
		// we spawn on this day - can we use this day + 9 as a spawning day?
		total += 1
		if test_day >= 0 {
			total += count_children(cache, depth+1, test_day-9)
		}
	}

	// fmt.Printf("[depth=%v, days=%v, returning %v]\n", depth, days, total)
	cache[days] = total
	return total

}

// rename this to the year and day in question
func (app *Application) Y2021D06P2() {
	s := strings.ReplaceAll(DAY_2021_06_DATA, " ", "")
	data := utils.SplitDataToListOfInts(s, ",")
	algo(256, data)

}

// rename and uncomment this to the year and day in question once complete for a gold star!
// func (app *Application) Y20XXDXXP1Render() {
// }

// rename and uncomment this to the year and day in question once complete for a gold star!
// func (app *Application) Y20XXDXXP2Render() {
// }

// this is what we will reflect and call - so both parts with run. It's up to you to make it print nicely etc.
// The app reference has a CLI for logging.
func (app *Application) Y2021D06() {
	app.Y2021D06P1()
	app.Y2021D06P2()
}
